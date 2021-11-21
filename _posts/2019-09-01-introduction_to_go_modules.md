---
layout: post
title:  "[译] Go Module 简介 / 试用 Go Module Proxy"
date:   2019-09-01 18:00:00 +0800
tags:   tech
---

* category
{:toc}


# Introduction to Go Modules[^IntroToGoModule] Go Module 简介

Posted on 2018-08-18 by Roberto Selbach

The upcoming version 1.11 of the Go programming language will bring experimental support for modules, a new dependency management system for Go.A few days ago, I wrote a quick post about it. Since that post went live, things changed a bit and as we’re now very close to the new release, I thought it would be a good time for another post with a more hands-on approach.So here’s what we’ll do: we’ll create a new package and then we’ll make a few releases to see how that would work.

golang 1.11 开始支持 `modules`功能，不过是实验性的，还不完善。之前几天我写过一篇快速上手。现在 golang 很快就正式发布，所以很适合再写一篇入门文档。文档中会创建一个 package 并发布几个 release 版本，看看在其他项目中如何引用这个 package 。


### Creating a Module 创建一个 module 


So first things first. Let’s create our package. We’ll call it “testmod”. An important detail here: this directory should be outside your $GOPATH because by default, the modules support is disabled inside it. Go modules is a first step in potentially eliminating $GOPATH entirely at some point.

首先创建一个 `testmod`的 package。注意，必须在$GOPATH目录以外的地方创建，否则，默认是禁用 modules 功能的。

```shell
$ mkdir testmod
$ cd testmod
```


Our package is very simple:

package 很简单，只有一个文件，内容如下：

```go
package testmod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string) string {
   return fmt.Sprintf("Hi, %s", name)
}
```


The package is done but it is still not a module. Let’s change that.

这个 package 完成了，但它还不是 module ，我们要执行下面的操作才行：

```shell
$ go mod init github.com/robteix/testmod
go: creating new go.mod: module github.com/robteix/testmod
```


This creates a new file named go.mod in the package directory with the following contents:

命令执行完毕，此时会在 package 所在目录创建一个文件，内容如下：

```txt
module github.com/robteix/testmod
```


Not a lot here, but this effectively turns our package into a module.We can now push this code to a repository:

短短一行文字，就把我们的 package 就变成 module 了。先把它 push 到仓库里存着吧：

```shell
$ git init 
$ git add * 
$ git commit -am "First commit" 
$ git push -u origin master
```


Until now, anyone willing to use this package would go get it

现在任何人都能执行以下命令使用这个 package 。

```shell
$ go get github.com/robteix/testmod
```


And this would fetch the latest code in master. This still works, but we should probably stop doing that now that we have a Better Way™. Fetching master is inherently dangerous as we can never know for sure that the package authors didn’t make change that will break our usage. That’s what modules aims at fixing.

这条命令会下载 master 分支的最新代码。
稍微思考一下，这样合适吗？
有更好的办法吗？
直接使用的 master 分支的代码不太靠谱，万一 package 作者做了一些不兼容的改动，影响我们使用了，怎么办？
其实 module 就是专门设计出来解决此问题的。



### Quick Intro to Module Versioning 版本管理 Module 简介


Go modules are versioned, and there are some particularities with regards to certain versions. You will need to familiarize yourself with the concepts behind semantic versioning.More importantly, Go will use repository tags when looking for versions, and some versions are different of others: e.g. versions 2 and greater should have a different import path than versions 0 and 1 (we’ll get to that.)As well, by default Go will fetch the latest tagged version available in a repository. This is an important gotcha as you may be used to working with the master branch.What you need to keep in mind for now is that to make a release of our package, we need to tag our repository with the version. So let’s do that.

