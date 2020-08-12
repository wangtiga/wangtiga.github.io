---
layout: post
title:  "todo"
date:   1990-01-01 12:00:00 +0800
tags:   todo
---

* category
{:toc}



## LINK

翻译工具 trans
  https://github.com/soimort/translate-shell
  wget git.io/trans

关于 golang 的资料整理
  https://github.com/golang101/golang101

关于 golang 可移植性 为什么原生 golang 代码编译后还依赖 libc？
  https://tonybai.com/2017/06/27/an-intro-about-go-portability/

golang 播放mp3
  https://github.com/hajimehoshi/go-mp3

jekyll
  https://jekyllrb.com/docs/step-by-step/01-setup/

ssr
  https://www.zfl9.com/ss-redir.html

  https://awesomeopensource.com/project/eycorsican/go-tun2socks

  https://code.google.com/archive/p/badvpn/wikis/Examples.wiki

golang error
  https://mp.weixin.qq.com/s/cE_q1LWapFFGYlphZJP-Cw

  https://github.com/pkg/errors

  https://go.googlesource.com/proposal/+/master/design/go2draft.md

golang io.Copy 与内存分配使用
  https://matt.aimonetti.net/posts/2013-07-golang-multipart-file-upload-example/

  https://github.com/wangtiga/test/tgin/upload

- Raspberry 4B + Kodi
  1. Kodi 显示中文要用特定字体，否则中文会乱码
  2. Raspberry 4B 仅在靠近 USB TypeC 接口的 MicroHDMI 才有音频输出。
  2. AirPlay 支持特定版本的 iOS 。
  3. Miracast 暂时不支持。
  4. DLAN 必须播放软件运行。目前 Bilibili 支持。
  5. Android 下载 Kore 快控遥控 Kodi 。
  6. iOS 下载 Kodi Remote 遥控 Kodi 。
  7. TODO 安装 HDV ，导入 m3u 地址观看电视台。
  8. TODO 记录相关开源程序版本，记录 m3u 开源库地址。

- hiwifi openwrt
  1. 恢复成默认开启 ssh 的官方版本。
  2. 刷砖后，恢复的方法。一台电脑与HIWIFI的Lan口网线直连；电脑设置固定的IP,开启 SFTP Server ，增加 Recover.bin 文件；开机HIWIFI过程，持续按信Reset键。
  3. TODO 刷 Openwrt 的过程。


  https://openwrt.io/docs/gee/

  https://imsqua.re/blog/post/squarefong/%E7%94%A8%E6%A0%91%E8%8E%93%E6%B4%BE4B%E5%81%9A%E9%AB%98%E6%B8%85%E6%92%AD%E6%94%BE%E7%9A%84%E7%9B%92%E5%AD%90?nsukey=MvwqDANXxg2aLD2FFhPoMxDS8ZT7%2B6pFZfmbuqj8%2F1eTsrccJQFojPAY74LcFWLbRs9H2Kv7CEQjOf%2BsOnTzppJ%2Bc3e533kglwQI9%2F6XXaqZuRrBXJ8FAYh4wNQl83JaykKSUFnM322bTXpuEOJpPIQ19RbKqC0p6C528pVCaJdITb2BTIPwnA7WTj2CcSYCEoeIIm9zD34EIiiaWwUfFA%3D%3D

  https://nyrra33.com/2019/08/10/unbrick-hc5661-router/

- Boox Note Plus
  1. 下载 Termux 终端模拟器， 安装 vim 编辑器，配合 vim-im 输入法，在 github.io 写日记。使用 ydict 实现终端内查词，且自动发音。
  2. 使用 Markor Android App 预览本地 markdown 文件。
  3. TypeC 口仅支持充电，无任何扩展功能，而且仅支持 USB TypeB + TypeC 的线充电，不支持 TypeC + TypeC 的线充电。
  4. TODO 下载 PDF 资源。
  5. send2boox.com 分发 pdf 等格式电子书资源。
  6. 有音频，听个响的级别，但查词发音时非常有用。

