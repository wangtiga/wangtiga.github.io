---
layout: post
title:  "如果 client 向 TIME WAIT 状态的 server 重传 FIN 时， server 会回复 RST 还是 ACK ？"
date:   2021-11-17 12:00:00 +0800
tags:   tech
---

* category
{:toc}



## 如果 client 向 `TIME_WAIT` 状态的 server 重传 FIN 时， server 会回复 RST 还是 ACK ？ [^TCPUnexcept] [^TCPRST]

答：server 回复 ACK 。

> 注：
> 本文示例的抓包数据中包含完整了的 TCP 连接建立和销毁过程，所以统一将先发 SYN 的端称为 client  。
> 网络其他文章一般按以下规则区分 server 和 client ：
> 1. 讨论 三次握手 建立连接的过程，一般称先发 SYN 的为 client ；
> 2. 讨论 四次握手 关闭连接的过程，一般称先发 FIN 的为 client ；


### 一 netstat 观察 tcp.port=60000 的连接状态


```sh
$ watch "sudo  netstat -n |sed -n -e 1,2p -e /60000/p"
```


### 二 tcpdump 抓包 tcp.port=60000 的所有 IP 包

```sh
$ sudo tcpdump -i any tcp port 60000 -ttttN
```


### 三 执行 tcp server 端代码，监听 60000 端口

```sh
$ python server.py
# 执行后程序会阻塞在 sock.accept() ，收到客户端 SYN 建立连接请求后，立即回复 ACK ，再回复 FIN 结束连接
```

server.py
```py
from socket import *

sock = socket(AF_INET,SOCK_STREAM)
sock.bind(('', 60000))
sock.listen(1000)
cli, addr = sock.accept()
cli.close()
```


### 四 执行 tcp client 端代码，连接 60000 端口

#### 4.1 创建 socket ，并连接 60000 端口

```sh
$ python
>>> from socket import *
>>> sock = socket(AF_INET,SOCK_STREAM)
>>> sock.connect(('localhost', 60000))
```

以上代码执行完毕后，会看到 TCP 连接处于半连接状态（client能发不能收，server 能收不能发；） TODO client send 试试
参考 log1 详细状态如下：

```sh
tiga@ubuntu:~$ sudo iptables -L INPUT
Chain INPUT (policy ACCEPT)
target     prot opt source               destination
tiga@ubuntu:~$
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

tiga@ubuntu:~$ sudo tcpdump -i any tcp port 60000 -ttttN
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on any, link-type LINUX_SLL (Linux cooked), capture size 262144 bytes
2021-09-03 16:17:09.379684 IP localhost.60600 > localhost.60000: Flags [S], seq 2048755344, win 65495, options [mss 65495,sackOK,TS val 441755898 ecr 0,nop,wscale 7], length 0
2021-09-03 16:17:09.379713 IP localhost.60000 > localhost.60600: Flags [S.], seq 2478887813, ack 2048755345, win 65483, options [mss 65495,sackOK,TS val 441755898 ecr 441755898,nop,wscale 7], length 0
2021-09-03 16:17:09.379741 IP localhost.60600 > localhost.60000: Flags [.], ack 1, win 512, options [nop,nop,TS val 441755898 ecr 441755898], length 0
2021-09-03 16:17:09.379995 IP localhost.60000 > localhost.60600: Flags [F.], seq 1, ack 1, win 512, options [nop,nop,TS val 441755898 ecr 441755898], length 0
2021-09-03 16:17:09.383583 IP localhost.60600 > localhost.60000: Flags [.], ack 2, win 512, options [nop,nop,TS val 441755902 ecr 441755898], length 0

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
Every 2.0s: sudo  netstat -n |sed -n -e 1,2p -e /60000/p                                                                                            ubuntu: Fri Sep  3 16:17:17 2021

激活Internet连接 (w/o 服务器)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
tcp        0      0 127.0.0.1:60000         127.0.0.1:60600         FIN_WAIT2
tcp        1      0 127.0.0.1:60600         127.0.0.1:60000         CLOSE_WAIT
───────────────────────────────────────────────────────────────────────────────────────────┬────────────────────────────────────────────────────────────────────────────────────────────
tiga@ubuntu:~/src/test/python/net$ cat server.py                                           │tiga@ubuntu:~$ 
from socket import *                                                                       │tiga@ubuntu:~$ python
                                                                                           │Python 2.7.17 (default, Feb 27 2021, 15:10:58) 
sock = socket(AF_INET,SOCK_STREAM)                                                         │[GCC 7.5.0] on linux2
sock.bind(('', 60000))                                                                     │Type "help", "copyright", "credits" or "license" for more information.sock.listen(1000)  
cli, addr = sock.accept()                                                                  │>>> 
cli.close()                                                                                │>>> sock = socket(AF_INET,SOCK_STREAM)
                                                                                           │>>> sock.connect(('localhost', 60000))
tiga@ubuntu:~/src/test/python/net$                                                         │>>> 
tiga@ubuntu:~/src/test/python/net$ python server.py                                        │
```

