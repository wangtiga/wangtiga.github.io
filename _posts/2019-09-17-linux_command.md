---
layout: post
title:  "Linux 命令记录"
date:   2019-09-17 18:17:00 +0800
tags:  linux
---

* category
{:toc}



















### docker

```shell

docker pull golang

docker run --rm -it --name go-http-demo golang bash

sudo docker image build --build-arg UID=$(id -u) --build-arg GID=$(id -g)  -t koa-demo:0.0.1 .

sudo docker run --rm -tiv `pwd`:/go koa-demo:0.0.1

sudo docker image build --build-arg UID=$(id -u) --build-arg GID=$(id -g) \
  -f bb.dockerfile -t testimg .

```

https://github.com/mbrt/go-docker-dev

http://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html

https://docs.docker.com/get-started/
















### shell

#### aria2 [https://zhuanlan.zhihu.com/p/30666881](https://zhuanlan.zhihu.com/p/30666881)

```shell
# start aria2c
aria2c --enable-rpc --rpc-listen-all

# webui aria2c
node node-server.js
```


#### rsync

目标主机的 ssh 服务在 8022 端口上

```shell
rsync -avz -e "ssh -p 2232" SRC/ user@remote.host:/DEST/ 
```

脚本

```shell
#!/bin/sh

set -x


echo ""
echo ""
echo "I will upload bserver to targethost, and restart it!" 
echo "  usage: deplay.sh hostitem"
echo ""

for hostitem in $@
do
        case $hostitem in
                149bus)
                        SOURCEDIR=/home/tiga/src/bserver/cmd/bin
                        TARGETDIR=/app/bserver/bin/
                        TARGETHOST=tiga@192.168.0.101
                        /bin/ls -alh $SOURCEDIR
                        rsync -azv $SOURCEDIR $TARGETHOST:$TARGETDIR
                        ssh $TARGETHOST sudo supervisorctl restart bserver 
                        ;;
                *)
                        echo "Sorry, I don't understand"
                        exit 1
                        ;;
        esac
done

echo 
echo "That's all done!"
```

#### gstreamer

- install

  ```shell
  # install gstreamer pkg-config
  # https://gstreamer.freedesktop.org/documentation/installing/on-linux.html?gi-language=c
  # https://github.com/notedit/media-server-go-demo
  # brew install gst-plugins-ugly to use x264enc
  # gstreamer1.0.plugins-bad gstreamer1.0.plugins-good 的区别参考
  # sudo apt-get install libgstreamer1.0-0 gstreamer1.0-plugins-base gstreamer1.0-libav gstreamer1.0-plugins-bad libgstreamer-plugins-bad1.0-dev
  # sudo apt-get install libgstreamer1.0-0 gstreamer1.0-plugins-base gstreamer1.0-libav gstreamer1.0-plugins-good libgstreamer-plugins-good1.0-dev
  ```

- video
  ```shell
  printf "file '%s'\n" *.mp4 > mylist.txt && ffmpeg -f concat -safe 0 -i mylist.txt -c copy `basename "$(pwd)"`.mp4
  ```

  ```shell
  # mkv to rtp h264 video
  gst-launch-1.0 filesrc location="/Users/tiga/tool/resource/sintel_trailer-480p.mkv"  ! decodebin ! \
      videoconvert ! x264enc !  \
      rtph264pay ssrc=3494657 mtu=1400 config-interval=-1 ! \
      udpsink  host=127.0.0.1 port=5000
  
  # videotestsrc rtp h264 video
  gst-launch-1.0 videotestsrc ! \
      videoconvert ! x264enc !  \
      rtph264pay ssrc=3494657 mtu=1400 config-interval=-1 ! \
      udpsink  host=127.0.0.1 port=5000
  
  
  # play rtp h264
  gst-launch-1.0 -v udpsrc port=5000 address=127.0.0.1 caps="application/x-rtp,media=(string)video,clock-rate=(int)90000,encoding-name=(string)H264,payload=(int)96" ! rtph264depay ! decodebin ! videoconvert ! autovideosink
  
  
  
  # mkv to rtp vp8 video
  gst-launch-1.0 filesrc location="/Users/tiga/tool/resource/sintel_trailer-480p.mkv"  ! decodebin ! \
      videoscale ! video/x-raw, width=320, height=240 ! queue ! \
      vp8enc error-resilient=partitions keyframe-max-dist=10 auto-alt-ref=true cpu-used=5 deadline=1 ! \
      rtpvp8pay ssrc=3494657 mtu=1400 ! \
      udpsink  host=127.0.0.1 port=5000
  
  # play rtp vp8
  gst-launch-1.0 -v udpsrc port=5000 address=127.0.0.1 caps="application/x-rtp,media=(string)video,clock-rate=(int)90000,encoding-name=(string)VP8,payload=(int)96" ! rtpvp8depay ! decodebin ! videoconvert ! autovideosink
  ```

