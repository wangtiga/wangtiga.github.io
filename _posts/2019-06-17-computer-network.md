---
layout: post
title: "计算机网络笔记"
date: 2019-06-17 00:00:00 +0800
tags: network
---

* category
{:toc}

 > 计算机网络笔记  
 > 《计算机网络》第五版（谢希仁）是对我十分有用的一本书，解答了很多疑惑。  
 
 
## 1.为什么出现路由器交换机之类的东西。 P11
 
 -   两台设备要通讯时，简单的用一根连接线接通就可以了。  
     但是试想3台设备，或者更多设备时，要怎么才能方便接通各个设备呢？
 -   可以像 图(b) 一样，每个设备都接上线，但这样不仅浪费线缆，每个设备上也要有这么多接口才行啊。  
     所以如果有个设备能收集转换数据到正确的设备中，那当然很方便。就像 图(c) 一样。  
     当然，交换机、路由器的功能远不止这些。  
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img1.png)
 
## 2.有关OSI七层协议。P27
 
 -   第一次看到教科书里讲的OSI七层协议，着实让我迷糊了好长一段时间。为什么要分层，为什么还偏就要分七层，每次都干什么的呢？  
     看了这本书的的简介以后确实让我有大跌眼镜的感觉。
     -   OSI七层 图(a)：概念清楚，理论完整，但既复杂又不实用，所以现实使用的设备没有按这个协议设计的。 * 但此模型仍然是其他协议的理论基础。
     -   TCP/IP体系 图(b)：这是真正抢占到市场中的协议，TELNET / TCP / IP 也都是我曾经听说过的东西。注意，网络接口层没有实际内容。
     -   五层协议 图(c)：统合OSI和TCP/IP的优点，作者又设计了方便学习原理的五层协议。再往后看，就能明白具体各层协议的原理和作用了。  
         ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img2.png)
 
## 3.数据在协议各层中的传输过程。P29
 
 -   APP层 应用程序数据  
     比如主机1要打开主机2中的网页文件  [](http://www.tsinghua.edu.cn/chn/yxsz/index.html)[http://www.tsinghua.edu.cn/chn/yxsz/index.html](http://www.tsinghua.edu.cn/chn/yxsz/index.html)
 -   5层应用层 数据部分(Get Host Connection等都是http协议要示的数据格式)：  
     Get  [](http://www.tsinghua.edu.cn/chn/yxsz/index.htm)[http://www.tsinghua.edu.cn/chn/yxsz/index.htm](http://www.tsinghua.edu.cn/chn/yxsz/index.htm)  HTTP/1.1.1  
     Host:  [www.tsinghua.edu.cn](http://www.tsinghua.edu.cn/)  
     Connection: close  
     User-Agent:Mozilla/5.0  
     Accept-Language:cn
 -   4层运输层 数据部分  
     (图P194 )  
     TCP首部 + 5层应用层的数据
 -   3层网络层 数据部分  
     (图P122)  
     IP 首部 + 4层运输层的数据
 -   2层数据链路层 数据部分  
     (图P89)  
     以太网MAC帧首部 + 3层网络层的数据
 -   1层物理层 数据部分  
     (全部是比特数据流)  
     100101011101011011
 
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img3.png)
 
## 4.TCP/IP的协议族
 
 TCP/IP的协议是目前实际使用的协议。
 
 -   网际层IP协议起到了承上启下的重要作用。
 -   向上的应用层(http smtp等)各协议都会转为IP数据包交由下层设备传输，
 -   向下的网络接口层（物理层、数据链路层）都可处理统一的IP数据包，方便处理不同异构网络连接。  
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img4.png)
 
## 5.数据通信系统模型P37
 
 一个数据通信系统可分为三部分，源系统、传输系统、目的系统  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img5.png)
 
## 6.信道复用P48
 
 我理解为一个信道供多个用户使用。（或者一根网线中，可传输多个用户的不同数据（理解可能不太准确））  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img6.png)
 
