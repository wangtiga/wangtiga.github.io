---
layout: post
title:  "Linux 命令记录"
date:   2019-09-17 18:17:00 +0800
tags: linux
---

* category
{:toc}




### shell


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


#### Mac OS 查看占用端口的程序

https://tonydeng.github.io/2016/07/07/use-lsof-to-replace-netstat/

```shell
sudo lsof -nP -i  :443
```
- -n 表示不显示主机名
- -P 表示不显示端口俗称
- 不加 sudo 只能查看以当前用户运行的程序



### ssh

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
$ chmod 600 ~/.ssh/authorized_keys
```



#### tmux 持久化远程 ssh 会话
```shell
# 启动
$ tmux

# 查看当前所有的 Tmux 会话
$ tmux ls

# 接入会话
$ tmux attach -t 0

# 窗格管理 快捷键
Ctrl+b %：划分左右两个窗格。
Ctrl+b "：划分上下两个窗格。
Ctrl+b <arrow key>：光标切换到其他窗格。<arrow key>是指向要切换到的窗格的方向键，比如切换到下方窗格，就按方向键↓。
Ctrl+b ;：光标切换到上一个窗格。
Ctrl+b o：光标切换到下一个窗格。

# 分离会话
$ tmux detach
# 或者快捷键 Ctrl+b d

# 重命名会话
$ tmux rename-session -t 0 <new-name>

# 使用指定名称启动
$ tmux new -s <session-name>

# 滚动屏幕 [scroll in tmux](https://www.freecodecamp.org/news/tmux-in-practice-scrollback-buffer-47d5ffa71c93/)
$ Ctrl-b [  : 进入copy mode, 使用 Down/Up 或 PageDown 和PageUp 键翻页, q 或 Enter 退出 copy mode 。
```

> [Tmux 使用教程](http://www.ruanyifeng.com/blog/2019/10/tmux.html)



#### Linux格式化 json base64

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
    679 ./_posts/2019-11-03-4735_2019_10.md
     84 ./_posts/2019-10-12-grinder.md
    251 ./_posts/2020-01-31-video_codec.md
    463 ./_posts/2019-12-22-redis_distlock_redlock.md
    410 ./_posts/2019-09-17-playing_with_go_module_proxies.md
    543 ./_posts/2019-09-12-slow_down_go_faster.md
    611 ./_posts/2019-06-17-computer-network.md
     58 ./_posts/2019-10-01-english_bing_today.md
    548 ./_posts/2019-10-01-english_word.md
    177 ./_posts/2019-10-09-influxdb_practice.md
    393 ./_posts/2020-01-01-gstream_basic_tutorial10_gstreamer_tools.md
    363 ./_posts/2020-04-30-zoom_avoid_using_webrtc.md
    601 ./_posts/2019-09-01-introduction_to_go_modules.md
    426 ./_posts/2019-06-12-4751_2014_10.md
    473 ./_posts/2019-11-03-4735_2019_10_sql.md
    654 ./_posts/2020-05-09-gstream_basic_tutorial08_shortcutting_pipeline.md
    986 ./_posts/2019-07-17-golang_intro.md
    352 ./_posts/2020-01-01-gstream_basic_tutorial07_multithreading_and_pad_availability.md
   3731 ./_posts/2019-06-01-effective-golang.md
    375 ./_posts/2019-08-24-tproxy.md
    650 ./_posts/2019-11-06-vim_faq.md
    137 ./_posts/2019-12-22-redis_distlock_setnx.md
    261 ./_posts/2019-08-13-mysql_utf8_utf8mb4.md
    585 ./_posts/2019-06-12-4751_2014_04.md
  14079 total
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


### nc netcat

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


### supervisorctl

http://supervisord.org/introduction.html

```shell
# start supervisord
sudo /usr/bin/python /bin/supervisord -c /etc/supervisord.conf

# reload new conf in /etc/supervisord.d/*.conf
sudo kill -HUP `pidof supervisord`

# restart pragram
sudo supervisorctl status all
sudo supervisorctl stop xxxserver
sudo supervisorctl start xxxserver
```

### golang

#### install golang in ubuntu
Install [recent Go](https://github.com/golang/go/wiki/Ubuntu):
```txt

$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt-get update
$ sudo apt-get install golang-go
```

#### vim-go 插件提示   "No AST for file"

> [github issue 2353](https://github.com/fatih/vim-go/issues/2353)

尝试更新 gopls 程序解决 2019/08/26

```shell
Same issue here. Just occasionally for me, 

gopls  version
version v0.1.0, built in $GOPATH mode

go get -u golang.org/x/tools/cmd/gopls
go install golang.org/x/tools/cmd/gopls

gopls version
version v0.1.3-cmd.gopls, built in $GOPATH mode
```



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