- audio

  ```shell
  # mkv to rtp opus audio
  gst-launch-1.0 filesrc location="/Users/tiga/tool/resource/sintel_trailer-480p.mkv"  ! decodebin ! \
      queue ! audioconvert  ! \
      opusenc bitrate=20000 ! \
      rtpopuspay ssrc=3494656 mtu=1400 pt=122 ! \
      udpsink host=127.0.0.1 port=5001
  
  # play rtp opus
  gst-launch-1.0 -v udpsrc port=5001 address=127.0.0.1 caps="application/x-rtp, encoding-name=OPUS, payload=96" ! rtpopusdepay ! decodebin ! autoaudiosink
  ```

- pcap

  ```shell
  # pcap to rtp 
  gst-launch-1.0 -v filesrc location=test11.pcap ! pcapparse dst-port=37904 ! application/x-rtp,payload=96 ! udpsink host=172.18.202.22 port=%d
  
  # play rtp h264 from pcap file
  gst-launch-1.0 filesrc location=264ffmpeg.pcap ! pcapparse caps="application/x-rtp,media=(string)video,clock-rate=(int)90000,encoding-name=(string)H264,payload=(int)96" ! \
      rtph264depay ! decodebin ! videoconvert ! autovideosink
  
  # play rtp opus from pcap file
  gst-launch-1.0 filesrc location="srcport_5000_rtptype_122.pcap" ! pcapparse caps="application/x-rtp,encoding-name=(string)OPUS,payload=(int)122" ! rtpopusdepay ! decodebin ! audioconvert ! audioresample ! autoaudiosink
  ```


#### protoc

```shell
$ protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/test.proto
```

https://developers.google.com/protocol-buffers/docs/gotutorial


#### influxdb

```shell
$ influx -precision rfc3339
Visit https://enterprise.influxdata.com to register for updates, InfluxDB server management, and monitoring.
Connected to http://localhost:8086 version 1.1.1
InfluxDB shell version: 1.1.1
> use mydb
Using database mydb
> show measurements
name: measurements
name
----
tab_monitor_cpu
tab_monitor_net

> drop measurement tab_monitor_cpu
> drop measurement tab_monitor_net
```


#### Mac OS 命令行中复制粘贴 clipboard

http://sweetme.at/2013/11/17/copy-to-and-paste-from-the-clipboard-on-the-mac-osx-command-line/

```sh
% pwd | pbcopy
% pbpaste
/Users/tiga

% curl -i -X GET "http://qq.com" | pbcopy
% pbpaste
HTTP/1.1 302 Moved Temporarily
Server: stgw/1.3.12.4_1.13.5
Date: Fri, 05 Nov 2021 06:53:50 GMT
Content-Type: text/html
Content-Length: 169
Connection: keep-alive
Location: https://www.qq.com/

<html>
<head><title>302 Found</title></head>
<body bgcolor="white">
<center><h1>302 Found</h1></center>
<hr><center>stgw/1.3.12.4_1.13.5</center>
</body>
</html>
```

#### Mac OS 查看当前监听的端口

https://stackoverflow.com/questions/4421633/who-is-listening-on-a-given-tcp-port-on-mac-os-x

```sh
% netstat -an | grep LISTEN
tcp6       0      0  *.6379                 *.*                    LISTEN 
tcp4       0      0  127.0.0.1.3306         *.*                    LISTEN 

% sudo lsof -iTCP -sTCP:LISTEN -n -P
Password:
COMMAND     PID USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
redis-ser 53189   ws    6u  IPv4 0xf0011292d78bdb4f      0t0  TCP *:6379 (LISTEN)
redis-ser 53189   ws    7u  IPv6 0xf0011292d7cb8f9f      0t0  TCP *:6379 (LISTEN)
transform 53904   ws    9u  IPv4 0xf0011292d750778f      0t0  TCP 127.0.0.1:8081 (LISTEN)
shorturl  53967   ws   12u  IPv6 0xf0011292d7cb70ff      0t0  TCP *:8888 (LISTEN)
```


