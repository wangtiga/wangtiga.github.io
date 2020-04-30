---
layout: post
title:  "[译]How Zoom's web client avoids using WebRTC (DataChannel Update)?"
date:   2020-04-29 12:00:00 +0800
tags:   todo
---


## How Zoom's web client avoids using WebRTC (DataChannel Update)? [^webrtcInZoomEN] [^webrtcInZoomCN]

## Zoom 的 web client 如何避免使用 WebRTC ？ [^webrtcInZoomEN] [^webrtcInZoomCN]



Posted by  [Philipp Hancke](https://webrtchacks.com/author/philipp-hancke/ "View all posts by Philipp Hancke")  on  [September 8, 2019](https://webrtchacks.com/zoom-avoids-using-webrtc/)

**Posted in:**  [Reverse-Engineering](https://webrtchacks.com/category/reverse-engineering/).

**Tagged:**  [Blackbox Exploration](https://webrtchacks.com/tag/blackbox-exploration/),  [DataChannel](https://webrtchacks.com/tag/datachannel/),  [webassembly](https://webrtchacks.com/tag/webassembly/),  [webrtc-nv](https://webrtchacks.com/tag/webrtc-nv/),  [websocket](https://webrtchacks.com/tag/websocket/),  [zoom](https://webrtchacks.com/tag/zoom/).

[3 comments](https://webrtchacks.com/zoom-avoids-using-webrtc/#comments)

_Editor?s Note: This post was originally published on October 23, 2018. Zoom recently started using WebRTC?s DataChannels so we have added some new details at the end  [in the DataChannels section](https://webrtchacks.com/zoom-avoids-using-webrtc/#datachannels)._

![](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/rube-goldberg-cartoon-og.jpg)

Rube Goldberg?s Professor Butts and the Self-Operating Napkin (1931)

[Zoom](https://zoom.us/)  has a web client that allows a participant to join meetings without downloading their app.  [Chris Koehncke](https://twitter.com/ckoehncke)  was excited to see how this worked (watch him at the upcoming  [KrankyGeek event](https://www.krankygeek.com/)!) so we gave it a try. It worked, removing the download barrier. The quality was acceptable and we had a good chat for half an hour.

可以通过 Web 客户端直接参与 Zoom 视频会议，不需要下载安装 App 。
Chris Koehncke 对于这个客户端的工作原理十分感兴趣。
我在这里尝试分析一下。
使用 Web 客户端时，完全避免下载安装 App 的麻烦。
我们使用了大概半小时，质量还是比较满意。


Opening  `chrome://webrtc-internals`  showed only  getUserMedia being used for accessing camera and microphone but no RTCPeerConnection like a WebRTC call should have. This got me very interested ? how are they making calls without WebRTC?

打开 `chrome://webrtc-internals` 页面，可以发现，仅仅调用了 getUserMedia 接口访问摄像头和麦克风，却没有通过 RTCPeerConnection 进行 WebRTC 呼叫。
这让我很好奇，它们不用 WebRTC 是怎么实现通话功能的？

[![](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-gum.png)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-gum.png)



## Why don't they use WebRTC? 为什么他们不使用 WebRTC ？

The relationship between Zoom and WebRTC is a difficult one as shown in this statement from their website:  
[![](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-webrtc-bashing-1024x247-1024x247.png)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-webrtc-bashing-1024x247.png)

TODO 

The Jitsi folks just did a  [comparison of the quality](https://jitsi.org/news/a-simple-congestion-test-for-zoom/)  recently in response to that. [Tsahi Levent-Levi had some useful comments](https://bloggeek.me/webrtc-vs-zoom-video-quality/)  on that as well.

Jitsi 近期做了一个质量对比。
Tsahi Levent-Levi 对此也有一些评价。


So let?s take a quick look at the ?excellent features? under quite interesting circumstances ? running in Chrome.

我们赶快人看看这么优秀的功能是如何在 chrome 中运行的吧。


## The Zoom Web client

Chromes network developer tools quickly showed two things:

Chrome 开发者工具的 network 标签显示以下两件事：

-   [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)  are used for data transfer
-   there are workers loading  [WebAssembly](https://webassembly.org/)  (wasm) files

- 使用 WebSockets 传输数据
- 这其中加载了很多 WebAssembly 文件

[![](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-workers.png)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-workers.png)

The WebAssembly file names quickly lead to a  [GitHub repository](https://github.com/zoom/sample-app-web/tree/master/lib/av)  where those files, including some of the other JavaScript components are hosted. The files are mostly the same as the ones used in production.

从 WebAssembly 文件名，可找到一个 [GitHub 仓库](https://github.com/zoom/sample-app-web/tree/master/lib/av) ，其中还包含很多 JavaScript 组件。仓库中的文件代码与生产环境中使用的基本一致。


### Media over WebSockets 通过 WebSocket 传输媒体数据

The overall design is quite interesting. It uses WebSockets to transfer the media which is certainly not an optimal choice. It is similar to using TURN/TCP in WebRTC ? it has a quality impact and will not work well in quite a number of cases. The general problem of doing realtime media over TCP is that packet loss can lead to resends and increased delay. Tsahi  [described this over at TestRTC a while ago](https://testrtc.com/happens-webrtc-shifts-turn-tcp/), showing the impact on bitrate and other things.

这个设计很有趣。
使用 WebSocket 传输媒体数据肯定不是最佳选择。
这有点像是在 WebRTC 中使用 TURN/TCP ？
这会影响很多场景下的使用效果。
使用 TCP 传输实时音视频数据的常见问题是，丢包会导致重传和越来越高的时延。
Tsahi 最近的一篇文章描述过这种用法对 bitrate 及其他方面的影响。


The primary advantage of using media over WebSockets is that it might pass firewalls where even TURN/TCP and TURN/TLS could be blocked. And it certainly avoids the issue of WebRTC TURN connections not getting past authenticated proxies. That was a  [long-standing issue](https://bugs.chromium.org/p/chromium/issues/detail?id=439560)  in Chrome?s WebRTC implementation that was only resolved last year.

使用 WebSocket 传输媒体数据的主要优势是，它能轻松穿越防火墙，而使用 TURN/TCP TURN/TLS 时常常被这些防火墙屏蔽。
而且，使用 WebSocket 还能避免，在使用认证代理时引发的 WebRTC TURN 连接问题。
[这里](https://bugs.chromium.org/p/chromium/issues/detail?id=439560)就描述了一个 Chrome WebRTC 去年才修复的一个问题。


Data received on the WebSockets goes into a WebAssembly (WASM) based decoder. Audio is fed to an AudioWorklet in browsers that support that. From there the decoded audio is played using the WebAudio ?magic? destination node.  

从 WebSocket 接收的数据会送入基于 WebAssembly (WASM) 实现的解码器。
音频送到浏览器的 AudioWorklet 中。
解码后的音频使用 WebAudio 播放。

TODO AudioWorklet WebAudio 都是 Chrome 浏览器本身提供的组件吗？ WASM 仅执行解码过程，播放还是由调用浏览器本身的组件执行？


[![](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-flow2.png)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-flow2.png)  

Video is painted to a canvas. This is surprisingly smooth and the quality is quite high.

视频由 canvas 绘制。画面十分顺滑，而且质量也很高。

In the other direction, WebAudio captures media from the  getUserMedia call and is sent to a WebAssembly encoder worker and then delivered via WebSocket. Video capture happens with a resolution of 640?360 and, unsurprisingly, is grabbed from a canvas before being sent to the WebAssembly encoder.

其他方面，
WebAudio 通过 getUserMedia 接口获取媒体数据，然后送到 WebAssembly 编码器编码后通过 WebSocket 发送。
显然，
通过 canvas 捕获的 640x360 分辨率的视频，然后送到 WebAssembly 编码器编码。

TODO getUserMedia 捕获到的视频放到 canvas 显示，音频送到 WebAudio 组件播放，然后 js 代码再从这两种组件提取媒体流数据，送到 WASM 编码？

[![](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-flow1.png)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-flow1.png)

The WASM files seem to contain the same encoders and decoders as Zooms native client, meaning the gateway doesn?t have to do transcoding. Instead it is probably little more than a websocket-to-RTP relay, similar to a TURN server. The encoded video is somewhat pixelated at times and Mr. Kranky even complained about staircase artifacts. While the CPU usage of the encoder is rather high (at 640?360 resolution) this might not matter as the user will simply blame Chrome and use the native client the next time.

WASM 文件中包含的编解码器与 Zoom Native Client 中一样。
所以媒体网关完全无需转码，它工作过程更像是 TURN 服务器，只是进行 websocket 到 RTP 的重传。

TODO 如何翻译 The encoded video is somewhat pixelated at times and Mr. Kranky even complained about staircase artifacts.

即使编码器占用的CPU太高，也没关系，用户只会报怨一下 Chrome 浏览器，然后下次直接使用  native client 。


### H.264

Delivering the media engine as WebAssembly is quite interesting, it allows for supporting codecs not supported by Chrome/WebRTC. This is not entirely novel ? [FFmpeg](https://www.ffmpeg.org/) compiled with  [emscripten](https://github.com/kripken/emscripten) has been done many times before and emscripten seems to have been used here as well. Delivering the encoded bytes via WebSockets allowed inspecting their content using Chrome?s excellent debugging tools and showed a H264 payload with an RTP header and some framing:

raw RTP data sent over the Websocket

将媒体引擎转换成 WebAssembly 很有趣。
这样就能编码 Chrome/WebRTC 不支持的编码格式了。
这够不够新奇？
使用 mscripten 编译的 FFmpeg 早就出现了。
这里似乎使用的就是 emscripten 。
使用 Chrome 的 debug tool 能看到通过 WebSocket 发送的字节流数据就是 RTP header + H264 payload :

通过 WebSocket 发送的就是原始 RTP 数据。


```txt
1

2

3

02000000

9062ae85bb9c9d7801000401bede0004124000003588b8021302135000000000

1c800000016764001eac1b1a68280bde54000000  ...
```

To my surprise, the  [Network Abstraction Layer Unit (NALU)](https://tools.ietf.org/html/rfc6184)  did not indicate  [H264-SVC](https://tools.ietf.org/html/rfc6190).

比较奇怪的是 Network Abstraction Layer Unit (NALU)  中没有包含 H264-SVC 。



## Comparison to WebRTC 与 WebRTC 对比

As a summary, let us compare where what Chrome uses in this case is different from the WebRTC standard (either the W3C or the various IETF drafts):

标准 WebRTC 与上述做法进行一个简单的比较：

**Feature** | **Zoom Web client** | **WebRTC/RTCWeb Specifications** |
----------- | ------------------- | -------------------------------- |
Encryption       | plain RTP over secure Websocket | DTLS-SRTP |
data channels    | n/a?                            | SCTP-based |
ICE              | n/a for Websockets              | RFC 5245 (RFC 8445) |
Audio codec      | unknown (yet)                   | Opus |
Multiple streams | not observed (yet)              | Chrome implements the spec finally  |
Simulcast        | not observed in the web client  | extension spec |



## WebRTC, Next Version

Even though WebRTC 1.0 is far from done (and most developer are still using something that is dubbed the ?legacy API?) there is a lot of discussion about the ?next version?.  

虽然 WebRTC 1.0 还远未完成（而且很多开发者仍然在使用一些 legacy API ) ，但有关下一个版本 WebRTC 的讨论声音已经出现很多了。


The overall design of the Zoom web client strongly reminded me of what Google?s Peter Thatcher presented as a proposal for WebRTC NV at the Working groups face-to-face meeting in Stockholm earlier this year. See  [the slides (starting on page 26)](https://www.w3.org/2011/04/webrtc/wiki/images/5/5c/WebRTCWG-2018-06-19.pdf).

有关 Zoom web client 的设计让我想起 Google 的 Peter Thatcher 在今年早些时候于 Stockholm 进行的以 WebRTC NV 为主题的展示会。
查看  [幻灯片 (从 26 页开始)](https://www.w3.org/2011/04/webrtc/wiki/images/5/5c/WebRTCWG-2018-06-19.pdf) 了解详情。

TODO 了解 WebRTC NG 2018 06 19 演讲。


If we were to rebuild WebRTC in 2018, we might have taken a similar approach to separate the components. Basically taking these steps:

-   compile webrtc.org encoders/decoders to wasm
-   connect decoders with canvas and WebAudio for ?playout?
-   connect encoders with  getUserMedia for input
-   send encoded media via (unreliable) datachannels for transport
-   connect  RTCDataChannel feedback metrics to audio/video encoder somehow


如果在 2018 年重新设计 WebRTC ，我可能会对以下组件重新拆分。基本步骤如下：

- 将 webrtc.org 的编码解码器编译为 wasm
- 支持解码器输出到 canvas WebAudio 播放
- 支持编码器从 getUserMedia 获取输出
- 支持通过任意 datachannel 作为 transport 发送编码后的媒体流
- 将 RTCDataChannel 反馈指标与 audio/video 编码器以某种方式连接


The approach is visualized on one of the slides from the working groups meeting materials:  

可以在 working groups meeting 的幻灯片资源中观看此种方式的图示：

[![](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/webrtc-pthatcher-1024x106.png)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/webrtc-pthatcher.png)

That proposal some has very obvious technical advantages over the Zoom approach. For instance, using  RTCDataChannels for transferring the data which gives a much much better congestion control properties than WebSockets, in particular if there is packet loss.

这上方法比 Zoom 的明显更有技术优势。
例如，使用 RTCDataChannels 传输数据比 WebSockets 提供了更好的拥塞控制能力，尤其是在丢包的情况下。


The big advantage of the design is that it would be possible to to separate the Encoder and Decoder (as well as related things like the RTP packetization) from the browser, allowing custom versions. The main problem is finding a good way to get the data processing off the main thread in a high-performance way including hardware acceleration. This has been one of the big challenges in Chrome in the early days and I remember many complaints about the sandbox making things difficult. Zoom seems to work ok but we only tried a 1:1 chat and the typical WebRTC app is a bit more demanding than that. Reusing building blocks like  MediaStreamTrack  for data transfer from and to workers would also be preferable to fiddling with Canvas elements and WebAudio.

该设计最大优点是，从浏览器中分离了编码器和解码器（还有打包 RTP 的过程），这样更容易使用自定义编解码器。
最重要的问题是如何以更高性能的方式将处理媒体数据的过程从主线程分享出来，包括硬件加速。
这是 Chrome 浏览器早期面临的一大挑战，我记得有许多人抱怨沙盒使得提高性能的方法复杂了很多。
Zoom 在这方面表示还不错，但是我们只尝试了一对一的视频聊天场景，而典型的 WebRTC 应用程序要求一般比这种高一些。
重复利用 MediaStreamTrack 这种基础组件进行数据传输可要比摆弄 Canvas 和 WebAudio 这些小玩意要好多了。


However, as that approach would have allowed  [building Hangouts in Chrome](https://webrtchacks.com/hangout-analysis-philipp-hancke/)  without having to open source the underlying media engine I am very glad WebAssembly was not a thing in 2011 when WebRTC was conceived. So thank you Google for open sourcing  [webrtc.org](https://webrtc.org/)  under a  [three clause BSD license](https://chromium.googlesource.com/external/webrtc/+/master/LICENSE).

无论如何，这种方法可以在 Chrome 中构建视频会议程序，而且不用开源底层媒体引擎。
我很高兴 2011 年 WebRTC 诞生时， WebAssembly 还不存在。
所以感谢 Google 以 three clause BSD license 授权将 webrtc.org 开源。


> TODO WebAssembly 在 2011 年没有出现，有什么可高兴的？


## Update September 2019: WebRTC DataChannel 2019 09 月更新: WebRTC DataChannel

As Mozilla?s  [Nils Ohlmeier](https://twitter.com/nilsohlmeier)  pointed out, Zoom switched to using WebRTC DataChannels for transferring media:

Nils Ohlmeier 指出， Zoom 已经改改用 WebRTC DataChannel 传输媒体数据了。

> Looks like  [@zoom_us](https://twitter.com/zoom_us?ref_src=twsrc%5Etfw)  has switched it's web client from web sockets to  [#WebRTC](https://twitter.com/hashtag/WebRTC?src=hash&ref_src=twsrc%5Etfw)  data channels. Performance a lot better compared to their old web client.  [pic.twitter.com/SQhP9XhHXP](https://t.co/SQhP9XhHXP)
> 
> ? Nils Ohlmeier (@nilsohlmeier)  [September 5, 2019](https://twitter.com/nilsohlmeier/status/1169691591922946049?ref_src=twsrc%5Etfw)

> 看来@zoom_us已经把它的Web客户端从web sockets改成了WebRTC data channel 。性能比旧的web客户端要好很多。
> 
> Nils Ohlmeier (@nilsohlmeier)  September 5, 2019

chrome://webrtc-internals  tells us a bit more about how this works. See a  [dump here](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-datachannel.txt)  which can be  [imported with my tool here](https://fippo.github.io/webrtc-dump-importer/).

打开 `chrome://webrtc-internals` 能看到其中的工作原理。
这里有一个[dump文件](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-datachannel.txt)，可以用[我的工具](https://fippo.github.io/webrtc-dump-importer/)导入dump文件。



### No STUN/TURN = fallback to WebSockets

Looking at the  RTCPeerConnection  configuration the most notable thing is that  iceServers  is configured as an empty array which means no STUN or TURN servers are used:

查看 RTTCPeerConnection 配置可以发现 iceServers 是一个空数组，也就是说这里没有使用 STUN 和 TURN 服务器。

[![Zoom datachannel peerconnection configuration](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-dc.png)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2018/10/zoom-dc.png)

This makes sense as TCP fallback is most still likely provided by the previous WebSocket implementation and no TURN infrastructure is required just for the web client. Indeed, when UDP is blocked the client falls back to WebSockets and is trying to establish a new RTCPeerConnection every ten seconds.

如果 web client 中不使用 TURN 服务器，那么 TCP fallback 很可能是利用之前 的 WebSocket 方案实现。
也就是说， UDP 被禁用时，客户端会继续通过 WebSocket 传输媒体数据，并且每过 10s 会尝试建立 RTCPeerConnection 。

> NOTE 如果保证 server 端有公网IP， UDP ICE 打洞理论上不会失败。只有客户网络中完全禁用 UDP 协议时，才会打洞失败。
> 很多公司网络都会配置策略禁用 UDP 协议，所以要基于 WebRTC 打造产品级的会议服务，必须提供 TCP 参会的方式。所以Web 端最好的方式还就是 WebSocket 了。

> NOTE 一般只有 音频 视频 游戏 P2P下载 等业务需要用UDP ，为保障网络通道，所以很多网管人员就直接禁用UDP了。
> 当然 DNS 53 端口 这类标准服务肯定会放行。


### Standard PeerConnection setup but with SDP munging 标准的 PeerConnection设置，但带有SDP munging 功能。

There are two PeerConnections, each of which creates an unreliable DataChannel, one labelled  ZoomWebclientAudioDataChannel, the other labelled  ZoomWebclientVideoDataChannel. Using two connections for this isn?t necessary but was probably easier to fit into the existing WebSocket architecture.

有两个 PeerConnection 都使用不可靠连接，一个标记为 ZoomWebclientAudioDataChannel ， 另一个标记为 ZoomWebclientVideoDataChannel 。
其实没必要用两个连接，可以是他们为兼容 WebSocket 中的实现方案吧。

> NOTE webrtc 可用 bundle 方式在一个连接中传输多路媒体，并且配合使用 rtcp-mux 还能在一个连接中同时包含 RTP RTCP 协议。

After that,  createOffer  is called followed by  setLocalDescription. While this is pretty standard there is an interesting thing going on here as the  a=ice-frag  line is changed before  setLocalDescription  is called, replacing Chrome?s rather short  [username fragment  _(ufrag_)](https://developer.mozilla.org/en-US/docs/Web/API/RTCIceCandidate/usernameFragment)  with a lengthy  _uuid_. This is called  _[SDP munging](https://webrtcglossary.com/sdp-munging/)_  and is generally frowned upon due potential interoperability issues. Surprisingly Firefox allows this which means more work for Nils and his Mozilla team. We?ll discuss what the purpose of this might be below.

随后，调用  createOffer 和 setLocalDescription  。
这些都是标准调用流程，没什么新鲜的。
但有一点特殊的是，在 setLocalDescription 前 a=ice-frag 行被改变过。
他们用了一个很长的 uuid 代替 chrome 中原本很短的 [username fragment  _(ufrag_)](https://developer.mozilla.org/en-US/docs/Web/API/RTCIceCandidate/usernameFragment) 。
这就是 _[SDP munging](https://webrtcglossary.com/sdp-munging/)_  ，但一般是被禁用的，因为有潜在的互操作性问题。
令人惊讶的是， Firefox 允许这样做，也就是说 Nils 和他的 Mozilla 团队为此要做更多工作。我们一会会讨论下这样做的目的。

> NOTE 互操作性问题，不同浏览器行为不一致，可能有的支持 SDP munging ，但有得就不支持了。可在些理解成兼容性。

### Simple server setup 简单的服务器设置 

Next we see the answer from the server. It is using  _[ice-lite](https://webrtcglossary.com/ice-lite/)_  which is common for servers as it is easy to implement as well as a single host candidate. That candidate is then also added via  addIceCandidate  which is a bit superfluous but explains the double occurrence in Nils screenshot. The answer also specifies  a=setup:passive  which means the browser is acting as the DTLS client, probably to reduce the server complexity.

下面我们看看服务端回复的 answer 。
服务端使用的是 _[ice-lite](https://webrtcglossary.com/ice-lite/)_ ，这很常见，因为服务端只有一个 candidate ，所以用 ice-lite 的方式实现更简单。
Nils 的截图中 answer sdp 中出现了两个一样的 candidate ， 是因为 zoom 人为调用了 addIceCandidate 添加的。
answer 中还出现了 a=setup:passive ，表示浏览器作为 DTLS 客户端的角色，这可能也是为了降低 server 的复杂性。


> NOTE candidate 候选项，服务端一般仅存在一个能与客户端通讯的公网地址，没必要启用其他复杂的 ICE 选项。


Quite noticeable is that we see the same server-side port 8801 both in Nils screenshot as well as the dump we gathered. This is no coincidence, Zoom?s native client runs on the same port. This means that all UDP packets have to be demultiplexed into sessions which is typically done by creating an association between the incoming UDP packets and the  _ufrag_ of the STUN requests from those packets. This is probably also the reason for munging the  _ufrag_. This is a bit silly ? demultiplexing works just as well with just the server side  _ufrag_.

值得注意的是， Nils 的截图和我收集的 dump 文件中，服务端端口都是 8801 。
这不是巧合，Zoom 的 native client 也使用相同的商品。
也就是说，所有的 UDP 包都会达到服务端的 8801 端口。
所以必须通过 STUN request 中的  _ufrag_ 分离收到的 UDP 包中的多路会话。
这就有点傻了吧？其实只用服务端 的 _ufrag_ 就足够分离多路会话了。


### Traffic inspection does not reveal anything new 流量分析并没有发现新情况

Inspecting the traffic is a bit more complicated. Since its JavaScript one can use the prototype override approach used by the  [WebRTC-Externals](https://webrtchacks.com/webrtc-externals/)  extension to snoop on any call to  RTCDataChannel.prototype.send. Unsurprisingly, the payload of the packets is the same as the one sent via WebSockets.

检查流量就比较复杂了。
因为它使用的 Javascript ，所以可利用 [WebRTC-Externals](https://webrtchacks.com/webrtc-externals/) 覆盖 prototype 方法来探测调用 RTCDataChannel.prototype.send 的调用情况。
不出所料，数据包的 payload 与 WebSocket 中完全一样。

This update to the Zoom web client will, as Nils pointed out, most likely increase the quality that was limited by the transfer via WebSockets over TCP quite a bit. While the client is now using WebRTC, it continues to avoid using the WebRTC media stack.

正如 Nils 所言， Zoom Web 客户端的此次更新明显改善了使用 WebSocket over TCP 方案造成的质量问题。
虽然它使用了 WebRTC ，但仍然没有使用 WebRTC media stack 。

> NOTE WebRTC media stack 指没有用 WebRTC 的编码解码和 RTP 打包等媒体数据处理相关的过程，仍然使用 wasm 编解码，打包 RTP 数据，然后仅仅通过 WebRTC 的 DataChannel 发送 RTP 数据。



author: [Philipp Hancke](https://webrtchacks.com/about/fippo/)



[^webrtcInZoomEN]: [How Zoom?s web client avoids using WebRTC (DataChannel Update)](https://webrtchacks.com/zoom-avoids-using-webrtc/)

[^webrtcInZoomCN]: [Zoom 的 Web 客户端与 WebRTC 有何不同？](https://toutiao.io/posts/c5vjvp/preview)

