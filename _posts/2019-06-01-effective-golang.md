---
layout: post
title:  "[译] Effective Go"
date:   2019-06-01 09:22:00 +0800
tags: golang
---

* category
{:toc}



# Effective Go 高效地使用 Go 语言 [^EffectiveGoEn] [^EffectiveGoCn1] [^EffectiveGoCn2]



## Introduction 简介

Go is a new language. 
Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. 
A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. 
On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program.
In other words, to write Go well, it's important to understand its properties and idioms.
It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.

Go 是一门新语言。尽管它也借鉴了现存语言的一些思想，使用 Go 完成一个高效（优秀）的程序，相比使用其他类似语言差异巨大。
直接将 C++ 或 Java 翻译成 Go 是写不出好程序的，Java 程序就要用 Java 写，不能用 Go 。
以 Go 的视角思考程序，会产生与众不同的优秀设计。
换种说法，要想写好 Go 程序，理解 Go 的特性和习惯十分重要。
另外，了解 Go 程序的构建习惯也很重要，像命名，格式化，代码组织结构等等。按这些标准规范编写的程序才更容易让其他 Go 开发者理解。
 

This document gives tips for writing clear, idiomatic Go code. It augments the  [language specification](https://golang.org/ref/spec), the  [Tour of Go](http://tour.golang.org/), and  [How to Write Go Code](https://golang.org/doc/code.html), all of which you should read first.

本文档描述如何编写清晰，符合惯例的 Go 代码技巧。
此读此文档之前，最好先看看 语言规划，Go 手册，如何编写 Go 代码几个文档。
因为本文档是这些内容的延伸阅读。



### Examples 示例

The  [Go package sources](https://golang.org/src/)  are intended to serve not only as the core library but also as examples of how to use the language. Moreover, many of the packages contain working, self-contained executable examples you can run directly from the[golang.org](http://golang.org/)  web site, such as  [this one](http://golang.org/pkg/strings/#example_Map)  (if necessary, click on the word "Example" to open it up). If you have a question about how to approach a problem or how something might be implemented, the documentation, code and examples in the library can provide answers, ideas and background.
 
Go package sources 不仅是核心库代码，也是演示如何编写Go语言的代码样板。
此外，其中许多 package 包含可以直接运行的示例，
你能直接在 golang.org 网站中运行。
如果你有类似 “如何解决某问题” 或者 “如何实现某功能” 的疑问，也许能在这此文档、代码和示例找到答案或一丝线索。



## Formatting 编码格式

Formatting issues are the most contentious but the least consequential. People can adapt to different formatting styles but it's better if they don't have to, and less time is devoted to the topic if everyone adheres to the same style. The problem is how to approach this Utopia without a long prescriptive style guide.

代码格式化（排版）也许是争议最多，但最不重要的问题了。
人们能适应不同的格式化风格，但如果所有人都坚持一种风格，就能减少大量有关这种问题的争论。
问题是，如何能在脱离冗长的代码风格指南的情况下，达到这种理想乌托邦呢。
 
With Go we take an unusual approach and let the machine take care of most formatting issues. The  `gofmt`  program (also available as  `go fmt`, which operates at the package level rather than source file level) reads a Go program and emits the source in a standard style of indentation and vertical alignment, retaining and if necessary reformatting comments. If you want to know how to handle some new layout situation, run  `gofmt`; if the answer doesn't seem right, rearrange your program (or file a bug about  `gofmt`), don't work around it.

在 Go 中，我们用了一种不同寻常的办法，那就是让机器解决大部分格式化问题。 
gofmt 程序（`go fmt` 命令多用于格式化 package ，`gofmt` 多用于格式化 source file）用于读取 Go 代码，并将源码缩进、对齐、注释等规范成标准风格。
如果不清楚如何处理某种代码格式，那就运行`gofmt;`，如果结果不太对，重新整理代码后重试一下（或者给 `gofmt`提一个bug）。
 
As an example, there's no need to spend time lining up the comments on the fields of a structure.  `Gofmt`  will do that for you. Given the declaration

看看下面的示例，我们不用浪费时间手动将结构体中的字段名对齐了，直接用 `gofmt`就能解决。
看看下面的声名：

```go
type T struct {
    name string // name of the object
    value int // its value
}
```

`gofmt`  will line up the columns:

`gofmt` 会直接按列对齐每行的字段类型:

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

All Go code in the standard packages has been formatted with  `gofmt`.

所有标准库中的Go代码都经过`gofmt`格式化过。

Some formatting details remain. Very briefly:

还有一些格式化要求，简介如下：


- Indentation 缩进

We use tabs for indentation and  `gofmt`emits them by default. Use spaces only if you must.

我们使用 tabs 缩进，`gofmt`默认也这样。
非特殊情况，不要使用空格缩进。


- Line length 每行代码宽度

Go has no line length limit. Don't worry about overflowing a punched card. If a line feels too long, wrap it and indent with an extra tab.

Go 不限制每行代码的长度。
不要担心穿孔卡片宽度不够（最早编程用穿孔卡片）。
如果觉得一行太长了，就换行，然后用几个 tabs 缩进一下就行。


- Parentheses 圆括号 

Go needs fewer parentheses than C and Java: control structures (`if`,  `for`,  `switch`) do not have parentheses in their syntax. Also, the operator precedence hierarchy is shorter and clearer, so

Go 相比 C 和 Java 很少使用圆括号，控制结构（如 if, for, switch）的语法中都不要求圆括号。
运算符的优先级别也更简洁清晰，比如：

```go
x<<8 + y<<16
```

means what the spacing implies, unlike in the other languages.

表示含义同空格分隔的一样，不像其他语言那么麻烦。

TODO 验证其他语言运算符号优先级有什么问题？



## Commentary 注释

Go provides C-style  `/* */`  block comments and C++-style  `//`  line comments.
Line comments are the norm;
block comments appear mostly as package comments, but are useful within an expression or to disable large swaths of code.

Go 提供 C 风格的块注释 `/*  */` ，还有C++风格的行注释 `//`。
行注释使用更普遍一些;
块注释较多用于 package 注释中，另外也用于行内注释，或者注释某一大段代码块。

> 译：行内注释即： `if a>b /*&& a>c*/` ，其中 `a>c` 的条件就失效了 


The program—and web server—`godoc`processes Go source files to extract documentation about the contents of the package.
Comments that appear before top-level declarations, with no intervening newlines, are extracted along with the declaration to serve as explanatory text for the item.
The nature and style of these comments determines the quality of the documentation  `godoc`  produces.

`godoc`即是一个程序也是一个 web 服务器，它能从Go源代码中提取有关 package 的文档。
所有出现在声名上方的注释，只有没有空行间隔，都会与该声名一起提取出来，作为说明文档。
注释的风格直接影响 `godoc` 生成的文档质量。


Every package should have a  _package comment_, a block comment preceding the package clause. For multi-file packages, the package comment only needs to be present in one file, and any one will do. The package comment should introduce the package and provide information relevant to the package as a whole. It will appear first on the  `godoc`  page and should set up the detailed documentation that follows.

每个 package 都应该有 _package 注释_ ，即要在 package 声名之前写一些注释。
对于含有多个文件的 package ，说明应该集中在一个文件中，哪个文件都行。
注释中应该介绍 package 的功能，并提供所有相关信息。
这些内容会出现在 `godoc` 生成的页面中，所以应该像现在的示例一样编写文档。

```go
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

If the package is simple, the package comment can be brief.

如果 package 很简单，package 注释也可以简略一点。

```go
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

Comments do not need extra formatting such as banners of stars. The generated output may not even be presented in a fixed-width font, so don't depend on spacing for alignment—`godoc`, like  `gofmt`, takes care of that. The comments are uninterpreted plain text, so HTML and other annotations such as  `_this_`  will reproduce  _verbatim_  and should not be used. One adjustment  `godoc`  does do is to display indented text in a fixed-width font, suitable for program snippets. 
The package comment for the  [`fmt`  package](https://golang.org/pkg/fmt/)  uses this to good effect.

不要对文字进行额外格式化（比如，不要用 Markdown 语法中那种用 星号 ﹡ 进行加粗显示）。
`godoc` 输出的文档可能不是等宽字体，也不要依赖空格进行对齐。
`godoc` 与 `gofmt` 处理缩进的方法一样（译：建议用 tab 制表符号进行缩进，官方代码 fmt 就是这样处理的）。
注释是不经任何处理的文本，因此类似 HTML 或 像 `_this_` 这种 Markdown 语法都会都会 _原样_ 输出，所以不要使用。
`godoc` 唯一进行特殊处理的地方就是，把缩进过的文本用等宽字体显示，这是为了方便显示一些程序代码片断。
[`fmt`  package](https://golang.org/pkg/fmt/)的注释就用了这种效果。.

Depending on the context,  `godoc`  might not even reformat comments, so make sure they look good straight up: use correct spelling, punctuation, and sentence structure, fold long lines, and so on.  

`godoc` 根据上下文房室是否对注释进行格式化，一定要让他们看起来整洁一些，确保拼写、标点、句子结构以及缩进都没有问题。

Inside a package, any comment immediately preceding a top-level declaration serves as a  _doc comment_  for that declaration. Every exported (capitalized) name in a program should have a doc comment.

package内部，所有紧挨声明之上的注释文字，都被当做该声名的 _文档注释_ 。所有导出变量（即大写字母开头的变量）都应该有文档注释。

Doc comments work best as complete sentences, which allow a wide variety of automated presentations. The first sentence should be a one-sentence summary that starts with the name being declared.

文档注释最好是一个完整句子，这样方便任意显示格式。注释的第一句话，应该以所声名的变量名称开头，做简要介绍。

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

If every doc comment begins with the name of the item it describes, you can use the  [doc](https://golang.org/cmd/go/#hdr-Show_documentation_for_package_or_symbol)  subcommand of the  [go](https://golang.org/cmd/go/)  tool and run the output through  `grep`. Imagine you couldn't remember the name "Compile" but were looking for the parsing function for regular expressions, so you ran the command,

如果每个文档注释都以它描述的变量名开头，就能方便的使用 `grep` 过滤 [`go`](https://golang.org/cmd/go/)  [`doc`](https://golang.org/cmd/go/#hdr-Show_documentation_for_package_or_symbol) 的输出内容。 
假设你想寻找正则表达式函数，但不记得函数名是"Compile"了，你可以使用下面的命令搜索文档。

```shell
$ go doc -all regexp | grep -i parse
```

If all the doc comments in the package began, "This function...",  `grep`  wouldn't help you remember the name. But because the package starts each doc comment with the name, you'd see something like this, which recalls the word you're looking for.

如果文档说明没有以它描述的函数名开头（即"Compile"），`grep` 就没法显示出准确的函数名。
但我们如果每个package中的文档注释都以它描述的变量名开头，你就能看到类似下面的输出结果：

```shell
$ go doc -all regexp | grep -i parse
    Compile parses a regular expression and returns, if successful, a Regexp
    MustCompile is like Compile but panics if the expression cannot be parsed.
    parsed. It simplifies safe initialization of global variables holding
$
// TODO 在 windows 7 go 1.9.1 中测试，godoc 输出的函数文档虽然逻辑上是一句话
// 但实际输出仍然是多行的，所以 grep 过滤时，不会显示 Compile 这行字符
// 这也就达不到上文说的目的了，不知道是不是我测试环境有问题？
```

Go's declaration syntax allows grouping of declarations. A single doc comment can introduce a group of related constants or variables. Since the whole declaration is presented, such a comment can often be perfunctory.

Go支持批量声名。
此时这组变量也共用一个文档说明。
因为所有声名代码也会显示出来，所以这种文档注释通过比较笼统。

```go
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

Grouping can also indicate relationships between items, such as the fact that a set of variables is protected by a mutex.

批量声名通常用于声名几个有关联性的数据项，比如下面这种，多个变量同时由一个mutex保护。

```go
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```

## Names 命名

Names are as important in Go as in any other language. They even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case. It's therefore worth spending a little time talking about naming conventions in Go programs.

Go语言中命名的重要性同其他语言一样。
命名甚至能影响语法：package中变量名称首字母大小写决定其是否对外部可见。
因此值我们花点时间了解有关Go的命名习惯。


### Package names 包名

When a package is imported, the package name becomes an accessor for the contents. After

当 package 被导入(import)后，其 package name 就是一个访生问器。出现下面代码后，

```go
import "bytes"
```

the importing package can talk about  `bytes.Buffer`. 
It's helpful if everyone using the package can use the same name to refer to its contents, which implies that the package name should be good: short, concise, evocative. 
By convention, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps. 
Err on the side of brevity, since everyone using your package will be typing that name. 
And don't worry about collisions  _a priori_. 
The package name is only the default name for imports; it need not be unique across all source code, and in the rare case of a collision the importing package can choose a different name to use locally. 
In any case, confusion is rare because the file name in the import determines just which package is being used.

我们就能使用 `bytes.Buffer` 这样的类型了。
一个好的 package 名称应该具备以下特点：简单又好记。
这样的名称有助于大家都用相同的 name 引用 package ，
packages 一般使用小字字母的单个单词命名，不使用下划线，不使用大小写字母（驼峰命名法）。
Err on the side of brevity, since everyone using your package will be typing that name.
因为每个使用 package 的人都会键入包名，所以像 `err` 这种简洁的名称就很方便。
不用担心因为优先级别造成命名冲突。
package name 只是 import 时的默认名称；没必要在所有源代码中都是唯一的。
偶尔遇到冲突时，起个别名就能解决问题（别名不是全局生效的）。
而且只有在使用 package  时，才会通过 import 的名称识别 package ，所以发生冲突的情况其实很少。
 

Another convention is that the package name is the base name of its source directory; the package in  `src/encoding/base64`  is imported as  `"encoding/base64"`  but has name  `base64`, not  `encoding_base64`  and not  `encodingBase64`.

另外一个惯例是，package 名称是源代码所有目录的名称；
比如 `src/encoding/base64` 导入时使用 `import "encoding/base64"` ，但真正调用时，使用"base64"作为名称。
既不是`encoding_base64`，也不是`encodingBase64`。

The importer of a package will use the name to refer to its contents, so exported names in the package can use that fact to avoid stutter. 
(Don't use the  `import .`notation, which can simplify tests that must run outside the package they are testing, but should otherwise be avoided.) 
For instance, the buffered reader type in the  `bufio`  package is called  `Reader`, not  `BufReader`, because users see it as  `bufio.Reader`, which is a clear, concise name. 
Moreover, because imported entities are always addressed with their package name,  `bufio.Reader`  does not conflict with  `io.Reader`. 
Similarly, the function to make new instances of  `ring.Ring`—which is the definition of a  _constructor_  in Go—would normally be called  `NewRing`, but since  `Ring`  is the only type exported by the package, and since the package is called  `ring`, it's called just  `New`, which clients of the package see as  `ring.New`. 
Use the package structure to help you choose good names.

> 译： exported name 直译为可导出名称，可理解为 C+＋ 或 Java 语言中的 public 方法或成员变量

使用者通过 package name 引用 package 中的内容，所以不同中 package 内的 exported name 也因此避免冲突。
(尽量不要使用 `import . xxx/xxx` 语法，这只是为了简化单元测试代码中，必须在 package 外调用 package 内部函数进行测试的一种方法，除此以外，应该尽量避免使用）
 比如在bufio中的 buffered reader 的 package name 是`Reader`，而不是`BufReader`，因为使用者通过 `bufio.Reader` 调用。
 因为调用者总会加上 package name 为前缀使用，所以 `bufio.Reader` 永远不会和 `io.Reader` 冲突。
 同样，一般用于创建一个`ring.Ring`的新实例的函数，我们起名为`NewRing`,但因为`Ring`中 package `ring` 中的导出类型，
 所以我们将函数命名为`New`就可以了。这样用户就能使用`ring.New`这种简洁的名称。
 利用 package 的目录结构帮你起个好名字。


Another short example is  `once.Do`;`once.Do(setup)`  reads well and would not be improved by writing  `once.DoOrWaitUntilDone(setup)`. Long names don't automatically make things more readable. A helpful doc comment can often be more valuable than an extra long name.
 
 还有个例子，`once.Do; once.Do(setup)`明显就比`once.DoOrWaitUntilDone(setup)`好多了。
 过长的名字反而可能影响可读性。好的 doc comment 可能比冗长的名称要有用得多。

> 译：结论我同意，但这个例子中，我觉得 DoOrWaitUntilDone() 更好,还不到20个字符的名字，不能算长 :) 



### Getters 访问器

Go doesn't provide automatic support for getters and setters. There's nothing wrong with providing getters and setters yourself, and it's often appropriate to do so, but it's neither idiomatic nor necessary to put  `Get`into the getter's name. If you have a field called  `owner`  (lower case, unexported), the getter method should be called  `Owner`(upper case, exported), not  `GetOwner`. The use of upper-case names for export provides the hook to discriminate the field from the method. A setter function, if needed, will likely be called  `SetOwner`. Both names read well in practice:

Go不提供默认的 Getter 和 Setter 。
这种东西由程序员自己实现就行。
但没必要在 Getter 函数名前加 Get 前缀。
如果你有一个名为 `owner` （小写，表示私有变量）的字段，那么其 Getter 函数名可起为 `Owner` （大小，表示公有函数），
没必要起这 `GetOwner` 这样的名称。因为我们仅凭大小写就能区分出字段和函数。
如果需要 Setter 函数，可以起名为 `SetOwner` 。
示例如下：

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### Interface names 接口名称

By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun:  `Reader`,  `Writer`,  `Formatter`,`CloseNotifier`  etc.

通常，仅有一个函数的 interface ，一般用它的函数名加 `-er` 后缀修饰成名词，比如：`Reader`,  `Writer`,  `Formatter`,`CloseNotifier` 等。


There are a number of such names and it's productive to honor them and the function names they capture.  `Read`,  `Write`,  `Close`,  `Flush`,  `String`  and so on have canonical signatures and meanings. To avoid confusion, don't give your method one of those names unless it has the same signature and meaning. Conversely, if your type implements a method with the same meaning as a method on a well-known type, give it the same name and signature; call your string-converter method  `String`not  `ToString`.

类似命名还有很多，遵循他们的字面意思，会更有效率一些。
`Read`,  `Write`,  `Close`,  `Flush`,  `String` 这些都是具有典型含义的名称。
为避免混淆，不要使用这些名称给你的方法命名，除非你的用途与它们含义完全一样。
	另外，如果你实现的一个方法与众所周知的方法有相同的含义，那就要使用相同的名称；比如，执行字符串转换的方法，就应该起名为 `String` ，而不是 `ToString` 。


### MixedCaps 驼峰命名法

Finally, the convention in Go is to use  `MixedCaps`  or  `mixedCaps`  rather than underscores to write multiword names.

通常，Go中倾向使用`MixedCaps`或`mixedCaps`这中驼峰命名法，很少使用下划线`_`分隔多个单词。

> 译：变量 函数等使用驼峰命名法，但 package 名称只用小写字母


## Semicolons 分号

Like C, Go is formal grammar uses semicolons to terminate statements, but unlike in C, those semicolons do not appear in the source. Instead the lexer uses a simple rule to insert semicolons automatically as it scans, so the input text is mostly free of them.

像C一样，Go也使用分号(;)断句，不同于C的是，源代码中可以不出现分号。
词法分析器(lexer)会自动插入分号，因此，大部分情况下，编写代码时不必手动输入分号。


The rule is this. If the last token before a newline is an identifier (which includes words like  `int`  and  `float64`), a basic literal such as a number or string constant, or one of the tokens

规则是这样的。如果行尾是标识符号, (比如 int and float64))，或者数字字符串这类字面量，或者是以下标记之一

```go
break continue fallthrough return ++ -- ) }
```


the lexer always inserts a semicolon after the token. This could be summarized as, “if the newline comes after a token that could end a statement, insert a semicolon”.

这个规则可以简单理解为，“在可以断句的地方，插入分号”。

> 译：初看这个说法，有点搞笑，但细想还真是这么回事。
> 经历一些项目后，不难发现，有些复杂逻辑背后的目标其实很简单，几个字就概括出来。
> 但实现成代码就会异常复杂。如果读者能了解复杂行为背后的目标，那就很容易理解了。
> 所以这个有点“搞笑”的话，应该也是golang开发者的一个目标吧。

A semicolon can also be omitted immediately before a closing brace, so a statement such as needs no semicolons. 

两个闭合的括号之后也能省略分号。比如下面这种情况就不需要分号：

```go
    go func() { for { dst <- <-src } }()
```

Idiomatic Go programs have semicolons only in places such as  `for`  loop clauses, to separate the initializer, condition, and continuation elements. They are also necessary to separate multiple statements on a line, should you write code that way.

Go 程序中常在`for`循环中使用分号分隔语句(`for i=0; i<1; i++ {}`)。
有时，也会用分号分隔一行代码存在多条语句的情况。

One consequence of the semicolon insertion rules is that you cannot put the opening brace of a control structure (`if`,  `for`,  `switch`, or  `select`) on the next line. If you do, a semicolon will be inserted before the brace, which could cause unwanted effects. Write them like this

由于自动插入分号的规则，不能在控制结构(`if`,  `for`,  `switch`, or  `select`)中换行写大括号。
如果一定要换行写大括号，词法分析器还会在行尾添加一个大括号，这会引发一些异常问题。
应该像下面这样写：

```go
if i < f() {
    g()
}
```

not like this

下面这种是错的：

```go
if i < f()  // wrong!
{           // wrong!
    g()
}
```

> 译： 大括号必须换行的要求原来是为了方便记法分析。我还以为只是单纯为了保持代码格式一致。



## Control structures 控制结构

The control structures of Go are related to those of C but differ in important ways. There is no  `do`  or  `while`  loop, only a slightly generalized  `for`;  `switch`  is more flexible;  `if`and  `switch`  accept an optional initialization statement like that of  `for`;  `break`  and  `continue`  statements take an optional label to identify what to break or continue; and there are new control structures including a type switch and a multiway communications multiplexer,  `select`. The syntax is also slightly different: there are no parentheses and the bodies must always be brace-delimited.

Go中的控制结构和C很像，但差异很大。
没有`do`或者`while`循环了。但有加强版本的`for`；有更灵活的`switch`；
`if`和`switch`都能使用类似`for`中的 initialization 语句；
`break`和`continue`标签仍然保留了下来；
新增加了用于类型选择和多路通信复用的`select`；
语法上有很大的变化，用于条件判断的小括号不需要了，但用于定界的大括号是必须存在的。
 


### If 条件控制

In Go a simple  `if`  looks like this:

Go中if语句一般是下面这样：

```go
if x > 0 {
    return y
}
```

Mandatory braces encourage writing simple  `if`  statements on multiple lines. It's good style to do so anyway, especially when the body contains a control statement such as a  `return`  or  `break`.

由于强制要求不能省略大括号，所以简单的`if`判断也要写成多行代码。
这么做是有好处的，尤其是代码中包含`return`或者`break`这样的控制语句时。

> 译：这种硬性要求在golang中有很多，但确实是有好处的。
> 比如这个要求，就能从根本是解决维护旧代码中，调整单行if语句时，由于忽略{}而常常出现的bug）
> `if`和`switch`支持 initialization 语句，这非常便于使用局部变量

`


Since  `if`  and  `switch`  accept an initialization statement, it's common to see one used to set up a local variable.

因为`if`和`switch`支持 initialization 语句，经常使用这个特性设置局部变量。

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

In the Go libraries, you'll find that when an  `if`  statement doesn't flow into the next statement—that is, the body ends in  `break`,  `continue`,  `goto`, or  `return`—the unnecessary`else`  is omitted.

在 Go 代码库中,会发现很多省略 `else` 子句的代码，因为 `if` 语句的执行体以 `break`,  `continue`,  `goto` 结束。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

This is an example of a common situation where code must guard against a sequence of error conditions. The code reads well if the successful flow of control runs down the page, eliminating error cases as they arise. Since error cases tend to end in  `return`  statements, the resulting code needs no  `else`  statements.

下面的救命是一种很常见的场景。
代码中检查了每个可能出错的环节，只要代码执行到函数最后，说明所有异常问题都没有发生。
在每个`if`条件处理中都用 return 返回 error ，所以代码中都不需要出现 `else` 语句。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

> 译：这种观点在《代码整洁之道》或者《重构》中也看到过。但周围还有很多搞技术的人不知道。老能看到俄罗斯套娃一样的代码。



### Redeclaration and reassignment 重定义和重新赋值

An aside: The last example in the previous section demonstrates a detail of how the  `:=`short declaration form works. The declaration that calls  `os.Open`  reads,

上面的示例代码也展示了`:=`符号的用法。
在`os.Open`这行代码中，声名了两个变量`f`和`err`：

```go
f, err := os.Open(name)
```

This statement declares two variables,  `f`and  `err`. A few lines later, the call to  `f.Stat`reads, which looks as if it declares  `d`  and  `err`.

在`f.Stat`代码中，看似又声名了两个变量`d`和`err`：

```go
d, err := f.Stat()
```

Notice, though, that  `err`  appears in both statements. This duplication is legal:  `err`  is declared by the first statement, but only  _re-assigned_  in the second. This means that the call to  `f.Stat`  uses the existing  `err`  variable declared above, and just gives it a new value.

注意，`err`出现在两个声名的代码中，但这是合法的。
第一次出现`err`是声名此变量，第二次出现`err`中对上一次声名的变量重新覆值。
也就是说`err`在`f.Stat()`调用之前就已经声名，`f.State()`只是赋予一个新值给`err`。


In a  `:=`  declaration a variable  `v`  may appear even if it has already been declared, provided:

-   this declaration is in the same scope as the existing declaration of  `v`  (if  `v`  is already declared in an outer scope, the declaration will create a new variable §),
-   the corresponding value in the initialization is assignable to  `v`, and
-   there is at least one other variable in the declaration that is being declared anew.

仅在以下几种情况下允许使用  `:=`  给变量 `v` 重新赋值：

-   声名代码与已经存在的变量`v`作用域相同。如果`v`已经在外层作用域声名过，那么这次声名会创建一个新的变量，
-   变量类型相符的value才能给`v`赋值
-   声名中至少包含一个新的变量


This unusual property is pure pragmatism, making it easy to use a single  `err`  value, for example, in a long  `if-else`  chain. You'll see it used often.

这种不常见的特性纯粹是因为实用主义。我们能在很长的`if-else`代码中仅仅使用一个`err`变量。
你应该能经常看到这种用法。


§ It's worth noting here that in Go the scope of function parameters and return values is the same as the function body, even though they appear lexically outside the braces that enclose the body.

§ 另外，虽然在词法上，函数参数与函数返回值写在函数体的大括号之外，但实际上它的作用域与函数体是一样的。



### For 循环

The Go  `for`  loop is similar to—but not the same as—C's. It unifies  `for`  and  `while`  and there is no  `do-while`. There are three forms, only one of which has semicolons.

Go的`for`循环结合了C中`for`和`while`的功能，不过不支持`do-while`的功能。
一共有三种形式，只有一种必需要用分号。

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

Short declarations make it easy to declare the index variable right in the loop.

在`for`循环中定义一个索引变量，简洁又方便。

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```


If you're looping over an array, slice, string, or map, or reading from a channel, a  `range`clause can manage the loop.

使用`range`遍历 array,slice,string,map 或者读取 channel 的消息：

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

If you only need the first item in the range (the key or index), drop the second:
 

如果只需用到第一个数据项(key/index)，直接省略第二个就行了。

```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

If you only need the second item in the range (the value), use the  _blank identifier_, an underscore, to discard the first:

如果需要用到第二个数据项(value)，用`blank`标识符(_)占位，忽略掉即可：

> 译：golang所有声名的变量必须使用，否则编译失败，所以不使用的变量，需要使用 _ 符号占位

```go
sum := 0
for _, value := range array {
    sum += value
}
```

The blank identifier has many uses, as described in  [a later section](https://golang.org/effective_go.html#blank).

`blank`标识符还有很多种用法，详细描述参考[本文 blank 这一节](https://golang.org/effective_go.html#blank).

For strings, the  `range`  does more work for you, breaking out individual Unicode code points by parsing the UTF-8. Erroneous encodings consume one byte and produce the replacement rune U+FFFD. (The name (with associated builtin type)  `rune`  is Go terminology for a single Unicode code point. See  [the language specification](https://golang.org/ref/spec#Rune_literals)  for details.) The loop
 
遍历字符串时,`range` 会自动解析UTF-8编码，分离 Unicode 码点。
错误的编码只消费一个 Byte ，并使用`rune`类型的 `U+FFFD` 代替 value。
（`rune`是内置类型，表示 Unicode code point ，详细解释参考 [Rune_literals](https://golang.org/ref/spec#Rune_literals)）

> 译：Unicode 编码每个字符 rune 占用的空间大小 byte 是不同的。
> UTF-8 编码方案中，英文字符占用 `1byte=8bit` ，但中文字符可能占用 `3byte=24bit` 。
> 如果像 C 语言遍历数组那样遍历 `char[]` 类型的 UTF-8 字符串，每个遍历项用 `char` 保存，因为 `char`固定占用 `1byte=ibit` ，那么遍历出一个不完整的中文字符串。


以下循环代码：

```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

prints

将会输出

```txt
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

Finally, Go has no comma operator and  `++`and  `--`  are statements not expressions. Thus if you want to run multiple variables in a  `for`  you should use parallel assignment (although that precludes  `++`  and  `--`).
 
最后，Go中没有逗号(comma)运算符，并且`++`和`--`是语句，不是表达式。
所以如果想在`for`中使用多个变量，只能使用批量赋值语句，避免使用`++`和`--`。

```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

### Switch 切换

Go's  `switch`  is more general than C's. The expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found, and if the  `switch`  has no expression it switches on`true`. It's therefore possible—and idiomatic—to write an  `if`-`else`-`if`-`else`  chain as a  `switch`.

Go中的`switch`比C用途更广。
它不要求表达式必须是常量或整型，只要从上往下找到第一个匹配的 case 即可，
如果`switch`没有表达式，那么找到第一个case表达式为`true`的。
当然`switch`的实现功能，也能用`if-else-if-else`实现。


```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

There is no automatic fall through, but cases can be presented in comma-separated lists.

switch 不会自动执行多个连续的 case 语句，但可以用逗号在一个 case 语句中分隔多个条件。

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

Although they are not nearly as common in Go as some other C-like languages,  `break`statements can be used to terminate a  `switch`  early. Sometimes, though, it's necessary to break out of a surrounding loop, not the switch, and in Go that can be accomplished by putting a label on the loop and "breaking" to that label. This example shows both uses.

虽然 Go 与其他 C 系语言的 switch 有一点区别，但`switch`中也能用`break`提前结束`switch`。
有些，还会想提前结束 switch 外层的 for 循环。
在 Go 语文中，可以在循环外定义一个 label 标签，然后执行`break Loop`时就会直接跳出循环。
请看下面的示例。

>  TODO 不理解为什么 Loop 标签要定义在 for 循环的前，而不是后面呢？不过还好自己实践过程遇到类似需要，都会封装一个函数执行 for 循环，然后在 case 中直接 return 。与本节最后一个示例一样，这种办法能避免用 label 标签。

```go
Loop:
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				break
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 >= len(src) {
				err = errShortInput
				break Loop
			}
			if validateOnly {
				break
			}
			size = 2
			update(src[n] + src[n+1]<<shift)
		}
	}
```


Of course, the  `continue`  statement also accepts an optional label but it applies only to loops.

当然`continue`语句也可以使用`label`，但`continue`仅能在循环中使用。


To close this section, here's a comparison routine for byte slices that uses two  `switch`statements:

看一个使用`switch`比较 byte slice 的 routine 示例结束本节吧：

```go
// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}
```



### Type switch 类型选择

A switch can also be used to discover the dynamic type of an interface variable. Such a  _type switch_  uses the syntax of a type assertion with the keyword  `type`  inside the parentheses. If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. It's also idiomatic to reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case.

switch 也可以用来识别 interface 的动态类型。
一般在小括号包裹的`type`关键字进行类型断言。如果在 switch 表达式内声名一个变量，变量类型就和 case 中一致。
当然，也能直接在 case 中使用这个变量名称，效果等同于在每个 case 中各声名了一个名称相同，但类型不同的变量。

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```



## Functions 函数

### Multiple return values 多返回值

One of Go's unusual features is that functions and methods can return multiple values. This form can be used to improve on a couple of clumsy idioms in C programs: in-band error returns such as  `-1`for  `EOF`  and modifying an argument passed by address.

另一个Go的亮点是，函数支持多返回值。
这个特点可用来解决C中遗存已久的麻烦：
通过返回值是 `-1` `EOF`，来判断操作成功或失败，
再通过传递指针参数，来返回额外额外的信息。


In C, a write error is signaled by a negative count with the error code secreted away in a volatile location. In Go,  `Write`  can return a count  _and_  an error: “Yes, you wrote some bytes but not all of them because you filled the device”. The signature of the  `Write`method on files from package  `os`  is:

在C中，判断`write()`失败原因的错误码隐藏在返回参数中：
`count>=0`表示写入成功的字节数；
`count<0`表示失败原因。
在Go中，`write`能同时返回两个参数 count 和 error 。
这能表达出C中无法区分的一种情况：“成功写入了个字节，但设备空间已满，其他数据写入失败”。
在 package `os` 中， `Write`方法定义如下 ：

```go
func (file *File) Write(b []byte) (n int, err error)
```


and as the documentation says, it returns the number of bytes written and a non-nil  `error`  when  `n != len(b)`. This is a common style; see the section on error handling for more examples.

像文档描述的一样，它返回成功写入的字节数 n ，如果 `n != len(b)` ，返回非 nil 的 `error` 。
这种风格很普遍，后面 error handle 一节有更多的示例，
 
A similar approach obviates the need to pass a pointer to a return value to simulate a reference parameter. Here's a simple-minded function to grab a number from a position in a byte slice, returning the number and the next position.

这个简单的方法，可以避免用指针模拟引用参数，作为函数返回值。
下面有一个示例，在 byte slice 中查找连续的 number 字符，然后返回这个数值，和下一次执行查找过程的数组起始下标。

```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```


You could use it to scan the numbers in an input slice  `b`  like this:

在输入 slice 数据中，找到数值的代码示例如下所示：

```go
    for i := 0; i < len(b); {
        x, i = nextInt(b, i)
        fmt.Println(x)
    }
```

### Named result parameters 可命名的返回参数

The return or result "parameters" of a Go function can be given names and used as regular variables, just like the incoming parameters. When named, they are initialized to the zero values for their types when the function begins; if the function executes a  `return`  statement with no arguments, the current values of the result parameters are used as the returned values.

Go中函数返回参数可以像普通变量一样命名并使用，就跟输入参数一样。
当函数开始时，命名返回参数会被初始化为 zero 值。
如果函数执行到一个没有参数的 return 语句，那么命名参数的当前值就作为函数返回值。

> 译： zero 值，不一定是数值 0 。根据具体类型，这个值可能是 nil, 0, "" 或其他。


The names are not mandatory but they can make code shorter and clearer: they're documentation. If we name the results of  `nextInt`  it becomes obvious which returned  `int`  is which.

命名不是强制的，善加利用能使代码更简洁：起到文档的效果。
如果我们给`nextInt`函数返回值命名，那就很容易知道每个返回参数是干什么用的了。

```go
func nextInt(b []byte, pos int) (value, nextPos int) {
```


Because named results are initialized and tied to an unadorned return, they can simplify as well as clarify. Here's a version of  `io.ReadFull`  that uses them well:

因为命名参数会自动初始化并返回，它能使代码十分干净。
看看下面这个版本的`io.ReadFull`函数棒不棒：

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```



### Defer 延迟执行

Go's  `defer`  statement schedules a function call (the  _deferred_  function) to be run immediately before the function executing the  `defer`  returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file.

Go的`defer`语句能让函数调用延迟到当前函数结束前执行不（即 `推迟执行` 函数）。
这种特性在其他语言中不太常见，但利用它进行资源回收，十分有用，尤其是函数有很多返回路径的情况。
最典型使用场景就是解锁 mutex 互斥量或者关闭文件。

```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```


Deferring a call to a function such as  `Close`has two advantages. First, it guarantees that you will never forget to close the file, a mistake that's easy to make if you later edit the function to add a new return path. Second, it means that the close sits near the open, which is much clearer than placing it at the end of the function.

延迟`Close`调用有两个好处，首先，它能保证不论后期怎么维护或调整代码，都不会忘掉关闭文件的事情，
其次，关闭和打开文件的代码可以紧挨着，这比在函数开头打开，函数末尾关闭清晰的多。
 
The arguments to the deferred function (which include the receiver if the function is a method) are evaluated when the  _defer_ executes, not when the  _call_  executes. Besides avoiding worries about variables changing values as the function executes, this means that a single deferred call site can defer multiple function executions. Here's a silly example.

传递给延迟函数的参数，是`defer`语句执行时的值，而不是真正 _调用_ 延迟函数时的值。
所以不必担心延迟函数执行时，相关参数值会改变。
同时，这也说明，即使只出现一次 defer 函数代码，也可能会启动多次延迟函数执行。
如下所示：

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```


Deferred functions are executed in LIFO order, so this code will cause  `4 3 2 1 0`  to be printed when the function returns. A more plausible example is a simple way to trace function execution through the program. We could write a couple of simple tracing routines like this:

延迟函数是按后进先出(LIFO)的顺序执行的。因此上面的代码会在函数返回时输出`4 3 2 1 0 `。
一个更合理的示例是，用defer追踪函数的执行。比如可以这样写一对简单的追踪程序。

```go
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
```


We can do better by exploiting the fact that arguments to deferred functions are evaluated when the  `defer`  executes. The tracing routine can set up the argument to the untracing routine. This example:

我们可以改造这个程序，让它用起来更方便：
在`defer`执行时，先调用`trace`。
在延迟函数执行时，会自动调用`un`并传递与调用`trace`函数一样的参数值。

```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

prints

以上程序会输出

```txt
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

For programmers accustomed to block-level resource management from other languages,  `defer`  may seem peculiar, but its most interesting and powerful applications come precisely from the fact that it's not block-based but function-based. In the section on  `panic`  and  `recover`  we'll see another example of its possibilities.

对于习惯了块级资源管理的程序员来说，`defer`看起来有些古怪。
因为`defer`是函数级的，但这也正是它有趣且强大的地方。
在后面的`panic`和`recover`一节，我们会看到更多示例。

> 译： block-level 块级资源管理是指类似 C+＋ 语言中，变量作用域在两个大括号之间`{}`，即程序块之间。
> 在此作用域中定义的变量，如果超出这个作用域，会立即销毁，即调用此变量的析构函数。

> 译：相比 C++ 中 class 的 destructer ，开始时我会觉得 defer 比较难用，上向说的那些功能，用 destructer 可以一行代码实现。
> 但考虑到 destructer 的时机并非确定，经过不同中版本的 C+＋ 编译器优化后，超出作用域的变量，不一定会立即销毁并调用 destructer 。
> Go 语言力求使行为特性清晰且明确，更不会混合 `defer` 功能与垃圾回收的逻辑的。了解一下 Go 和 Java 语言垃圾回收的逻辑，就更能明白 Go 不可以使用 destructer 机制来实现 defer 蔪。

TODO Golang FAQ 中相关 defer 情况说明呢？
 

## Data 数据

### Allocation with  `new` 使用 `new` 分配变量存储空间（内存）

Go has two allocation primitives, the built-in functions  `new`  and  `make`. They do different things and apply to different types, which can be confusing, but the rules are simple. Let's talk about  `new`  first. It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not  _initialize_  the memory, it only  _zeros_it. That is,  `new(T)`  allocates zeroed storage for a new item of type  `T`  and returns its address, a value of type  `*T`. In Go terminology, it returns a pointer to a newly allocated zero value of type  `T`.

Go中有两种分配原语，内置函数是`new`和`make`。
这俩函数很容易混淆，但它们分别用于完全不同的目的，区别很大。
区分的规则也很简单。
先说`new`,这是内置的分配内存的函数，它不会初始化内存，只会将其清零(zeros)。
即`new(T)`会分配类型为`T`的内存空间，并把这个空间的数据清零后，返回类型为`*T`的内存地址。 

Since the memory returned by  `new`  is zeroed, it's helpful to arrange when designing your data structures that the zero value of each type can be used without further initialization. This means a user of the data structure can create one with  `new`  and get right to work. For example, the documentation for  `bytes.Buffer`  states that "the zero value for  `Buffer`  is an empty buffer ready to use." Similarly,  `sync.Mutex`  does not have an explicit constructor or  `Init`  method. Instead, the zero value for a  `sync.Mutex`  is defined to be an unlocked mutex.


因为`new`返回的内存数据都经过 zero ，我们的自定义结构体都可以不初始化了。
也就是说，我们用`new`创建一个指定类型的变量后，就能直接使用了。
比如关于`bytes.Buffer`的文档就这样描述：“zero的`Buffer`就是随时可用的空 buffer”。
同样，`sync.Mutex`也没有显示初始化的`Init`方法。
zero 的 `sync.Mutex` 就是解锁状态的mutex。

The zero-value-is-useful property works transitively. Consider this type declaration.

zero值能直接使用的特性非常棒。 看看下面的类型声名。

```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

Values of type  `SyncedBuffer`  are also ready to use immediately upon allocation or just declaration. In the next snippet, both  `p`  and  `v`  will work correctly without further arrangement.

`SyncedBuffer`类型的变量一经声名就能直接使用。
下面的代码片断中，`p`和`v`都能直接使用，不需要其他初始化代码了。

```go
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

### Constructors and composite literals 构造函数与复合字面量

Sometimes the zero value isn't good enough and an initializing constructor is necessary, as in this example derived from package  `os`.

有时 zero 值不能直接使用，我们需要更进一步的初始化，即构造函数(constructor)。
下面示例是来自`package os`。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

There's a lot of boiler plate in there. We can simplify it using a  _composite literal_, which is an expression that creates a new instance each time it is evaluated.

这样有些啰嗦。
我可以简化成只用一句 _composite literal 复合字面量_ 表达式就创建一个实例并赋值。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

Note that, unlike in C, it's perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns. In fact, taking the address of a composite literal allocates a fresh instance each time it is evaluated, so we can combine these last two lines.

注意，这跟 C 语言不一样，我们能返回局部变量的地址：函数返回时，变量存量空间仍然保留。
实际上，composite literal 执行的时候，就已经分配了地址空间了。
我们能把最后两行合并。

```go
    return &File{fd, name, nil, 0}
```

The fields of a composite literal are laid out in order and must all be present. However, by labeling the elements explicitly as  _field_`:`_value_  pairs, the initializers can appear in any order, with the missing ones left as their respective zero values. Thus we could say

composite literal 中必须按顺序写出相关结构的所有字段。
如果显示指定字段名，我们就能按任意顺序，初始化任意的字段，没有列出的字段，初始化为 zero 。
像下面这样：

```go
    return &File{fd: fd, name: name}
```

As a limiting case, if a composite literal contains no fields at all, it creates a zero value for the type. The expressions  `new(File)`  and  `&File{}`  are equivalent.

如果 composite literal 未包含任何字段，就赋值为zero。 
表达式`new(File)`和`&File{}` 是等效的。

Composite literals can also be created for arrays, slices, and maps, with the field labels being indices or map keys as appropriate. In these examples, the initializations work regardless of the values of  `Enone`,  `Eio`, and  `Einval`, as long as they are distinct.

composite literal 也能创建  arrays, slices, maps ，字段名会自动适配为 array 的 索引 或 map 的 键。
下面的示例中，只要 `Enone,Eio,Einval` 的值不同，就能正确初始化。

```go
const (
    Enone = 0
    Eio = 1
    Einval = 3 // 取值可以不连续
    // Einval  = "4" // 如果是字符串，就不能编译通过
)
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```



### Allocation with  `make` 使用 `make` 分配变量存储空间（内存）

Back to allocation. The built-in function  `make(T,` _args_`)`  serves a purpose different from  `new(T)`. It creates slices, maps, and channels only, and it returns an  _initialized_(not  _zeroed_) value of type  `T`  (not  `*T`). The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use. A slice, for example, is a three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity, and until those items are initialized, the slice is  `nil`. For slices, maps, and channels,  `make`  initializes the internal data structure and prepares the value for use. For instance,

回到 allocation （资源分配）的话题。
内置函数`make(T, args)`与`new(T)`使用目的完全不同。
它仅用于创建 slices, maps, channels ，并返回 _initialized 初始化过_ 的（不是 _zeroed 置零_ ）的值，且类型为`T`的变量（不是`*T`）。
出现这种差异的本质原因是这三种类型是引用类型，必须在使用前初始化。
比如， slice 由三项描述符组成，分别指向 data （数组的数据），length （长度），capacity（容量），
在三个描述符未初始化前， slice 的值是 nil 。
对于 slices, maps, channels 来说，`make`用于初始化结构体内部数据并赋值。比如，

```go
make([]int, 10, 100)
```

allocates an array of 100 ints and then creates a slice structure with length 10 and a capacity of 100 pointing at the first 10 elements of the array. (When making a slice, the capacity can be omitted; see the section on slices for more information.) In contrast,  `new([]int)`  returns a pointer to a newly allocated, zeroed slice structure, that is, a pointer to a  `nil`  slice value.

分配了一个包含100个 int 的 array，并创建了一个 length 为10， capacity 为100的 slice ，指向 array 的前10个元素。
（创建 slice 时， capacity 可以省略，查看有关 slice 的章节，了解更多信息。）
与之对照，`new([]int)`返回一个 zero 的 slice 结构体，也就是一个指向值为 nil 的 slice 。

These examples illustrate the difference between  `new`  and  `make`.

下面代码阐明了`new`和`make`的不同。

```go
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```

Remember that  `make`  applies only to maps, slices and channels and does not return a pointer. To obtain an explicit pointer allocate with  `new`  or take the address of a variable explicitly.

记住，`make`仅用于 maps, slices, channels ，返回的也不是指针。
只有使用`new`或者对变量执行取地址操作`&File{}`才能得到指针。



### Arrays 数组

Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation, but primarily they are a building block for slices, the subject of the next section. To lay the foundation for that topic, here are a few words about arrays.

在详细规划内存布局时， array 是很有用的，有时它还能避免过多的内存分配，它的主要作用是构造 slice ，不过那就是下一节的主题了，这里先说几句做铺垫。

There are major differences between the ways arrays work in Go and C. In Go,

-   Arrays are values. Assigning one array to another copies all the elements.
-   In particular, if you pass an array to a function, it will receive a  _copy_  of the array, not a pointer to it.
-   The size of an array is part of its type. The types  `[10]int`  and  `[20]int`  are distinct.

下面是 C 与 Go 中有关 array 的主要区别。在 Go 中，
- Arrays 是值类型，两个 array 之间赋值会复制所有元素。
- 具体来讲，如果函数参数是数据，函数将接收一个 array 的完整副本（深拷贝），而不是指针。
- array 大小是类型的一部分。 `[10]int`和`[20]int`是不同类型。
 
The value property can be useful but also expensive; if you want C-like behavior and efficiency, you can pass a pointer to the array.

值类型有用，但代价高；
如果你想要类 C 的行为和效率，可以传递array的指针做参数。

```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator
```

But even this style isn't idiomatic Go. Use slices instead.

但这种方式不常用，在 Go 语言中经常使用的 slice 。

### Slices 切片

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

slice 对 array 做了封装，提供更通用、强大、方便的管理数据序列的接口。
除了转换矩阵这种需要明确维度的操作外，Go中大部分编程操作通过 slice 完成。

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array. A  `Read`function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the  `Read`  method of the  `File`  type in package  `os`:
 
slice 保存了对底层 array 的引用，如果你把一个 slice 赋值给另外一个 slice ，两个 slice 会引用同一个 array 。
如果一个函数接收 slice 参数，那么函数内部对 slice 的修改，都能影响调用方的参数，这和传递底层 array 指针的效果类似。
比方说，`Read`函数可以使用 slice 作为参数，slice 的长度刚好用来限制能读取的最大数据量，这种方法很适合代替以 data 指针 与 count 容量 作为参数的方式。以下是 `package os`中`File`类型的`Read`方法定义：

```go
func (f *File) Read(buf []byte) (n int, err error)
```

The method returns the number of bytes read and an error value, if any. To read into the first 32 bytes of a larger buffer  `buf`,  _slice_  (here used as a verb) the buffer.

这个方法返回成功读取的字节数 n，以及标明是否遇到错误的 err 。
用下面这种方法，可实现仅读取文件前32字节，并将读取到的数据填充到缓冲区 buf 中的前32字节中。
这里使用了 slicing 切割缓冲区的方法。

```go
    n, err := f.Read(buf[0:32])
```

Such slicing is common and efficient. In fact, leaving efficiency aside for the moment, the following snippet would also read the first 32 bytes of the buffer.

这种 slicing 方式常见而高效。
如果不考虑效率问题，下面的代码也能实现读取前32字节到缓冲区的目的。

```go
    var n int
    var err error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
```

The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The  _capacity_  of a slice, accessible by the built-in function  `cap`, reports the maximum length the slice may assume. Here is a function to append data to a slice. If the data exceeds the capacity, the slice is reallocated. The resulting slice is returned. The function uses the fact that  `len`  and  `cap`are legal when applied to the  `nil`  slice, and return 0.
 
只要不超出 slice的底层数组的长度限制，也能改变 slice 的长度，只要对 slice 做一次切割(slicing)就行。
使用内置函数`cap`返回 slice 的容量(capacity)，这是 slice 当前能使用的最大长度。
下面的函数能向 slice 中追加数据。如果数据超出最大容量，则为 slice 重新分配空间。返回值就是追加数据后的 slice 。
函数`len`和`cap`能正确处理值为`nil`的 slice ，并返回 0。

```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

We must return the slice afterwards because, although  `Append`  can modify the elements of  `slice`, the slice itself (the run-time data structure holding the pointer, length, and capacity) is passed by value.

我们必须在最后返回 slice ，是因为 `Append` 能修改`slice`指向的底层 array 中的元素值，但 slice 本身（保存 data指针，length, capacity的数据结构）是作为值传递的。

The idea of appending to a slice is so useful it's captured by the  `append`  built-in function. To understand that function's design, though, we need a little more information, so we'll return to it later.

向 slice 中追加数据的操作用途很大，所以内置了`append`函数实现此功能。
我们还需要更多信息才好理解这个函数的设计，所以，一会还会谈到它。


### Two-dimensional slices 二维切片

Go's arrays and slices are one-dimensional. To create the equivalent of a 2D array or slice, it is necessary to define an array-of-arrays or slice-of-slices, like this:

Go的 array 和 slice 是一维的。想要创建二维 array 或 slice ，需要定义包含 array 的 array 或者包含 slice 的 slice 。

```go
type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte     // A slice of byte slices.
```

Because slices are variable-length, it is possible to have each inner slice be a different length. That can be a common situation, as in our  `LinesOfText`  example: each line has an independent length.

因为 slice 是变长，所以每个内部 slice 也能有不同的长度。这种用法很常见，比如下面的`LinesOfText`示例，每行长度都不一样。

```go
text := LinesOfText{
	[]byte("Now is the time"),
	[]byte("for all good gophers"),
	[]byte("to bring some fun to the party."),
}
```

Sometimes it's necessary to allocate a 2D slice, a situation that can arise when processing scan lines of pixels, for instance. There are two ways to achieve this. One is to allocate each slice independently; the other is to allocate a single array and point the individual slices into it. Which to use depends on your application. If the slices might grow or shrink, they should be allocated independently to avoid overwriting the next line; if not, it can be more efficient to construct the object with a single allocation. For reference, here are sketches of the two methods. First, a line at a time:

处理像素描述行时，就会需要二维（2D）的 slice 。有两种方法来实现。
一种是，每行独立分配 slice ；
另一种是，分配一个 array ， 将其分割成多块交由 slice 管理。
根据自己应用的实际情况选择使用哪种方法。
如果 slice 空间会增加或收缩， 应该选用第一种独立分配 slice 的方法，防止越界覆盖下一秆数据。
否则，第二种方法能一次分配所有空间，更高效一些。下面是两种方法的示例。

第一种方法，每次一行：

```go
// Allocate the top-level slice.
picture := make([][]uint8, YSize) // One row per unit of y.
// Loop over the rows, allocating the slice for each row.
for i := range picture {
	picture[i] = make([]uint8, XSize)
}
```

And now as one allocation, sliced into lines:

第二种方法，一次分配，再分割成多行：

```go
// Allocate the top-level slice, the same as before.
picture := make([][]uint8, YSize) // One row per unit of y.
// Allocate one large slice to hold all the pixels.
pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
// Loop over the rows, slicing each row from the front of the remaining pixels slice.
for i := range picture {
	picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```



### Maps 键值对

Maps are a convenient and powerful built-in data structure that associate values of one type (the  _key_) with values of another type (the  _element_  or  _value_). The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. Slices cannot be used as map keys, because equality is not defined on them. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

maps 是内置的方便而强大的数据类型，用于将一种类型的值（键，key）与另一种类型的值（元素，element, value）进行关联。
key 可以是任何能用等号(=)比较的类型，如 integer, floating point 和 complex numbers, strings, pointers, interface (只要动态类型支持等号比较)， structs 和 arrays。
slice 不能用做 maps 的 key ，因为无法用等号比较 slice 的值。
和 slice 类似， map 在底层保存某个数据类型的引用。
如果将 map 作为函数参数，并且在函数内部改变了 map 的值，那这种改变也会影响调用者的参数。
 
Maps can be constructed using the usual composite literal syntax with colon-separated key-value pairs, so it's easy to build them during initialization.

map 可以由使用分号分隔 key 和 value 对（键值对）的 composite literal （复合字面量）声名。

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

Assigning and fetching map values looks syntactically just like doing the same for arrays and slices except that the index doesn't need to be an integer.

设定和获取 map 值与 array / slice 的做法一样，只是索引(index)不必是 integer 了。

```go
offset := timeZone["EST"]
```

An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return  `0`. A set can be implemented as a map with value type  `bool`. Set the map entry to  `true`  to put the value in the set, and then test it by simple indexing.

如果尝试获取 map 中不存在的 key ，将返回 value 类型的 zero 值。
比如，如果 map 的 value 是 integer，那么查询不存在的 key 时，返回值是 0 。
set(集合） 类型可以用 value 是 bool 的 map 进行模拟。将 value 设置为 true 表示元素加入 set ，直接用索引操作就能确认 key 是否存在。

```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
```

Sometimes you need to distinguish a missing entry from a zero value. Is there an entry for  `"UTC"`  or is that 0 because it's not in the map at all? You can discriminate with a form of multiple assignment.
 
有时，需要区分 key 不存在（即zero值）与 value 是0值的情况。
比如，返回 0 时，是因为 key 为 "UTC" 还是因为 key 根本不存在于 map 中？
可以用多返回值(multiple assignment)来区分这些情况。

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

For obvious reasons this is called the “comma ok” idiom. In this example, if  `tz`  is present,  `seconds`  will be set appropriately and  `ok`  will be true; if not,  `seconds`  will be set to zero and  `ok`  will be false. Here's a function that puts it together with a nice error report:

按照惯例，在 seconds 后面加一个“, ok” 。
在下面的示例中，如果`tz`存在，则`seconds`就是对应的值，并且`ok`会被设置为 true ；否则，`seconds`会设置为 zero 值，`ok`被设置为 false。

```go
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

To test for presence in the map without worrying about the actual value, you can use the  [blank identifier](https://golang.org/effective_go.html#blank)  (`_`) in place of the usual variable for the value.

如果只想确认 map 中是否存在指定key，不关心其值是多少，可以使用 [blank identifier](https://golang.org/doc/effective_go.html#blank)

```go
_, present := timeZone[tz]
```

To delete a map entry, use the  `delete`  built-in function, whose arguments are the map and the key to be deleted. It's safe to do this even if the key is already absent from the map.

使用内置`delete`函数删除 map 中的元素，参数是 map 和需要被删除的 key 。即使 key 不存在，也能安全调用`delete`函数。

```go
delete(timeZone, "PDT")  // Now on Standard Time
```

### Printing 输出

Formatted printing in Go uses a style similar to C's  `printf`  family but is richer and more general. The functions live in the  `fmt`  package and have capitalized names:  `fmt.Printf`,  `fmt.Fprintf`,  `fmt.Sprintf`  and so on. The string functions (`Sprintf`  etc.) return a string rather than filling in a provided buffer.

Go 的格式化输出与 C 的 `printf`很像，但功能更丰富。相关函数位于 `fmt` package 中，以首字母大写命名，如`fmt.Printf` ， `fmt.Fprintf`， `fmt.Sprintf`等等。
字符串函数，如(Sprintf 等)会返回一个 string ，而不是向 buffer 填充字符串。

You don't need to provide a format string. 
For each of  `Printf`,  `Fprintf`  and  `Sprintf`there is another pair of functions, for instance  `Print`  and  `Println`. 
These functions do not take a format string but instead generate a default format for each argument. 
The  `Println`  versions also insert a blank between arguments and append a newline to the output while the  `Print`versions add blanks only if the operand on neither side is a string. 
In this example each line produces the same output.
 
也可以不提供 format string （格式化字符串）。
每个`Printf, Fprintf, Sprintf`都有一个对应函数，如`Print Println`。
这些函数不需要 format string 参数，因为它会给每个参数生成一个默认格式。
`Print`会在两个参数之间增加空格（只要两个参数都不是字符串），
`Println`不仅在参数之间增加空格，还会在行尾增加一个换行符号。下面的示例中，每行的输出结果都一样。

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

The formatted print functions  `fmt.Fprint`and friends take as a first argument any object that implements the  `io.Writer`interface; the variables  `os.Stdout`  and  `os.Stderr`  are familiar instances.

`fmt.Fprint`这类格式化输出函数的第一个参数必须是实现了`io.Writer`接口的对象；比如常见的`os.Stdout`和`os.Stderr`。

Here things start to diverge from C. First, the numeric formats such as  `%d`  do not take flags for signedness or size; instead, the printing routines use the type of the argument to decide these properties.

与 C 不同的是。`%d`这样的格式化符号不需要表示符号或大小的标记。
输出函数能直接根据参数类型，决定这些属性。

> 译：比如不存在 %ld 表示 long int，而 %d 表示int这种情况

```go
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
```

prints

输出

```go
18446744073709551615 ffffffffffffffff; -1 -1
```

If you just want the default conversion, such as decimal for integers, you can use the catchall format  `%v`  (for “value”); the result is exactly what  `Print`  and  `Println`would produce. Moreover, that format can print  _any_  value, even arrays, slices, structs, and maps. Here is a print statement for the time zone map defined in the previous section.

你还能用 通用格式化符号`%v` ，这个符号有一套默认输出格式，如对于整数来说，直接输出十进制整数；其实`Print`和`Println`的输出结果就这样的。
这个格式化符号甚至能打印 arrays, slices structs 和 maps 。下面的代码输出 time zone map 类型。

```go
fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)
```

which gives output:

输出：

```txt
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

For maps,  `Printf`  and friends sort the output lexicographically by key.
 
注意，maps 的 key 是乱序输出的。

When printing a struct, the modified format  `%+v`  annotates the fields of the structure with their names, and for any value the alternate format  `%#v`  prints the value in full Go syntax.

输出 struct 时，使用`%+v`这样的格式化输出符号能把字段名称一起输出，而`%#v`则按完整的Go语法规则输出值。

```go
type T struct {
    a int
    b float64
    c string
}
t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
```

prints

输出

```go
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
```

(Note the ampersands.) That quoted string format is also available through  `%q`  when applied to a value of type  `string`  or  `[]byte`. The alternate format  `%#q`  will use backquotes instead if possible. (The  `%q`format also applies to integers and runes, producing a single-quoted rune constant.) Also,  `%x`  works on strings, byte arrays and byte slices as well as on integers, generating a long hexadecimal string, and with a space in the format (`% x`) it puts spaces between the bytes.

注意 t 是 struct 指针，所以输出结果有与符号`&`。
 
使用`%q`格式化`string`或者`[]byte`时，也会输出双引号`""`。
使用`%#q`格式化符号，则会输出反引号`` ` ``。
`%q`也可用于 integers 和 runes 类型，此时会输出单引号`'`。
另外，`%x`也可用于 strings, byte arrays, byte slices, integers，其输出为十六进制字符串。如果在格式化符号前增加空格（`% x`），则输出的每个 bytes 之间也会以空格分隔。

> 译：以下示例是译者增加，参考： https://blog.golang.org/strings

```go
package main
import"fmt"
func main() {
var x uint64 = 18
var str string = "1汉字string"
var byt []byte = []byte("2汉字byte")
var rne []rune = []rune("3汉字rune")
 
fmt.Printf("%d, %x, %v\n", x, x, x)
fmt.Printf("%q, %#q, %x, % x\n", x, x, x, x)
fmt.Printf("%q, %#q, %x, % x\n", str, str, str, str)
fmt.Printf("%q, %#q, %x, % x\n", byt, byt, byt, byt)
fmt.Printf("%q, %#q, %x, % x\n", rne, rne, rne, rne)
}
```

> 输出

```go
18, 12, 18
'\x12', '\x12', 12,  12
"1汉字string", `1汉字string`, 31e6b189e5ad97737472696e67, 31 e6 b1 89 e5 ad 97 73 74 72 69 6e 67
"2汉字byte", `2汉字byte`, 32e6b189e5ad9762797465, 32 e6 b1 89 e5 ad 97 62 79 74 65
['3' '汉' '字' 'r' 'u' 'n' 'e'], ['3' '汉' '字' 'r' 'u' 'n' 'e'], [33 6c49 5b57 72 75 6e 65], [ 33  6c495b57  72  756e  65]
```
 
Another handy format is  `%T`, which prints the  _type_  of a value.

还有一个常用格式化符号是`%T`，用于出变量类型。

```go
fmt.Printf("%T\n", timeZone)
```

prints

输出

```go
map[string]int
```

If you want to control the default format for a custom type, all that's required is to define a method with the signature  `String() string`  on the type. For our simple type  `T`, that might look like this.

如果要控制自定义类型的默认输出格式，只需要给自定义类型增加一个`String() string`方法签名（signature）。
假设自定义类型是`T`，代码实现就是下面这样。

```go
package main
import"fmt"
 
type TPointer struct {
    a int
    b float64
    c string
}
func (t *TPointer) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
type TValue struct {
    a int
    b float64
    c string
}
func (t TValue) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
func main() {
    fmt.Printf("%v\n", TPointer{ 7, -2.35, "tPointer abc\tdef" })
    fmt.Printf("%v\n", &TPointer{ 7, -2.35, "tPointer abc\tdef" })
     
    fmt.Printf("%v\n", TValue{ 7, -2.35, "tValue abc\tdef" })
    fmt.Printf("%v\n", &TValue{ 7, -2.35, "tValue abc\tdef" })
}
```

to print in the format

输出以下格式

```go
{7 -2.35 tPointer abc   def}
7/-2.35/"tPointer abc\tdef"
7/-2.35/"tValue abc\tdef"
7/-2.35/"tValue abc\tdef"
```

(If you need to print  _values_  of type  `T`  as well as pointers to  `T`, the receiver for  `String`must be of value type; this example used a pointer because that's more efficient and idiomatic for struct types. See the section below on  [pointers vs. value receivers](https://golang.org/effective_go.html#pointers_vs_values)  for more information.)

注意，String() 方法签名的接收者是指针`*T`时，fmt.Printf 的参数也必须是指针，否则不会按自定义格式输出。
String() 的接收者是值类型`T`时，没有这种问题。但是用指针`*T`效率更高。详细情况参考[pointers vs. value receivers ](https://golang.org/doc/effective_go.html#pointers_vs_values)

Our  `String`  method is able to call  `Sprintf`because the print routines are fully reentrant and can be wrapped this way. There is one important detail to understand about this approach, however: don't construct a  `String`  method by calling`Sprintf`  in a way that will recur into your  `String`  method indefinitely. This can happen if the  `Sprintf`  call attempts to print the receiver directly as a string, which in turn will invoke the method again. It's a common and easy mistake to make, as this example shows.

`Sprintf`是可重入函数，所以在`String()`方法中可以再次调用`Sprintf`。但是要小心，别在`String()`方法中引发`String()`方法的调用，这会无限循环调用`String()`。
在`Sprintf`中直接将接收者当作 string 输出时，就会引起上面所述问题。这是一种常见的错误。
示例如下：

```go
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
}
```

It's also easy to fix: convert the argument to the basic string type, which does not have the method.

这个问题好解决，把参数强转成 string 类型即可，因为 string 类型没有使用 MyString 的  String() 签名方法，也就不会引起无限循环调用的问题了。

```go
type MyString string
func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}
```

In the  [initialization section](https://golang.org/effective_go.html#initialization)  we'll see another technique that avoids this recursion.

在[ initialization section](https://golang.org/doc/effective_go.html#initialization)一节，我们能用其他方法解决这个问题。

Another printing technique is to pass a print routine's arguments directly to another such routine. The signature of  `Printf`  uses the type  `...interface{}`  for its final argument to specify that an arbitrary number of parameters (of arbitrary type) can appear after the format.

另外一点值得说明的技术是，print 函数参数传递的过程。
`Printf`使用`...interface{}`作为最后一个参数，表示接收任意数量，任意类型的参数。

```go
func Printf(format string, v ...interface{}) (n int, err error) {
```

Within the function  `Printf`,  `v`  acts like a variable of type  `[]interface{}`  but if it is passed to another variadic function, it acts like a regular list of arguments. Here is the implementation of the function  `log.Println`we used above. It passes its arguments directly to  `fmt.Sprintln`  for the actual formatting.

在`Printf`函数中，可以把参数`v`当做`[]interface{}`使用。
但如果把`v`传递到其他函数使用，就要将其转为列表参数(regular list of arguments)。
下面是`log.Println`的实现代码，它将参数直接传递到`fmt.Sprintln`进行实际的格式化操作。

```go
// Println prints to the standard logger in the manner of fmt.Println.
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output takes parameters (int, string)
}
```

We write  `...`  after  `v`  in the nested call to  `Sprintln`  to tell the compiler to treat  `v`  as a list of arguments; otherwise it would just pass  `v`  as a single slice argument.

我们在调用`Sprintfln`时在参数`v`后面加了几个`...`，用来指明编译器将`v`作为列表变量(list of arguments)；如果不加`...`，`v`参数会被当做 slice 类型传递。

There's even more to printing than we've covered here. See the  `godoc`  documentation for package  `fmt`  for the details.

还有很多有关 print 的知识点没有提及，详细内容可能参考`godoc`中到`fmt`的说明。

By the way, a  `...`  parameter can be of a specific type, for instance  `...int`  for a min function that chooses the least of a list of integers:

顺带说一句，`...`参数也可以用来指明具体类型，比如下面以`...int`为参数的 min 函数，从一列 integers 中选取最小值。

```go
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```



### Append 追加
 
> NOTE append 会修改 slice 的特性是有坑的。 比如下面这种情况，

```go
a := []int{1}
b = append(a, 2,3,4,5) // a b 两个 slice 有可能指向一个底层 array 也有可以指向两个不同 array
fmt.Printf("%#v, %#v\n", a, b)
```


Now we have the missing piece we needed to explain the design of the  `append`  built-in function. The signature of  `append`  is different from our custom  `Append`  function above. Schematically, it's like this:

现在我们分析一下内建函数`append`的设计。这个`append`与我们之前自定义的`Append`有些区别，它的定义如下：

```go
func append(slice []_T_, elements ..._T_) []_T_
```

where  _T_  is a placeholder for any given type. You can't actually write a function in Go where the type  `T`  is determined by the caller. That's why  `append`  is built in: it needs support from the compiler.

`T`是表示任意类型的占位符。Go中无法实现一个参数类型`T`由调用者指定的函数。这正是为何`append`是内置类型的原因，因为它需要编译器支持。

> TODO 这么看来 Go 支持 泛型 功能的道路还很远呢

What  `append`  does is append the elements to the end of the slice and return the result. The result needs to be returned because, as with our hand-written  `Append`, the underlying array may change. This simple example

`append`的作用就是在 slice 中增加一个 element ，然后返回新的 slice 。
必须返回一个结果是因为，slice 底层的 array 可能改变。简洁示例如下：

```go
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```

prints  `[1 2 3 4 5 6]`. So  `append`  works a little like  `Printf`, collecting an arbitrary number of arguments.

结果输出[1 2 3 4 5 6]。`append`和`Printf`都能接收任意个参数。

But what if we wanted to do what our  `Append`  does and append a slice to a slice? Easy: use  `...`  at the call site, just as we did in the call to  `Output`  above. This snippet produces identical output to the one above.

如果我们把在 slice 后面追加一个 slice 怎么做呢？很简单，把 `...`  放到参数后面就行，和上面示例中`std.Output`用法。下面示例代码也输出[1 2 3 4 5 6]。

x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)

Without that  `...`, it wouldn't compile because the types would be wrong;  `y`  is not of type  `int`.

没有`...`是无法编译通过的，因为类型不正确，`y`的类型是`[]int`，而不是`int`。

## Initialization 初始化

Although it doesn't look superficially very different from initialization in C or C++, initialization in Go is more powerful. Complex structures can be built during initialization and the ordering issues among initialized objects, even among different packages, are handled correctly.

表面上看，Go 的初始化过程和 C/C++ 区别不大，其实 Go 的功能很强大的。
初始化过程不仅能构造复杂的结构体，还能正确处理不同 package 之间的初始化顺序。

### Constants 常量

Constants in Go are just that—constant. They are created at compile time, even when defined as locals in functions, and can only be numbers, characters (runes), strings or booleans. Because of the compile-time restriction, the expressions that define them must be constant expressions, evaluatable by the compiler. For instance,  `1<<3`  is a constant expression, while  `math.Sin(math.Pi/4)`  is not because the function call to  `math.Sin`  needs to happen at run time.

Go中的常量就是不会改动的变量(constant)。
常量必须是 numbers, characters(runes), strings, booleans。
即使在函数中定义的局部常量，也是在编译时期(compile time)创建的。
由于编译时的限制，定义常量的表达式必须能由编译器计算。
比如 `1<<3`是可用的常量表达式，而`math.Sin(math.Pi/4)`就不行，因为`math.Sin`是函数调用，必须在运行时(run time)执行。
 
In Go, enumerated constants are created using the  `iota`  enumerator. Since  `iota`  can be part of an expression and expressions can be implicitly repeated, it is easy to build intricate sets of values.

Go 中可以用`iota`创建枚举变量。
`iota`是表达式的一部分，能自动叠加，这种特性方便定义复杂的常量集合。

> 译： implicitly repeated 每行自动加1

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func main() {
	fmt.Println(YB, ByteSize(1e13))
}
```

The ability to attach a method such as  `String`  to any user-defined type makes it possible for arbitrary values to format themselves automatically for printing. Although you'll see it most often applied to structs, this technique is also useful for scalar types such as floating-point types like  `ByteSize`.

在自定义类型的增加`String`方法，能在 printing 输出时自动格式化。
虽然这个特性经常用于 struct 中，但其实也能用在`ByteSize`这种浮点数上。

> 译： scalar  标量，与向量相对

The expression  `YB`  prints as  `1.00YB`, while  `ByteSize(1e13)`  prints as  `9.09TB`.

表达式`ByteSize(YB)`会输出`1.00TB`，而`ByteSize(1e13) `会输出`9.09TB`。

The use here of  `Sprintf`  to implement  `ByteSize`'s  `String`  method is safe (avoids recurring indefinitely) not because of a conversion but because it calls  `Sprintf`  with  `%f`, which is not a string format:  `Sprintf`will only call the  `String`  method when it wants a string, and  `%f`  wants a floating-point value.

这个`ByteSize`的`String`方法实现是安全的（不会出现无限循环调用），并非因为类型转换，而是因为这里调用`Sprintf`时使用的参数`%f`，`Sprintf`只在期望 string 类型时，调用`String`方法，而使用`%f`时，期望的是 floating-point 类型。

> 译：并非这个原因，即表达式 `b/YB` 的结果转换成 float64 类型后，就没有了 ByteSize 类型的 String 方法



### Variables 变量

Variables can be initialized just like constants but the initializer can be a general expression computed at run time.

变量与常量的初始化方法类似，但变量初值是在 run time 计算的。

```go
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```



### The init function 初始化 init 函数

Finally, each source file can define its own niladic  `init`  function to set up whatever state is required. (Actually each file can have multiple  `init`  functions.) And finally means finally:  `init`  is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.

每个源文件都能定义`init`函数来设置一些初始状态。
（实际上每个文件可以包含多个`init`函数。）
在 package 中所有声名的变量及其 import （导入）的 package 都初始化完毕后，才会执行`init`函数。
 
Besides initializations that cannot be expressed as declarations, a common use of  `init`  functions is to verify or repair correctness of the program state before real execution begins.

`init`中可用于处理无法在 declaration （声明）中初始化的表达式，所以通常会在`init`中检查修正程序运行状态。

```go
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```



## Methods 方法

### Pointers vs. Values 指针类型与值类型

As we saw with  `ByteSize`, methods can be defined for any named type (except a pointer or an interface); the receiver does not have to be a struct.

可以给所有命名的类型（除 pointer 和 interface 外）定义 method ；receiver 不一定是 struct 。
就像上面`ByteSize`的例子就说明了这个特性。

> 译：这里的  method 可理解为类的成员方法； receiver 可理解为类的 this 指针。

In the discussion of slices above, we wrote an  `Append`  function. We can define it as a method on slices instead. To do this, we first declare a named type to which we can bind the method, and then make the receiver for the method a value of that type.

比如之前讨论到 slice 时提到的`Append`函数其实可以定义成 slice 的 method 。
为达到这个目的，我们先要定义一个类型，然后将这个类型作为 method 的 receiver 。

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above.
}
```

This still requires the method to return the updated slice. We can eliminate that clumsiness by redefining the method to take a  _pointer_  to a  `ByteSlice`  as its receiver, so the method can overwrite the caller's slice.

这种方式仍然需要返回更新后的 slice。将method 的 receiver 类型改成`ByteSlice`指针，就能在 method 中改变 receiver 的值。

```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

In fact, we can do even better. If we modify our function so it looks like a standard  `Write`  method, like this,

我们还能做的更好一点，如果把 Append 修改成下面这种标准`Write`方法的格式，

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```

then the type  `*ByteSlice`  satisfies the standard interface  `io.Writer`, which is handy. For instance, we can print into one.

于是，` *ByteSlice`类型就符合`io.Writer`掊口的定义。
这是很实用的技巧，
比如，能这样写入数据到 ByteSlice ：

```go
    var b ByteSlice
    fmt.Fprintf(&b, "This hour has %d days\n", 7)
```

We pass the address of a  `ByteSlice`  because only  `*ByteSlice`  satisfies  `io.Writer`. The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.

示例中使用`ByteSlice`的指针作为 Fprintf 的参数是因为`*ByteSlice `类型实现了`io.Writer`接口需要的方法（即`Write`方法的接收者类型是` *ByteSlice`）。
`pointer methods`，使用 指针 作为方法接收者，则必须通过 指针 调用此方法。
`value methods`，使用 值 作为方法接收者，则既能通过 值 也能通过 指针 调用此方法。

This rule arises because pointer methods can modify the receiver; invoking them on a value would cause the method to receive a copy of the value, so any modifications would be discarded. The language therefore disallows this mistake. There is a handy exception, though. When the value is addressable, the language takes care of the common case of invoking a pointer method on a value by inserting the address operator automatically. In our example, the variable  `b`  is addressable, so we can call its  `Write`  method with just  `b.Write`. The compiler will rewrite that to  `(&b).Write`  for us.

产生以上限制的原因是，`pointer methods`可以修改 方法接收者 。但使用 值 调用方法时，被修改的变量是 接收者 的一个拷贝，所以修改操作被忽略了。
Golang 语法不允许出现这样的错误。
不过，这有个例外情况。当 value 是`addressable`的，Golang 编译器会自动将通过 值 调用`pointer methods`的代码转换成通过 指针 调用。
在我们的示例中，虽然`Write`方法是`pointer methods`，但变量`b`是`addressable`的，所以直接写`b.Write()`这样的代码，也能调用`Write`方法。因为编译器替我们将代码改写成了`(&b).Write()`。
 

By the way, the idea of using  `Write`  on a slice of bytes is central to the implementation of  `bytes.Buffer`.

顺便一提，以上通过`Write`方法操作 slice bytes 的想法，已经在内置类`bytes.Buffer`中实现。



#### 译 有关 Pointer 与 Value 的问题总结

##### 1.为什么下面的代码没有复现此处所说问题？如果是因为 addressable ，那怎么样才能复现出上述问题呢？ TODO

参考[在线演示](https://play.golang.org/p/fCGBoXTD3wq)

```go
package main

import "fmt"

type ByteSlice struct {
	Tips string
}

func (p *ByteSlice) WriteByPointer(val string) () {
	p.Tips += " " + val
	fmt.Printf("WriteByPointer %s\n", p.Tips)
}
func (p ByteSlice) WriteByValue(val string) () {
	p.Tips += " " + val
	fmt.Printf("WriteByValue %s\n", p.Tips)
}

func main() {
	var tVal1 ByteSlice = ByteSlice{}
	tVal1.Tips = "tVal1"
	tVal1.WriteByPointer("t1")
	tVal1.WriteByValue("t2")
	
	
	var tPointer1 *ByteSlice = &ByteSlice{}
	tPointer1.Tips = "tPointer1"
	tPointer1.WriteByPointer("t3")
	tPointer1.WriteByValue("t4")
}
```

输出

```go
WriteByPointer tVal1 t1
WriteByValue tVal1 t1 t2
WriteByPointer tPointer1 t3
WriteByValue tPointer1 t3 t4
```

##### 2.什么是 addressable? 简单理解为，常量无法寻址，但变量肯定会存储在内存某个地方，可以被寻址
- [官方描述](https://golang.org/ref/spec#Address_operators)
- [译文](http://colobu.com/2017/02/01/golang-summaries/)
- [原文](http://www.tapirgames.com/blog/golang-summaries)
- 下面的值不能被寻址(addresses):
```txt
bytes in strings：字符串中的字节
map elements：map中的元素
dynamic values of interface values (exposed by type assertions)：接口的动态值
constant values：常量
literal values：字面值
package level functions：包级别的函数
methods (used as function values)：方法
intermediate values：中间值
function callings
explicit value conversions
all sorts of operations, except pointer dereference operations, but including:
channel receive operations
sub-string operations
sub-slice operations
addition, subtraction, multiplication, and division, etc.
注意， `&T{}`相当于`tmp := T{}; (&tmp)`的语法糖，所以`&T{}`合法不意味着`T{}`可寻址。
```
- 下面的值可以寻址:
```txt
variables
fields of addressable structs
elements of addressable arrays
elements of any slices (whether the slices are addressable or not)
pointer dereference operations
```



## Interfaces and other 接口和其他类型

### Interfaces 接口

Interfaces in Go provide a way to specify the behavior of an object: if something can do  _this_, then it can be used  _here_. We've seen a couple of simple examples already; custom printers can be implemented by a  `String`  method while  `Fprintf`  can generate output to anything with a  `Write`  method. Interfaces with only one or two methods are common in Go code, and are usually given a name derived from the method, such as  `io.Writer`  for something that implements  `Write`.

Golang 提供 `interface` 接口来实现 'object‘对象类似的功能：能做什么事，就是什么人。
我们其实已经看到过多个示例了。比如， 通过实现 `String()` method 来实现自定义输出格式的功能；还有使用 `Fprintf` 打印实现 `Write()` method 的类型。
只有一两个 method 的 interface 在Go代码中很常见。
并且 interface 的命名往往源于其实现的 method 方法名称，比如，实现了 `Write()` method 的 interface 称做`io.Writer`。

> 译： if something can d this, then is can be used here. 能不能通俗的翻译成：有奶便是娘？

A type can implement multiple interfaces. For instance, a collection can be sorted by the routines in package  `sort`  if it implements  `sort.Interface`, which contains  `Len()`,  `Less(i, j int) bool`, and  `Swap(i, j int)`, and it could also have a custom formatter. In this contrived example  `Sequence`  satisfies both.

并且一个`type`类型可以实现多个 interface 。
比如，如果一个数组集合（译：这里应该是专指数组集合，如 []string []int等）实现了`sort.Interface` interface 要求的 `Len()`, `Less(i, j int) bool`, 和 `Swap(i, j int)` 三个 method ,那它就能用`sort.Sort()`实现排序功能。
同时，还能再实现`fmt.Stringer` interface 要求的 `String()` method ,满足自定义输出格式功能。
下面这个刻意为之的例子中，Sequence type 就实现了 `sort.Interface` 和 `fmt.Stringer`  要求的几个method。

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sort"
)

func main() {
	seq := Sequence{6, 2, -1, 44, 16}
	sort.Sort(seq)
	fmt.Println(seq)
}

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	s = s.Copy() // Make a copy; don't overwrite argument.
	sort.Sort(s)
	str := "["
	for i, elem := range s { // Loop is O(N²); will fix that in next example.
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}
```



### Conversions 类型转换

The  `String`  method of  `Sequence`  is recreating the work that  `Sprint`  already does for slices. (It also has complexity O(N²), which is poor.) We can share the effort (and also speed it up) if we convert the  `Sequence`  to a plain  `[]int`  before calling  `Sprint`.

下面 `Sequence` 类型的 `String()` method 重用了 `fmt.Sprint([]int{})` 函数。
我们把 `Sequence` 转换成 `[]int` 类型，就能直接调用 `fmt.Sprint([]int{})` 函数了。

```go
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

This method is another example of the conversion technique for calling  `Sprintf`safely from a  `String`  method. Because the two types (`Sequence`  and  `[]int`) are the same if we ignore the type name, it's legal to convert between them. The conversion doesn't create a new value, it just temporarily acts as though the existing value has a new type. (There are other legal conversions, such as from integer to floating point, that do create a new value.)

这就是，在 `String()` method 中使用类型转换技术调用 `Sprintf` 方法的示例。
因为两个类型(`Sequence` `[]int`)本质是一样的，只是名称不同，所以可以合法（且安全）的在两个类型之前转换。
这次转换不会创建新值，他只是临时把已经存在的值当成另一个类型使用。
（还有另一种合法的转换方式，比如把 int 转换成 floating point 类型，此时就会创建一个新值。）


It's an idiom in Go programs to convert the type of an expression to access a different set of methods. As an example, we could use the existing type  `sort.IntSlice`  to reduce the entire example to this:

理所当然，Go 程序中也能对集合 set 类型执行类型转换。下面就是 Sequence 的另一种实现方法，因为使用了 `sort.IntSlice(s)`，所以比之前的方法少写了很多代码。

```go
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

Now, instead of having  `Sequence`  implement multiple interfaces (sorting and printing), we're using the ability of a data item to be converted to multiple types (`Sequence`,  `sort.IntSlice`  and  `[]int`), each of which does some part of the job. That's more unusual in practice but can be effective.

现在，不用给 Sequence 类型实现 Len() Less() Swap() 三个 method ，只是通过几次类型转换，我们就实现了相关的功能。当然，这种技术虽然管用，但实践中并不常用类型转换来实现排序功能。

### Interface conversions and type assertions 接口类型转换与类型断言

[Type switches](https://golang.org/effective_go.html#type_switch)  are a form of conversion: they take an interface and, for each case in the switch, in a sense convert it to the type of that case. Here's a simplified version of how the code under  `fmt.Printf`  turns a value into a string using a type switch. If it's already a string, we want the actual string value held by the interface, while if it has a  `String`  method we want the result of calling the method.

[Type switches](https://golang.org/doc/effective_go.html#type_switch)是一种类型转换：使用 switch 遍历 interface{} 的类型，每个 case 子句，都会将其转换成指定类型。下面简化的 fmt.Printf() 代码就是使用 type switch 中将 value 转换成 string 的示例。如果其本身就是 string 类型，则直接使用原始的字符串值；如果实现了 Stringer 接口要求的方法，则使用 String() 返回的字符串。

TODO `case string, *string:` 时 str 就不是 string 类型的

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

The first case finds a concrete value; the second converts the interface into another interface. It's perfectly fine to mix types this way.

第一个 case 返回原始字符串； 第二个 case 使用 String() 方法返回的字符串。
使用这种方式处理混合类型简直太完美了。


What if there's only one type we care about? If we know the value holds a  `string`and we just want to extract it? A one-case type switch would do, but so would a  _type assertion_. A type assertion takes an interface value and extracts from it a value of the specified explicit type. The syntax borrows from the clause opening a type switch, but with an explicit type rather than the  `type`  keyword:

如果我们只关心一种类型呢？
如果我们知道这个 value 中包含 string ，并且我们只想取出 string 使用，怎么做呢？
这时使用 `type assertion` 类型断言，就能直接将 interface 中指定类型的值取出来。type assertion 借鉴了 type switch 的语法，但使用显式的类型名，替换了 type 关键字：

```go
value.(typeName)
```

and the result is a new value with the static type  `typeName`. That type must either be the concrete type held by the interface, or a second interface type that the value can be converted to. To extract the string we know is in the value, we could write:

如果 typeName 是 interface 实现的类型， 或者该值可以转换成 typeName 为类型时，这里就能直接返回静态类型 typeName 的值。 可以像下面这样取出 string 类型的值：

```go
str := value.(string)
```

But if it turns out that the value does not contain a string, the program will crash with a run-time error. To guard against that, use the "comma, ok" idiom to test, safely, whether the value is a string:

但是，如果 value 不能转换为 string ，程序就会出现一个运行时错误，并崩溃。为避免这种问题，惯例是使用 `comma, ok` 语法判断 value 是否能转换为 string：

```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

If the type assertion fails,  `str`  will still exist and be of type string, but it will have the zero value, an empty string.

如果 type assertion 失败， str 变量仍然是 zero value 的 string 类型，也就是空字符串。

As an illustration of the capability, here's an  `if`-`else`  statement that's equivalent to the type switch that opened this section.

下面使用 type assertion 的 if else 语句 和 type switch 作用相同，做为对比，会更好理解一些。

```go
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
```



### Generality 通用性

If a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the type itself. Exporting just the interface makes it clear the value has no interesting behavior beyond what is described in the interface. It also avoids the need to repeat the documentation on every instance of a common method.

如果有个仅实现了 interface 的  type ，没有导出任何 method ，那也没必要把这个 type 导出。只导出 interface 可以保证 value 不会有超出 interface 描述的行为。同时也避免在所有 instance 中重复写一遍文档。

In such cases, the constructor should return an interface value rather than the implementing type. As an example, in the hash libraries both  `crc32.NewIEEE`  and  `adler32.New`  return the interface type  `hash.Hash32`. Substituting the CRC-32 algorithm for Adler-32 in a Go program requires only changing the constructor call; the rest of the code is unaffected by the change of algorithm.

在这种情况下， constructor 应该返回 interface 而不是 type 。比如在 hash 库中， crc32.NewIEEE 和 adler32.New 都返回 interface 类型的 hash.Hash32 。所以 在 Go 程序中，将 CRC-32 算法替换为 Adler-32 算法时，只需要改变一行 constructor 调用的代码，其他代码不需要任何变动。


A similar approach allows the streaming cipher algorithms in the various  `crypto`packages to be separated from the block ciphers they chain together. The  `Block`interface in the  `crypto/cipher`  package specifies the behavior of a block cipher, which provides encryption of a single block of data. Then, by analogy with the  `bufio`  package, cipher packages that implement this interface can be used to construct streaming ciphers, represented by the  `Stream`  interface, without knowing the details of the block encryption.

类似方法，可以让 streaming cipher 算法与不同 crypto package 中的 block cipher 自由组合。 在 crypt/cipher 中的 Block interface 规定了 block cipher 的行为，即，专门用于块数据(大小已知且固定的数据)的加密功能。然后，与 bufio package 类似，实现了 Block interface 的 cipher package 能很容易与 Stream interface 结合实现 streaming cipher ，而且不必了解具体 block encryption 的细节。

The  `crypto/cipher`  interfaces look like this:

crypto/cipher interfaces 定义如下 ：

```go
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}
```

Here's the definition of the counter mode (CTR) stream, which turns a block cipher into a streaming cipher; notice that the block cipher's details are abstracted away:

下面是 counter mode (CTR) stream, 将 block cipher 转换成 streaming cipher，注意 block cipher 是抽象的 interface：

```go
// NewCTR returns a Stream that encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block's block size.
func NewCTR(block Block, iv []byte) Stream
```

`NewCTR`  applies not just to one specific encryption algorithm and data source but to any implementation of the  `Block`interface and any  `Stream`. Because they return interface values, replacing CTR encryption with other encryption modes is a localized change. The constructor calls must be edited, but because the surrounding code must treat the result only as a  `Stream`, it won't notice the difference.

NewCTR 只仅能用于某一个特定的加密算法和数据，也能用于所有实现  Block interface 的加密算法，来生成相关的 Stream 实例。因为这个函数返回的是 interface value，替换 CTR 为其他加密模式的过程，只需要少量改动。 调用 NewCTR() 这个构造函数的代码肯定要修改，但其周围的代码不用动，因为它们用到的只是 Stream interface ，完全不关心具体加密算法的区别。



### Interfaces and methods 接口和方法

Since almost anything can have methods attached, almost anything can satisfy an interface. One illustrative example is in the  `http`  package, which defines the  `Handler`interface. Any object that implements  `Handler`  can serve HTTP requests.

几乎所有类型都能添加 method ，所以它们同样能满足 interface 的定义要求。
http package 中的 Handler interface 就是一个好例子。
所有实现了 Handler 的接口都能处理 HTTP 请求。

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`ResponseWriter`  is itself an interface that provides access to the methods needed to return the response to the client. Those methods include the standard  `Write`method, so an  `http.ResponseWriter`  can be used wherever an  `io.Writer`  can be used.`Request`  is a struct containing a parsed representation of the request from the client.

`ResponseWriter` interface 用于向 client 返回响应消息。
其中定义了标准的 Write method ，因此 http.ResponseWriter 能当作 io.Writer 使用。
`Request` struct 包含解析后的 client 请求数据。

For brevity, let's ignore POSTs and assume HTTP requests are always GETs; that simplification does not affect the way the handlers are set up. Here's a trivial but complete implementation of a handler to count the number of times the page is visited.

下面的代码片段可用于统计网页访问量，为简单起见，这里忽略 POST 请求，假设只有 Get 类型的 HTTP 请求。

```go
// Simple counter server.
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

(Keeping with our theme, note how  `Fprintf`can print to an  `http.ResponseWriter`.) For reference, here's how to attach such a server to a node on the URL tree.

(注意看， `Fprintf` 可以把响应信息输出到 http.ResponseWriter 中。) 统计特定 url 的访问频率，可以参考下面的代码。

```go
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

But why make  `Counter`  a struct? An integer is all that's needed. (The receiver needs to be a pointer so the increment is visible to the caller.)

这个小功能只需要一个 integer 变量，为什么要把 Counter 定义成 struct 呢？ ( receiver 必须是 pointer 时，才能被 caller 修改。)

```go
// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```

What if your program has some internal state that needs to be notified that a page has been visited? Tie a channel to the web page.

如果你需要在网页访问量变化时，修改程序的内部状态，那就要用 channel 了。

```go
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
```

Finally, let's say we wanted to present on  `/args`  the arguments used when invoking the server binary. It's easy to write a function to print the arguments.

如果我们要在请求 `/args` url 时，打印出服务端二进制进程的启动参数，怎么实现呢？只要写个输出函数就行了。

```go
func ArgServer() {
    fmt.Println(os.Args)
}
```

How do we turn that into an HTTP server? We could make  `ArgServer`  a method of some type whose value we ignore, but there's a cleaner way. Since we can define a method for any type except pointers and interfaces, we can write a method for a function. The  `http`  package contains this code:

但怎么把这个函数用到 HTTP server 中呢？可以给 ArgServer （按 Handler interface 的定义）加两个参数，但这个方法不够潇洒。
除了 pointers 和 interfaces 外，任何类似都能定义 method ，甚至能给 function 定义 method 。
http package 中就有这样的代码：

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```

`HandlerFunc`  is a type with a method,  `ServeHTTP`, so values of that type can serve HTTP requests. Look at the implementation of the method: the receiver is a function,  `f`, and the method calls  `f`. That may seem odd but it's not that different from, say, the receiver being a channel and the method sending on the channel.

HandlerFunc 是一个 function ，但它还带有一个 ServerHTTP method ，所以这种类型的 value 可以用来处理 HTTP 请求。仔细观察 ServerHTTP method 的实现， receiver 是一个 function 类型的 f ，并且在 method 中调用了 function f 。

To make  `ArgServer`  into an HTTP server, we first modify it to have the right signature.

把 ArgServer 加上正确的签名（按 Handler interface 的定义，加两个参数），就能用到 HTTP server 中。

```go
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
```

`ArgServer`  now has same signature as  `HandlerFunc`, so it can be converted to that type to access its methods, just as we converted  `Sequence`  to  `IntSlice`  to access  `IntSlice.Sort`. The code to set it up is concise:

现在的 ArgServer function 与 HandlerFunc type 有相同的定义 ，所以也能转换成 HandlerFunc 调用自己的 method 了。就像之前把 Sequence type 转换成 IntSlice type 来调用 IntSlice.Sort 一样。 代码十分简洁的：

```go
http.Handle("/args", http.HandlerFunc(ArgServer))
```

When someone visits the page  `/args`, the handler installed at that page has value  `ArgServer`  and type  `HandlerFunc`. The HTTP server will invoke the method  `ServeHTTP`  of that type, with  `ArgServer`  as the receiver, which will in turn call  `ArgServer`  (via the invocation  `f(w, req)`  inside  `HandlerFunc.ServeHTTP`). The arguments will then be displayed.

有人访问 `/args` 页面时，就会调用到 HandlerFunc 类型的 ArgServer valur 的 ServerHTTP 方法。 ArgServer 作为 receiver ，通过 HandlerFunc.ServeHTTP 方法中的 f(w, req) 被调用。此时，服务端的启动参数会作为响应显示在浏览器中。

In this section we have made an HTTP server from a struct, an integer, a channel, and a function, all because interfaces are just sets of methods, which can be defined for (almost) any type.

这一节中，我们分别使用 struct, integer, channe, function 实现了 HTTP server 。interface 能在任意 type 中定义使用。 interface 是 methods 的集合。



## The blank identifier 空白标识符

We've mentioned the blank identifier a couple of times now, in the context of  [`for``range`  loops](https://golang.org/effective_go.html#for)  and  [maps](https://golang.org/effective_go.html#maps). The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly. It's a bit like writing to the Unix  `/dev/null`  file: it represents a write-only value to be used as a place-holder where a variable is needed but the actual value is irrelevant. It has uses beyond those we've seen already.

之前介绍 for-range 和 map 时我们提到过 blank identifier。
blank identifier 可以被赋予和声明为任意类型的值，不过这些值会被丢弃。
这些点像 Unix 中的 /dev/null 文件：它是一个只能写入的占位符，表示这里需要一个变量，但变量的值没有实际用途。
之前已经见过它的用法了。

### The blank identifier in multiple assignment 空白标识符和多重赋值

The use of a blank identifier in a  `for`  `range`loop is a special case of a general situation: multiple assignment.

在 for-range 循环中偶尔会用到 blank identifier ，更多情况是 multiple assignment 时使用 blank identifier 。

If an assignment requires multiple values on the left side, but one of the values will not be used by the program, a blank identifier on the left-hand-side of the assignment avoids the need to create a dummy variable and makes it clear that the value is to be discarded. For instance, when calling a function that returns a value and an error, but only the error is important, use the blank identifier to discard the irrelevant value.

如果赋值语句左边需要多个返回值，但其中一个值不需要使用，就可以用 blank identifier 占位，从而避免创建一个无用的变量，并且也更能清晰的表达此变量要被丢弃的含义。比如，调用 functon 时返回一个 value 和 一个 error，但我们只关心是否出错，这里就能用 blank identifier 丢弃 value 。

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
	fmt.Printf("%s does not exist\n", path)
}
```

Occasionally you'll see code that discards the error value in order to ignore the error; this is terrible practice. Always check error returns; they're provided for a reason.

有时也会用 blank identifier 忽略错误信息；但这是一个可怕的习惯。
应该检查所有的 error 信息；因为 error 信息中会包含出错的详细原因。

```go
// Bad! This code will crash if path does not exist.
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
```

### Unused imports and variables 不用的变量和包 

It is an error to import a package or to declare a variable without using it. Unused imports bloat the program and slow compilation, while a variable that is initialized but not used is at least a wasted computation and perhaps indicative of a larger bug. When a program is under active development, however, unused imports and variables often arise and it can be annoying to delete them just to have the compilation proceed, only to have them be needed again later. The blank identifier provides a workaround.

声名变量或导入包时，如果不使用，会出现编译错误。
无用的 import 使代码臃肿并拖慢编译时间，初始化一个不用的变量，会浪费计算资源，甚至有时也预示着一个严重的 bug 。
但开发过程中，为了正常编译代码，必须频繁删除无用的 package / variable 是很恼火的事。
这种情况就要用到 blank identifier 了。

> 译：实践过程这个特性确实帮助我发现过多次 bug 。尤其是合并代码 或 使用了复制粘贴的代码时。


This half-written program has two unused imports (`fmt`  and  `io`) and an unused variable (`fd`), so it will not compile, but it would be nice to see if the code so far is correct.

下面这个写了一半的程序有两个未使用的 package (fmt 和 io) 和一个未使用的 variable (fd)，所以肯定是编译不通过的，但除此以外，代码都是正常的。

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
}
```

To silence complaints about the unused imports, use a blank identifier to refer to a symbol from the imported package. Similarly, assigning the unused variable  `fd`to the blank identifier will silence the unused variable error. This version of the program does compile.

用 `_`(blank identifier) 引用 package 中一个函数(symbol)，或者把变量赋值给 blank identifier ，就能忽略编译错误。下面的代码就是能正常编译的版本。

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

By convention, the global declarations to silence import errors should come right after the imports and be commented, both to make them easy to find and as a reminder to clean things up later.

按照惯例，这种用来忽略编译错误的全局声名必须在后面加上注释，一来提醒自己及时清理，二来方便找到这种代码。



### Import for side effect 导入包过程的副作用

An unused import like  `fmt`  or  `io`  in the previous example should eventually be used or removed: blank assignments identify code as a work in progress. But sometimes it is useful to import a package only for its side effects, without any explicit use. For example, during its  `init`function, the  [`net/http/pprof`](https://golang.org/pkg/net/http/pprof/)  package registers HTTP handlers that provide debugging information. It has an exported API, but most clients need only the handler registration and access the data through a web page. To import the package only for its side effects, rename the package to the blank identifier:

参考前面的示例可知， import `fmt` 或 `io` 等 package 后，必须使用，否则编译无法通过。除非使用 blank identify(`_`) 符号忽略。
有时候 import package 时，只是为了利用 import 过程产生的副作用(side effective)，并不会真正使用它。
比如[`net/http/pprof`](https://golang.org/pkg/net/http/pprof/) package 中， init() 函数注册了 HTTP 的 handlers ，用于显示 debuging information 。
虽然 pprof 也有导出一些 API，但多数情况只是访问一个 web page 显示一些调试信息，不经常用到这些 API 。
把 import 时的包名改成 blank identify(`_`) 就是为了达到这种副作用：自动调用 package 中 init() 函数的，但又不使用 package 中导出的 API。

```go
import _ "net/http/pprof"
```

This form of import makes clear that the package is being imported for its side effects, because there is no other possible use of the package: in this file, it doesn't have a name. (If it did, and we didn't use that name, the compiler would reject the program.)

这种写法能明确这行代码的目的：
不给 import 的 package 起名，因为我们只想利用 import package 过程的副作用，不会在当前代码中使用这个 package ，所以它不需要有名字。
（如果不用 `_` 忽略这个 package ，又不在代码中使用这个 package ，编译会失败）



### Interface checks 接口检查

As we saw in the discussion of  [interfaces](https://golang.org/effective_go.html#interfaces_and_types)above, a type need not declare explicitly that it implements an interface. Instead, a type implements the interface just by implementing the interface's methods. In practice, most interface conversions are static and therefore checked at compile time. For example, passing an  `*os.File`  to a function expecting an  `io.Reader`  will not compile unless  `*os.File`  implements the  `io.Reader`  interface.

之前讨论 [interfaces](https://golang.org/effective_go.html#interfaces_and_types) 时说过，定义 type 时不需要显式声名实现了哪些 interface ，只要实现了 interface 中定义的 method （方法），也就实现了 interface 。
实践中，多数情况转换 interface 的过程属于静态类型转换，所以能在编译期进行检查。
比如在函数调用时，如果 `*os.File` 没有实现 `io.Reader` 定义的方法，还把  `*os.File` 当成 `io.Reader` 类型的 interface 使用肯定会编译失败。


Some interface checks do happen at run-time, though. One instance is in the  [encoding/json](https://golang.org/pkg/encoding/json/)  package, which defines a  [Marshaler](https://golang.org/pkg/encoding/json/#Marshaler)  interface. When the JSON encoder receives a value that implements that interface, the encoder invokes the value's marshaling method to convert it to JSON instead of doing the standard conversion. The encoder checks this property at run time with a  [type assertion](https://golang.org/effective_go.html#interface_conversions)like:

有些 interface check 会发生在 run-time (运行时)。
比如 `encoding/json` package 中的 Marshaler interface 。
当 Json encoder （编码器）收到一个实现了此 interface 的 value 时， encoder 就会调用此 value 的 marshaling method 把它转换成 JSON ，此时，标准的（默认的） JSON 转换方法会被忽略。
这个过程中， encoder 就会用 type assertion （类型断言） 在 run-time (运行时）检查。

```go
m, ok := val.(json.Marshaler)
```

If it's necessary only to ask whether a type implements an interface, without actually using the interface itself, perhaps as part of an error check, use the blank identifier to ignore the type-asserted value:

如果只需要检查某种 type 是否实现了 interface ，并不需要实际使用这个 interface ，可以用 blank identifier 忽略 type-asserted 的返回值。

```go
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```


One place this situation arises is when it is necessary to guarantee within the package implementing the type that it actually satisfies the interface. If a type—for example,  [json.RawMessage](https://golang.org/pkg/encoding/json/#RawMessage)—needs a custom JSON representation, it should implement`json.Marshaler`, but there are no static conversions that would cause the compiler to verify this automatically. If the type inadvertently fails to satisfy the interface, the JSON encoder will still work, but will not use the custom implementation. To guarantee that the implementation is correct, a global declaration using the blank identifier can be used in the package:

当需要确定某个 package 中的 type 必须实现了指定的 interface 时，就会出现下述情况。
比如， json.RawMessage 需要自定义输出的 JSON 格式，所以它应该实现 json.Marshaler ,但因为没有静态类型转换的代码，所以编译器不会自动 verify （检查）。
如果因为某些意外情况，导致 json.RawMessage type 不满足 Marshaler interface 的要求，则 JSON encoder 还会继续正常执行下行，只是没有调用到 RawMessage 中定制的代码。
为防止类似情况发生，保证正确实现 interface 的要求，可以 package 中增加如下的一个全局声名。

```go
var _ json.Marshaler = (*RawMessage)(nil)
```


In this declaration, the assignment involving a conversion of a  `*RawMessage`  to a  `Marshaler`  requires that  `*RawMessage`implements  `Marshaler`, and that property will be checked at compile time. Should the  `json.Marshaler`  interface change, this package will no longer compile and we will be on notice that it needs to be updated.

这行变量声名代码使用了类型转换，所以要求 `*RawMessage` 必须实现 json.Marshaler interface 的定义，而且会在编译过程检查。如果 Marshaler 定义发生变化时， RawMessage 也必须同时调整，否则编译时就报错。


The appearance of the blank identifier in this construct indicates that the declaration exists only for the type checking, not to create a variable. Don't do this for every type that satisfies an interface, though. By convention, such declarations are only used when there are no static conversions already present in the code, which is a rare event.

这里使用了 blank identifier 是因为，这行代码的目的只是进行类型检查，不需要使用变量。
不建议给所有类型都增加上述代码来检查是否符合 interface 的实现。
只有在代码中没有 static conversion （静态类型转换）时，才需要增加这种声名，但实际上这种情况很少出现。



## Embedding 内嵌

Go does not provide the typical, type-driven notion of subclassing, but it does have the ability to “borrow” pieces of an implementation by  _embedding_  types within a struct or interface.

Go 没有提供 subclassing （子类） 的概念实现“继承”的功能，但它能在 struct 和 interface 中 embedding types （嵌入 type），从而继承这些 type 的功能。


Interface embedding is very simple. We've mentioned the  `io.Reader`  and  `io.Writer`interfaces before; here are their definitions.

interface embedding 很简单。我们之前提到过 io.Reader 和 io.Writer ，定义如下：


```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```


The  `io`  package also exports several other interfaces that specify objects that can implement several such methods. For instance, there is  `io.ReadWriter`, an interface containing both  `Read`  and  `Write`. We could specify  `io.ReadWriter`  by listing the two methods explicitly, but it's easier and more evocative to embed the two interfaces to form the new one, like this:

在 io package 中还导出了其他几个同样实现了这些 method 的 interface ，比如 `io.ReadWriter` ，这是一个同时包含 Read 和 Write 功能的 interface 。
在定义 `io.ReadWriter` interface 时，可以再把 Read Write method 的定义依次写出来。
但更简单而且容易理解的方法是，把 Reader 和 Writer 这两个 interface embed （内嵌）到 `io.ReadWriter` 中。
像下面这样：

```go
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

This says just what it looks like: A  `ReadWriter`  can do what a  `Reader`  does  _and_what a  `Writer`  does; it is a union of the embedded interfaces (which must be disjoint sets of methods). Only interfaces can be embedded within interfaces.

这么做，显而易见地表达了 ReadWriter 的作用，即同时拥有 Reader 和 Writer 的能力。
ReadWriter 是两个 embedded interface 的并集（同时包含两者定义的method）。
注意，只有 interface 能 embedded 到 interface 中。


The same basic idea applies to structs, but with more far-reaching implications. The  `bufio`  package has two struct types,`bufio.Reader`  and  `bufio.Writer`, each of which of course implements the analogous interfaces from package  `io`. And  `bufio`  also implements a buffered reader/writer, which it does by combining a reader and a writer into one struct using embedding: it lists the types within the struct but does not give them field names.

相同的逻辑也能应用到 struct 上，但有更多细节需要注意。
bufio package 有 `bufio.Reader` 和 `bufio.Writer` 这样两个 struct type，而且都实现了与 io package 中类似的接口。
bufio 使用 embedding struct 的方法实现了带缓冲的(buffered) reader/writer ：即在 struct 中列出 type name ，但并不赋予 field name。

```go
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```

The embedded elements are pointers to structs and of course must be initialized to point to valid structs before they can be used. The  `ReadWriter`  struct could be written as

内嵌在 ReadWriter 中的元素是 struct 类型的指针，所以使用 ReadWriter 类型的变量前也必须将指针指向可用的 struct 变量。


ReadWriter 也能像下面这样定义：

```go
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
```

but then to promote the methods of the fields and to satisfy the  `io`  interfaces, we would also need to provide forwarding methods, like this:

但是为了让 ReadWriter 同时满足 `io` interface 中的定义，我们要定义一个 method 转发调用请求到真正的 method 中，像下面这样：

```go
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```

By embedding the structs directly, we avoid this bookkeeping. The methods of embedded types come along for free, which means that  `bufio.ReadWriter`  not only has the methods of  `bufio.Reader`  and  `bufio.Writer`, it also satisfies all three interfaces:  `io.Reader`,  `io.Writer`, and`io.ReadWriter`.

如果直接 embedding struct ，我们就不用罗列每个 method 并实现转发请求的代码了。
现在 bufio.ReadWriter 不仅拥有 bufio.Reader 和 bufio.Writer 的方法，它同时满足 io.Reader io.Writer 和 io.ReadWriter 三个 interface 的定义。

> 译： outer type 会自动“继承” inner type ， ReadWriter 是 outer type ， Reader 和 Writer 是 inner type。


There's an important way in which embedding differs from subclassing. When we embed a type, the methods of that type become methods of the outer type, but when they are invoked the receiver of the method is the inner type, not the outer one. In our example, when the  `Read`  method of a  `bufio.ReadWriter`  is invoked, it has exactly the same effect as the forwarding method written out above; the receiver is the  `reader`  field of the  `ReadWriter`, not the`ReadWriter`  itself.

实现上， embedding 和 subclassing 还是有很大的区别的。
outer type 虽然“继承”了 inner type 的应运，但 invoked(调用) method 的过程， receiver 还是 inner type，而不是 outer type。
比如，上面两种实现方案中，我们调用  bufio.ReadWriter 的 Read method 的实际效果是十分相似的，两者的 receiver field 都是 bufio.Reader 类型，并不是 ReadWriter 。

> TODO　embed struct pointer 与 embed struct 有什么区别吗？


Embedding can also be a simple convenience. This example shows an embedded field alongside a regular, named field.

善用 embedding 能给我们提供很多便利。下面的代码中将 embedded field 与 named field 组合使用。

```go
type Job struct {
    Command string
    *log.Logger
}
```

The  `Job`  type now has the  `Print`,  `Printf`,  `Println`  and other methods of  `*log.Logger`. We could have given the  `Logger`  a field name, of course, but it's not necessary to do so. And now, once initialized, we can log to the  `Job`:

Job type 现在“继承”了 `*log.Logger`的 Log Logf 和 其他所有的方法。注意我们没有给 Logger field 命名。现在，初始化Job类型的变量后，就能用来记录日志了(log)。

```go
job.Println("starting now...")
```

The  `Logger`  is a regular field of the  `Job`struct, so we can initialize it in the usual way inside the constructor for  `Job`, like this,

Logger 也是 Job struct 中 regular （普通的） field ，所以我们像下面这样定义一个 constructor （构造函数）来初始化 Job 类型的变量：

> 译： `*log.Logger` 与 `Command` 其实都是 regular field

```go
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
```

or with a composite literal,

或者也可以直接使用 composite literal （复合字面量），

```go
job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
```


If we need to refer to an embedded field directly, the type name of the field, ignoring the package qualifier, serves as a field name, as it did in the  `Read`  method of our  `ReadWriter`  struct. Here, if we needed to access the  `*log.Logger`  of a  `Job`  variable  `job`, we would write  `job.Logger`, which would be useful if we wanted to refine the methods of  `Logger`.

我们可以用不带 package qualifier （包名限定符）的 type name 直接访问 embedded field ，跟访问其他普通的 field 和 method 一样简单。
比如，这里我们可以这样重定义 `Logf` method 的行为：
用 `job.Logger` 访问 job 变量的 `*log.Logger` field ，然后转发 Job.Logf mehtod 请求到 Logger.Logf 。

```go
func (job *Job) Logf(format string, args ...interface{}) {
    job.Logger.Logf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
// go1.13 的代码是下面这样，感觉是弄错了，无法与正文对应
// func (job *Job) Printf(format string, args ...interface{}) {
//     job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
// }
```

Embedding types introduces the problem of name conflicts but the rules to resolve them are simple. First, a field or method  `X`hides any other item  `X`  in a more deeply nested part of the type. If  `log.Logger`contained a field or method called  `Command`, the  `Command`  field of  `Job`  would dominate it.

embedding type 时会出现 name conflict (命名冲突)的问题，解决冲突的规则其实也很简单。

- 首先，优先使用 outer type 的 field/method 覆盖 inner type 中的重名 field/method。

比如，如果 log.Logger 也存在一个 Command field ， Job.Command 会 dominate （覆盖）它。



Second, if the same name appears at the same nesting level, it is usually an error; it would be erroneous to embed  `log.Logger`  if the  `Job`  struct contained another field or method called  `Logger`. However, if the duplicate name is never mentioned in the program outside the type definition, it is OK. This qualification provides some protection against changes made to types embedded from outside; there is no problem if a field is added that conflicts with another field in another subtype if neither field is ever used.

- 其次，如果冲突发生在相同的 nesting level （嵌套层级），通常会报错。

比如， Job struct 中包含一个名为 Logger 的 field 或 method 时，再 embed 一个 log.Logger 就会出错。

- 可是，冲突的 field 和 method 在程序中除 type definition （类型定义）之外，再没有被 mentioned (引用过)，也不会有问题。

这样的限制能在 type 定义改变时，提供一种保护机制。只要不显式使用冲突的 field ，就不会有问题。

> 译： name conflict 时，不使用冲突的 field ，虽然编译时不报错，但实际运行时，还会让自己踩坑的。比如，用于 json 解析用的 struct 中有重名字段，在 Marshar/Unmarshar 时就会有一些潜在问题。
> TODO 提供代码示例 name conflict 时踩的坑

## Concurrency 并发

### Share by communicating 使用通信来共享数据

Concurrent programming is a large topic and there is space only for some Go-specific highlights here.

并发编程是一个很大的话题，限于篇幅，这是只讨论 Go 中特有的东西。

Concurrent programming in many environments is made difficult by the subtleties required to implement correct access to shared variables. Go encourages a different approach in which shared values are passed around on channels and, in fact, never actively shared by separate threads of execution. Only one goroutine has access to the value at any given time. Data races cannot occur, by design. To encourage this way of thinking we have reduced it to a slogan:

并发编程中，不同环境中要确保并发访问共享变量时不出错是很困难的。
Go 另辟蹊径，仅通过 channel （信道）传递值来共享数据，而不会在多个独立执行的线程中共享数据。
任何时间都只有一个 goroutine 访问 value （数据）。从设计层面避免 Data Race （数据竞态）。
为了鼓励这种思考方式，我们提出下面这句口号：

> Do not communicate by sharing memory; instead, share memory by communicating.

> 不要通过共享内存（数据）来通信，而是通过通信来共享数据（内存）。

This approach can be taken too far. Reference counts may be best done by putting a mutex around an integer variable, for instance. But as a high-level approach, using channels to control access makes it easier to write clear, correct programs.

这种方法意义深远。
例如，使用 mutex （互斥量）配合一个 integer （整型）变量很容易实现 Reference Count （引用计数）。
但使用 channel 这种 high level (高阶，通用)的方案来控制数据访问的过程能写出逻辑清晰，不易出错的程序。

One way to think about this model is to consider a typical single-threaded program running on one CPU. It has no need for synchronization primitives. Now run another such instance; it too needs no synchronization. Now let those two communicate; if the communication is the synchronizer, there's still no need for other synchronization. Unix pipelines, for example, fit this model perfectly. Although Go's approach to concurrency originates in Hoare's Communicating Sequential Processes (CSP), it can also be seen as a type-safe generalization of Unix pipes.

可以这样理解这个模型。一个运行在单核CPU环境上的单线程程序是不需要同步原语的。然后，启动相同的程序，（两者之间没有交互）此时仍然是不需要同步原语。如果两个程序之间需要互相通信，但通信的过程是同步的，那此时仍然不需要同步原语。Unix Pipeline（管道）完全符合这种模型。Go的并发编程方法源于 Communicating Sequential Processes（CSP）顺序通信处理（Hoare的论文），可以把它理解成一种类型安全的 Unix Pipe。



### Goroutines Go 程

They're called  _goroutines_  because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations. A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

称其为 _Goroutine_ 是因为现存术语 Threads（线程）、Coroutines（协程）、Process（进程）都不能精确传达 Goroutine 的内涵。
Goroutine 是一种很简单的模型：
它是一个与其他 Goroutine 共享地址空间且 concurrently （并发）运行的函数。
它十分轻量，消耗很少的栈空间。
因为 stack 初始时非常小，仅在需要时才在 heap 中分配或释放 storage 空间。
所以它们的使用代价很小。

> TODO 区分 堆、栈 的具体分别，及这里具体要讲的是什么情况。

Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.

Goroutine 在 OS thread 上是多路复用。
如果一个 Goroutine 因等待 IO 操作而 block （阻塞），那么 CPU 就会切换到其他 Goroutine 运行。
从设计层面就隐藏了创建和管理线程的开销。

> TODO 多个 goroutine 有可能运行在一个 CPU 上，那么一个 goroutine 有可能这一秒在 CPU1 上运行，下一秒在CPU2 上运行吗？ 


Prefix a function or method call with the  `go`keyword to run the call in a new goroutine. When the call completes, the goroutine exits, silently. (The effect is similar to the Unix shell's  `&`  notation for running a command in the background.)

在 function 或 method 调用前增加一个 `go` 关键字，就能启动一个 Goroutine 。
当调用（函数执行）完毕后， Goroutine 就安静的退出了。
这个效果就像 UnixShell 中的 `&` 符号一样，让命令进程在后台运行。

```go
go list.Sort()  // run list.Sort concurrently; don't wait for it.
```

A function literal can be handy in a goroutine invocation.

也可以在调用 goroutine 时使用 function literal （函数字面量）。

```go
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

In Go, function literals are closures: the implementation makes sure the variables referred to by the function survive as long as they are active.

These examples aren't too practical because the functions have no way of signaling completion. For that, we need channels.

在 Go 中，function literal （函数字面量）是 closure （闭包）：这种实现保证在 function 中引用到的 variable 与 function 的生命周期是一样长的。

上面的示例并不实用，因为无法得知 function 是否执行完毕，为此，我们需要 channel 。



### Channels 信道

Like maps, channels are allocated with  `make`, and the resulting value acts as a reference to an underlying data structure. If an optional integer parameter is provided, it sets the buffer size for the channel. The default is zero, for an unbuffered or synchronous channel.

像 maps 一样，channel 也使用 make 分配内存空间，并且 channel 中收发的 value 也是对 underlying data 底层数据结构的引用。
如果调用 make 时，传递了 integer 参数，可用于设置 buffer size （channel的缓存大小）。
默认值是 0 ，即无缓冲的同步 channel 。

```go
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

Unbuffered channels combine communication—the exchange of a value—with synchronization—guaranteeing that two calculations (goroutines) are in a known state.

无缓冲 channel 在通信过程同步交换数据。
它能确保两个 goroutine 的执行过程处于确定状态。

There are lots of nice idioms using channels. Here's one to get us started. In the previous section we launched a sort in the background. A channel can allow the launching goroutine to wait for the sort to complete.

channel 有很多惯用的用法，我们现在开始介绍。
上一节中我们在后台启动了一个排序算法。
使用一个 channel 就能等待 goroutine 的 sort 操作执行完毕。

```go
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

Receivers always block until there is data to receive. If the channel is unbuffered, the sender blocks until the receiver has received the value. If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.

接收过程会一直阻塞，直到收到数据为止。
如果 channel 无缓冲，那么发送方也会一直阻塞，直到有接收方收到数据。
如果 channel 有缓冲，则发送方仅仅会阻塞到数据被 copy 到 buffer 为止；如果 buffer 满了，那么发送方也会阻塞到有任一接收方从 缓冲区取走一个数据为止。


A buffered channel can be used like a semaphore, for instance to limit throughput. In this example, incoming requests are passed to  `handle`, which sends a value into the channel, processes the request, and then receives a value from the channel to ready the “semaphore” for the next consumer. The capacity of the channel buffer limits the number of simultaneous calls to  `process`.

带缓冲的 channel 可当作信号量使用。
比如限制吞吐量。
下面的示例中，所有请求被传递到 handle 中， 
handle 函数先发送一个 value 到 channel ，然后调用 process 处理请求；
最后再从 channel 接收一个 value ，以便腾出 channel 缓冲空间，来正常处理下一次请求。
channel 的缓冲大小就决定了同时调用 process 的数量。

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```

Once  `MaxOutstanding`  handlers are executing  `process`, any more will block trying to send into the filled channel buffer, until one of the existing handlers finishes and receives from the buffer.

一旦执行 process 的 handlers 函数超过 MaxOutStanding ，任何向 channel send 的操作都会阻塞（ block ），直到某个 handlers 执行完毕，从 channel recv 一个数据。


This design has a problem, though:  `Serve`creates a new goroutine for every incoming request, even though only  `MaxOutstanding`  of them can run at any moment. As a result, the program can consume unlimited resources if the requests come in too fast. We can address that deficiency by changing  `Serve`  to gate the creation of the goroutines. Here's an obvious solution, but beware it has a bug we'll fix subsequently:

这个设计有个缺陷，虽然 serve 为每个 request 都创建了一个 goroutine ， 但还是只能有最大 MaxOutStanding 个 goroutine 同时运行。结果请求过快时，程序会无限地消耗资源。
我们可以修改 serve 函数限制创建 goroutine 的数量来解决这个问题，但这样修复后，还是会有一个 bug 的：

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // Buggy; see explanation below.
            <-sem
        }()
    }
}
```

The bug is that in a Go  `for`  loop, the loop variable is reused for each iteration, so the  `req`  variable is shared across all goroutines. That's not what we want. We need to make sure that  `req`  is unique for each goroutine. Here's one way to do that, passing the value of  `req`  as an argument to the closure in the goroutine:

这个 bug 就是，在 Go 的 for 循环中， loop 变量（req）在每迭代时都会复用。
我们要保证每个 goroutine 的 req 都是独立唯一的。
可以将 req 的值作为 goroutine 所使用的 closure（闭包）函数的参数。

> 译：所以多个 goroutine 会使用相同的 req 变量

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
}
```

Compare this version with the previous to see the difference in how the closure is declared and run. Another solution is just to create a new variable with the same name, as in this example:

比较前后两个版本，并注意， closure 是如何定义并运行的。
还有一个办法，就是创建一个同名的新变量，示例如下：

```go
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // Create new instance of req for the goroutine.
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

It may seem odd to write

这么看起来确实有些古怪：

```go
req := req
```

but it's legal and idiomatic in Go to do this. You get a fresh version of the variable with the same name, deliberately shadowing the loop variable locally but unique to each goroutine.

但这是合法的，而且是 Go 中的惯用手段。
我们生成了一个同名新变量，并且暂时隐藏了 loop 变量（局部变量），所以每个 gorouting 使用的都是独立的不同变量。

Going back to the general problem of writing the server, another approach that manages resources well is to start a fixed number of  `handle`  goroutines all reading from the request channel. The number of goroutines limits the number of simultaneous calls to  `process`. This  `Serve`function also accepts a channel on which it will be told to exit; after launching the goroutines it blocks receiving from that channel.

回到编写 serve 的问题。
另外一个管理资源的方法是启动固定数量的 handle goroutine。
每个 goroutine 从 channel 读取 Request 进行处理， goroutine 的数量就是同时调用 process  的数量。
另外 serve 函数还增加了一个 channel 参数用于处理退出（exit）的情况。
goroutine 启动后就阻塞地从 channel 接收数据。

```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // Start handlers
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // Wait to be told to exit.
}
```



### Channels of channels 信道中的信道

One of the most important properties of Go is that a channel is a first-class value that can be allocated and passed around like any other. A common use of this property is to implement safe, parallel demultiplexing.

Go 中的一个重要特性是 channel 是一等公民，它可以像普通类型一样分配并传递。
这种特性非常适合实现安全、并行的多路分解。

> TODO 这里的 demultiplexing 具体指什么？有没有相关示例？


In the example in the previous section,  `handle`  was an idealized handler for a request but we didn't define the type it was handling. If that type includes a channel on which to reply, each client can provide its own path for the answer. Here's a schematic definition of type  `Request`.

上一节中， `handle` 所处理的请求类型并没有明确定义。
如果在 Request 中增加一个 channel 变量，那每个 client 都能有自己的专属通道来回复响应消息了。
下面是 Request 的定义。

> TODO idealized handler 表达的含义是什么？ 是想说明示例仅仅是凭空想象，连成功或失败这类错误码都不能返回，所以没有实际使用价值吗？

```go
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
```

The client provides a function and its arguments, as well as a channel inside the request object on which to receive the answer.

其中包含一个参数，一个函数以及一个用于接收响应的 channel 。

> 译：client 提供计算方法(func)，计算参数(arg)；serve 负责计算，并把结果送到 result channel 。

```go
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests <- request
// Wait for response.
fmt.Printf("answer: %d\n", <-request.resultChan)
```

On the server side, the handler function is the only thing that changes.

server 端 handler 函数只有一点变化。

```go
func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan <- req.f(req.args)
    }
}
```

There's clearly a lot more to do to make it realistic, but this code is a framework for a rate-limited, parallel, non-blocking RPC system, and there's not a mutex in sight.

这些代码包含了 限流，并行，非阻塞 RPC 系统的基本框架，显然，这里没有使用 mutex。
但是要使它具有实际使用价值，还有很多事情要做。



### Parallelization 并行

Another application of these ideas is to parallelize a calculation across multiple CPU cores. If the calculation can be broken into separate pieces that can execute independently, it can be parallelized, with a channel to signal when each piece completes.

以上 idea 的另一个应用是在多CPU核心上进行并行计算。
如果计算过程能分解为独立运行的几部分，那么就能使其并行化，并在每一部分任务完成时使用 channel 发信号。


Let's say we have an expensive operation to perform on a vector of items, and that the value of the operation on each item is independent, as in this idealized example.

假设我们要对 vector 的每一项数据执行一个极其消耗资源的操作，并且这个操作是独立的。像下面的示例一样：

```go
type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1    // signal that this piece is done
}
```

We launch the pieces independently in a loop, one per CPU. They can complete in any order but it doesn't matter; we just count the completion signals by draining the channel after launching all the goroutines.

我们给每个 CPU 分配一部分计算任务，它们之间以任意顺序运行者没有影响。
然后在启动 goroutine 后统计 channel 中接收到的信号数量，来标识所有计算过程都已经完成。

```go
const numCPU = 4 // number of CPU cores

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // Drain the channel.
    for i := 0; i < numCPU; i++ {
        <-c    // wait for one task to complete
    }
    // All done.
}
```

Rather than create a constant value for numCPU, we can ask the runtime what value is appropriate.
The function  [runtime.NumCPU](https://golang.org/pkg/runtime#NumCPU)  returns the number of hardware CPU cores in the machine, so we could write

这里用一个常量 numCPU 来表示使用的 CPU core （CPU核心）数量，但我们其实能在运行时动态获取 CPU core 数量。

[runtime.NumCPU](https://golang.org/pkg/runtime#NumCPU) 函数返回机器中 CPU core 数量。
所以可以这样写：

```go
var numCPU = runtime.NumCPU()
```

There is also a function  [runtime.GOMAXPROCS](https://golang.org/pkg/runtime#GOMAXPROCS), which reports (or sets) the user-specified number of cores that a Go program can have running simultaneously. It defaults to the value of  `runtime.NumCPU`  but can be overridden by setting the similarly named shell environment variable or by calling the function with a positive number. Calling it with zero just queries the value. Therefore if we want to honor the user's resource request, we should write

还有一个 `runtime.GOMAXPROCS()`函数能返回用户设置的允许 Go 程序能同时使用的 CPU core 数量。
它默认就是 runtime.NumCPU()，但是这个值能通过 Shell 环境变量修改。
另外，在调用 runtime.GOMAXPROCS() 时传一个 >0 的整数，也能修改这个配置。
因此，如果希望优先使用用户配置的 CPU core 数量，应该这样写：

> 译： 传0时，表示查询当前配置。

```go
var numCPU = runtime.GOMAXPROCS(0)
```

Be sure not to confuse the ideas of concurrency—structuring a program as independently executing components—and parallelism—executing calculations in parallel for efficiency on multiple CPUs. Although the concurrency features of Go can make some problems easy to structure as parallel computations, Go is a concurrent language, not a parallel one, and not all parallelization problems fit Go's model. For a discussion of the distinction, see the talk cited in  [this blog post](http://blog.golang.org/2013/01/concurrency-is-not-parallelism.html).

注意，不要混淆并发（concurrency）和并行（parallelism）。
并发：将程序构造成可独立执行的组件。
并行：同时在多个 CPU 上执行计算任务，提高效率。
利用 Go 的并发特性能很容易实现并行计算，但 Go 是并发的语言，不是并行的。
而且不是所有并行问题都适合使用 Go 解决。
并发与并行的详细区别，参考这篇[博客文章](https://blog.golang.org/2013/01/concurrency-is-not-parallelism.html)



### A leaky buffer 漏桶缓冲池

The tools of concurrent programming can even make non-concurrent ideas easier to express. Here's an example abstracted from an RPC package. The client goroutine loops receiving data from some source, perhaps a network. To avoid allocating and freeing buffers, it keeps a free list, and uses a buffered channel to represent it. If the channel is empty, a new buffer gets allocated. Once the message buffer is ready, it's sent to the server on  `serverChan`.

并发编程语言也能很容易表达非并发的想法。
这有一个摘自 RPC package 的例子。
client goroutine 循环从某处接收数据，比如网络。
为了避免频繁的 allocating （分配）和 freeing （释放） buffer ，这里使用了一个 free list （空闲链表），list 使用带缓冲的 channel 实现。
如果 channel 是空的，就 allocate 一个新的 buffer。
一旦 message buffer 准备完毕，就会通过 serverChan 发送到 server 。

```go
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // Grab a buffer if available; allocate if not.
        select {
        case b = <-freeList:
            // Got one; nothing more to do.
        default:
            // None free, so allocate a new one.
            b = new(Buffer)
        }
        load(b)              // Read next message from the net.
        serverChan <- b      // Send to server.
    }
}
```

The server loop receives each message from the client, processes it, and returns the buffer to the free list.

server 端循环接收 client 端的  message ，处理完毕后，把 buffer 回收到 free list 。

func server() {
    for {
        b := <-serverChan    // Wait for work.
        process(b)
        // Reuse buffer if there's room.
        select {
        case freeList <- b:
            // Buffer on free list; nothing more to do.
        default:
            // Free list full, just carry on.
        }
    }
}

The client attempts to retrieve a buffer from  `freeList`; if none is available, it allocates a fresh one. The server's send to  `freeList`  puts  `b`  back on the free list unless the list is full, in which case the buffer is dropped on the floor to be reclaimed by the garbage collector. (The  `default`  clauses in the  `select`  statements execute when no other case is ready, meaning that the  `selects`  never block.) This implementation builds a leaky bucket free list in just a few lines, relying on the buffered channel and the garbage collector for bookkeeping.

client 端首先尝试从 freeList 中获取 buffer ，如果没有获取到，就 allocate 一个新的 buffer 。
server 端会首先尝试将 b 放到 freeList 中，如果 list 已经满了， buffer b 会被丢弃，然后等 garbage collector 自动回收。
(在 select 语句中使用的 default 关键字会在没有其他 case 可执行时，执行，这样 select 就永远不会 block（阻塞）了。)
这里利用带缓冲的 channel 和 garbage collector 进行资源管理，只用几行代码就实现了一个 leaky bucket free list 。

> 译：leaky bucket 漏斗、漏桶。这里想介绍的是漏桶算法，它主要用于流量控制。但这里的示例好像是在演示"内存空间控制"。
> TODO 搜索 leaky bucket 查询 wiki 等资料详细了解。极客时间与 Google Cloud 中也能搜索到相关资料。
> 限制接口调用频率(rate limit) 会用到这些算法：固定窗口，滑动窗口，漏桶，令牌桶。
> 另外， linux 中 traffic control 控制用户上网流量也会使用令牌桶 (token bucket) 。

## Errors 错误

Library routines must often return some sort of error indication to the caller. As mentioned earlier, Go's multivalue return makes it easy to return a detailed error description alongside the normal return value. It is good style to use this feature to provide detailed error information. For example, as we'll see,  `os.Open`  doesn't just return a  `nil`  pointer on failure, it also returns an error value that describes what went wrong.

library routines（库程序）经常会返回各种类型的 error 给调用者。
之前提到过，Go 的多值返回特性非常方便将 错误描述信息 与 正常返回值 区分开。
利用这种方式返回错误信息是个好习惯。
比如，`os.Open`执行失败时，不仅返回了一个 `nil` 指针，还返回了一个 error 值来描述详细的错误信息。

By convention, errors have type  `error`, a simple built-in interface.

按照惯例，Go语言内置了 `error` 类型的 interface 专门用于此。

```go
type error interface {
    Error() string
}
```

A library writer is free to implement this interface with a richer model under the covers, making it possible not only to see the error but also to provide some context. As mentioned, alongside the usual  `*os.File`return value,  `os.Open`  also returns an error value. If the file is opened successfully, the error will be  `nil`, but when there is a problem, it will hold an  `os.PathError`:

Library 可以利用这个 error interface 实现更丰富的功能，使之不仅提供错误描述信息，还包含更详细的上下文信息。
刚才说`os.Open`返回值中，除了`os.File`类型的变量外，还有一个 error 值。
文件打开正常时，返回的 error 是 nil ，但打开失败时，它返回的其实是`os.PathError`类型的 error 。

```go
// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // The associated file.
    Err error    // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

`PathError`'s  `Error`  generates a string like this:

PathError 的 Error() 方法返回的错误信息如下：

```txt
open /etc/passwx: no such file or directory
```

Such an error, which includes the problematic file name, the operation, and the operating system error it triggered, is useful even if printed far from the call that caused it; it is much more informative than the plain "no such file or directory".

这个错误信息中，不仅包含了 文件名称，执行的操作，还有它 operating systm error 信息。
这比仅仅一句`no such file or directory`信息可有用多了。
即使输出日志的代码与调用`os.Open()`失败的代码相差很远，这种错误信息也能帮我们定位问题。


When feasible, error strings should identify their origin, such as by having a prefix naming the operation or package that generated the error. For example, in package  `image`, the string representation for a decoding error due to an unknown format is "image: unknown format".

错误信息应该尽可能包含它的来源，比如增加上产生错误的 operation（操作）或 package （包名）前缀。
比如在 package `image`中，表示由于未知格式而产生的解码失败的错误信息格式是“image: unknow format"。

Callers that care about the precise error details can use a type switch or a type assertion to look for specific errors and extract details. For  `PathErrors`  this might include examining the internal  `Err`  field for recoverable failures.

如果调用者十分关心准确的错误信息，可以使用 type switch （类型选择）或 type assertion （类型断言）来识别具体的错误并抽取出详细信息。
比如获取到`PathErrors`的内部`Err` field（字段），就能对错误现场进行一些必要的处理。

```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
```

The second  `if`  statement here is another  [type assertion](https://golang.org/effective_go.html#interface_conversions). If it fails,  `ok`  will be false, and  `e`  will be  `nil`. If it succeeds,  `ok`  will be true, which means the error was of type  `*os.PathError`, and then so is  `e`, which we can examine for more information about the error.

第二个`if`语句是 type assertion 。
如果断言失败， `ok`的值是 false ，并且`e`的值是`nil`。
如果断言成功，`ok`的值是true，并且`e`的类型是`*os.PathError`，我们能从这个类型的错误中获取更多更详细的信息。



### Panic 惊恐

The usual way to report an error to a caller is to return an  `error`  as an extra return value. The canonical  `Read`  method is a well-known instance; it returns a byte count and an  `error`. But what if the error is unrecoverable? Sometimes the program simply cannot continue.

通常情况，向调用者返回一个 `error` 值来表示出现了错误。
`Read()`方法就是一个广为人知的例子，它返回一个 byte count 和一个 error 。
如果错误是不可恢复的呢，有时程序确实无法执行下去了，怎么办？

For this purpose, there is a built-in function  `panic`  that in effect creates a run-time error that will stop the program (but see the next section). The function takes a single argument of arbitrary type—often a string—to be printed as the program dies. It's also a way to indicate that something impossible has happened, such as exiting an infinite loop.

为此，我们提供了内置函数`panic`，它会产生一个运行时错误，并让程序停止运行（详细情况参考下一节）。
该函数接受一个任意类型的参数（通常是一个字符串）用于在程序终止时（译：作为提示信息）输出。
它还可用于表明发生了意料之外的事情，比如从无限循环中退出了。

```go
// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
    z := x/3   // Arbitrary initial value
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // A million iterations has not converged; something is wrong.
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
```

This is only an example but real library functions should avoid  `panic`. If the problem can be masked or worked around, it's always better to let things continue to run rather than taking down the whole program. One possible counterexample is during initialization: if the library truly cannot set itself up, it might be reasonable to panic, so to speak.

以上只是示例，真实的库函数应该尽量避免`panic`。
如果能屏蔽或解决问题，最好让程序正常运行，尽量减少程序停止工作的情况。
一个反例是，在初始化的过程，如果某个库确实无法让它正常工作，而且有正当的理由产生 panic ，那就让它挂掉吧。

```go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

### Recover 恢复

When  `panic`  is called, including implicitly for run-time errors such as indexing a slice out of bounds or failing a type assertion, it immediately stops execution of the current function and begins unwinding the stack of the goroutine, running any deferred functions along the way. If that unwinding reaches the top of the goroutine's stack, the program dies. However, it is possible to use the built-in function  `recover`  to regain control of the goroutine and resume normal execution.

当`panic`被调用时，会立即停止执行当前函数，并展开当前 goroutine 的调用栈，然后执行当前函数中的 deferred function。
如果 unwinding 到 goroutine top （顶端），the program die（程序就挂了）。
但我们可以使用内置的`recover`函数夺回 goroutine 的控制权，并让程序恢复正常运行。


A call to  `recover`  stops the unwinding and returns the argument passed to  `panic`. Because the only code that runs while unwinding is inside deferred functions,  `recover`  is only useful inside deferred functions.

调用`recover`将停止 unwinding （展开、回溯）并返回调用 panic 函数时传递的参数。
因为 unwinding 过程能执行的代码只有 defer 调用的函数，所以 `recover`必须放在 defer 调用的函数中才有效。


One application of  `recover`  is to shut down a failing goroutine inside a server without killing the other executing goroutines.

`recover`的用途之一就是终止异常的 goroutine ，并防止进程被 kill （杀死）。

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

In this example, if  `do(work)`  panics, the result will be logged and the goroutine will exit cleanly without disturbing the others. There's no need to do anything else in the deferred closure; calling  `recover`  handles the condition completely.

在示例中，如果`do(work)`发生 panic ，会立即记录这一情况到日志中，并干净利落的结束异常的 goroutine ，其他 goroutine 不会有影响。
要实现这种效果，只需要在 defer 运行的函数中调用一个`recover`即可。

Because  `recover`  always returns  `nil`  unless called directly from a deferred function, deferred code can call library routines that themselves use  `panic`  and  `recover`  without failing. As an example, the deferred function in  `safelyDo`  might call a logging function before calling  `recover`, and that logging code would run unaffected by the panicking state.

除非在 defer 的函数中调用`recover`，否则，在其他任何地方调用`recover`总会返回 nil 值。
defer 的代码中，如果调用的第三方 library 也用了`panic`和`recover`，是不会报错的。
示例中，`safelyDo()`函数中，如果在调用`recover`之前调用了 logging 函数，logging 过程的代码是不受 panicking 状态所影响的。

> TODO 验证在 defer func 中调用的函数 panic 会发生什么？


With our recovery pattern in place, the  `do`function (and anything it calls) can get out of any bad situation cleanly by calling  `panic`. We can use that idea to simplify error handling in complex software. Let's look at an idealized version of a  `regexp`package, which reports parsing errors by calling  `panic`  with a local error type. Here's the definition of  `Error`, an  `error`  method, and the  `Compile`  function.

合理的使用 recovery （恢复模式），能干净利落地让`do`函数从`panic`异常中恢复。
我们可以利用这个想法简化软件开发中的错误处理。
我们以`regexp` package 为例，这个包中，parsing（解析）失败时，会调用`panic`报告错误，panic 中传一个自定义的 error 类型。
下面是`Error`类型，`error` method，和`Compile` function 的定义。

```go
// Error is the type of a parse error; it satisfies the error interface.
type Error string
func (e Error) Error() string {
    return string(e)
}

// error is a method of *Regexp that reports parsing errors by
// panicking with an Error.
func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

// Compile returns a parsed representation of the regular expression.
func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    // doParse will panic if there is a parse error.
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // Clear return value.
            err = e.(Error) // Will re-panic if not a parse error.
        }
    }()
    return regexp.doParse(str), nil
}
```


If  `doParse`  panics, the recovery block will set the return value to  `nil`—deferred functions can modify named return values. It will then check, in the assignment to  `err`, that the problem was a parse error by asserting that it has the local type  `Error`. If it does not, the type assertion will fail, causing a run-time error that continues the stack unwinding as though nothing had interrupted it. This check means that if something unexpected happens, such as an index out of bounds, the code will fail even though we are using  `panic`  and  `recover`  to handle parse errors.

如果`doParse` panic ， recovery 部分的代码会设置返回参数 regexp 为 nil 值， defer 的函数中能修改有名称的返回参数。
然后，通过断言 `err` 是否为自定义的`Error`类型来识别 parse error （解析异常）。
如果不是`Error`类型， type assertiong （类型断言）失败，又会引发一个新的 run-time （运行时）错误，继而重新引发 stack unwinding （堆栈展开），就像从来没被打断过一样。
这块代码说明，即使使用了`panic`和`recover`来处理 parse error ，如果发生`索引越界`一类的异常错误，程序仍然会挂掉。


With error handling in place, the  `error`method (because it's a method bound to a type, it's fine, even natural, for it to have the same name as the builtin  `error`  type) makes it easy to report parse errors without worrying about unwinding the parse stack by hand:

合理利用`error` method 来处理错误，能简化报告 parse error 错误的过程，而且省去手动 unwinding the parse stack （展开调用栈）的麻烦。

```go
if pos == 0 {
    re.error("'*' illegal at start of expression")
}
```

Useful though this pattern is, it should be used only within a package.  `Parse`  turns its internal  `panic`  calls into  `error`  values; it does not expose  `panics`  to its client. That is a good rule to follow.

尽管这种方法很好用，也尽量不要在 package 外部使用。
`Parse`将内部的`panic`调用转换成一个变通的`error`错误变量； 这样就不会给客户端暴露出`panic`。
这是一个值得遵守的好规则。


By the way, this re-panic idiom changes the panic value if an actual error occurs. However, both the original and new failures will be presented in the crash report, so the root cause of the problem will still be visible. Thus this simple re-panic approach is usually sufficient—it's a crash after all—but if you want to display only the original value, you can write a little more code to filter unexpected problems and re-panic with the original error. That's left as an exercise for the reader.

顺便一提，这种 re-panic 的手法会改变实际产生 panic 的 value 。
但新旧两次错误信息都会在 crash report（崩溃报告）中显示，所以问题的 root cause （根源）仍然是可见的。
所以这种简单的 re-panic 的方法已经够用了——毕竟它只能崩溃一次——但如果你只想显示原始的错误信息，那就还要再写一点代码来过滤无用的信息，然后再使用 origin error 来 re-panic。
这个练习就留给读者了。



## A web server 一个 web 服务器

{% raw %}

Let's finish with a complete Go program, a web server. This one is actually a kind of web re-server. Google provides a service at  `chart.apis.google.com`  that does automatic formatting of data into charts and graphs. It's hard to use interactively, though, because you need to put the data into the URL as a query. The program here provides a nicer interface to one form of data: given a short piece of text, it calls on the chart server to produce a QR code, a matrix of boxes that encode the text. That image can be grabbed with your cell phone's camera and interpreted as, for instance, a URL, saving you typing the URL into the phone's tiny keyboard.

我们一起用 Go 实现一个完整的 WebServer 程序吧。
这实现上是一个 re-server （转发服务器）。
Google 在 chart.apis.google.com 提供了能将 formatting data 转换成 chart and graphs （图表）的接口服务。
但它并不易用，因为需要把 data 放到 URL 的参数中组成一个查询请求。
下面的程序用来简化成图表的过程：
输入一小段文本，它就能调用 chart server生成QRCode（二维码），一种编码的点格矩阵。
手机摄像头能识别二维码，并解码为字符串，比如URL，这就省去用手机键盘输入的麻烦了。


Here's the complete program. An explanation follows.

下面是完整的程序，随后会有详细的说明。

```go
package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`
```

The pieces up to  `main`  should be easy to follow. The one flag sets a default HTTP port for our server. The template variable  `templ`  is where the fun happens. It builds an HTML template that will be executed by the server to display the page; more about that in a moment.

`main`函数之前的代码比较好理解，使用`flag`设置程序监听的 HTTP port 。
`temp`是 template 模板变量，它用于创建 HTML 模板，生成 client 端显示的页面，一会儿会详细讨论。


The  `main`  function parses the flags and, using the mechanism we talked about above, binds the function  `QR`  to the root path for the server. Then  `http.ListenAndServe`  is called to start the server; it blocks while the server runs.

`main`函数调用`flag.Parse()`解析程序启动时的命令行参数。
同时将`QR`函数 bind 到 HTTP server 的 root path('/')，然后调用`http.ListenAndServe`启动 HTTP Server，服务器程序运行到此处时阻塞(block）。


`QR`  just receives the request, which contains form data, and executes the template on the data in the form value named  `s`.

`QR`函数接收 form 表单数据，取出其中名`s`值，并调用 `temp1.Execute`传递到HTML模板中。


The template package  `html/template`  is powerful; this program just touches on its capabilities. In essence, it rewrites a piece of HTML text on the fly by substituting elements derived from data items passed to  `templ.Execute`, in this case the form value. Within the template text (`templateStr`), double-brace-delimited pieces denote template actions. The piece from  `{{html "{{if .}}"}}`  to  `{{html "{{end}}"}}`  executes only if the value of the current data item, called  `.`  (dot), is non-empty. That is, when the string is empty, this piece of the template is suppressed.

`html/template`package十分强大，这里只是仅仅用了很简单的功能。
本质上，它通过传递给`temp1.Execute`参数动态生成HTML文本，这个例子中，它传递的参数是 form 表单的值。
模板文本(templateStr)中，双大括号之间的文本就是模板代码(denote template action)。
从`{{if .}}`到`{{end}}`之间的文本只在当前数据项非空时，才会执行，即这里的`.(dot)`的值非空时。
也就是说，当temp1.Execute的第二个参数字符串为空时，这部分文本会被忽略。

> 译："忽略"表示不会显示到 HTML 页面中。


The two snippets  `{{html "{{.}}"}}`  say to show the data presented to the template—the query string—on the web page. The HTML template package automatically provides appropriate escaping so the text is safe to display.

两处`{{.}}`表示将数据（即query string）直接显示在HTML页面上。`html/template`会自动对文本进行转义，保证显示的文件是安全的。

The rest of the template string is just the HTML to show when the page loads. If this is too quick an explanation, see the  [documentation](https://golang.org/pkg/html/template/)  for the template package for a more thorough discussion.

剩下的 template 字符串会在页面加载完毕时，直接显示出来。
如果这段解释不好理解，可以参考[文档](https://golang.org/pkg/html/template/)获得更详细的说明。

And there you have it: a useful web server in a few lines of code plus some data-driven HTML text. Go is powerful enough to make a lot happen in a few lines.

Go语言就是这么强大，我们只用了几行代码就实现了能动态生成HTML页面的Web服务器。



Build version go1.13



## 更新说明 2019.09.20
- 美化一些排版格式
  * 2019.06.01 的版本，没用 vim 编辑器，也没有 Note Plus 这个电子墨水屏，排版过程十分麻烦。
  * 甚至有一部分内容是先写纸上，再打字输入到电脑。
  * 导致自己积极性不高，翻译周期十分长，大概2018年初始就开始翻译了。
- 补充一些必要的内容
  * Golang 已经从 go1.11 升级到 go1.13 ,文档内容有一些些细节差异。
  * 还要把英文原文保留下来，方便以后中更新翻译错误时，不好找原文。
  * 自己查阅资料时，有原文也更方便检索。
- 顺便温故知新
  * 复习技术细节。
  * 提高英文阅读的理解能力，反复翻译同一个文字的过程，肯定也能加深复杂句型的理解能力吧。



## 更新说明 2019.06.01
- 英文术语尽量不翻译
  * 方便记忆理解。也方便上Google检索相关技术问题。
- 使用自己的语言描述，尽量避免直译
  * 直译的内容用Google翻译就能得到。网上也能找到很多类似的内容，对我来说，并不好理解。
- 系统学习一遍Golang相关的技术。
- ` raw ` 和 ` endraw ` 不是 markdown 语法
  * 这是为了在 jekyll 中显示 `{{}}` 这样的特殊符号使用。


{% endraw %}

[^EffectiveGoEn]: [Effective Go 官方文档](https://golang.org/doc/effective_go.html)
[^EffectiveGoCn1]: [Effective Go studygolang](https://studygolang.com/articles/3228)
[^EffectiveGoCn2]: [Effective Go docscn.studygolang](http://docscn.studygolang.com/doc/effective_go.html)



