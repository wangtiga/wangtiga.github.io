---
layout: post
title:  "[译] Go Module 简介"
date:   2019-09-01 18:00:00 +0800
tags: golang
---

* category
{:toc}


# Introduction to Go Modules[^IntroToGoModule]

# Go Module 简介

Posted on 2018-08-18 by Roberto Selbach

The upcoming version 1.11 of the Go programming language will bring experimental support for modules, a new dependency management system for Go.A few days ago, I wrote a quick post about it. Since that post went live, things changed a bit and as we’re now very close to the new release, I thought it would be a good time for another post with a more hands-on approach.So here’s what we’ll do: we’ll create a new package and then we’ll make a few releases to see how that would work.

golang 1.11 开始支持 `modules`功能，不过是实验性的，还不完善。之前几天我写过一篇快速上手。现在 golang 很快就正式发布，所以很适合再写一篇入门文档。文档中会创建一个 package 并发布几个 release 版本，看看在其他项目中如何引用这个 package 。


## Creating a Module

## 创建一个 module 


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



## Quick Intro to Module Versioning

## Module 版本管理简介


Go modules are versioned, and there are some particularities with regards to certain versions. You will need to familiarize yourself with the concepts behind semantic versioning.More importantly, Go will use repository tags when looking for versions, and some versions are different of others: e.g. versions 2 and greater should have a different import path than versions 0 and 1 (we’ll get to that.)As well, by default Go will fetch the latest tagged version available in a repository. This is an important gotcha as you may be used to working with the master branch.What you need to keep in mind for now is that to make a release of our package, we need to tag our repository with the version. So let’s do that.

go module 是基于版本控制的。所以需要你熟悉[semantic versioning (语义化版本version 语法规则)](https://semver.org/)  。
另外 go 使用 repository tags 来查找版本号。
，有些不同版本之间差异很大，比如 大于 version 2 的版本与 version 0  version 1 之间的 import path 可能都不一样。
默认情况下， go 会 fetch 最新的 tag version。
所以说，如果使用 go module 管理发布 package ，就必须在 git repository 中使用 tag 管理 version。



## Making our first release

## 发布一个 module 版本


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



## Using our module

## 使用 moduel


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



## Making a bugfix release

## 制作一个 bugfix release （补丁版本）


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



## Updating modules

## 更新 module 


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



## Major versions

## 主版本号


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



## Updating to a major version

## 更新主版本号


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



## Tidying it up

## 稍稍整理一下


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



## Vendoring

## 关于 vendor 目录


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

> 译：2019.09.07 go1.3 已经发布，可以执行 `go env -w GOPROXY=https://goproxy.cn,direct` 此命令，直接使用七牛年提供的依赖代理服务。

> 译：对于 go1.3 之前的部分版本，也可以这样启用 GOPROXY `export GOPRXOY=https://goproxy.cn`。
启用 GOPRXOY 后，就这样能在国内网络环境中顺利使用 `golang.org/x/net` 库。

> 译： 非 gomod 目录中，设置了 GOPROXY  时，使用 `go get -u golang.org/x/net` 命令仍然不走代理。需要强制启用 gomod `go env -w GO111MODULE=on`。



## Conclusion

## 总结


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


Posted in
  Go
  Programming

Tagged
  modules
  vgo

Copyright © Roberto Selbach Teixeira. All rights reserved.

[^IntroToGoModule]: [intro to go module](https://roberto.selbach.ca/intro-to-go-modules/)

<!--stackedit_data:
eyJoaXN0b3J5IjpbMTc2MjQ4NzI4Ml19
-->