## 7.局域网使用广播形势进行通信，采用总线型的拓扑结构，使用 CSMA/CD协议进行通信，这是数据链路层的协议。P76
 
 > TODO 至于为什么局域网采用广播,即一对多的方式进行通信，而不是一对一？而协议又为什么选定CSMA/CD，我都没有搞清楚。。惭愧啊，有待日后研究
 
 下图是常见的局域网拓扑，一直也搞不懂这几种拓扑有什么区别？越看越觉得好像每种拓扑都差不多，又有点不同，但有说不出哪点不一样。  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img7.png)
 
## 8.适配器（网卡）的工作原理。P79
 
 工作在 3层网络层 和 2层数据链路层 之间的设备。  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img8.png)
 
## 9.CSMA/CD协议要点。P83
 
 局域网使用总线型的拓扑，广播的方式通信，当一台计算机发送数据时，总线上的其他计算机都能检测到这个数据。也就是说局域网中，同一时刻只能有一台计算机发送数据，否则计算机之间就会受到干扰。那么如果才能实现一对一的通讯呢？  
 我们可以“发送前先监控”，即每个站在发送数据前先检测一下是否有其他站在发送数据。如果有，则暂时不要发送，等待信道上空闲的时候再发送。这叫“载波监听”。  
 显然，使用CSMA/CD协议时，一个站不可能同时进行发送和接收，因此使用CSMA/CD协议的以太网不能进行全双式通信，只能双向交替通信（半双工通信）。
 
 -   1）适配器从网络层获得一个分组，加上以太网的首部和尾部，组成以太网帧，放入适配器的缓存，准备发送。
 -   2）叵适配器检测到信道空闲（96比特时间没有检测到信道上有信号），就发送帧，若检测到信道忙，则继续检测并等待信道转为空闲（加上96比特时间），然后发送帧。
 -   3）发送过程继续检测信道，若一直末检测到碰撞，就顺利把这个帧成功发送完毕。若检测到碰撞，则中止数据的发送，并发送人工干扰信号。
 -   4）在中止发送后，适配器就执行指数退避算法。等待r倍512时间后，返回到步骤2).
 
## 10.以太网的MAC层P86
 
 局域网中，硬件地址又称物理地址或MAC地址，也就是适配器（网卡）的地址。  
 适配器有过滤功能。但适配器从网络上每收到一个MAC帧，就先用硬件检查MAC帧的目的地址，如果是发往本站的帧则收下，然后再进行其他处理。否则就将此帧丢弃，不再进行其他的处理。这样做就不浪费主机的处理机和内存资源。这里“发往本站的帧”包括以下三种帧：
 
 -   1）单播（unicast）帧，（一对一），即收到帧的MAC地址与本站的硬件地址相同。
 -   2）广播（broadcast）帧，（一对全体），即发送给局域网上所有站点的帧（全1地址）。
 -   3）多播（multicast）帧，（一对多），即发送给本局域网上一部分站点的帧。  
     所有适配器至少识别单播和广播。有的适配器可用编程方法识别多播地址。
 
## 11.混杂方式（promiscuous mode）P88
 
 工作在混杂方式的适配器只要“听到”有帧在以太网上传输就都悄悄接收。这样实际上是“窃听”其他站点的通信而不中断其他站点的通信。黑客可用此方法获取网上用户的口令。网络管理员也可以用这种方法分析以太网上的流量。
 
## 12.MAC帧的格式P89
 
 以太网上传送数据以帧为单位，各帧间还必须有一定的间隙。因此，接收端只要找到帧开始定界符，其后面连续到达的比特流都属于同一个MAC帧。可见以太网不需要使用帧结束定界符，也不需要使用字节插入来证明透明传输。  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img12.png)
 
