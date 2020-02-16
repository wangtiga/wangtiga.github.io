---
layout: post
title:  "Linux 命令记录"
date:   2019-09-17 18:17:00 +0800
tags: linux
---

* category
{:toc}




### shell

#### Linux查看实时网络速率

```shell
iptraf-ng
watch -n 1 "ifconfig eth0"
sar -n DEV 1 2
# nload
```

> [cnblogs](https://www.cnblogs.com/klb561/articles/9080151.html)


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
```


#### 9.有哪些git参考资源？

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

