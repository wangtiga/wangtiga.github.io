---
layout: post
title:  "[译] 在MySQL一定要用 utf8mb4 字符集，千万别用 utf8 "
date:   2019-08-13 20:05:00 +0800
tags: translate
---

* category
{:toc}

## In MySQL, never use “utf8”. Use “utf8mb4” Adam Hooper 14, 2016 [^MySQL_utf8_utf8mb4]

Today’s bug: I tried to store a UTF-8 string in a MariaDB “utf8”-encoded database, and Rails raised a bizarre error:
`Incorrect string value: ‘\xF0\x9F\x98\x83 <…’ for column ‘summary’ at row 1`

今天遇到一个bug:我在尝试向使用"utf8"字符串的MariaDB数据库中插入一个UTF-8字符串时， Rails 框架报以下错误
`Incorrect string value: ‘\xF0\x9F\x98\x83 <…’ for column ‘summary’ at row 1`

译：Rails 是 Ruby 编程语言的一个开发框架
译：MariaDB 是一个开源数据库，与MySQL同一个作者

This is a UTF-8 client and a UTF-8 server, in a UTF-8 database with a UTF-8 collation. The string, “😃 <…”, is valid UTF-8.

这是使用 UTF-8 编码实现的 client 和 server, 数据库也是 UTF-8 编码实现的，而且数据库的数据也使用 UTF-8 编码。而且要存储的字符串“😃 <…”也是合法的 UTF-8 编码。

But here’s the rub: MySQL’s “**utf8**”  _isn’t UTF-8_.

矛盾就在这里：MySQL的 “**utf8**” 并 _不是_ 真正的 UTF-8 字符集。

The “utf8” encoding only supports three bytes per character. The real UTF-8 encoding — which everybody uses, including you — needs up to four bytes per character.

MySQL 所用的"utf8"编码只支持每个字符(character)3个字节(byte)。但真正的UTF-8编码中每个字符最大可能会超过4个 byte 。