##### 4.2 server 端的 python server.py 脚本运行结束了；
##### 4.3 tcpdump 窗口显示，三次握手对应的 SYN/ACK 已经结束；server 端先发了 FIN 关闭连接，而 client 端回复了 ACK ；
##### 4.4 netstat 窗口显示，server 端连接处于 `FIN_WAIT2` 状态，等待接收 FIN ；而 client 端连接处于 `CLOSE_WAIT` 状态，等待应用层调用 close ，从而发出 FIN ；

### 五 增加 iptables 规则，在 client 发 FIN 时，丢弃 server 回复的 ACK 包

```sh
$ sudo iptables -A INPUT -p tcp -s 127.0.0.1  --sport 60000 -j DROP
```

#### 5.1 client 端关闭连接，触发 FIN 包

在刚才手动执行 python 脚本的窗口，执行一行代码即可

```sh
$ python
>>> sock.close()
```

以上代码执行完毕后，会看到 TCP 连接已经结束（client 端等待接收被 drop 的 `LAST_ACK`，server 端已经进入 `TIME_WAIT`；
参考 log2 详细状态如下：

```sh
tiga@ubuntu:~$ sudo iptables -L INPUT
Chain INPUT (policy ACCEPT)
target     prot opt source               destination
tiga@ubuntu:~$ sudo iptables -A INPUT -p tcp -s 127.0.0.1  --sport 60000 -j DROP 
tiga@ubuntu:~$ 
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
tiga@ubuntu:~$ sudo tcpdump -i any tcp port 60000 -ttttN
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on any, link-type LINUX_SLL (Linux cooked), capture size 262144 bytes
2021-09-03 16:17:09.379684 IP localhost.60600 > localhost.60000: Flags [S], seq 2048755344, win 65495, options [mss 65495,sackOK,TS val 441755898 ecr 0,nop,wscale 7], length 0
2021-09-03 16:17:09.379713 IP localhost.60000 > localhost.60600: Flags [S.], seq 2478887813, ack 2048755345, win 65483, options [mss 65495,sackOK,TS val 441755898 ecr 441755898,nop,wscale 7], length 0
2021-09-03 16:17:09.379741 IP localhost.60600 > localhost.60000: Flags [.], ack 1, win 512, options [nop,nop,TS val 441755898 ecr 441755898], length 0
2021-09-03 16:17:09.379995 IP localhost.60000 > localhost.60600: Flags [F.], seq 1, ack 1, win 512, options [nop,nop,TS val 441755898 ecr 441755898], length 0
2021-09-03 16:17:09.383583 IP localhost.60600 > localhost.60000: Flags [.], ack 2, win 512, options [nop,nop,TS val 441755902 ecr 441755898], length 0

2021-09-03 16:17:37.399376 IP localhost.60600 > localhost.60000: Flags [F.], seq 1, ack 2, win 512, options [nop,nop,TS val 441783918 ecr 441755898], length 0
2021-09-03 16:17:37.399402 IP localhost.60000 > localhost.60600: Flags [.], ack 2, win 512, options [nop,nop,TS val 441783918 ecr 441783918], length 0
2021-09-03 16:17:37.603600 IP localhost.60600 > localhost.60000: Flags [F.], seq 1, ack 2, win 512, options [nop,nop,TS val 441784122 ecr 441755898], length 0
2021-09-03 16:17:37.603620 IP localhost.60000 > localhost.60600: Flags [.], ack 2, win 512, options [nop,nop,TS val 441784122 ecr 441783918], length 0

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
Every 2.0s: sudo  netstat -n |sed -n -e 1,2p -e /60000/p                                                                                            ubuntu: Fri Sep  3 16:17:47 2021

激活Internet连接 (w/o 服务器)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
tcp        0      0 127.0.0.1:60000         127.0.0.1:60600         TIME_WAIT
tcp        1      1 127.0.0.1:60600         127.0.0.1:60000         LAST_ACK  
───────────────────────────────────────────────────────────────────────────────────────────┬────────────────────────────────────────────────────────────────────────────────────────────
tiga@ubuntu:~/src/test/python/net$ cat server.py                                           │tiga@ubuntu:~$ 
from socket import *                                                                       │tiga@ubuntu:~$ python
                                                                                           │Python 2.7.17 (default, Feb 27 2021, 15:10:58) 
sock = socket(AF_INET,SOCK_STREAM)                                                         │[GCC 7.5.0] on linux2
sock.bind(('', 60000))                                                                     │Type "help", "copyright", "credits" or "license" for more information.sock.listen(1000)  
cli, addr = sock.accept()                                                                  │>>> 
cli.close()                                                                                │>>> sock = socket(AF_INET,SOCK_STREAM)
                                                                                           │>>> sock.connect(('localhost', 60000))
tiga@ubuntu:~/src/test/python/net$                                                         │>>> sock.close()
tiga@ubuntu:~/src/test/python/net$ python server.py                                        │>>>
tiga@ubuntu:~/src/test/python/net$                                                         │
```

