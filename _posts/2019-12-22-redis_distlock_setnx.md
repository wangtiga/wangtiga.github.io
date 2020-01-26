---
layout: post
title:  "[译]Redis 分布式锁 setnx"
date:   2019-12-22 21:00:00 +0800
tags:   tech
---

* category
{:toc}



# SETNX  key  value [^RedisDistLockSetnx]


**Available since 1.0.0.**

**Time complexity:**  O(1)

Set  `key`  to hold string  `value`  if  `key`  does not exist. In that case, it is equal to  [SET](https://redis.io/commands/set). When  `key`  already holds a value, no operation is performed.  [SETNX](https://redis.io/commands/setnx)  is short for "**SET**  if  **N**ot e**X**ists".

当 key 不存在时，才设置 value 到 key 上。
如果 key 已经存在，什么也不干。


## Return value

[Integer reply](https://redis.io/topics/protocol#integer-reply), specifically:

-   `1`  if the key was set
-   `0`  if the key was not set

## Examples

redis> SETNX mykey "Hello"

(integer) 1

redis> SETNX mykey "World"

(integer) 0

redis> GET mykey

"Hello"

redis>

## Design pattern: Locking with  `SETNX`

**Please note that:**

1. The following pattern is discouraged in favor of  [the Redlock algorithm](http://redis.io/topics/distlock)  which is only a bit more complex to implement, but offers better guarantees and is fault tolerant.
2. We document the old pattern anyway because certain existing implementations link to this page as a reference. Moreover it is an interesting example of how Redis commands can be used in order to mount programming primitives.
3. Anyway even assuming a single-instance locking primitive, starting with 2.6.12 it is possible to create a much simpler locking primitive, equivalent to the one discussed here, using the  [SET](https://redis.io/commands/set)  command to acquire the lock, and a simple Lua script to release the lock. The pattern is documented in the  [SET](https://redis.io/commands/set)  command page.

1. 建议使用 Redlock 算法，因为 Redlock 只比下面的方法复杂一点点，但容错性更好。
2. 本文还描述这个旧的方案，是因为已经有很多链接引用了当前网址。另外，本文也可以作为一个示例，演示如何使用 Redis 命令实现编程原语的过程。
3. 假定只是想实现 single-instance (单例锁) 原语，2.6.12 版本后可用 SET 命令，用更简单的方法实现与本文等效的功能。使用 SET 命令加锁，使用 Lua 脚本解锁。详细文档参考 [SET](https://redis.io/commands/set)。.

That said,  [SETNX](https://redis.io/commands/setnx)  can be used, and was historically used, as a locking primitive. For example, to acquire the lock of the key  `foo`, the client could try the following:

所以说，用 SETNX 实现 锁原语，纯粹是由于历史原因。

下面是实现加锁的示例代码：
```
SETNX lock.foo <current Unix time + lock timeout + 1>

```

If  [SETNX](https://redis.io/commands/setnx)  returns  `1`  the client acquired the lock, setting the  `lock.foo`  key to the Unix time at which the lock should no longer be considered valid. The client will later use  `DEL lock.foo`  in order to release the lock.

如果 SETNX 返回 1 ，表示客户端 加锁成功，并把 lock.foo 键的 value 设置为 加锁失效时间点的时间戳。
解锁时，执行 DEL lock.foo 命令。

If  [SETNX](https://redis.io/commands/setnx)  returns  `0`  the key is already locked by some other client. We can either return to the caller if it's a non blocking lock, or enter a loop retrying to hold the lock until we succeed or some kind of timeout expires.

如果 SETNX 返回 0 ，表示加锁失败，锁已经由其他客户端持有。

### Handling deadlocks

In the above locking algorithm there is a problem: what happens if a client fails, crashes, or is otherwise not able to release the lock? It's possible to detect this condition because the lock key contains a UNIX timestamp. If such a timestamp is equal to the current Unix time the lock is no longer valid.

以上加锁算法有个问题：如果客户端因为 crash 或其他原因导致无法释放锁怎么办？
可以通过存储在 key 中的时间戳来检测这种情况。
如果当前时间超过 value 的时间，就认为锁已经超时失效。


When this happens we can't just call  [DEL](https://redis.io/commands/del)  against the key to remove the lock and then try to issue a  [SETNX](https://redis.io/commands/setnx), as there is a race condition here, when multiple clients detected an expired lock and are trying to release it.

这时，不能简单的顺序调用 DEL SETNX 来重新加锁。
当多个客户端同时检测到过期锁并释放过期锁时，会出现竞态条件。

以下描述了 C1 C2 同时顺序调用 DEL SETNX 导致两端都认为自己占有锁了，但这是错误的。

-   C1 and C2 read  `lock.foo`  to check the timestamp, because they both received  `0`  after executing  [SETNX](https://redis.io/commands/setnx), as the lock is still held by C3 that crashed after holding the lock.
-   C1 sends  `DEL lock.foo`
-   C1 sends  `SETNX lock.foo`  and it succeeds
-   C2 sends  `DEL lock.foo`
-   C2 sends  `SETNX lock.foo`  and it succeeds
-   **ERROR**: both C1 and C2 acquired the lock because of the race condition.

Fortunately, it's possible to avoid this issue using the following algorithm. Let's see how C4, our sane client, uses the good algorithm:

使用下面的算法能避免上述问题。

-   C4 使用 SETNX 尝试加锁
-   加锁失败，使用 GET 检查锁是否超时
-   锁超时，使用 GETSET 尝试加锁
-   假设 C5 先于 C4 使用 GETSET 占有锁，那么 C4 从 GETSET 返回一个非过期时间，所以能检测到自己加锁失败了。
-   此时 C4 可以从头开始，使用 SETNX 尝试加锁，重复刚才的过程。

-   C4 sends  `SETNX lock.foo`  in order to acquire the lock
    
-   The crashed client C3 still holds it, so Redis will reply with  `0`  to C4.
    
-   C4 sends  `GET lock.foo`  to check if the lock expired. If it is not, it will sleep for some time and retry from the start.
    
-   Instead, if the lock is expired because the Unix time at  `lock.foo`  is older than the current Unix time, C4 tries to perform:
    
    ```
    GETSET lock.foo <current Unix timestamp + lock timeout + 1>
    
    ```
    
-   Because of the  [GETSET](https://redis.io/commands/getset)  semantic, C4 can check if the old value stored at  `key`  is still an expired timestamp. If it is, the lock was acquired.
    
-   If another client, for instance C5, was faster than C4 and acquired the lock with the  [GETSET](https://redis.io/commands/getset)  operation, the C4  [GETSET](https://redis.io/commands/getset)  operation will return a non expired timestamp. C4 will simply restart from the first step. Note that even if C4 set the key a bit a few seconds in the future this is not a problem.
    

In order to make this locking algorithm more robust, a client holding a lock should always check the timeout didn't expire before unlocking the key with  [DEL](https://redis.io/commands/del)  because client failures can be complex, not just crashing but also blocking a lot of time against some operations and trying to issue  [DEL](https://redis.io/commands/del)  after a lot of time (when the LOCK is already held by another client).

为了让加锁算法更健壮，持有锁的客户端应该在解锁前检查锁是否超时后再执行 DEL 命令。
因为客户端会发生各种异常情况，比如，占有锁的客户端，执行某些操作耗费很久，导致自己持有锁的时间已经超过 SETNX 时设置的 expire 时间。

[^RedisDistLockSetnx]:[SETNX  key  value](https://redis.io/commands/setnx)

