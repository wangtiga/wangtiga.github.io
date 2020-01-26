

# Guide to pkg-config [^GuidToPkgConfig]

### Dan Nicholson

-   [Overview](https://people.freedesktop.org/~dbn/pkg-config-guide.html#overview)
-   [Why?](https://people.freedesktop.org/~dbn/pkg-config-guide.html#why)
-   [Concepts](https://people.freedesktop.org/~dbn/pkg-config-guide.html#concepts)
-   [Writing pkg-config files](https://people.freedesktop.org/~dbn/pkg-config-guide.html#writing)
-   [Using pkg-config files](https://people.freedesktop.org/~dbn/pkg-config-guide.html#using)
-   [Frequently asked questions](https://people.freedesktop.org/~dbn/pkg-config-guide.html#faq)

## Overview

This document aims to give an overview to using the  pkg-config  tool from the perspective of both a user and a developer. It reviews the concepts behind  pkg-config, how to write  pkg-config  files to support your project, and how to use  pkg-config  to integrate with 3rd party projects.

More information on  pkg-config  can be found at the  [website](http://pkg-config.freedesktop.org/)  and in the  pkg-config(1)  manual page.

This document assumes usage of  pkg-config  on a Unix-like operating system such as Linux. Some of the details may be different on other platforms.

## Why?

Modern computer systems use many layered components to provide applications to the user. One of the difficulties in assembling these parts is properly integrating them.  pkg-config  collects metadata about the installed libraries on the system and easily provides it to the user.

Without a metadata system such as  pkg-config, it can be very difficult to locate and obtain details about the services provided on a given computer. For a developer, installing  pkg-config  files with your package greatly eases adoption of your API.

没有 pkg-config 这样的元数据工具，想查找一台计算机中安装了哪些依赖库会很困难。
对 lib 的开发者来说，提供了 pkg-config 文件，能让别人使用你的 API 时更方便。



## Concepts 概念

The primary use of  pkg-config  is to provide the necessary details for compiling and linking a program to a library. This metadata is stored in  pkg-config  files. These files have the suffix  .pc  and reside in specific locations known to the  pkg-config  tool. This will be described in more detail later.

pkg-config 的主要用途是提供编译和链接 library 时所需的相关信息。
这些元信息保存在 pkg-config 文件中。
这些文件都以 `.pc` 为后缀。
保存在 pkg-config 工具指定的路径中。
稍后会详细描述。


The file format contains predefined metadata keywords and freeform variables. An example may be illustrative:

文件中可定义 metadata keyword (元数据关键字) 和 变量。
示例如下：

```txt
prefix=/usr/local
exec_prefix=${prefix}
includedir=${prefix}/include
libdir=${exec_prefix}/lib

Name: foo
Description: The foo library
Version: 1.0.0
Cflags: -I${includedir}/foo
Libs: -L${libdir} -lfoo
```

The keyword definitions such as  `Name:`  begin with a keyword followed by a colon and the value.
The variables such as  `prefix=`  are a string and value separated by an equals sign.
The keywords are defined and exported by  pkg-config .
The variables are not necessary, but can be used by the keyword definitions for flexibility or to store data not covered by  pkg-config.

`Name: foo` 定义一个 keyword ，`:` （冒号）前的 `Name` 是 keyword （名称），后面是关键字的 value （值）。

`prefix=` 定义一个 keyword ，`=` （等号）前的 `prefix` 是变量名，后面是变量的 value （值）。

变量可以在 keyword 中使用，增加灵活性。


Here is a short description of the keyword fields. A more in depth description of these fields and how to use them effectively will be given in the  [Writing pkg-config files](https://people.freedesktop.org/~dbn/pkg-config-guide.html#writing)  section.

下面是有关 keyword 的简介， Writing pkg-config files 一节有详细描述。

-   **Name**: A human-readable name for the library or package. This does not affect usage of the  pkg-config  tool, which uses the name of the  .pc  file.
-   **Name**:  用于人类阅读的 library 或 package 名称。不影响 pkg-config 工具。

-   **Description**: A brief description of the package.
-   **Description**: 有关 package 的简要描述。

-   **URL**: An URL where people can get more information about and download the package.
-   **URL**: 能下载并查阅有关 package 信息的 URL 。

-   **Version**: A string specifically defining the version of the package.
-   **Version**: 标识 package 版本的字符串。

-   **Requires**: A list of packages required by this package. The versions of these packages may be specified using the comparison operators =, <, >, <= or >=.
-   **Requires**: 当前 package 所依赖的 package 。 可以用 =, <, >, <= or >= 符号指定版本号。

-   **Requires.private**: A list of private packages required by this package but not exposed to applications. The version specific rules from the  Requires  field also apply here.
-   **Requires.private**: 当前 package 所依赖的 package 。这些依赖不会暴露给外部。配置规则与 Requires 一样。（TODO 仅仅是查询时看不到？还是它依赖的动态库，会存放在特殊目录，不会被多个相同 package 不同版本的库所覆盖？）

-   **Conflicts**: An optional field describing packages that this one conflicts with. The version specific rules from the  Requires  field also apply here. This field also takes multiple instances of the same package. E.g.,  Conflicts: bar < 1.2.3, bar >= 1.3.0.
-   **Conflicts**: 与当前库有冲突的 package 。配置规则与 Requires 一样，而且可以配置相同 package 的多个版本。比如，`Conflicts: bar < 1.2.3, bar >= 1.3.0` 。

-   **Cflags**: The compiler flags specific to this package and any required libraries that don't support  pkg-config. If the required libraries support  pkg-config, they should be added to  Requires  or  Requires.private.
-   **Cflags**: 用于设置当前 package 的编译选项(compiler flags)，或者当前 package 依赖的其他不支持 pkg-config 的 library 。如果依赖的 library 支持 pkg-config ，应该使用 `Requires`  或 `Requires.private` 配置。

-   **Libs**: The link flags specific to this package and any required libraries that don't support  pkg-config. The same rule as  Cflags  applies here.
-   **Libs**: 用于设置当前 package 的链接选项(link flags)，或者当前 package 依赖的其他不支持 pkg-config 的 library 。配置规则与 `Cflags` 一样。


-   **Libs.private**: The link flags for private libraries required by this package but not exposed to applications. The same rule as  Cflags  applies here.
-   **Libs.private**: 当前 package 所依赖的 package 所使用的 链接选项(link flags)。这些依赖不会暴露给外部。配置规则与 `Cflags` 一样。



## Writing pkg-config files

When creating  pkg-config  files for a package, it is first necessary to decide how they will be distributed. Each file is best used to describe a single library, so each package should have at least as many  pkg-config  files as they do installed libraries.

The package name is determined through the filename of the  pkg-config  metadata file. This is the portion of the filename prior to the  .pc  suffix. A common choice is to match the library name to the  .pc  name. For instance, a package installing  libfoo.so  would have a corresponding  libfoo.pc  file containing the  pkg-config  metadata. This choice is not necessary; the  .pc  file should simply be a unique identifier for your library. Following the above example,  foo.pc  or  foolib.pc  would probably work just as well.

The  Name,  Description  and  URL  fields are purely informational and should be easy to fill in. The  Version  field is a bit trickier to ensure that it is usable by consumers of the data.  pkg-config  uses the algorithm from  [RPM](http://rpm.org/)  for version comparisons. This works best with a dotted decimal number such as  1.2.3  since letters can cause unexpected results. The number should be monotonically increasing and be as specific as possible in describing the library. Usually it's sufficient to use the package's version number here since it's easy for consumers to track.

Before describing the more useful fields, it will be helpful to demonstrate variable definitions. The most common usage is to define the installation paths so that they don't clutter the metadata fields. Since the variables are expanded recursively, this is very helpful when used in conjunction with autoconf derived paths.

```txt
prefix=/usr/local
includedir=${prefix}/include

Cflags: -I${includedir}/foo
```

The most important  pkg-config  metadata fields are  Requires,  Requires.private,  Cflags,  Libs  and  Libs.private. They will define the metadata used by external projects to compile and link with the library.

Requires  and  Requires.private  define other modules needed by the library. It is usually preferred to use the private variant of  Requires  to avoid exposing unnecessary libraries to the program that is linking with your library. If the program will not be using the symbols of the required library, it should not be linking directly to that library. See the discussion of  [overlinking](https://wiki.openmandriva.org/en/Overlinking_issues_in_packaging)  for a more thorough explanation.

Since  pkg-config  always exposes the link flags of the  Requires  libraries, these modules will become direct dependencies of the program. On the other hand, libraries from  Requires.private  will only be included when static linking. For this reason, it is usually only appropriate to add modules from the same package in  Requires.

The  Libs  field contains the link flags necessary to use that library. In addition,  Libs  and  Libs.private  contain link flags for other libraries not supported by  pkg-config. Similar to the  Requires  field, it is preferred to add link flags for external libraries to the  Libs.private  field so programs do not acquire an additional direct dependency.

Finally, the  Cflags  contains the compiler flags for using the library. Unlike the  Libs  field, there is not a private variant of  Cflags. This is because the data types and macro definitions are needed regardless of the linking scenario.

## Using pkg-config files

Assuming that there are  .pc  files installed on the system, the  pkg-config  tool is used to extract the metadata for usage. A short description of the options can be seen by executing  pkg-config --help. A more in depth discussion can be found in the  pkg-config(1)  manual page. This section will provide a brief explanation of common usages.

`.pc` 文件安装到系统后，可使用 pkg-config 工具获取和系统中依赖库的 metadata 来使用。
`pkg-config --help` 可以看到一些简单的使用说明。
`pkg-config(1) manual page` 手册中能看到更详细的讨论。
这里也会讲解一些常见的用法。

Consider a system with two modules,  foo  and  bar. Their  .pc  files might look like this:

假设系统中有两个 modules 分别是 foo 和 bar 。
它们的 .pc 文件如下所示：

foo.pc:
```txt
prefix=/usr
exec_prefix=${prefix}
includedir=${prefix}/include
libdir=${exec_prefix}/lib

Name: foo
Description: The foo library
Version: 1.0.0
Cflags: -I${includedir}/foo
Libs: -L${libdir} -lfoo
```

bar.pc:
```
prefix=/usr
exec_prefix=${prefix}
includedir=${prefix}/include
libdir=${exec_prefix}/lib

Name: bar
Description: The bar library
Version: 2.1.2
Requires.private: foo >= 0.7
Cflags: -I${includedir}
Libs: -L${libdir} -lbar
```

The version of the modules can be obtained with the  --modversion  option.

使用 `--modversion` 选项获取 module 的版本信息。

```shell
$ pkg-config --modversion foo
1.0.0
$ pkg-config --modversion bar
2.1.2
```

To print the link flags needed for each module, use the  --libs  option.

使用 `--libs` 选项获取 module 的 link flags 。

```shell
$ pkg-config --libs foo
-lfoo
$ pkg-config --libs bar
-lbar
```

Notice that  pkg-config  has suppressed part of the  Libs  field for both modules. This is because it treats the  -L  flag specially and knows that the  ${libdir}  directory  /usr/lib  is part of the system linker search path. This keeps  pkg-config  from interfering with the linker operation.

Also, although  foo  is required by  bar, the link flags for  foo  are not output. This is because  foo  is not directly needed by an application that only wants to use the  bar  library. For statically linking a  bar  application, we need both sets of linker flags:

bar 依赖 foo ，但 bar 的 link flags 中并没有 foo 。
这是因为使用 bar 的 library 时，并不直接使用 foo 。
如果程序需要静态链接 bar 时，就要设置两个 flags ：

```shell
$ pkg-config --libs --static bar
-lbar -lfoo
```

pkg-config  needs to output both sets of link flags in this case to ensure that the statically linked application will find all the necessary symbols. On the other hand, it will always output all the  Cflags.

```shell
$ pkg-config --cflags bar
-I/usr/include/foo
$ pkg-config --cflags --static bar
-I/usr/include/foo
```

Another useful option,  --exists, can be used to test for a module's availability.

```shell
$ pkg-config --exists foo
$ echo $?
0
```

One of the nicest features of  pkg-config  is providing version checking. It can be used to determine if a sufficient version is available.

```shell
$ pkg-config --libs "bar >= 2.7"
Requested 'bar >= 2.7' but version of bar is 2.1.2
```

Some commands will provide more verbose output when combined with the  --print-errors  option.

```shell
$ pkg-config --exists --print-errors xoxo
Package xoxo was not found in the pkg-config search path.
Perhaps you should add the directory containing `xoxo.pc'
to the PKG_CONFIG_PATH environment variable
No package 'xoxo' found
```

The message above references the  PKG_CONFIG_PATH  environment variable. This variable is used to augment  pkg-config's search path. On a typical Unix system, it will search in the directories  /usr/lib/pkgconfig  and  /usr/share/pkgconfig. This will usually cover system installed modules. However, some local modules may be installed in a different prefix such as  /usr/local. In that case, it's necessary to prepend the search path so that  pkg-config  can locate the  .pc  files.

```shell
$ pkg-config --modversion hello
Package hello was not found in the pkg-config search path.
Perhaps you should add the directory containing `hello.pc'
to the PKG_CONFIG_PATH environment variable
No package 'hello' found
$ export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
$ pkg-config --modversion hello
1.0.0
```

A few  [autoconf](http://www.gnu.org/software/autoconf/)  macros are also provided to ease integration of  pkg-config  modules into projects.

-   **PKG_PROG_PKG_CONFIG([MIN-VERSION])**: Locates the  pkg-config  tool on the system and checks the version for compatibility.
-   **PKG_CHECK_EXISTS(MODULES, [ACTION-IF-FOUND], [ACTION-IF-NOT-FOUND])**: Checks to see whether a particular set of modules exists.
-   **PKG_CHECK_MODULES(VARIABLE-PREFIX, MODULES, [ACTION-IF-FOUND], [ACTION-IF-NOT-FOUND])**: Checks to see whether a particular set of modules exists. If so, it sets  <VARIABLE-PREFIX>_CFLAGS  and  <VARIABLE-PREFIX>_LIBS  according to the output from  pkg-config --cflags  and  pkg-config --libs.

## Frequently asked questions

1.  My program uses library  x. What do I do?

The  pkg-config  output can easily be used on the compiler command line. Assuming the  x  library has a  x.pc  pkg-config  file:

cc `pkg-config --cflags --libs x` -o myapp myapp.c

The integration can be more robust when used with  [autoconf](http://www.gnu.org/software/autoconf/)  and  [automake](http://www.gnu.org/software/automake/). By using the supplied  PKG_CHECK_MODULES  macro, the metadata is easily accessed in the build process.

configure.ac:
```shell
PKG_CHECK_MODULES([X], [x])
```

Makefile.am:
```shell
myapp_CFLAGS = $(X_CFLAGS)
myapp_LDADD = $(X_LIBS)
```

If the  x  module is found, the macro will fill and substitute the  `X_CFLAGS`  and  `X_LIBS`  variables. If the module is not found, an error will be produced. Optional 3rd and 4th arguments can be supplied to  `PKG_CHECK_MODULES`  to control actions when the module is found or not.

7.  My library  z  installs header files which include  libx  headers. What do I put in my  z.pc  file?

If the  x  library has  pkg-config  support, add it to the  Requires.private  field. If it does not, augment the  Cflags  field with the necessary compiler flags for using the  libx  headers. In either case,  pkg-config  will output the compiler flags when  --static  is used or not.

9.  My library  z  uses  libx  internally, but does not expose  libx  data types in its public API. What do I put in my  z.pc  file?

Again, add the module to  Requires.private  if it supports  pkg-config. In this case, the compiler flags will be emitted unnecessarily, but it ensures that the linker flags will be present when linking statically. If  libx  does not support  pkg-config, add the necessary linker flags to  Libs.private.

----------

Dan Nicholson <dbn.lists (at) gmail (dot) com>

Copyright (C) 2010 Dan Nicholson.  
This document is licensed under the  [GNU General Public License, Version 2](http://www.gnu.org/licenses/old-licenses/gpl-2.0.txt)  or any later version.



[^GuidToPkgConfig]: [Guide to pkg-config](https://people.freedesktop.org/~dbn/pkg-config-guide.html)