go module 是基于版本控制的。所以需要你熟悉[semantic versioning (语义化版本version 语法规则)](https://semver.org/)  。
另外 go 使用 repository tags 来查找版本号。
，有些不同版本之间差异很大，比如 大于 version 2 的版本与 version 0  version 1 之间的 import path 可能都不一样。
默认情况下， go 会 fetch 最新的 tag version。
所以说，如果使用 go module 管理发布 package ，就必须在 git repository 中使用 tag 管理 version。



### Making our first release 发布一个 module 版本


Now that our package is ready, we can release it to the world. We do this by using version tags. Let’s release our version 1.0.0:

package 已经准备完毕，我们可以把他发布到互联网上了。使用 version tags 先发布一个 1.0.0 的版本吧。

```shell
$ git tag v1.0.0
$ git push --tags
```


This creates a tag on my Github repository marking the current commit as being the release 1.0.0.Go doesn’t enforce that in any way, but a good idea is to also create a new branch (“v1”) so that we can push bug fixes to.

以上命令会在 Gihub 仓库中创建一个名为 v1.0.0 的 tag ，这就把当前 commit 作为 v1.0.0 版本发布出去了。
当然 Go 不强制使用 tag ，也可以创建一个 branch ("v1") 发布版本，这样还能向 branch 中 push 一些 bugfix （补丁）代码。命令如下所示：

```shell
$ git checkout -b v1
$ git push -u origin v1
```


Now we can work on master without having to worry about breaking our release.

使用 tag 或 branch 发布版本后，就可以放心大胆的向 master 提交代码了。

- git tag 用来给版本打标记
- git branch 用来给版本做补丁



### Using our module 使用 moduel


Now we’re ready to use the module. We’ll create a simple program that will use our new package:

下面演示如何使用刚才创建的 module 。
先创建一个简单的程序，其中调用 testmod 的 Hi 方法，代码如下：

```go
package main

import (
    "fmt"

    "github.com/robteix/testmod"
)

func main() {
    fmt.Println(testmod.Hi("roberto"))
}
```


Until now, you would do a `go get github.com/robteix/testmod` to download the package, but with modules, this gets more interesting. First we need to enable modules in our new program.

增加一个使用 testmod package 的 golang 代码
在以前没有 go module 功能的时候，必须使用`go get github.com/robteix/testmod`把引用的 testmod package 下载到本地使用。
但是我们按下面的方法启用 module 功能，就会方便很多。
先在我们的程序中启用 module 吧。

```shell
$ go mod init mod
```


As you’d expect from what we’ve seen above, this will have created a new go.mod file with the module name in it:

跟之前一样，这条命令创建了一个名为 `mod` 的 module ，同时生成了一个 go.mod 文件，内容如下：

```txt
module mod
```


Things get much more interesting when we try to build our new program:

这时尝试编译这个程序，就会发现一些变化。

```shell
$ go build
go: finding github.com/robteix/testmod v1.0.0
go: downloading github.com/robteix/testmod v1.0.0
```


As we can see, the go command automatically goes and fetches the packages imported by the program. If we check our go.mod file, we see that things have changed:

如上所示， go 命令自动下载了当前程序引用的 package 。
再次检查 go.mod 文件的内容，也比刚才多了一点内容。

```txt
module mod
require github.com/robteix/testmod v1.0.0
```


And we now have a new file too, named go.sum, which contains hashes of the packages, to ensure that we have the correct version and files.

另外，还比刚才多了一个 go.sum 文件，其中包含 package 的 hash （摘要）值，用于确保 package 的版本和文件内容都是正确的。

```txt
github.com/robteix/testmod v1.0.0 h1:9EdH0EArQ/rkpss9Tj8gUnwx3w5p0jkzJrd5tRAhxnA=
github.com/robteix/testmod v1.0.0/go.mod h1:UVhi5McON9ZLc5kl5iN2bTXlL6ylcxE9VInV71RrlO8=
```

译：
go.sum 文件存在的目的是保证在不同主机，不现网络环境中执行`go build`时下载的代码内容都是一致的。
即使因为有人更改了相同版本号的 package 代码的内容，我们也能通过 go.sum 文件内容的变化察觉出不一致。

译：
简要总结，golang 会在编译时，自动查找import的外部程序，下载package的代码，记录版本信息及版本对应的 hash 值。



### Making a bugfix release 制作一个 bugfix release （补丁版本）


Now let’s say we realized a problem with our package: the greeting is missing ponctuation! People are mad because our friendly greeting is not friendly enough. So we’ll fix it and release a new version:

我们发现刚才做的 testmod package 有个问题：问候语后面缺少一个标点符号。万一有人因为我们的欢迎辞词不够友善而生气就不好了。所以赶紧修复这个 bug 并发布一个新版本吧。

```go
// Hi returns a friendly greeting
func Hi(name string) string {
-       return fmt.Sprintf("Hi, %s", name)
+       return fmt.Sprintf("Hi, %s!", name)
}
```


We made this change in the v1 branch because it’s not relevant for what we’ll do for v2 later, but in real life, maybe you’d do it in master and then back-port it. Either way, we need to have the fix in our v1 branch and mark it as a new release.

我们直接在 v1 branch 上进行修复，因为这点改动跟我们以后要开发的 v2 版本完全不相关。
当然，实际上，你也可以直接在 master 分支上修复，然后花时间合并到某个分支中。
不论怎么，我们都必须在 v1 branch 中修复，并发布一个新版本。

```shell
$ git commit -m "Emphasize our friendliness" testmod.go
$ git tag v1.0.1
$ git push --tags origin v1
```

译：在 testmod package 修复bug并且发布新的 tags 版本



### Updating modules 更新 module 


By default, Go will not update modules without being asked. This is a Good Thing™ as we want predictability in our builds. If Go modules were automatically updated every time a new version came out, we’d be back in the uncivilized age pre-Go1.11. No, we need to tell Go to update a modules for us.We do this by using our good old friend go get:

默认情况 Go （go build 时）不会自动更新 module 。
这样能保证我们每次都能顺利 build 项目。
如果 go module 每次都自动把项目引用的第三方 package 更新到最新版本，那就跟 Go1.11 之前的同什么两样了（太容易因为三方版本代码更新而出问题）。
执行下面的命令，就能按需更新 module 。


- run `go get -u` to use the latest minor or patch releases (i.e. it would update from 1.0.0 to, say, 1.0.1 or, if available, 1.1.0)

- run `go get -u=patch` to use the latest patch releases (i.e., would update to 1.0.1 but not to 1.1.0)

- run `go get package@version` to update to a specific version (say, github.com/robteix/testmod@v1.0.1)


- run `go get -u` 更新到 latest minor or patch releases (从 1.0.0 更新到 1.0.1 或  1.1.0)

- run `go get -u=patch` 更新到 latest patch releases (从 1.0.1 更新到 1.1.0)

- run `go get package@version` 更新到指定版本 (say, github.com/robteix/testmod@v1.0.1)

> run `go get package@commitID` 更新到指定 commit 的版本 (比如，github.com/gin-gonic/gin@393a63f3b020df89d42695064443760c7d0a0dc8)
解决[github.com/ugorji/go/codec 冲突的问题](https://cloud.tencent.com/developer/article/1417112) [ugorj issues279](https://github.com/ugorji/go/issues/279)
```shell
build xxxxxxx: cannot load github.com/ugorji/go/codec: ambiguous import: found github.com/ugorji/go/codec in multiple modules:

        github.com/ugorji/go v1.1.4 (/home/dodo/go/pkg/mod/github.com/ugorji/go@v1.1.4/codec)

        github.com/ugorji/go/codec v0.0.0-20181022190402-e5e69e061d4f (/home/dodo/go/pkg/mod/github.com/ugorji/go/codec@v0.0.0-20181022190402-e5e69e061d4f)
```

In the list above, there doesn’t seem to be a way to update to the latest major version. There’s a good reason for that, as we’ll see in a bit.Since our program was using version 1.0.0 of our package and we just created version 1.0.1, any of the following commands will update us to 1.0.1:

上面这几条命令都不能更新到 latest major version 。
具体原因一会儿解释。
在我们的 usetestmod 项目中，使用的是 testmod v1.0.0 版本。
而 testmod 的最新版本是 1.0.1 ，所以无论执行上面哪个命令，都会更新到 v1.0.1 的版本

```shell
$ go get -u
$ go get -u=patch
$ go get github.com/robteix/testmod@v1.0.1
```


After running, say, `go get -u ` our go.mod is changed to:

运行 `go get -u` 命令后，go.mod 文件内容就变成下面这样：

```txt
module mod
require github.com/robteix/testmod v1.0.1
```



### Major versions 主版本号


According to semantic version semantics, a major version is different than minors. Major versions can break backwards compatibility. From the point of view of Go modules, a major version is a different package completely. This may sound bizarre at first, but it makes sense: two versions of a library that are not compatible with each other are two different libraries.Let’s make a major change in our package, shall we? Over time, we realized our API was too simple, too limited for the use cases of our users, so we need to change the Hi() function to take a new parameter for the greeting language:

语义化版本规则中， major version （主版本号）与 minor version （次版本号）区别很大。
Major version 改变后，程序不再兼容低版本（丢失了向后兼容性）。
对 Go module 来说，不同 major version 的 package 完全可以当做不同的 package 看待了。
这听起来有些怪，但很好理解：两个版本不同版本，而且互相不兼容的库，自然可以看成做两个完全不同的库。


要不，我们在自己的 testmod 库里试着做一次 major version 的变动？
比方说，一段时间后，我们发现这个 package 中的 API 太简单了，限制了用户的使用场景。因此，我们决定给 `Hi()` 函数增加一个参数，用于指定欢迎辞的语音：


```go
package testmod

import (
    "errors"
    "fmt" 
) 

// Hi returns a friendly greeting in language lang
func Hi(name, lang string) (string, error) {
    switch lang {
    case "en":
        return fmt.Sprintf("Hi, %s!", name), nil
    case "pt":
        return fmt.Sprintf("Oi, %s!", name), nil
    case "es":
        return fmt.Sprintf("¡Hola, %s!", name), nil
    case "fr":
        return fmt.Sprintf("Bonjour, %s!", name), nil
    default:
        return "", errors.New("unknown language")
    }
}
```


Existing software using our API will break because they (a) don’t pass a language parameter and (b) don’t expect an error return. Our new API is no longer compatible with version 1.x so it’s time to bump the version to 2.0.0.I mentioned before that some versions have some peculiarities, and this is the case now. Versions 2 and overshould change the import path. They are different libraries now.We do this by appending a new version path to the end of our module name.

（经过以上改动后）之前已经使用了旧 API 的软件，肯定无法编译通过，因为他们既没传 lang 参数，也没处理返回的 error 参数。
新的 API 无法与 v1.x 的版本兼容，所以应该把版本号升级到 2.0.0 。
这就是之前提到的应该升级 major version 的变动。
v2.x 还应该调整 import path （包导入路径）。
他们已经是两个不同的 library 库了。
我们只需要在 module 名后加个版本标识，就能调整 import path 。


```txt
module github.com/robteix/testmod/v2
```


The rest is the same as before, we push it, tag it as v2.0.0 (and optionally create a v2 branch.)

剩下的工作和之前一样，push 到代码仓库中，打一个 v2.0.0 的 tag （或者创建一个 v2 的 branch 也行）。

```shell
$ git commit testmod.go -m "Change Hi to allow multilang"
$ git checkout -b v2 # optional but recommended
$ echo "module github.com/robteix/testmod/v2" > go.mod
$ git commit go.mod -m "Bump version to v2"
$ git tag v2.0.0
$ git push --tags origin v2 # or master if we don't have a branch
```



### Updating to a major version 更新主版本号


Even though we have released a new incompatible version of our library, existing software will not break, because it will continue to use the existing version 1.0.1. `go get -u` will not get version 2.0.0.At some point, however, I, as the library user, may want to upgrade to version 2.0.0 because maybe I was one of those users who needed multi-language support.I do it but modifying my program accordingly:

虽然我们给自己的 library 发布了一个新的不兼容版本，那些使用旧版本的软件也不受任何影响。因为它们编译的时候还会使用旧版本的代码。
`go get -u` 命令也不会获取到 v2.0.0 的代码。
但也有一些 library 的使用者希望升级到 v2.0.0 的代码，因为他可能需要用到多语言支持的功能。
对这些用户，可以按下面的方法修改他们的程序。

```go
package main

import (
    "fmt"
    "github.com/robteix/testmod/v2" 
)

func main() {
    g, err := testmod.Hi("Roberto", "pt")
    if err != nil {
        panic(err)
    }
    fmt.Println(g)
}
```


And then when I run `go build`, it will go and fetch version 2.0.0 for me. Notice how even though the import path ends with “v2”, Go will still refer to the module by its proper name (“testmod”).As I mentioned before, the major version is for all intents and purposes a completely different package. Go modules does not link the two at all. That means we can use two incompatible versions in the same binary:

此时运行`go build`，就会自动下载使用 v2.0.0 的代码了。
注意，虽然即使 import path 以 v2 结尾， Go 还会使用正确的 module 名称，即 testmod 。
我之前说过， major version 升级时，package 已经是个焕然一新的 package 了。
Go module 也不会把新旧 package 当成一个使用的。
所以，我们其实也能在一个程序里同时使用新旧两个版本的 package 。


```go
package main

import (
    "fmt"
    "github.com/robteix/testmod"
    testmodML "github.com/robteix/testmod/v2"
)

func main() {
    fmt.Println(testmod.Hi("Roberto"))
    g, err := testmodML.Hi("Roberto", "pt")
    if err != nil {
        panic(err)
    }
    fmt.Println(g)
}
```


This eliminates a common problem with dependency management: when dependencies depend on different versions of the same library.

以上方法解决了依赖管理中常见的一个问题，即，如何处理依赖相同 library 的不同版本。



### Tidying it up 稍稍整理一下


Going back to the previous version that uses only testmod 2.0.0, if we check the contents of go.mod now, we’ll notice something:

回到之前仅使用了 testmod v2.0.0 的版本，我们会发现 go.mod 文件的内容变成下面的样子：

```txt
module mod
require github.com/robteix/testmod v1.0.1
require github.com/robteix/testmod/v2 v2.0.0
```


By default, Go does not remove a dependency from go.mod unless you ask it to. If you have dependencies that you no longer use and want to clean up, you can use the new tidy command:

默认情况 Go 不会在 go.mode 文件中自动移除 dependency （依赖），必须手动修改。
可以使用 `go mod tidy`命令自动移除已经不再使用的 dependency 。

```shell
$ go mod tidy
```


Now we’re left with only the dependencies that are really being used.

现在，我们只剩下一些实际引用到的 dependency 了。



### Vendoring 关于 vendor 目录


Go modules ignores the `vendor/` directory by default. The idea is to eventually do away with vendoring1. But if we still want to add vendored dependencies to our version control, we can still do it:

开启 Go module 功能后，默认忽略 `vendor/` 目录。
因为 module 已经能代替 vendor 的功能。
但如果硬要把 go.mod 中的依赖添加到 vendor 目录，然后提交放到自己的代码仓库中管理，可以执行下面的命令。

TODO 原文中使用 vendoring1 ，是不是引用了什么网址

```shell
$ go mod vendor
```


This will create a `vendor/` directory under the root of your project containing the source code for all of your dependencies.Still, go build will ignore the contents of this directory by default. If you want to build dependencies from the vendor/ directory, you’ll need to ask for it.

这会在当前项目的根目录中创建一个 `vendor/` 目录，其中包含所有依赖。
当然，因为启用了 go mod ,所以`go build`时会忽略 `vendor/` 目录的代码。
如果想使用 vendor 目录的代码编译程序，可以执行下面的命令：

```shell
$ go build -mod vendor
```


I expect many developers willing to use vendoring will run go build normally on their development machines and use -mod vendor in their CI.
Again, Go modules is moving away from the idea of vendoring and towards using a Go module proxy for those who don’t want to depend on the upstream version control services directly.
There are ways to guarantee that go will not reach the network at all (e.g. GOPROXY=off) but these are the subject for a future blog post.

我估计会有很多开发者使用这种方案，保持开发环境与CI持续集成环境中所引用的第三方库代码一致。
但请注意，Go module 相比 vendor 前进了一大步，我们完全可以使用 Go module proxy 的特性来避免过于依赖第三方提供的源码版本控制服务（译：指 Github Gitlag 等）。
还有多种方法可以让 go 完全不访问网络，就能下载到三方依赖库，编译我们的程序（比如 GOPROXY=off）。
但有关这些主题，会在以后中的博客中讨论。

> 译：
这段没完全理解。
自己搭建一个 GOPROXY 服务，应该就能自己维护一个第三方依赖库，减少对 github.com 等第三方代码仓库中的依赖了吧。
那么作者说的比 vendor 更好的是哪些解决方法呢？
会是 go.mod 中使用 replace 指令吗？

> 译：2019.09.07 go1.13 已经发布，可以执行 `go env -w GOPROXY=https://goproxy.cn,direct` 此命令，直接使用七牛年提供的依赖代理服务。

> 译：对于 go1.13 之前的部分版本，也可以这样启用 GOPROXY `export GOPRXOY=https://goproxy.cn`。
启用 GOPRXOY 后，就这样能在国内网络环境中顺利使用 `golang.org/x/net` 库。

> 译：[go1.16](https://golang.org/doc/go1.16#go-command) 之前版本中，在非 gomod 目录，即使设置了 GOPROXY 时，使用 `go get -u golang.org/x/net` 命令时仍然不会走代理下载源码。需要强制启用 gomod ，即 `go env -w GO111MODULE=on`。



### Conclusion 总结


This post may seem a bit daunting, but I tried to explain a lot of things together. The reality is that now Go modules is basically transparent. We import package like always in our code and the go command will take care of the rest.When we build something, the dependencies will be fetched automatically. It also eliminates the need to use $GOPATH which was a roadblock for new Go developers who had trouble understanding why things had to go into a specific directory.Vendoring is (unofficially) being deprecated in favour of using proxies.1 I may do a separate post about the Go module proxy. (Update: it’s live.)

这篇文章可能有点长，可能是因为我把很多概念都放在一块说明了。
实际上是，使用 Go module 的过程，对用户来说基本是透明的（比较易用）。
我们在代码中 import 一个 package ，剩下的整改都交给 go 命令来完成了。
当我们执行`go build`编译时，相关依赖会自动下载。
另外，它还减少了 Go 语言初学者阻塞，以前的版本，我们必须了解 `$GOPATH` 这个全局变更，理解其工作原理，并设置到一个正常的目录。
vendor （非官方）已经被弃用，现在尽量可以使用 proxy 管理依赖吧代替。
我以后还会单独写一篇文章介绍 Go module proxy.

译：怎么用 proxy 代替 vendor ？自建Proxy服务就有点复杂了吧。待翻译 go module proxy 时了解吧。

TODO 原文中 proxies 1 这里应该也引用了一个网址


I think this came out a bit too strong and people left with the impression that vendoring is being removed right now. It isn’t. Vendoring still works, albeit slightly different than before. There seems to be a desire to replace vendoring with something better, which may or may not be a proxy. But for now this is just it: a desire for a better solution. Vendoring is not going away until a good replacement is found (if ever.) ↩ ↩

可能有些人会觉得 go module 这么强大，现在应该就能把 vendor 功能移除掉了吧。
并不行。
vendor 还可以正常使用，虽然使用方式发生了一点变化。
我们希望有一个更好的方式替代 vendor ，也许是 proxy ，也许不是。
但这仅仅是一个希望。在更好的方式出现前，vendor 肯定会一直存在。

Copyright © Roberto Selbach Teixeira. All rights reserved.

[^IntroToGoModule]: [intro to go module](https://roberto.selbach.ca/intro-to-go-modules/)




# Playing with Go module proxies[^PlayingWithGoModuleProxies] 试用 Go module proxy


Posted on 2018-08-29 by Roberto Selbach

I wrote a  [brief introduction to Go modules](https://roberto.selbach.ca/intro-to-go-modules/)  and in it I talked briefly about Go modules proxies and now that  [Go 1.11 is out](https://blog.golang.org/go1.11), I thought I’d play a bit these proxies to figure our how they’re supposed to work.


### Why 为什么

One of the goals of Go modules is to provide reproducible builds and it does a very good job by fetching the correct and expected files from a repository.

Go Module 的目标之一是，通过从代码仓库下载正确的依赖代码，提供一个稳定、可重用的编译环境。


But what if the  [servers](https://blog.github.com/2016-02-03-january-28th-incident-report/)  [are](https://blog.bitbucket.org/2012/09/19/post-mortem-on-our-availability-earlier-today/)  [offline](https://about.gitlab.com/2017/02/10/postmortem-of-database-outage-of-january-31/)? 
What if the repository  [simply vanishes](https://www.theregister.co.uk/2016/03/23/npm_left_pad_chaos/)?

如果代码仓库服务器故障了怎么办？[^Github] [^Bitbucket] [^Gitlab]
如果代码仓库消失了怎么办？[^NPM]

One way teams deal with these risks is by vendoring the dependencies, which is fine. But Go modules offers another way: the use of a module proxy.

可以用 `vendor` 目录管理依赖代码，来防止这种风险。
但 Go module 提供了另一种办法： 使用 module proxy 。



### The Download Protocol 下载协议

When Go modules support is enabled and the  `go`  command determines that it needs a module, it first looks at the local cache (under  `$GOPATH/pkg/mods`). If it can’t find the right files there, it then goes ahead and fetches the files from the network (i.e. from a remote repo hosted on Github, Gitlab, etc.)

启用 Go module 后，如果 `go` 命令需要下载使用某个 module ，它首先在本地缓存( `$GOPATH/pkg/mods` 目录)中查询。
如果没找到到相关代码，就会自动从网络中下载（比如 Github, Gitlab 等代码仓库）。


If we want to control what files  `go`  can download, we need to tell it to go through our proxy by setting the  `GOPROXY`  environment variable to point to our proxy’s URL. For instance:

如果要控制 `go` 从哪里下载代码， 可以设置 `GOPROXY` 环境变量，让 `go` 通过指定的代理服务器下载代码。比如下面这样：

```shell
export GOPROXY=http://gproxy.mycompany.local:8080
```


The proxy is nothing but a web server that responds to the module download protocol, which is a very simple API to query and fetch modules. The web server may even serve static files.

proxy 就是响应 module 下载协议的接口服务器。
下载协议的 API 也很简单，主要用于查询和下载 module 代码。
甚至用静态文件服务器实现这个接口服务器就行。


A typical scenario would be the  `go`  command trying to fetch  `github.com/pkg/errors`:

使用 `go` 命令下载 `github.com/pkg/errors` 的流程如下：

![](https://roberto.selbach.ca/wp-content/uploads/2018/08/goproxyseq.png)

```seq
go    ->> proxy : `Get /github.com/pkg/errors/@v/list`
proxy ->> go    : `<list of available versions>`
go    ->> proxy : `Get /github.com/pkg/errors/@v/v0.8.0.info`
proxy ->> go    : `{"Version": "v0.8.0", "Time": "2018-08-27T08:54:46.436183-04:00"}`
go    ->> proxy : `Get /github.com/pkg/errors/@v/v0.8.0.zip`
proxy ->> go    : `<bytes of zip archive containing module files>`
```

> 译：补充一个 go1.12 执行 `go get` 的过程

```shell
$ go version 
go version go1.12.5 windows/amd64

$ go get -insecure github.com/pkg/errors
```

以上 go get 命令会外发几个 HTTP 请求下载 pkg/errrors 的代码

经抓包发现，外发的请求就是以下两个

```shell
$ curl -k -i --raw -o 0.dat "https://github.com/pkg/errors/info/refs?service=git-upload-pack" -H "Host: github.com" -H "User-Agent: git/2.11.1.windows.1" -H "Accept: */*" -H "Accept-Encoding: gzip" -H "Accept-Language: C, *;q=0.9" -H "Pragma: no-cache"
$ curl -k -i --raw -o 1.dat -X POST "https://github.com/pkg/errors/git-upload-pack" -H "Host: github.com" -H "User-Agent: git/2.11.1.windows.1" -H "Accept-Encoding: gzip" -H "Content-Type: application/x-git-upload-pack-request" -H "Accept: application/x-git-upload-pack-result"
```


The first thing  `go`  will do is ask the proxy for a list of available versions. It does this by making a  `GET`  request to  `/{module name}/@v/list`. The server then responds with a simple list of versions it has available:

`go` 首先发送 `GET /{module name}/@v/list` 请求 proxy 获取可用的版本。

然后，服务器返回一个可用的版本列表。

```txt
v0.8.0
v0.7.1
```


The  `go`  will determine which version it wants to download — the latest unless explicitly told otherwise[1](https://roberto.selbach.ca/go-proxies/#fn-1943-0). It will then request information about that given version by issuing a  `GET`  request to  `/{module name}/@v/{module revision}`  to which the server will reply with a JSON representation of the  `struct`:

`go` 命令会选择它需要的版本下载 - 如果没有[显式指定版本](https://roberto.selbach.ca/go-proxies/#fn-1943-0) 就会获取当前最新版本。
然后发送  `GET /{module name}/@v/{module revision}` 获取指定版本的相关信息。
服务器返回以下格式的 JSON 格式字符串。

```go
type RevInfo struct {
    Version string    // version string
    Name    string    // complete ID in underlying repository
    Short   string    // shortened ID, for use in pseudo-version
    Time    time.Time // commit time
}
```

So for instance, we might get something like this:

实际收到的数据格式如下所示：

```json
{
    "Version": "v0.8.0",
    "Name": "v0.8.0",
    "Short": "v0.8.0",
    "Time": "2018-08-27T08:54:46.436183-04:00"
}
```

The  `go`  command will then request the module’s  `go.mod`  file by making a  `GET`  request to  `/{module name}/@v/{module revision}.mod`. The server will simply respond with the contents of the  `go.mod`  file (e.g.  `module github.com/pkg/errors`.) This file may list additional dependencies and the cycle restarts for each one.

然后，`go` 命令会发送 `GET /{module name}/@v/{module revision}.mod` 下载 `go.mod` 文件。
服务端收到请求后，会返回对应版本的 `go.mod` 文件内容 (比如 `module github.com/pkg/errors`) 。
文件中可能包含一些额外的依赖，`go` 会自动对这些依赖执行类似的步骤。


Finally, the  `go`  command will request the actual module by getting  `/{module name}/@v/{module revision}.zip`. The server should respond with a byte blob (`application/zip`) containing a zip archive with the module files where each file  _must_  be prefixed by the full module path  _and_  version (e.g.  `github.com/pkg/errors@v0.8.0/`), i.e. the archive should contain:

最终，`go` 命令会请求 `/{module name}/@v/{module revision}.zip` 下载实际的 module 内容。
服务端返回一个 zip 打包的二进制文件 (`application/zip`) 。
压缩包内的 module 文件 _必须_ 以完整的 module 路径和 _版本_ 为前缀 (比如  `github.com/pkg/errors@v0.8.0/`)，示例如下：

```txt
github.com/pkg/errors@v0.8.0/example_test.go
github.com/pkg/errors@v0.8.0/errors_test.go
github.com/pkg/errors@v0.8.0/LICENSE
...
```

And  _not_:

_不能_ 像下面这样：

```txt
errors/example_test.go
errors/errors_test.go
errors/LICENSE
...
```


This seems like a lot when written like this, but it’s in fact a very simple protocol that simply fetches 3 or 4 files:

说得好像有点复杂，实际上这个协议很简单，只是请求了 3 到 4 个文件：


- 1.The list of versions (only if  `go`  does not already know which version it wants)
- 2.The module metadata
- 3.The  `go.mod`  file
- 4.The module zip itself

- 1.当前版本列表 （仅在 `go` 不确定自己需要哪个版本时，才会请求）
- 2.module 的元数据 （译：即此 module 版本的详细描述信息）
- 3.`go.mod` 文件
- 4.module 的 zip 压缩包



### Creating a simple local proxy 创建一个简单的代理

To try out the proxy support, let’s create a very basic proxy that will serve static files from a directory. First we create a directory where we will store our in-site copies of our dependencies. Here’s what I have in mine:

要实现一个简单的代理，使用一个静态文件服务器就可以（译：直接返回指定目录的静态文件的服务器，类似 nginnx ）。
所以，先要创建一个目录，保存我们依赖的那些静态文件就行。
就是下面这样：

```txt
$ find . -type f
./github.com/robteix/testmod/@v/v1.0.0.mod
./github.com/robteix/testmod/@v/v1.0.1.mod
./github.com/robteix/testmod/@v/v1.0.1.zip
./github.com/robteix/testmod/@v/v1.0.0.zip
./github.com/robteix/testmod/@v/v1.0.0.info
./github.com/robteix/testmod/@v/v1.0.1.info
./github.com/robteix/testmod/@v/list
```


These are the files our proxy will serve. You can  [find these files on Github](https://github.com/robteix/go-proxy-blog)  if you’d like to play along. For the examples below, let’s assume we have a  `devel`  directory under our home directory; adapt accordingly.


以上这些文件都会由我们的 proxy 返回。
如果你也想把 proxy 运行起来看看效果，可以在  [Github](https://github.com/robteix/go-proxy-blog) 下载代码。
使用以下命令将 proxy 下载到 home 目录的 devel 文件夹中。

```shell
$ cd $HOME/devel
$ git clone https://github.com/robteix/go-proxy-blog.git
```

Our proxy server is simple (it could be even simpler, but I wanted to log the requests):

这个 proxy 服务很简单（如果去掉那些记录请求日志的代码，它还能更简单）：

```go
package main

import (
    "flag"
    "log"
    "net/http"
)

func main() {
    addr := flag.String("http", ":8080", "address to bind to")
    flag.Parse()

    dir := "."
    if flag.NArg() > 0 {
        dir = flag.Arg(0)
    }

    log.Printf("Serving files from %s on %s\n", dir, *addr)

    h := handler{http.FileServer(http.Dir(dir))}

    panic(http.ListenAndServe(*addr, h))
}

type handler struct {
    h http.Handler
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println("New request:", r.URL.Path)
    h.h.ServeHTTP(w, r)
}
```

Now run the code above:

执行下面的命令，运行以上代码：

```go
$ go run proxy.go -http :8080 $HOME/devel/go-proxy-blog
2018/08/29 14:14:31 Serving files from /home/robteix/devel/go-proxy-blog on :8080
```

```go
$ curl http://localhost:8080/github.com/robteix/testmod/@v/list
v1.0.0
v1.0.1
```

Leave the proxy running and move to a new terminal. Now let’s create a new test program. we create a new directory  `$HOME/devel/test`  and create a file named  `test.go`  inside it with the following code:

让 proxy 保持运行，我们再打开一个新的终端窗口，创建一个测试程序。
创建一个目录 `$HOME/devel/test` 并在其中增加一个 `test.go` 文件，文件内容如下：

```go
package main

import (
    "github.com/robteix/testmod"
)

func main() {
    testmod.Hi("world")
}
```


And now, inside this directory, let’s enable Go modules:

在这个目录中，我们启用 Go module ：

```shell
$ go mod init test
```


And we set the  `GOPROXY`  variable:

然后设置 `GOPROXY` 环境变量：

```shell
export GOPROXY=http://localhost:8080
```


Now let’s try building our new program:

编译我们的程序：

```shell
$ go build
go: finding github.com/robteix/testmod v1.0.1
go: downloading github.com/robteix/testmod v1.0.1
```

And if you check the output from our proxy:

这里，你会看到我们的 proxy 服务输出以下日志：

```log
2018/08/29 14:56:14 New request: /github.com/robteix/testmod/@v/list
2018/08/29 14:56:14 New request: /github.com/robteix/testmod/@v/v1.0.1.info
2018/08/29 14:56:14 New request: /github.com/robteix/testmod/@v/v1.0.1.mod
2018/08/29 14:56:14 New request: /github.com/robteix/testmod/@v/v1.0.1.zip
```

So as long as  `GOPROXY`  is set,  `go`  will only download files from out proxy. If I go ahead and delete the repository from Github, things will continue to work.

因为刚才设置了 `GOPROXY` ， `go` 程序会自动从我们的 proxy 中下载相关依赖。
这时，如果我去 Github 删除了 testmod 仓库，编译过程不会受到任何影响。



### Using a local directory 使用 local directory 

It is interesting to note that we don’t even need our  `proxy.go`  at all. We can set  `GOPROXY`  to point to a directory in the filesystem and things will still work as expected:

更有意思的是，我们甚至完全不需要 `proxy.go` 。
直接把 `GOPROXY` 设置为本地文件系统的路径，编译过程也能正常进行。

```shell
export GOPROXY=file://home/robteix/devel/go-proxy-blog
```

If we do a  `go build`  now[2](https://roberto.selbach.ca/go-proxies/#fn-1943-1), we’ll see  _exactly_  the same thing as with the proxy:

这时执行 `go build` ，也能正常编译：

```shell
$ go build
go: finding github.com/robteix/testmod v1.0.1
go: downloading github.com/robteix/testmod v1.0.1
```

Of course, in real life, we probably will prefer to have a company/team proxy server where our dependencies are stored, because a local directory is not really much different from the local cache that  `go`  already maintains under  `$GOPATH/pkg/mod`, but still, nice to know that it works.

当然，实际开发过程中，我们肯定会使用 公司/团队 统一配置的 GoProxy 服务保存依赖库代码。
因为把 GOPROXY 设置成本地目录与 `go` 命令在 `$GOPATH/pkg/mod` 维护的缓存效果差不多。
知道这其中的工作过程，也没什么坏处。

There is a  [project called Athens](https://github.com/gomods/athens)  that is building a proxy and that aims — if I don’t misunderstand it — to create a central repository of packages à la npm.

有一个  [名为 Athens 的项目](https://github.com/gomods/athens) 实现了一个 proxy ，它们的目标是创建一个类似 npm 的集中化 package 代码仓库。

----------

1.  Remember that  `somepackage`  and  `somepackage/v2`  are treated as different packages. [↩](https://roberto.selbach.ca/go-proxies/#fnref-1943-0)

2.  That’s not strictly true as now that we’ve already built it once,  `go`  has cached the module locally and will not go to the proxy (or the network) at all. You can still force it by deleting  `$GOPATH/pkg/mod/cache/download/github.com/robteix/testmod/`  and  `$GOPATH/pkg/mod/github.com/robteix/testmod@v1.0.1`) [↩](https://roberto.selbach.ca/go-proxies/#fnref-1943-1)

Posted in
 [Go](https://roberto.selbach.ca/category/golang/)
 [Programming](https://roberto.selbach.ca/category/programming/)

Tagged
 [modules](https://roberto.selbach.ca/tag/modules/) 
 [vgo](https://roberto.selbach.ca/tag/vgo/)

Copyright © Roberto Selbach Teixeira. All rights reserved.

### 说明

- TODO module 下载的源代码不在 GOPATH 了，那么影响使用 vim-go 等插件阅读源码吗？即， gopls 这些源码分析软件是不是也要更新？

- 扩展阅读 [干货满满的 Go Modules 和 goproxy.cn](https://mp.weixin.qq.com/s/AsdCDodxZFxs2SkhSwOvpg)


[^Github]: [Github Incident](https://blog.github.com/2016-02-03-january-28th-incident-report/) 

[^Bitbucket]: [Bitbucket Post Mortem](https://blog.bitbucket.org/2012/09/19/post-mortem-on-our-availability-earlier-today/) 

[^Gitlab]: [Gitlab Post Mortem](https://about.gitlab.com/2017/02/10/postmortem-of-database-outage-of-january-31/)? 

[^NPM]: [npm Vanishes](https://www.theregister.co.uk/2016/03/23/npm_left_pad_chaos/)?

[^PlayingWithGoModuleProxies]: [Playing with Go module proxies](https://roberto.selbach.ca/go-proxies/)


