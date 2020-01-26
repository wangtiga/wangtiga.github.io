---
layout: post
title:  "[译] 试用 Go Module Proxy"
date:   2019-09-17 18:17:00 +0800
tags: golang
---

* category
{:toc}


# Playing with Go module proxies[^PlayingWithGoModuleProxies]

# 试用 Go module proxy


Posted on  [2018-08-29](https://roberto.selbach.ca/go-proxies/) by  [Roberto Selbach](https://roberto.selbach.ca/author/robteix/)

I wrote a  [brief introduction to Go modules](https://roberto.selbach.ca/intro-to-go-modules/)  and in it I talked briefly about Go modules proxies and now that  [Go 1.11 is out](https://blog.golang.org/go1.11), I thought I’d play a bit these proxies to figure our how they’re supposed to work.


# Why

# 为什么

One of the goals of Go modules is to provide reproducible builds and it does a very good job by fetching the correct and expected files from a repository.

Go Module 的目标之一是，通过从代码仓库下载正确的依赖代码，提供一个稳定、可重用的编译环境。


But what if the  [servers](https://blog.github.com/2016-02-03-january-28th-incident-report/)  [are](https://blog.bitbucket.org/2012/09/19/post-mortem-on-our-availability-earlier-today/)  [offline](https://about.gitlab.com/2017/02/10/postmortem-of-database-outage-of-january-31/)? 
What if the repository  [simply vanishes](https://www.theregister.co.uk/2016/03/23/npm_left_pad_chaos/)?

如果代码仓库服务器故障了怎么办？[^Github] [^Bitbucket] [^Gitlab]
如果代码仓库消失了怎么办？[^NPM]

One way teams deal with these risks is by vendoring the dependencies, which is fine. But Go modules offers another way: the use of a module proxy.

可以用 `vendor` 目录管理依赖代码，来防止这种风险。
但 Go module 提供了另一种办法： 使用 module proxy 。



# The Download Protocol

# 下载协议

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

# 以上 go get 命令会外发几个 HTTP 请求下载 pkg/errrors 的代码
# 经抓包发现，外发的请求就是以下两个
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



# Creating a simple local proxy

# 创建一个简单的代理

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



# Using a local directory

# 使用 local directory 

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

## 说明

- TODO module 下载的源代码不在 GOPATH 了，那么影响使用 vim-go 等插件阅读源码吗？即， gopls 这些源码分析软件是不是也要更新？
- 扩展阅读 [干货满满的 Go Modules 和 goproxy.cn](https://mp.weixin.qq.com/s/AsdCDodxZFxs2SkhSwOvpg)


[^Github]: [Github Incident](https://blog.github.com/2016-02-03-january-28th-incident-report/) 

[^Bitbucket]: [Bitbucket Post Mortem](https://blog.bitbucket.org/2012/09/19/post-mortem-on-our-availability-earlier-today/) 

[^Gitlab]: [Gitlab Post Mortem](https://about.gitlab.com/2017/02/10/postmortem-of-database-outage-of-january-31/)? 

[^NPM]: [npm Vanishes](https://www.theregister.co.uk/2016/03/23/npm_left_pad_chaos/)?

[^PlayingWithGoModuleProxies]: [Playing with Go module proxies](https://roberto.selbach.ca/go-proxies/)



<!--stackedit_data:
eyJoaXN0b3J5IjpbNzI4MjQzNDkzLC0xODg4MzY5MzE3XX0=
-->
