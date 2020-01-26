> 来源 
> [原文 webrtc-issue-debug](https://blog.codeship.com/webrtc-issues-and-how-to-debug-them/?__s=sgkgganpatrhthvch4js)
> [翻译 webrtc-issue-debug](https://webrtc.org.cn/issues-debug-1/)

作者：Lee Sylvester（[原文链接](https://blog.codeship.com/webrtc-issues-and-how-to-debug-them/?__s=sgkgganpatrhthvch4js)）

翻译：刘通

原标题：WebRTC Issues and How to Debug Them

[![webrtc-logo](http://webrtc.org.cn/wp-content/uploads/2017/08/webrtc-logo-251x300.png)](http://webrtc.org.cn/wp-content/uploads/2017/08/webrtc-logo.png)

WebRTC是由Google的天才工程师创造的一项惊人且颇具开创性的技术。在浏览器之间创建无插件的连接，WebRTC的典型应用是网络视频聊天。

但是，WebRTC不仅仅可以用于音频和视频的传输，它也能够应用于其他高速数据传输。简而言之，它让我们得以窥探端到端游戏，文件传输和其他真正无服务器应用程序的未来。

### **WebRTC****的阴暗面**

我在Xirsys工作时，我看到了一些使用WebRTC技术的非常了不起的应用程序。但是，尽管一些人已经在为下一个伟大的发明努力，但是大多数开发人员还处于起步阶段。因为WebRTC本身就很难！

我们都熟悉典型的Web应用程序。我们有一个客户端，它将数据发送到服务器且从其接收数据；我们还有一个服务器应用，它发送数据给数据库且从其接收数据。许多网络应用的处理过程是线性且可预知的。如果出现问题，我们通常会知道检查哪里，以及它可能发生的原因。

但是，WebRTC，并没有这么简单。

### **异步**

如果你之前编写过多线程应用，那么你可能知道一些会导致问题的地方。竞争条件和受损数据是一对，但是通常大家只是把这些归结于那些难于发现及修复的问题。

WebRTC本质上就是异步的。而且考虑到从远端机器发送和接收数据的情况，它也必须是异步的。这与AJAX调用的不同步性有所不同；相比之下，这些确实很简单。相反的，根据许多这样的调用来考虑，这些调用分布在本地和远端数据不断变化的情况中。

当然，胆小的人不适合来做这件事。

### **主要在客户端**

WebRTC是一个重视客户端的技术。当客户端应用无法正常工作时，通常第一步是询问TURN服务提供商是否有任何日志显示它无法正常工作。

但是需要指出的是，绝大多数WebRTC故障发生在服务器甚至从未被联系过时。至少从服务器的角度来看，没有联系意味着无法登陆。这代表必须挖掘自己的代码来找到答案并知道从哪里开始寻找。

### **NAT****穿越是个雷区**

[![issue2](http://webrtc.org.cn/wp-content/uploads/2018/05/issue2-300x225.jpg)](http://webrtc.org.cn/wp-content/uploads/2018/05/issue2.jpg)

构建Web应用程序通常是一项简单的工作。你写一些会放在Web服务器后面的东西，其中可能有也可能没有服务器端逻辑。一旦写好之后，你就可以部署到您的服务器中去。

在这种情况下，可能发生的最糟糕的事情是你忘了咋爱你IPTable中打开正确的端口。我们都可以做到这些，通常都很简单。但是WebRTC可没有这么容易。

这里的问题是Web服务器本身（硬件而不是软件）是以公开提供数据的目的部署的。他们为此配置了网络硬件并提供公共IP地址。但是，使用WebRTC的目的是连接并发送和接收来自用户机的数据。

这些机器通常情况下，位于旨在保护所述机器免受公众请求的网络上，很有可能不具备公共IP地址，并且经常会引入一些复杂的障碍。

虽然连接到简单的Web服务器与创建HTTP请求一样容易，但WebRTC提供了多种连接类型，我们可以尝试各种连接类型以成功建立连接。


### **WebRTC debug****从何处入手**

有几种调试WebRTC应用的方法，也有一些你需要的重要工具。我们在这里将介绍最重要的工具，同时介绍大多数刚从事WebRTC的开发人员经常遇到的典型问题场景。

然而，在我们太详细讨论之前，先来看看典型的WebRTC连接过程。

### **WebRTC****连接过程**

所有WebRTC连接都需要来自信令协议的一点帮助。这只是需要引入第三方通信协议以允许在主叫方和被叫方之间交换数据的高端说法。

这些数据包括了有关每台机器的浏览器音频和视频功能的信息，以及有关连接到的网络的一些信息。此信息用于配置WebRTC连接选项和流类型。

这种信息共享的典型技术包括COMET和Web Sockets。但是你可以选择使用任何技术。

使用信令技术，以及两台互相连接的用户机，你就可以开始启动WebRTC连接过程了。

提议及应答

WebRTC连接通过使用提议及应答进行初始化。从客观角度看：
```text
# 客户端向对等端发送一个连接提议。
# 对等端收到要连接的提议。
# 对等端发送给客户端一个应答。
# 客户端接收到对等端的应答。
```

这一切更像是我们日常生活的理解。我的意思是，谁会希望在接受提议之前就开启了浏览器通话呢？

但实际上要比这更复杂。

1. 在发送连接提议之前，客户端会创建一个RTCPeerConnection对象的实例，并使用以下命令生成SDP数据包：rtcPeerConnection.createOffer()。这会将客户的机器和浏览器的功能细化。

2. 然后通过RTCPeerConnection对象的setLocalDescription方法将SDP数据包设置为客户端的本地描述。

3. 然后将SDP数据包发送给对等端，对等端通过其RTCPeerConnection对象实例的setRemoteDescription方法将其设置为远程描述。

4. 对等端然后使用RTCPeerConnection对象的createAnswer方法生成自己的SDP数据包，并将其设置为其本地描述。

5. 对等端将SDP数据包作为应答响应发送给客户端，并将其设置为其远程描述。

这个时候，客户端和对等端就都知道对方的会话能力了。

### **连接性**

为了创建连接，WebRTC使用所谓的STUN（Session Traversal Utilities for NAT）和TURN（Traversal Using Relay around NAT）。

这些听起来相当复杂，但实际上是非常简单的协议，旨在创建两个候选之间的连接。

STUN服务器

第一个要说的是STUN。当机器希望公布自己的连接性时，它们需要通知连接器自己的公共IP地址。毕竟，典型用户计算机通常没有自己的公共域名。

但问题是，机器的本地网络地址与公共网络地址有很大的不同。这是NAT的一部分，用于将传入和传出数据包的IP地址转换为本地和公共IP地址。

为了解决这个问题，用户的机器向STUN服务器发出请求。当请求数据包通过本地NAT时，数据包报头中的本地IP地址会被转换为公共IP地址。

然后，STUN服务器会接收到这个数据包，将其报头中的公共IP地址复制到它的主体中，然后将其发回给发送者。然后数据包通过用户的NAT，从而将报头IP转换回本地IP地址。

由于数据包主体保持不动，机器现在能够识别其公共IP，并将其发送给希望建立连接的对等端。

[![issue21](http://webrtc.org.cn/wp-content/uploads/2018/05/issue21-300x175.png)](http://webrtc.org.cn/wp-content/uploads/2018/05/issue21.png)

TURN服务器

TURN服务器是作为STUN协议的扩展而构建的。使用相同的报头结构，但是提供了额外的功能，被称为commands。

TURN服务器作为代理服务器被使用。客户端使用allocation（分配）建立到对等端的安全连接。此分配只是对等端可以连接到的服务器上的协商UDP端口，用于侦听来自客户端的流量。一旦在服务器上建立了客户端和对等端之间的连接，数据就可以通过两种的方式进行自由传递。

TURN服务器的设计方式使客户端有更多的权力和机会发送数据，而并非对等端。处于这个原因，并且取决于本地网络配置，客户端/对等端连接可能比其他方式更为成功。

[![issue22](http://webrtc.org.cn/wp-content/uploads/2018/05/issue22-300x184.png)](http://webrtc.org.cn/wp-content/uploads/2018/05/issue22.png)

### **Debug**

你已经阅读到这里了。这篇文章本应是只关于调试WebRTC连接性的，但你可能已经意识到了，也会涉及很多可能导致问题的地方。除非你非常的幸运，否则是绝对会出错的。

进行WebRTC连接相关的工作时会有许多潜在问题。第一个就是根本无法进行连接。在这种情况下，只要在本地按照连接过程工作就是一个良好的开端。

### **调试你本地会话设置**

正如我前面提到的，WebRTC是一个客户端巨兽。建立ICE连接的过程基本上全都是在浏览器中完成的。STUN和TURN服务器相对而言就非常简单了，因此大多数的问题都是你自己导致的。

但是不要灰心，这会刺激你尽最大努力来解决这些烦人的错误。

在寻找连接失败的根本原因时，首先应该检查通过信令服务器的数据。需要记住的是，媒体和网络会话描述数据包都是通过这里发送的，所以你手头上有许多数据来测试问题出在哪里。

最明显的着手点是确定哪些数据包已被发送和接收。注意，提议和应答需要从一个机器发送给另一个才能进行连接。
``` text
# 对等端收到提议了吗？客户端收到应答了吗？如果二者有一个没有时候到的话，那么连接就无法成功建立。

# 同样的，ICE候选数据包是否生成了？它们被分享了吗？各方是否在addIceCandidate中添加了另一方的数据包？

# 如果所有数据包都已被发送和接收了，那么检查远程流是否通过RTCPeerConnection中的onaddstream处理程序添加到HTML视频元素中？之前我确实没有提到过，但我确信在构建应用程序的时候一定要查看它。
```

如果以上几条检查过了都没有问题，那么就需要对会话数据进行深入排查了。

### **会话描述协议**

所有用于建立连接的数据都是SDP格式的。当你第一次看到它的时候可能会有些困惑，但只要给你一点帮助，它就可以成为一个非常有启发性的资源。

ICE候选SDP数据包更重要的一个方面是typ参数。对于WebRTC，这可以是以下三种选择之一：

```text
# typ host

# typ srflx

# typ relay
```

#### **typ host**

host类型表示与本地网络上设备的连接。通常情况下，所有host连接都不会使用STUN或TURN。原因在于，由于连接设备都在同一网络上，所以根本不需要本地到公共IP地址的转换，因此可以直接进行连接。

如果你希望连接到同一网络上的机器，则两台机器会共享host SDP数据包。如果连接仍然没能建立，那么你可以就认为你的代码在某个地方出现了大量的错误。

需要注意的是，仅仅是因为两台机器位于同一个网络中，它不会假设可以建立直接连接。host SDP ICE候选必须由双方提供。我见到过许多机器必须需要TURN连接才能与它自己相连。

#### **typ srflx**

srflx是服务器反身性（Server Reflexive）的缩写，用于表示获取公共IP地址的术语。因此，srflx代表只需要STUN的对等连接。当双方都提供了srflx数据包的时候，这意味着双方都应该通过STUN-only设置变成可连接的，但并不意味着双方可以通过这种方式实现相互连接。

但是，他们很有可能可以连接。

#### **typ relay**

relay用于描述TURN的连接性。当双方都提供这样的数据包时，那么连接是绝对可能进行的。

上面列出的数据包类型不是连续的，也就是说，设备可以提供srflx数据包，但不提供relay数据包，反之亦然。当双方都不能提供匹配的数据包类型时，连接根本就不可能建立。



### **测试设备连接性**

我们可以以清晰可读的形式来描述设备网络和媒体的功能。Google提供了一个测试工具，你可以在[https://test.webrtc.org/](https://test.webrtc.org/)上面使用它。

只需点击“Start”按钮即可以获取机器的完整描述。Xirsys在其仪表板中提供了相同的工具，但其被配置为使用STUN和TURN服务器而不是Google的服务器。

**WebRTC Internals**

现在，你已经完成了上文中所讲的所有内容，但仍然无法弄清楚为什么还是无法建立连接？值得庆幸的是，Chrome提供了额外的工具来帮助调试你的连接性，以及一些图表来帮助你。这些都是通过Chrome的WebRTC Internal功能实现的。

想要使用WebRTC Internal的话，只需要打开一个新的标签页并输入以下协议和URL就可以了：chrome://webrtc-internals。

如果你已经有一个WebRTC应用正在运行的话，应该会立即看到一堆数据。否则，就需要先在另外一个选项卡中运行你的WebRTC应用程序。

Internals应用程序提供了对RTCPeerConnection实例的API调用的细分，以及来自任何getUserMedia实例的信息。后者不提供大量信息，但前者是调试应用程序的宝贵工具。

### **ICE****配置**

在Internals页面的顶部，你会看到用于创建连接的ICE字符串。如果你没有在应用配置中正确地提供ICE数据，那么你会立刻在此处的数据中注意到这点。

当然，缺乏ICE数据或者没有TURN服务器配置的ICE数据可能会导致问题。如果是这种情况，请务必追踪你的设置并确保ICE字符串在开始提议/应答流程之前提供。

[![issue41](http://webrtc.org.cn/wp-content/uploads/2018/05/issue41-300x180.png)](http://webrtc.org.cn/wp-content/uploads/2018/05/issue41.png)

### **RTCPeerConnection****事件**

下一组数据将按照请求的顺序由RTCPeerConnection实例的事件和函数调用组成。 任何这些请求的失败将以红色突出显示。

如果您看到一个红色突出显示的addIceCandidateFailed事件，请不要太担心。 这种错误可能会发生，但仍然可能会成功连接。

请注意，你需要查找连接失败的线索。由于每个时间都是按顺序发生的，因此你会对可能导致故障的步骤有一个大致的概念，从而找到代码或运行环境中的错误。

如果连接成功，则应在列表末尾存在值为complete的iceconnectionstatechange事件。 这是你希望达到的结果，表明一切进展顺利。

### **统计数据列表**

下一组数据与您的连接统计数据相关。 当连接成功并提供连接内的延迟和数据传输时，这将非常有用。

两个最有用的统计数据是ssrc和bweforvideo数据。

ssrc，表示流源（Stream Source），提供给每个音频和视频媒体轨道。它详细介绍了相关轨道的吞吐量，并包含了诸如RTCP往返时间等有用的细节。

bweforvideo数据是对于对等连接的带宽估计报告。

[![issue42](http://webrtc.org.cn/wp-content/uploads/2018/05/issue42-300x179.png)](http://webrtc.org.cn/wp-content/uploads/2018/05/issue42.png)

### **getStats****函数**

很多时候无法访问WebRTC Internals页面，比如应用程序用户遇到错误时。

在这种情况下，可以通过RTCPeerConnection对象的getStats函数和记录对象的各种处理程序来获取由WebRTC Internals提供的相同事件数据。getStats函数接受回调处理程序，并为其提供一个详细对象，列出WebRTC Internals接口中存在的每个统计值。

```javascript
rtcPeerConnection.getStats(function(stats) {
    document.getElementById("lostpackets").innerText = stats.packetsLost;
});
```

RTCPeerConnection的oniceconnectionstatechange处理程序是一个非常有用的工具。 此处理程序将在连接发生更改时收到连接的条件，并且可能是以下任一项：
```text
# new — WebRTC引擎正在等待通过调用RTCPeerConnection.addIceCandidate（）来接收远程候选对象。

# checking – WebRTC引擎已收到远程候选，并正在比较本地和远程候选以尝试找到合适的匹配。

# connected – 已经确定了一对适当匹配的候选并建立了连接。根据Trickle ICE协议，候选项仍可以继续共享。

# completed – WebRTC引擎已经完成了收集候选项，已经检查了所有候选对，并找到了所有组件的连接。

# failed – WebRTC引擎查找了所有候选对，但是未能找到合适的匹配。

# disconnected –RTCPeerConnection中至少有一个轨道已断开连接。这可能会间歇性地触发并在不可靠的网络上自行解决。在这种情况下，连接状态可能会改回“已连接”。

# closed –RTCPeerConnection实例已关闭，并且不再处理请求。
```

通常，如果状态变为failed，可能需要循环两端机器上的每个通过的ICE候选包，以辨别失败发生的原因；例如，如果一方只提供host和srflx数据包，而另一方提供host和relay数据包，但双方都在不同的网络上。

### **遇到空白视频窗口**

通常，用户可能会遇到音频正常，但是其中一个用户，或者两个用户的视频窗口都是黑屏的连接问题。

现在，我们将用逻辑分析来确定这是一个本地问题。由于连接本身不知道在它之间传输的数据，并且由于音频明显从一个设备流向另一个设备，因此该问题必须与分配给视频组件的流相关联。

在这种情况下，原因通常是在连接状态准备好之前将视频流分配给视频标签的结果。确保只有在RTCPeerConnection实例的状态显示为completed后才会分配流。

### **外部问题解决**

除了RTCPeerConnection API和WebRTC Internals，寻找连接问题的另一个有用工具是通过使用网络数据包抓包分析器，例如Wireshark。

在WebRTC连接时运行Wireshark，将在Wireshark主窗口中记录STUN协议数据包。你可以在筛选字段中输入stun，然后按回车键来过滤这些数据包。

[![issue43](http://webrtc.org.cn/wp-content/uploads/2018/05/issue43-300x181.png)](http://webrtc.org.cn/wp-content/uploads/2018/05/issue43.png)

在记录连接数据时，通常首先要确定连接类型。如果只记录Binding请求和响应，则只有srflx或STUN连接请求会被记录下来。另一方面，TURN连接将记录STUN Binding请求和TURN特定请求，例如Allocation和CreatePermission数据包。

对于TURN连接，请检查是否提供了成功的Allocation响应。如果所有Allocation响应都显示为失败，则可能是ICE字符串中提供的凭证无效或过期。试着更新凭证并再次尝试。

如果您在Wireshark捕获中看到CreatePermission Success Response数据包，那么通常可以放心的假定一切都顺利。如果看到了ChannelBind数据包就更好了，这表明高速TURN连接已建立。

### **移动数据连接问题**

不幸的是，在移动设备上创建WebRTC连接可能会有些挫败。大量的支持请求已经成为我的烦心事，这些应用可以在WiFi上正常工作，但在3G或4G上就会完全失败。

调试此类应用程序可能会很痛苦，因为您无法在移动设备上使用Wireshark等应用程序，并且移动Safari无法提供像WebRTC Internals这样的好界面。

在这种情况下，最好停下手上的工作并且好好考虑一下。如果该应用在WiFi上运行良好，则不太可能是应用本身的问题，而是移动服务提供商的问题。由于移动设备上调试WebRTC应用程序并不是一件有意思的事，所以你可以考虑花钱使用移动网络安全装置，就像[这个](https://www.amazon.co.uk/gp/product/B06XC16QC1/ref=as_li_qf_asin_il_tl?ie=UTF8&tag=designrealm-21&creative=6738&linkCode=as2&creativeASIN=B06XC16QC1&linkId=b4dc0bfe06738f54ccd1e16e0f7a4bdd)。

这里的要点是尝试使用电脑端技术来模拟移动环境，以便利用您可以使用的调试工具。使用上述设备，您甚至可以在4G接收区强制使用3G，以查看可能发生的不同连接结果。

### **结语**

虽然WebRTC是个很难对付的家伙，但有很多可能的途径可以找到应用程序中令人讨厌的错误，许多有用的技巧和技巧的博客以及提供免费支持的公司。

另外，我非常鼓励你阅读STUN和TURN的RFC规范，以及WebRTC本身的规范。这些文件可能有点难以实现，但通过了解或者只了解其中的一小部分内容，就有助于使你的工作变得更加简单。


<!--stackedit_data:
eyJkaXNjdXNzaW9ucyI6eyJNU2VpYzRub2pRdGN4czJWIjp7In
N0YXJ0Ijo4MTg0LCJlbmQiOjgyMzAsInRleHQiOiLnoa7kv53l
j6rmnInlnKhSVENQZWVyQ29ubmVjdGlvbuWunuS+i+eahOeKtu
aAgeaYvuekuuS4umNvbXBsZXRlZOWQjuaJjeS8muWIhumFjea1
geOAgiJ9fSwiY29tbWVudHMiOnsiUHhaamd4ZzVJcWJFSzlLQi
I6eyJkaXNjdXNzaW9uSWQiOiJNU2VpYzRub2pRdGN4czJWIiwi
c3ViIjoiZ2g6Mzk3NjcwMzIiLCJ0ZXh0Ijoi5oiR5Lus55qEYn
Vn5LiO5q2k5Y+v6IO95pyJ5YWz77yM5L2G5aaC5p6c56Gu6K6k
5ZGi77yaXG7nnIvku6PnoIHvvJ9cbuWGmWRlbW/pqozor4HvvJ
8iLCJjcmVhdGVkIjoxNTYxMzcwODczODQ4fX0sImhpc3Rvcnki
OlstMTc0NzI1NTc3MV19
-->