## 13.网桥-数据链路层扩展以太网P93
 
 网桥工作在数据链路层，根据MAC帧的目的地址对收到的帧进行转发和过滤。当网桥收到一个帧时，并不是向所有的接口转发此帧，是先检查此帧的目的MAC地址，然后确定将该帧转发到哪个接口。或者把它丢弃。网桥依靠转发表转发帧。（我把交换机理解为有多个接口的网桥）  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img13.png)
 
## 14.虚拟局域网P99
 
 下图是四个交换机的网络拓扑。其中10个工作站分配在三个楼层，  
 这10个工作站划分为三个工作组，也就是三个虚拟局域网。  
 VLAN1:(A1,A2,A3,A4) VLAN2:(B1,B2,B3) VLAN3:(C1,C2,C3)
 
 -   vlan最主要的作用是划分了广播域:
 -   如果划分了三个虚拟局域网，那么VLAN1中的A1发送广播数据，只有A2、A3、A4能收到广播，而B!和C1是收不到广播数据的。因为这两个工作站和A1不在一个虚拟局域网。
 -   如果不划分虚拟局域网，那么A1发送的广播就能被另外9个工作站收到(这个结论是我个人理解，不知道是否有误，因为我觉得交换机2连接交换机1的接口，与交换机1连接A1的接口并无区别）。
 -   划分这个广播域又有什么用呢?
 -   防止广播风暴  
     A1广播的数据,只能在A1所在的虚拟局域网中广播.若VLAN!中出现广播风暴,也不会影响其他网络.如果没有划分虚拟局域网,广播风暴就将交换机1、2、3、4下在所有电脑都影响到。
 -   安全  
     比如VLAN里都是领导，给他们单独划分一个VLAN,那么其他电脑比如B1、B1、C1、C2即使自己更改IP到与VLAN1相同的网段,也无法与之通信.因为,不同VLAN间的报文传输时是相互隔离的,如果不同VLAN间要相互通信,必须通过路由器,或者三层交换.也就是下图中交换机1的位置必须是一个路由器或者三层交换机,并且配置了相应的路由转发策略.比如将源IP是VLAN2(192.168.2.0/24)的数据转发到VLAN1的(192.168.1.0/24)的网络.(猜测实际实现时,就是更改以太网帧中的VLAN标记)
     
     > 参考:  [http://wenku.baidu.com/view/74977bf9fab069dc502201f9](http://wenku.baidu.com/view/74977bf9fab069dc502201f9)  
     > ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img14-1.png)  
     > ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img14-2.png)
     
 
## 15.路由器P112
 
 工作在网络层，连接异构网络，使用IP协议  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img15.png)
 
## 16.ARP工作原理P120
 
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img16.png)
 
