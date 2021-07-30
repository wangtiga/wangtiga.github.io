---
layout: post
title:  "[译]Redis 分布式锁 redlock"
date:   2019-12-22 21:00:00 +0800
tags:   tech
---

* category
{:toc}



# Distributed locks with Redis[^RedisDistLockRedlock]


Distributed locks are a very useful primitive in many environments where different processes must operate with shared resources in a mutually exclusive way.

分布式锁用于在多个进程间，在处理抢占共享资源时，它是十分有用的编程原语。

There are a number of libraries and blog posts describing how to implement a DLM (Distributed Lock Manager) with Redis, but every library uses a different approach, and many use a simple approach with lower guarantees compared to what can be achieved with slightly more complex designs.

网上有很多 library 和博客文章描述如何使用 Redis 实现 DLM(分布式锁管理器）。
方法各不相同，有些简单，有些非常复杂。


This page is an attempt to provide a more canonical algorithm to implement distributed locks with Redis. We propose an algorithm, called  **Redlock**, which implements a DLM which we believe to be safer than the vanilla single instance approach. We hope that the community will analyze it, provide feedback, and use it as a starting point for the implementations or more complex or alternative designs.

本文尝试描述一种，使用 Redis 实现分布式锁的经典算法。
我们建议的算法称作 **Redlock** ，这种算法实现的 DLM 更安全。
希望社区中的人能分析讨论这个算法，并提供一些反馈。
也可以以此为起点实现更复杂的算法。

TODO vanilla single instance approach 如何翻译

## Implementationsa 实现

Before describing the algorithm, here are a few links to implementations already available that can be used for reference.

描述此算法前，这里有一些现成的代码可供参考。

-   [Redlock-rb](https://github.com/antirez/redlock-rb)  (Ruby implementation). There is also a  [fork of Redlock-rb](https://github.com/leandromoreira/redlock-rb)  that adds a gem for easy distribution and perhaps more.
-   [Redlock-py](https://github.com/SPSCommerce/redlock-py)  (Python implementation).
-   [Aioredlock](https://github.com/joanvila/aioredlock)  (Asyncio Python implementation).
-   [Redlock-php](https://github.com/ronnylt/redlock-php)  (PHP implementation).
-   [PHPRedisMutex](https://github.com/malkusch/lock#phpredismutex)  (further PHP implementation)
-   [cheprasov/php-redis-lock](https://github.com/cheprasov/php-redis-lock)  (PHP library for locks)
-   [Redsync](https://github.com/go-redsync/redsync)  (Go implementation).
-   [Redisson](https://github.com/mrniko/redisson)  (Java implementation).
-   [Redis::DistLock](https://github.com/sbertrang/redis-distlock)  (Perl implementation).
-   [Redlock-cpp](https://github.com/jacket-code/redlock-cpp)  (C++ implementation).
-   [Redlock-cs](https://github.com/kidfashion/redlock-cs)  (C#/.NET implementation).
-   [RedLock.net](https://github.com/samcook/RedLock.net)  (C#/.NET implementation). Includes async and lock extension support.
-   [ScarletLock](https://github.com/psibernetic/scarletlock)  (C# .NET implementation with configurable datastore)
-   [Redlock4Net](https://github.com/LiZhenNet/Redlock4Net)  (C# .NET implementation)
-   [node-redlock](https://github.com/mike-marcacci/node-redlock)  (NodeJS implementation). Includes support for lock extension.


## Safety and Liveness guarantees 安全和活性保证

> TODO liveness 直译是活跃性，但可以理解为 活性 ，需要找几个实例来强化理解


We are going to model our design with just three properties that, from our point of view, are the minimum guarantees needed to use distributed locks in an effective way.

我们认为要实现一个高效的分布式锁，至少要保证以下三点：

1.  Safety property: Mutual exclusion. At any given moment, only one client can hold a lock.
2.  Liveness property A: Deadlock free. Eventually it is always possible to acquire a lock, even if the client that locked a resource crashes or gets partitioned.
3.  Liveness property B: Fault tolerance. As long as the majority of Redis nodes are up, clients are able to acquire and release locks.

1. Safety property 安全性：互相独占。任意时间，仅有一个客户端能占有锁。
2. Liveness property A 活性A ：不会死锁。即使占有锁的客户端 crash or partitioned 都不会必须死锁。
3. Liveness property B 活性B ：容错性。只要 Redis 主节点存在，客户端就能执行 加锁 解锁 的操作。

TODO partitioned 分区， 这里是否是说 redis cluster 部署时， server 端对 key 在存储区域执行的 分裂 ？



## Why failover-based implementations are not enough 为什么灾备措施不够？

> failover 故障转移，指一套系统失效，另外一套系统能在短时间接替工作。

To understand what we want to improve, let’s analyze the current state of affairs with most Redis-based distributed lock libraries.

为了理解 redlock 算法具体改进了哪方面，我们先来分析现有的基于 Redis 实现的分布式锁。

> TODO affair

The simplest way to use Redis to lock a resource is to create a key in an instance. The key is usually created with a limited time to live, using the Redis expires feature, so that eventually it will get released (property 2 in our list). When the client needs to release the resource, it deletes the key.

最简单的方法是为每个要加锁的资源创建一个 key 。
通常这个 key 都有一个过期时间 Time To Live （TTL) ，利用 Redis expire 机制实现。
所以这种方案能保证锁肯定会被释放，即满足 Liveness property A 。
当客户端需要主动释放锁时，它可以直接把 key 删除。


Superficially this works well, but there is a problem: this is a single point of failure in our architecture. What happens if the Redis master goes down? Well, let’s add a slave! And use it if the master is unavailable. This is unfortunately not viable. By doing so we can’t implement our safety property of mutual exclusion, because Redis replication is asynchronous.

多数情况下这样能正常工作，但它存在 single point of failure （单点失效） 问题。
如果 Redis master 挂了怎么办？可以加一个 slave 从节点，在 master 不可用时接替工作。
但这样不能解决问题。因为 Redis replication 是异步的，所以此方案无法满足 Safety property 。

TODO Superficially viable

There is an obvious race condition with this model:

这个 model 存在显而易见的竟态条件：

1.  Client A acquires the lock in the master.
2.  The master crashes before the write to the key is transmitted to the slave.
3.  The slave gets promoted to master.
4.  Client B acquires the lock to the same resource A already holds a lock for.  **SAFETY VIOLATION!**

1.  Client A 在 master 中得到锁。
2.  master 在将 key 写入 slave 前挂了。
3.  slave 提升为 master
4.  Client B 在 新 master 中回锁成功。 B 与 A 同时占有相同的锁。 **SAFETY 规则被破坏!** 


Sometimes it is perfectly fine that under special circumstances, like during a failure, multiple clients can hold the lock at the same time. If this is the case, you can use your replication based solution. Otherwise we suggest to implement the solution described in this document.

如果在类似的异常情况下，允许多个 Client 同时占有锁，那么使用这种方案，问题不大。
假设你也允许发生这种情况，那么可以继续使用基于 replication 的方案吧。
否则，我们建议使用本文描述的方案。


## Correct implementation with a single instance 单点时，正确的实现方案

Before trying to overcome the limitation of the single instance setup described above, let’s check how to do it correctly in this simple case, since this is actually a viable solution in applications where a race condition from time to time is acceptable, and because locking into a single instance is the foundation we’ll use for the distributed algorithm described here.

在解决上述方案的缺陷前，我们先看看如果实现上述方案。

TODO overcome  viable colution

To acquire the lock, the way to go is the following:

使用如下命令加锁：

```redis
SET resource_name my_random_value NX PX 30000
```

The command will set the key only if it does not already exist (NX option), with an expire of 30000 milliseconds (PX option). The key is set to a value “my_random_value”. This value must be unique across all clients and all lock requests.

NX 选项，只在 key 不存在时，才会成功执行 SET 命令。
PX 选项，设置 key 的过期时间为 30000 ms ，即 30s 后自动删除 key 。
key 的值设置为 "my_randowm_value" ，即随机字符串。这个字符串每个客户端的每个请求都应该不同。


Basically the random value is used in order to release the lock in a safe way, with a script that tells Redis: remove the key only if it exists and the value stored at the key is exactly the one I expect to be. This is accomplished by the following Lua script:

设置随机字符串是为了确保只有真正占有锁的客户端才能执行解锁操作，即确保安全性。
使用下面的 Lua script 能确保：只有在 key 中存储的 value 与指定参数完全一致时，才执行删除操作。

```redis
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
```

This is important in order to avoid removing a lock that was created by another client. For example a client may acquire the lock, get blocked in some operation for longer than the lock validity time (the time at which the key will expire), and later remove the lock, that was already acquired by some other client. Using just DEL is not safe as a client may remove the lock of another client. With the above script instead every lock is “signed” with a random string, so the lock will be removed only if it is still the one that was set by the client trying to remove it.

这很重要，它能避免移除其他 client 占用的锁。
比如， clientA 加锁后，执行了一些阻塞的操作，总的执行时间超出 lockKey 的有效期后，另外一个 clientB 占有了这个锁。
这时，如果在 clientA 中直接执行 DEL 删除 lockKey 是非常危险的。
使用以上脚本删除 lockKey 来解锁，能保证只删除自己占有的锁。
因为每次加锁成功后，写入的随机字符串起到了签名的作用。



What should this random string be? I assume it’s 20 bytes from /dev/urandom, but you can find cheaper ways to make it unique enough for your tasks. For example a safe pick is to seed RC4 with /dev/urandom, and generate a pseudo random stream from that. A simpler solution is to use a combination of unix time with microseconds resolution, concatenating it with a client ID, it is not as safe, but probably up to the task in most environments.

如何生成随机字符串？
使用 clientID + unixTimeStamp + /dev/urandom 。


The time we use as the key time to live, is called the “lock validity time”. It is both the auto release time, and the time the client has in order to perform the operation required before another client may be able to acquire the lock again, without technically violating the mutual exclusion guarantee, which is only limited to a given window of time from the moment the lock is acquired.

key 的过期时间，也可称为 "锁的有效时"。
过期后，会自动释放锁，所以占有锁的客户端应该尽量在此时间内执行完毕。


So now we have a good way to acquire and release the lock. The system, reasoning about a non-distributed system composed of a single, always available, instance, is safe. Let’s extend the concept to a distributed system where we don’t have such guarantees.



## The Redlock algorithm

In the distributed version of the algorithm we assume we have N Redis masters. Those nodes are totally independent, so we don’t use replication or any other implicit coordination system. We already described how to acquire and release the lock safely in a single instance. We take for granted that the algorithm will use this method to acquire and release the lock in a single instance. In our examples we set N=5, which is a reasonable value, so we need to run 5 Redis masters on different computers or virtual machines in order to ensure that they’ll fail in a mostly independent way.

刚才的算法，适用于 single instance Redis master 的环境中，如果我们使用了 N 个 Redis master ，就要考虑一些额外的问题。


In order to acquire the lock, the client performs the following operations:

1.  It gets the current time in milliseconds.
2.  It tries to acquire the lock in all the N instances sequentially, using the same key name and random value in all the instances. During step 2, when setting the lock in each instance, the client uses a timeout which is small compared to the total lock auto-release time in order to acquire it. For example if the auto-release time is 10 seconds, the timeout could be in the ~ 5-50 milliseconds range. This prevents the client from remaining blocked for a long time trying to talk with a Redis node which is down: if an instance is not available, we should try to talk with the next instance ASAP.
3.  The client computes how much time elapsed in order to acquire the lock, by subtracting from the current time the timestamp obtained in step 1. If and only if the client was able to acquire the lock in the majority of the instances (at least 3), and the total time elapsed to acquire the lock is less than lock validity time, the lock is considered to be acquired.
4.  If the lock was acquired, its validity time is considered to be the initial validity time minus the time elapsed, as computed in step 3.
5.  If the client failed to acquire the lock for some reason (either it was not able to lock N/2+1 instances or the validity time is negative), it will try to unlock all the instances (even the instances it believed it was not able to lock).


1. 获取当前时间， 单位 millisecond 毫秒。
2. 顺序对 N 个 instance 执行加锁操作，每个 instance 使用的 key 和 random value 都是一样的。
这个过程， client timeout 应该比总的 lock auto-release time 要小。
比如，如果 auto-release 是 10 second ， timeout 应该是  5-50 ms 左右。
这样是为了减少已经停机的 Redis 对客户端加锁过程的影响：发现无效的 instance 后，我们应该尽快尝试下一个 instance 。
3. 只有在多数（至少3个） instance 中加锁成功，并且总的加锁耗费时间小于 lock auto-release 时间要时，才认为加锁操作是有效的。 client 应使用加锁成功时的时间减去 step 1 记录的时间，计算出 total time elapsed （加锁过程耗费的时间）。
4. validity time （剩余有效时间）。时间内。所以加锁成功后，剩余的有效时间不足 10 second 了，所以 lock auto-release time 减 total time elapsed 等于 validity time 。
5. 如果 client 加锁失败（比如不到 N/2 + 1 个 instance 加锁成功 或者 validity time 是负数），应该在所有 instance 上执行解锁操作（即使这个 instance 上加锁失败，也要执行解锁操作）。

> 译： redlock 算法主要优化在于，对 判断加锁成功的条件 和 计算加锁有效时间的过程 做了更细致的处理。对吗？

> TODO redis cluster 官方实现是如何支持分布式的？是客户端对 key 进行分片，还是服务端对 key 进行分片？
> 如果是服务端分片，对于客户端来说，应该就不用关注这个问题了吧。



## Is the algorithm asynchronous?

The algorithm relies on the assumption that while there is no synchronized clock across the processes, still the local time in every process flows approximately at the same rate, with an error which is small compared to the auto-release time of the lock. This assumption closely resembles a real-world computer: every computer has a local clock and we can usually rely on different computers to have a clock drift which is small.

这个算法假设每个 computer 上的 local time 差异很小，相比 auto-release time 要小很多。

At this point we need to better specify our mutual exclusion rule: it is guaranteed only as long as the client holding the lock will terminate its work within the lock validity time (as obtained in step 3), minus some time (just a few milliseconds in order to compensate for clock drift between processes).

为了补偿不同 process 的  clock 差异，计算 validity time 时，要减去一个很小的值（约几 millisecond 单位左右）。

> CLOCK_DRIFT
> 因为存在时间差，所以 lock 的实际有效时间肯定比预期要短。因为个别 computer 的时钟走的快，它的 key 就会更早过期失败，所以总有有效时间自然就短了。 

For more information about similar systems requiring a bound  _clock drift_, this paper is an interesting reference:  [Leases: an efficient fault-tolerant mechanism for distributed file cache consistency](http://dl.acm.org/citation.cfm?id=74870).



## Retry on failure

When a client is unable to acquire the lock, it should try again after a random delay in order to try to desynchronize multiple clients trying to acquire the lock for the same resource at the same time (this may result in a split brain condition where nobody wins). Also the faster a client tries to acquire the lock in the majority of Redis instances, the smaller the window for a split brain condition (and the need for a retry), so ideally the client should try to send the SET commands to the N instances at the same time using multiplexing.

当客户加锁失败时，应该等待随机的一段时间后，再重试。这样能防止在同一时间点，有很多客户端同时对一个资源执行加锁操作（这可能导致出现 split brain 脑裂的后果）。
客户端越快在 Redis majority 占有锁，发生 split brain 的时间窗口也越短（重试的次数自然也越少），
最理想的情况是，客户端同时向 N 个 instance 发送 SET 命令。

TODO 了解 split brain 的预防手段

It is worth stressing how important it is for clients that fail to acquire the majority of locks, to release the (partially) acquired locks ASAP, so that there is no need to wait for key expiry in order for the lock to be acquired again (however if a network partition happens and the client is no longer able to communicate with the Redis instances, there is an availability penalty to pay as it waits for key expiration).

有必要再次强调，因为实际环境中，占多数的是加锁失败的客户端，所以占有锁的客户端要尽量释放锁，这样其他端就不必等待锁超时过期后才能加锁了。
发生网络分区时，客户端无法与 Redis instance 通信，所以在等待锁超时过期这段时间，要承担服务不可用的后果。



## Releasing the lock

Releasing the lock is simple and involves just releasing the lock in all instances, whether or not the client believes it was able to successfully lock a given instance.

解锁的过程很简单，但是要在所有 instance 中执行解锁的操作，无论客户端是否在指定 instance 上加锁成功。

TODO 为什么要在 all instance 中执行解锁？避免 客户端认为自己没加锁，但服务端认为客户端加锁成功的情况出现？同时给所有 instance 发解锁的命令，不浪费资源吗？虽然会有些 instance 没加锁成功，但也会收到解锁命令，这些命令其实可以不发。但为可用性，这点浪费是值得的。



## Safety arguments 有关安全性的论述

Is the algorithm safe? We can try to understand what happens in different scenarios.

这个算法安全吗？我们可以尝试分析下两种不同场景。

To start let’s assume that a client is able to acquire the lock in the majority of instances. All the instances will contain a key with the same time to live. However, the key was set at different times, so the keys will also expire at different times. But if the first key was set at worst at time T1 (the time we sample before contacting the first server) and the last key was set at worst at time T2 (the time we obtained the reply from the last server), we are sure that the first key to expire in the set will exist for at least  `MIN_VALIDITY=TTL-(T2-T1)-CLOCK_DRIFT`. All the other keys will expire later, so we are sure that the keys will be simultaneously set for at least this time.

client 在  majority of instance 加锁成功后。
看似所有 instance 中的 key 都有相同的  time to live 。
实际上这个时间是有些差异的，所以它们的过期时间也不是同一时间点。
假设在第一个 instance 加锁成功的时间点是 T1 （这个时间点是在发送请求到 instance 之前记录的）。
而最近一个 instance 加锁成功的时间点是 T2 （这个时间点是在接收到 instance 响应之后记录的）。

我们能肯定，在第一个 instance 中的 key 最快会在`MIN_VALIDITY=TTL-(T2-T1)-CLOCK_DRIFT`后过期。
然后其他 instance 中的 key 也很快相继过期。


During the time that the majority of keys are set, another client will not be able to acquire the lock, since N/2+1 SET NX operations can’t succeed if N/2+1 keys already exist. So if a lock was acquired, it is not possible to re-acquire it at the same time (violating the mutual exclusion property).

在设置了 majority of key 后，其他 client 无法 acquire lock ，因为有 N/2+1 个 key 已经存在，所以肯定会有 N/2+1 个 SET NX 操作失败。
因为 lock 成功后，不会再次 acquire lock 成功。


However we want to also make sure that multiple clients trying to acquire the lock at the same time can’t simultaneously succeed.

但我们还要保证多个 client 同时 acquire lock 时，只有一个 client 成功，其他 client 必须失败。


If a client locked the majority of instances using a time near, or greater, than the lock maximum validity time (the TTL we use for SET basically), it will consider the lock invalid and will unlock the instances, 
so we only need to consider the case where a client was able to lock the majority of instances in a time which is less than the validity time. 
In this case for the argument already expressed above, for  `MIN_VALIDITY`  no client should be able to re-acquire the lock. 
So multiple clients will be able to lock N/2+1 instances at the same time (with "time" being the end of Step 2) only when the time to lock the majority was greater than the TTL time, making the lock invalid.

如果一个 client 在 lock the majority of instance 的过程就耗费接近或超过 maxium validity time （即 SET 命令的 TTL) ，它会认为这次 lock 失败，并 unlock 所有 instance 。
因此，我们只需要考虑 client lock the majority of instance 耗时少于 validity time 的情况。这种情况下，任何 client 都无法 re-acquire the lock 。
所以，只有在 lock the majority of instance 过程超过 TTL time 时，才会出现多个 client 同时 lock N/2+1 个 instance  的情况。但因为超过 TTL ，所以即使出现这种情况也会被认为 lock 失败。

Are you able to provide a formal proof of safety, point to existing algorithms that are similar, or find a bug? That would be greatly appreciated.

你能找到达到同样安全性的类似算法吗？或者你能指出这个算法的 bug 吗？ 欢迎提出你的看法。



## Liveness arguments 活性争论（参数）

> Discovery and failure detection = liveness  发现+故障检测=活性


The system liveness is based on three main features:

系统活性基于以下三个功能：

1.  The auto release of the lock (since keys expire): eventually keys are available again to be locked.
2.  The fact that clients, usually, will cooperate removing the locks when the lock was not acquired, or when the lock was acquired and the work terminated, making it likely that we don’t have to wait for keys to expire to re-acquire the lock.
3.  The fact that when a client needs to retry a lock, it waits a time which is comparably greater than the time needed to acquire the majority of locks, in order to probabilistically make split brain conditions during resource contention unlikely.

1. 超时后自动解锁：保证最终可用性。
2. 实际上，很少出现要等待 key 过期才能加锁的情况。因为通常 client 加锁成功时，会在使用完毕后主动解锁，而 client 加锁失败时，也会主动移除锁（移除的是小于 N/2+1 个 instance 上的锁）。
3. 为减小资料竞争时出现 split brain （脑裂）的概念，在 client  retry lock （重试加锁）的情况，应该等待一段时间。等待的时间长度应该稍大于 acquire the majority of lock 的时间（即等待超过加锁成功所需要的时间）。



However, we pay an availability penalty equal to  [TTL](https://redis.io/commands/ttl)  time on network partitions, so if there are continuous partitions, we can pay this penalty indefinitely. This happens every time a client acquires a lock and gets partitioned away before being able to remove the lock.

但是，如果发生 network partition （网络分区），就会持续 TTL 时间不可用状态。
如果一直处于 partition 状态 ，那就一直不可用。
只要在 client 加锁成功之后，解锁之前发生 partition ，就会出现不可用的情况。


Basically if there are infinite continuous network partitions, the system may become not available for an infinite amount of time.

简单说，如果一直处于 network partition 状态 ，那就一直不可用。



## Performance, crash-recovery and fsync 性能，故障恢复与同步

Many users using Redis as a lock server need high performance in terms of both latency to acquire and release a lock, and number of acquire / release operations that it is possible to perform per second. In order to meet this requirement, the strategy to talk with the N Redis servers to reduce latency is definitely multiplexing (or poor man's multiplexing, which is, putting the socket in non-blocking mode, send all the commands, and read all the commands later, assuming that the RTT between the client and each instance is similar).

使用 Redis 作为 lock server 的用户都关心它的恨不能，即 acquire/release lock 的耗时与每秒执行 acquire/release lock 操作的数量（tps)。
为满足这种要求，与 N 个 Redis server 同时通讯时，提高性能显而易见的方法就是 multiplexing （假设 client 与每个 instance 之间的 RTT 相同， pool man multiplexing 简易版的多路复用就是 设置 socket 为 non-blocking 模式，然后 send all command ，然后 read all command ）

However there is another consideration to do about persistence if we want to target a crash-recovery system model.

但，如果考虑故障恢复的话，就要考虑一下持久化的问题了。


Basically to see the problem here, let’s assume we configure Redis without persistence at all. A client acquires the lock in 3 of 5 instances. One of the instances where the client was able to acquire the lock is restarted, at this point there are again 3 instances that we can lock for the same resource, and another client can lock it again, violating the safety property of exclusivity of lock.

问题简要描述如下。
假设 Redis 没有开启持久化。
一个 client 在 五分之三个 instance 中 acquire lock 成功。
随后，三个 instance 中有一个重启了。
由于没有开启持久化，只剩两个 instance 保存了 client 加锁成功的状态。
也就是说，如果有另外一个 client acquire lock  资源，刚好能在 2+1=3 个 instance 中 acquire lock 成功，这就破坏 exclusivity lock （排它锁） 的 safety property （安全性）。


If we enable AOF persistence, things will improve quite a bit. For example we can upgrade a server by sending SHUTDOWN and restarting it. Because Redis expires are semantically implemented so that virtually the time still elapses when the server is off, all our requirements are fine. 
However everything is fine as long as it is a clean shutdown. What about a power outage? If Redis is configured, as by default, to fsync on disk every second, it is possible that after a restart our key is missing. In theory, if we want to guarantee the lock safety in the face of any kind of instance restart, we need to enable fsync=always in the persistence setting. This in turn will totally ruin performances to the same level of CP systems that are traditionally used to implement distributed locks in a safe way.

如果启用 AOF 持久化功能，事情就化复杂了。
我们能发送 SHUTDOWN 命令重启服务器升级。
Redis 是 sematically implement 的 expire  命令，重启后也没影响。
正常关闭没有影响，那么断电呢？如果配置 Redis 每秒写入磁盘保持数据同步，很可能出现重启后丢失 key 的情况。
理论上，如果想在任意情况重启 Redis 都能保证 lock safety ，必须在持久化设置中启用  sync=always 。
想在 same level of CP system 中保证分布式锁的 safe 特性，就必须牺牲性能。



> Redis sematically implement expire ，Redis 只在执行 get/set 等操作实际用到相关 key 时，才会检查其 expire 时间，所以重启 Redis 的过程不影响 expire 功能。能保证，过期后的 key 肯定不能使用，没过期的 key 肯定能使用。

> TODO 加锁后关闭 RedisServer instance ，解锁后 启动这个 RedisServer instance 。启动时还没超过 auto release time 。如果偶尔一两个 instance 出现类似故意，影响不大，如果超过 2/N 个 instance 出现此现象，影响大吗？目前能想到的，只 影响新的 client require lock ，即必须过了 auto release time 后，才能新 client 能正常加锁。


However things are better than what they look like at a first glance. Basically the algorithm safety is retained as long as when an instance restarts after a crash, it no longer participates to any  **currently active**  lock, so that the set of currently active locks when the instance restarts, were all obtained by locking instances other than the one which is rejoining the system.

其实也没有初看起来那么差。
当 Redis instance A 从  crash 中重启后，还是能保证算法的 safety 特性的。
因为重启的系统 instance A 不会参与所有 **当前活跃的** lock ，这些活跃的 lock 肯定只会从除 instance A 之外的其他 instanck 中获取。

To guarantee this we just need to make an instance, after a crash, unavailable for at least a bit more than the max  [TTL](https://redis.io/commands/ttl)  we use, which is, the time needed for all the keys about the locks that existed when the instance crashed, to become invalid and be automatically released.

为保证 safety ，只要让 crash 后的 Redis instance 等待超过 max TTL 时间后再加入 system 中工作。这样就能保证所有 crash 前的 key 都会自动 release 。


Using  _delayed restarts_  it is basically possible to achieve safety even without any kind of Redis persistence available, however note that this may translate into an availability penalty. For example if a majority of instances crash, the system will become globally unavailable for  [TTL](https://redis.io/commands/ttl)  (here globally means that no resource at all will be lockable during this time).

使用 _delayed restarts_ 就能保证任意 Redis persistence 类型的 safety ，但要注意，这也会有其他缺点，即丢失了一些可用性。
如果超过 N/2+1 数量个 instance crash 了，那么系统就会出现 TTL 时间的全局不可用状态。（这里的 globally 表示在此 TTL 时间内，所有 key 都不能执行 lock 操作。）



## Making the algorithm more reliable: Extending the lock 让这个算法更加可用：扩展这个 lock

> 译：下文主要内容是，减少默认加锁时间（lock validity time），通过多次延长过期时间（extend expire time），达到减少故障恢复所需要时间的目的，也即增大可用性。

If the work performed by clients is composed of small steps, it is possible to use smaller lock validity times by default, and extend the algorithm implementing a lock extension mechanism. Basically the client, if in the middle of the computation while the lock validity is approaching a low value, may extend the lock by sending a Lua script to all the instances that extends the TTL of the key if the key exists and its value is still the random value the client assigned when the lock was acquired.

如果 client 加锁成功后执行的操作可以划分成多个 step 执行，那么加锁算法还有优化的余地。
首先，要减小默认的 lock validity time 。
然后，在 client 加锁成功后，即将达到 validity 快要过期的时间点，执行延长 key 过期时间的操作。
要确保 key 存在，并且 random value 一致才能执行延期操作，所以可给所有 instance 发 lua script 实现。

The client should only consider the lock re-acquired if it was able to extend the lock into the majority of instances, and within the validity time (basically the algorithm to use is very similar to the one used when acquiring the lock).

只有 client 在 validity time 之内执行延时操作，并且在 majority of instance 上执行延期操作成功后（超过 N/2+1 个 instance ），才认为 lock re-acquired 成功。

> TODO 如果仅在少于 N/2+1 个 instance 上延时成功，是不是还要对所有 instance 执行 release lock 的操作。这个逻辑有点复杂。


However this does not technically change the algorithm, so the maximum number of lock reacquisition attempts should be limited, otherwise one of the liveness properties is violated.

但是这个优化并没有从本质上改变算法，所以应该限制延期操作的最大执行次数，否则会破坏 liveness property 。


## Want to help?

If you are into distributed systems, it would be great to have your opinion / analysis. Also reference implementations in other languages could be great.

Thanks in advance!

## Analysis of Redlock

1.  Martin Kleppmann  [analyzed Redlock here](http://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html). I disagree with the analysis and posted  [my reply to his analysis here](http://antirez.com/news/101).



> ## 说明
> 使用分布式锁之前，应该想清楚是否真得需要分布式锁。
> 用这种锁的代价并不小，如果是没想清楚就用了这种锁，就会跟我一样走弯路。
> 
> ### 1.不需要用分布式锁，内存锁就足够了
>   之所以要用分布式锁，是因为会对公共数据执行并行的 读写 操作。
>   执行读写操作的请求可能会分发到不同的进程节点处理。但实际上，有必要让这些请求放到不同进程中处理吗？
>   可以使用一种路由机制，让所有请求都路由到同一个进程节点。这样读写操作就都在同一个进程，也就不需要分布式锁了。
>   codis作者描述：`lua脚本里涉及操作多个key，Codis能做的就是将这个脚本分配到参数列表中的第一个key的机器上执行。所以这种场景下，你需要自己保证你的脚本所用到的key分布在同一个机器上，这里可以采用hashtag的方式。` [^CodisDeveloper]
>   我猜测这种限制的原因跟我上面的想法类似。
>
> ### 2.不需要用锁，原子操作就足够了
>   如果只对数据做简单的修改，没有复杂的状态判断，也没有数据追加。只要保证原子写就够用。
>   也就是，我们只需要保证最终一致即可。
>   如果很修改请求都是同一个 client ，只要让 client 保证按顺序发出修改请求，也是不需要加锁的。顺序发出请求的意思是：收到上一交请求的响应后，再发出下一次请求。
>   
> ### 3.如果加锁时间过长，是否应该考虑使用队列来处理任务
>   加锁时间过长时，其效果跟使用队列没有任何区别。
>   个人主观认为，加锁超过 5s 都算长。当然，具体业务情况，要具体讨论。
> 
> 进程 crash 后，销毁锁的方法：
> - 1.超时自动清理
>   redis key Time To Live
> - 2.每个加锁的进程维护一个 session 
>   zookeeper ephemeral nodes
> 

[^RedisDistLockRedlock]:[分布式锁 redlock](https://redis.io/topics/distlock)

[^CodisDeveloper]:[Codis作者黄东旭：细说分布式Redis架构设计和那些踩过的坑](https://dbaplus.cn/news-141-270-1.html)