#### Mac OS 修改 PATH

https://scriptingosx.com/2017/05/where-paths-come-from/

1. vi /etc/paths 文件添加 /xx/xx/bin 目录
2. eval `/usr/libexec/path_helper -s` 更新 PATH
3. source /etc/profile  更新 PATH

```sh
% cat /etc/paths
/usr/local/bin
/usr/bin
/bin
/usr/sbin
/sbin
```



#### Mac OS 防止耳机播放键打开 music 应用

https://www.zhihu.com/question/38813017/answer/146715121

```sh
# 关闭耳机暂停键打开 itunes music 
sudo launchctl unload -w /System/Library/LaunchAgents/com.apple.rcd.plist

# 打开 
sudo launchctl load -w /System/Library/LaunchAgents/com.apple.rcd.plist
```


#### Mac OS 查看占用端口的程序

https://tonydeng.github.io/2016/07/07/use-lsof-to-replace-netstat/

```shell
sudo lsof -nP -i  :443
```
- -n 表示不显示主机名
- -P 表示不显示端口俗称
- 不加 sudo 只能查看以当前用户运行的程序

#### Mac OS brew 换源

```shell
git -C "$(brew --repo)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git
git -C "$(brew --repo homebrew/core)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git
```

参考 https://sspai.com/post/56009


#### 更改 Mac OS history 最大历史命令记录数

```shell
# history size 
export HISTFILESIZE=1000000 
export HISTSIZE=1000000 
history 0  # 显示所有历史命令
```

参考  https://apple.stackexchange.com/questions/246621/cant-increase-mac-osx-bash-shell-history-length



#### ssh