## 17.协议格式 
### 17.1 IP数据报的格式P122

 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img17-1.png)
 
 <table border="1">
   <caption border="1">IPv4 Header Format</caption>
   <tbody border="1">
     <tr>
       <td colspan="01">位偏移</td>
       <td colspan="08">00-07</td>
       <td colspan="08">08-15</td>
       <td colspan="08">16-23</td>
       <td colspan="08">24-32</td>
     </tr>	
     <tr>
       <td colspan="01">-</td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
     </tr>
     <tr>
       <td colspan="01">0</td>
       <td colspan="04">版本</td>
       <td colspan="04">首部长度</td>
       <td colspan="06">区分服务</td>
       <td colspan="02">显式拥塞通告</td>
       <td colspan="16">全长</td>
     </tr>
     <tr>
       <td colspan="01">32</td>
       <td colspan="16">标识符 Identification</td>
       <td colspan="03">标志 Flag</td>
       <td colspan="13">分片偏移 Fragment Offset</td>
     </tr>
     <tr>
       <td colspan="01">64</td>
       <td colspan="08">存活时间 TimeToLive</td>
       <td colspan="08">协议 Protocol</td>
       <td colspan="16">首部检验和</td>
     </tr>
     <tr>
       <td colspan="01">96</td>
       <td colspan="32">源IP地址</td>
     </tr>
     <tr>
       <td colspan="01">128</td>
       <td colspan="32">目的IP地址</td>
     </tr>
     <tr>
       <td colspan="01">160</td>
       <td colspan="32">选项（如首部长度&gt;5）</td>
     </tr>
     <tr>
       <td colspan="01">160 or 192+</td>
       <td colspan="32">数据</td>
     </tr>
   </tbody>
 </table>
 
 -   1）版本，4位，IP协议版本，广泛使用的是IPv4，IPv6尚未普及
 -   2）首部长度，4位，首部长度，最大值为15，单位32位字（4字节）
 -   3）服务类型（区分服务+显式拥塞通告），8位，并未实际应用
 -   4）总长度，16位，首部和数据长之和的长度，最大值65535，单位字节
 -   5）标识，16位，不是序号，仅供数据报分片时，原属同一数据报的分片，标识相同。但部分系统实现时，是全局递增序号。
 -   6）标志，3位，仅后两位有意义，  
     （XX DF MF）
     *   MF=1表示后面还有分片的数据报，MF=0表示这是若十数据报片中最后一个。
     *   DF=1表示不允许路由器对引数据报分片，DF=0表示允许
 -   7）片偏移，13位，以8字节为偏移单位，表示当前分片在原分组（原数据报）中的相对位置
 -   8）生存时间，8位，数据报在网络中的寿命，由源点设置，每经过一个路由器减1，为0时会被丢弃
 -   9）协议，8位，指明数据报携带的数据使用何种协议（TCP/ICMP/UDP/IPv6）等
 -   10）首部检验和，16位，只检验数据报首部，不包括数据部分
 -   11）源地址，32位
 -   12）目的地址，32位
 -   13）可选字段（选项），长度可变，1字节到40字节不等。很少使用
 -   14）数据
 
### 17.2 TCP协议格式  
 
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img17-2.png)
  
 <table border="1">
   <caption border="1">TCP Header Format</caption>
   <tbody border="1">
     <tr>
       <td colspan="01">位偏移</td>
       <td colspan="08">00-07</td>
       <td colspan="08">08-15</td>
       <td colspan="08">16-23</td>
       <td colspan="08">24-32</td>
     </tr>
     <tr>
       <td colspan="01">-</td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
     </tr>
     <tr>
       <td colspan="01">0</td>
       <td colspan="16">源端口</td>
       <td colspan="16">目的端口</td>
     </tr>
     <tr>
       <td colspan="01">32</td>
       <td colspan="32">序号 Sequence Number</td>
     </tr>
     <tr>
       <td colspan="01">64</td>
       <td colspan="32">确认号 Acknowledgment Number(if ACK set)</td>
     </tr>
     <tr>
       <td colspan="01">96</td>
       <td colspan="04">数据偏移 Data offset</td>
       <td colspan="06">保留+</td>
       <td colspan="01">URG</td>
       <td colspan="01">ACK</td>
       <td colspan="01">PSH</td>
       <td colspan="01">RST</td>
       <td colspan="01">SYN</td>
       <td colspan="01">FIN</td>
       <td colspan="16">窗口大小</td>
     </tr>
     <tr>
       <td colspan="01">128</td>
       <td colspan="16">校验和</td>
       <td colspan="16">紧急指针</td>
     </tr>
     <tr>
       <td colspan="01">64</td>
       <td colspan="08">存活时间</td>
       <td colspan="08">协议</td>
       <td colspan="16">首部检验和</td>
     </tr>
     <tr>
       <td colspan="01">160</td>
       <td colspan="32">选项（如首部长度&gt;5）</td>
     </tr>
     <tr>
       <td colspan="01">160 or 192+</td>
       <td colspan="32">数据</td>
     </tr>
   </tbody>
 </table>
 
 -   1）源端口，16位，源端口号，对方回信使用
 -   2）目的端口，16位，终点交付报文使用
 -   3）序号，32位，取舍范围[0,2^32 -1 ]，序号增加到最大时，会回到0。上一数据报的 序号 加 其长度，就是本次数据报的序号
 -   4）确认号，32位，期望收到的对方下一报文段的第一数据字节的序号
 -   5）数据偏移，4位，指明TCP报文段 数据起始处 距离 TCP报文段起始处 有多远。因为TCP首部选项长度可变，所以需要此字段。单位32位字（4字节）
 -   6）保留，6位，保留，未使用，必须置0
 -   7）紧急URG，1位，为1时表明报文段是紧急数据，需要尽快传送
 -   8）确认ACK，1位，为1时确认号字段才有效
 -   9）推送PUSH，1位，为1时，接收方会立即将数据交付应用进程（而非等到缓存填满才交付）
 -   10）复位RST，1位，表示TCP连接出错，必须释放连接
 -   11）同步SYN，1位，为1时，表示请求建立连接
 -   12）终止FIN，1位，为1时，表示释放连接
 -   13）窗口，16位，取值[0, 2^16 -1]，表示发送方的接收窗口，单位字节。值是动态调整的。
 -   14）检验和，16位，检验首部和数据，类似UDP，需要加伪首部
 -   15）紧急指针，16位，URG=1时有效，表明本报文段中紧急数据的字节数
 -   16）选项，长度可变，最长40字节，无选项时，TCP首部长度20字节
 
