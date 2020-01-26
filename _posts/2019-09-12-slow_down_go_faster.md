---
layout: post
title:  "[译] 如何磨刀不误砍柴功"
date:   2019-09-12 22:37:00 +0800
tags: translate
---

* category
{:toc}



# How to Slow Down to Go Faster [^SlowDownGoFaster]

# 如何磨刀不误砍柴功[^SlowDownGoFaster]



## Key Takeaways

## 主要观点

-   Rushing makes us neither faster, nor more productive; it increases stress and distracts focus. We need creativity, effectiveness, and focus.
-   加班不能更快完成任务，更不能提高生产力；它增加压力并且分散注意力。我们需要的是创造力，高效率和专注；
-   Hire better talents, do together, practice together and learn together to improve professionalism and cultivate craftsmanship in your organization.
-   雇佣更好的人才一起工作，互相学习，共同提高专业技术水平。
-   Improve adaptation of your team and efficiency of your processes by doing plans & revising them often, collecting & analyzing and eliminating waste.
-   制订合适的计划并根据实际情况调整，收集分析反馈，提升团队的适应性和效率，减少无用功。
-   Without having a quality codebase, you cannot be agile. Push defects down, release frequently, test first and refactor and focus on simple design.
-   没有高质量的基础代码库，是敏捷不起来的。关注简单的设计，优先测试与重构，频繁发布版本，减少缺陷。
-   Working software doesn’t have to be well-crafted. Only good professionals can build well-crafted software, and only well-crafted software lets you build faster than ever.
-   能正常工作的软件并不一定需要精心设计。只有优秀的专业人士才能创造出精心设计的软件。只有精心设计的软件才能让迅速地增加功能。


Going fast without control could be the biggest enemy of software development. The three main areas where you should slow down in are people, process, and product. Before digging into details, let me start with a story.

不加思考的埋头苦干是软件开发过程最大的忌晦。
你应该放慢 人 流程 产品 这三方面的工作。
在详细展开之前，我们先听一个故事。


I guess it was 2011. I joined a team responsible for building an online marketing platform. My main responsibility was adding new features to the system as fast as I could. Hey, I was the senior developer. We call developers “senior" when they are able to develop faster than others, right? However when I joined, we noticed that it was almost impossible to go fast due to technical debt and design challenges. At every single attempt to go faster, we noticed that we increased the complexity and destroyed the quality. It seemed to me that the only way to gear up was to rewrite the whole system from scratch.

那应该是 2011 年。
我加入一个负责开发在线营销平台的开发团队。
我的主要职责是以最快的速度给系统增加功能。
我可是资深工程师。
就是因为我开发速度比别人快，才能被称为"资深"的，对吧？
由于技术债务和设计上的挑战，开发过程几乎不可能快起来。
我们发现每次尝试加速开发速度时，都会增加复杂性并破坏软件质量。
看来只能用敏捷开发(scratch)的方式重写整个系统了。


I remember, I called the product manager and said we needed to rewrite the whole system. After 30 seconds of silence at phone, the project manager answered, "you are saying that your team wrote your product so poor quality that the same team has to rewrite the same product again, but better this time. Right? Sorry man, it is unacceptable. You should have written better."

我记得，当时打电话给产品经理说，我们需要重写整个系统。
电话那头沉默了 30 秒，项目经回答说，
"你说，你们团队开发的产品质量很差，现在还想把这个产品重写一次。是吗？
不好意思，无法接受。你们之前就应该做好的。"

## Zombie Software Needs Rewrites, Again and Again

## Zombie 软件需要反复重构

