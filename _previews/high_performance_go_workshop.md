---
layout: post
title:  "[译]High Performance Go Workshop"
date:   2020-01-01 12:00:00 +0800
tags:   tech
---

* category
{:toc}




# High Performance Go Workshop [^HighPerformanceWorkShopEN] [^HighPerformanceWorkShopGithub] [^HighPerformanceWorkShopCN1] [^HighPerformanceWorkShopCN2] 


> 其他值得参考的文章 [^GoMillionTCP] [^PerformanceIntruction]

Dave Cheney[dave@cheney.net](mailto:dave@cheney.net)Version Dotgo-2019-3-G660848,2019-04-26



## Overview

The goal for this workshop is to give you the tools you need to diagnose performance problems in your Go applications and fix them.

Through the day we’ll work from the small — learning how to write benchmarks, then profiling a small piece of code. Then step out and talk about the execution tracer, the garbage collector and tracing running applications. The remainder of the day will be a chance for you to ask questions, experiment with your own code.

You can find the latest version of this presentation at

[http://bit.ly/dotgo2019](http://bit.ly/dotgo2019)

## Welcome

Hello and welcome! 🎉

The goal for this workshop is to give you the tools you need to diagnose performance problems in your Go applications and fix them.

Through the day we’ll work from the small — learning how to write benchmarks, then profiling a small piece of code. Then step out and talk about the execution tracer, the garbage collector and tracing running applications. The remainder of the day will be a chance for you to ask questions, experiment with your own code.

### Instructors

-   Dave Cheney  [dave@cheney.net](mailto:dave@cheney.net)
    

### License and Materials

This workshop is a collaboration between  [David Cheney](https://twitter.com/davecheney)  and  [Francesc Campoy](https://twitter.com/francesc).

This presentation is licensed under the  [Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/)  licence.

### Prerequisites

The are several software downloads you will need today.

#### The workshop repository

Download the source to this document and code samples at  [https://github.com/davecheney/high-performance-go-workshop](https://github.com/davecheney/high-performance-go-workshop)

#### Laptop, power supplies, etc.

The workshop material targets Go 1.12.

[Download Go 1.12](https://golang.org/dl/)

If you’ve already upgraded to Go 1.13 that’s ok. There are always some small changes to optimisation choices between minor Go releases and I’ll try to point those out as we go along.

#### Graphviz

The section on pprof requires the  `dot`  program which ships with the  `graphviz`  suite of tools.

-   Linux:  `[sudo] apt-get install graphviz`
    
-   OSX:
    
-   MacPorts:  `sudo port install graphviz`
    
-   Homebrew:  `brew install graphviz`
    
-   [Windows](https://graphviz.gitlab.io/download/#Windows)  (untested)
    

#### Google Chrome

The section on the execution tracer requires Google Chrome. It will not work with Safari, Edge, Firefox, or IE 4.01. Please tell your battery I’m sorry.

[Download Google Chrome](https://www.google.com/chrome/)

#### Your own code to profile and optimise

The final section of the day will be an open session where you can experiment with the tools you’ve learnt.

### One more thing …

This isn’t a lecture, it’s a conversation. We’ll have lots of breaks to ask questions.

If you don’t understand something, or think what you’re hearing not correct, please ask.

## 1. The past, present, and future of Microprocessor performance 微处理器性能的过去，现在和未来

This is a workshop about writing high performance code. In other workshops I talk about decoupled design and maintainability, but we’re here today to talk about performance.

I want to start today with a short lecture on how I think about the history of the evolution of computers and why I think writing high performance software is important .

今天演讲的主要内容主要是：
有关计算机发展历史,我的一些思考;
为什么编写高性能的软件很重要。

The reality is that software runs on hardware, so to talk about writing high performance code, first we need to talk about the hardware that runs our code.

因为软件是在硬件上运行的，所以，要想讨论如何编写高性能软件的话题，我们先说一说运行代码的硬件。

### 1.1. Mechanical Sympathy 机械共情

Sympathy![image 20180818145606919](https://dave.cheney.net/high-performance-go-workshop/images/image-20180818145606919.png)

There is a term in popular use at the moment, you’ll hear people like Martin Thompson or Bill Kennedy talk about “mechanical sympathy”.

你可能从 马丁·汤普森（Martin Thompson）或比尔·肯尼迪（Bill Kennedy）讨论过 “Mechanical Sympathy” 这一术语。

The name "Mechanical Sympathy" comes from the great racing car driver Jackie Stewart, who was a 3 times world Formula 1 champion. He believed that the best drivers had enough understanding of how a machine worked so they could work in harmony with it.

"Mechanical Sympathy" 最早由曾三度获得世界一级方程式赛车冠军的 赛车手杰基·斯图尔特（Jackie Stewart）提出。
他认为，好的驾驶员肯定对机器的工作原理有足够了解，这样他们才能与机器和谐工作。

To be a great race car driver, you don’t need to be a great mechanic, but you need to have more than a cursory understanding of how a motor car works.

要成为一名出色的赛车手，您不需要成为一名出色的机械师，但您需要对汽车的工作原理有一个粗略的了解。

I believe the same is true for us as software engineers. I don’t think any of us in this room will be a professional CPU designer, but that doesn’t mean we can ignore the problems that CPU designers face.

我觉得软件工程师也是一样。
我们不必成为专业的CPU设计者，但需要了解CPU设计人员所面临的问题。

### 1.2. Six orders of magnitude

There’s a common internet meme that goes something like this;

![jalopnik](https://dave.cheney.net/high-performance-go-workshop/images/jalopnik.png)

Of course this is preposterous, but it underscores just how much has changed in the computing industry.

As software authors all of us in this room have benefited from Moore’s Law, the doubling of the number of available transistors on a chip every 18 months, for 40 years. No other industry has experienced a  six order of magnitude  [1](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#_footnotedef_1 "View footnote.") improvement in their tools in the space of a lifetime.

这里所有软件开发者都受益于摩尔定律，即40年来每18个月将芯片上可用晶体管的数量增加一倍。
没有其他行业在使用寿命内对工具的改进达到六个数量级。

But this is all changing.

但这种好处马上将消失。



### 1.3. Are computers still getting faster? 计算机会越来越快吗？

So the fundamental question is, confronted with statistic like the ones in the image above, should we ask the question  _are computers still getting faster_?

我们应该关心的问题是：计算机会一直越来越快吗？

If computers are still getting faster then maybe we don’t need to care about the performance of our code, we just wait a bit and the hardware manufacturers will solve our performance problems for us.

如果计算机的速度仍在不断提高，那么也许我们不需要关心代码的性能，只需稍等一下，硬件制造商就会为我们解决性能问题。



#### 1.3.1. Let’s look at the data

This is the classic data you’ll find in textbooks like  _Computer Architecture, A Quantitative Approach_  by John L. Hennessy and David A. Patterson. This graph was taken from the 5th edition

下面的经典数据可以在 《 计算机体系结构》，《约翰·亨尼西和大卫·帕特森的定量方法》找到 。 此图摘自第5版。

![2313.processorperf](https://community.cadence.com/cfs-file/__key/communityserver-blogs-components-weblogfiles/00-00-00-01-06/2313.processorperf.jpg)

In the 5th edition, Hennessey and Patterson argue that there are three eras of computing performance

-   The first was the 1970’s and early 80’s which was the formative years. Microprocessors as we know them today didn’t really exist, computers were built from discrete transistors or small scale integrated circuits. Cost, size, and the limitations in the understanding of material science were the limiting factor.
    
-   From the mid 80s to 2004 the trend line is clear. Computer integer performance improved on average by 52% each year. Computer power doubled every two years, hence people conflated Moore’s law — the doubling of the number of transistors on a die, with computer performance.
    
-   Then we come to the third era of computer performance. Things slow down. The aggregate rate of change is 22% per year.


在第5版中，轩尼诗（Hennessey）和帕特森（Patterson）认为计算性能存在三个时代:

- 1970 ~ 1985 是计算机形成的年代。微处理器不存在。计算机由集成电路组成　。成本，尺寸，及材料科学的发展是主要限制因素。 
- 1985 ~ 2004 是调整发展的年代。计算机性能每年提高52％。计算能力每两年翻一翻。因此人们将摩尔定律和计算机性能混为一谈，摩尔定律是指芯片上晶休管的数量每两年翻一翻。
- 2004 ~ 至今(2012)　增涨放缓，每年总共增加 22% 。


That previous graph only went up to 2012, but fortunately in 2012  [Jeff Preshing](http://preshing.com/20120208/a-look-back-at-single-threaded-cpu-performance/)  wrote a  [tool to scrape the Spec website and build your own graph](https://github.com/preshing/analyze-spec-benchmarks).

上图数据只到2012年。但　Jeff Preshing 在2012年写了一个爬取　Spec 网站数生成图表的工具，下图就是此工具生成的　1995～2017　年间的　Spec　数据　。


![intgraph](https://dave.cheney.net/high-performance-go-workshop/images/int_graph.png)

So this is the same graph using Spec data from 1995 til 2017.

To me, rather than the step change we saw in the 2012 data, I’d say that  _single core_  performance is approaching a limit. The numbers are slightly better for floating point, but for us in the room doing line of business applications, this is probably not that relevant.

观察 2012 年的数据可以发现，_单核_ 整数运算单元的性能已经接近极限。
虽然浮点数运算单元的数据可能会稍好一些，但对我们做业务程序的人来说，区别不大。

> NOTE: 整数运算单元、浮点数运算单元 integer performance  floating point performance 是有区别的。详细情况 TODO google CSAPP 了解浮点数计算


#### 1.3.2. Yes, computer are still getting faster, slowly 计算机还在变快，但是在慢慢变快 [^CPUMax4GHz]


> The first thing to remember about the ending of Moore’s law is something Gordon Moore told me. He said "All exponentials come to an end". — [John Hennessy](https://www.youtube.com/watch?v=Azt8Nc-mtKM)

This is Hennessy’s quote from Google Next 18 and his Turing Award lecture. His contention is yes, CPU performance is still improving. However, single threaded integer performance is still improving around 2-3% per year. At this rate its going to take 20 years of compounding growth to double integer performance. Compare that to the go-go days of the 90’s where performance was doubling every two years.

Why is this happening?

> 戈登·摩尔告诉我，摩尔定律中指数增涨的过程即将结束。 — [John Hennessy](https://www.youtube.com/watch?v=Azt8Nc-mtKM)

这是轩尼诗（Hennessy）在Google Next 18上的引用以及他在图灵奖上的演讲。 他认为CPU性能仍在提高,但是单线程 integer performance 每年仅提高2-3％左右。 以这种速度，它将需要20年的复合增长才能使整数运算的性能翻倍。 相比之下，90年代的发展趋势是每两年翻一番。

到底发生了什么呢？



### 1.4. Clock speeds 时钟速度

![stuttering](https://dave.cheney.net/high-performance-go-workshop/images/stuttering.png)

This graph from 2015 demonstrates this well. The top line shows the number of transistors on a die. This has continued in a roughly linear trend line since the 1970’s. As this is a log/lin graph this linear series represents exponential growth.

However, If we look at the middle line, we see clock speeds have not increased in a decade, we see that cpu speeds stalled around 2004

The bottom graph shows thermal dissipation power; that is electrical power that is turned into heat, follows a same pattern—clock speeds and cpu heat dissipation are correlated.


2015 年的这张图很好地说明了这一点。

第一条线显示了芯片上的晶体管数量。自 1970 年代以来，一直以线性趋势持续增长。 由于这是 log/lin 图，因此该图表示的是指数增长。

如果我们看中间的线，会发现时钟速度近十年来没有增加， CPU 速度在 2004 年左右停滞了。

最下面一条线，表示散热功率；即变成热量的电能，它和时钟速度的走向差不多，所以时钟速度和 cpu 散热也有一些关系。



### 1.5. Heat 发热

Why does a CPU produce heat? It’s a solid state device, there are no moving components, so effects like friction are not (directly) relevant here.

为什么 CPU 会产生热量？它是固定不动的设备，也没有什么需要来回活动的零件，所以这里产生的热量肯定和摩擦生热无关。


This digram is taken from a great  [data sheet produced by TI](http://www.ti.com/lit/an/scaa035b/scaa035b.pdf). In this model the switch in N typed devices is attracted to a positive voltage P type devices are repelled from a positive voltage.

下图来自 TI 公司的数据。
在这个模型中，N device 中的开关被吸引到正电压上，P device 被排斥在正电压上。

![cmos inverter](https://dave.cheney.net/high-performance-go-workshop/images/cmos-inverter.png)

The power consumption of a CMOS device, which is what every transistor in this room, on your desk, and in your pocket, is made from, is combination of three factors.

1.  Static power. When a transistor is static, that is, not changing its state, there is a small amount of current that leaks through the transistor to ground. The smaller the transistor, the more leakage. Leakage increases with temperature. Even a minute amount of leakage adds up when you have billions of transistors!
    
2.  Dynamic power. When a transistor transitions from one state to another, it must charge or discharge the various capacitances it is connected to the gate. Dynamic power per transistor is the voltage squared times the capacitance and the frequency of change. Lowering the voltage can reduce the power consumed by a transistor, but lower voltages causes the transistor to switch slower.
    
3.  Crowbar, or short circuit current. We like to think of transistors as digital devices occupying one state or another, off or on, atomically. In reality a transistor is an analog device. As a switch a transistor starts out  _mostly_  off, and transitions, or switches, to a state of being  _mostly_  on. This transition or switching time is very fast, in modern processors it is in the order of pico seconds, but that still represents a period of time when there is a low resistance path from Vcc to ground. The faster the transistor switches, its frequency, the more heat is dissipated.
    

我们身边能看到的所有 CMOS 设备功耗主要由以下三部分组成。

- 1.静态功耗。当晶体管没有状态变化时，只有少量电流泄漏到大地。晶体管越小，泄漏越多；温度越高，泄漏越多。成长上亿的晶体管泄漏的电量累积到一起是非常巨大的。
- 2.动态功耗。当晶体管进行状态转换时，要对栅极上的电容充放电。每个晶体管的动态功耗是 电压x电容x频率^2 。低电压会压低晶体管的能耗。但低电压也会使晶体管的开关速度变慢。
- 3.短路。我们经常把晶体管当成数字设备，但它实际上是模拟设备。一个启动时是 off 状态，状态切换时是 on 状态的 switch 。这个转换过程很快，现在处理器大约只需要　一皮秒(pico second)，但当从 Vcc 到 ground 的电阻路径很低时，这个时间也不算短了。这个　switch 转换得越快，频率越高，它的温度也越高。

> TODO gate 栅极 技术概念是指什么?




### 1.6. The end of Dennard scaling ( 丹纳德 扩展的终结 )

To understand what happened next we need to look to a paper written in 1974 co-authored by  [Robert H. Dennard](https://en.wikipedia.org/wiki/Robert_H._Dennard). Dennard’s Scaling law states roughly that as transistors get smaller their  [power density](https://en.wikipedia.org/wiki/Power_density)  stays constant. Smaller transistors can run at lower voltages, have lower gate capacitance, and switch faster, which helps reduce the amount of dynamic power.

参考 Robert H. Dennard 1974年论文可知。
根据 Dennard’s Scaling 定律，在晶体管小到一定程度后，其 power density 保持恒定。
晶体管越小，所需要的电压越低，栅极电容也越小，并且开关速度更快，这样总的动态功耗反而会降低。

> NOTE: power density (功率密度) watts (瓦特 功率) hot plate (热铁板)  nuclear reactor (核反应堆) rocket nozzle (火箭喷射器)  

So how did that work out?

真的是这样吗？

![Screen Shot 2014 04 14 at 8.49.48 AM](http://semiengineering.com/wp-content/uploads/2014/04/Screen-Shot-2014-04-14-at-8.49.48-AM.png)

It turns out not so great. As the gate length of the transistor approaches the width of a few silicon atom, the relationship between transistor size, voltage, and importantly leakage broke down.

并非这样。当晶体管的栅极长度小到 silicon atom （硅原子）宽度时。晶体管大小、电压、leakage　之间的规律发生了变化。


It was postulated at the  [Micro-32 conference in 1999](https://pdfs.semanticscholar.org/6a82/1a3329a60def23235c75b152055c36d40437.pdf)  that if we followed the trend line of increasing clock speed and shrinking transistor dimensions then within a processor generation the transistor junction would approach the temperature of the core of a nuclear reactor. Obviously this is was lunacy. The Pentium 4  [marked the end of the line](https://arstechnica.com/uncategorized/2004/10/4311-2/)  for single core, high frequency, consumer CPUs.

在1999的 Micro-32 会议曾经推测，如果继续提高时钟速度，减小晶体管尺寸，按以上趋势图的发展，将生产出晶体箮结温度达到核反应堆温度的处理理。
这有点荒谬了。
单核心、高频率的消费类CPU奔腾4处理器，是最后一个符合上面趋势线的CPU。



Returning to this graph, we see that the reason clock speeds have stalled is because cpu’s exceeded our ability to cool them. By 2006 reducing the size of the transistor no longer improved its power efficiency.

我们再回过头讨论刚才这张图。
时钟速度停止增长的主要原因是冷却 CPU 的技术能力跟不上。
所以到 2006 年时，减小晶体管尺寸已经无法提高功率了。


We now know that CPU feature size reductions are primarily aimed at reducing power consumption. Reducing power consumption doesn’t just mean “green”, like recycle, save the planet. The primary goal is to keep power consumption, and thus heat dissipation,  [below levels that will damage the CPU](https://en.wikipedia.org/wiki/Electromigration#Practical_implications_of_electromigration).

我们知道，减小 CPU 尺寸的主要原因是为了降低能耗。
降低能耗并非为了“绿色环保”，不是为了节约资源，保护地球环境。
主要目标只是将保持能耗，在现有散热能力下，防止热量过高损坏CPU。


![stuttering](https://dave.cheney.net/high-performance-go-workshop/images/stuttering.png)

But, there is one part of the graph that is continuing to increase, the number of transistors on a die. The march of cpu features size, more transistors in the same given area, has both positive and negative effects.

但是图中晶体箮的数量仍然在持续增加。
CPU尺寸变大，其中就能放更多晶体箮。这即有正面影响也有负面影响。

Also, as you can see in the insert, the cost per transistor continued to fall until around 5 years ago, and then the cost per transistor started to go back up again.

上图中左上角的插图显示，晶体管单价在2012年前一直在下降，随后2015年单价又开始上长。

> NOTE: 图中显示的是1美元能购买的晶体管数量，所以跟单价趋势刚好相反。



![moores law](https://whatsthebigdata.files.wordpress.com/2016/08/moores-law.png)

Not only is it getting more expensive to create smaller transistors, it’s getting harder. This report from 2016 shows the prediction of what the chip makers believed would occur in 2013; two years later they had missed all their predictions, and while I don’t have an updated version of this report, there are no signs that they are going to be able to reverse this trend.

制造更小的晶体管不仅更贵了，也更难了。
上面2016年的报告显示，芯片制造商预测2013年 physical gate length 从 20 nanometers　开始每年减小 2 nanometer 。
但再看两年后2015的预测图发现，它们显示没有达到之前预期。
虽然我没有相关报告的最新版本，但没有迹象表明谁能扭转这一趋势。

It is costing intel, TSMC, AMD, and Samsung billions of dollars because they have to build new fabs, buy all new process tooling. So while the number of transistors per die continues to increase, their unit cost has started to increase.

英特尔，台积电，AMD和三星等厂商在建厂，购置生产工具的花费高达数十亿美元。
即使单个芯片的晶体管数量在增加，这些芯片的单位成本也仍然在上涨。


> Even the term gate length, measured in nano meters, has become ambiguous. Various manufacturers measure the size of their transistors in different ways allowing them to demonstrate a smaller number than their competitors without perhaps delivering. This is the Non-GAAP Earning reporting model of CPU manufacturers.

> 术语 gate length (栅极长度)的单位 nano meter（纳米）定义也有些模糊。
> 每个厂商测量晶体管大小的方式都不一样，所以它们总能展示出比竞厂更小尺寸的样品。
> 这是CPU制造商 Non-GAAP 收益报告模型。


### 1.7. More cores

![y5cdp7nhs2uy](https://i.redd.it/y5cdp7nhs2uy.jpg)

With thermal and frequency limits reached it’s no longer possible to make a single core run twice as fast. But, if you add another cores you can provide twice the processing capacity — if the software can support it.

由于温度和频率的限制，想让单核心运行速度变快两倍已经不太容易了。
但是，假如软件能同时利用好两个 core ，那只要再加一个 CPU 就能轻松让运行速度快两倍。

In truth, the core count of a CPU is dominated by heat dissipation. The end of Dennard scaling means that the clock speed of a CPU is some arbitrary number between 1 and 4 Ghz depending on how hot it is. We’ll see this shortly when we talk about benchmarking.

CPU 的核心数量主要由散热情况决定。
由 Dennard’s Scaling 定律可知，CPU 时钟速度肯定是在 1 到 4 Ghz 之间，具体大小由它的热度决定。
一会讨论基准测试时，我们就能看到这一点。


### 1.8. Amdahl’s law (阿姆达尔定律)

CPUs are not getting faster, but they are getting wider with hyper threading and multiple cores. Dual core on mobile parts, quad core on desktop parts, dozens of cores on server parts. Will this be the future of computer performance? Unfortunately not.

CPU虽然没有变得更快，但由于超纯种和多核心技术的发展，CPU变的更“宽”了。
在移动设备上使用双核处理器，桌面设备上使用四核处理器，在服务器中使用几十核心的处理器。
在以后的日子，只要增加核心数量 ，计算机性能就能一直提升吗？
当然不可能了。

Amdahl’s law, named after the Gene Amdahl the designer of the IBM/360, is a formula which gives the theoretical speedup in latency of the execution of a task at fixed workload that can be expected of a system whose resources are improved.

Amdahl 定律，是以 IBM/360 的设计师 Gene Amdahl 名字命名的。
此定律中的公式能计算出，在任务工作量不变的情况下，能无限增加系统资源，最快能提前多久完成任务。

> NOTE: “提前多久完成任务” 表达的含义与 “能提升工作效率多少倍” “提速多少倍” 一样。

![AmdahlsLaw](https://upload.wikimedia.org/wikipedia/commons/e/ea/AmdahlsLaw.svg)

Amdahl’s law tells us that the maximum speedup of a program is limited by the sequential parts of the program. If you write a program with 95% of its execution able to be run in parallel, even with thousands of processors the maximum speedup in the programs execution is limited to 20x.

Amdahl 定律告诉我们，能提速多少，取决于程序当中能够顺序执行的部分有多少。
假设你的程序中有 95% 的代码都能并行执行，即使有上千个处理器，最多也只能提速 20 倍。

Think about the programs that you work on every day, how much of their execution is parralisable?

想想你每天写的程序，它们当中有多少是可以并行执行的呢？


> TODO 超线程、整数运算单元、浮点数运算单元 integer performance  floating point performance
> 
> 参考链接：https://www.zhihu.com/question/20277695/answer/14588735
>
> Intel的超线程技术，目的是为了更充分地利用一个单核CPU的资源。CPU在执行一条机器指令时，并不会完全地利用所有的CPU资源，而且实际上，是有大量资源被闲置着的。超线程技术允许两个线程同时不冲突地使用CPU中的资源。比如一条整数运算指令只会用到整数运算单元，此时浮点运算单元就空闲了，若使用了超线程技术，且另一个线程刚好此时要执行一个浮点运算指令，CPU就允许属于两个不同线程的整数运算指令和浮点运算指令同时执行，这是真的并行。我不了解其它的硬件多线程技术是怎么样的，但单就超线程技术而言，它是可以实现真正的并行的。但这也并不意味着两个线程在同一个CPU中一直都可以并行执行，只是恰好碰到两个线程当前要执行的指令不使用相同的CPU资源时才可以真正地并行执行。




### 1.9. Dynamic Optimisations 动态优化

With clock speeds stalled and limited returns from throwing extra cores at the problem, where are the speedups coming from? They are coming from architectural improvements in the chips themselves. These are the big five to seven year projects with names like  [Nehalem, Sandy Bridge, and Skylake](https://en.wikipedia.org/wiki/List_of_Intel_CPU_microarchitectures#Pentium_4_/_Core_Lines).

既然时钟速度停滞不前，通过增加 CPU core 数量带来的提速又十分有限，那么近年来的性能提升又来自哪里呢？
这主要是由于芯片本身的架构改进。像 Nehalem, Sandy Bridge, Skylake 这些微处理器架构项目一般都要持久五到七年时间。

Much of the improvement in performance in the last two decades has come from architectural improvements:

可以说，过去二十年间的性能提升大都来源于架构的改进。



#### 1.9.1. Out of order execution 乱序执行

Out of Order, also known as super scalar, execution is a way of extracting so called  _Instruction level parallelism_  from the code the CPU is executing. Modern CPUs effectively do SSA at the hardware level to identify data dependencies between operations, and where possible run independent instructions in parallel.

乱序，也称为超标量，是一种能在运行中的 CPU 代码中，执行行指令级并行优化的方法。
现代 CPU 能高效执行 SSA 过程，因为它能在硬件层识别各种数据操作之间的依赖关系，并尽可能并行执行。

TODO [SSA Static single assignment 静态单赋值](https://en.wikipedia.org/wiki/Static_single_assignment_form)


However there is a limit to the amount of parallelism inherent in any piece of code. It’s also tremendously power hungry. Most modern CPUs have settled on six execution units per core as there is an n squared cost of connecting each execution unit to all others at each stage of the pipeline.

但是每段代码的并行量是有限的。而且十分费电。
现代CPU中，每个 core 配置六个执行单元，在执行指令流水线的过程中，将每个执行单元与其他执行单元连接到一起的成本是 N^2 。

#### 1.9.2. Speculative execution 预测执行

Save the smallest micro controllers, all CPUs utilise an  _instruction pipeline_  to overlap parts of in the instruction fetch/decode/execute/commit cycle.

除了最小的微型控制器外，所有的 CPU 都能在执行 fetch/decode/execute/commit 指令周期的过程中，利用 _指令流水线_ 重叠（并行）执行其中部分指令。

> NOTE [What are the smallest microcontrollers?](https://electronics.stackexchange.com/questions/84800/what-are-the-smallest-microcontrollers)


![800px Fivestagespipeline](https://upload.wikimedia.org/wikipedia/commons/thumb/2/21/Fivestagespipeline.png/800px-Fivestagespipeline.png)

The problem with an instruction pipeline is branch instructions, which occur every 5-8 instructions on average. When a CPU reaches a branch it cannot look beyond the branch for additional instructions to execute and it cannot start filling its pipeline until it knows where the program counter will branch too. Speculative execution allows the CPU to "guess" which path the branch will take  _while the branch instruction is still being processed!_

可是，指令流水线的分支指令平均每 5-8 个指令周期才会执行一次。
当 CPU 达到分支时，它不能从当前分支之外寻找要执行的指令，必须从 program counter 获取到下一个要切换的分支后，才能填充流水线。
预测执行功能，能让 CPU 在执行 分支指令 的过程中，猜测下一次执行的分支路径。

> NOTE Program counter 程序计数器，也叫指令指针，用于保存程序下一次要执行的指令。
> A program counter is a register in a computer processor that contains the address (location) of the instruction being executed at the current time. As each instruction gets fetched, the program counter increases its stored value by 1.
> [What is program counter? And how it work?](https://www.quora.com/What-is-program-counter-And-how-it-work)


If the CPU predicts the branch correctly then it can keep its pipeline of instructions full. If the CPU fails to predict the correct branch then when it realises the mistake it must roll back any change that were made to its  _architectural state_. As we’re all learning through Spectre style vulnerabilities, sometimes this rollback isn’t as seamless as hoped.

如果 CPU 预测到正确的分支，就能保持它的分支流水线一直是满的。
如果 CPU 预测错了分支，它就必须在发现错误时，立即回滚之前对 _architectural state_ 的改动。
像我们在 Spectre style vulnerabilities 学习到的那样，有时这种回滚并非无缝的。


Speculative execution can be very power hungry when branch prediction rates are low. If the branch is misprediction, not only must the CPU backtrace to the point of the misprediction, but the energy expended on the incorrect branch is wasted.

如果分支预测的正确率很低时，是十分费电的。
分支预测失败时，不仅 CPU 要回溯到之前的状态，花在分支在的能量也浪费了。


All these optimisations lead to the improvements in single threaded performance we’ve seen, at the cost of huge numbers of transistors and power.

这些在单线程性能上的提升，都是以消耗大量晶体管和电力为代价的。

> Cliff Click has a  [wonderful presentation](https://www.youtube.com/watch?v=OFgxAFdxYAQ)  that argues out of order and speculative execution is most useful for starting cache misses early thereby reducing observed cache latency.

> Cliff Click 有一个精彩的演示，论证了 乱序执行和分支预测 能降低 cache latency .

TODO starting cache misses early 如何理解？ 因为提前预测执行了代码，所以把相关数据和代码也都缓存到 L1 L2 L3 等多级缓存中较近的一级中了。所以当真正使用相关指令时，也就能更快找到所需要的数据和指令，进而减少 cache latency 。

TODO [CPU 为何非得要用乱序执行和预测执行呢？](https://www.v2ex.com/t/420690)

TODO [CPU Cache 机制以及 Cache miss](https://www.cnblogs.com/jokerjason/p/10711022.html)

NOTE CoolShell CPU Cache [^CPUCache]


### 1.10. Modern CPUs are optimised for bulk operations 现代 CPU 已针对批量操作进行了优化

> Modern processors are a like nitro fuelled funny cars, they excel at the quarter mile. Unfortunately modern programming languages are like Monte Carlo, they are full of twists and turns. — David Ungar

> 现代处理器就像是硝基燃料的汽车，它们在四分之一英里处表现出色。不幸的是，现代编程语言就像蒙特卡洛一样，充满了曲折。—大卫·昂加（David Ungar）

NOTE [Monte Carlo 蒙特卡罗方法是一种计算方法。原理是通过大量随机样本，去了解一个系统，进而得到所要计算的值。](http://www.ruanyifeng.com/blog/2015/07/monte-carlo-method.html)

This a quote from David Ungar, an influential computer scientist and the developer of the SELF programming language that was referenced in a very old presentation [I found online](http://www.ai.mit.edu/projects/dynlangs/wizards-panels.html).

这句话引用自有影响力的计算机科学家，SELF编程语言的开发人员 David Ungar ，我在网上找到了一个很旧的演示文稿，并引用了该引用。

Thus, modern CPUs are optimised for bulk transfers and bulk operations. At every level, the setup cost of an operation encourages you to work in bulk. Some examples include

-   memory is not loaded per byte, but per multiple of cache lines, this is why alignment is becoming less of an issue than it was in earlier computers.
    
-   Vector instructions like MMX and SSE allow a single instruction to execute against multiple items of data concurrently providing your program can be expressed in that form.
    
现代 CPU 针对批量传输和运算操作进行了优化。
建议尽量把操作合并到一次调用来执行。

- 内存是按缓存行的倍数加载，不再像以前一样按字节加载。
- 类似 MMX SSE 等向量指令允许一条指令同时针对多个数据项并发执行。当然，这也需要上层程序支持。



### 1.11. Modern processors are limited by memory latency not memory capacity 现代处理器的主要瓶颈在内存延迟，而非内存大小

If the situation in CPU land wasn’t bad enough, the news from the memory side of the house doesn’t get much better.

如果 CPU 负载不是特别高，那内存延迟的影响也就没那么大。

Physical memory attached to a server has increased geometrically. My first computer in the 1980’s had kilobytes of memory. When I went through high school I wrote all my essays on a 386 with 1.8 megabytes of ram. Now its commonplace to find servers with tens or hundreds of gigabytes of ram, and the cloud providers are pushing into the terabytes of ram.

服务器上的物理内存已经程几何级数增长。
我在1980年的第一台电脑只有几千字节内存。
高中时所有论文都是在只有 1.8 MB 内存的 386 机器上编写。
现在很容易找到拥有几十上百 GB 内在的服务器，云服务提供商甚至使用了 TB 大小的内存。

> NOTE gemoetrically “几何级数”，又称“等比级数”。跟算法课程中的“复杂度量级”不同。

> 常量阶O(1) 对数阶O(logn) 线性阶O(n) 线性对数阶O(n logn) 平方阶O(n^2) 立方阶O(n^3) k次方阶O(n^k) 指数阶O(2^n) 阶乘阶O(n!) 。
> [数据结构与算法之美](https://time.geekbang.org/column/article/40036)


![mem gap](https://www.extremetech.com/wp-content/uploads/2018/01/mem_gap.png)

However, the gap between processor speeds and memory access time continues to grow.

处理器速度与内存访问延迟之间的差异越来越大。（上图中2000到2010年间，CPU速度提高10倍速，Memory访问延迟仅提高2倍）

[Table 2.2 Example Time Scale of System Latencies](https://pbs.twimg.com/media/BmBr2mwCIAAhJo1.png)

Event        | Latency | Scaled |
------------ | ---------------- |
1 CPU cycle  | 0.3 ns  | 1s     |
Level 1 cache access | 0.9 ns | 3s |
Level 2 cache access | 2.8 ns | 9s |
Level 3 cache access | 12.9 ns | 43 s |
Main memory access (DRAM, from CPU) | 120 ns | 6 min |
Solid-state disk I/O (flash memory) | S0- -150 us | 2- -6 days |
Rotational disk I/O | 1-10 ms | 1-12 months |
Internet: San Francisco to New York | 40 ms | 4 years |
Internet: San Francisco to United Kingdom | 81 ms | 8 years |
Internet: San Francisco to Australia | 183 ms | 19 years |
TCP packet retransmit | 1-3 s | 105- 317 years |
OS virtualization system reboot | 4S | 423 years |
SCSI command time-out | 30 s | 3 millennia |
Hardware (HW) virtualization system reboot | 40 s | 4 millennia |
Physical system reboot | 5m | 32 millennia |



But, in terms of processor cycles lost waiting for memory, physical memory is still as far away as ever because memory has not kept pace with the increases in CPU speed.

Memory 跟不上 CPU 速度快速增长的步伐，所以处理器要空转多个时钟周期来等待访问内存的过程。

So, most modern processors are limited by memory latency not capacity.

所以说，现代处理器主要瓶颈在内存延迟，而不是内存大小。



### 1.12. Cache rules everything around 至关重要的缓存

![latency](https://www.extremetech.com/wp-content/uploads/2014/08/latency.png)

> data access range  数据访问范围，大小

> memory latency 内存延迟


For decades the solution to the processor/memory cap was to add a cache-- a piece of small fast memory located closer, and now directly integrated onto, the CPU.

近几十年来，提升处理器/内存瓶颈的主要方法就是加缓存 －－ 最开始是在CPU附近增加一小块高速内存，现在直接把高速内存集成到CPU内了。

But;

-   L1 has been stuck at 32kb per core for decades
    
-   L2 has slowly crept up to 512kb on the largest intel parts
    
-   L3 is now measured in 4-32mb range, but its access time is variable


但是：

- 一级缓存一直保持在每核心 32kb 大小，几十年没什么变化

- 二级缓存增长缓慢，在 interl 中最大 512kb

- 三级缓存在 4-32mb 之间，但它的访问时间是变化的

> TODO 三级缓存访问时间可变，是指不同厂商使用的三级缓存速度不同？还是说不凬时间点访问三级缓存所需要的时间是变化的？
    

![E5v4blockdiagram](https://i3.wp.com/computing.llnl.gov/tutorials/linux_clusters/images/E5v4blockdiagram.png)

By caches are limited in size because they are  [physically large on the CPU die](http://www.itrs.net/Links/2000UpdateFinal/Design2000final.pdf), consume a lot of power. To halve the cache miss rate you must  _quadruple_  the cache size.

因为这块高速内存占用CPU的物理空间太多，且消费电量较高，所以缓存一直不大。
如果想让缓存丢失率减小一半，至少要让缓存大小增加四倍。




### 1.13. The free lunch is over 免费午餐的时代结束了

In 2005 Herb Sutter, the C++ committee leader, wrote an article entitled  [The free lunch is over](http://www.gotw.ca/publications/concurrency-ddj.htm). In his article Sutter discussed all the points I covered and asserted that future programmers will not longer be able to rely on faster hardware to fix slow programs or slow programming languages.

 C+＋ 委员会领导人 Herb Sutter 于 2005 年写过一篇名为《The free lunch is over》的文章。
 文中讨论了我刚才讲的所有知识点，并且断言，以后的程序员再也不能仅仅靠升级更快的硬件来给应用程序或编程语言提升性能了。


Now, more than a decade later, there is no doubt that Herb Sutter was right. Memory is slow, caches are too small, CPU clock speeds are going backwards, and the simple world of a single threaded CPU is long gone.

十几年后的如今，可以肯定 Herb Sutter 是正确的。
内存太慢，缓存又太小，CPU时钟速率还更慢了，单线程CPU的世界已经过去了。

Moore’s Law is still in effect, but for all of us in this room, the free lunch is over.

摩尔定律仍然有效，但对在座各位来说，免费午餐时代已经过去了。



### 1.14. Conclusion 结论

> The numbers I would cite would be by 2010: 30GHz, 10billion transistors, and 1 tera-instruction per second. — [Pat Gelsinger, Intel CTO, April 2002](https://www.cnet.com/news/intel-cto-chip-heat-becoming-critical-issue/)

> Intel CTO  Pat Gelsinger 在 2002 4月曾预测2010年的CPU性能指标： 30GHz, 100亿个晶体管，每秒种执行 一千亿 条指令。
>
> NOTE 引用的文章中，Pat Gelsinger 是说在没有散热问题的情况下，才能达到 30GHz 的频率。虽然 2004 年时已经生产出 3.8GHz 主频的 [Prescott](https://zh.wikipedia.org/wiki/%E5%A5%94%E8%85%BE4#cite_note-3) ，但至今 2020 年，单核心CPU主频普遍低于 4GHz [^InterlCPUList] 。
> 
> - 1k = 1000
> - 1m = 1000k = 100 万
> - 1g = 1000m = 1 亿
> - 1t = 1000g = 一千亿 
> 
> NOTE 其他相关文章 [Favorite Forecast Fallacies](https://semiengineering.com/favorite-forecast-fallacies/)  [请问目前主流CPU的每秒计算次数能达到多少？能够上亿吗？](https://www.zhihu.com/question/39604940)


It’s clear that without a breakthrough in material science the likelihood of a return to the days of 52% year on year growth in CPU performance is vanishingly small. The common consensus is that the fault lies not with the material science itself, but how the transistors are being used. The logical model of sequential instruction flow as expressed in silicon has lead to this expensive endgame.

显然，如果材料科学方面没有技术突破，想让 CPU 性能恢复到同比 52％ 增长的可能性很小。
普遍的共识是，问题不在于材料科学本身，而在于晶体管的使用方式。
只要使用硅表示的顺序指令流的逻辑模型，必然导致这种代价。

There are many presentations online that rehash this point. They all have the same prediction — computers in the future will not be programmed like they are today. Some argue it’ll look more like graphics cards with hundreds of very dumb, very incoherent processors. Others argue that Very Long Instruction Word (VLIW) computers will become predominant. All agree that our current sequential programming languages will not be compatible with these kinds of processors.

互联网上有很多资料强调过这一观点。
它们都预，未来的计算机将不会像今天这样编程。
有人认为未来的处理器将由上百个型号不一的低端处理器组成。也有人认为超长指令字(VLIW)将成为主流。
但大家一致同意，现在的顺序编程语言无法适应未来的处理器。

My view is that these predictions are correct, the outlook for hardware manufacturers saving us at this point is grim. However, there is  _enormous_  scope to optimise the programs today we write for the hardware we have today. Rick Hudson spoke at GopherCon 2015 about  [reengaging with a "virtuous cycle"](https://talks.golang.org/2015/go-gc.pdf)  of software that works  _with_  the hardware we have today, not indiferent of it.

我认为这些预测是对的，靠硬件厂商提升我们的软件性能是不太靠谱了。
但现有硬件上的程序还有很大优化空间。
Rick Hudson (里克·哈德森)在2015年GopherCon谈到“良性循环”的概念。
即硬件和软件应该应该互相配合，迭代优化升级。

> NOTE: 以前的硬件频率增长快，那软件应该充分发按高频的特性；以后的硬件 CPU 核心数量越来越多，那软件应该充分利用多核心的优势。


Looking at the graphs I showed earlier, from 2015 to 2018 with at best a 5-8% improvement in integer performance and less than that in memory latency, the Go team have decreased the garbage collector pause times by  [two orders of magnitude](https://blog.golang.org/ismmkeynote). A Go 1.11 program exhibits significantly better GC latency than the same program on the same hardware using Go 1.6. None of this came from hardware.

回顾之前的图标可看出，从 2015 年到 2018 年间， integer performance 性能仅提高了 5-8% ， memory latency 提高得更少。即使这种情况下， Go 开发团队仍然将 garbage collector 暂停时间提高了两个数量级。
同样的代码，在同样的硬件中，使用 Go 1.11 编译时其 GC latency 明显优于 Go 1.6 版本。
这些提升可不来自硬件。

![intgraph](https://dave.cheney.net/high-performance-go-workshop/images/int_graph.png)


So, for best performance on today’s hardware in today’s world, you need a programming language which:

为了在现今世界的硬件中获得更好的性能，你需要的编程语言应该是下面这样：

-   Is compiled, not interpreted, because interpreted programming languages interact poorly with CPU branch predictors and speculative execution.
    
-   You need a language which permits efficient code to be written, it needs to be able to talk about bits and bytes, and the length of an integer efficiently, rather than pretend every number is an ideal float.
    
-   You need a language which lets programmers talk about memory effectively, think structs vs java objects, because all that pointer chasing puts pressure on the CPU cache and cache misses burn hundreds of cycles.
    
-   A programming language that scales to multiple cores as performance of an application is determined by how efficiently it uses its cache and how efficiently it can parallelise work over multiple cores.

- 它应该是编译型语言，而非解释型语言，因为解释型语言没法发挥 CPU 的分支预测和乱序执行功能优势。

- 这种语言应该支持编写更高效率的代码，它应该能操作 bit 和 byte ，而且能区分 integer 和 float 数值类型，以便更高效率地处理 integer 

- 这种语言应该能让程序员讨论内存，思考 struct 与 java object 的区别，因为所有的 pointer chasing 都会给 CPU  cache 带来很大压力，而 cache miss 会消耗上百个时钟周期。

- 这种编程语言应该支持，通过增加CPU核心数量来提升程序性能。所以它要能高效地利用 cache ，并高效地利用多核心并行工作。
    

Obviously we’re here to talk about Go, and I believe that Go inherits many of the traits I just described.

我们来到这里讨论 Go 语言，显然是因为 Go 具备很多我刚才描述的那些特点。


#### 1.14.1. What does that mean for us?

> There are only three optimizations: Do less. Do it less often. Do it faster.
> 
> The largest gains come from 1, but we spend all our time on 3. — [Michael Fromberger](https://twitter.com/creachadair/status/1039602865831010305)

> 有三种优化手段：少做。再少做一点。做快点。

> 收益最大的是第一种，但我们把时间都花到第三种手段上了。  — Michael Fromberger


The point of this lecture was to illustrate that when you’re talking about the performance of a program or a system is entirely in the software. Waiting for faster hardware to save the day is a fool’s errand.

这次讲座的目的是想说明，当我们谈论软件或系统性能优化时，肯定是在说完全基于软件的优化手段。
妄想等硬件变快大幅提高软件性能的想法是愚蠢的。

But there is good news, there is a tonne of improvements we can make in software, and that is what we’re going to talk about today.

好消息是，软件上还有非常大的优化空间，这就是我们今天要讲的内容。


#### 1.14.2. Further reading 延伸阅读

-   [The Future of Microprocessors, Sophie Wilson](https://www.youtube.com/watch?v=zX4ZNfvw1cw)  JuliaCon 2018
    
-   [50 Years of Computer Architecture: From Mainframe CPUs to DNN TPUs, David Patterson](https://www.youtube.com/watch?v=HnniEPtNs-4)
    
-   [The Future of Computing, John Hennessy](https://web.stanford.edu/~hennessy/Future%20of%20Computing.pdf)
    
-   [The future of computing: a conversation with John Hennessy](https://www.youtube.com/watch?v=Azt8Nc-mtKM)  (Google I/O '18)
    

## 2. Benchmarking 基准测试

> Measure twice and cut once. — Ancient proverb

> 测量两次，切一次。  — 谚语

> NOTE 这是它的直接翻译。它的深刻含义是指，做事情要精心准备，特别是当你只有一次机会的时候。例如，当人们切割木头的时候，必须仔细测量尺寸，因为你只有一次机会，否则你最后的尺寸不是大了就是小了。所以，当我们提醒别人做事情要三思后行时，在英文中就说“Measure Twice, Cut Once”。充足的准备是成功所必需的。有准备未必成功，但是没准备，失败的可能性很大。[http://learn-english-writing.blogspot.com/](http://learn-english-writing.blogspot.com/2011/12/measure-twice-cut-once.html)


Before we attempt to improve the performance of a piece of code, first we must know its current performance.

在我们试图提高一段代码的性能时，必须先了解它当前的性能。

This section focuses on how to construct useful benchmarks using the Go testing framework, and gives practical tips for avoiding the pitfalls.

本节将重点介绍如何使用 Go 测试框架构建基准测试，并给出避坑实践指南。


### 2.1. Benchmarking ground rules 基准测试规则

Before you benchmark, you must have a stable environment to get repeatable results.

进行基准测试前，你必须有一个稳定的运行环境，才能得到可重复的结果。


-   The machine must be idle -- don’t profile on shared hardware, don’t browse the web while waiting for a long benchmark to run.
    
-   Watch out for power saving and thermal scaling. These are almost unavoidable on modern laptops.
    
-   Avoid virtual machines and shared cloud hosting; they can be too noisy for consistent measurements.
    

- 机器必须是空闲的 -- 不要使用公用的硬件环境，也不要在等待运行较长时间的基准测试过程浏览网页。

- 注意系统的节能配置和热力缩放。这些问题在现代笔记本中几乎无法避免。

- 避免使用虚拟机和公共的云主机；这些环境干扰因素太多，无法保证测量结果一致性。


If you can afford it, buy dedicated performance test hardware. Rack it, disable all the power management and thermal scaling, and never update the software on those machines. The last point is poor advice from a system adminstration point of view, but if a software update changes the way the kernel or library performs —-think the Spectre patches—- this will invalidate any previous benchmarking results.

如果有钱，就买专门用于性能测试的硬件。放机架上，禁用所有电源管理和热力缩放功能，并且永远不要升级这台机器的软件。
从系统管理员的角度来说，最后一条建议非常糟糕。但如果升级软件后，改变了系统内核或第三方库，那么在此之前所有基准测试结果都无效了，比如 Spectre 漏洞的补丁就会除低系统性能。

> NOTE Spectre 漏洞的修复补丁会降低系统性能，参考 [如何看待 2018 年 1 月 2 日爆出的 Intel CPU 设计漏洞？](https://www.zhihu.com/question/265012502/answer/288239171)
> 芯片微码更新不足以修复漏洞，必须修改系统或者购买新设计的 CPU。
目前 Linux 内核的解决方案是重新设计页表（KPTI 技术，前身为 KAISER）。之前普通程序和内核程序共用页表，靠 CPU 来阻止普通程序的越权访问。新方案让内核使用另外一个页表，而普通程序的页表中只保留一些必要的内核信息（例如调用内核的地址）。这个方案会导致每次普通程序和内核程序之间的切换（例如系统内核调用或者硬件中断）都需要切换页表，引起 CPU 的 TLB 缓存刷新。TLB 缓存刷新相对来说是非常耗时的，因此会降低系统的效率。
> KAISER 技术对系统性能的影响一般是 5%，最高可达 30%。一些高级的芯片功能（例如 PCID）可以支持其他技术，从而减少性能影响。Linux 已经在 4.14 版本的开发过程中添加了对 PCID 的支持。
> 在 Linux 系统中，KPTI 只有在英特尔芯片上才会启用，因此 AMD 芯片不受影响，且用户可以通过手动修改开关的方式关闭 KPTI 。  


For the rest of us, have a before and after sample and run them multiple times to get consistent results.

优化前，优化后，都要运行多次基准测试，来保证前后样本结果的一致性。


### 2.2. Using the testing package for benchmarking 使用 testing package 进行基准测试

The  `testing`  package has built in support for writing benchmarks. If we have a simple function like this:

`testing` package 专门编写基准测试的内置包。
如果我们有一个下面这样简单的函数要测试：

```go
func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	default:
		return Fib(n-1) + Fib(n-2)
	}
}
```

The we can use the  `testing`  package to write a  _benchmark_  for the function using this form.

可以像下面这样，用 `testing` package  编写 _基准测试_  测试刚才函数。

```go
func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // run the Fib function b.N times
	}
}
```

The benchmark function lives alongside your tests in a  `_test.go`  file.

这个基准测试函数一般放在以 `_test.go` 结尾的单元测试中。

Benchmarks are similar to tests, the only real difference is they take a  `*testing.B`  rather than a  `*testing.T`. Both of these types implement the  `testing.TB`  interface which provides crowd favorites like  `Errorf()`,  `Fatalf()`, and  `FailNow()`.

基准测试与单元测试很像，唯一区别是，基准测试函数参数是 `*testing.B`，而单元测试函数参数是 `*testing.T`。
两者都实现了 `testing.TB` 接口定义的 `Errorf()`,  `Fatalf()`, 和  `FailNow()` 方法。




#### 2.2.1. Running a package’s benchmarks 运行 package 的基准测试

As benchmarks use the  `testing`  package they are executed via the  `go test`  subcommand. However, by default when you invoke  `go test`, benchmarks are excluded.

因为 benchmark 使用 `testing` package ，自然是使用 `go test` 的子命令运行基准测试的。
但，执行 `go test` 命令时，默认不会执行基准测试。

To explicitly run benchmarks in a package use the  `-bench`  flag.  `-bench`  takes a regular expression that matches the names of the benchmarks you want to run, so the most common way to invoke all benchmarks in a package is  `-bench=.`. Here is an example:

使用 `-bench` 标志可明确运行 package 中的基准测试。
`-bench` 使用正则匹配要运行基准测试的 package 名称。
一般常用 `-bench=.` 运行 package 中的所有基准测试。
下面是示例：

```shell
% go test -bench=. ./examples/fib/
goos: darwin
goarch: amd64
BenchmarkFib20-8           30000             40865 ns/op
PASS
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     1.671s
```

`go test`  will also run all the tests in a package before matching benchmarks, so if you have a lot of tests in a package, or they take a long time to run, you can exclude them by providing  go test’s `-run`  flag with a regex that matches nothing; ie.

`go test` 在匹配基准测试前，会先运行 package 中的所有单元测试，如果 package 中的单元测试很多，或者运行单元测试需要很长时间，可以给 go test 的 `-run` 标记加一个匹配结果为空的参数，就能跳过单元测试。


```shell
go test -run=^$
```


#### 2.2.2. How benchmarks work 基准测试是如何工作的

Each benchmark function is called with different value for  `b.N`, this is the number of iterations the benchmark should run for.

`b.N` 参数是基准的迭代次数，每个基准测试函数都会使用不同的 `b.N` 值被调用多次。

`b.N`  starts at 1, if the benchmark function completes in under 1 second --the default-- then  `b.N`  is increased and the benchmark function run again.

`b.N` 默认取值是1，如果基准测试函数在1秒内完成，会增加 `b.N` 的值，然后再次运行基准测试。


`b.N`  increases in the approximate sequence; 1, 2, 3, 5, 10, 20, 30, 50, 100, and so on. The benchmark framework tries to be smart and if it sees small values of  `b.N`  are completing relatively quickly, it will increase the the iteration count faster.

`b.N` 的增加过程类似这个数列：1, 2, 3, 5, 10, 20, 30, 50, 100, 等等。
基准测试框架如果发现很小的 `b.N` 值能很快完成，会增加迭代 `b.N` 数时的取值(NOTE:跳过较小的数)。


Looking at the example above,  `BenchmarkFib20-8`  found that around 30,000 iterations of the loop took just over a second. From there the benchmark framework computed that the average time per operation was 40865ns.

从上面的例子可以看出 `BenchmarkFib20-8` 过程大概 30000 次迭代只需要 1 秒钟。
所以基准测试框架计算出平均每次操作是 40865ns 。

The  `-8`  suffix relates to the value of  `GOMAXPROCS`  that was used to run this test. This number, like  `GOMAXPROCS`, defaults to the number of CPUs visible to the Go process on startup. You can change this value with the  `-cpu`  flag which takes a list of values to run the benchmark with.

后缀 `-8` 的取值与运行测试时的 `GOMAXPROCS` 有关。
跟 `GOMAXPROCS` 环境变量一样，它的默认取值是 Go 进程启动时的 CPU 数量。
你能用 `-cpu` 参数改变运行基准测试时的取值

```shell
% go test -bench=. -cpu=1,2,4 ./examples/fib/
goos: darwin
goarch: amd64
BenchmarkFib20             30000             39115 ns/op
BenchmarkFib20-2           30000             39468 ns/op
BenchmarkFib20-4           50000             40728 ns/op
PASS
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     5.531s
```

This shows running the benchmark with 1, 2, and 4 cores. In this case the flag has little effect on the outcome because this benchmark is entirely sequential.

以上是分别使用 1，2 和 4 个CPU核心运行基准测试的结果。
因为被测函数是顺序执行的 (NOTE:Fib没有利用多核心优化)  ，所以这里的基准测试结果变化不大。


#### 2.2.3. Improving benchmark accuracy 提高基准测试的精度

The  `fib`  function is a slightly contrived example —-unless your writing a TechPower web server benchmark—- it’s unlikely your business is going to be gated on how quickly you can compute the 20th number in the Fibonaci sequence. But, the benchmark does provide a faithful example of a valid benchmark.

这个 `fib` 是人为的例子 --除非你写 TechPower 网络服务的基准测试-- 否则你的业务瓶颈不在可能限制在计算第 20 个斐波那契序列的值。
但这个示例对于演示基准测试足够了。

Specifically you want your benchmark to run for several tens of thousand iterations so you get a good average per operation. If your benchmark runs for only 100’s or 10’s of iterations, the average of those runs may have a high standard deviation. If your benchmark runs for millions or billions of iterations, the average may be very accurate, but subject to the vaguaries of code layout and alignment.

如果你的基准测试运行成千上万次，应该能得到一个较真实的平均操作时间。
如果你的基准测试只运行几十几百次，那么得到的平均操作时间应该有较大的误差。
如果你的基准测试运行几百万甚至几十亿次，结果会非常精确。但这时有可能受到代码布局和对齐的影响。

> NOTE 这里 vaguaries of code layout and alignment 可能是指数据对齐。数据对齐有可能对CPU性能产生影响，当数据结构大小刚才好 Cache Line 对齐，有可能提高性能。 [^CPUCache] [^MemoryAndNativeCodePerformance]


To increase the number of iterations, the benchmark time can be increased with the  `-benchtime`  flag. For example:

可以使用 `-benchtime` 参数增加基准测试时间，进而增加迭代次数。比如：

```shell
% go test -bench=. -benchtime=10s ./examples/fib/
goos: darwin
goarch: amd64
BenchmarkFib20-8          300000             39318 ns/op
PASS
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     20.066s
```

Ran the same benchmark until it reached a value of  `b.N`  that took longer than 10 seconds to return. As we’re running for 10x longer, the total number of iterations is 10x larger. The result hasn’t changed much, which is what we expected.

此时会选取一个使基准测试的运行时间至少超过 10s 的 `b.N` 值。 
我们的运行时间增加了10倍，运行次数也增加了10倍。
但平均操作时间的结果并没有变化，这也正是我们期望看到的结果。


Why is the total time reporteded to be 20 seconds, not 10?

但为什么总耗时是 20s 而不是 10s 呢？

If you have a benchmark which runs for millons or billions of iterations resulting in a time per operation in the micro or nano second range, you may find that your benchmark numbers are unstable because thermal scaling, memory locality, background processing, gc activity, etc.

如果你的基准测试会运行上百万甚至数亿次，而每次操作在微秒和纳秒范围内。
你会发现，基准测试结果会因为 thermal scaling, memory locality, background processing, gc 活动等变得十分不稳定。

For times measured in 10 or single digit nanoseconds per operation the relativistic effects of instruction reordering and code alignment will have an impact on your benchmark times.

对于单次操作耗时在 10纳秒以内的情况，基准测试受指令重排和代码对齐的影响很大。

To address this run benchmarks multiple times with the  `-count`  flag:

为解决这种问题，可以使用 `-count` 参数指定运行基准测试的数次。

```shell
% go test -bench=Fib1 -count=10 ./examples/fib/
goos: darwin
goarch: amd64
BenchmarkFib1-8         2000000000               1.99 ns/op
BenchmarkFib1-8         1000000000               1.95 ns/op
BenchmarkFib1-8         2000000000               1.99 ns/op
BenchmarkFib1-8         2000000000               1.97 ns/op
BenchmarkFib1-8         2000000000               1.99 ns/op
BenchmarkFib1-8         2000000000               1.96 ns/op
BenchmarkFib1-8         2000000000               1.99 ns/op
BenchmarkFib1-8         2000000000               2.01 ns/op
BenchmarkFib1-8         2000000000               1.99 ns/op
BenchmarkFib1-8         1000000000               2.00 ns/op
```

A benchmark of  `Fib(1)`  takes around 2 nano seconds with a variance of +/- 2%.

函数 `Fib(1)` 耗时大概2纳秒，方差为 +/- 2% 。 

New in Go 1.12 is the  `-benchtime`  flag now takes a number of iterations, eg.  `-benchtime=20x`  which will run your code exactly  `benchtime`  times.

Go 1.12 版本中 `-benchtime` 参数支持设置迭代的次数，比如 `-benchtime=20x` 可以让准确的让基准测试运行20次。

Try running the fib bench above with a  `-benchtime`  of 10x, 20x, 50x, 100x, and 300x. What do you see?

尝试使用 10x, 20x, 50x, 100x, 和 300x 为 `-benchtime` 分别运行基准测试，看看结果是什么样？

If you find that the defaults that  `go test`  applies need to be tweaked for a particular package, I suggest codifying those settings in a  `Makefile`  so everyone who wants to run your benchmarks can do so with the same settings.

如果你希望调整执行 `go test` 所用的默认参数，建议把这些配置写到 `Makefile` 中，以便所有人运行基准测试时，都使用相同的配置。


### 2.3. Comparing benchmarks with benchstat 使用 benchstat 比较基准测试

In the previous section I suggested running benchmarks more than once to get more data to average. This is good advice for any benchmark because of the effects of power management, background processes, and thermal management that I mentioned at the start of the chapter.

上一节中，我建议多运行几次基准测试，以便得到更准确的平均结果。
为防止 power management, background process 和 thermal management 的影响，这是个很好的建议。

I’m going to introduce a tool by Russ Cox called  [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat).

下面我要介绍的是 Russ Cox 的 benchstat 工具。

```shell
% go get golang.org/x/perf/cmd/benchstat
```

Benchstat can take a set of benchmark runs and tell you how stable they are. Here is an example of  `Fib(20)`  on battery power.

Benchstat 可以分析出一组基准测试结果的稳定性如何。
下面是在 battery power 上执行 `Fib(20)` 的基准测试结果。

```shell
% go test -bench=Fib20 -count=10 ./examples/fib/ | tee old.txt
goos: darwin
goarch: amd64
BenchmarkFib20-8           50000             38479 ns/op
BenchmarkFib20-8           50000             38303 ns/op
BenchmarkFib20-8           50000             38130 ns/op
BenchmarkFib20-8           50000             38636 ns/op
BenchmarkFib20-8           50000             38784 ns/op
BenchmarkFib20-8           50000             38310 ns/op
BenchmarkFib20-8           50000             38156 ns/op
BenchmarkFib20-8           50000             38291 ns/op
BenchmarkFib20-8           50000             38075 ns/op
BenchmarkFib20-8           50000             38705 ns/op
PASS
ok      _/Users/dfc/devel/high-performance-go-workshop/examples/fib     23.125s
% benchstat old.txt
name     time/op
Fib20-8  38.4µs ± 1%
```

`benchstat`  tells us the mean is 38.8 microseconds with a +/- 2% variation across the samples. This is pretty good for battery power.

`benchstat` 发现样本均值在 38.8ms ，方差 +/- 1% 。
这个结果对 battery power 来说很不错了。

> NOTE: 描述有误，以命令输出的内容为准，均值应该是 38.4µs ± 1%。 38.8 microseconds 比 38.4µs 大太多了。


-   The first run is the slowest of all because the operating system had the CPU clocked down to save power.
    
-   The next two runs are the fastest, because the operating system as decided that this isn’t a transient spike of work and it has boosted up the clock speed to get through the work as quick as possible in the hope of being able to go back to sleep.
    
-   The remaining runs are the operating system and the bios trading power consumption for heat production.

- 第一次运行结果是最慢的，因为操作系统为了省电，把 CPU 时钟速率降到最低了。

- 紧接下来的两次是最快的，因为操作系统发现这不是一个短暂的临时任务，为了尽快完成任务，回到睡眠状态，它调高了CPU时钟速率。

-  剩下的结果都是伴随操作系统与 BIOS 间协调能耗与散热的情况执行的。

> TODO: 后面这次结果时高时低，是想说明什么呢？ 
> 因为主频调高后，能耗高，热量高，导致降频，然后能耗变低，热量低，所以结果时高时低吗？ 但主频升高后，应该会维护一段时间才对。所以这里结果时高时低应该和能耗处理无关，
> 也许是运行多次后，由于每次运算都是重复的，导致更好得利用了 CacheLine 等特性，才产生时高时低的结果吗？没想通。

    

#### 2.3.1. Improve  `Fib` 改善 `Fib`

Determining the performance delta between two sets of benchmarks can be tedious and error prone. Benchstat can help us with this.

确定两组基准测试之间的性能偏差可能会很繁琐，而且容易出错。Benchstat 可以帮助我们解决这个问题。

Saving the output from a benchmark run is useful, but you can also save the  _binary_  that produced it. This lets you rerun benchmark previous iterations. To do this, use the  `-c`  flag to save the test binary。I often rename this binary from  `.test`  to  `.golden`.

将基准测试的输出结果保存起来也许会很有用的，但也可以保存产生测试结果的二进制可执行文件。
这可以让你重新运行之前的基准测试。
要做到这一点，使用 `-c` 标志来保存产生测试结果的二进制可执行文件。
我经常把这个二进制文件从 `.test` 重命名为 `.golden` 。

> NOTE golden 珍贵

```shell
% go test -c
% mv fib.test fib.golden
```

The previous  `Fib`  fuction had hard coded values for the 0th and 1st numbers in the fibonaci series. After that the code calls itself recursively. We’ll talk about the cost of recursion later today, but for the moment, assume it has a cost, especially as our algorithm uses exponential time.

前面的 Fib 函数硬编码了 fibonacci 中的第0和第1个数字。之后的代码会进行递归调用。
我们稍后会讲到递归的成本，暂时先假设它是有成本的，而且时间复杂度是指数阶的。

As simple fix to this would be to hard code another number from the fibonacci series, reducing the depth of each recusive call by one.

最简单的优化方法是，直接硬编码其他几个 fibonacci 数量，直接减少递归的尝试。

```go
func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}
```

This file also includes a comprehensive test for  `Fib`. Don’t try to improve your benchmarks without a test that verifies the current behaviour.

代码中还有 `Fib` 相关的基准测试。
在没有验证过改进后的函数代码前，不要调整基准测试的代码。

To compare our new version, we compile a new test binary and benchmark both of them and use  `benchstat`  to compare the outputs.

为了比较我们新版本的函数，可以分别编译两个基准测试的二进制文件，并使用 benchstat 来比较两个基准测试的输出结果。

```shell
% go test -c
% ./fib.golden -test.bench=. -test.count=10 > old.txt
% ./fib.test -test.bench=. -test.count=10 > new.txt
% benchstat old.txt new.txt
name     old time/op  new time/op  delta
Fib20-8  44.3µs ± 6%  25.6µs ± 2%  -42.31%  (p=0.000 n=10+10)
```

There are three things to check when comparing benchmarks

比较基准测试结果时，主要关注这三点

-   The variance ± in the old and new times. 1-2% is good, 3-5% is ok, greater than 5% and some of your samples will be considered unreliable. Be careful when comparing benchmarks where one side has a high variance, you may not be seeing an improvement.

- 新旧结果中各自的方差值。最好的方差是 1-2% ，其次是 3-5% ， 超过 5% 的方差结果，说明这个样本结果不可信。无论新旧基准测试结果中哪一个方差值偏高，这次优化提升的结果(delta)都是不准确的了。
    
-   p value. p values lower than 0.05 are good, greater than 0.05 means the benchmark may not be statistically significant.

- p 值。低于 0.05 的 p 值是有意义的，只要基准测试结果 p 值超过 0.05 ，都没有意义。
    
-   Missing samples. benchstat will report how many of the old and new samples it considered to be valid, sometimes you may find only, say, 9 reported, even though you did  `-count=10`. A 10% or lower rejection rate is ok, higher than 10% may indicate your setup is unstable and you may be comparing too few samples.

- 样本数量不足。 benchstat 会报告新旧基准测试结果中它认为有效的样本数量，有时你会发现，即使指定了 `-count=10` 参数，却仅显示 9 个样本数量。缺少的样本数量在 10% 以下，都是可以接受的。缺失的样本数量大于 10% ，可能就是你参数配置不正确。
    

### 2.4. Avoiding benchmarking start up costs 减少基准测试的启动成本

Sometimes your benchmark has a once per run setup cost.  `b.ResetTimer()`  will can be used to ignore the time accrued in setup.

有时，在运行基准测试前，可能需要执行一些比较耗时的初始化配置。
使用 `b.ResetTimer()` 可以忽略这些初始化配置浪费的时间。

```go
func BenchmarkExpensive(b *testing.B) {
        boringAndExpensiveSetup()
        b.ResetTimer() 
        for n := 0; n < b.N; n++ {
                // function under test
        }
}
```

Reset the benchmark timer

重置基准测试定时器

If you have some expensive setup logic  _per loop_  iteration, use  `b.StopTimer()`  and  `b.StartTimer()`  to pause the benchmark timer.

如果在循环迭代 `b.N` 的过程每次都要执行一些耗时的操作，可以搭配使用 `b.StopTimer()`  和 `b.StartTimer()`  暂停基准测试过程的计时器。

```go
func BenchmarkComplicated(b *testing.B) {
        for n := 0; n < b.N; n++ {
                b.StopTimer() 
                complicatedSetup()
                b.StartTimer() 
                // function under test
        }
}
```

Pause benchmark timer

暂停计时器使用 `b.StopTimer()` 

Resume timer

恢复计时器使用 `b.StartTimer()` 



### 2.5. Benchmarking allocations 基准测试过程的内存分配

Allocation count and size is strongly correlated with benchmark time. You can tell the  `testing`  framework to record the number of allocations made by code under test.

内存分配的次数和大小与基准测试的耗时密切相关。
你让 `testing` 框架记录记录被测试代码执行内存分配的次数。

```go
func BenchmarkRead(b *testing.B) {
        b.ReportAllocs()
        for n := 0; n < b.N; n++ {
                // function under test
        }
}
```

Here is an example using the  `bufio`  package’s benchmarks.

以下示例是运行标准库中 `bufio` package 的基准测试的结果。

```shell
% go test -run=^$ -bench=. bufio
goos: darwin
goarch: amd64
pkg: bufio
BenchmarkReaderCopyOptimal-8            20000000               103 ns/op
BenchmarkReaderCopyUnoptimal-8          10000000               159 ns/op
BenchmarkReaderCopyNoWriteTo-8            500000              3644 ns/op
BenchmarkReaderWriteToOptimal-8          5000000               344 ns/op
BenchmarkWriterCopyOptimal-8            20000000                98.6 ns/op
BenchmarkWriterCopyUnoptimal-8          10000000               131 ns/op
BenchmarkWriterCopyNoReadFrom-8           300000              3955 ns/op
BenchmarkReaderEmpty-8                   2000000               789 ns/op            4224 B/op          3 allocs/op
BenchmarkWriterEmpty-8                   2000000               683 ns/op            4096 B/op          1 allocs/op
BenchmarkWriterFlush-8                  100000000               17.0 ns/op             0 B/op          0 allocs/op
```

You can also use the  `go test -benchmem`  flag to force the testing framework to report allocation statistics for all benchmarks run.

还可以用 `go test -benchmem` 标志强制  testing 框架记录基准测试过程的内存分配情况。

```shell
% go test -run=^$ -bench=. -benchmem bufio
goos: darwin
goarch: amd64
pkg: bufio
BenchmarkReaderCopyOptimal-8            20000000                93.5 ns/op            16 B/op          1 allocs/op
BenchmarkReaderCopyUnoptimal-8          10000000               155 ns/op              32 B/op          2 allocs/op
BenchmarkReaderCopyNoWriteTo-8            500000              3238 ns/op           32800 B/op          3 allocs/op
BenchmarkReaderWriteToOptimal-8          5000000               335 ns/op              16 B/op          1 allocs/op
BenchmarkWriterCopyOptimal-8            20000000                96.7 ns/op            16 B/op          1 allocs/op
BenchmarkWriterCopyUnoptimal-8          10000000               124 ns/op              32 B/op          2 allocs/op
BenchmarkWriterCopyNoReadFrom-8           500000              3219 ns/op           32800 B/op          3 allocs/op
BenchmarkReaderEmpty-8                   2000000               748 ns/op            4224 B/op          3 allocs/op
BenchmarkWriterEmpty-8                   2000000               662 ns/op            4096 B/op          1 allocs/op
BenchmarkWriterFlush-8                  100000000               16.9 ns/op             0 B/op          0 allocs/op
PASS
ok      bufio   20.366s
```



### 2.6. Watch out for compiler optimisations 小心编译器的优化

This example comes from  [issue 14813](https://github.com/golang/go/issues/14813#issue-140603392).

这下例子来自 [issue 14813](https://github.com/golang/go/issues/14813#issue-140603392) 。

```go
const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

func BenchmarkPopcnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt(uint64(i))
	}
}
```

How fast do you think this function will benchmark? Let’s find out.

你觉得这个函数的基准测试结果有多快？我们看看结果吧。

```shell
% go test -bench=. ./examples/popcnt/
goos: darwin
goarch: amd64
BenchmarkPopcnt-8       2000000000               0.30 ns/op
PASS
```

0.3 of a nano second; that’s basically one clock cycle. Even assuming that the CPU may have a few instructions in flight per clock tick, this number seems unreasonably low. What happened?

只用了 0.3 纳秒，也就是只有一个时钟周期。
即使 CPU 在一个时钟滴答内能执行多个指令，这个结果也太小了。
到底为什么会这样呢？

> CPU时钟周期耗时，内存访问耗时等，可参考本文 1.11 节 Table 2.2 Example Time Scale of System Latencies

> [Concept of clock tick and clock cycles](https://stackoverflow.com/questions/25743995/concept-of-clock-tick-and-clock-cycles)
> 
> clock tick 时钟滴答 指系统时钟，是系统能识别的最小时间单位。
> 
> clock cycle 时钟周期 则是 CPU 执行一次完整的处理器脉冲所花费的时间。这是能从 CPU 主频计算出来的。比如 2GHz 的处理器，每秒钟能执行 2,000,000,000 clock cycles 。 


To understand what happened, we have to look at the function under benchmake,  `popcnt`.  `popcnt`  is a leaf function — it does not call any other functions — so the compiler can inline it.

要了解原因，我们得看看基准测试时 `popcnt` 到底做了什么。
`popcnt`是叶子函数 - 它不调用任意子函数 - 所以编译器会内联此函数。

Because the function is inlined, the compiler now can see it has no side effects.  `popcnt`  does not affect the state of any global variable. Thus, the call is eliminated. This is what the compiler sees:

因为这个函数是内联的，而编译器发现它没有任何副作用。`popcnt`没有修改任何全局变量的值。因此这个调用被省略了。
编译器看到的代码实际是下面这样。

```go
func BenchmarkPopcnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// optimised away
	}
}
```

On all versions of the Go compiler that i’ve tested, the loop is still generated. But Intel CPUs are really good at optimising loops, especially empty ones.

虽然我所测试过的所有 Go 编译器中，都会产生循环代码。但 Intel CPU 太擅长优化循环语言了，尤其是空循环。



#### 2.6.1. Exercise, look at the assembly 练习，看看汇编语言 

Before we go on, lets look at the assembly to confirm what we saw

继续讲其他内容前，我们先看看汇编语言确认下刚才的判断。

```go
% go test -gcflags=-S
```

Use `gcflags="-l -S"` to disable inlining, how does that affect the assembly output

可以使用 `gcflags="-l -S"` 关闭内联，这会影响编译出的汇编代码。

> Optimisation is a good thing
> 
> The thing to take away is the same optimisations that  _make real code fast_, by removing unnecessary computation, are the same ones that remove benchmarks that have no observable side effects.
> 
> This is only going to get more common as the Go compiler improves.


> 优化是一件好事
> 
> 优化是为了 _更快执行真正有用的代码_ 。
> 移除无用的计算是一种优化；像基准测试这样，移除没有副作用的代码也是一种优化。
>
> 随着Go编译器的改进，这种优化会更加常见。



#### 2.6.2. Fixing the benchmark 修复基准测试

Disabling inlining to make the benchmark work is unrealistic; we want to build our code with optimisations on.

在基准测试中关闭内联不太现实；因为我们编译代码时，肯定希望能启用内联功能优化代码。

> NOTE 如果基准测试关内联，但编译代码时又开内联，那基准测试结果就没有参考价值了。


To fix this benchmark we must ensure that the compiler cannot  _prove_  that the body of  `BenchmarkPopcnt`  does not cause global state to change.

为了修复这一现象，我们只需要让编译器无法 _证明_ `BenchmarkPopcnt` 没有修改全局变量即可。

```go
var Result uint64

func BenchmarkPopcnt(b *testing.B) {
	var r uint64
	for i := 0; i < b.N; i++ {
		r = popcnt(uint64(i))
	}
	Result = r
}
```

This is the recommended way to ensure the compiler cannot optimise away body of the loop.

这样就能保证编译器不会优化循环体内的代码了。


First we  _use_  the result of calling  `popcnt`  by storing it in  `r`. Second, because  `r`  is declared locally inside the scope of  `BenchmarkPopcnt`  once the benchmark is over, the result of  `r`  is never visible to another part of the program, so as the final act we assign the value of  `r`  to the package public variable  `Result`.

首先，我们把调用 `popcnt` 返回的结果保存在变量 `r` 中。
因为`r`是`BenchmarkPopcnt`是局部变量，所以一旦基准测试完毕，变量`r`的值对程序内其他代码都不可见，因此，我们还要把`r`的值分配到 package 公开变量 `Result` 。

Because  `Result`  is public the compiler cannot prove that another package importing this one will not be able to see the value of  `Result`  changing over time, hence it cannot optimise away any of the operations leading to its assignment.

因为`Result`是公开变量，所以编译器无法判断其他 package 何时会在导入当前 package 后访问`Result`的取值，因此编译器不会随意去除公共变量的赋值语句进行优化的。


What happens if we assign to  `Result`  directly? Does this affect the benchmark time? What about if we assign the result of  `popcnt`  to  `_`?

如果我们使用局部变量`r`，直接赋值给`Result`会发生什么呢？这会影响基准测试的时间吗？
如果我们把 `popcnt` 的结果赋值给 `_`  又会怎样呢？

In our earlier  `Fib`  benchmark we didn’t take these precautions, should we have done so?

在以前的 `Fib` 基准测试中，我们没有防范这些情况，那我们是否应该考虑并防范这些问题呢？


### 2.7. Benchmark mistakes 错误的基准测试

The  `for`  loop is crucial to the operation of the benchmark.

`for`循环是基准测试的关系部分。

Here are two incorrect benchmarks, can you explain what is wrong with them?

下面两个基准测试中，你能解释一下他们分别错在哪里吗？

```go
func BenchmarkFibWrong(b *testing.B) {
	Fib(b.N)
}
```

```go
func BenchmarkFibWrong2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(n)
	}
}
```

TODO 验证：1.应该是没有修改全局状态，所以被优化了。 2. b.N 表示执行 N 次函数，而非调用 Fib 时传参数为N 。

Run these benchmarks, what do you see?

运行这些基准测试，你看到了什么样的结果？



### 2.8. Profiling benchmarks 带性能分析的基准测试

> Profiling 是指在程序执行过程中，收集能够反映程序执行状态的数据。在软件工程中，性能分析（performance analysis，也称为 profiling），是以收集程序运行时信息为手段研究程序行为的分析方法，是一种动态程序分析的方法。[^qcraoPPROF]

The  `testing`  package has built in support for generating CPU, memory, and block profiles.

`testing` package 支持生成 CPU memory block profile 等类型的性能分析文件。

-   `-cpuprofile=$FILE`  writes a CPU profile to  `$FILE`.
    
-   `-memprofile=$FILE`, writes a memory profile to  `$FILE`,  `-memprofilerate=N`  adjusts the profile rate to  `1/N`.  设置内存采样速率。 [golang.org MemProfileRate](https://golang.org/pkg/runtime/#pkg-variables)
    
-   `-blockprofile=$FILE`, writes a block profile to  `$FILE`.
    

Using any of these flags also preserves the binary.

```
% go test -run=XXX -bench=. -cpuprofile=c.p bytes
% go tool pprof c.p
```

### 2.9. Discussion 讨论

Are there any questions?

Perhaps it is time for a break.



## 3. Performance measurement and profiling 性能评估与分析

In the previous section we looked at benchmarking individual functions which is useful when you know ahead of time where the bottlekneck is. However, often you will find yourself in the position of asking

> Why is this program taking so long to run?


上一节介绍的基准测试用于分析单个函数的性能瓶颈。
但我们经常遇到的问题的

> 为什么这个程序要运行这么长时间？


Profiling  _whole_  programs which is useful for answering high level questions like. In this section we’ll use profiling tools built into Go to investigate the operation of the program from the inside.

只有对 _整个_ 程序进行分析，才能回答这个问题。
这一节，我们使用 Go 内置的 profile tool 来调整程序内部的运行情况。



### 3.1. pprof

The first tool we’re going to be talking about today is  _pprof_.  [pprof](https://github.com/google/pprof)  descends from the  [Google Perf Tools](https://github.com/gperftools/gperftools)  suite of tools and has been integrated into the Go runtime since the earliest public releases.

`pprof`  consists of two parts:

-   `runtime/pprof`  package built into every Go program
    
-   `go tool pprof`  for investigating profiles.


第一个要讨论的工具是 _pprof_ 。  [pprof](https://github.com/google/pprof)  来自 [Google Perf Tools](https://github.com/gperftools/gperftools) ，在首次公开发版时，就包集成到了 Go 运行时内。

`pprof` 包含以下两部分：

-   `runtime/pprof`  是一个可在 Go 程序代码中引用 package 。
    
-   `go tool pprof`  是进行性能分析的工具。



### 3.2. Types of profiles 性能分析的种类

pprof supports several types of profiling, we’ll discuss three of these today:

-   CPU profiling.
    
-   Memory profiling.
    
-   Block (or blocking) profiling.
    
-   Mutex contention profiling. 
    
pprof 支持分析以下几种类型：

- CPU 性能剖析
- 内存性能剖析
- 阻塞剖析
- 锁（互斥量）争用剖析




#### 3.2.1. CPU profiling [^qcraoPPROF]

CPU profiling is the most common type of profile, and the most obvious.

When CPU profiling is enabled the runtime will interrupt itself every 10ms and record the stack trace of the currently running goroutines.

Once the profile is complete we can analyse it to determine the hottest code paths.

The more times a function appears in the profile, the more time that code path is taking as a percentage of the total runtime.

最常用的是 CPU profile。

启用 CPU profile 后，runtime 每隔 10ms 中断一次，然后记录当前的堆栈。

profile 完毕后，就能 analyse（分析） 出 hottest code path （热点代码）。

在 profile 中出现次数越多的函数，在 code path 中所占比重就最大。




#### 3.2.2. Memory profiling [^qcraoPPROF]

Memory profiling records the stack trace when a  _heap_  allocation is made.

Stack allocations are assumed to be free and are  _not tracked_  in the memory profile.

Memory profiling, like CPU profiling is sample based, by default memory profiling samples 1 in every 1000 allocations. This rate can be changed.

Because of memory profiling is sample based and because it tracks  _allocations_  not  _use_, using memory profiling to determine your application’s overall memory usage is difficult.

> Personal Opinion:  I do not find memory profiling useful for finding memory leaks. There are better ways to determine how much memory your application is using. We will discuss these later in the presentation.



Memory profile 是在堆(Heap)分配的时候，记录一下调用堆栈。

栈(Stack)分配由于会随时释放，因此不会被内存分析所记录。

与 CPU profile 类似，默认情况下，Memory profile 每 1000 次分配，取样一次，这个数值可以改变。

由于内存分析是取样方式，并且也因为其记录的是分配的内存，而不是使用的内存。因此使用内存性能分析工具来准确判断程序具体的内存使用是比较困难的。

> 个人观点:  Memory profile 不能用于查找内存泄漏。有其他更好的方法来跟踪程序占用的内存大小。我们后面讨论。



#### 3.2.3. Block profiling

Block profiling is quite unique to Go.

A block profile is similar to a CPU profile, but it records the amount of time a goroutine spent waiting for a shared resource.

This can be useful for determining  _concurrency_  bottlenecks in your application.

Block profiling can show you when a large number of goroutines  _could_  make progress, but were  _blocked_. Blocking includes:

-   Sending or receiving on a unbuffered channel.
    
-   Sending to a full channel, receiving from an empty one.
    
-   Trying to  `Lock`  a  `sync.Mutex`  that is locked by another goroutine.

Block profiling is a very specialised tool, it should not be used until you believe you have eliminated all your CPU and memory usage bottlenecks.


Block profile 是 Go 语言中特有的一种分析方法。

Block profile 与 CPU profile 很像，但它记录的是 goroutine 等待共享资源所花费的时间。

在分析程序 _并发_ 瓶颈时，十分有用。

Block profile 可分析出哪些时间，出现了大量 goroutine 同时处于 block 状态的情况。

可能引起 block 的原因如下：

- 发送或接收无缓冲的 channel 时。
- 向已满的 channel 中写数据，或从空的 channel 中读数据时。
- 尝试 Lock 一个已经被其他 goroutine 锁住的 sync.Mutex 时。

Block profile 很特殊，在排除 CPU 和 Memory 的性能瓶颈前，不要使用它来分析。



#### 3.2.4. Mutex profiling

Mutex profiling is simlar to Block profiling, but is focused exclusively on operations that lead to delays caused by mutex contention.

I don’t have a lot of experience with this type of profile but I have built an example to demonstrate it. We’ll look at that example shortly.

Mutex profile 与 Block profile 类似，但它专门分析互斥量争用所导致的延迟。

我没有多少 Mutext profile 相关的使用经验，但后面会有一个示例演示具体用法。



### 3.3. One profile at at time 每次只分析一种类型

Profiling is not free.

profile 是有代价的。

Profiling has a moderate, but measurable impact on program performance—especially if you increase the memory profile sample rate.

执行 profile 的过程对程序有一定性能损耗，特别是在提高 memory profile 采样率时。

Most tools will not stop you from enabling multiple profiles at once.

但是多数工具都不禁止你同时开启多个 profile 。

Do not enable more than one kind of profile at a time.

If you enable multiple profile’s at the same time, they will observe their own interactions and throw off your results.

同时开启多种 profile 时，在分析结果会互相影响。


### 3.4. Collecting a profile 收集分析结果

The Go runtime’s profiling interface lives in the  `runtime/pprof`  package.  `runtime/pprof`  is a very low level tool, and for historic reasons the interfaces to the different kinds of profile are not uniform.

Go 运行时 profile 接口位于 `runtime/pprof` package 中。
这是一个很低层的接口。
由于历史原因，不同类型的 profile 接口也不统一。

As we saw in the previous section, pprof profiling is built into the  `testing`  package, but sometimes its inconvenient, or difficult, to place the code you want to profile in the context of at  `testing.B`  benchmark and must use the  `runtime/pprof`  API directly.

pprof profile 是内置在 `testing` package 中的，如果不方便在 `testing.B` 基准测试中放置 profile 代码时，可以直接调用 `runtime/pprof` API 。


A few years ago I wrote a [pkg/profile](https://github.com/pkg/profile) package, to make it easier to profile an existing application.

几年前，我实现了一个生成 profile 的 package [pkg/profile](https://github.com/pkg/profile) 。

```go
import "github.com/pkg/profile"

func main() {
	defer profile.Start().Stop()
	// ...
}
```

We’ll use the profile package throughout this section. Later in the day we’ll touch on using the  `runtime/pprof`  interface directly.

本节我们就会用这个 pkg/profile package 进行演示。后面几天，才会直接使用 `runtime/pprof` 接口。




### 3.5. Analysing a profile with pprof 使用 pprof 分析剖析结果 

Now that we’ve talked about what pprof can measure, and how to generate a profile, let’s talk about how to use pprof to analyse a profile.

刚才讨论过，pprof 能剖析哪些内容，以有如何生成 profile 文件。
现在我们看看如何用 pprof 分析 profile 文件吧。

The analysis is driven by the  `go pprof`  subcommand

go tool pprof /path/to/your/profile

This tool provides several different representations of the profiling data; textual, graphical, even flame graphs.

此工具可以：文本，图形，甚至火焰图等几种方式展现 profile 数据。


If you’ve been using Go for a while, you might have been told that  `pprof`  takes two arguments. Since Go 1.9 the profile file contains all the information needed to render the profile. You do no longer need the binary which produced the profile. 🎉

如果你用过较早版本的 Go，可能遇到 `pprof` 命令要求提供两个参数的情况。即同时提供  profile 文件 与 生成 profile 时运行的二进制程序文件，才能输出分析结果。
从 Go 1.9 版本开始，只需要一个 profile 文件中就能执行分析，并输出分析结果。



#### 3.5.1. Further reading 延伸阅读

-   [Profiling Go programs](http://blog.golang.org/profiling-go-programs)  (Go Blog)
    
-   [Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs)
 


#### 3.5.2. CPU profiling (exercise)

Let’s write a program to count words:

写一个计算单词个数的程序：

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	defer profile.Start().Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	words := 0
	inword := false
	for {
		r, err := readbyte(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}
```

Let’s see how many words there are in Herman Melville’s classic  [Moby Dick](https://www.gutenberg.org/ebooks/2701)  (sourced from Project Gutenberg)

我们看看 赫尔曼·梅尔维尔经典小说 《Moby Dick白鲸记》 有多少个单词吧。

```txt
% go build && time ./words moby.txt
"moby.txt": 181275 words

real    0m2.110s
user    0m1.264s
sys     0m0.944s
```

Let’s compare that to unix’s  `wc -w`

再与 unix 中的标准单词计数程序 `wc -w` 比较一下

```txt
% time wc -w moby.txt
215829 moby.txt

real    0m0.012s
user    0m0.009s
sys     0m0.002s
```

So the numbers aren’t the same.  `wc`  is about 19% higher because what it considers a word is different to what my simple program does. That’s not important - both programs take the whole file as input and in a single pass count the number of transitions from word to non word.

Let’s investigate why these programs have different run times using pprof.

`wc` 命令算出的单词数量比我们的多了 19% ， 这是因为两个程序识别单词的标准不一样导致。
这个单独并不重启，我们主要关注的问题是：两个程序都从文件中读取所有数据然后计算单词数，为什么我们的程序花费的时间更长呢？

我们用 pprof 工具分析一下看看。




#### 3.5.3. Add CPU profiling 增加 CPU profile

First, edit  `main.go`  and enable profiling

编辑 `main.go` 文件，启用 profile 

```go
import (
        "github.com/pkg/profile"
)

func main() {
        defer profile.Start().Stop()
        // ...
```

Now when we run the program a  `cpu.pprof`  file is created.

现在，我们运行程序，一个 `cpu.pprof` 文件就会自动生成 。


```txt
% go run main.go moby.txt
2018/08/25 14:09:01 profile: cpu profiling enabled, /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile239941020/cpu.pprof
"moby.txt": 181275 words
2018/08/25 14:09:03 profile: cpu profiling disabled, /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile239941020/cpu.pprof
```

Now we have the profile we can analyse it with  `go tool pprof`

现在我们就能用 `go tool pprof` 工具直接分析这个 profile 文件了。

```txt
% go tool pprof /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile239941020/cpu.pprof
Type: cpu
Time: Aug 25, 2018 at 2:09pm (AEST)
Duration: 2.05s, Total samples = 1.36s (66.29%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1.42s, 100% of 1.42s total
      flat  flat%   sum%        cum   cum%
     1.41s 99.30% 99.30%      1.41s 99.30%  syscall.Syscall
     0.01s   0.7%   100%      1.42s   100%  main.readbyte
         0     0%   100%      1.41s 99.30%  internal/poll.(*FD).Read
         0     0%   100%      1.42s   100%  main.main
         0     0%   100%      1.41s 99.30%  os.(*File).Read
         0     0%   100%      1.41s 99.30%  os.(*File).read
         0     0%   100%      1.42s   100%  runtime.main
         0     0%   100%      1.41s 99.30%  syscall.Read
         0     0%   100%      1.41s 99.30%  syscall.read
```

The  `top`  command is one you’ll use the most. We can see that 99% of the time this program spends in  `syscall.Syscall`, and a small part in  `main.readbyte`.

最常用的命令是 `top` 。
这个程序 99% 的时间用花在 `syscall.Syscall` 上，
另一部分时间花在了 `main.readbyte` 。


We can also visualise this call the with the  `web`  command. This will generate a directed graph from the profile data. Under the hood this uses the  `dot`  command from Graphviz.

我们还可以用 `web` 命令观察剖析结果。这会根据 profile 中的数据生成一张图表。命令内部会调用 Graphviz 工具的 `dot` 命令生成的图表（译：所以在执行 pprof 分析的主机中，要安装 Graphviz 相关命令）。


However, in Go 1.10 (possibly 1.11) Go ships with a version of pprof that natively supports a http sever

但是，在 Go 1.10 (也可能是 1.11）版本中， pprof 内置了一个 HTTP 服务器。


```txt
% go tool pprof -http=:8080 /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile239941020/cpu.pprof
```

Will open a web browser;

-   Graph mode
    
-   Flame graph mode

执行上面的命令，就会自动打开浏览器，支持

- 图表模式
- 火焰图模式
    

On the graph the box that consumes the  _most_  CPU time is the largest — we see  `syscall.Syscall`  at 99.3% of the total time spent in the program. The string of boxes leading to  `syscall.Syscall`  represent the immediate callers — there can be more than one if multiple code paths converge on the same function. The size of the arrow represents how much time was spent in children of a box, we see that from  `main.readbyte`  onwards they account for near 0 of the 1.41 second spent in this arm of the graph.

在图表中，方框最大的图形占用CPU也最多。
方框上的连线，表示此函数的调用方。如果有多条代码路径调用相同的函数，就会出现多条线。
箭头的大小表示被调用方（子方框）花费的时间。

在本图中，可以观察到：
占用程序 99.3% 时间的是`syscall.Syscall` 。
`main.readbyte`方框占用了 1.41 秒，但`readbyte`函数本身只占用了 0 秒，其子函数 `File.Read() 占用了 1.41 秒。


_Question_: Can anyone guess why our version is so much slower than  `wc`?

_问题_： 谁知道为什么我们的程序比 `wc` 慢了这么多呢？


> 深度解密Go语言之pprof[^qcraoPPROF]

 列名 | 含义
----- | -------------------------------------------
flat  | 本函数的执行耗时, the time in a function
flat% | flat 占 CPU 总时间的比例。程序总耗时 16.22s, Eat 的 16.19s 占了 99.82%
sum%  | 前面每一行的 flat 占比总和
cum   | 累计量。指该函数加上该函数调用的函数总耗时。cumulative time a function and everything below it.
cum%  | cum 占 CPU 总时间的比例

> 方框中文字的含义 syscall.Syscall 760ms(84.44%) of 820ms(91.11%) 
> 
> 760ms 表示 flat 时间； 820ms 表示 cumulate 时间；



#### 3.5.4. Improving our version 优化我们的程序

The reason our program is slow is not because Go’s  `syscall.Syscall`  is slow. It is because syscalls in general are expensive operations (and getting more expensive as more Spectre family vulnerabilities are discovered).

Each call to  `readbyte`  results in a syscall.Read with a buffer size of 1. So the number of syscalls executed by our program is equal to the size of the input. We can see that in the pprof graph that reading the input dominates everything else.

我们的程序慢，并非由于 Go 的 `syscall.Syscall` 接口慢导致的。
而是因为 syscall 系统调用原本就是开销巨大的一种操作（为了修复越来越多的安全漏洞，这种开销也会越来越大）。

每次 `readbyte` 都会触发一次 syscall.Read 系统调用，而且 buffer size 是 1 ，所以系统调用的次数就是文件在字节数。
所以能从 pprof 图表中看到，主要耗时都在读取数据的过程。

> NOTE Spectre family vulnerabilities 幽灵是一个存在于分支预测实现中的硬件缺陷及安全漏洞，含有预测执行功能的现代微处理器均受其影响，漏洞利用是基于时间的旁路攻击，允许恶意进程获得其他程序在映射内存中的数据内容。 [维基百科](https://zh.wikipedia.org/zh-cn/%E5%B9%BD%E7%81%B5%E6%BC%8F%E6%B4%9E)



```go
func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// defer profile.Start(profile.MemProfile).Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	b := bufio.NewReader(f)
	words := 0
	inword := false
	for {
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}
```

By inserting a  `bufio.Reader`  between the input file and  `readbyte`  will

Compare the times of this revised program to  `wc`. How close is it? Take a profile and see what remains.

在文件内容与 `readbyte` 之间加一个 `bufio.Reader` 缓冲，减少系统调用的次数。
将优化后的版本与 `wc` 比较下，看看还有多少差距？执行一次 profile 看看还有哪些可以改进的地方。


> NOTE 还是比 wc 慢，开启 CPU profile 显示 10ms 时间都在 main ，找不到优化点。难道是 golang 本身就是慢？是慢在内存管理上吗？

```txt
~/tmp$ time wc -w ./2701-0.txt 
215830 ./2701-0.txt

real	0m0.018s
user	0m0.018s
sys	0m0.000s
~/tmp$ time wc -w ./2701-0.txt 
215830 ./2701-0.txt

real	0m0.018s
user	0m0.018s
sys	0m0.000s
~/tmp$ time wc -w ./2701-0.txt 
215830 ./2701-0.txt

real	0m0.018s
user	0m0.018s
sys	0m0.000s
~/tmp$ time wc -w ./2701-0.txt 
215830 ./2701-0.txt

real	0m0.018s
user	0m0.018s
sys	0m0.000s
~/tmp$ time wc -w ./2701-0.txt 
215830 ./2701-0.txt

real	0m0.032s
user	0m0.032s
sys	0m0.000s
~/tmp$ time ./t4 ./2701-0.txt 
"./2701-0.txt": 181276 words

real	0m0.020s
user	0m0.020s
sys	0m0.000s
~/tmp$ time ./t4 ./2701-0.txt 
"./2701-0.txt": 181276 words

real	0m0.022s
user	0m0.023s
sys	0m0.000s
~/tmp$ time ./t4 ./2701-0.txt 
"./2701-0.txt": 181276 words

real	0m0.022s
user	0m0.019s
sys	0m0.004s
~/tmp$ time ./t4 ./2701-0.txt 
"./2701-0.txt": 181276 words

real	0m0.033s
user	0m0.033s
sys	0m0.000s
~/tmp$ time ./t4 ./2701-0.txt 
"./2701-0.txt": 181276 words

real	0m0.022s
user	0m0.023s
sys	0m0.000s
```



#### 3.5.5. Memory profiling

The new  `words`  profile suggests that something is allocating inside the  `readbyte`  function. We can use pprof to investigate.

在 `readbyte` 函数内部还有内存分配的操作，我们可以用 pprof 分析下。

```go
defer profile.Start(profile.MemProfile).Stop()
```

Then run the program as usual

```txt
% go run main2.go moby.txt
2018/08/25 14:41:15 profile: memory profiling enabled (rate 4096), /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile312088211/mem.pprof
"moby.txt": 181275 words
2018/08/25 14:41:15 profile: memory profiling disabled, /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile312088211/mem.pprof
```

![pprof1](https://blog.zeromake.com/public/img/high-performance-go-workshop/pprof-1.svg)


As we suspected the allocation was coming from  `readbyte` — this wasn’t that complicated, readbyte is three lines long:

从上图中可看出来，内存分配操作确实来自 `readbyte` － 这个函数代码只有3行，所以分析起来也很容易。

Use pprof to determine where the allocation is coming from.

通过 pprof 很容易发现内存分配操作来自这里。

```go
func readbyte(r io.Reader) (rune, error) {
        var buf [1]byte // Allocation is here 就是这里在分配内存
        _, err := r.Read(buf[:])
        return rune(buf[0]), err
}
```


We’ll talk about why this is happening in more detail in the next section, but for the moment what we see is every call to readbyte is allocating a new one byte long  _array_  and that array is being allocated on the heap.

现在我们能确定，每次 readbyte 调用都会在 堆(heap) 上分配一个字节长的 数组(array) 。
后面我们会详细讨论为什么发生这种现象。


What are some ways we can avoid this? Try them and use CPU and memory profiling to prove it.

有什么办法能避免这次内存分配操作呢？
试试优化一下，然后用 CPU 和 Memory profile 分析分析。



##### Alloc objects vs. inuse objects

Memory profiles come in two varieties, named after their  `go tool pprof`  flags

-   `-alloc_objects`  reports the call site where each allocation was made.
    
-   `-inuse_objects`  reports the call site where an allocation was made  _iff_  it was reachable at the end of the profile.

Memory profile 中大概有两个类型，在 `go tool pprof` 将其命名为

-   `-alloc_objects`  统计了所有执行内存分配操作的调用点。
    
-   `-inuse_objects`  统计了所有执行内存分配且"直到生成 profile 后，还能继续被访问的内存"的调用点。


> 有两种内存分析策略：[^qcraoPPROF]
> 一种是当前的（这一次采集）内存或对象的分配，称为 inuse；
> 另一种是从程序运行到现在所有的内存分配，不管是否已经被 gc 过了，称为 alloc
    

To demonstrate this, here is a contrived program which will allocate a bunch of memory in a controlled manner.

为了方便说明，制作了下面这个以特定方式控制内存分配过程的演示程序。

```go
const count = 100000

var y []byte

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	y = allocate()
	runtime.GC()
}

// allocate allocates count byte slices and returns the first slice allocated.
func allocate() []byte {
	var x [][]byte
	for i := 0; i < count; i++ {
		x = append(x, makeByteSlice())
	}
	return x[0]
}

// makeByteSlice returns a byte slice of a random length in the range [0, 16384).
func makeByteSlice() []byte {
	return make([]byte, rand.Intn(2^14))
}
```

The program is annotation with the  `profile`  package, and we set the memory profile rate to  `1`--that is, record a stack trace for every allocation. This is slows down the program a lot, but you’ll see why in a minute.

我们会用 `profile` package 来解释这个程序。
另外，我们将 memory profile rate 设置为 1 ，所以会记录每一次内存分配时的 stack trace 。
这会使程序变慢，但你很快就会明白为什么要这样做了。

```txt
% go run main.go
2018/08/25 15:22:05 profile: memory profiling enabled (rate 1), /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile730812803/mem.pprof
2018/08/25 15:22:05 profile: memory profiling disabled, /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile730812803/mem.pprof
```

Lets look at the graph of allocated objects, this is the default, and shows the call graphs that lead to the allocation of every object during the profile.

我们先看 allocate object 类型的分析结果。
这里会显示所有执行内存分配操作的调用图表（调用栈关系）。


```txt
% go tool pprof -http=:8080 /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile891268605/mem.pprof
```

![](https://blog.zeromake.com/public/img/high-performance-go-workshop/pprof-2.svg)

Not surprisingly more than 99% of the allocations were inside  `makeByteSlice`. Now lets look at the same profile using  `-inuse_objects`

你可能会有些惊讶，有超过 99% 的内存分配都发生在 `makeByteSlice` 。
现在再看 inuse object 类型的分析结果吧。


```txt
% go tool pprof -http=:8080 /var/folders/by/3gf34_z95zg05cyj744_vhx40000gn/T/profile891268605/mem.pprof
```

![](https://blog.zeromake.com/public/img/high-performance-go-workshop/pprof-3.svg)

What we see is not the objects that were  _allocated_  during the profile, but the objects that remain  _in use_, at the time the profile was taken — this ignores the stack trace for objects which have been reclaimed by the garbage collector.

这张图中看不到多少 _allocated_ 的对象，只有一些还在 _in use_ 的对象（在 profile 结束前还在使用中的对象）。
也就是说，那些已经被垃圾回收的对象，不会显示在这里。



#### 3.5.6. Block profiling

The last profile type we’ll look at is block profiling. We’ll use the  `ClientServer`  benchmark from the  `net/http`  package

最后一个介绍的是 block profile 类型。
我们使用 net/http package 中的  ClientServer  基准测试来介绍。

```txt
% go test -run=XXX -bench=ClientServer$ -blockprofile=/tmp/block.p net/http
% go tool pprof -http=:8080 /tmp/block.p
```

![](https://blog.zeromake.com/public/img/high-performance-go-workshop/pprof-4.svg)




#### 3.5.7. Thread creation profiling

Go 1.11 (?) added support for profiling the creation of operating system threads.

Add thread creation profiling to  `godoc`  and observe the results of profiling  `godoc -http=:8080 -index`.




#### 3.5.8. Framepointers

Go 1.7 has been released and along with a new compiler for amd64, the compiler now enables frame pointers by default.

Go 1.7 与新版 amd64 编译器一起发布，新版本编译器默认支持 帧指针。

The frame pointer is a register that always points to the top of the current stack frame.

帧指针是一直指向当前堆栈帧顶部的寄存器。

Framepointers enable tools like  `gdb(1)`, and  `perf(1)`  to understand the Go call stack.

帧指针能帮助 gdb(1) perf(1) 这样的调试工具理解 Go 的调用栈。

We won’t cover these tools in this workshop, but you can read and watch a presentation I gave on seven different ways to profile Go programs.

这里不会介绍这些工具，下面的链接中有对 Go program 进行 profile 的详细介绍。

-   [Seven ways to profile a Go program](https://talks.godoc.org/github.com/davecheney/presentations/seven.slide)  (slides)
    
-   [Seven ways to profile a Go program](https://www.youtube.com/watch?v=2h_NFBFrciI)  (video, 30 mins)
    
-   [Seven ways to profile a Go program](https://www.bigmarker.com/remote-meetup-go/Seven-ways-to-profile-a-Go-program)  (webcast, 60 mins)
    

#### 3.5.9. Exercise

-   Generate a profile from a piece of code you know well. If you don’t have a code sample, try profiling  `godoc`.
    
    ```
    % go get golang.org/x/tools/cmd/godoc
    % cd $GOPATH/src/golang.org/x/tools/cmd/godoc
    % vim main.go
    ```
    
-   If you were to generate a profile on one machine and inspect it on another, how would you do it?

- 尝试对熟悉的代码执行一次 profile 。如果找不到合适的代码进行测试，可以试试对 godoc 进行 profile 。
- 如果要在一台机器上生成 profile ，而在另一台机器上分析，如果完成？
    

## 4. Compiler optimisations 编译器优化

This section covers some of the optimisations that the Go compiler performs.

For example;

-   Escape analysis
    
-   Inlining
    
-   Dead code elimination

are all handled in the front end of the compiler, while the code is still in its AST form; then the code is passed to the SSA compiler for further optimisation.


下面介绍 Go 编译器执行的一些优化。

比如：

- 逃逸分析
- 内联
- 去掉无效代码


### 4.1. History of the Go compiler Go 编译器的历史

The Go compiler started as a fork of the Plan9 compiler tool chain circa 2007. The compiler at that time bore a strong resemblance to Aho and Ullman’s  [_Dragon Book_](https://www.goodreads.com/book/show/112269.Principles_of_Compiler_Design).

Go 编译器源于 Plan9 编译器工具链的分支。
这时的编译器与[Dragon Book 龙书 编译器设计原理](https://www.goodreads.com/book/show/112269.Principles_of_Compiler_Design)十分相似。

In 2015 the then Go 1.5 compiler was mechanically translated from  [C into Go](https://golang.org/doc/go1.5#c).

2015 年时， Go 1.5 编译器，已经完成 [从 C 翻译到 Go 实现自举](https://golang.org/doc/go1.5#c) 。

A year later, Go 1.7 introduced a  [new compiler backend](https://blog.golang.org/go1.7)  based on  [SSA](https://en.wikipedia.org/wiki/Static_single_assignment_form)  techniques replaced the previous Plan 9 style code generation. This new backend introduced many opportunities for generic and architecture specific optimistions.

一年后, Go 1.7 基于 [SSA](https://en.wikipedia.org/wiki/Static_single_assignment_form) 技术实现了 [一个新的编译器后端](https://blog.golang.org/go1.7)  从而代替了此前 Plan 9 时代的代码风格。新的编译器后端能为通用体系架构及特定体系架构上提供很多优化空间。



### 4.2. Escape analysis 逃逸分析

The first optimisation we’re doing to discuss is  _escape analysis_.

首先要讨论的优化手段是 _逃逸分析_ 。

To illustrate what escape analysis does recall that the  [Go spec](https://golang.org/ref/spec)  does not mention the heap or the stack. It only mentions that the language is garbage collected in the introduction, and gives no hints as to how this is to be achieved.

在具体介绍逃逸分析前，可回顾下[GO语言编程规范 Go spec](https://golang.org/ref/spec) ，会发现其中完全没有提到 heap 和 stack 。
它仅在引言中提到这一门自动垃圾回收的语言，至于如何实现自动垃圾回收，则全未提及。

A compliant Go implementation of the Go spec  _could_  store every allocation on the heap. That would put a lot of pressure on the the garbage collector, but it is in no way incorrect — for several years, gccgo had very limited support for escape analysis so could effectively be considered to be operating in this mode.

把所有分配的变量都保存到 heap 中，也是一种符合 Go 语文编程规范的实现。
这种方案对垃圾回收器的压力很大，但也不能说这种方案不对 －－ 很多年以来 gccgo 对逃逸分析的支持非常有限，基本可认为 Go 语文就是这样工作的。

However, a goroutine’s stack exists as a cheap place to store local variables; there is no need to garbage collect things on the stack. Therefore, where it is safe to do so, an allocation placed on the stack will be more efficient.

但是在 goroutine 的 stack 空间保存本地变量太方便了（开销小）；stack 中没有垃圾回收。
因此，在确定安全的情况下，在 stack 中分配变量会更有效率。

In some languages, for example C and C++, the choice of allocating on the stack or on the heap is a manual exercise for the programmer—heap allocations are made with  `malloc`  and  `free`, stack allocation is via  `alloca`. Mistakes using these mechanisms are a common cause of memory corruption bugs.

在某些语文中，比如 C 和 C++ ，变量在 stack 还是在 heap 中分配，是由程序员决定的：

调用  `malloc` 或 `free` 函数分配的空间在 heap 分配；

调用 `alloca` 函数分配的空间在 stack 分配。

大量内存损坏的 bug 都是由于这一机制导致的。



In Go, the compiler automatically moves a value to the heap if it lives beyond the lifetime of the function call. It is said that the value  _escapes_  to the heap.

在 Go 语文中，一个变量的生命周期超出函数调用的范围后，编译器会自动把变量移动到 heap 中。
也就是说，变量 _逃逸_ 到 heap 了。

```go
type Foo struct {
	a, b, c, d int
}

func NewFoo() *Foo {
	return &Foo{a: 3, b: 1, c: 4, d: 7}
}
```

In this example the  `Foo`  allocated in  `NewFoo`  will be moved to the heap so its contents remain valid after  `NewFoo`  has returned.

在上面的示例中， `NewFoo` 函数中分配了 `Foo` 变量，随后变量又被移动到 heap 中，以便在函数返回后，继续使用这个变量的值。

This has been present since the earliest days of Go. It isn’t so much an optimisation as an automatic correctness feature. Accidentally returning the address of a stack allocated variable is not possible in Go.

这种特性在 Go 语言发展的早期就实现了。
与其说这是一种优化，不如说这是一种自动纠错功能。
因为，在 Go 语言中，再也不会意外返回一个 stack 的变量地址（而造成 bug 了）。


But the compiler can also do the opposite; it can find things which would be assumed to be allocated on the heap, and move them to stack.

实际是，编译器还会反其道而行之；
它也会找到那些原本分配在 heap 的变量，移动到 stack 上分配。

Let’s have a look at an example

我们看个例子：

```go
func Sum() int {
	const count = 100
	numbers := make([]int, count)
	for i := range numbers {
		numbers[i] = i + 1
	}

	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func main() {
	answer := Sum()
	fmt.Println(answer)
}
```

`Sum`  adds the `int`s between 1 and 100 and returns the result.

`Sum` 返回 1 到 100 的 int 值累加到一起的 和。

Because the  `numbers`  slice is only referenced inside  `Sum`, the compiler will arrange to store the 100 integers for that slice on the stack, rather than the heap. There is no need to garbage collect  `numbers`, it is automatically freed when  `Sum`  returns.

因为 `numbers` slice 只在 `Sum` 函数中使用，所以编译器会把保存这 100 个整数的 slick 变量分配到 stack 上，而不是 heap 上。
所以，这里不需要垃圾回收 `numbers` ，因为在 `Sum` 函数返回后，stack 的变量会自动销毁。



#### 4.2.1. Prove it!

To print the compilers escape analysis decisions, use the  `-m`  flag.

```
% go build -gcflags=-m examples/esc/sum.go
# command-line-arguments
examples/esc/sum.go:22:13: inlining call to fmt.Println
examples/esc/sum.go:8:17: Sum make([]int, count) does not escape
examples/esc/sum.go:22:13: answer escapes to heap
examples/esc/sum.go:22:13: io.Writer(os.Stdout) escapes to heap
examples/esc/sum.go:22:13: main []interface {} literal does not escape
<autogenerated>:1: os.(*File).close .this does not escape
```

Line 8 shows the compiler has correctly deduced that the result of  `make([]int, 100)`  does not escape to the heap. The reason it did no

The reason line 22 reports that  `answer`  escapes to the heap is  `fmt.Println`  is a  _variadic_  function. The parameters to a variadic function are  _boxed_  into a slice, in this case a  `[]interface{}`, so  `answer`  is placed into a interface value because it is referenced by the call to  `fmt.Println`. Since Go 1.6 the garbage collector requires  _all_  values passed via an interface to be pointers, what the compiler sees is  _approximately_:

```
var answer = Sum()
fmt.Println([]interface{&answer}...)
```

We can confirm this using the  `-gcflags="-m -m"`  flag. Which returns

```
% go build -gcflags='-m -m' examples/esc/sum.go 2>&1 | grep sum.go:22
examples/esc/sum.go:22:13: inlining call to fmt.Println func(...interface {}) (int, error) { return fmt.Fprintln(io.Writer(os.Stdout), fmt.a...) }
examples/esc/sum.go:22:13: answer escapes to heap
examples/esc/sum.go:22:13:      from ~arg0 (assign-pair) at examples/esc/sum.go:22:13
examples/esc/sum.go:22:13: io.Writer(os.Stdout) escapes to heap
examples/esc/sum.go:22:13:      from io.Writer(os.Stdout) (passed to call[argument escapes]) at examples/esc/sum.go:22:13
examples/esc/sum.go:22:13: main []interface {} literal does not escape
```

In short, don’t worry about line 22, its not important to this discussion.

#### 4.2.2. Exercises

-   Does this optimisation hold true for all values of  `count`?
    
-   Does this optimisation hold true if  `count`  is a variable, not a constant?
    
-   Does this optimisation hold true if  `count`  is a parameter to  `Sum`?
    

#### 4.2.3. Escape analysis (continued)

This example is a little contrived. It is not intended to be real code, just an example.

```
type Point struct{ X, Y int }

const Width = 640
const Height = 480

func Center(p *Point) {
	p.X = Width / 2
	p.Y = Height / 2
}

func NewPoint() {
	p := new(Point)
	Center(p)
	fmt.Println(p.X, p.Y)
}
```

`NewPoint`  creates a new  `*Point`  value,  `p`. We pass  `p`  to the  `Center`  function which moves the point to a position in the center of the screen. Finally we print the values of  `p.X`  and  `p.Y`.

```
% go build -gcflags=-m examples/esc/center.go
# command-line-arguments
examples/esc/center.go:11:6: can inline Center
examples/esc/center.go:18:8: inlining call to Center
examples/esc/center.go:19:13: inlining call to fmt.Println
examples/esc/center.go:11:13: Center p does not escape
examples/esc/center.go:19:15: p.X escapes to heap
examples/esc/center.go:19:20: p.Y escapes to heap
examples/esc/center.go:19:13: io.Writer(os.Stdout) escapes to heap
examples/esc/center.go:17:10: NewPoint new(Point) does not escape
examples/esc/center.go:19:13: NewPoint []interface {} literal does not escape
<autogenerated>:1: os.(*File).close .this does not escape
```

Even though  `p`  was allocated with the  `new`  function, it will not be stored on the heap, because no reference  `p`  escapes the  `Center`  function.

_Question_: What about line 19, if  `p`  doesn’t escape, what is escaping to the heap?

Write a benchmark to provide that  `Sum`  does not allocate.

### 4.3. Inlining

In Go function calls in have a fixed overhead; stack and preemption checks.

Some of this is ameliorated by hardware branch predictors, but it’s still a cost in terms of function size and clock cycles.

Inlining is the classical optimisation that avoids these costs.

Until Go 1.11 inlining only worked on  _leaf functions_, a function that does not call another. The justification for this is:

-   If your function does a lot of work, then the preamble overhead will be negligible. That’s why functions over a certain size (currently some count of instructions, plus a few operations which prevent inlining all together (eg. switch before Go 1.7)
    
-   Small functions on the other hand pay a fixed overhead for a relatively small amount of useful work performed. These are the functions that inlining targets as they benefit the most.
    

The other reason is that heavy inlining makes stack traces harder to follow.

#### 4.3.1. Inlining (example)

```
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func F() {
	const a, b = 100, 20
	if Max(a, b) == b {
		panic(b)
	}
}
```

Again we use the  `-gcflags=-m`  flag to view the compilers optimisation decision.

```
% go build -gcflags=-m examples/inl/max.go
# command-line-arguments
examples/inl/max.go:4:6: can inline Max
examples/inl/max.go:11:6: can inline F
examples/inl/max.go:13:8: inlining call to Max
examples/inl/max.go:20:6: can inline main
examples/inl/max.go:21:3: inlining call to F
examples/inl/max.go:21:3: inlining call to Max
```

The compiler printed two lines.

-   The first at line 3, the declaration of  `Max`, telling us that it can be inlined.
    
-   The second is reporting that the body of  `Max`  has been inlined into the caller at line 12.
    

_Without_  using the  `//go:noinline`  comment, rewrite  `Max`  such that it still returns the right answer, but is no longer considered inlineable by the compiler.

#### 4.3.2. What does inlining look like?

Compile  `max.go`  and see what the optimised version of  `F()`  became.

```
% go build -gcflags=-S examples/inl/max.go 2>&1 | grep -A5 '"".F STEXT'
"".F STEXT nosplit size=2 args=0x0 locals=0x0
        0x0000 00000 (/Users/dfc/devel/high-performance-go-workshop/examples/inl/max.go:11)     TEXT    "".F(SB), NOSPLIT|ABIInternal, $0-0
        0x0000 00000 (/Users/dfc/devel/high-performance-go-workshop/examples/inl/max.go:11)     FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (/Users/dfc/devel/high-performance-go-workshop/examples/inl/max.go:11)     FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (/Users/dfc/devel/high-performance-go-workshop/examples/inl/max.go:11)     FUNCDATA        $3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (/Users/dfc/devel/high-performance-go-workshop/examples/inl/max.go:13)     PCDATA  $2, $0
```

This is the body of  `F`  once  `Max`  has been inlined into it — there’s nothing happening in this function. I know there’s a lot of text on the screen for nothing, but take my word for it, the only thing happening is the  `RET`. In effect  `F`  became:

```
func F() {
        return
}
```

What are FUNCDATA and PCDATA?

The output from  `-S`  is not the final machine code that goes into your binary. The linker does some processing during the final link stage. Lines like  `FUNCDATA`  and  `PCDATA`  are metadata for the garbage collector which are moved elsewhere when linking. If you’re reading the output of  `-S`, just ignore  `FUNCDATA`  and  `PCDATA`  lines; they’re not part of the final binary.

#### 4.3.3. Discussion

Why did I declare  `a`  and  `b`  in  `F()`  to be constants?

Experiment with the output of What happens if  `a`  and  `b`  are declared as are variables? What happens if  `a`  and  `b`  are passing into  `F()`  as parameters?

`-gcflags=-S`  doesn’t prevent the final binary being build in your working directory. If you find that subsequent runs of  `go build …​`  produce no output, delete the  `./max`  binary in your working directory.

#### 4.3.4. Adjusting the level of inlining

Adjusting the  _inlining level_  is performed with the  `-gcflags=-l`  flag. Somewhat confusingly passing a single  `-l`  will disable inlining, and two or more will enable inlining at more aggressive settings.

-   `-gcflags=-l`, inlining disabled.
    
-   nothing, regular inlining.
    
-   `-gcflags='-l -l'`  inlining level 2, more aggressive, might be faster, may make bigger binaries.
    
-   `-gcflags='-l -l -l'`  inlining level 3, more aggressive again, binaries definitely bigger, maybe faster again, but might also be buggy.
    
-   `-gcflags=-l=4`  (four `-l`s) in Go 1.11 will enable the experimental  [_mid stack_  inlining optimisation](https://github.com/golang/go/issues/19348#issuecomment-393654429).
    

#### 4.3.5. Mid Stack inlining

Since Go 1.12 so called  _mid stack_  inlining has been enabled (it was previously available in preview in Go 1.11 with the  `-gcflags='-l -l -l -l'`  flag).

We can see an example of mid stack inlining in the previous example. In Go 1.11 and earlier  `F`  would not have been a leaf function — it called  `max`. However because of inlining improvements  `F`  is now inlined into its caller. This is for two reasons; . When  `max`  is inlined into  `F`,  `F`  contains no other function calls thus it becomes a potential  _leaf function_, assuming its complexity budget has not been exceeded. . Because  `F`  is a simple function—​inlining and dead code elimination has eliminated much of its complexity budget—​it is eligable for  _mid stack_  inlining irrispective of calling  `max`.

Mid stack inlining can be used to inline the fast path of a function, eliminating the function call overhead in the fast path.  [This recent CL which landed in for Go 1.13](https://go-review.googlesource.com/c/go/+/152698)  shows this technique applied to  `sync.RWMutex.Unlock()`.

##### [](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#further_reading_3)[Further reading](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#further_reading_3)

-   [Mid-stack inlining in the Go compiler presentation by David Lazar](https://docs.google.com/presentation/d/1Wcblp3jpfeKwA0Y4FOmj63PW52M_qmNqlQkNaLj0P5o/edit#slide=id.p)
    
-   [Proposal: Mid-stack inlining in the Go compiler](https://github.com/golang/proposal/blob/master/design/19348-midstack-inlining.md)
    

### 4.4. Dead code elimination

Why is it important that  `a`  and  `b`  are constants?

To understand what happened lets look at what the compiler sees once its inlined  `Max`  into  `F`. We can’t get this from the compiler easily, but it’s straight forward to do it by hand.

Before:

```
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func F() {
	const a, b = 100, 20
	if Max(a, b) == b {
		panic(b)
	}
}
```

After:

```
func F() {
	const a, b = 100, 20
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}
	if result == b {
		panic(b)
	}
}
```

Because  `a`  and  `b`  are constants the compiler can prove at compile time that the branch will never be false;  `100`  is always greater than  `20`. So the compiler can further optimise  `F`  to

```
func F() {
	const a, b = 100, 20
	var result int
	if true {
		result = a
	} else {
		result = b
	}
	if result == b {
		panic(b)
	}
}
```

Now that the result of the branch is know then then the contents of  `result`  are also known. This is call  _branch elimination_.

```
func F() {
        const a, b = 100, 20
        const result = a
        if result == b {
                panic(b)
        }
}
```

Now the branch is eliminated we know that  `result`  is always equal to  `a`, and because  `a`  was a constant, we know that  `result`  is a constant. The compiler applies this proof to the second branch

```
func F() {
        const a, b = 100, 20
        const result = a
        if false {
                panic(b)
        }
}
```

And using branch elimination again the final form of  `F`  is reduced to.

```
func F() {
        const a, b = 100, 20
        const result = a
}
```

And finally just

```
func F() {
}
```

#### 4.4.1. Dead code elimination (cont.)

Branch elimination is one of a category of optimisations known as  _dead code elimination_. In effect, using static proofs to show that a piece of code is never reachable, colloquially known as  _dead_, therefore it need not be compiled, optimised, or emitted in the final binary.

We saw how dead code elimination works together with inlining to reduce the amount of code generated by removing loops and branches that are proven unreachable.

You can take advantage of this to implement expensive debugging, and hide it behind

```
const debug = false
```

Combined with build tags this can be very useful.

#### 4.4.2. Further reading

-   [Using // +build to switch between debug and release builds](http://dave.cheney.net/2014/09/28/using-build-to-switch-between-debug-and-release)
    
-   [How to use conditional compilation with the go build tool](http://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool)
    

### 4.5. Compiler flags Exercises

Compiler flags are provided with:

```
go build -gcflags=$FLAGS
```

Investigate the operation of the following compiler functions:

-   `-S`  prints the (Go flavoured) assembly of the  _package_  being compiled.
    
-   `-l`  controls the behaviour of the inliner;  `-l`  disables inlining,  `-l -l`  increases it (more  `-l`  's increases the compiler’s appetite for inlining code). Experiment with the difference in compile time, program size, and run time.
    
-   `-m`  controls printing of optimisation decision like inlining, escape analysis.  `-m`-m` prints more details about what the compiler was thinking.
    
-   `-l -N`  disables all optimisations.
    

If you find that subsequent runs of  `go build …​`  produce no output, delete the  `./max`  binary in your working directory.

#### 4.5.1. Further reading

-   [Codegen Inspection by Jaana Burcu Dogan](http://go-talks.appspot.com/github.com/rakyll/talks/gcinspect/talk.slide#1)
    

### 4.6. Bounds check elimination

Go is a bounds checked language. This means array and slice subscript operations are checked to ensure they are within the bounds of the respective types.

For arrays, this can be done at compile time. For slices, this must be done at run time.

```
var v = make([]int, 9)

var A, B, C, D, E, F, G, H, I int

func BenchmarkBoundsCheckInOrder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		A = v[0]
		B = v[1]
		C = v[2]
		D = v[3]
		E = v[4]
		F = v[5]
		G = v[6]
		H = v[7]
		I = v[8]
	}
}
```

Use  `-gcflags=-S`  to disassemble  `BenchmarkBoundsCheckInOrder`. How many bounds check operations are performed per loop?

```
func BenchmarkBoundsCheckOutOfOrder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		I = v[8]
		A = v[0]
		B = v[1]
		C = v[2]
		D = v[3]
		E = v[4]
		F = v[5]
		G = v[6]
		H = v[7]
	}
}
```

Does rearranging the order in which we assign the  `A`  through  `I`  affect the assembly. Disassemble  `BenchmarkBoundsCheckOutOfOrder`  and find out.

#### 4.6.1. Exercises

-   Does rearranging the order of subscript operations affect the size of the function? Does it affect the speed of the function?
    
-   What happens if  `v`  is moved inside the  `Benchmark`  function?
    
-   What happens if  `v`  was declared as an array,  `var v [9]int`?
    

## 5. Execution Tracer

The execution tracer was developed by  [Dmitry Vyukov](https://github.com/dvyukov)  for Go 1.5 and remained under documented, and under utilised, for several years.

Unlike sample based profiling, the execution tracer is integrated into the Go runtime, so it does just know what a Go program is doing at a particular point in time, but  _why_.

### 5.1. What is the execution tracer, why do we need it?

I think its easiest to explain what the execution tracer does, and why it’s important by looking at a piece of code where the pprof,  `go tool pprof`  performs poorly.

The  `examples/mandelbrot`  directory contains a simple mandelbrot generator. This code is derived from  [Francesc Campoy’s mandelbrot package](https://github.com/campoy/mandelbrot).

```
cd examples/mandelbrot
go build && ./mandelbrot
```

If we build it, then run it, it generates something like this

![mandelbrot](https://dave.cheney.net/high-performance-go-workshop/images/mandelbrot.png)

#### 5.1.1. How long does it take?

So, how long does this program take to generate a 1024 x 1024 pixel image?

The simplest way I know how to do this is to use something like  `time(1)`.

```
% time ./mandelbrot
real    0m1.654s
user    0m1.630s
sys     0m0.015s
```

Don’t use  `time go run mandebrot.go`  or you’ll time how long it takes to  _compile_  the program as well as run it.

#### 5.1.2. What is the program doing?

So, in this example the program took 1.6 seconds to generate the mandelbrot and write to to a png.

Is that good? Could we make it faster?

One way to answer that question would be to use Go’s built in pprof support to profile the program.

Let’s try that.

### 5.2. Generating the profile

To turn generate a profile we need to either

1.  Use the  `runtime/pprof`  package directly.
    
2.  Use a wrapper like  `github.com/pkg/profile`  to automate this.
    

### 5.3. Generating a profile with runtime/pprof

To show you that there’s no magic, let’s modify the program to write a CPU profile to  `os.Stdout`.

```

import "runtime/pprof"

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()
```

By adding this code to the top of the  `main`  function, this program will write a profile to  `os.Stdout`.

```
cd examples/mandelbrot-runtime-pprof
go run mandelbrot.go > cpu.pprof
```

We can use  `go run`  in this case because the cpu profile will only include the execution of  `mandelbrot.go`, not its compilation.

#### 5.3.1. Generating a profile with github.com/pkg/profile

The previous slide showed a super cheap way to generate a profile, but it has a few problems.

-   If you forget to redirect the output to a file then you’ll blow up that terminal session. 😞 (hint:  `reset(1)`  is your friend)
    
-   If you write anything else to  `os.Stdout`, for example,  `fmt.Println`  you’ll corrupt the trace.
    

The recommended way to use  `runtime/pprof`  is to  [write the trace to a file](https://godoc.org/runtime/pprof#hdr-Profiling_a_Go_program). But, then you have to make sure the trace is stopped, and file is closed before your program stops, including if someone `^C’s it.

So, a few years ago I wrote a  [package](https://godoc.org/github.gom/pkg/profile)  to take care of it.

```

import "github.com/pkg/profile"

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
```

If we run this version, we get a profile written to the current working directory

```
% go run mandelbrot.go
2017/09/17 12:22:06 profile: cpu profiling enabled, cpu.pprof
2017/09/17 12:22:08 profile: cpu profiling disabled, cpu.pprof
```

Using  `pkg/profile`  is not mandatory, but it takes care of a lot of the boilerplate around collecting and recording traces, so we’ll use it for the rest of this workshop.

#### 5.3.2. Analysing the profile

Now we have a profile, we can use  `go tool pprof`  to analyse it.

```
% go tool pprof -http=:8080 cpu.pprof
```

In this run we see that the program ran for 1.81s seconds (profiling adds a small overhead). We can also see that pprof only captured data for 1.53 seconds, as pprof is sample based, relying on the operating system’s  `SIGPROF`  timer.

Since Go 1.9 the  `pprof`  trace contains all the information you need to analyse the trace. You no longer need to also have the matching binary which produced the trace. 🎉

We can use the  `top`  pprof function to sort functions recorded by the trace

```
% go tool pprof cpu.pprof
Type: cpu
Time: Mar 24, 2019 at 5:18pm (CET)
Duration: 2.16s, Total samples = 1.91s (88.51%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1.90s, 99.48% of 1.91s total
Showing top 10 nodes out of 35
      flat  flat%   sum%        cum   cum%
     0.82s 42.93% 42.93%      1.63s 85.34%  main.fillPixel
     0.81s 42.41% 85.34%      0.81s 42.41%  main.paint
     0.11s  5.76% 91.10%      0.12s  6.28%  runtime.mallocgc
     0.04s  2.09% 93.19%      0.04s  2.09%  runtime.memmove
     0.04s  2.09% 95.29%      0.04s  2.09%  runtime.nanotime
     0.03s  1.57% 96.86%      0.03s  1.57%  runtime.pthread_cond_signal
     0.02s  1.05% 97.91%      0.04s  2.09%  compress/flate.(*compressor).deflate
     0.01s  0.52% 98.43%      0.01s  0.52%  compress/flate.(*compressor).findMatch
     0.01s  0.52% 98.95%      0.01s  0.52%  compress/flate.hash4
     0.01s  0.52% 99.48%      0.01s  0.52%  image/png.filter
```

We see that the  `main.fillPixel`  function was on the CPU the most when pprof captured the stack.

Finding  `main.paint`  on the stack isn’t a surprise, this is what the program does; it paints pixels. But what is causing  `paint`  to spend so much time? We can check that with the  _cummulative_  flag to  `top`.

```
(pprof) top --cum
Showing nodes accounting for 1630ms, 85.34% of 1910ms total
Showing top 10 nodes out of 35
      flat  flat%   sum%        cum   cum%
         0     0%     0%     1840ms 96.34%  main.main
         0     0%     0%     1840ms 96.34%  runtime.main
     820ms 42.93% 42.93%     1630ms 85.34%  main.fillPixel
         0     0% 42.93%     1630ms 85.34%  main.seqFillImg
     810ms 42.41% 85.34%      810ms 42.41%  main.paint
         0     0% 85.34%      210ms 10.99%  image/png.(*Encoder).Encode
         0     0% 85.34%      210ms 10.99%  image/png.Encode
         0     0% 85.34%      160ms  8.38%  main.(*img).At
         0     0% 85.34%      160ms  8.38%  runtime.convT2Inoptr
         0     0% 85.34%      150ms  7.85%  image/png.(*encoder).writeIDATs
```

This is sort of suggesting that  `main.fillPixed`  is actually doing most of the work.

You can also visualise the profile with the  `web`  command, which looks like this:

```txt
Type: cpuTime: Sep 17, 2017 at 12:22pm (AEST)Duration: 1.81s, Total samples = 1.53s (84.33%)Showing nodes accounting for 1.53s, 100% of 1.53s totalmainpaintmandelbrot.go1s (65.36%)runtimemainproc.go0 of 1.53s (100%)mainmainmandelbrot.go0 of 1.53s (100%)1.53smainfillPixelmandelbrot.go0.27s (17.65%)of 1.27s (83.01%)1s(inline)image/pngEncodewriter.go0 of 0.26s (16.99%)0.26smainseqFillImgmandelbrot.go0 of 1.27s (83.01%)1.27sruntimemallocgcmalloc.go0.13s (8.50%)of 0.16s (10.46%)runtime(*mcache)nextFreemalloc.go0 of 0.03s (1.96%)0.03simage/png(*encoder)writeImagewriter.go0 of 0.19s (12.42%)main(*img)Atmandelbrot.go0 of 0.18s (11.76%)0.11simage/pngfilterwriter.go0.01s (0.65%)0.01scompress/zlib(*Writer)Writewriter.go0 of 0.07s (4.58%)0.07simage/png(*Encoder)Encodewriter.go0 of 0.26s (16.99%)image/png(*encoder)writeIDATswriter.go0 of 0.19s (12.42%)0.19simage/pngopaquewriter.go0 of 0.07s (4.58%)0.07sruntimeconvT2Inoptriface.go0 of 0.18s (11.76%)0.18ssyscallSyscallasm_darwin_amd64.s0.05s (3.27%)0.16sruntimememmovememmove_amd64.s0.02s (1.31%)0.02scompress/flate(*compressor)deflatedeflate.go0.01s (0.65%)of 0.07s (4.58%)compress/flate(*compressor)findMatchdeflate.go0 of 0.01s (0.65%)0.01scompress/flate(*compressor)writeBlockdeflate.go0 of 0.05s (3.27%)0.05sruntimemmapsys_darwin_amd64.s0.02s (1.31%)compress/flate(*huffmanBitWriter)writehuffman_bit_writer.go0 of 0.05s (3.27%)compress/flate(*dictWriter)Writedeflate.go0 of 0.05s (3.27%)0.05scompress/flate(*huffmanBitWriter)writeTokenshuffman_bit_writer.go0 of 0.05s (3.27%)compress/flate(*huffmanBitWriter)writeBitshuffman_bit_writer.go0 of 0.01s (0.65%)0.01scompress/flate(*huffmanBitWriter)writeCodehuffman_bit_writer.go0 of 0.04s (2.61%)0.04sruntimesystemstackasm_amd64.s0 of 0.03s (1.96%)runtime(*mcache)nextFreefunc1malloc.go0 of 0.02s (1.31%)0.02sruntime(*mheap)allocfunc1mheap.go0 of 0.01s (0.65%)0.01scompress/flatematchLendeflate.go0.01s (0.65%)runtime(*mcentral)growmcentral.go0 of 0.02s (1.31%)runtime(*mheap)allocmheap.go0 of 0.01s (0.65%)0.01sruntimeheapBitsinitSpanmbitmap.go0 of 0.01s (0.65%)0.01sruntimememclrNoHeapPointersmemclr_amd64.s0.01s (0.65%)bufio(*Writer)Flushbufio.go0 of 0.05s (3.27%)image/png(*encoder)Writewriter.go0 of 0.05s (3.27%)0.05sbufio(*Writer)Writebufio.go0 of 0.05s (3.27%)0.05scompress/flate(*Writer)Writedeflate.go0 of 0.07s (4.58%)compress/flate(*compressor)writedeflate.go0 of 0.07s (4.58%)0.07s0.01s0.07scompress/flate(*huffmanBitWriter)writeBlockhuffman_bit_writer.go0 of 0.05s (3.27%)0.05s0.05s0.01s0.05s0.04s0.07simage/png(*encoder)writeChunkwriter.go0 of 0.05s (3.27%)0.05sos(*File)Writefile.go0 of 0.05s (3.27%)0.05s0.19s0.26s0.07sinternal/poll(*FD)Writefd_unix.go0 of 0.05s (3.27%)syscallWritesyscall_unix.go0 of 0.05s (3.27%)0.05s1.27sos(*File)writefile_unix.go0 of 0.05s (3.27%)0.05s0.05s0.03sruntime(*mcache)refillmcache.go0 of 0.02s (1.31%)0.02sruntime(*mcentral)cacheSpanmcentral.go0 of 0.02s (1.31%)0.02s0.02s0.01sruntime(*mheap)alloc_mmheap.go0 of 0.01s (0.65%)0.01sruntime(*mheap)allocSpanLockedmheap.go0 of 0.01s (0.65%)runtime(*mheap)growmheap.go0 of 0.01s (0.65%)0.01s0.01sruntime(*mheap)sysAllocmalloc.go0 of 0.01s (0.65%)0.01sruntimesysMapmem_darwin.go0 of 0.01s (0.65%)0.01sruntimenewMarkBitsmheap.go0 of 0.01s (0.65%)0.01sruntimenewArenaMayUnlockmheap.go0 of 0.01s (0.65%)runtimesysAllocmem_darwin.go0 of 0.01s (0.65%)0.01s0.01s0.01s0.01ssyscallwritezsyscall_darwin_amd64.go0 of 0.05s (3.27%)0.05s0.05s
```

### 5.4. Tracing vs Profiling

Hopefully this example shows the limitations of profiling. Profiling told us what the profiler saw;  `fillPixel`  was doing all the work. There didn’t look like there was much that could be done about that.

So now it’s a good time to introduce the execution tracer which gives a different view of the same program.

#### 5.4.1. Using the execution tracer

Using the tracer is as simple as asking for a  `profile.TraceProfile`, nothing else changes.

```

import "github.com/pkg/profile"

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
```

When we run the program, we get a  `trace.out`  file in the current working directory.

```
% go build mandelbrot.go
% % time ./mandelbrot
2017/09/17 13:19:10 profile: trace enabled, trace.out
2017/09/17 13:19:12 profile: trace disabled, trace.out

real    0m1.740s
user    0m1.707s
sys     0m0.020s
```

Just like pprof, there is a tool in the  `go`  command to analyse the trace.

```
% go tool trace trace.out
2017/09/17 12:41:39 Parsing trace...
2017/09/17 12:41:40 Serializing trace...
2017/09/17 12:41:40 Splitting trace...
2017/09/17 12:41:40 Opening browser. Trace viewer s listening on http://127.0.0.1:57842
```

This tool is a little bit different to  `go tool pprof`. The execution tracer is reusing a lot of the profile visualisation infrastructure built into Chrome, so  `go tool trace`  acts as a server to translate the raw execution trace into data which Chome can display natively.

#### 5.4.2. Analysing the trace

We can see from the trace that the program is only using one cpu.

```
func seqFillImg(m *img) {
	for i, row := range m.m {
		for j := range row {
			fillPixel(m, i, j)
		}
	}
}
```

This isn’t a surprise, by default  `mandelbrot.go`  calls  `fillPixel`  for each pixel in each row in sequence.

Once the image is painted, see the execution switches to writing the  `.png`  file. This generates garbage on the heap, and so the trace changes at that point, we can see the classic saw tooth pattern of a garbage collected heap.

The trace profile offers timing resolution down to the  _microsecond_  level. This is something you just can’t get with external profiling.

go tool trace

Before we go on there are some things we should talk about the usage of the trace tool.

-   The tool uses the javascript debugging support built into Chrome. Trace profiles can only be viewed in Chrome, they won’t work in Firefox, Safari, IE/Edge. Sorry.
    
-   Because this is a Google product, it supports keyboard shortcuts; use  `WASD`  to navigate, use  `?`  to get a list.
    
-   Viewing traces can take a  **lot**  of memory. Seriously, 4Gb won’t cut it, 8Gb is probably the minimum, more is definitely better.
    
-   If you’ve installed Go from an OS distribution like Fedora, the support files for the trace viewer may not be part of the main  `golang`  deb/rpm, they might be in some  `-extra`  package.
    

### 5.5. Using more than one CPU

We saw from the previous trace that the program is running sequentially and not taking advantage of the other CPUs on this machine.

Mandelbrot generation is known as  _embarassingly_parallel_. Each pixel is independant of any other, they could all be computed in parallel. So, let’s try that.

```
% go build mandelbrot.go
% time ./mandelbrot -mode px
2017/09/17 13:19:48 profile: trace enabled, trace.out
2017/09/17 13:19:50 profile: trace disabled, trace.out

real    0m1.764s
user    0m4.031s
sys     0m0.865s
```

So the runtime was basically the same. There was more user time, which makes sense, we were using all the CPUs, but the real (wall clock) time was about the same.

Let’s look a the trace.

As you can see this trace generated  _much_  more data.

-   It looks like lots of work is being done, but if you zoom right in, there are gaps. This is believed to be the scheduler.
    
-   While we’re using all four cores, because each  `fillPixel`  is a relatively small amount of work, we’re spending a lot of time in scheduling overhead.
    

### 5.6. Batching up work

Using one goroutine per pixel was too fine grained. There wasn’t enough work to justify the cost of the goroutine.

Instead, let’s try processing one row per goroutine.

```
% go build mandelbrot.go
% time ./mandelbrot -mode row
2017/09/17 13:41:55 profile: trace enabled, trace.out
2017/09/17 13:41:55 profile: trace disabled, trace.out

real    0m0.764s
user    0m1.907s
sys     0m0.025s
```

This looks like a good improvement, we almost halved the runtime of the program. Let’s look at the trace.

As you can see the trace is now smaller and easier to work with. We get to see the whole trace in span, which is a nice bonus.

-   At the start of the program we see the number of goroutines ramp up to around 1,000. This is an improvement over the 1 << 20 that we saw in the previous trace.
    
-   Zooming in we see  `onePerRowFillImg`  runs for longer, and as the goroutine  _producing_  work is done early, the scheduler efficiently works through the remaining runnable goroutines.
    

### 5.7. Using workers

`mandelbrot.go`  supports one other mode, let’s try it.

```
% go build mandelbrot.go
% time ./mandelbrot -mode workers
2017/09/17 13:49:46 profile: trace enabled, trace.out
2017/09/17 13:49:50 profile: trace disabled, trace.out

real    0m4.207s
user    0m4.459s
sys     0m1.284s
```

So, the runtime was much worse than any previous. Let’s look at the trace and see if we can figure out what happened.

Looking at the trace you can see that with only one worker process the producer and consumer tend to alternate because there is only one worker and one consumer. Let’s increase the number of workers

```
% go build mandelbrot.go
% time ./mandelbrot -mode workers -workers 4
2017/09/17 13:52:51 profile: trace enabled, trace.out
2017/09/17 13:52:57 profile: trace disabled, trace.out

real    0m5.528s
user    0m7.307s
sys     0m4.311s
```

So that made it worse! More real time, more CPU time. Let’s look at the trace to see what happened.

That trace is a mess. There were more workers available, but the seemed to spend all their time fighting over the work to do.

This is because the channel is  _unbuffered_. An unbuffered channel cannot send until there is someone ready to receive.

-   The producer cannot send work until there is a worker ready to receive it.
    
-   Workers cannot receive work until there is someone ready to send, so they compete with each other when they are waiting.
    
-   The sender is not privileged, it cannot take priority over a worker that is already running.
    

What we see here is a lot of latency introduced by the unbuffered channel. There are lots of stops and starts inside the scheduler, and potentially locks and mutexes while waiting for work, this is why we see the  `sys`  time higher.

### 5.8. Using buffered channels

```

import "github.com/pkg/profile"

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
```

```
% go build mandelbrot.go
% time ./mandelbrot -mode workers -workers 4
2017/09/17 14:23:56 profile: trace enabled, trace.out
2017/09/17 14:23:57 profile: trace disabled, trace.out

real    0m0.905s
user    0m2.150s
sys     0m0.121s
```

Which is pretty close to the per row mode above.

Using a buffered channel the trace showed us that:

-   The producer doesn’t have to wait for a worker to arrive, it can fill up the channel quickly.
    
-   The worker can quickly take the next item from the channel without having to sleep waiting on work to be produced.
    

Using this method we got nearly the same speed using a channel to hand off work per pixel than we did previously scheduling on goroutine per row.

Modify  `nWorkersFillImg`  to work per row. Time the result and analyse the trace.

### 5.9. Mandelbrot microservice

It’s 2019, generating Mandelbrots is pointless unless you can offer them on the internet as a serverless microservice. Thus, I present to you,  _Mandelweb_

```
% go run examples/mandelweb/mandelweb.go
2017/09/17 15:29:21 listening on http://127.0.0.1:8080/
```

[http://127.0.0.1:8080/mandelbrot](http://127.0.0.1:8080/mandelbrot)

#### 5.9.1. Tracing running applications

In the previous example we ran the trace over the whole program.

As you saw, traces can be very large, even for small amounts of time, so collecting trace data continually would generate far too much data. Also, tracing can have an impact on the speed of your program, especially if there is a lot of activity.

What we want is a way to collect a short trace from a running program.

Fortuntately, the  `net/http/pprof`  package has just such a facility.

#### 5.9.2. Collecting traces via http

Hopefully everyone knows about the  `net/http/pprof`  package.

```
import _ "net/http/pprof"
```

When imported, the  `net/http/pprof`  will register tracing and profiling routes with  `http.DefaultServeMux`. Since Go 1.5 this includes the trace profiler.

`net/http/pprof`  registers with  `http.DefaultServeMux`. If you are using that  `ServeMux`  implicitly, or explicitly, you may inadvertently expose the pprof endpoints to the internet. This can lead to source code disclosure. You probably don’t want to do this.

We can grab a five second trace from mandelweb with  `curl`  (or  `wget`)

```
% curl -o trace.out http://127.0.0.1:8080/debug/pprof/trace?seconds=5
```

#### 5.9.3. Generating some load

The previous example was interesting, but an idle webserver has, by definition, no performance issues. We need to generate some load. For this I’m using  [`hey`  by JBD](https://github.com/rakyll/hey).

```
% go get -u github.com/rakyll/hey
```

Let’s start with one request per second.

```
% hey -c 1 -n 1000 -q 1 http://127.0.0.1:8080/mandelbrot
```

And with that running, in another window collect the trace

```
% curl -o trace.out http://127.0.0.1:8080/debug/pprof/trace?seconds=5
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 66169    0 66169    0     0  13233      0 --:--:--  0:00:05 --:--:-- 17390
% go tool trace trace.out
2017/09/17 16:09:30 Parsing trace...
2017/09/17 16:09:30 Serializing trace...
2017/09/17 16:09:30 Splitting trace...
2017/09/17 16:09:30 Opening browser.
Trace viewer is listening on http://127.0.0.1:60301
```

#### 5.9.4. Simulating overload

Let’s increase the rate to 5 requests per second.

```
% hey -c 5 -n 1000 -q 5 http://127.0.0.1:8080/mandelbrot
```

And with that running, in another window collect the trace

% curl -o trace.out http://127.0.0.1:8080/debug/pprof/trace?seconds=5
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                Dload  Upload   Total   Spent    Left  Speed
100 66169    0 66169    0     0  13233      0 --:--:--  0:00:05 --:--:-- 17390
% go tool trace trace.out
2017/09/17 16:09:30 Parsing trace...
2017/09/17 16:09:30 Serializing trace...
2017/09/17 16:09:30 Splitting trace...
2017/09/17 16:09:30 Opening browser. Trace viewer is listening on http://127.0.0.1:60301

#### 5.9.5. Extra credit, the Sieve of Eratosthenes

The  [concurrent prime sieve](https://github.com/golang/go/blob/master/doc/play/sieve.go)  is one of the first Go programs written.

Ivan Daniluk  [wrote a great post on visualising](http://divan.github.io/posts/go_concurrency_visualize/)  it.

Let’s take a look at its operation using the execution tracer.

#### 5.9.6. More resources

-   Rhys Hiltner,  [Go’s execution tracer](https://www.youtube.com/watch?v=mmqDlbWk_XA)  (dotGo 2016)
    
-   Rhys Hiltner,  [An Introduction to "go tool trace"](https://www.youtube.com/watch?v=V74JnrGTwKA)  (GopherCon 2017)
    
-   Dave Cheney,  [Seven ways to profile Go programs](https://www.youtube.com/watch?v=2h_NFBFrciI)  (GolangUK 2016)
    
-   Dave Cheney,  [High performance Go workshop](https://dave.cheney.net/training#high-performance-go)]
    
-   Ivan Daniluk,  [Visualizing Concurrency in Go](https://www.youtube.com/watch?v=KyuFeiG3Y60)  (GopherCon 2016)
    
-   Kavya Joshi,  [Understanding Channels](https://www.youtube.com/watch?v=KBZlN0izeiY)  (GopherCon 2017)
    
-   Francesc Campoy,  [Using the Go execution tracer](https://www.youtube.com/watch?v=ySy3sR1LFCQ)
    

## 6. Memory and Garbage Collector 内存和垃圾回收器 GC

Go is a garbage collected language. This is a design principle, it will not change.

Go 是一门自动垃圾回收的语言。
这是设计原则，不会改变。

As a garbage collected language, the performance of Go programs is often determined by their interaction with the garbage collector.

作为一门垃圾回收的语言，Go 程序的性能常常是由垃圾回收器决定的。

Next to your choice of algorithms, memory consumption is the most important factor that determines the performance and scalability of your application.

使用的算法与内存消费是决定程序的性能与扩展性的主要因素。


This section discusses the operation of the garbage collector, how to measure the memory usage of your program and strategies for lowering memory usage if garbage collector performance is a bottleneck.

本节讨论垃圾收集器的工作方法，
以及 如何测试程序的内存使用情况，
还有 当垃圾收集器性能成为瓶颈时，如何降低内存使用量的策略。



### 6.1. Garbage collector world view 垃圾收集器面面观

The purpose of any garbage collector is to present the illusion that there is an infinite amount of memory available to the program.

垃圾收集器的作用是，给程序造成一种错觉，以为有无限内存可用。

You may disagree with this statement, but this is the base assumption of how garbage collector designers work.

你有可能不同意这种观点，但垃圾收集器的作者就是以此为目标来设计的。


A stop the world, mark sweep GC is the most efficient in terms of total run time; good for batch processing, simulation, etc. However, over time the Go GC has moved from a pure stop the world collector to a concurrent, non compacting, collector. This is because the Go GC is designed for low latency servers and interactive applications.

> stop the world STW 暂停时间，GC时要暂停整个程序。

标记清除的GC方案，其总STW时间最短；适用于 batch processing, simulation 等场景。
但现在 Go 的 GC 方案已经从纯粹的 STW 后执行垃圾收集优化为 concurrent, non compacting 的垃圾收集。
这是因为， Go 的 GC 主要为低延迟服务器和交互式应用程序而设计。

The design of the Go GC favors  _lower latency_  over  _maximum throughput_; it moves some of the allocation cost to the mutator to reduce the cost of cleanup later.

Go 的 GC 设计目标是 最大吞吐量 上 降低延迟 ；
它把分配资源的开销转稼到修改过程(mutator)，以此来降低清理资源时的开销。


### 6.2. Garbage collector design 垃圾回收器的设计

The design of the Go GC has changed over the years

-   Go 1.0, stop the world mark sweep collector based heavily on tcmalloc.
    
-   Go 1.3, fully precise collector, wouldn’t mistake big numbers on the heap for pointers, thus leaking memory.
    
-   Go 1.5, new GC design, focusing on  _latency_  over  _throughput_.
    
-   Go 1.6, GC improvements, handling larger heaps with lower latency.
    
-   Go 1.7, small GC improvements, mainly refactoring.
    
-   Go 1.8, further work to reduce STW times, now down to the 100 microsecond range.
    
-   Go 1.10+,  [move away from pure cooperative goroutine scheduling](https://github.com/golang/proposal/blob/master/design/24543-non-cooperative-preemption.md)  to lower the latency when triggering a full GC cycle.
    

Go GC 的设计一直在改进。

- Go 1.0 基于 tcmalloc 实现的 STW 标记清除。
- Go 1.3 更精确的垃圾收集器，能处理 heap 上的大数指针，避免内存泄露问题。
- Go 1.5 全新 GC 设计，改善大吞量下延迟表现。
- Go 1.6 GC 改进，处理大 heap 时，延迟更低。
- Go 1.7 小副 GC 改进，主要是重构。
- Go 1.8 进一步降低 STW 时间为 100 微秒。
- Go 1.10+ 去除 完全协作式 Goroutine 调度，降低触发 full GC cycle 时的延迟。

> 协作式调度 与 抢占式调度
>
> - 协作：被调度方主动弃权；
>   还细分为 主动出让 和 抢战标记 几种方法。
>   缺点：
>       对用户不友好。
>       易出现长久无法停止的代码，无法及时垃圾回收，其他 Goroutine 无法调度。
>
> - 抢占：调试器强制将被调试方被动中断；
>




### 6.3. Garbage collector monitoring  观察垃圾收集器的工作过程

A simple way to obtain a general idea of how hard the garbage collector is working is to enable the output of GC logging.

打开 GC 日志开关，就能看到垃圾收集器的工作过程。

These stats are always collected, but normally suppressed, you can enable their display by setting the  `GODEBUG`  environment variable.

这些统计信息是持续在收集的，但通常不会显示出来，
您也可以通过 `GODEBUG` 环境变量打开显示开关。

```log
% env GODEBUG=gctrace=1 godoc -http=:8080
gc 1 @0.012s 2%: 0.026+0.39+0.10 ms clock, 0.21+0.88/0.52/0+0.84 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 2 @0.016s 3%: 0.038+0.41+0.042 ms clock, 0.30+1.2/0.59/0+0.33 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 3 @0.020s 4%: 0.054+0.56+0.054 ms clock, 0.43+1.0/0.59/0+0.43 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 4 @0.025s 4%: 0.043+0.52+0.058 ms clock, 0.34+1.3/0.64/0+0.46 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 5 @0.029s 5%: 0.058+0.64+0.053 ms clock, 0.46+1.3/0.89/0+0.42 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 6 @0.034s 5%: 0.062+0.42+0.050 ms clock, 0.50+1.2/0.63/0+0.40 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 7 @0.038s 6%: 0.057+0.47+0.046 ms clock, 0.46+1.2/0.67/0+0.37 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 8 @0.041s 6%: 0.049+0.42+0.057 ms clock, 0.39+1.1/0.57/0+0.46 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 9 @0.045s 6%: 0.047+0.38+0.042 ms clock, 0.37+0.94/0.61/0+0.33 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
```

The trace output gives a general measure of GC activity. The output format of  `gctrace=1`  is described in  [the  `runtime`  package documentation](https://golang.org/pkg/runtime/#hdr-Environment_Variables).

根据输出的追踪日志就能判断 GC 的工作状态。
在 [`runtime`  package 文档](https://golang.org/pkg/runtime/#hdr-Environment_Variables) 中会描述 `gctrace=1` 的输出格式。


DEMO: Show  `godoc`  with  `GODEBUG=gctrace=1`  enabled

DEMO: 启用 `GODEBUG=gctrace=1` 时查看 godoc

> Use this env var in production, it has no performance impact.

> 可以在生产环境中使用这个环境变量，它对性能没有影响。

Using  `GODEBUG=gctrace=1`  is good when you  _know_  there is a problem, but for general telemetry on your Go application I recommend the  `net/http/pprof`  interface.

当你十分了解问题时，可通过 `GODEBUG=gctrace=1` 分析。
但我更建议你用  `net/http/pprof` 接口检测自己的 Go 程序。

```go
import _ "net/http/pprof"
```

Importing the  `net/http/pprof`  package will register a handler at  `/debug/pprof`  with various runtime metrics, including:

-   A list of all the running goroutines,  `/debug/pprof/heap?debug=1`.
    
-   A report on the memory allocation statistics,  `/debug/pprof/heap?debug=1`.

`net/http/pprof`  will register itself with your default  `http.ServeMux`.

Be careful as this will be visible if you use  `http.ListenAndServe(address, nil)`.

DEMO:  `godoc -http=:8080`, show  `/debug/pprof`.

 
引用 `net/http/pprof` 包会在 HTTP 处理器中注册 `/debug/pprof` 开头的地址，其中包含各种运行时指标信息。

TODO pprof/heap 地址可能写错了。

- 访问  `/debug/pprof/heap?debug=1` 返回所有运行中的 goroutine 列表。

- 访问  `/debug/pprof/heap?debug=1` 返回内存分配的统计信息报告。


注意：
`net/http/pprof`  会注册到默认的 `http.ServeMux` 上。
只要调用了 `http.ListenAndServe(address, nil)` 代码，这些统计信息就能获取到。

DEMO:  `godoc -http=:8080`, show  `/debug/pprof`.



#### 6.3.1. Garbage collector tuning

The Go runtime provides one environment variable to tune the GC,  `GOGC`.

The formula for GOGC is

goal=reachable⋅(1+GOGC100)goal=reachable⋅(1+GOGC100)

For example, if we currently have a 256MB heap, and  `GOGC=100`  (the default), when the heap fills up it will grow to

512MB=256MB⋅(1+100100)512MB=256MB⋅(1+100100)

-   Values of  `GOGC`  greater than 100 causes the heap to grow faster, reducing the pressure on the GC.
    
-   Values of  `GOGC`  less than 100 cause the heap to grow slowly, increasing the pressure on the GC.
    

The default value of 100 is  _just a guide_. you should choose your own value  _after profiling your application with production loads_.

### 6.4. Reducing allocations

Make sure your APIs allow the caller to reduce the amount of garbage generated.

Consider these two Read methods

```
func (r *Reader) Read() ([]byte, error)
func (r *Reader) Read(buf []byte) (int, error)
```

The first Read method takes no arguments and returns some data as a  `[]byte`. The second takes a  `[]byte`  buffer and returns the amount of bytes read.

The first Read method will  _always_  allocate a buffer, putting pressure on the GC. The second fills the buffer it was given.

Can you name examples in the std lib which follow this pattern?

### 6.5. strings and []bytes

In Go  `string`  values are immutable,  `[]byte`  are mutable.

Most programs prefer to work  `string`, but most IO is done with  `[]byte`.

Avoid  `[]byte`  to string conversions wherever possible, this normally means picking one representation, either a  `string`  or a  `[]byte`  for a value. Often this will be  `[]byte`  if you read the data from the network or disk.

The  [`bytes`](https://golang.org/pkg/bytes/)  package contains many of the same operations — `Split`,  `Compare`,  `HasPrefix`,  `Trim`, etc — as the  [`strings`](https://golang.org/pkg/strings/)  package.

Under the hood  `strings`  uses same assembly primitives as the  `bytes`  package.

### 6.6. Using  `[]byte`  as a map key

It is very common to use a  `string`  as a map key, but often you have a  `[]byte`.

The compiler implements a specific optimisation for this case

```
var m map[string]string
v, ok := m[string(bytes)]
```

This will avoid the conversion of the byte slice to a string for the map lookup. This is very specific, it won’t work if you do something like

```
key := string(bytes)
val, ok := m[key]
```

Let’s see if this is still true. Write a benchmark comparing these two methods of using a  `[]byte`  as a  `string`  map key.

### 6.7. Avoid string concatenation

Go strings are immutable. Concatenating two strings generates a third. Which of the following is fastest?

```
		s := request.ID
		s += " " + client.Addr().String()
		s += " " + time.Now().String()
		r = s
```

```
		var b bytes.Buffer
		fmt.Fprintf(&b, "%s %v %v", request.ID, client.Addr(), time.Now())
		r = b.String()
```

```
		r = fmt.Sprintf("%s %v %v", request.ID, client.Addr(), time.Now())
```

```
		b := make([]byte, 0, 40)
		b = append(b, request.ID...)
		b = append(b, ' ')
		b = append(b, client.Addr().String()...)
		b = append(b, ' ')
		b = time.Now().AppendFormat(b, "2006-01-02 15:04:05.999999999 -0700 MST")
		r = string(b)
```

```
		var b strings.Builder
		b.WriteString(request.ID)
		b.WriteString(" ")
		b.WriteString(client.Addr().String())
		b.WriteString(" ")
		b.WriteString(time.Now().String())
		r = b.String()
```

DEMO:  `go test -bench=. ./examples/concat`

### 6.8. Preallocate slices if the length is known

Append is convenient, but wasteful.

Slices grow by doubling up to 1024 elements, then by approximately 25% after that. What is the capacity of  `b`  after we append one more item to it?

```
func main() {
	b := make([]int, 1024)
	b = append(b, 99)
	fmt.Println("len:", len(b), "cap:", cap(b))
}
```

If you use the append pattern you could be copying a lot of data and creating a lot of garbage.

If know know the length of the slice beforehand, then pre-allocate the target to avoid copying and to make sure the target is exactly the right size.

Before

```
var s []string
for _, v := range fn() {
        s = append(s, v)
}
return s
```

After

```
vals := fn()
s := make([]string, len(vals))
for i, v := range vals {
        s[i] = v
}
return s
```

### 6.9. Using sync.Pool

The  `sync`  package comes with a  `sync.Pool`  type which is used to reuse common objects.

`sync.Pool`  has no fixed size or maximum capacity. You add to it and take from it until a GC happens, then it is emptied unconditionally. This is  [by design](https://groups.google.com/forum/#!searchin/golang-dev/gc-aware/golang-dev/kJ_R6vYVYHU/LjoGriFTYxMJ):

> If before garbage collection is too early and after garbage collection too late, then the right time to drain the pool must be during garbage collection. That is, the semantics of the Pool type must be that it drains at each garbage collection. — Russ Cox

sync.Pool in action

```
var pool = sync.Pool{New: func() interface{} { return make([]byte, 4096) }}

func fn() {
	buf := pool.Get().([]byte) // takes from pool or calls New
	// do work
	pool.Put(buf) // returns buf to the pool
}
```

`sync.Pool`  is not a cache. It can and will be emptied  _at_any_time_.

Do not place important items in a  `sync.Pool`, they will be discarded.

The design of sync.Pool emptying itself on each GC may change in Go 1.13 which will help improve its utility.

> This CL fixes this by introducing a victim cache mechanism. Instead of clearing Pools, the victim cache is dropped and the primary cache is moved to the victim cache. As a result, in steady-state, there are (roughly) no new allocations, but if Pool usage drops, objects will still be collected within two GCs (as opposed to one). — Austin Clements

[https://go-review.googlesource.com/c/go/+/166961/](https://go-review.googlesource.com/c/go/+/166961/)

### 6.10. Exercises

-   Using  `godoc`  (or another program) observe the results of changing  `GOGC`  using  `GODEBUG=gctrace=1`.
    
-   Benchmark byte’s string(byte) map keys
    
-   Benchmark allocs from different concat strategies.
    

## 7. Tips and trips

A random grab back of tips and suggestions

This final section contains a number of tips to micro optimise Go code.

### 7.1. Goroutines

The key feature of Go that makes it a great fit for modern hardware are goroutines.

Goroutines are so easy to use, and so cheap to create, you could think of them as  _almost_  free.

The Go runtime has been written for programs with tens of thousands of goroutines as the norm, hundreds of thousands are not unexpected.

However, each goroutine does consume a minimum amount of memory for the goroutine’s stack which is currently at least 2k.

2048 * 1,000,000 goroutines == 2GB of memory, and they haven’t done anything yet.

Maybe this is a lot, maybe it isn’t given the other usages of your application.

#### 7.1.1. Know when to stop a goroutine

Goroutines are cheap to start and cheap to run, but they do have a finite cost in terms of memory footprint; you cannot create an infinite number of them.

Every time you use the  `go`  keyword in your program to launch a goroutine, you must  **know**  how, and when, that goroutine will exit.

In your design, some goroutines may run until the program exits. These goroutines are rare enough to not become an exception to the rule.

If you don’t know the answer, that’s a potential memory leak as the goroutine will pin its stack’s memory on the heap, as well as any heap allocated variables reachable from the stack.

Never start a goroutine without knowing how it will stop.

#### 7.1.2. Further reading

-   [Concurrency Made Easy](https://www.youtube.com/watch?v=yKQOunhhf4A&index=16&list=PLq2Nv-Sh8EbZEjZdPLaQt1qh_ohZFMDj8)  (video)
    
-   [Concurrency Made Easy](https://dave.cheney.net/paste/concurrency-made-easy.pdf)  (slides)
    
-   [Never start a goroutine without knowning when it will stop](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_never_start_a_goroutine_without_knowning_when_it_will_stop)  (Practical Go, QCon Shanghai 2018)
    

### 7.2. Go uses efficient network polling for some requests

The Go runtime handles network IO using an efficient operating system polling mechanism (kqueue, epoll, windows IOCP, etc). Many waiting goroutines will be serviced by a single operating system thread.

However, for local file IO, Go does not implement any IO polling. Each operation on a  `*os.File`  consumes one operating system thread while in progress.

Heavy use of local file IO can cause your program to spawn hundreds or thousands of threads; possibly more than your operating system allows.

Your disk subsystem does not expect to be able to handle hundreds or thousands of concurrent IO requests.

To limit the amount of concurrent blocking IO, use a pool of worker goroutines, or a buffered channel as a semaphore.

```
var semaphore = make(chan struct{}, 10)

func processRequest(work *Work) {
	semaphore <- struct{}{} // acquire semaphore
	// process request
	<-semaphore // release semaphore
}
```

### 7.3. Watch out for IO multipliers in your application

If you’re writing a server process, its primary job is to multiplex clients connected over the network, and data stored in your application.

Most server programs take a request, do some processing, then return a result. This sounds simple, but depending on the result it can let the client consume a large (possibly unbounded) amount of resources on your server. Here are some things to pay attention to:

-   The amount of IO requests per incoming request; how many IO events does a single client request generate? It might be on average 1, or possibly less than one if many requests are served out of a cache.
    
-   The amount of reads required to service a query; is it fixed, N+1, or linear (reading the whole table to generate the last page of results).
    

If memory is slow, relatively speaking, then IO is so slow that you should avoid doing it at all costs. Most importantly avoid doing IO in the context of a request—don’t make the user wait for your disk subsystem to write to disk, or even read.

### 7.4. Use streaming IO interfaces

Where-ever possible avoid reading data into a  `[]byte`  and passing it around.

Depending on the request you may end up reading megabytes (or more!) of data into memory. This places huge pressure on the GC, which will increase the average latency of your application.

Instead use  `io.Reader`  and  `io.Writer`  to construct processing pipelines to cap the amount of memory in use per request.

For efficiency, consider implementing  `io.ReaderFrom`  /  `io.WriterTo`  if you use a lot of  `io.Copy`. These interface are more efficient and avoid copying memory into a temporary buffer.

### 7.5. Timeouts, timeouts, timeouts

Never start an IO operating without knowing the maximum time it will take.

You need to set a timeout on every network request you make with  `SetDeadline`,  `SetReadDeadline`,  `SetWriteDeadline`.

### 7.6. Defer is expensive, or is it?

`defer`  is expensive because it has to record a closure for defer’s arguments.

```
defer mu.Unlock()
```

is equivalent to

```
defer func() {
        mu.Unlock()
}()
```

`defer`  is expensive if the work being done is small, the classic example is  `defer`  ing a mutex unlock around a struct variable or map lookup. You may choose to avoid  `defer`  in those situations.

This is a case where readability and maintenance is sacrificed for a performance win.

Always revisit these decisions.

### 7.7. Avoid Finalisers

Finalisation is a technique to attach behaviour to an object which is just about to be garbage collected.

Thus, finalisation is non deterministic.

For a finaliser to run, the object must not be reachable by  _anything_. If you accidentally keep a reference to the object in the map, it won’t be finalised.

Finalisers run as part of the gc cycle, which means it is unpredictable when they will run and puts them at odds with the goal of reducing gc operation.

A finaliser may not run for a long time if you have a large heap and have tuned your appliation to create minimal garbage.

### 7.8. Minimise cgo

cgo allows Go programs to call into C libraries.

C code and Go code live in two different universes, cgo traverses the boundary between them.

This transition is not free and depending on where it exists in your code, the cost could be substantial.

cgo calls are similar to blocking IO, they consume a thread during operation.

Do not call out to C code in the middle of a tight loop.

#### 7.8.1. Actually, maybe avoid cgo

cgo has a high overhead.

For best performance I recommend avoiding cgo in your applications.

-   If the C code takes a long time, cgo overhead is not as important.
    
-   If you’re using cgo to call a very short C function, where the overhead is the most noticeable, rewrite that code in Go — by definition it’s short.
    
-   If you’re using a large piece of expensive C code is called in a tight loop, why are you using Go?
    

Is there anyone who’s using cgo to call expensive C code frequently?

##### [](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#further_reading_7)[Further reading](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#further_reading_7)

-   [cgo is not Go](http://dave.cheney.net/2016/01/18/cgo-is-not-go)
    

### 7.9. Always use the latest released version of Go

Old versions of Go will never get better. They will never get bug fixes or optimisations.

-   Go 1.4 should not be used.
    
-   Go 1.5 and 1.6 had a slower compiler, but it produces faster code, and has a faster GC.
    
-   Go 1.7 delivered roughly a 30% improvement in compilation speed over 1.6, a 2x improvement in linking speed (better than any previous version of Go).
    
-   Go 1.8 will deliver a smaller improvement in compilation speed (at this point), but a significant improvement in code quality for non Intel architectures.
    
-   Go 1.9-1.12 continue to improve the performance of generated code, fix bugs, and improve inlining and improve debuging.
    

Old version of Go receive no updates.  **Do not use them**. Use the latest and you will get the best performance.

#### 7.9.1. Further reading

-   [Go 1.7 toolchain improvements](http://dave.cheney.net/2016/04/02/go-1-7-toolchain-improvements)
    
-   [Go 1.8 performance improvements](http://dave.cheney.net/2016/09/18/go-1-8-performance-improvements-one-month-in)
    

#### 7.9.2. Move hot fields to the top of the struct

### 7.10. Discussion

Any questions?

## Final Questions and Conclusion

> Readable means reliable — Rob Pike

Start with the simplest possible code.

_Measure_. Profile your code to identify the bottlenecks,  _do not guess_.

If performance is good,  _stop_. You don’t need to optimise everything, only the hottest parts of your code.

As your application grows, or your traffic pattern evolves, the performance hot spots will change.

Don’t leave complex code that is not performance critical, rewrite it with simpler operations if the bottleneck moves elsewhere.

Always write the simplest code you can, the compiler is optimised for  _normal_  code.

Shorter code is faster code; Go is not C++, do not expect the compiler to unravel complicated abstractions.

Shorter code is  _smaller_  code; which is important for the CPU’s cache.

Pay very close attention to allocations, avoid unnecessary allocation where possible.

> I can make things very fast if they don’t have to be correct. — Russ Cox

Performance and reliability are equally important.

I see little value in making a very fast server that panics, deadlocks or OOMs on a regular basis.

Don’t trade performance for reliability.

----------

[1](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#_footnoteref_1). Hennessy et al: 1.4x annual performance improvment over 40 years.

Version dotgo-2019-3-g660848  
Last updated 2019-04-26 02:55:54 UTC


> Note: exec js in chrome console, to  remove anchor in https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html

```js
var ehs = document.querySelectorAll( "h2" );
ehs.forEach( e => {e.removeChild(e.childNodes[0])});

var ehs = document.querySelectorAll( "h2 a.link" );
ehs.forEach ( e => { 
    e.removeAttribute("href");
});

var ehs = document.querySelectorAll( "h3" );
ehs.forEach( e => {e.removeChild(e.childNodes[0])});

var ehs = document.querySelectorAll( "h3 a.link" );
ehs.forEach ( e => { 
    e.removeAttribute("href");
});

var ehs = document.querySelectorAll( "h4" );
ehs.forEach( e => {e.removeChild(e.childNodes[0])});

var ehs = document.querySelectorAll( "h4 a.link" );
ehs.forEach ( e => { 
    e.removeAttribute("href");
});
```


[^HighPerformanceWorkShopEN]:[High Performance Work Shop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#mechanical_sympathy)

[^HighPerformanceWorkShopGithub]:[High Performance Work Shop Github](https://github.com/davecheney/gophercon2018-performance-tuning-workshop)


[^HighPerformanceWorkShopCN1]:[译文1 High Performance Work Shop](https://www.yuque.com/ksco/uiondt/nimz8b)

[^HighPerformanceWorkShopCN2]:[译文2 High Performance Work Shop](https://blog.zeromake.com/pages/high-performance-go-workshop/)

[^CPUCache]:[CPUCache](https://coolshell.cn/articles/20793.html)

[^MemoryAndNativeCodePerformance]: [内存与本机代码的性能](https://www.infoq.cn/article/2013/07/Native-Performance)

[^PerformanceIntruction]:[性能调优攻略](https://coolshell.cn/articles/7490.html)


[^GoMillionTCP]: [百万 Go TCP 连接的思考: epoll方式减少资源占用](https://colobu.com/2019/02/23/1m-go-tcp-connection/)


[^InterlCPUList]: [英特尔微处理器列表](https://zh.wikipedia.org/wiki/%E8%8B%B1%E7%89%B9%E5%B0%94%E5%BE%AE%E5%A4%84%E7%90%86%E5%99%A8%E5%88%97%E8%A1%A8)。

[^CPUMax4GHz]: [为什么主流CPU的频率止步于4G?我们触到频率天花板了吗？](https://zhuanlan.zhihu.com/p/30409360)

[^qcraoPPROF]: [深度解密Go语言之pprof](https://qcrao.com/2019/11/10/dive-into-go-pprof/)