### 17.3 UDP协议格式

![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img17-3.png)
  
 <table border="1">
   <caption border="1">UDP Header Format</caption>
   <tbody border="1">
     <tr>
       <td colspan="01">位偏移</td>
       <td colspan="08">00-07</td>
       <td colspan="08">08-15</td>
       <td colspan="08">16-23</td>
       <td colspan="08">24-32</td>
     </tr>
     <tr>
       <td colspan="01">-</td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
       <td colspan="01"></td>
     </tr>
     <tr>
       <td colspan="01">0</td>
       <td colspan="16">源端口</td>
       <td colspan="16">目的端口</td>
     </tr>
     <tr>
       <td colspan="01">32</td>
       <td colspan="16">长度</td>
       <td colspan="16">校验和</td>
     </tr>
   </tbody>
 </table>
 
 -   1）源端口，16位，源端口号，对方回信使用
 -   2）目的端口，16位，终点交付报文使用
 -   3）长度，16位，UDP数据报长度，单位字节，最小值是8（仅有首部），最大65535
 -   4）检验和，16位，检测数据报传输过程是否有差错，计算时要加上前面12字节的伪首部，IP Header

## 18.防止数据兜圈子的方法。
 
 -   网络(IP)层：IP协议中有生存时间字段(TTL)，指明数据报在因特网中最多可经由的路由器个数。也叫“跳数限制”。P124
 -   物理(MAC)层：透明网桥使用生成树的方法；不透明网桥使用源路由网桥。P96
 
## 19.IP层转发分组的流程。P127
 
 分组转发算法：
 
 -   1）从数据报首部提取目的主机的IP地址D,得出目的网络地址N。
 -   2）若N就是此路由器直接连接的某个网络地址，则进行直接交付。不需要经过其他路由器，直接反数据报交付给目的主机（这里包括把目的主机地址D转换为具体的硬件地址，把数据报封装为MAC帧，再发送此帧，再发送此帧）；否则就是间接交会，执行3）。
 -   3）若路由表中有目的地址为D的特定主机路由，则把数据传送给路由表中指明的下一跳路由；否则，执行4）。
 -   4）若路由不中有到达网络N的路由，则把数据报传送给路由表中所指明的默认路由器；否则，执行5）。
 -   5）若路由表中有一个默认路由，则反数据报传送给路由表中所指明的默认路由器；否则，执行6）。
 -   6）报告转发分组出错。

 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img19.png)
 
 