#### 5.2 tcpdump 窗口显示，server 端已收到 FIN 并回复 ACK ，但 ACK 被 iptables 丢弃，所以 client 没收到 ACK 就多次重传 FIN ；

NOTE 处于 `TIME_WAIT` 状态收到多次 FIN 时，都会回复 ACK ；如果连接关闭进入 CLOSED 状态再收到 FIN ，就肯定不会回复 ACK 了。这时根据系统配置，可能回复 RST 或直接丢包忽略。[^TCPRST]


#### 5.3 netstat 窗口显示，server 端已进入 `TIME_WAIT` 中， 而 client 处于 `LAST_ACK` 状态；



### 六 删除 iptables 规则，让 client 重传 FIN 时，能收到 server 回复的 ACK 包，使连接正确关闭

```sh
$ sudo iptables -D INPUT -p tcp -s 127.0.0.1  --sport 60000 -j DROP
```

耐心等待一段时间（TODO 重传规则计算时间）

参考 log3 详细状态如下：

```sh
tiga@ubuntu:~$ sudo iptables -L INPUT
Chain INPUT (policy ACCEPT)
target     prot opt source               destination
tiga@ubuntu:~$ sudo iptables -A INPUT -p tcp -s 127.0.0.1  --sport 60000 -j DROP 
tiga@ubuntu:~$ sudo iptables -D INPUT -p tcp -s 127.0.0.1  --sport 60000 -j DROP 
tiga@ubuntu:~$ 
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
tiga@ubuntu:~$ sudo tcpdump -i any tcp port 60000 -ttttN
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on any, link-type LINUX_SLL (Linux cooked), capture size 262144 bytes
2021-09-03 16:17:09.379684 IP localhost.60600 > localhost.60000: Flags [S], seq 2048755344, win 65495, options [mss 65495,sackOK,TS val 441755898 ecr 0,nop,wscale 7], length 0
2021-09-03 16:17:09.379713 IP localhost.60000 > localhost.60600: Flags [S.], seq 2478887813, ack 2048755345, win 65483, options [mss 65495,sackOK,TS val 441755898 ecr 441755898,nop,wscale 7], length 0
2021-09-03 16:17:09.379741 IP localhost.60600 > localhost.60000: Flags [.], ack 1, win 512, options [nop,nop,TS val 441755898 ecr 441755898], length 0
2021-09-03 16:17:09.379995 IP localhost.60000 > localhost.60600: Flags [F.], seq 1, ack 1, win 512, options [nop,nop,TS val 441755898 ecr 441755898], length 0
2021-09-03 16:17:09.383583 IP localhost.60600 > localhost.60000: Flags [.], ack 2, win 512, options [nop,nop,TS val 441755902 ecr 441755898], length 0

2021-09-03 16:17:37.399376 IP localhost.60600 > localhost.60000: Flags [F.], seq 1, ack 2, win 512, options [nop,nop,TS val 441783918 ecr 441755898], length 0
2021-09-03 16:17:37.399402 IP localhost.60000 > localhost.60600: Flags [.], ack 2, win 512, options [nop,nop,TS val 441783918 ecr 441783918], length 0
2021-09-03 16:17:37.603600 IP localhost.60600 > localhost.60000: Flags [F.], seq 1, ack 2, win 512, options [nop,nop,TS val 441784122 ecr 441755898], length 0
2021-09-03 16:17:37.603620 IP localhost.60000 > localhost.60600: Flags [.], ack 2, win 512, options [nop,nop,TS val 441784122 ecr 441783918], length 0

2021-09-03 16:18:27.875887 IP localhost.60600 > localhost.60000: Flags [F.], seq 1, ack 2, win 512, options [nop,nop,TS val 441834394 ecr 441755898], length 0
2021-09-03 16:18:27.875940 IP localhost.60000 > localhost.60600: Flags [.], ack 2, win 512, options [nop,nop,TS val 441834394 ecr 441783918], length 0

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
Every 2.0s: sudo  netstat -n |sed -n -e 1,2p -e /60000/p                                                                                            ubuntu: Fri Sep  3 16:18:30 2021

激活Internet连接 (w/o 服务器)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
tcp        0      0 127.0.0.1:60000         127.0.0.1:60600         TIME_WAIT

───────────────────────────────────────────────────────────────────────────────────────────┬────────────────────────────────────────────────────────────────────────────────────────────
tiga@ubuntu:~/src/test/python/net$ cat server.py                                           │tiga@ubuntu:~$ 
from socket import *                                                                       │tiga@ubuntu:~$ python
                                                                                           │Python 2.7.17 (default, Feb 27 2021, 15:10:58) 
sock = socket(AF_INET,SOCK_STREAM)                                                         │[GCC 7.5.0] on linux2
sock.bind(('', 60000))                                                                     │Type "help", "copyright", "credits" or "license" for more information.sock.listen(1000)  
cli, addr = sock.accept()                                                                  │>>> 
cli.close()                                                                                │>>> sock = socket(AF_INET,SOCK_STREAM)
                                                                                           │>>> sock.connect(('localhost', 60000))
tiga@ubuntu:~/src/test/python/net$                                                         │>>> sock.close()
tiga@ubuntu:~/src/test/python/net$ python server.py                                        │>>>
tiga@ubuntu:~/src/test/python/net$                                                         │
```

