---
layout: post
title:  "Go语言入门资料"
date:   2019-07-17 18:42:00 +0800
tags: golang
---

* category
{:toc}

## 什么是GO
### hello world [demo online](https://play.golang.org/p/5_WIDFovQT4)
```go
package main // 当前 package ， main 是程序入口 package
import (
	"fmt" // 当前文件引入的外部 package
)

func main() { // main() 是程序的入口 function
	fmt.Println("Hello, playground")
}
```
> 推荐 https://play.golang.org/ 
> 小功能在线调试；问题交流时方便分享代码；


### demo C1000K Go/C++ 测试对比[^C1000kPracticeGuide]

单机100 万的连接

服务器每 10 秒向所有客户端发 500 byte的消息

- C++ [code](https://github.com/xiaojiaqi/C1000kPracticeGuide/blob/master/code/cppserver/src/main.cc)
  * 1000 行代码，epoll pipe cpu 亲缘性
  * Load average: 1.89, 1.59, 1.07
  * Memory: 14761404 / 16428352 k
  * CPU: 99% 77% 76% 75%
- Go [code](https://github.com/xiaojiaqi/C1000kPracticeGuide/blob/master/code/goserver/server.go)
  * 100 行代码，无优化
  * Load average: 1.11, 1.14, 1.05
  * Memory: 15590444 / 16428352 k
  * CPU: 99% 33% 100% 64%


### Go的主要特点有哪些？[^TheWayToGo]
- 类型安全 和 内存安全
- 以非常直观和极低代价的方案实现 高并发
- 高效的垃圾回收机制
- 快速编译（同时解决C语言中头文件太多的问题）
- 为多核计算机提供性能提升的方案
- UTF-8编码支持

## Go语言在Google的起源发展
### 问题[^GoAtGoogleEN] [^GoAtGoogleCN]
-   Google 有大量软硬件，上百万行 c++ java python 代码
-   巨量软件运行在无数的硬件中， 存在大量分布式，集群系统
-   开发工作变得 缓慢而笨拙
    *  C++项目编译时间在 20 ~ 40 分钟

### 解决 [^TheWayToGo]
#### Pike、Griesemer、Thompson 是Golang的主要设计者
![Pike、Griesemer、Thompson](https://box.kancloud.cn/2015-10-23_5629d848a5508.png)
- 罗布*派克(Rob Pike) - Unix编程环境
- 罗伯特*格里泽默(Robert Griesemer)  - ChromeV8
- 肯*汤普逊(Ken Thompson) - Unix

#### 现代化的系统编程语言
- 提供原生的并发支持
GO 主要面向 并发网络服务 的现代化系统编程
- 并发编程的的复兴
CPU速度会提升，但单个硬件的性能无法满足要求。需要发掘多核，即并行的潜力。
- 与 应用编程语言 相对
	* System programming系统编程：用于生产软件的程序，如编译器；用于给其他软件提供服务的程序，如接口服务器、OS、游戏引擎。
	* Application programming应用编程： 用于面向使用者的程序，如 Word、QQ、微信等。

#### 提升开发效率的实用型编程语言
- 易于开发
  *   类似动态语言的特性，垃圾回收、interface等
  *   简洁而统一的代码风格
  *   丰富的官方工具 gofmt 格式化 gofix升级golang godoc生成文档
- 快速编译
  *   编译和链接到机器代码的速度很快
  *   构建一个程序的时间只需要数百毫秒到几秒
  *   整个 Go 语言标准库的编译时间一般都在 20 秒以内
- 高效执行
  *   网络通信、并发和并行编程的极佳支持，更好地利用分布式和多核计算机
- 易于部署
*   编译出的一个可执行文件
*   支持常见平台
*   纯Go代码仅依赖几个基本库

#### 取长补短
吸收 C/Java/JavaScript/Unix-plan9-limbo 的优点
![golang_inheritance](https://box.kancloud.cn/2015-10-23_5629d84ac42d4.png)
- 和 C++、Java 和 C# 一样属于 C 系。
- 来自C语言的基本语法，和面向过程的编程风格
- 来自Java的Interface
- 来自C#等语言的 package 管理
- 在声明和包的设计方面，Go 语言受到 Pascal、Modula 和 Oberon 系语言的影响
- 在并发原理的设计上，Go 语言从同样受到 Tony Hoare 的 CSP（通信序列进程 Communicating Squential Processes）理论影响的 Limbo 和 Newsqueak 的实践中借鉴了一些经验

### 发展[^GoClimbingEN] [^GoClimbingCN]
#### 2007年 Google 内部 20% 兼职项目
##### Rob Pike 在 2007 年  
Rob Pike 在 2007 年 9 月 25 号，星期二，下午 3：12 回复给 Robert Griesemer、Ken Thompson 的有关编程语言讨论主题的邮件
```txt
在开车回家的路上我得到了些灵感
给这门编程语言取名为“go”，它很简短，易书写。
工具类可以命名为：goc、 gol、goa。
交互式的调试工具也可以直接命名为“go”。语言文件后缀名为 .go 等
```

##### Ian Lance Taylor 在 2008 年
Ian Lance Taylor 在 2008 年 6月 7 日（星期六）的晚上 7：06 写给 Robert Griesemer、Rob Pike、 Ken Thompson 的关于 Go gcc 编译器前端的邮件
```txt
我的同事向我推荐了这个网站 http://…/go_lang.html 。这似乎是一门很有趣的编程语言。我为它写了一个 gcc 编译器前端。虽然这个工具仍缺少很多的功能，但它确实可以编译网站上展示的那个素数筛选程序了。
```
```txt  
 Ian Lance Taylor 的加入以及第二个编译器 (gcc go) 的实现 在带来震惊的同时，也伴随着喜悦。这对 Go 项目来说不仅仅是鼓励，更是一种对可行性的证明。语言的第二次实现对制定语言规范和确定标准库的过程至关重要，同时也有助于保证其高可移植性，这也是 Go 语言承诺的一部分。自此之后 Ian Lance Taylor 成为了设计和实现 Go 语言及其工具的核心人物。
```
#####  Russ Cox 在2008年
Russ Cox 在2008年带着他的语言设计天赋和编程技巧加入了刚成立不久的 Go 团队
```txt
Russ 发现 Go 方法的通用性意味着函数也能拥有自己的方法，这直接促成了 http.HandlerFunc 的实现，这是一个让 Go 一下子变得无限可能的特性。Russ 还提出了更多的泛化性的想法，比如 io.Reader 和 io.Writer 接口，奠定了所有 I/O 库的整体结构。
```
#### 2009年 正式对外发布 
源代码最初托管在 http://code.google.com 上，之后几年才逐步的迁移到 GitHub 上。
聘请了安全专家 Adam Langley 帮助 Go 走向 Google 外面的世界
```txt
Adam 为 Go 团队做了许多不为外人知晓的工作，包括创建最初的 http://golang.org 网站以及 build dashboard。不过他最大的贡献当属创建了 cryptographic 库。起先，在我们中的部分人看来，这个库无论在规模还是复杂度上都不成气候。但是就是这个库在后期成为了很多重要的网络和安全软件的基础，并且成为了 Go 语言开发历史的关键组成部分。许多网络基础设施公司，比如 Cloudflare，均重度依赖 Adam 在 Go 项目上的工作，互联网也因它变得更好。我记得当初 beego 设计的时候，session 模块设计的时候也得到了 Adam 的很多建议，因此，就 Go 而言，我们由衷地感谢 Adam。
```
#### 2016年 被评选为 TIOBE 2016 年最佳语言
[https://www.bnext.com.tw/article/42761/tiobe-2016-program-language](https://www.bnext.com.tw/article/42761/tiobe-2016-program-language)
[https://www.tiobe.com/tiobe-index/go/](https://www.tiobe.com/tiobe-index/go/)

#### Go 在云计算取得了领导地位[^GolangProject]
-   Docker 就是使用 Go 进行项目开发，并促进了计算机领域的容器行业，进而出现了像 Kubernetes 这样的项目。现在，我们完全可以说 Go 是容器语言，这是另一个完全出乎意料的结果。

-   kubernetes 已经成为了所有云计算公司的底层架构，而且越来越多的互联网公司系统架构迁移到 k8s 上面，例如 github、阿里、腾讯、百度、滴滴、京东等大型企业纷纷拥抱，而这个系统就是 Go 开发的




## 开发环境与开发工具
### 开发工具IDE
- Sublime
- VSCode
- Goland
- Vim
- 在线调试工具 Go Playground
[ https://play.golang.org/](https://play.golang.org/)

### 安装[^GolangInstallEN] [^GolangInstallCN]
不同系统稍有区别，以Linux 为例
```shell
# 下载
$ wget https://studygolang.com/dl/golang/go1.11.linux-amd64.tar.gz <https://studygolang.com/dl/golang/go1.11.linux-amd64.tar.gz>

# 解压
$ tar -C /usr/local -xzf go1.11.linux-amd64.tar.gz

# 将可执行目录添加到 PATH 即可
$ export PATH=$PATH:/usr/local/go/bin

# 命令行输入 go version ，输出版本号就表示安装成功
$ go version
go version go1.9.1 windows/amd64
```
$GOPATH 和 $GOROOT是常用的环境变量
```shell
#  命令行输入 go env
$ go env
set GOARCH=amd64
set GOBIN=
set GOEXE=.exe
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=F:\ws\gopath
set GORACE=
set GOROOT=C:\Go
set GOTOOLDIR=C:\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set CC=gcc
set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0
set CXX=g++
set CGO_ENABLED=1
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
```
### 常用 go 命令
在命令行或终端输入go即可查看所有支持的命令
-   go get：获取远程包（需 提前安装 git）
-   go run：直接运行程序
-   go build：测试编译，检查是否有编译错误
-   go fmt：格式化源码（部分IDE在保存时自动调用）
-   go install：编译包文件并编译整个程序
-   go test：运行测试文件
-   go doc：查看文档

### 代码目录结构
```shell
$GOPATH
    bin
        go intall 等命令生成的可执行文件
        test.exe
    pkg
        go insstall / go build 命令生成的 lib 文件
        linux_amd64
            test.a
        windows_amd64
            test.a
    src
        go get 命令自动将源码下载到此目录
        github.com
            wangtiga
                golangintro
                    .git
                    demo
                        websocket_chat
                            vendor
                                当前项目引用的第三方源码
                                    可用 govendor.exe 管理
                                    目录结构同  gopath/src
                                github.com
                                    gorilla
                                        websocket
                                            client.go
                                            doc.go
                                            README.go
                                vendor.json
                            client.go
                            main.go
                            README.md
        gitee.com
```

## demo WebSocket Chat
使用 `gorilla/websocket`[^GolangWebSocket] 实现的简易聊天室 ，任意成员发送数据，其他人都能看到 
```shell
# 设置环境变量
$ export GOPATH=/d/ws/tmp/gopath 
$ export PATH=/d/ws/tmp/gopath/bin:$PATH

# 下载源码
$ go get github.com/wangtiga/golangintro
# 编译并安装 demo 的linux版本到 $GOPATH/bin
$ CGO_ENABLED=0 GOOS=linuxGOARCH=amd64 go install github.com/wangtiga/golangintro/demo/websocket_chat 

# 编译并安装 demo 到 $GOPATH/bin 
$ go install github.com/wangtiga/golangintro/demo/websocket_chat??&& websocket_chat.exe
2019/02/16 15:49:30 ListenAndServe2019/02/16 15:49:30 ListenAndServe?on??:8080 on??:8080
```

### 涉及后台开发常用技术
-   http
-   interface struct
-   channel goroutine
-   vendor第三方源码


## Go 语言特性[^GolangProgramXSW]
### 自动垃圾回收[^GolangProfiling] [^GolangFAQ_GC]
#### [C示例](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/garbage_collection.c)
```c
#include <stdio.h>

void foo() {
	// 给 p 动态分配内存
	char* p = malloc(10);

	// work with p
	p[0] = 65;
	p[1] = 66;
	p[2] = 66;
	p[3] = '\0';
	printf("%s\n", p);

	// 需要显式调用  free() 释放内存空间
	free(p);
	p=NULL;
}

int main() {
	foo();
	return 0;
}
```
- 手动资源管理一直是个大麻烦
引发的问题往往十分严重，又难以解决
dangling pointer / wild pointer / core dump / segment fault / stack overflow
- 即允许手动销毁资源又提供析构函数这样自动销毁资源的机制
增加程序设计的负担

#### 权衡
- concurrent 编程中，最困难的问题就是 object lifetime problem
    要保证安全地清理在 thread 之间传递的 object ，是很困难的
- 垃圾回收机制会带来很大的成本
    资源的消耗、回收的延迟以及复杂的实现等
    最近 garbage collection 技术的发展也让我们有自信能以较低的代价实现这个功能（延迟较低）
    优秀的 garbage collection 不容易。但我们实现一次，就免于每个人重复劳动
- 给程序员减轻负担很重要
    特别是对于程序员的编程体验来说，是要大于它所带来的成本的，因为这些成本大都是加诸在编程语言的实现者身上

#### [Go 示例](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/declare.go)
```go
package main

import (
	"fmt"
)

func main() {
	foo()
}

func foo() {
	var p []byte = make([]byte, 100)
	p[0] = byte(65)
	p[1] = byte(66)
	p[2] = byte(66)

	fmt.Println(string(p))
}
```

- 只提供 GC 一种方式销毁资源，即自动垃圾回收
- 不仅如此，Golange有意简化变量生存周期的管理
    [看看这个代码是否有Bug](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/garbage_collection2.go)
- 一直在探索新方法降低开销和延迟
     collector 造成的 pause time 停顿时间，即使很大的 heaps 已经降低到 sub-millisecond （< 1ms，即亚微秒，不到1微秒）

- 提供 `go tool pprof`[^GolangProfiling] 分析内存、CPU的使用情况


### 语法声名[^GolangSyntax] [^GolangFAQ_DeclarationsBackwards]
#### [C示例](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/declare1.c)
- 观察下面C代码中， f1 f2 f3 三个变量分别表示什么类型 ?
- [使用 typedef 提升可读性的版本](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/declare2.c)
- [编译器分析语法时，是忽略变量名称的](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/declare3.c)

```c
#include <stdio.h>

// https://blog.golang.org/gos-declaration-syntax
// https://stackoverflow.com/questions/1591361/understanding-typedefs-for-function-pointers-in-c

int *p1;
int p2[3];
int *p3, p4; // p3 和 p4 的类型是一样的吗？

// fp1 是一个函数指针
int (*fp1)(int a, int b); 
int (*fp2)(
    int (*arg1)(int x, int y), // arg1 is pointer to function
    int arg2);
int (* (*fp3)(
        int (*arg1)(int x, int y), // arg1 is pointer to function
        int arg2) ) (int a, int a); // return argument of fp3  is pointer to function
```

- 螺旋式阅读
人不易读，使用代码也不易解析。不好写工具

#### [Go示例](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/declare.go)

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    var f1 func(a int, b int) int

    var f2 func(
        arg1 func(x int, y int) int,
        arg2 int,
    ) int

    var f3f func(
        arg1 func(x int, y int) int,
        arg2 int,
    ) func(a int, b int) int

    // 声名一个匿名函数，并立即调用，将返回结果存入 sum 变量中
    sum := func(a, b int) int { return a + b }(3, 4)
}
```

- 类型在前的声名
    简洁易读，也容易进行语法解析 
    方便自动化工具 如 go fmt 和 go fix 等  

- 可以这样记忆
 
```go
//  x: int  
var x int  

//  p: pointer to int  
var p *int  

//  a: array[3] of int 
var a [3]int 
```

### 错误处理[^GoAtGoogleEN] [^GoAtGoogleCN]
#### [Java示例](https://github.com/wangtiga/golangintro/blob/master/demo/echo/java/EchoServer.java)

```java
    try {
         serverSocket = new ServerSocket(10007);
    }
    catch (IOException e) {
         System.err.println("Could not listen on port: 10007.");
         System.exit(1);
    }
```

- try-catch-finally
错误处理扭曲程序的控制流

#### [Go示例](https://github.com/wangtiga/golangintro/blob/master/demo/echo/golang/echoserver.go)
Go源码中的错误处理风格是很好的示例，比如[ioutil.readAll](https://github.com/golang/go/blob/release-branch.go1.12/src/io/ioutil/ioutil.go#L18)

```go
	addr := "127.0.0.1:80"
	listener, err := net.Listen("tcp", addr)
	if nil != err {
		panic("starting TCP server Fail, err=" + err.Error())
	}
	defer listener.Close()
```

- Go支持多值返回，但一般习惯[ error 是最后一个变量](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/multi_return.go#L15)

- panic / recover 类似 throw / try catch  

```txt
Go does provide mechanisms for handling exceptional situations such as division by zero. A pair of built-in functions calledpanicandrecoverallow the programmer to protect against such things. However, these functions are intentionally clumsy, rarely used, and not integrated into the library the way, say, Java libraries use exceptions
```

- defer 表达式在当前函数结束时调用，用于手动释放资源

- 区分 错误 和 异常[^GolangErrorAndException]

```txt
不会终止程序逻辑运行的归类为错误，比如用户输入密码错误。
会终止程序逻辑运行的归类为异常，比如系统运行环境被环境，某些文件丢失。
```

- 显式处理错误
Go中不包含异常，是我们故意为之的。编码带来的清晰度和简单性可以弥补其冗长的缺点。
Go检查错误的方式更加繁琐。但这种显式的设计使得控制流更加直截了当。

#### 总结
- 错误处理，培养及时思考并应对错误的习惯
显式地错误检查会迫使程序员在错误出现的时候对错误进行思考并进行相应的处理。

- 异常处理，会推迟甚至遗忘问题
异常机制只是将错误处理推卸到了调用堆栈之中，直到错过了修复问题或准确诊断错误情况的时机，这就使得程序员更容易去忽略错误而不是处理错误了。

### 类型和接口[^GoAtGoogleEN] [^GoAtGoogleCN]
#### Java / C++ 中的 Object Oriented
```java
class Foo implements IFoo {
   // Java 语法
   // ...
}
class Foo : public IFoo {
  // C++ 语法
  // ...
}
```

**如何定义一个 IFile接口？**
要求IFile同时支持 Read() Write() Close()等几种操作

- 方法1：仅定义一个IFiler?
无法照顾到不同使用者的需求。比如，若需要随机读功能时如何实现？
可以在IFiler中增加成员方法 Seek() ，但不需要随机读的场景，也就没必要实现 Seek() 方法。
也可以实现一个新的接口ISeekFiler。继承自IFiler，只比IFiler多一个Seek()方法。但这会产生[Javaio包中复杂的层级关系](https://docs.oracle.com/javase/7/docs/api/java/io/package-tree.html)。


- 方法2：每个功能定义一个接口?
把 IFiler 拆分成 IReader / IWriter / ICloser 三个接口后，紧跟着就要考虑下面的问题。
是否需要定义一个 IFiler 接口继承这三个接口？这就有可能增加接口抽象层次。
是否需要定义 ISeeker 接口呢？定义了ISeeker，实际业务却一直用不到就尴尬了。这很容易过度设计，产生垃圾代码。
其实这两个问题需要根据实际使用场景来确定。这就需要库的实现者预知未来。但随着业务发展，改变是必然。

看看Golang作者之一Rob Pike怎么说的[要组合，不要继承](https://talks.golang.org/2012/splash.article#TOC_15.)[^GoAtGoogleEN] [^GoAtGoogleCN]：
> 类型层次结构容易造成非常脆弱的代码。层次结构必需在早期进行设计，通常会是程序设计的第一步，而一旦写出程序后，早期的决策就很难进行改变了。所以，类型层次结构这种模型会促成早期的过度设计，因为程序员要尽力对软件可能需要的各种可能的用法进行预测，不断地为了避免挂一漏万，不断的增加类型和抽象的层次。这种做法有点颠倒了，系统各个部分之间交互的方式本应该随着系统的发展而做出相应的改变，而不应该在一开始就固定下来。

- 另一种不同的思路
    关注方法，而不是接口
    随系统的发展而做出改变，而非一开始就固定

#### Go 的  Glang Object Oriented
《Go语言编程 3.5.1》[^GolangProgramXSW]：[interface.go](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/interface.go)
MP3Player 实现了 Play() 方法，所以就能当作 Player 接口使用。不需要显式声明自己实现了哪个接口。
```go
func testMain() {
    var p Player
    p = &MP3Player{}
    p.Play(source)
}

type Player interface {
    Play(source string)
}

type MP3Player struct {
    stat     int
    progress int
}

func (p *MP3Player) Play(source string) {
    fmt.Println("Playing MP3 music", source)
}
```

- [EffectiveGo embedding](https://golang.org/doc/effective_go.html#embedding)[^EffectiveGoEN]
针对刚才IFiler的要求，可以定义 Reader Writer Closer Seeker 四个接口，再让 IFiler `继承`这四个接口。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type IFiler interface {
    Reader
    Writer
    Closer
    Seeker
}
```

因为 struct 不用显式声名自己实现了哪个 interface ，只要实现 interface 定义的 method，就能当作 interface 使用。
所以Go巧妙地保持了自己的灵活性，还不用担心过度设计。

-  GO源码中[io package](https://github.com/golang/go/blob/release-branch.go1.12/src/io/io.go#L77) 就是这样做的
其中计算文件的 sha1 值 [crypto/sha1 example](https://github.com/golang/go/blob/release-branch.go1.12/src/crypto/sha1/example_test.go)

```go
func ExampleNew_file() {
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    h := sha1.New()
    if _, err := io.Copy(h, f); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("% x", h.Sum(nil))
}
```

```go
// 相关定义如下
// sha1.New() 的定义
func New() hash.Hash {}

// io.Copy() 的定义
func Copy(dst Writer, src Reader) (written int64, err error) {}

// digest struct 的定义 
type digest struct {}
func (d *digest) Write(p []byte) (nn int, err error) {}

// Hash interface 的定义
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	io.Writer

	Size() int
	// .... 省略
}
```

#### 总结
- 组合而非继承
    实现了接口要求的方法，就实现了接口
    不关心继承树
    关注方法，而不是接口
    随系统的发展而做出改变，而非一开始就固定


### 并发编程[^SevenConcurrencyModelsSevenWeeks]  [^SevenConcurrencyModelsSevenWeeksSource]  [^GoroutinesVSThreads] [^ConcurrencyIsNotParallelism] [^WhatIsConcurrencyParallelism]

#### 并发的含义
- demo 输出字母和数字

* [thread实现](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/thread2.c)

```c
#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

void *numbers();
void *alphabets();
int main()
{
    int rc1, rc2;
    pthread_t thread1, thread2;

    /* 创建线程，每个线程独立执行函数function count */
    if((rc1 = pthread_create(&thread1, NULL, &numbers, NULL))) {
        printf("Thread creation failed: %d\n", rc1);
    }
    if((rc2 = pthread_create(&thread2, NULL, &alphabets, NULL))) {
        printf("Thread creation failed: %d\n", rc2);
    }

    sleep(3); // 3 second
    printf("main terminated in c\n");
    return 0;
}
void *numbers()
{
    for (int i = 1; i <= 5; i++) {
        usleep(250 * 1000); // 250ms
        printf("%d ", i); fflush(stdout);
    }
}
void *alphabets()
{
    for (char i = 'a'; i <= 'e'; i++) {
        usleep(400 * 1000);
        printf("%c ", i); fflush(stdout);
    }
}
```

* [goroutine实现](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/thread2.go)   

```go
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated in golang")
}
```

* 运行结果： `1 a 2 3 b 4 c 5 d e`

- thread
   * 内存(stack size)占用较大 >1MB
   * thread恢复运行时，需要还原大量 CPU 寄存器，会降低性能
   * thread的创建和销毁过程都需要 系统调用，这个操作也很慢
- goroutine
   * 初始内存(stack size)仅 2KB ，并按需要增加或减少
   * 减少 thread 切换，使用 channel 阻塞时，OS 不会 block thread
   * 能大量的使用 goroutine
   * 串行的代码逻辑，相比事件回调的代码，易写易维护
   * 配合 channel 减少 lock 代码

#### 单词统计 的需求
-  1.如何统计文本中各个单词出现的频率？
fileA: `one potato  two potato three potato four potato`

-  2.如何提高效率，尽量少的时间得到统计结果？
	当要统计文本数据量很大时，可把数据分成几份，使用两个CPU同时执行统计任务，即并行统计。详细过程如下：
	* 2.1.将fileA分隔成以下两部分
		fileA1: `one potato two`
		fileA2: `three potato four potato`
	* 2.2.执行统计
		启动CPU1统计fileA1的数据，将统计结果写入 resultA
		启动CPU2统计fileA2的数据，将统计结果写入 resultA
		最终结果为：`resultA = (one=1 two=1 three=1 potato=3)`
		这个过程还要考虑下面的问题：
		详细合并过程是怎样的？两个CPU同时写resultA会不会有什么问题？

> 注：要求统计结果实时合并到全局统计结果中。否则，如果每个统计过程都使用独立的内存变量保存结果，那就不存在数据竞争，也就没法展开下面的讨论了。
> 总体来说，使用共享的全局变量合并统计结果，相对节省省内存空间。两者其他的优缺点这里不再详细考虑。

- 3.合并结果的一个办法：[demo lock 共享结果](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/count_word_bylock.go)

```go
package main

import (
	"fmt"
	"sync"
)

var gResultsLock sync.Mutex
var gResults map[string]int = map[string]int{}

func main() {
	words1 := []string{"one", "potato", "two", "potato"}
	words2 := []string{"three", "potato"}

	go countRoutine(words1)
	go countRoutine(words2)

	fmt.Printf("Result: %#v\n", gResults)
}

func countRoutine(words []string) {
	for _, item := range words {
		gResultsLock.Lock()
		gResults[item] += 1
		gResultsLock.Unlock()
	}
}
```

- 4.合并结果的一个办法：[demo channel 共享结果](https://github.com/wangtiga/golangintro/blob/master/demo/example_piece/count_word_bychannel.go)

```go
package main

import "fmt"
import "time"

func main() {
	words1 := []string{"one", "potato", "two", "potato"}
	words2 := []string{"three", "potato"}

	ch := make(chan map[string]int, 4)
	go countRoutine(ch, words1)
	go countRoutine(ch, words2)
	results := map[string]int{}
	for {
		select {
		case nResult := <-ch:
			for key, num := range nResult {
				results[key] += num
			}
		case <-time.After(time.Duration(2) * time.Second):
			fmt.Printf("Result: %#v\n", results)
			return
		}
	}
}

func countRoutine(ch chan map[string]int, words []string) {
	results := map[string]int{}
	for _, item := range words {
		results[item] += 1
	}
	ch <- results
}
```


##### 线程和锁模型
* 是其他技术的基础：互斥量 mutex，条件变量 condition variables  和内存屏障 memory barriers  等技术都基于锁模型
* 适用面广
* 正确使用时，效率很高
* 难于测试，bug偶现，哲学家进餐问题，无锁代码运行一周才出问题

##### CSP通信顺序进程 
- CSP
    Hoare's 的 通信顺序进程 CSP 就是一种较为成功的并发编程模型
    CSP 模型非常适合面向过程的语言 procedural language
    communicate by sharing memory, share memory by communicating

- channel
类似 Unix 管道
可理解为 阻塞 的 先进先出的队列

```go
// 声名
// make的第2个参数指定了队列的最大缓存长度
// ch 中的数据量大于5个时，write 操作阻塞
ch := make(chan int, 5)  

// 写入
// 写入数值4到 channel
ch <- 4 

// 读取
// 从channel读取数值，并存入 val 变更
val := <- ch
```
- 注意 使用channel 没有锁，但也会“死锁”
    了解 select 的用法解决

### 反射
- Java出现后流行的一种概念，通过反射动态操作类型值
- 这些库代码中会频繁使用
   比如 encoding json / encoding  xml / db orm 等
- 能实现强大的功能，但代码复杂。
- 《Go语言编程》[^GolangProgramXSWSource]中的[示例](https://github.com/qiniu/gobook/blob/master/chapter9/reflect.go)

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
}
```

### 语言交互性
- 无数代码库仍然只有C语言版本
    比如访问 Oracle 数据库的功能，只能使用 C 语言的库
- 作为传承于C言的GO，提供Cgo的技术调用C语言的代码
- CGo 作用与 Java 中的 JNI 一样
- 《Go语言编程》[^GolangProgramXSWSource]中的[示例](https://github.com/qiniu/gobook/blob/master/chapter9/cgo2.go)

```go
package main 

/*
#include <stdio.h>
void hello() {
    printf("Hello, Cgo! -- From C world.\n");
}
*/
import "C"

func Hello() {
  C.hello()
}

func main() {
  Hello()
}
```

## 学习资源
- [官方文档](https://golang.org/doc)内容十分全面
   包含 入门、安装、开发环境、effective go 、问题诊断 （profiling/debuging 等工具）、常见问题FAQ，wiki
- 所有关于golang的问题都能从FAQ[^GolangFAQ]中找到资料 
- 资源
    * Golang 示例
        <https://gobyexample.com/>
    * Golang 向导
        <https://tour.golang.org/list>
    * Go 入门指南
        <https://www.kancloud.cn/kancloud/the-way-to-go/72432>
    * 官方文档
        <https://golang.org/doc/>
- other [^BuildWebApplicationWithGolang]


## TODO 其他待整理的资料
- 区分 并发与并行
- 如何实现非阻塞的 channel ，了解 select 的用法
- [slice执行append函数的坑](https://play.golang.org/p/XE9ymg4Pn0T)
- [interface 与 nil 的坑](https://play.golang.org/p/trnp_70gD4j)
  [ref1](https://deepzz.com/post/why-nil-error-not-equal-nil.html) [ref2](https://speakerdeck.com/campoy/understanding-nil)
- [time.Time 比较的坑](https://play.golang.org/p/hiwJiO3gqjc)
- 搜集有关批评Go错误处理的讨论
- golang不支持泛型（ 类似 c++ Template )

### 3.Go的垃圾回收提供了哪些优化的特性 ?
查找这段描述的来源，好像并不是 GolangGCJourney[^GolangGCJourney]
- 在Java中，buffer字段需要再次进行内存分配，因为需要另一层的间接访问形式。
- 在Go中，该缓冲区同包含它的struct一起分配到了一块单独的内存块中，无- 需间接形式。
- 对于系统编程，这种设计可以得到更好的性能并减少回收器（collector）需要了解的项目数。要是在大规模的程序中，这么做导致的差别会非常巨大。


### 2. 说明 writer 和 reader 接口的使用优势
使用 Java 和 Golang 实现HTTP上传文件，保存到硬盘并且计算hash的过程，说明 writer 和 reader 接口的使用优势。
[java-calculate-sha 的代码](https://stackoverflow.com/questions/1741545/java-calculate-sha-256-hash-of-large-file-efficiently)
[java-md5-hashing 的代码/](https://www.javainterviewpoint.com/java-md5-hashing-example/)
[Java Class hierarchy io RandomAccessFile](https://docs.oracle.com/javase/7/docs/api/java/io/package-tree.html)
[Java Class hierarchy MessageDigest](https://docs.oracle.com/javase/7/docs/api/java/security/package-tree.html)
`DigestInputStream, DigestOutputStream, DataInput, DataOutput`


###  1.区分错误与异常的例子[^GolangErrorAndException]
```txt
// 关于知乎回答内容的总结1
    游戏玩家通过购买按钮，用铜钱购买宝石
    胖客户端结构
        胖客户端检查缓存中的铜钱数量，条件允许才发请求到服务端
            铜钱数量足够的时候购买按钮为可用的亮起状态
            铜钱数量不足，购买按钮不可用
        服务端收到请求后检查铜钱数量
            数量足够就进一步完成宝石购买过程
            数量不足就抛出异常，终止请求过程并断开客户端的连接
                当服务端收到不合理的请求（铜钱不足以购买宝石）时，抛出异常比返回错误更为合理，因为这个请求只可能来自两种客户端：外挂或者有BUG的客户端
                如果不通过抛出异常来终止业务过程和断开客户端连接，那么程序的错误就很难被第一时间发现，攻击行为也很难被发现。
    瘦客户端结构
        瘦客户端不会存有太多状态数据和用户数据也不清楚业务逻辑
            用户点击购买按钮，客户端发送购买请求
        服务端收到请求后检查铜钱数量
            数量足够就继续完成业务并返回成功信息
            数量不足就返回错误，提示数量不足
                铜钱不足就变成了业务逻辑范围内的一种失败情况，
                不能提升为异常，否则铜钱不足的用户一点购买按钮都会出错掉线

// 关于知乎回答内容的总结2
错误和异常需要分类和管理，不能一概而论
错误和异常的分类可以以是否终止业务过程作为标准
错误是业务过程的一部分，异常不是
不要随便捕获异常，更不要随便捕获再重新抛出异常
Go语言项目需要把Goroutine分为两类，区别处理异常
在捕获到异常时，需要尽可能的保留第一现场的关键数据
```



[^GolangFAQ]: [Golang相关的常见问题](https://golang.org/doc/faq)
[^GolangFAQ_GC]:[Golang为什么提供GarbageCollection](https://golang.org/doc/faq#garbage_collection)
[^GolangProfiling]:[Profiling Go Programs](https://blog.golang.org/profiling-go-programs)
[^GolangGCJourney]: [Getting to Go: The Journey of Go's Garbage Collector](https://blog.golang.org/ismmkeynote)

[^GolangFAQ_DeclarationsBackwards]: [Golang语法中为什么把变量类型放在最后？](https://golang.org/doc/faq#declarations_backwards)

[^GoAtGoogleEN]: [Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article)
[^GoAtGoogleCN]:[Go在谷歌：以软件工程为目的的语言设计](https://www.oschina.net/translate/go-at-google-language-design-in-the-service-of-software-engineering?cmp)

[^GoClimbingEN]: [Go Ten Years and Climbing](https://commandcenter.blogspot.com/2017/09/go-ten-years-and-climbing.html)
[^GoClimbingCN]: [Go 语言发展史](https://zhuanlan.zhihu.com/p/34263871)

[^TheWayToGo]: [Go 入门指南 The way to go](https://www.kancloud.cn/kancloud/the-way-to-go/72432)
[^BuildWebApplicationWithGolang]: [Build Web Application With Golang](https://github.com/astaxie/build-web-application-with-golang)
[^GolangProgramXSW]: [Go语言编程 许式伟](https://book.douban.com/subject/11577300/)
[^GolangProgramXSWSource]: [Go语言编程 许式伟 源码](https://github.com/qiniu/gobook)

[^GolangProjectIteye]: [iteye.com 列举的Golang相关开源项目](https://www.iteye.com/news/29895)
[^GolangProject]: [官网列举的Golang相关开源项目](https://github.com/golang/go/wiki/Projects)
[^C1000kPracticeGuide]:[C1000kPracticeGuide](https://github.com/xiaojiaqi/C1000kPracticeGuide)


[^GolangInstallEN]: [ 英文： https://golang.org/doc/install](https://golang.org/doc/install)
[^GolangInstallCN]: [ 中文： https://studygolang.com/dl](https://studygolang.com/dl)

[^GolangWebSocket]: [Golang WebSocket](https://github.com/gorilla/websocket)
[^GolangSyntax]: [Gos Declaration Syntax](https://blog.golang.org/gos-declaration-syntax)
[^GolangErrorAndException]: [ZhiHu：Go 语言的错误处理机制是一个优秀的设计吗？](https://www.zhihu.com/question/27158146/answer/44676012)
[^EffectiveGoEN]:[EffectiveGo英文版](https://golang.org/doc/effective_go.html)


[^SevenConcurrencyModelsSevenWeeks]:[七周七并发模型](https://book.douban.com/subject/26337939/)
[^SevenConcurrencyModelsSevenWeeksSource]:[七周七并发中的部分源码](https://zhuanlan.zhihu.com/p/20852001)
[^GoroutinesVSThreads]:[Goroutines VS Threads](http://tleyden.github.io/blog/2014/10/30/goroutines-vs-threads/)
[^ConcurrencyIsNotParallelism]:[Concurrency Is Not Parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
[^WhatIsConcurrencyParallelism]:[ 还在疑惑并发和并行](https://laike9m.com/blog/huan-zai-yi-huo-bing-fa-he-bing-xing)

<!--stackedit_data:
eyJoaXN0b3J5IjpbODYxMzAxNDI5LDU5NTQ1NDk4NCwxMzcyMz
A3NjU2LDYyMTM2NDE2OCwtODQ5NDEyNTM0LC02MjA2MTc1MjNd
fQ==
-->