## 20.子网掩网产生原因。P128
 
 今天看来ARPANET早期，IP地址设计不够合理。
 
 -   1）IP地址空间的利用率有时很低。A、B类网络占用主机数很太，有些单位会空闲很多IP，但考虑到以后的发展又不会去申请一个占用主机数少的IP。
 -   2）给每一个物理网络分配一个网络号会使路由表太大，网络性能变坏。
 -   3）两级IP不够灵活。比如一个单位需要开通一个新的网络，但在申请到一个新IP之前，新增加的网络不能连接到因特网上工作。  
     解决办法：划分子网。从网络的主机号中若干位作为子网号。  
     原IP地址 ::= {<网络号>,[<主机号>]}  
     含子网IP地址 ::= {<网络号>,[<子网号>,<主机号>]}  
     如下图，不管网络有没有划分子网，只要把子网掩码和IP地址进行逐位“与”运算，就能立即得出网络地址。  
     
  ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img20.png)

## 21.划分子网后，路由器转发分组的算法。P134
 
 -   1）从收到的数据报首部提取目的IP地址D。
 -   2）先判断是否为直接交付。对路由器直接相连的网络逐个进行检查（？相连的网络怎么检查，物理接口吗,不知道这个和查找路由表有什么区别。Ws2013.09.24）：用各网络的子网掩码和D逐位相“与”（AND操作），看结果是否和相应的网络地址匹配。若匹配，则把分组进行直接交付（当然还需要把D转换成物理地址，把数据报封闭成帧发送出去），转发任务结束。否则就是间接交付，执行3）。
 -   3）若路由表中有目的地址为D的特定主机路由，则把数据送给路由表中指定的下一跳路由；否则，执行4）。
 -   4）对路由表中的每一行（目的网络地址，子网掩码，下一跳地址），用其中的子网掩码和D逐位相“与”（AND操作），其结果为N，若N与该行的目的网络地址匹配，则把数据报传送给该行指明的下一跳路由器；否则执行5）。
 -   5）若路由表中有一个默认路由，则把数据报传送给路由器中所指明的默认路由；否则，执行6）。
 -   6）报告转发分组出错。  
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img21.png)
 
## 22.无分类编赴CIDR（构造超网）(为解决IP地址不够用的问题而产生)P35
 
 特点：
 
 -   1）CIDR消除传统的A类、B类和C类地址以及划分子网的概念。  
     记法：IP地址 ：：= ｛《网络前缀》，《主机号》｝  
     斜线记法： 128.14.35.7/20 = 10000000 00001110 00100011 00000111
 -   2）CIDR把网络前缀相同的连续IP地址组成一个”CIDR地址块”。  
     这个地址所在的地址块中的最小地址和最大地址可以很方便地得出。  
     如：  
     最小地址 128.14.32.0 10000000 00001110 00100000 0000000  
     最大地址 128.14.47.255 10000000 00001110 00101111 11111111  
     
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img22.png)
 
## 23.traceroute 命令可跟踪一个分组从源点到终点的路径。(win下是tracert) P143
 
Traceroute发送一个TTL为1的数据报,到达第一个路由器后，路由器收到它，接着把TTL值减1，由于TTL等于0了，路由器就丢弃数据报，并向源主机发送ICMP时间超过差错报告报文。 源主机接着发送第一个数据报P2，将TTL设置为2，P2到达第二个路由器R2后，TTL也会变为0，R2会向源主机发送一个ICMP时间超过差错报文。 当最后一个数据报刚刚到达目的主机时，数据报的TTl是1，主机不转发数据报，也不把TTL值减1.但因IP数据报中封装的是无法交付的运输层UDP数据报，因此目的主机要向源主机发送ICMP终点不可达差错报告报文。 这样通过路由器和目的主机发来的ICMP报文正好给出了源主机想知道的路由信息。  
 每一行有三个时间出现，是因为对应一每一个TTL值，源主机要发送三次同样的IP数据报。  
![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img23.png)
 