MySQL developers never fixed this bug. They released a workaround in  [2010](https://dev.mysql.com/doc/relnotes/mysql/5.5/en/news-5-5-3.html): a new character set called “**utf8mb4**”.

MySQL 的开发者[2010发布的版本](https://dev.mysql.com/doc/relnotes/mysql/5.5/en/news-5-5-3.html)，提供了一个新的 “**utf8mb4**” 字符集绕过了这个问题。所以这个bug一直也没修复。

Of course, they never advertised this (probably because the bug is so embarrassing). Now, guides across the Web suggest that users use “utf8”. All those guides are wrong.

但是，他们从来没有宣告并强调过这个问题（也许是因为太尴尬了吧？）。现在，用户在网上看到 guide 上建议使用 "utf8" ，但这个建议是错的。

In short:

-   MySQL’s “utf8mb4” means “UTF-8”.
-   MySQL’s “utf8” means “a proprietary character encoding”. This encoding can’t encode many Unicode characters.

简要说明如下：

-   MySQL 的 "utf8mb4" 表示真正的 "UTF-8"。
-   MySQL 的 "utf8" 表示 "一种私有的字符编码集"。很多 Unicode 字符不能在这个编码集中使用。


I’ll make a sweeping statement here:  _all_  MySQL and MariaDB users who are currently using “utf8” should  _actually_  use “utf8mb4”. Nobody should ever use “utf8”.

我在这里强烈建议： _所有_ MySQL 和 MariaDB 的用户都 _应该_ 用 "utf8mb4" 字符集表示 "utf8" 字符。没有任何人需要使用 "utf8"。


### What’s encoding? What’s UTF-8?

### 编码是什么意思？UTF-8又是什么意思？

Joel on Software wrote  [my favorite introduction](http://www.joelonsoftware.com/articles/Unicode.html). I’ll abridge it.

"Joel 谈软件"写过一篇文章 [my favorite introduction](http://www.joelonsoftware.com/articles/Unicode.html)。我简要概括一下。

Computers store text as ones and zeroes. The first letter in this paragraph was stored as “01000011” and your computer drew “C”. Your computer chose “C” in two steps:

计算机使用 0 和 1 存储文本信息。这段文字第一个字符是 C ，在计算机中存储为 "01000011" ，然后你的计算机绘制字符 C 的样子显示出来。计算机之所以选择字符 "C" 的样子绘制到屏幕，是经过以下几步骤得出的结果：

1.  Your computer read “01000011” and determined that it’s the number 67. That’s because 67 was  _encoded_  as “01000011”.
2.  Your computer looked up character number 67 in the  [Unicode](http://unicode.org/)  _character set_, and it found that 67 means “C”.

- 1.  计算机读取到二进制数据 "01000011" ，并认出它实际就是数值 67 。这是因为 67 被 _编码_ 为 "01000011" 。
- 2.  计算机在 [Unicode](http://unicode.org/) _字符集_ 中查询数值等于 67 的字符，它发现 67 表示字符 "C" 。 

The same thing happened on my end when I typed that “C”:

1.  My computer mapped “C” to 67 in the Unicode character set.
2.  My computer  _encoded_  67, sending “01000011” to this web server.

当我用键盘输入字符 "C" 的时候，也会发生类似的过程：

- 1.  计算机从 Unicode 字符集中将字符 "C" 映射为数值 67 。
- 2.  计算机将数值 67 编码，并发送 "01000011" 到 Web 服务器。

_Character sets_  are a solved problem. Almost every program on the Internet uses the Unicode character set, because there’s no incentive to use another.

_字符集_ 是已经确定的方案。目前 Internet 中几乎所有程序都使用 Unicode 字符集，因为找不到什么理由用另外一套字符串。

But  _encoding_  is more of a judgement call. Unicode has slots for over a million characters. (“C” and “💩” are two such characters.) The simplest encoding, UTF-32, makes each character take 32 bits. That’s simple, because computers have been treating groups of 32 bits as numbers for ages, and they’re really good at it. But it’s not useful: it’s a waste of space.

但 _编码方案_ 就存在很多种方法了。Unicode 定义了上万个字符。（"C" 和 “💩” 只是其中两个字符。）最简单的编码方案当属 UTF-32 ，它规定每个字符都占用 32 bit 。这种方案简单是因为，计算机系统多年来都是以 32 bit 一组数据进行处理的。但它也有缺点，太浪费空间了。
译：前几年 CPU 一直是 32 bit 的处理器，最近大多已经升级为 64 bit 处理器了。但处理 32 bit 一组的数据对计算机来说仍然很容易。 @2019

TODO: 翻译 judgement call


UTF-8 saves space. In UTF-8, common characters like “C” take 8 bits, while rare characters like “💩” take 32 bits. Other characters take 16 or 24 bits. A blog post like this one takes about four times less space in UTF-8 than it would in UTF-32. So it loads four times faster.



UTF-8 非常节省空间。在 UTF-8 中，像 "C" 这样的普通字符只占用 8 个 bit ，只有少量类似 “💩” 这样的字符才会占用 32 bit 。还有一些其他字符只占用 16 或 24 bit 。像这样一篇英文的博客文章，使用 UTF-32 编码占用空间 是 使用 UTF-8 编码占用空间 的四倍。所以加载时间也有可能慢四倍。

译：英文字符多数只占用 8 bit ，但中文字符多数占用 16 bit ，所以不是四位的关系。 

You may not realize it, but our computers agreed on UTF-8 behind the scenes. If they didn’t, then when I type “💩” you’ll see a mess of random data.

各位读者可以不了解这些细节，但我们的计算机必须了解并遵守 UTF-8 的规则。否则，当我打出 “💩” 字符时，你看到的会是一坨乱七八槽的数据。 

MySQL’s “utf8” character set doesn’t agree with other programs. When they say “💩”, it balks.

MySQL 的 utf8 字符集与标准 UTF-8 的规则有差异，所以当它看到字符 “💩” 时，就懵了。。

### A bit of MySQL history

### 关于 MySQL 的一点历史

Why did MySQL developers make “utf8” invalid? We can guess by looking at commit logs.

为什么 MySQL 开发者实现 "utf8" 编码时会出错呢？我们可以从代码的 commit logs 中做一些猜想。

MySQL supported UTF-8 since  [version 4.1](http://mysql.localhost.net.ar/doc/refman/4.1/en/news-4-1-0.html). That was 2003 — before today’s UTF-8 standard,  [RFC 3629](https://tools.ietf.org/html/rfc3629).

MySQL 从 [version 4.1](http://mysql.localhost.net.ar/doc/refman/4.1/en/news-4-1-0.html) 版本开始支持 UTF-8 编码。这是在如今的 UTF-8 标准 [RFC3629](https://tools.ietf.org/html/rfc3629) 制定之前的 2003 年。

The previous UTF-8 standard,  [RFC 2279](https://www.ietf.org/rfc/rfc2279.txt), supported up to six bytes per character. MySQL developers coded RFC 2279 in the  [the first pre-pre-release version of MySQL 4.1](https://github.com/mysql/mysql-server/commit/55e0a9cb01af4b01bc4e4395de9e4dd2a1b0cf23)  on March 28, 2002.

在这之前的 UTF-8 标准,  [RFC 2279](https://www.ietf.org/rfc/rfc2279.txt)，支持每个字符6个 bytes 。 MySQL 开发者在 2002 年 三月 28 号编码实现了 RFC 2279 ， 可以参考[the first pre-pre-release version of MySQL 4.1](https://github.com/mysql/mysql-server/commit/55e0a9cb01af4b01bc4e4395de9e4dd2a1b0cf23) 。


Then came a cryptic, one-byte tweak to MySQL’s source code in September: “UTF8 now works with up to  [3 byte sequences only](https://github.com/mysql/mysql-server/commit/43a506c0ced0e6ea101d3ab8b4b423ce3fa327d0).”

在九月份时，MySQL 代码突然出现了一处神秘改动：“UTF8 现在最多只占用 [3 byte sequences only](https://github.com/mysql/mysql-server/commit/43a506c0ced0e6ea101d3ab8b4b423ce3fa327d0).”

Who committed this? Why? I can’t tell. MySQL’s code repository seems to have lost old author names when it adopted Git. (MySQL used to use BitKeeper, like the Linux kernel.) There’s nothing on the mailing list around September 2003 that explains the change.

谁提交的这行代码？为什么要这样改？我也不知道。MySQL 的代码仓库转移到 Git 的过程丢失了 author 名称。（MySQL 以前与 Linux 内核一样使用 BitKeeper 管理代码。）在 2003 年的 mailing list 中也没有关于这次改动的任何解释。

But I can guess.

但我能大概猜测出来。

译： 也可能这里的原因 [^max_byte_utf8_1]  [^max_byte_utf8_2]

Back in 2002, MySQL gave users a  [speed boost](http://dev.mysql.com/doc/refman/5.7/en/static-format.html)  if users could guarantee that every row in a table had the same number of bytes. To do that, users would declare text columns as “CHAR”. A “CHAR” column always has the same number of characters. If you feed it too few characters, it adds spaces to the end; if you feed it too many characters, it truncates the last ones.

在 2002 时，MySQL 实现了一加速功能 [speed boost](http://dev.mysql.com/doc/refman/5.7/en/static-format.html) ， 只要用户能保证 table 中每个 row 占用相同的 byte 空间，就能利用到这个加速特性。为了保证 row 占用空间不变，用户可以把 text column 声名为 "CHAR" 类型， “CHAR” 类型的 column 只占用固定个数的字符。如果你使用的字符比它占用的空间少，它会自动在字符末尾填充空格；但如果你使用的字符比它占用的空间多，它就会把多余的字符截断。

When MySQL developers first tried UTF-8, with its back-in-the-day six bytes per character, they likely balked: a CHAR(1) column would take six bytes; a CHAR(2) column would take 12 bytes; and so on.

当 MySQL 开发者第一次尝试支持 UTF-8 编码时，他们犹豫了，那个时候，UTF-8 编码中每个字符占用 6 个 byte ，也就是说：一个 CHAR(1) column 类型就要占用 6 个 byte ；而 CHAR(2) column 类型就占用12 byte;以此类推。

Let’s be clear: that initial behavior, which was never released, was  _correct_. It was well documented and widely adopted, and anybody who understood UTF-8 would agree that it was right.

说明一下：这个最初版本虽然是 _正确_ 的，但从来没有正确公开发布过。它有广泛的文档说明，而且每个了解 UTF-8 的人都同意这样的做法。

But clearly, a MySQL developer (or businessperson) was concerned that a user or two would do two things:


1.  Choose CHAR columns. (The CHAR format is a relic nowadays. Back then, MySQL was faster with CHAR columns. Ever since 2005, it’s not.)
2.  Choose to encode those CHAR columns as “utf8”.

注意，MySQL 开发者（也许是商人）很关心用户做的下面这两件事情。

- 1.  选择 CHAR column 。（CHAR 类型现在已经算上古遗物了，基本没人用。但在 2005 年之前， CHAR 类型在 MySQL 中速度很快。）
- 2.  选择 CHAR column 的编码类型为 "utf8" 。

My guess is that MySQL developers broke their “utf8” encoding to help these users: users who both 1) tried to optimize for space and speed; and 2) failed to optimize for speed and space.

我猜测 MySQL 开发者破坏 "utf8" 编码方案是为了帮助这样的用户：1) 尝试优化空间和速度；优化空间和速度失败了。

TODO 译： 没理解这段话的意思。谁？为什么？怎么样？优化空间和速度的？

Nobody won. Users who wanted speed and space were  _still_  wrong to use “utf8” CHAR columns, because those columns were still bigger and slower than they had to be. And developers who wanted correctness were wrong to use “utf8”, because it can’t store “💩”.

但无人受益。想要优化空间和速度的 _仍然_ 在使用错误的 "utf8" CHAR column 类型，因为这些 column 仍然占用很大空间而且很慢。而希望使用正确编码类型的人，却误用了 "utf8" ，因为它存储不了 “💩” 。

Once MySQL published this invalid character set, it could never fix it: that would force every user to rebuild every database. MySQL finally released UTF-8 support in  [2010](https://dev.mysql.com/doc/relnotes/mysql/5.5/en/news-5-5-3.html), with a different name: “utf8mb4”.

MySQL 发布这个错误字符集就，就永远没法修复了：因为必须让所有人 rebuild 所有的数据库才能处理。但 MySQL 最终还是在 [2010](https://dev.mysql.com/doc/relnotes/mysql/5.5/en/news-5-5-3.html) 年发布了一个新的编码类型 "utf8mb4" 来支持真正的 UTF-8 编码。

### Why it’s so frustrating
### 为什么这么令人沮丧

Clearly I was frustrated this week. My bug was hard to find because I was fooled by the name “utf8”. And I’m not the only one — almost every article I found online touted “utf8” as, well, UTF-8.

我这周感觉十分受挫。因为被 "utf8" 这个编码名称给误导了，我解决 bug 的过程太难了。而且我不是唯一一个，我在互联网上找到的所有关于 "utf8" 的讨论，都被误导了－－嗯，应该说是关于 UTF-8 的讨论才对。

The name “utf8” was always an error. It’s a proprietary character set. It created new problems, and it didn’t solve the problem it meant to solve.

这个 "utf8" 编码完全是个错误。它是专有的编码集（不符合规范）。它既没解决原本要解决的问题，还引入了新的问题。

It’s false advertising.

TODO 翻译 advertising 在这个上下文中的含义。应该不是广告的意思吧。

### My take-away lessons
### 我的教训

- 1.  Database systems have subtle bugs and oddities, and you can avoid a lot of bugs by avoiding database systems.
- 1.  数据库系统本身也是会有坑的，比如一些古怪的 bug ，正确选择数据库，可以避开很多问题。
- 2.  If you  _need_  a database, don’t use MySQL or MariaDB. Use  [PostgreSQL](http://www.postgresql.org/).
- 2.  如果 _必须_ 使用数据库，不要使用 MySQL 或 MariaDB ，请使用  [PostgreSQL](http://www.postgresql.org/)。
- 3.  If you  _need_  to use MySQL or MariaDB, never use “utf8”. Always use “utf8mb4” when you want UTF-8.  [Convert your database](https://mathiasbynens.be/notes/mysql-utf8mb4#utf8-to-utf8mb4)  now to avoid headaches later.
- 3.  如果 _必须_ 使用 MySQL 或 MariaDB ，永远不要使用 “utf8” 。需要 UTF-8 编码的地方，请使用 “utf8mb4” 。如果已经使用了 “utf8” ，参考这个文章，[转换数据库编码](https://mathiasbynens.be/notes/mysql-utf8mb4#utf8-to-utf8mb4)，避免踩坑。




## 说明
-  使用vim-im插件终于可以在  Boox Note Plus 的 Termux 中使用 vim + 极点五笔输入中文了。这篇文章主要在 NotePlus 上完成，特此纪念。
- MySQL utf8 问题自己遇到两次，只知道解决方法，不知道具体原因。这篇文章终于让自己了解一些历史原因。
- 这个 Bug 算是 MySQL 的历史包袱吧？自己写的程序应该也有很多类似问题，只不过用的人少，没被发现而已。
- UTF-8编码最大占用几个 byte ？ 
   答：2019年实际使用的场景 4 byte 足够了。 
   1998 年制定的 RFC-2279 标准所定义的 original unicode  即[Basic_Multilingual_Plane](https://en.wikipedia.org/wiki/Plane_(Unicode)#Basic_Multilingual_Plane) U+0800 ~ U+FFFF，只需要 3 byte 就够了。
   而 2003 年制定的 RFC-3629 标准  U+10000 ~ U+1FFFFF ，需要 4 byte 。
   详细说明参考 stackoverflow [^max_byte_utf8_1] 。
   所以准确的说， MySQL 在 2002 年实现 utf8 时，业界标准 BMP 定义的 unicode 确实只需要 3 byte ，但是后来标准变成了 4 byte ，但 MySQL 没有及时按新标准调整。所以才出现本文的问题。

```txt
Without further context, I would say that the maximum number of bytes for a character in UTF-8 is

answer: 6 bytes

The author of the accepted answer correctly pointed this out as the "original specification". That was valid through RFC-2279 1. As J. Cocoe pointed out in the comments below, this changed in 2003 with RFC-3629 2, which limits UTF-8 to encoding for 21 bits, which can be handled with the encoding scheme using four bytes.

answer if covering all unicode: 4 bytes

But, in Java <= v7, they talk about a 3-byte maximum for representing unicode with UTF-8? That's because the original unicode specification only defined the basic multi-lingual plane (BMP), i.e. it is an older version of unicode, or subset of modern unicode. So

answer if representing only original unicode, the BMP: 3 bytes

But, the OP talks about going the other way. Not from characters to UTF-8 bytes, but from UTF-8 bytes to a "String" of bytes representation. Perhaps the author of the accepted answer got that from the context of the question, but this is not necessarily obvious, so may confuse the casual reader of this question.

Going from UTF-8 to native encoding, we have to look at how the "String" is implemented. Some languages, like Python >= 3 will represent each character with integer code points, which allows for 4 bytes per character = 32 bits to cover the 21 we need for unicode, with some waste. Why not exactly 21 bits? Because things are faster when they are byte-aligned. Some languages like Python <= 2 and Java represent characters using a UTF-16 encoding, which means that they have to use surrogate pairs to represent extended unicode (not BMP). Either way that's still 4 bytes maximum.

answer if going UTF-8 -> native encoding: 4 bytes

So, final conclusion, 4 is the most common right answer, so we got it right. But, mileage could vary.
```

[en.wikipedia UTF-8#Description](https://en.wikipedia.org/wiki/UTF-8#Description)

 Bits of code point  |  First code point  |  Last code point  | Bytes in sequence | Byte 1   |  Byte 2  |  Byte 3  | Byte 4   |
-------------------- | ------------------ | ----------------- | ----------------- | -------- | -------- | -------- | -------- |
7                    |    U+0000          |      U+007F       |      1            | 0xxxxxxx |          |          |          |                        
11                   |    U+0080          |      U+07FF       |      2            | 110xxxxx | 10xxxxxx |          |          |   
16                   |    U+0800          |      U+FFFF       |      3            | 1110xxxx | 10xxxxxx | 10xxxxxx |          | 
21                   |    U+10000         |      U+1FFFFF     |      4            | 11110xxx | 10xxxxxx | 10xxxxxx | 10xxxxxx |


[^MySQL_utf8_utf8mb4]: [In MySQL, never use “utf8”. Use “utf8mb4” Adam Hooper 14, 2016 ](https://medium.com/@adamhooper/in-mysql-never-use-utf8-use-utf8mb4-11761243e434) 
[^max_byte_utf8_1]:  [what-is-the-maximum-number-of-bytes-for-a-utf-8-encoded-character](https://stackoverflow.com/questions/9533258/what-is-the-maximum-number-of-bytes-for-a-utf-8-encoded-character)
[^max_byte_utf8_2]:  [max-bytes-in-a-utf-8-char/](https://stijndewitt.com/2014/08/09/max-bytes-in-a-utf-8-char/)

<!--stackedit_data:
eyJoaXN0b3J5IjpbNjIyMTQyNTEwLC0yODQxOTI5NDAsLTEwMD
kyODQ3NDcsLTg1MzUwMzA3OSw4MTk2MTgzNDAsLTI4NTMxMTQ4
MywtNzMxMjU5ODI0LC00MjA0OTU5OTYsOTE2Njc0MzUyLC0xNj
IwNjkxNTcyXX0=
-->