- use public key autologin remote host, ref [ruanyifeng](https://www.ruanyifeng.com/blog/2011/12/ssh_remote_login.html)

```shell
ssh wangtiga@hostip 'mkdir -p .ssh && cat >> .ssh/authorized_keys' < ~/.ssh/id_rsa.pub "
```

- enable

```shell
$ sudo vi /etc/ssh/sshd_config
PubkeyAuthentication yes

# The default is to check both .ssh/authorized_keys and .ssh/authorized_keys2
# but this is overridden so installations will only check .ssh/authorized_keys
AuthorizedKeysFile .ssh/authorized_keys


$ sudo systemctl reload sshd
```

- faq

```shell
$  sudo tail -f /var/log/secure
Jul 1 11:33:23 vm-19 sshd[123456]: Authentication refused: bad ownership or modes for file /home/tiga/.ssh/authorized_keys

$ chmod 700 ~/.ssh
$ chmod 700 ~/.ssh/*
$ chmod 600 ~/.ssh/authorized_keys
```

#### expect

把以下文件保存到 ~/tool/bin/sshdomain ，然后再把 ~/tool/bin 添加到 /etc/paths 中，
重启 terminal 后， Mac 中就能直接输入 sshdomain 登录到 ssh 后台了


```sh
#!/usr/bin/expect

# https://segmentfault.com/a/1190000019464936

# 设置超时时间，单位秒
set timeout 10

# 主要功能是给ssh运行进程加个壳，用来传递交互指令
# ssh -A 是转发密钥设置，用于有堡垒机的场景，一般情况下不需要
spawn ssh -A  yourname@yourdomain.com

# 判断上次输出结果里是否包含 Password: 的字符串，如果有则立即返回，否则就等待一段时间后返回，这里等待时长就是前面设置的 10秒
expect "*password"

# 发送密码 \r 表示字符串结束
send "yourpassword\r"

# 执行完成后保持交互状态，把控制权交给控制台，这个时候就可以手工操作了。
# 如果没有这一句登录完成后会退出，而不是留在远程终端上。
interact
```



#### tmux 持久化远程 ssh 会话

```shell
# 启动
$ tmux

# 查看当前所有的 Tmux 会话
$ tmux ls

# 接入会话
$ tmux attach -t 0

# 分离会话
$ tmux detach
# 或者快捷键 Ctrl+b d

# 重命名会话
$ tmux rename-session -t 0 <new-name>

# 使用指定名称启动
$ tmux new -s <session-name>

# 滚动屏幕 [scroll in tmux](https://www.freecodecamp.org/news/tmux-in-practice-scrollback-buffer-47d5ffa71c93/)
# Ctrl-b [  # 进入copy mode, 使用 Down/Up 或 PageDown 和PageUp 键翻页, q 或 Enter 退出 copy mode 。
```

```shell
# 窗格管理 panel 快捷键
# Ctrl+b % # 划分左右两个窗格。
# Ctrl+b " # 划分上下两个窗格。
# Ctrl+b <arrow key> # 光标切换到其他窗格。<arrow key>是指向要切换到的窗格的方向键，比如切换到下方窗格，就按方向键↓。
# Ctrl+b ; # 光标切换到上一个窗格。
# Ctrl+b o # 光标切换到下一个窗格。
# Ctrl+b Ctrl+<arrow key> # 按住 Ctrl 不动，松开 b 键，然后同再按 <arrow key> 可以调整空格大小
# Ctrl+b M+<arrow key> #  M 在 Mac 中表示 opt 按键

# 窗口管理 window 快捷键
# Ctrl+b c # 创建新窗口
# Ctrl+b n # 切换到下一个窗口
# Ctrl+b p # 切换到上一个窗口
```

> [Tmux 使用教程](http://www.ruanyifeng.com/blog/2019/10/tmux.html)

> man tmux 能查看更详细的使用说明


#### Mac 解码

- 解码 base64

```shell
% echo -n "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" | base64 -d
{"alg":"HS256","typ":"JWT"}
```

[jwt](https://go-zero.dev/cn/jwt.html?h=jwt)


- 计算 md5

```shell
% echo -n "123456" | md5
e10adc3949ba59abbe56e057f20f883e

% echo  "123456" | md5  # 没有 -n 参数,所以实际上计算的是 123456\n 的哈希值
f447b20a7fcbf53a5d5be013ea0b15af
```

> echo -n 表示输出字符串时,不自动追加 "\n" 这样的换行符


#### Linux 格式化 json base64

使用 jq 命令格式化 json 数据
```shell
# -c 表示去除空格和换行 -compact-output
echo "ewoJImNvbmZJZCI6CSI1Nzc1NjgiLAoJImFwcElkIjoJIjhhMmFmOTg4NTM2NDU4YzMwMTUzN2Q3MTk3MzIwMDA0IiwKCSJ1c2VySWQiOgkiODEwMzEwMzciLAoJInJlY29yZCI6CSJubyIsCgkibW9kZWwiOgkic2luZ2xlIiwKCSJtZW1iZXJJZCI6CSI3MjgxMTUyIgp9" | base64 -d | jq '.' -c

# 无参数时，表示格式化后输出
echo "ewoJImNvbmZJZCI6CSI1Nzc1NjgiLAoJImFwcElkIjoJIjhhMmFmOTg4NTM2NDU4YzMwMTUzN2Q3MTk3MzIwMDA0IiwKCSJ1c2VySWQiOgkiODEwMzEwMzciLAoJInJlY29yZCI6CSJubyIsCgkibW9kZWwiOgkic2luZ2xlIiwKCSJtZW1iZXJJZCI6CSI3MjgxMTUyIgp9" | base64 -d | jq '.' -c | jq '.'

# vim 中使用 python 格式化 json
!%python -m json.tool
```

> [jq tutorial](https://stedolan.github.io/jq/tutorial/)

#### 估算代码行数

```shell
$ find . -name "*.go" -not -path "./vendor/*" | xargs wc -l
```

```txt
. 当前文件夹
-name "*.go* 以 .go 结尾的文件
-not -path "./vendor/*" 忽略 vendor 目录的文件
xargs wc -l 统计 find 命令返回的文件行数
```

```shell
$ find ./_posts/  | xargs wc -l
```

```txt
wc: ./_posts/: Is a directory
      0 ./_posts/
    268 ./_posts/2019-07-14-datarate_bandwidth_datasize.md
    601 ./_posts/2019-09-01-introduction_to_go_modules.md
    261 ./_posts/2019-08-13-mysql_utf8_utf8mb4.md
   1130 total
```

#### Linux查看实时网络速率

```shell
speedometer    # 命令行下实现了速率波形图
iptraf-ng
watch -n 1 "ifconfig eth0"
sar -n DEV 1 2
nload
```

> [cnblogs](https://www.cnblogs.com/klb561/articles/9080151.html)


[use IPERF test network Speed BandWidth](https://www.slashroot.in/iperf-how-test-network-speedperformancebandwidth)

- server

```shell
$ iperf -s -u
------------------------------------------------------------
Server listening on UDP port 5001                           
Receiving 1470 byte datagrams                               
UDP buffer size:  208 KByte (default)                       
------------------------------------------------------------
```

- client

```shell
$ iperf -c 192.168.1.123 -u -b 30m -t 60s
------------------------------------------------------------
Client connecting to ubuntucard.tiga.wang, UDP port 5001
Sending 1470 byte datagrams, IPG target: 392.00 us (kalman adjust)
UDP buffer size:  160 KByte (default)
------------------------------------------------------------
[  3] local 192.168.1.100 port 64768 connected with 192.168.1.123  port 5001
[ ID] Interval       Transfer     Bandwidth
[  3]  0.0-60.0 sec   215 MBytes  30.0 Mbits/sec
[  3] Sent 153063 datagrams
[  3] Server Report:
[  3]  0.0-60.0 sec   215 MBytes  30.0 Mbits/sec   0.000 ms    0/153063 (0%)
```

#### 在 linux 中制作  ubuntu 系统安装盘

首先在 Linux 系统中打开终端，确认 U 盘路径：
```shell
sudo fdisk -l
```

格式化 U 盘，为了格式化首先需要 umount U 盘： 
```shell
# /dev/sdb 是 U 盘设备
sudo umount /dev/sdb*
```

格式化 U 盘：
```shell
sudo mkfs.vfat /dev/sdb -I
```

使用 dd 命令向 sdb 设备（即U盘）写入 ISO 镜像：
```shell
# sudo dd if=xxx.iso of=U盘路径
sudo dd if=~/images/ubuntu.iso of=/dev/sdb
```

输完上述DD命令后回车执行，系统就开始制作启动盘了，
期间终端命令窗口不会有任何反馈，
但能通过U盘运行指示灯看到U盘在进行读写操作，
这个过程可能持续5、6分钟才完成。
当看到终端命令窗口有返回消息即制作完成。

> [CSDN博主「develbai」](https://blog.csdn.net/master5512/article/details/69055662)



#### awk

> [awk](https://www.runoob.com/linux/linux-comm-awk.html)

#### tcpdump

```shell
# 读取pcap文件，并显示摘要
# -tttt 指定显示时间格式，
# -n 显示原始 ip port ，不转换成域名或对应端的的服务名显示。 localhost.telnet 改为 127.0.0.1.23:
$ tcpdump -r ./lo_33434_8845.pcap -ttttN 
2021-07-12 14:59:07.797176 IP test.33434 > test.8845: Flags [S], seq 4268287690, win 65495, options [mss 65495,sackOK,TS val 4248848834 ecr 0,nop,wscale 7], length 0
2021-07-12 14:59:07.797181 IP test.8845 > test.33434: Flags [S.], seq 2411423277, ack 4268287691, win 65483, options [mss 65495,sackOK,TS val 4248848834 ecr 4248848834,nop,wscale 7], length 0
2021-07-12 14:59:07.797185 IP test.33434 > test.8845: Flags [.], ack 1, win 512, options [nop,nop,TS val 4248848834 ecr 4248848834], length 0


# 监听 lo 网卡上的数据，过滤显示 syn 或 fin 包
$ sudo tcpdump "tcp and port 8845 and tcp[tcpflags] & (tcp-syn|tcp-fin) != 0" -i lo -N  
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on lo, link-type EN10MB (Ethernet), capture size 262144 bytes
18:09:01.064281 IP test.49190 > test.8845: Flags [S], seq 1792207613, win 65495, options [mss 65495,sackOK,TS val 4260242101 ecr 0,nop,wscale 7], length 0
18:09:01.064297 IP test.8845 > test.49190: Flags [S.], seq 582524750, ack 1792207614, win 65483, options [mss 65495,sackOK,TS val 4260242101 ecr 4260242101,nop,wscale 7], length 0
18:09:01.067032 IP test.49190 > test.8845: Flags [F.], seq 535, ack 222, win 512, options [nop,nop,TS val 4260242104 ecr 4260242104], length 0
18:09:01.067178 IP test.8845 > test.49190: Flags [F.], seq 222, ack 536, win 512, options [nop,nop,TS val 4260242104 ecr 4260242104], length 0
```

```shell
# 按最后一列的值排序shell 按 column 排序
awk '{print $NF,$0}' file.txt | sort -nr | cut -f2- -d' '
# awk 命令中 `$NF` 表示的最后一个Field（列），即输出最后一个字段的内容

# 删除 60minute 之前创建的文件
find /tmp/ws/ -type f -name 'dump*' -mmin +60 -exec rm -rf {} \;

# crontab -e 添加定时任务  其中第1列表示分钟
# DT:delete pcap files each 5 minute
5 * * * * /usr/bin/find /tmp/ws/ -type f -name 'dump*' -mmin +60 -exec rm -rf {} \;

# 定时抓包
# -G 每隔N秒切分文件，需要和-w配合
# -s 抓取完整的包，0为抓全部包
# -w 把抓包结果输出到文件
# 下面两个是date命令的参数
# %F 年-月-日
# %T 时:分:秒
tcpdump -i any -G 1 -s 0 -w /tmp/ws/dump_any'%F_%T.cap'

# 按大小抓包，每个抓包文件最大 100MB
# -C 每个抓包文件大小，超过则会分割文件，单位MB
tcpdump -i any -C 100 -s 0 -w /tmp/ws/dump_any.pcap
```

> [Linux Tools Quick Tutorial crontab ](https://linuxtools-rst.readthedocs.io/zh_CN/latest/tool/crontab.html)


#### nc netcat

```shell
# listen tcp
$ nc -l 8080

# connect tcp
$ nc 192.168.1.100 80

# listen udp
$ nc -l -u 1234

# connect udp
# nc -v -u 192.168.105.150 53
```

[10 useful ncat (nc) Command Examples for Linux Systems](https://www.linuxtechi.com/nc-ncat-command-examples-linux-systems/)


#### supervisorctl

http://supervisord.org/introduction.html

```shell
# start supervisord
sudo /usr/bin/python /bin/supervisord -c /etc/supervisord.conf

# reload new conf in /etc/supervisord.d/*.conf
sudo kill -HUP `pidof supervisord`
sudo kill -1 `pidof supervisord`

# restart pragram
sudo supervisorctl status all
sudo supervisorctl stop xxxserver
sudo supervisorctl start xxxserver
```

```txt
Linux系统中的信号 http://www.infoq.com/cn/articles/linux-signal-system

在下列情况下，我们的应用进程可能会收到系统信号：
用户空间的其他进程调用了类似kill(2)函数
进程自身调用了类似about(3)函数
当子进程退出时，内核会向父进程发送SIGCHLD信号
当父进程退出时，所有子进程会收到SIGHUP信号
当用户通过键盘终端进程（ctrl+c）时，进程会收到SIGINT信号
当进程运行出现问题时，可能会收到SIGILL、SIGFPE、SIGSEGV等信号
当进程在调用mmap(2)的时候失败（可能是因为映射的文件被其他进程截短），会收到SIGBUS信号
当使用性能调优工具时，进程可能会收到SIGPROF。这一般是程序未能正确处理中断系统函数（如read(2)）。
当使用write(2)或类似数据发送函数时，如果对方已经断开连接，进程会收到SIGPIPE信号。
如需了解所有系统信号，参见signal(7)手册。

-SIGKILL (9) 
-HUP (1)
-INT (2)
-TERM (15)
```

[Intro pages](https://man7.org/linux/man-pages/index.html)

- intro(1) introduction to user commands
- intro(2) introduction to system calls
- intro(3) introduction to library functions
- intro(4)
- intro(5)
- intro(6)
- intro(7) introduction to overview and miscellany section. (describes conventions and protocols, character set standards, the standard filesystem layout.)
- intro(8)



### golang


#### golang 为什么将 struct{} 用做 context.Value() 的 key ？

1. 可隐藏  context 的字段。用 string/int 做 key 时，只要知道变量值就从 ctx 获取 value ，但定义一个小写的 struct ，外部 package 就无法随意取value 了。
2. string 有内存分配开销；

```go
type metadataKey struct{}
type Metadata map[string]string
// NewContext creates a new context with the given metadata
func NewContext(ctx context.Context, md Metadata) context.Context {
        return context.WithValue(ctx, metadataKey{}, md)
}
"vendor/github.com/micro/go-micro/v2/metadata/metadata.go" 126 行
```

[Use struct{} as keys for context.Value() in Go](https://gist.github.com/ww9/4ad7b2ddfb94816a30dfdf2218e02f48)



#### AgoraIO SDK cannot find module rtctokenbuilder 

```diff
--- a/live_auth.go
+++ b/live_auth.go
@@ -12,25 +12,25 @@ import (
        "fmt"

-       "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder"
+       rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
```

- 编译报错

  同一份代码，windows 机器上编译正常，但 linux 上编译提示以下错误：
  
  `live_auth.go|16 col 2| cannot find module providing package github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder: import lookup disabled by -mod=readonly (Go version in go.mod is 1.13, so vendor directory was not used.) `

- 环境：

  go version go1.16.2 但 go.mod 文件仍然使用 go 1.13

- 原因：

  原始目录名是大小写的 RtcTokenBuilder 。

  所以 import "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder" 时,

  由于 linux 系统中区分大小写，找不到名为 rtctokenbuilder 的目录，因为磁盘上保存的目录是 RtcTokenBuilder ；

  而 windows 系统不区分大小写，所以使用 RtcTokenBuilder 或  rtctokenbuilder 都能找到这个文件夹；
 
- github 上的 [Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder/RtcTokenBuilder.go](https://github.com/AgoraIO/Tools/blob/master/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder/RtcTokenBuilder.go) 原始文件

  其文件名是大小写 `RtcTokenBuilder.go`  ，但 package 包名又是小写的 `package rtctokenbuilder` ，不符合 effective go 推荐的做法。

  ```go
  package rtctokenbuilder
  
  import (
  	"fmt"
  	accesstoken "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/AccessToken"
  )
  ```

#### golang append `too many arguments to append`

```go
package main

import "fmt"
import "encoding/json"

func main() {
        buf, _ := json.Marshal(appendArray1(1, 2, 3, 4, 5, 6))
        fmt.Printf("buf %v\n", string(buf))

        buf, _ = json.Marshal(appendArray1(1, 2, 3, 4, 5, 6))
        fmt.Printf("buf %v\n", string(buf))

}

func appendArray1(a1, a2 int, args ...int) []int {
        ret := []int{}
        ret = append(ret, a1, a2)
        ret = append(ret, args...)
        return ret
}

func appendArray2(a1, a2 int, args ...int) []int {
        ret := []int{}
        ret = append(ret, a1, a2, args...)
        return ret
}
```

```sh
$ go run main.go 
# command-line-arguments
./main.go:24:14: too many arguments to append
```

#### install golang in ubuntu
Install [recent Go](https://github.com/golang/go/wiki/Ubuntu):
```txt

$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt-get update
$ sudo apt-get install golang-go
```

#### golang goroutine id

[golang faq](https://golang.org/doc/faq#no_goroutine_id)

Why is there no goroutine ID?

Goroutines do not have names; they are just anonymous workers. They expose no unique identifier, name, or data structure to the programmer. Some people are surprised by this, expecting the go statement to return some item that can be used to access and control the goroutine later.

The fundamental reason goroutines are anonymous is so that the full Go language is available when programming concurrent code. By contrast, the usage patterns that develop when threads and goroutines are named can restrict what a library using them can do.

Here is an illustration of the difficulties. Once one names a goroutine and constructs a model around it, it becomes special, and one is tempted to associate all computation with that goroutine, ignoring the possibility of using multiple, possibly shared goroutines for the processing. If the net/http package associated per-request state with a goroutine, clients would be unable to use more goroutines when serving a request.

Moreover, experience with libraries such as those for graphics systems that require all processing to occur on the "main thread" has shown how awkward and limiting the approach can be when deployed in a concurrent language. The very existence of a special thread or goroutine forces the programmer to distort the program to avoid crashes and other problems caused by inadvertently operating on the wrong thread.

For those cases where a particular goroutine is truly special, the language provides features such as channels that can be used in flexible ways to interact with it.





### git


#### 1.如何配置 git 默认编辑器

```shell
$ git config --list 
$ git config --global core.editor "vim"
```

#### 2.如何使用 fiddler 查看 go get 命令的执行过程

- 1.Fiddler Options -> HTTPS tab
   [ * ] Enable Capture HTTPS CONNECTs 
   [ * ] Decrypt HTTPS traffic

- 2.Fiddler Options -> Connections tab
   Fiddler listens on port: 10809

- 3.配置系统代理和 git 代理
```shell
export https_proxy="http://localhost:10809/"
export http_proxy="http://localhost:10809/"
git config --global https.proxy http://localhost:10809
git config --global http.proxy http://localhost:10809
```
关闭
```shell
unset https_proxy
unset http_proxy
git config --global --unset https.proxy
git config --global --unset http.proxy
```

- 4.关闭 https 证书验证
```shell
git config --global http.sslVerify false
go get -insecure https://xxxxx
```
> 用于防止出现此问题`fatal: unable to access 'https://github.com/pkg/errors/': SSL certificate problem: unable to get local issuer certificate`

- 5.获取git代码
```shell
go get -insecure github.com/pkg/errors
```


#### 3.如何查询某个文件的提交历史？

[https://stackoverflow.com/questions/278192/view-the-change-history-of-a-file-using-git-versioning](https://stackoverflow.com/questions/278192/view-the-change-history-of-a-file-using-git-versioning)

```shell
git log --follow -p --path-to-file
```

`--follow` 选项能让 rename 之前的文件历史也显示出来


#### 4.如何让当前系统记住 git 在各平台的密码？

- Linux系统可以使用 netrc 配置文件保存密码
```shell
$ cat ~/.netrc 
machine github.com 
login wangtiga@gmail.com
password xxxxxxxx
```

- 可以使用[git credentials](https://git-scm.com/docs/gitcredentials) 命令的配置
```shell
// Find a helper.
$ git help -a | grep credential-
credential-foo
// Read its description.
$ git help credential-foo
// Tell Git to use it.
$ git config --global credential.helper foo
```

- mingw 中默认使用 git-credential-store
其文档说明，配置在 ~/.git-credentials 中，打开文件内容如下
```txt
https://wangtiga:xxxxxxxx@github.com
https://wangtiga:xxxxxxxx@gitee.com
```



#### 5.如何合并特定 commitId 的变更到当前分支 [cherry-pick]

```shell
$ git checkout master
Switched to branch 'master'
$ git cherry-pick 99daed2
error: could not apply 99daed2... commit
hint: after resolving the conflicts, mark the corrected paths
hint: with 'git add <paths>' or 'git rm <paths>'
hint: and commit the result with 'git commit'
```

> [git cherry-pick](https://git-scm.com/docs/git-cherry-pick)

> [git tutorial](https://backlog.com/git-tutorial/cn/stepup/stepup7_4.html)


#### 6.如何设置git使用某个http proxy 下载代码？
有时候我的网络环境不仅无法访问Google，连Github都很慢，这时不得不用ssr了
```shell
# 设置代理
git config --global https.proxy https://127.0.0.1:1080
git config --global http.proxy http://127.0.0.1:1080

# 查看配置
git config --global -l

# 取消代理
git config --global --unset https.proxy
git config --global --unset http.proxy
```

> TODO 待确认 使用windows的 ssr 时，“允许来自局域网的连接”并“启用系统代理” 时，在linux主机中如果想使用这个代理。 https.proxy 和 http.proxy 配置的值都是一样的吗？

> $ git config --global https.proxy http://192.168.177.38:1080

> $ git config --global http.proxy http://192.168.177.38:1080


#### 7.如何修改已经commit的log message？
```shell
git commit --amend --author="wangtiga <wangtiga@gmail.com>"
```


#### 8.如何查看git最近几次提交的差异？
```shell
git diff @~7..@~6 -U1000
git show @^0
git show @^1
```

#### 9.如何删除在 remote  中不存在的本地分支?

```shell
# 在 hostA 上删除远程分支 serverfix 后,在 hostB 上也许还残留 serverfix 分支
git push origin --delete serverfix

# 可用以下命令清理本地缓存的远程分支信息
git fetch --prune origin
```

[stackoverflow](https://stackoverflow.com/questions/32147093/git-delete-remotes-remote-refs-do-not-exist)


#### 有哪些git参考资源？

- stackoverflow

https://stackoverflow.com/questions/880957/pull-all-commits-from-a-branch-push-specified-commits-to-another/881014#881014

- git-scm

https://git-scm.com/book/zh/v1/Git-%E5%88%86%E6%94%AF-%E8%BF%9C%E7%A8%8B%E5%88%86%E6%94%AF

- git学习资源汇总(写给Git初学者的7个建议)

http://www.open-open.com/news/view/b7227e

- 在线试用git

https://try.github.io/levels/1/challenges/10

- GIT理解图

https://gitee.com/drinkjava2/UnderstandGIT


- Linux 相关知识

https://huataihuang.gitbooks.io/cloud-atlas/content/