## 24.路由选择协议 P144
 
 因特网选用分次的路由选择协议原因：1）网络规模大，若让网络所有路由器都知道所有网络怎样到达，这种路由表非常大，路由信息就会使网络通信链路饱和。2）许多单位不愿意外界了解自己单位网络的布局细节。但同时希望连接因特网。
 
 -   1）内部网关协议(IGP)RIP  
     RIP维护从它自己到每一个目的网络的距离记录（跳数）。RIP选择一条具有最少路由器的路由，而不是最低延时但路由器较多的路由。  
     特点：  
     -   1.仅和相邻路由器交换信息。
     -   2.交换当前路由交换的全部信息，即路由表（我到本自治系统中所有网络的最短距离，以及到每个网络应经过的下一跳路由）。
     -   3.按固定时间间隔交换路由信息。  
         下图是距离向量算法（好消息传播的快，坏消息传播的慢）：  
         ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img24-1.png)
 -   2）内部网关协议(IGP)OSPF(p152)  
     开放最短路径优先。为克服RIP坏消息传播的慢的缺点开发出来的。  
     要点：  
     -   1.向本自治系统中所有路由器发送信息。
     -   2.发送的信息就是与本路由器相邻的所有路由器的链路状态。
     -   3.只有当链路状态发生变化时路由器才用洪泛法发送信息。
 
 OSPF可以得到全网的拓扑结构图。然后每个路由器构造自己的路由表，算出最短路径。（例如用Dijkstra算法）。  
 为了能用于规模很大的网络，OSPF将一个自治系统再划分为若干个更小的范围，区域(200主机数内）。利用洪泛法交换链路状态时，作用范围仅限于区域，以减少通信量。  
 ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img24-2.png)
 
 -   3）外部网关协议(EGP) BGP
     -   边界网关协议BGP是力求寻找一条能够到达目的网络且比较好的路由，并非寻找一条最佳路由。主要原因是：1因特网规模大，使得AS之间路由选择非常困难。各AS内部网络协议也可能不同，同样是代价为1000的路由在不同AS中可能代表不同意义。2AS之间路由选择必须考虑有关策略。比如国内的站点间的数据没必要到国外去兜圈子。
     -   发言人：配置BGP时，每个AS至少一个路由器作为该AS的“BGP发言人”。各AS间的发言人建立连接互换网络可达性信息后，根据策略从收到的路由信息中构造连通图（树形，无加路）找出到达各AS的较好路由。  
         ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img24-3.png)
 
## 25.路由器组成.(P160)
 
 -   1）路由器的结构。
     -   具有多个输入和多个输出端口的专用计算机，任务是转发分组。从路由器输入端口收到分组，按照分组要去的目的地（即目的网络），把分组从路由器合适的输出端口转发给下一跳路由器。
     -   路由选择部分：根据路由选择协议构造路由表。路由表包含从 目的网络 到 下一跳 （IP地址A）的映射。我想去下一个网络N，那么此数据包应该转发到IP地址为A的路由器R。
     -   分组转发：交换结构根据转发表将输入端口进入的分组从一个合适的输出端口转发出去。转发表从路由表得到，包含到达的 目的网络 到 输出端口 和 某些MAC地址信息。  
         ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img25-1-1.png)  
         ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img25-1-2.png)
 -   2）交换结构  
     交结结构将分组从一个输入端口转发到某个合适的输出端口，是路由器的关键构件。  
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img25-2.png)
 
