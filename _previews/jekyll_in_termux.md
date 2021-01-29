---
layout: post
title:  "Jekyll In Termux"
date:   2021-01-24 09:22:00 +0800
tags:   tech
---


* category
{:toc}




## termux

Install Termux from [F-Droid](https://f-droid.org/).

F-Droid 是一个 Android 平台上 FOSS（Free and Open Source Software，自由开源软件）的目录，
并提供下载安装支持。
使用客户端可以更轻松地浏览、安装及跟进设备上的程序更新。


### install & [Remote_Access] https://wiki.termux.com/wiki/Remote_Access)

```sh
pkg upgrade 
pkg install openssh git vim
sshd
passwd
ssh anyname@host -p 8022
```

- how to ssh from external ip?

  默认端口是 8022 ；
  默认能从外部IP访问到SSH,如果异常，可能VPN相关软件影响，关闭并重启手机解决；

- how to add user in termux?

  termux 是单用户，所以使用任意用户名都能登录

- install jekyll

  ```sh
  pkg install ruby
  gem install jekyll
  # gem install "all extention in _config.yaml"
  ```


#### .bashrc

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

alias ll='ls -alFh'
alias gll='git branch -avv'
alias gpush='git add . && git commit -m "updated from newdevice" && git push'
alias y='ydict -q -v -v -c '
```


## jeykll

```shell
#gem list --local  | grep jekyll
#gem uninstall jekyll --version 3.8.6
#gem install jekyll --version 3.8.6
#bundle exec jekyll serve

# Quick-start Instructions https://jekyllrb.com/
gem install bundler jekyll
jekyll new my-awesome-site
cd my-awesome-site
bundle exec jekyll serve
# => Now browse to http://localhost:4000
```

以下是个人的浅显理解，没有验证过真伪

- Gem
  
  gem 可理解为包管理器，是用 ruby 语言实现的

- Gemfile 

  记录当前项目依赖的第三方库，及相关版本。

- Gemfile.lock 

  某些命令后，当前使用的第三方库的版本会记录在这个文件，下次再执行相关命令时，不会自动更新版本。除非删除这个文件

- Bundler

  bundle 是包含所有相关三方库的命令，可把它理解成一个运行环境

- jekyll 。

  jekyll 也一个 ruby 程序，主要实现将 markdown 翻译成 html 语言，生成 blog 所需要的网页文件

- `_config.yml`

  记录当前 blogs 的环境配置，比如博客的名称，联系人的邮箱等。

  TODO 也会记录当前 blogs 使用的第三方库，可能是给 github pages 后台服务使用的吧？
  但是为什么要在 Gemfile 与 `_config.yml` 两个地方分别记录第三方库呢？




https://stackoverflow.com/questions/58195772/how-can-fix-the-jekyll-installation-error-in-termux

https://wiki.termux.com/wiki/List_of_Ruby_packages_known_to_work

[setting-up-a-github-pages-site-with-jekyll](https://docs.github.com/en/github/working-with-github-pages/setting-up-a-github-pages-site-with-jekyll)

[Jekyll ruby 101](https://jekyllrb.com/docs/ruby-101/)

[termux wiki Jekyll](https://wiki.termux.com/wiki/Jekyll)

[一个功能较全面的 jekyll 主题](https://github.com/professordeng/log)