#### 6.1 tcpdump 窗口显示，server 端第三次收到 FIN 并回复 ACK ，这次 client 收到 ACK 并正确关闭连接，释放了系统资源；
#### 6.2 netstat 窗口显示，server 端已进入 `TIME_WAIT` 中， 而 client 已经进入 CLOSED 状态，释放了资源，所以看不到相关连接了；


### 相关命令集合

1. 有 linux 才能方便地用 iptables 模拟丢包；
2. 有 python 才能方便地单步执行代码；
3. 观测程序发网络包时，用 gdb 调试 c 或用 dlv 调试 go 都不如直接用 ptyhon 的交互式命令行写代码方便。

```sh
sudo iptables -L INPUT
sudo iptables -A INPUT -p tcp -s 127.0.0.1  --sport 60000 -j DROP
# sudo iptables -D INPUT -p tcp -s 127.0.0.1  --sport 60000 -j DROP
watch "sudo  netstat -n |sed -n -e 1,2p -e /60000/p" # ss -anp 命令比 netstat 更好用
sudo tcpdump -i any tcp port 60000 -ttttN
sudo sysctl -w net.ipv4.tcp_fin_timeout=60 # 设置 TIME_WAIT 时长
cat /proc/sys/net/ipv4/tcp_fin_timeout
```


[^TCPUnexcept]: [TCP 三次握手、四次挥手出现意外情况时，为保证稳定，是如何处理的？](https://segmentfault.com/a/1190000021740112)

[^TCPRST]: [TCP RST 产生的几种情况](https://zhuanlan.zhihu.com/p/30791159)
 
