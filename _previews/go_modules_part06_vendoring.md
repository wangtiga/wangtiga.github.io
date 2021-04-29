
# Modules Part 06: Vendoring  [^GolangModulesPart06Vendoring]

William Kennedy

April 13, 2020

### Series Index

[Why and What](https://www.ardanlabs.com/blog/2019/10/modules-01-why-and-what.html)  
[Projects, Dependencies and Gopls](https://www.ardanlabs.com/blog/2019/12/modules-02-projects-dependencies-gopls.html)  
[Minimal Version Selection](https://www.ardanlabs.com/blog/2019/12/modules-03-minimal-version-selection.html)  
[Mirrors, Checksums and Athens](https://www.ardanlabs.com/blog/2020/02/modules-04-mirros-checksums-athens.html)  
[Gopls Improvements](https://www.ardanlabs.com/blog/2020/04/modules-05-gopls-improvements.html)  
[Vendoring](https://www.ardanlabs.com/blog/2020/04/modules-06-vendoring.html)


### Introduction

It’s no secret that I am a fan of vendoring when it’s reasonable and practical to use it for your application projects. I believe vendoring gives your application projects the most durability since the project owns every line of source code it needs to build the applications. If you want a reproducible build without needing to rely on external services (like module mirrors) and being connected to the network, vendoring is the solution.

These are other benefits of vendoring:

-   If dependencies are removed from the VCS or somehow proxy servers lose modules, you are covered.
-   Upgrading dependencies can be seen by running diffs and you maintain a history.
-   You will have the ability to trace and debug your dependencies and test changes if necessary.
    -   Once you run a  `go mod tidy`  and  `go mod vendor`  your changes will be replaced.

In this post, I will provide a history of Go’s support for vendoring and the changes in default behavior that have existed over time. I will also share how Go’s tooling is capable of maintaining backwards compatibility between versions. Finally, I will share how you may need to (over time) manually upgrade the version listed in the  `go.mod`  file to change the default behavior of future Go releases.

### Running Different Versions Of Go

To show you the differences in default behavior between Go 1.13 and Go 1.14, I need to be able to run both versions of the tooling on my machine at the same time. I’ve already installed Go 1.14.2 on my machine at the time I published this post and I access that version using the traditional  `go`  front end. However for this post, I also need to run a Go 1.13 environment. So how can I do that without disrupting my current development environment?

Luckily, the Go team publishes version downloads that give you a specific binary for any version of Go you want to use, including  [Go Tip](https://pkg.go.dev/golang.org/dl/gotip?tab=doc).

**Figure 1**  
![](https://www.ardanlabs.com/images/goinggo/117_figure1.png)

Figure 1 shows a screenshot of the Go 1.13.10  [page](https://pkg.go.dev/golang.org/dl/go1.13.10?tab=doc)  from the download server. It shows the instructions for building a binary that can be used to build and test your Go code using Go 1.13.10.

**Listing 1**

```
$ cd $HOME
$ go get golang.org/dl/go1.13.10

OUTPUT
go: downloading golang.org/dl v0.0.0-20200408221700-d6f4cf58dce2
go: found golang.org/dl/go1.13.10 in golang.org/dl v0.0.0-20200408221700-d6f4cf58dce2


$ go1.13.10 download

OUTPUT
Downloaded   0.0% (    14448 / 121613848 bytes) ...
Downloaded   9.5% ( 11499632 / 121613848 bytes) ...
Downloaded  30.8% ( 37436528 / 121613848 bytes) ...
Downloaded  49.2% ( 59849840 / 121613848 bytes) ...
Downloaded  69.3% ( 84262000 / 121613848 bytes) ...
Downloaded  90.3% (109804656 / 121613848 bytes) ...
Downloaded 100.0% (121613848 / 121613848 bytes)
Unpacking /Users/bill/sdk/go1.13.10/go1.13.10.darwin-amd64.tar.gz ...
Success. You may now run 'go1.13.10'


$ go1.13.10 version

OUTPUT
go version go1.13.10 darwin/amd64


$ go version

OUTPUT
go version go1.14.2 darwin/amd64

```

Listing 1 shows how after running the  `go get`  command for version 1.13.10 of Go and performing the  `download`  call, I can now use Go 1.13.10 on my machine without any disruption to my Go 1.14.2 installation.

If you want to remove any version of Go from your machine, you can find the specific binaries in your  `$GOPATH/bin`  folder and all the supporting files will be found in  `$HOME/sdk`.

**Listing 2**

```
$ cd $GOPATH/bin
$ l

OUTPUT
-rwxr-xr-x   1 bill  staff   7.0M Apr 11 10:51 go1.13.10
-rwxr-xr-x   1 bill  staff   2.3M Jan  6 11:02 gotip


$ cd $HOME
$ l sdk/

OUTPUT
drwxr-xr-x  22 bill  staff   704B Apr 11 10:52 go1.13.10
drwxr-xr-x  24 bill  staff   768B Feb 26 01:59 gotip

```

### Quick Vendoring Tutorial

The Go tooling did a great job minimizing the workflow impacts to manage and vendor an application project’s dependencies. It requires two commands:  `tidy`  and  `vendor`.

**Listing 3**

```
$ go mod tidy

```

Listing 3 shows the  `tidy`  command that helps to keep the dependencies listed in your module files accurate. Some editors (like VS Code and GoLand) provide support to update the module files during development but that doesn’t mean the module files will be clean and accurate once you have everything working. I recommend running the  `tidy`  command before you commit and push any code back to your VCS.

If you want to vendor those dependencies as well, then run the  `vendor`  command after  `tidy`.

**Listing 4**

```
$ go mod vendor

```

Listing 4 shows the  `vendor`  command. This command creates a vendor folder inside your project that contains the source code for all the dependencies (direct and indirect) that the project requires to build and test the code. This command should be run after running  `tidy`  to keep your vendor folder in sync with your module files. Make sure to commit and push the vendor folder to your VCS.

### GOPATH or Module Mode

In Go 1.11, a new mode was added to the Go tooling called “module mode”. When the Go tooling is operating in module mode, the module system is used to find and build code. When the Go tooling is operating in GOPATH mode, the traditional GOPATH system continues to be used to find and build code. One of the bigger struggles I have had with the Go tooling is knowing what mode will be used by default between the different versions. Then knowing what configuration changes and flags I need to keep my builds consistent.

To understand the history and the semantic changes that have occurred over the past 4 versions of Go, it’s good to have a refresher on these modes.

**Go 1.11**

A new environment variable was introduced called  `GO111MODULE`  whose default was  `auto`. This variable would determine if the Go tooling would use module mode or GOPATH mode depending on where the code was located (inside or outside of GOPATH). To force one mode or the other, you would set this variable to  `on`  or  `off`. When it came to vendor folders, module mode would ignore a vendor folder by default and build dependencies against the module cache.

**Go 1.12**

The default setting for  `GO111MODULE`  remains  `auto`  and the Go tooling continues to determine module mode or GOPATH mode depending on where the code is located (inside or outside of GOPATH). When it comes to vendor folders, module mode would still ignore a vendor folder by default and build dependencies against the module cache.

**Go 1.13**

The default setting for  `GO111MODULE`  remains  `auto`  but the Go tooling is no longer sensitive to whether the working directory is within the GOPATH. Module mode would still ignore a vendor folder by default and build dependencies against the module cache.

**Go 1.14**

The default setting for  `GO111MODULE`  remains  `auto`  and the Go tooling is still no longer sensitive to whether the working directory is within the GOPATH. However, if a vendor folder exists, it will be used by default to build dependencies instead of the module cache [1]. In addition, the  `go`  command verifies that the project’s  `vendor/modules.txt`  file is consistent with its  `go.mod`  file.

### Backwards Compatibility Between Versions

The change in Go 1.14 to use the vendor folder by default over the module cache is the behavior I wanted for my projects. Initially I thought I could just use Go 1.14 to build against my existing projects and it would be enough, but I was wrong. After my first build with Go 1.14 and not seeing the vendor folder being respected, I learned that the Go tooling reads the  `go.mod`  file for version information and maintains backwards compatibility with that version listed. I had no idea, but it is clearly expressed in the release notes for Go 1.14.

[https://golang.org/doc/go1.14#go-command](https://golang.org/doc/go1.14#go-command)

_When the main module contains a top-level vendor directory and its  `go.mod`  file specifies Go 1.14 or higher, the go command now defaults to  `-mod=vendor`  for operations that accept that flag._

In order to use the new default behavior for vendoring, I was going to need to upgrade the version information in the  `go.mod`  file from Go 1.13 to Go 1.14. This is something I quickly did.

### Small Demo

To show you the behavior of Go 1.13 and Go 1.14, and how the tooling maintains backwards compatibility, I am going to use the  [service](https://github.com/ardanlabs/service)  project. I will show you how changing the version listed in  `go.mod`  will change the default behavior of the Go tooling.

To start, I will clone the service project outside of my GOPATH.

**Listing 5**

```
$ cd $HOME/code
$ git clone https://github.com/ardanlabs/service
$ cd service
$ code .

```

Listing 5 shows the commands to clone the project and open the project in VS Code.

**Listing 6**

```
$ ls -l vendor/

OUTPUT
total 8
drwxr-xr-x   3 bill  staff    96 Mar 26 16:01 contrib.go.opencensus.io
drwxr-xr-x  14 bill  staff   448 Mar 26 16:01 github.com
drwxr-xr-x  20 bill  staff   640 Mar 26 16:01 go.opencensus.io
drwxr-xr-x   3 bill  staff    96 Mar 26 16:01 golang.org
drwxr-xr-x   3 bill  staff    96 Mar 26 16:01 gopkg.in
-rw-r--r--   1 bill  staff  2860 Mar 26 16:01 modules.txt

```

Listing 6 shows the listing of the vendor folder for the service project. You can see directories for some of the popular VCS sites that exist today as well as several vanity domains. All the code the project depends on to build and test are located inside the vendor folder.

Next, I will manually change the  `go.mod`  file back to version 1.13. This will allow me to show you the behavior I experienced when I used Go 1.14 for the first time against this project.

**Listing 7**

```
module github.com/ardanlabs/service

go 1.13   // I just changed this from go 1.14 to go 1.13

```

Listing 7 shows the change I am making to the  `go.mod`  file (switching out  `go 1.14`  for  `go 1.13`).

**_Note: There is a  `go mod`  command that can be used to change the version in the  `go.mod`  file:  `go mod edit -go=1.14`_**

#### Go 1.13

On this first build, I will use Go 1.13.10 to build the  `sales-api`  application. Remember, the  `go.mod`  file is listing Go 1.13 as the compatible version for this project.

**Listing 8**

```
$ cd service/cmd/sales-api
$ go1.13.10 clean -modcache
$ go1.13.10 build

OUTPUT
go: downloading contrib.go.opencensus.io/exporter/zipkin v0.1.1
. . .
go: finding github.com/leodido/go-urn v1.2.0

```

Listing 8 shows how I navigated to the application folder, cleaned out my local module cache and then performed a build using Go 1.13.10. Notice how the Go tooling downloaded all the dependencies back into my module cache in order to build the binary. The vendor folder was ignored.

To get Go 1.13 to respect the vendor folder, I need to use the  `-mod=vendor`  flag when building and testing.

**Listing 9**

```
$ go1.13.10 clean -modcache
$ go1.13.10 build -mod=vendor

OUTPUT

```

Listing 9 shows how I am now using the  `-mod=vendor`  flag on the build call. This time the module cache is not re-populated with the missing modules and the code in the vendor folder is respected.

#### Go 1.14

This time I will run the build command using Go 1.14.2 without the use of the  `-mod=vendor`  flag.

**Listing 10**

```
$ go clean -modcache
$ go build

OUTPUT
go: downloading github.com/openzipkin/zipkin-go v0.2.2
. . .
go: downloading github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e


```

Listing 10 shows what happens when I use Go 1.14 to build the project. The vendor folder is not being respected because the tooling is operating with Go 1.13 semantics. This is because the  `go.mod`  file is listing Go 1.13 as the compatible version for the project. When I saw this for the first time I was shocked. This is what started my investigation.

If I switch the  `go.mod`  file to version 1.14, the default mode of the Go 1.14 tooling will switch to respect the vendor folder by default.

**Listing 11**

```
module github.com/ardanlabs/service

go 1.14   // I just changed this from go 1.13 to go 1.14

```

Listing 11 shows the change to the  `go.mod`  file back to 1.14. I will clear the module cache again and run the build command again using Go 1.14.

**Listing 12**

```
$ go clean -modcache
$ go build

OUTPUT

```

Listing 12 shows that the module cache is not re-populated this time on the call to  `go build`  using Go 1.14. Which means the vendor folder is being respected, without the need of the  `-mod=vendor`  flag. The default behavior has changed because the module file is listing Go 1.14.

### Future Changes For Vendoring and Modules

Thanks to  [John Reese](https://twitter.com/johnpreese), here is a link to a discussion about the tooling maintaining backwards compatibility between different versions of Go based on what is listed in the  `go.mod`  file. John was instrumental in making sure the post was accurate and flowed correctly.

[https://github.com/golang/go/issues/30791](https://github.com/golang/go/issues/30791)

There is more support coming for vendoring that will follow in future releases. One such feature being discussed is about validating the code in the vendor folder to find situations where the code has been changed.

[https://github.com/golang/go/issues/27348](https://github.com/golang/go/issues/27348)

I have to thank  [Chris Hines](https://twitter.com/chris_csguy)  for reminding me about the default behaviors in the previous versions of Go and how that has been promoted with each new release. Chris also provided some interesting links that share some history and other cool things coming to the Go tooling for modules. Chris was instrumental in making sure the post was accurate and flowed correctly.

[https://github.com/golang/go/issues/33848](https://github.com/golang/go/issues/33848)  
[https://github.com/golang/go/issues/36460](https://github.com/golang/go/issues/36460)

### Conclusion

This post is a result of me being surprised that the version listed in  `go.mod`  was affecting the default behavior of the Go tooling. In order to gain access to the new default vendoring behavior in Go 1.14 that I wanted, I had to manually upgrade the version listed in  `go.mod`  from 1.13 to 1.14.

I haven’t formed any concrete opinions about the Go tooling using the version information in  `go.mod`  to maintain backwards compatibility between versions. The Go tooling has never been tied to the Go compatibility promise and so this was unexpected to me. Maybe this is the start of something great and moving forward the Go tooling can grow without the Go community worrying if their builds, tests and workflows will break when a new version of the Go tooling is released.

If you have any opinions, I’d love to hear them on Twitter.

### Footnotes

[1] Beginning with Go 1.11, the  `-mod=vendor`  flag caused the  `go`  command to load packages from the vendor directory, instead of modules from the module cache. (The vendor directory contains individual packages, not complete modules.) In Go 1.14, the default value of the  `-mod`  flag changes depending on the contents of the main module: if there is a vendor directory and the  `go.mod`  file specifies  `go 1.14`  or higher,  `-mod`  defaults to  `-mod=vendor`. If the  `go.mod`  file is read-only,  `-mod`  defaults to  `-mod=readonly`. We also added a new value,  `-mod=mod`, meaning “load modules from the module cache” (that is, the same behavior that you get by default if none of the other conditions holds). Even if you are working in a main module for which the default behavior is  `-mod=vendor`, you can explicitly go back to the module cache using the  `-mod=mod`  flag. - Bryan Mills


[^GolangModulesPart06Vendoring]: [Vendoring](https://www.ardanlabs.com/blog/2020/04/modules-06-vendoring.html)