## 28.IP多播
 
 -   1）利用IP多播，可让服务器只发出一个分组数据，就能有多个主机接收到此分组。  
     多播数据报的目的地址写入的不是特定主机的IP地址，而是多播组的标识符，然后设法让加入此多播组的主机的IP地址与多播组的标识符关联起来。  
     多播组的标识符就是IP地址中的D类地址。(224.0.1.0 ~ 238.255.255.255)  
     因为物理层MAC地址只有23位可用作多播，而网络层IP地址可用28位作多播地址，因此收到多播数据报的主机，还要在IP层利用软件进行过滤。  
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img28-1.png)
 -   2）IP多播需要两种协议（网际组管理协议(IGMP)，多播路由选择协议）  
     网际组管理协议：1新主机加入多播组时，主机应向多播组发送IGMP报文，多播路由器也要将这种组成员关系转发给其他路由器。2组成员关系是动态的，多播路由器要周期试探组成员，只要有一个主机对某组响应，则认为这个组是活跃的。  
     多播路由选择协议：1（适用小多播组）洪泛与剪除。反向路径广播，只转发从源点经最短路径传来的广播数据。2（适用多播成员地理位置分散）隧道技术。路由器对多播数据再次封装，加上普通数据报首，成为单播数据。3（适用多播组变化范围大）基于核心的发现技术。（没看懂，略过）
 
## 29.VPN（P171）
 
 -   1）本地地址：对于仅在机构内部使用的计算机可以由本机构自行分配其IP地址，这些仅在本机构有效的IP地址就是 本地地址。
 -   2）全球地址：需要向因特网管理机构申请全球唯一的IP地址，叫做全球地址。
 -   3）专用地址：有时机构内部的某个主机需要和因特网连接，那么这些仅在内部使用的本地地址就可能和因特网中某个IP地址重合，为了解决这一问题，RFC指明了一些 专用地址。  
     专用地址只能用作本地地址，不能用作全球地址，因特网中所有的路由器，对目的地址是专用地址的数据报一律不转发。  
     RFC1918指明的专用地址：
     -   1 10.0.0.0 ~ 10.255.255.255 10/8, 24位块
     -   2 172.168.0.0 ~172.31.255.255 172.16/12, 20位块
     -   3 192.168.0.0 ~ 192.168.255.255 192.168/16， 16位块
 -   4）专用网：采用专用地址的互连网络 称为 专用互联网 或 本地互联网。
 -   5）虚拟专用网（VPN virtual private network）：利用公用因特网作为本机构专用网之间的通信载体，叫VPN。  
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img29.png)
 
## 30.网络地址转换NAT
 
 -   专用网内部一些主机已经分配到了本地IP，但又想和因特网上的主机通信。解决这个问题有两个办法：1设法申请一些全球IP 2采用网络地址转换。
 -   在专用网中添加一个连接到因特网的路由器，路由器至少有一个有效的外部全球IP地址，并在路由器上安装NAT软件，这样所有使用本地地址的主机在和外界通信时都要在NAT路由器上将其本地地址转换成全球IP地址，才能和因特网连接。（这种叫做NAT路由器的东西才是我们一般家庭经常使用的“路由器”吧）。  
     ![](https://wangtiga.github.io/assets/img/2019-06-17-computer-network/img30.png)
 
 # 第五章 运输层（略，未完待续）
 
 ## 31.运输层为应用进程间提供逻辑通信。P180
 
 ## 32.端口号P183
 
 ## 33.UDP用户数据报协议P185
 
 ## 34.TCP传输控制（流）P188
 
 ## 35.可靠传输的工作原理
 
 -   1）停止等待P191
 -   2）连续ARQP193
 -   3）TCP格式P194
 
 ## 36.可靠传输的实现P197
 
 -   1）字节为单位滑动窗口
 -   2）超时重传
 
 ## 37.TCP流量控制P203
 
 -   1）拥塞控制
 -   2）随机早检测
 
 ## 38.TCP连接管理
 
 -   1）三次握手
 -   2）连接释放
 -   3）有限状态机
 
<!--stackedit_data:
eyJwcm9wZXJ0aWVzIjoiZXh0ZW5zaW9uczpcbiAgcHJlc2V0Oi
BnZm1cbiIsImhpc3RvcnkiOlstMTkyNTc1MDExNywyMDE2NTI0
MDQxXX0=
-->