According to [Standish Group’s Chaos Report](https://www.projectsmart.co.uk/white-papers/chaos-report.pdf), 94% of software projects are redeveloped from scratch more than once. That’s huge. Similarly, when I look at the products I worked on in the past, I see that almost all of them were rewritten from scratch with newer technologies, architecture and design. Rewriting is so common in the sector that often enterprise companies see it as the only option in project management and innovation. We write, rewrite and rewrite, again and again.

在[Standish Group’s Chaos 报告](https://www.projectsmart.co.uk/white-papers/chaos-report.pdf)中显示，94% 的软件项目至少通过敏捷开发重构过一次。
这的确挺多的了。
当我回头看自己过去参与的产品时，我发现几乎所有项目都普经使用新的技术，构架，设计重构过。
在企业项目管理与创新中，重构是很普遍的。我们开发完软件后，再一遍又一遍地重写。


We have to understand our biggest enemies in software development. In the software world, being fast is vital. Not only being first on the market is important, but also responding to customers by adding new features and eliminating bugs fast keep customer satisfaction high. But we all have a problem with “speed". We assume going faster and smarter and efficiently is related with giving deadlines for the targets. We think that we go faster by working more or with more people. Therefore we either add new people, or do overtime to gear up production. Rushing makes us neither faster, nor more productive. Rushing increases stress, distracts focus and destroys productivity. We need creativity, effectiveness, and focus instead.

我们必须明白在软件开发中，最大的阻碍是什么。
软件的世界中，快很重要。
不仅要快速进入市场，为了保持较高的客户满意度，还要及时添加新功能，修复 Bug ，快速响应客户要求。
但我们对快的理解是有问题的。
一般认为在截止日期前完成目标，就是快，就是有效率。
大家还认为，工作时间越长，参与工作的人越多，就能越快完成工作。
所以，一般都会通过增加人手或加班的方式追赶产品进度。
但这种方式，既不能加快进度，也不能提高效率。
匆忙地工作只会增加压力，分散注意力，降低生产力。
我们需要的是保持创造力和高效率，并且专注；

TODO 翻译 productivity 为 效率 ，有问题吗？

TODO 翻译 rush 为 加班，匆忙有问题吗？


Software development is damn hard and complex. We cannot get rid of complexity. Therefore we have to live with it. The need for speed creates an unstable, unsustainable environment, makes us stressed, less focused and less productive. It just doesn’t work. Team capacity, masterplan, estimation, fixed working hours, deadlines and velocity concepts are imaginary; incompetency is the reality. Delivery times have a direct dependency on people’s skills, the efficiency of processes and the quality of output. Most of the time, developers give hidden deadlines for themselves, without any real need.

软件开发是十分困难的复杂的。
我们无法摆脱这种复杂性。
想快，就必然会增加压力，分散注意力，降低生产力，从而造成一种不稳定，不可持续的局面。
团队能力，总体规划，固定的工作时间，截止日期，速度，这些都是虚拟的概念；能力不足才是真实情况。
交付日期直接依赖个人能力，流程效率，输出质量这些因素。
多数情况开发人员都会提供一个隐藏的截止日期，但并无实际需要。

TODO unstable unsustainableenvironment 翻译为 不可持续的困境

TODO developers give hidden deadlines for themselves, without any real need. 如何翻译？


At the end of the day, we get legacy software. The pressure of deadlines combined with incompetence leads to legacy software, the dead-end of a working software. I’ve been using a different term for legacy software for a while: zombie software. Zombie software fits better because this kind is literally dead, but seems to live in production. It works at production and people gain money from it, but it needs the blood, life and energy of software developers to continue working somehow. Developers are too scared to touch it, therefore if it works, no one wants to change it.

在截止日期的压力与能力不足的状态下，终于得到了 legacy software , 这是 working software 的死胡同(译:软件能使用，但是已经无法维护下去)。
我给 legacy software 起了个别名： zombie software (译:僵尸软件)。
这个名字更合适是因为它们的字面意思非常匹配：形容看起来还活着，事实上已经死亡的产品。
这种产品能正常工作，而且还能给人们赚钱，但它透支软件开发者的鲜血，生命及全部精力来维持运转。
开发者不敢碰它的代码，只要它能正常工作，谁都不会轻易改动代码。

TODO legacy software 传统软件 ？

TODO working software 可以正常工作的软件 ？

TODO dea end 尽头 死胡同 ？


[Robert C. Martin has a perfect saying about symptoms of legacy software at twitter](https://twitter.com/unclebobmartin/status/555013005751377920): "If your software is getting harder and harder to develop, you are doing something wrong.” While rushing, we are destroying the quality so much that every step we take forward makes the whole progress and flow slower than before. I believe slowing down until we reach a sustainable page is the only way of going faster.

[在 twitter 的 Robert C. Martin 有一个关于 legacy software 观点](https://twitter.com/unclebobmartin/status/555013005751377920): "如果你们的软件开发过程越来越困难，就说明出问题了。” 如果开发过程十分匆忙，每前进一步都会让质量变差一点，这会使后面的开发过程越来越慢。我认为放慢脚步，达到一种可持续发展的状态，才能让整个项目的进度变快。



## Rushing is Evil in Software Development

## 加班是软件开发过程的大忌

TODO rush 译成 加班？

TODO evil 译成 大忌？


As [Robert C. Martin mentions on the primary value of software at CleanCoders](https://cleancoders.com/episode/clean-code-episode-9/), “The ability of a software system to tolerate and facilitate such ongoing change is the primary value of software". Rushing is evil in software development. Any attempt to rush causes dramatic damage in productivity, focus, people’s effectiveness, adaptation capability, and tolerance of software.

[Robert C. Martin 在 CleanCoders 谈到软件系统价值时说过](https://cleancoders.com/episode/clean-code-episode-9/), 
“软件系统能够适应并促进持续变化的能力是它的主要价值”。
加班是软件开发过程的大忌。
所有加班的企图都会破坏生产力，干扰注意力，降低效率，降低软件的容错性及灵活性。

TODO adaptation capability 译成 灵活性？

TODO tolerance of software 译成 容错性？


For instance, we always have time for fixing bugs, but no time for writing tests. We don’t refactor and write tests because we don’t have enough time. But we have time for debugging, hacking code and fixing bugs.

比如，我们总是有时间修复 bug ，但是没时间写单元测试。
我们不重构，不写单元测试是因为没时间。
但是我们却有时间调试，有时间分析代码，有时间修复 bug 。


We focus on processes so much that we often forget the main asset in software development: people. Processes help people to improve the way they build products, increase motivation and cultivate a healthy environment. In the end, the efficiency of processes is important, but people are crucial.

我们往往过于关注进度，而忽略软件开发过程的重要资源：人。
提升人的积极性，营造一个健康的工作环境，能加快进度。
进度很重要，但人更重要。

TODO process 译为 流程 还是 进度 ？

TODO process help ... to 如何翻译？


We have to admit that nobody and nothing is perfect. Customers, bosses, managers, team mates, business people, even you yourself, are far from being perfect. Requirements, documents, tools, code, the system you built, and the design, can also never be perfect. Therefore, we have to stop rushing and speeding up without control. The only way of going faster at a sustainable pace is slowing down, in three important areas:

必须承认没有人是完美的。
客户，老板，领导，同事，商业伙伴，甚至你自己，都远远称不上完美。
需求，文档，工具，代码，系统，设计，也从末完美过。
因此，我们必须停止加班，立即放松下来。
想要保持较快的可持续速度的唯一秘决是：放慢脚步，主要指以下几方面：

-   People for improving professionalism and craftsmanship
-   The process for improving adaptation and efficiency
-   Product for improving automation and quality

-   提升人们的专业技术能力
-   提升灵活性和效率
-   提升产品自动化能力和质量

TODO the process 如何翻译？ 提升流程 ？

TODO 如何区分 快 和 慢 ？ 只要加班都算是 rush 吗？



## Areas to Slow down in When It Comes to People

## 当这种人来的时候，就会自然而然慢下来

Processes and tools do not build products, but people do. We have to admit, "talent hiring" is the most important functionality of an organization. It has direct impact on the future of the company and the product itself.

流程和工具不能创造产品，只有人能创造产品。
必须承认，人才招聘对企业非常重要。
它直接影响企业和产品的未来。


Hire the best talent for your organization. By saying “the best”, I do not mean the smartest, or most experienced people around. I look for passion, discipline and motivation at a minimum. If all three exists in a talent, the other skills can grow with ease. Hiring is a win-win process, so both sides should gain from the process. So you should slow down your hiring process and invest on improving it. People join companies in which they believe. So model the behavior you want to see. And through your company culture, your vision and people, make talent believe in you.

聘用最聪明的人。
使用"最"来形容，并非要求最聪明或者经验最丰富这样。
至少要足够热情，自律，积极。
如果一个聪明人具备这三种特点，学习其他技能会很快的。
招聘是双赢的过程，双方都能从中获得好处。
因此放慢招聘的过程，投入并改进这个流程。
人们选择加入一个公司，是因为信任。
在你的工作环境中搜集并提炼一些你认可的特点。
通过公司文化，愿景和人才，来吸引其他人才。

TODO model the behavoir 如何翻译？ 对行为建模？ 译为 提炼行为 可能不太合适。


Ego is a cyanide and kills your organization slowly. Never allow ego enter the door in your organization. From lovable fools to genius jerks, never allow extremes to join your team. Never hire people with ego too. With these people, you can never build a company culture which people admire.

自负有毒，会把你的公司慢慢搞垮。
永远不能让自负的人进门。
不要让可爱的傻瓜或讨厌的人精这种过于极端的人加入团队。
不要聘用自负的人。
有这些人在，永远也建立不了让人羡慕的企业文化。

TODO From ... to ..., never allow extremes join 这句到底是说不让过于极端的人进入团队？那极端的人跟自负的人（ego）有什么联系呢？


Stop working alone and start working together. Never allow silos to occur, because silos or hero developers are the symptoms of dysfunctional organizations. Sit together, closely. Define team standards together. Work in pairs and mobs; review together. Let the responsibility be shared in the team.

不要单独工作，要让团队在一起工作。
出现 silo 或 hero 开发者，正是企业功能失调的症状。
让团队成员做的一起紧密合作，互相监督，每团队每个人都承担起一定的责任。

TODO silo 是什么？ 筒仓 表达什么含义？

Practicing together is the most efficient way to improve your skills. While collaborating, not only do we inspire people, but also we learn from each other. Organize code retreats, randoris and coding dojos regularly in your team. Spend 30 mins of each working day for just practicing.

与团队成员一起实践，才是提高技能最有效的方法。
团队配合，不仅是启发激励别人，还要互相学习。
在团队中定期进行代码审查和代码竞赛。
每周花30分钟就够了。

TODO randoris 什么意思？

TODO dojo 道场 是什么意思？竞赛场？


Let the knowledge flow among people. Learn together. I’ve been organizing Brown Bag / Lunch & Learn sessions since 2010 every week in the teams I worked in. I heard twice, at different times, from my colleagues, “joining sessions every Wednesday allows me to improve myself and that motivates me a lot”. That reflects the power and impact of regular internal meetups in companies.

让知识在人群之间流动起来。2010年开始，我在所在团队内每周组织 Brown Bag / Lunch & 学习课程。
我曾经两次听同事说："每周三的学习都在激励和提升自己"。
这正说明了定期的内部培训带来的正能量。

TODO the power and impact 译成正能量，应该还好吧？


Collect and deliver feedback. To gather collective feedback, you can organize Grand Retrospectives as I have done for years. By the way, Grand Retrospective is a new type of retrospective for digging into problems, with more than 20 people.

收集并提供反馈。
可以像我之前几件的做法一样，举办大型总结（回顾）会。
总结大会一般会超过20人，专门于用于回顾过去的工作，发扬问题。


Teaching and sharing is the best way to master a topic. Be a speaker and give back to the community.

教学和分享是成为某一领域专家的最佳途径。
做一个输出成果的人(发言人)，向社区提供反馈。


Developers seem to hate documentation, but in reality, it is the opposite. Any output people read is simply a documentation; from the production code itself to test code, from commit messages to commit graph, from log messages to error messages, developers document a lot unintentionally. So whatever you document, since people read it to understand, do it better.

开发人员很讨厌文档吗？
实际上刚好相反。
任何可以阅读的内容都可以看做是文档；
从软件产品代码本身，到测试代码，从 commit message 到 commit graph ，从程序输出的日志到错误提示，开发人员有意无意地记录了很多文档。
因此不论你记录什么内容，目的都是为了让别人更容易理解，一定要认真做。


You are not children. Companies are not your parents. We have to own your career and invest in yourself. If investing means spending time and money, then do it for yourself.

你不是小孩了，公司也不是你的父母。
我们必须有自己的事业，并投资自己。
投资是要花费时间和精力的，赶紧动手做去吧。



## How Can We Optimize Processes by Slowing Down?

## 慢下来后，怎么才能优化流程呢？


Every single day, we face with new challenges. These challenges should not be just about market needs or new requirements. Technical challenges also have a great impact on our progress.

每天都要对新的挑战。
这些挑战不应该只来自市场需要或新需求。
技术上的挑战也会对我们产生很大影响。


Plans are nothing, but planning is everything. Do plan and revise often. Especially at the early phases of startups, you need extreme agility. One alignment per day, like via daily Scrum or daily standups, is not enough. You have to collaborate closely, work in pairs, and align more than once everyday. Keep your iteration length short, as short as one week. Create multi feedback loop channels by organizing regular review & demo sessions.

计划是死的，人是活的。做好规划以后，一定要按实际情况调整。
尤其是创业公司的早期阶段，你必须极其敏捷。
要每天跟踪一次进度，通过 Scrum 站立会议的形式。
不仅如此，你还要保持和团队成员密切合作，互相之间随时同步进度。

译： plan (名词，死的计划) are nothing, but planning (动词，做计划的行为) is everything.  像这种短句，网上有很单个词的翻译，但整句的翻译应该是怎么呢？只能自己理解了。

译： Scrum 是敏捷开发的方法学，用于迭代式增量软件开发。 站立会议一般就是 Scrum 会议的一种。详细说明可参考 wiki 。


Define short term goals and long term purposes. Short term goals create focus for your team, and long term purposes prevent distraction from the focus.

定义好短期目标和长期目标。
短期目标给团队创造焦点。
长期目标防止分散注意力。


If you want to understand where you are going wrong, start by visualizing the flows, both technical and business-wise. Visualize failures and fuckups to boost your learning from past experiences.

如果想知道自己犯了哪些错误，首先看技术和业务方面的流程。
直视错误，能从过往的经验中加速成长。



Never make decisions from your gut feeling. Always collect data, analyze and make your decisions based on data. It is also important to allow every developer to access product & code metrics. This increases collective ownership and common sense around product development.

不要凭直觉做决定。
持续收集并分析数据，一切决定都要有数据支撑。
更重要的是，要让开发者能访问产品和代码的指标数据。
这能在产品开发过程中增强集体责任感。



Waste is anything you produce that has no business value. Detect and eliminate waste in the office, in your code and in the processes you follow. Boy scouts leave campgrounds cleaner than they found it. The same philosophy is valid in software development too. Follow the boy scout rule and leave your code cleaner. When you open a file to add new functionality and notice an issue there, fix it without getting any permission. Do not forget to write tests before fixing the issues. This makes you feel confident and comfortable touching your code.

任何没有商业价值的产出都是浪费。
检测并消除办公室，代码等所有流程中的浪费现象。
童子军在离开露营地时，比发现它时还要干净。
相关的理念在软件开发中也适用。
遵守童子军规则，让你的代码保持整洁。
当你打开代码添加新功能时，如果发现问题（不整洁的代码，或 bug ），请直接修复，不需要征求别人同意。
别忘了在修复前编写单元测试。
这能让你写代码的过程更加自信和舒适。

> 译： boy scout rule 童子军规则，《程序员应该知道的97件事》中有提及，保持提交代码时，比 Checkout Out （下载 得到） 代码时要整洁。


You can detect waste at every single point in the software development lifecycle. Obey your definition of done and eliminate "90% done, 90+ remaining" tasks. Never allow long living branches. Long living branches are considered evil. Do not verify your code by manual testing. Manual testing mainly validates the happy path. All other scenarios can only be validated by the test code. So take it seriously.

软件开发生命周期的每个阶段都能发现浪费的现象。
要遵守完成任务的标志，杜绝"90% 90+"现象发生。
永远不要出现长期存在的分支代码。
出现这种分支，说明肯定出了大问题。
不要手动测试自己的代码。
人工测试只检查 happy path （默认的正常场景）。
其他场景只能通过测试代码来验证。
一定要认真对待。

> 译："90% done, 90+ remain" Ninety ninety rule 。
> 前90%的任务要用90%的时间完成，剩下10%的任务还要用90%的时间完成。 -- Tom Cargill, Bell Labs
> 这里表示任务完成标志不清晰。或者给任务的工作目标不清晰。导致完成的工作比实际应该做的少了很多，造成拖延工期的现象发生。

> 译：Happy path 正常路径，测试过程的默认场景，不含任何异常和错误信息的场景，即用户正常使用程序时的流程。
> 异常路径手动测试十分繁琐，所以编写单元测试，或自动化测试代码来覆盖这些场景，会极大减少工作质量，增加畜类。
> 反观自己之前尝试编写单元测试过程中，一直想写一个验证主要业务流程的测试代码。我是不是方向搞反了呢。



## How Can Slowing down Improve the Quality of Products?

## 慢下来后，如何提高产品质量？


One thing is clear. Without having a quality codebase, you cannot be agile, sorry. The first thing you need to do is eliminate technical debt and resolve bugs. If you need to stop building the features for a while, and focus on eliminating bugs.

有一点很清楚。
如果没有一个高质量的基础代码库，你是敏捷不起来的。
首先要消除技术债务并解决 bug 。
专注解决 bug 的时候，你可能有一段时间不能构建新功能。


"Fixing bugs and deploying to servers afterwards" is not a proper procedure today. It contains risks and danger. We need a better and more disciplined way of doing it. When you want to fix a bug, first write a test and reproduce the problem programmatically. Then fix the bug and see that the tests are passing. Deploying to production is safe afterwards.

"修复 bug  ，然后直接部署到服务器上" 如今已经不是一个正常的流程了。
这种过程出错风险很高。
我们需要一个更规范的流程。
修复 bug 前，我们首先要编写能一个能自动重现 bug 并进行测试的代码。
然后再修复 bug ，并观察能否通过之前的测试代码。
最后再发布到生产环境，这个过程就安全多了。


I worked in teams which were spending almost all their time bug fixing and maintaining codebase. These teams suffered from production instability. To continue developing new features while fixing bugs, you need to separate your team into virtual teams. For instance, we select two teammates every iteration to deliver direct technical support and to continue fixing bugs. We call them Batman & Robin. No matter what kind of features you are rushing, bugs have to be fixed without any break.

我在团队的工作，几乎所有时间都用于修复 bug 并维护基础代码库了。
这个团队的产品之前极其不稳定。
为了在修复 bug  时开发新功能，需要把你拉团队分成几个虚拟小组。
比如，我们会在每次迭代时都选出两个团队成员专门提供技术支持并持续修复 bug 。
我们称这两人为 Batman 和 Robin 。
有了这两个人，无论你们有什么样着急的功能要开发，都不会中断修复 bug 的任务。

译： Batman & Robin 是两个英雄人物。 Batman 是蝙蝠侠和罗宾汉。 具体的英雄事迹需要查阅 wiki 。

译： 这样的两个应急成员应该是团队中的技术骨干才能胜任吧。


Today, developers commonly use one practice to slow down the progress, in order to speed up. That is using pull requests. They stop the production line, do verifications and code reviews to improve the code quality. They never deploy unreviewed code to production.

现在，为了提升效率，开发者会执行一个放慢进度的流程。
这就是 pull request （也是 merge request ，合并请求，简称 PR MR ）。
这个过程会停止开发过程，执行代码审核及验证，用于提高代码质量。
这样就永远不会有未经审核的代码上线了。


Our ultimate aim should be to achieve continuous delivery, and release frequently. From git branching mechanisms to deployment strategies, from feedback mechanisms to automated testing practices, it requires a different mindset.

我们的最终目标是，达到持续交付，频繁发布版本的状态。
从 git 分支机制到部署策略，从反馈机制到自动化测试，都需要不同的思维方式。


The practices you use at your SDLC indicate how fast you develop. For git branching mechanism, the philosophy "commit early, commit often, perfect later, publish once” and trunk based development with feature toggling let you eliminate waste.

在软件系统生命周期中的实践方法，决定了你的开发速度。
git 分支的使用原则"尽早提交，频繁提交，延迟优化，一次发布"。
使用功能开关的方式基于主干分支进行开发。
善用这些方法，都能节约你的时间。

译： SDLC(Software Development Life Cycle) 软件系统生命周期，从规划，创建，测试到部署的全过程。

译："commit early often, perfect later, publish once" 是 git 最佳实践，网上有相关英文文章。
 个人理解如下：
 - 尽量保持自己的代码与其他成员的代码差异尽量少。
 - 差异越少，冲突也越少。
 - 所以，时刻保持本地与远程代码同步，时刻保持自己和主干分支同步。



I’ve been using TDD for years. Many people complain to me about the impact of TDD on programming speed. [Joe Rainsberger shared his thoughts about TDD on twitter](https://twitter.com/jbrains/status/167297606698008576): “Worried that TDD will slow down your programmers? Don't. They probably need slowing down.”

我已经使用多年的 TDD （测试驱动开发）。 许多人报怨说 TDD 影响开发速度。 [Joe Rainsberger 在 twitter 分享过他的看法](https://twitter.com/jbrains/status/167297606698008576): “担心 TDD 降低开发人员的速度? 别担心了，他们可能确实需要慢下来。”

译： TDD (Test Driven Development) 测试驱动开发


TDD has more refactoring than testing, more thinking than coding, more simplicity than elegancy. That's why it leads to better quality. Develop by TDD; have just enough tests and simple design.

TDD 提倡重构高于测试，思考多于编码，简单好于高雅。
这就是它能提高质量的原因。
使用 TDD 的方法进行开发，能保证足够的测试和简单的设计。


Have you ever reached 100% code coverage? I achieved this at a three-month project. I wrote unit tests to every single line of production code. At that time, I felt like a hero with super powers. But when we deployed to a pre-production environment for UAT, we noticed that four features were not working. I had to write integration and functional tests to detect the bugs and solve them. Then I realized that unit tests do not guarantee good design and working functionalities. Stop calculating code coverage. Code coverage rates are nothing but stupid information radiators.

你能达到100%的代码覆盖率？
我曾经在一个为期三个月的项目中达到过。
这个项目中的每行代码，我都写了单元测试。
我觉得那里自己就像一个强大的超级英雄。
但是，代码部署到 UAT 预生产环境时，有四个功能无法正常工作。
我只好编写集成测试和功能测试，来检测并修复 bug 。
然后我才明白，单元测试不能保证设计正确且功能可用。
别计算代码覆盖率了，代码覆盖率就是个垃圾。

译： UAT (User Aceptance Test) 用户验证测试。在小公司里，一般就是指生产环境了。

TODO 译： 代码覆盖率仅指单元测试吗？

TODO 译： information radiator 信息散热器怎么理解？



Fix the bugs if you need 30 mins or less to fix them. Additionally, use 20% of your time eliminating technical debt.

如果你能在30分钟内修复 bug 。那么，请再花 20% 的时候消除技术债务。

译： technical debt ，技术债务。指开发人员为快速完成功能，使用了不合理的设计。使后续功能升级和维护过程的更加麻烦。要额外浪费更多时间解决不合理设计引发的问题。


We usually write code for not changing it in the future. Therefore when we design software, similarly we select technologies and tools for not changing them in the future. But we are wrong. Refactoring should be at every stage in our development process. As Kent Beck says, we have to do easy changes to make the changes easy. To achieve that, we govern all our microservices in a mono repository. Mono repository is for making refactoring easy, and that is what we really need.

一般我们写完的代码很少再改动。
因此我们设计软件时，选择的技术和工具也很少改动。
但这是错的。应试在开发过程的每个阶段进行重构。
像 Kent Beck 说的，为了让每次改动都很容易，应该多做一些微小的调整。
为了达到这个目的，尽量把所有的微服务模块都集中到一个代码仓库中。
单一的仓库中能让重构过程更简单，我们确实需要这样。

Any design decision taken before it is required is wrong. Therefore we wait till the most responsible moment to take action. We use hexagonal architecture to active loose coupling and high cohesion at high level design of the system. That also leads well-designed monoliths. By the way, monoliths are not evil, but bad design is. We always start with a monolith and with the help of ports & adaptors, we extract some functionalities as microservices. As [Martin Fowler says](https://martinfowler.com/bliki/MonolithFirst.html)[in the “Monolith First” article at his blog](https://martinfowler.com/bliki/MonolithFirst.html), “going directly to a microservices architecture is risky, so consider a monolithic system first. Split to microservices when and if you need it.”

超出真实需要的设计决策都是错误的。
一定要在最关键的时刻才行动。
把系统设计成复杂蜂巢结构来达到高内聚，低耦合的特性。
这可能会把系统做成精心设计的大石头。
做成巨石没事，设计不对就会出大问题。
我们先把系统做 拥有灵活对接（适配）接口 的大石头，然后再按功能分解成微服务模块。
正如 [Martin Fowler 在 "Monolith First" 这篇博客文章所说](https://martinfowler.com/bliki/MonolithFirst.html), “直接设计成微服务架构是很危险的, 先做成一个巨大的系统。然后按实际需要再分隔成微服务模块"。

译： loose coupling and high cohesion 高内聚，低耦合

译： hexagonal 6边形，译成 蜂巢 怎么样？

译： monolith 巨石，译成 庞然大物 怎么样？



## Slowing down to Go Fast as a Philosophy

## 慢下来后再加速是一种哲学


[Andreas Möller mentioned how he feels about software development in a tweet](https://twitter.com/localheinz/status/472386653173342208): “I don't want to write code that just works: I want to write code that is clean, maintainable, easy to understand and well tested.”

[Andreas Möller 曾经在 tweet 提到他对于软件开发的看法](https://twitter.com/localheinz/status/472386653173342208): “我不想只为应付工作而写代码：我希望能写出整洁，可维护，易于理解，便于测试的代码。”


To achieve that, we need to focus on three areas: people, process and product. By slowing down on people, we aim to improve professionalism and craftsmanship. By slowing down on process, we aim to improve adaptation and efficiency. And by slowing down on product, we aim to improve automation and quality. When we focus on these areas, we start to cultivate a development culture enabling software development fast.

为达到这个目的，我们要关注这三方面：人，流程和产品。
让人们放慢脚步，我们的目标是加强专业性，并提高工艺水平。
放慢流程，是为了增强灵活性，提高效率。
放慢产品，是为了提升自动化水平并提高质量。
当我们开始关注这些方面时，也就培养了一种能加速软件开发过程的文化。


We should not forget the followings. Working software does not have to be well-crafted. Only good professionals can build well-crafted software. And only well-crafted software lets you build features faster than ever.

也还要忘记下面几点。
能正常工作的软件并不一定需要精心设计。
只有优秀的专业人士才能创造出精心设计的软件。
只有精心设计的软件才能让你更快地增加功能。


## About the Author

## 关于作者

**![](https://res.infoq.com/articles/slow-down-go-faster/en/resources/1Lemi-Orhan-Ergin-1546546016896.jpg)Lemi Orhan Ergin**  is a [software craftsman](http://www.lemiorhanergin.com/) with a passion for raising the bar of his profession and sharing his experiences with communities. He is the co-founder of [Craftbase](http://www.craftbase.io/) and [founder of](https://www.meetup.com/software-craftsmanship-turkey) [Turkish Software Craftsmanship Community](https://www.meetup.com/software-craftsmanship-turkey). He has been programming since 2001. He is an active practitioner and mentor in Scrum, extreme programming, engineering practices and web development technologies.

Lemi Orhan Ergin 是一位热衷于提升专业水平并热爱分享的[软件工匠](http://www.lemiorhanergin.com/) 。
他是 [Craftbase](http://www.craftbase.io/) 的联合创始人，也是 [Turkish Software Craftsmanship Community](https://www.meetup.com/software-craftsmanship-turkey) 的创始人。
从2001年开始编程至今。
他是 Scrum,极限编程，工程实践和 web 开发技术的实践和倡导者。


[^SlowDownGoFaster]: [slow-down-go-faster](https://www.infoq.com/articles/slow-down-go-faster/?utm_source=wanqu.co&utm_campaign=Wanqu+Daily&utm_medium=website)

<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE4NDUyMzQwMTksLTMzNTg5ODcyNV19
-->