- BenQ 显示器 
  1. USB TypeC 口支持接收4K视频，且支持TNT过程充电，但充电很慢。
  2. 但无法与 Switch 直连，因为 Switch 官方限制，显示器必须有 dock 组件。主 GoBigger 这种三方小厂反而能直连 Switch TypeC 。
  3. 有音频，听个响的级别。

- Gobigger 显示器
  1. MiniHDMI 不方便。
  2. TypeC 充电 + 投屏 ＋ 十指触摸 使用 TNT 。
  3. 显示器必须配合送的保护套才能立起来使用，但保护套是磁吸的，这个支撑过程很脆弱，感觉随时跨台。
  4. 有音频，听个响的级别。


## TODO

(B) dynamic program https://leetcode.com/problems/minimum-path-sum/

(B) golang you-get 了解音视频编码 [csdn 知名博主 雷霄骅 leixiaohua1020](https://blog.csdn.net/leixiaohua1020/article/details/50534150#comments)

(B) MP4 文件合并原理。 录制 RTP 流的实现方案

(C) Golang 中 `go test xxx_test.go` 与  `go test github.com/wangtiga/test` 的区别？为什么vscode使用时，会提示 `cycle reference`

(C) Golang 屏蔽了对 thread 的控制。即使用 Golang 只能控制和使用 goroutine ，但无法控制 thread。 为什么这样做？ 另外，为什么不提供获取 goroutine id 的接口，而且有意屏蔽？

(C) 翻译 [P2P NAT](https://bford.info/pub/net/p2pnat/)


(D) ydict 内置音频播放，去除 mpv 等三方播放组件的依赖。 

(D) 人生马拉松 我们为什么生病

(D) Google提供的基于GO语言和WebSocket的信令服务器[Collider](https://webrtc.org.cn/webrtc_server/)

(D) logrus hook log [wklken/logging-go](https://github.com/wklken/logging-go)

(D) 数组协变逆变、设计缺陷； 泛型、类型原地泛型化  @golang 

(D) 泛型 [Generics — Problem Overview](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)

(D) error check [errors-are-values](https://blog.golang.org/errors-are-values)

(D) android 中映射键盘快捷键。 比如将应用横屏显示？alt+tab 切换另一个任务

(D) vim-im 插件中，如何让这个插件仅在 insert 模式有效？因为在 note plus 中，normal 模式开启 vim-im 插件时，无法使用 / 搜索

(D) aria2 [https://zhuanlan.zhihu.com/p/30666881](https://zhuanlan.zhihu.com/p/30666881)
```shell
# start aria2c
aria2c --enable-rpc --rpc-listen-all

# webui aria2c
node node-server.js
```



## DONE

x (A) 了解 TOOD [语法](https://github.com/todotxt/todo.txt/blob/master/README.md) @markdown

x (A) 练习4751.04 并收集错误 @exam

x (A) 练习4751.10 并收集错误 @exam


x (B) 整理 golang effective @golang

x (B) 整理KB KiB单位的说明

x (B) 整理TCP/IP笔记

x (B) ydict 优化，缓存查询结果，支持离线使用； 是否有必要参考现成开源词典格式？


x (C) 翻译 [slow-down-go-faster](https://www.infoq.com/articles/slow-down-go-faster/?utm_source=wanqu.co&utm_campaign=Wanqu+Daily&utm_medium=website)


x (D) github.io blogs

x (D) android 中如何将网页导出pdf/mobi/epub等电子书格式方便 boox 查阅并批注?答: chrome 打印功能导出pdf十分方便

x (D) markdown 如何显示 mod 取余数符号 @markdown 答：$\mod$

x (D) termux 中输入中文 [hello-termux](https://tonybai.com/2017/11/09/hello-termux/) 也提到了这个问题。 答：使用vim-im插件解决可以输入中文，我用的 [yuweijun修改版](http://4e00.com/blog/vim/2019/03/20/vim-killer-plugin-vim-im-chinese-input-method.html) 默认是极点五笔，跟我的习惯完全一致了




<!--stackedit_data:
eyJoaXN0b3J5IjpbMTMwNjE0NjYwMywyMDcyNzE1MDExLC00ND
kwNTcyNDYsMTYzNzI0NDE2NywyMzA1OTkzMzQsOTg1MTY2NzQ0
LC0yMDc3MTIyNTg0XX0=
-->
