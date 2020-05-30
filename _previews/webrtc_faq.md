---
layout: post
title:  "WebRTC FAQ"
date:   2020-01-01 12:00:00 +0800
tags:   tech
---

* category
{:toc}




### 12.RTP RTCP Multiplex SSRC 

参考以下内容

https://github.com/pion/webrtc/internal/mux/muxfunc.go

[RFC 7983](https://tools.ietf.org/html/rfc7983)

[RFC 5761](https://tools.ietf.org/html/rfc5761)

[RFC 8108](https://tools.ietf.org/html/rfc8108)


### 11.相关术语
https://blog.csdn.net/DittyChen/article/details/78065974

SR：发送者报告，描述作为活跃发送者万岁的发送和接收统计数据

RR：接收者报告，描述非活跃发送者成员的接收统计数据

SDES：源描述项，其中包括规范名CNAME

BYE：表明参与者将结束会话

APP：应用描述功能


关键帧请求

PLI 是Picture LossIndication，SLI 是Slice Loss Indication。发送方接收到接收方反馈的PLI或SLI需要重新让编码器生成关键帧并发送给接收端。

FIR：Full Intra Request，关键帧重传请求. Intra的含义是图像内编码，不需要其他图像信息即可解码；Inter指图像间编码，解码需要参考帧。故Intra Frame其实就是指I帧，Inter Frame指P帧或B帧。



重传请求

NACK：Negative Acknowledgemen 否定确认，NACK重传（丢包）

RPSI：Reference Picture Selection Indication


[码率控制](https://zhuanlan.zhihu.com/p/90094059)

TMMBR是Temporal Max MediaBitrate Request，表示临时最大码率请求。表明接收端当前带宽受限，告诉发送端控制码率。

REMB是ReceiverEstimatedMax Bitrate，接收端估计的最大码率。

TMMBN是Temporal Max MediaBitrate Notification

JITTER 抖动时间间隔, 数据包传输时长的变化就是拌动. 上一个包传输耗时 10ms ,下一个包传输耗时 20ms ,抖动就是 20-10=10ms. 第三个包传输耗时 25ms, 抖动就是 25-20=5ms



[关键帧](https://zhuanlan.zhihu.com/p/34162672)

Intra Picture 关键帧，作为随机访问的参考点，当成静态图像

Predictive Frame  参考帧，前向预测编码帧

Bi-directionalinterpolated prediction frame 前后参考帧，双向预测内插编码帧













### 10.SDP offer answer  中 ssrc 为什么不一样？以哪个为准 ？

### 9.jitter buffer

https://zhuanlan.zhihu.com/p/90094059

### 8.如何录制 vp8 opus 到一个文件中？
用 gstream 推流到 confserver ，然后 chrome 播放时，音频异响， sdp 协商的是 48k opus 2 。
如果能在 confserver 录制 chrome 推上来的视频到文件中，然后再用 confclient 推流到 confserver 。编码肯定就没问题了，


### 7.相关  RFC 文档都有哪些，什么作用？

RFC版本号 名称 备注

#### 1889

RTP: A Transport Protocol for Real-Time Applications

早期RTP协议，RTP v1

#### 1890

RTP Profile for Audio and Video Conferences with Minimal Control

RTP的负载类型定义，对应于RTP v1

#### 2198

RTP Payload for Redundant Audio Data

发送音频冗余数据的机制，FEC的雏形

#### 3550

RTP: A Transport Protocol for Real-Time Applications

现用RTP协议，RTP v2

#### 3551

RTP Profile for Audio and Video Conferences with Minimal Control

RTP的负载类型定义，对应于RTP v2

#### 3611

RTP Control Protocol Extended Reports (RTCP XR)

RTCP的拓展报文即XR报文定义

#### 3640

RTP Payload Format for Transport of MPEG-4 Elementary Streams

RTP负载为MPEG-4的格式定义

#### 3711

The Secure Real-time Transport Protocol (SRTP)

RTP媒体流采用AES-128对称加密

#### rfc6184 

Internet Engineering Task Force (IETF)  

- 3984

RTP Payload Format for H.264 Video

RTP负载为H264的格式定义，已被6184取代


#### 4103

RTP Payload for Text Conversation

RTP负载为文本或者T.140的格式定义

#### 4585

Extended RTP Profile for Real-time Transport Control Protocol (RTCP)-Based Feedback (RTP/AVPF)

NACK定义，通过实时的RTCP进行丢包重传

#### 4587

RTP Payload Format for H.261 Video Streams

H261的负载定义

#### 4588

RTP Retransmission Payload Format

RTP重传包的定义

#### 4961

Symmetric RTP / RTP Control Protocol (RTCP)

终端收发端口用同一个，叫做对称的RTP，便于DTLS加密

#### 5104

Codec Control Messages in the RTP Audio-Visual Profile with Feedback (AVPF)

基于4585实时RTCP消息，来控制音视频编码器的机制

#### 5109

RTP Payload Format for Generic Forward Error Correction

Fec的通用规范。

#### 5124

Extended Secure RTP Profile for Real-time Transport Control Protocol (RTCP)-Based Feedback (RTP/SAVPF)

SRTP的丢包重传

#### 5285

A General Mechanism for RTP Header Extensions

RTP 扩展头定义，可以扩展1或2个字节，比如CSRC，已被8285协议替代

#### 5450

Transmission Time Offsets in RTP Streams

计算RTP的时间差，可以配合抖动计算

#### 5484

Associating Time-Codes with RTP Streams

RTP和RTCP中时间格式的定义

#### 5506

Support for Reduced-Size Real-Time Transport Control Protocol (RTCP): Opportunities and Consequences

RTCP压缩

#### 5669

The SEED Cipher Algorithm and Its Use with the Secure Real-Time Transport Protocol (SRTP)

SRTP的对称加密算法的种子使用方法

#### 5691

RTP Payload Format for Elementary Streams with MPEG Surround Multi-Channel Audio

对于MPEG-4中有多路音频的RTP负载格式的定义

#### 5760

RTP Control Protocol (RTCP) Extensions for Single-Source Multicast Sessions with Unicast Feedback

RTCP对于单一源进行多播的反馈机制

#### 5761

Multiplexing RTP Data and Control Packets on a Single Port

RTP和RTCP在同一端口上传输

#### 6051

Rapid Synchronisation of RTP Flows

多RTP流的快速同步机制，适用于MCU的处理

#### 6128

RTP Control Protocol (RTCP) Port for Source-Specific Multicast (SSM) Sessions

RTCP对于多播中特定源的反馈机制

#### 6184

RTP Payload Format for H.264 Video

H264的负载定义

#### 6188

The Use of AES-192 and AES-256 in Secure RTP

SRTP拓展定义AES192和AES256

#### 6189

ZRTP: Media Path Key Agreement for Unicast Secure RTP

ZRTP的定义，非对称加密，用于密钥交换

#### 6190

RTP Payload Format for Scalable Video Coding

H264-SVC的负载定义

#### 6222

Guidelines for Choosing RTP Control Protocol (RTCP) Canonical Names (CNAMEs)

RTCP的CNAME的选定规则，可根据RFC 4122的UUID来选取

#### 6798 6843 6958 7002 7003 7097

RTP Control Protocol (RTCP) Extended Report (XR) Block for XXX

RTCP的XR报文，关于各个方面的定义

#### 6904

Encryption of Header Extensions in the Secure Real-time Transport Protocol

SRTP的RTP头信息加密

#### 7022

Guidelines for Choosing RTP Control Protocol (RTCP) Canonical Names (CNAMEs)

RTCP的CNAME的选定规则，修订6222

#### 7160

Support for Multiple Clock Rates in an RTP Session

RTP中的码流采样率变化的处理规则，音频较常见

#### 7164

RTP and Leap Seconds

RTP时间戳的校准机制

#### 7201

Options for Securing RTP Sessions

RTP的安全机制的建议，什么时候用DTLS，SRTP，ZRTP或者RTP over TLS等

#### 7202

Securing the RTP Framework: Why RTP Does Not Mandate a Single Media Security Solution

RTP的安全机制的补充说明

#### 7656

A Taxonomy of Semantics and Mechanisms for Real-Time Transport Protocol (RTP) Sources

RTP在webrtc中的应用场景

#### 7667

RTP Topologies

在MCU等复杂系统中，RTP流的设计规范

#### 7741

RTP Payload Format for VP8 Video

负载为vp8的定义

#### 7798

RTP Payload Format for High Efficiency Video Coding (HEVC)

负载为HEVC的定义

#### 8082

Using Codec Control Messages in the RTP Audio-Visual Profile with Feedback with Layered Codecs

基于4585实时RTCP消息，来控制分层的音视频编码器的机制，对于5104协议的补充

#### 8083

Multimedia Congestion Control: Circuit Breakers for Unicast RTP Sessions

RTP的拥塞处理之码流环回的处理

#### 8108

Sending Multiple RTP Streams in a Single RTP Session

单一会话，单一端口传输所有的RTP/RTCP码流，对现有RTP/RTCP机制的总结

#### 8285

A General Mechanism for RTP Header Extensions

RTP 扩展头定义，可以同时扩展为1或2个字节



### 6.为什么在 firefox 中仅拉取媒体流时，即 sdp recvonly 时会失败？

### 5.如何在一个 RTCPeerConnection 中拉取多个成员的 video/audio
TODO 这个问题优先级最低
先考虑用简单的方式实现

http://yunxin.163.com/blog/webrtc-2/

- 1）单个PeerConnection里有多个media stream（bundle）；

  媒体处理代码简单；只有一次 DTLS 握手，服务端性能好；

  需要兼容不同浏览器。
  下午媒体流过多时， SDP 长度过长，如果频繁进出， SDP 耗费较多带宽。

- 2）为不同的media stream创建不同的PeerConnection。

  不同浏览器兼容性好；

  服务端维护每个用户上下行对应关系，代码逻辑复杂。
  
#### 5.1.SDP 协商过程  Plan B 与 JSEP(Unified Plan) 区别？

https://www.douban.com/note/658626071/

https://ouchunrun.github.io/2018/10/23/%E2%80%9CUnified%20Plan%E2%80%9D%20%E8%BF%87%E6%B8%A1%E6%8C%87%E5%8D%97/

webrtc为了一个peerconnection上支持多个media stream SDP有两个方法来表征.

- Plan B  https://tools.ietf.org/html/draft-uberti-rtcweb-plan-00

  Plan B是一个m line里多路media stream(通过msid区分).

- JSEP(Unified Plan)  https://tools.ietf.org/html/draft-ietf-rtcweb-jsep-24

  JSEP是一个m line对应一个media stream.
  JSEP里边多个m line的使用有点tricky需要循环利用.

目前Chrome还没支持JSEP,正在支持中. Firefox已经支持JSEP.

JSEP demo

offer
```sdp
v=0
o=- 4380437079560121890 2 IN IP4 127.0.0.1
s=-
t=0 0

// NOTE BUNDLE
a=group:BUNDLE 0 1

a=msid-semantic: WMS gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S
m=audio 62066 UDP/TLS/RTP/SAVPF 111 103 104 9 0 8 106 105 13 110 112 113 126
c=IN IP4 192.168.0.105
a=rtcp:9 IN IP4 0.0.0.0
a=candidate:1221703924 1 udp 2122260223 192.168.0.105 62066 typ host generation 0 network-id 1 network-cost 10
a=candidate:3419488047 1 udp 2122194687 10.10.5.131 62686 typ host generation 0 network-id 2 network-cost 50
a=candidate:106054660 1 tcp 1518280447 192.168.0.105 9 typ host tcptype active generation 0 network-id 1 network-cost 10
a=candidate:2236793823 1 tcp 1518214911 10.10.5.131 9 typ host tcptype active generation 0 network-id 2 network-cost 50
a=ice-ufrag:Z82v
a=ice-pwd:TGo0VBBobIDFROjA68tzc6SZ
a=ice-options:trickle
a=fingerprint:sha-256 B4:97:1B:41:55:29:A4:39:8D:A7:C3:57:F6:AE:FC:B0:34:BC:AB:7C:7A:48:CD:B9:3C:2D:96:AC:26:15:33:C2
a=setup:actpass
a=mid:0
a=extmap:1 urn:ietf:params:rtp-hdrext:ssrc-audio-level
a=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time
a=extmap:3 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01
a=extmap:4 urn:ietf:params:rtp-hdrext:sdes:mid
a=extmap:5 urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id
a=extmap:6 urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id
a=sendrecv
a=msid:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S 734751a3-05c2-4399-a57a-092888c42f6a
a=rtcp-mux

// NOTE useinbandfec
a=rtpmap:111 opus/48000/2
a=rtcp-fb:111 transport-cc
a=fmtp:111 minptime=10;useinbandfec=1

a=rtpmap:103 ISAC/16000
a=rtpmap:104 ISAC/32000
a=rtpmap:9 G722/8000
a=rtpmap:0 PCMU/8000
a=rtpmap:8 PCMA/8000
a=rtpmap:106 CN/32000
a=rtpmap:105 CN/16000
a=rtpmap:13 CN/8000
a=rtpmap:110 telephone-event/48000
a=rtpmap:112 telephone-event/32000
a=rtpmap:113 telephone-event/16000
a=rtpmap:126 telephone-event/8000
a=ssrc:2060113113 cname:m0iJpygIF0G096YE
a=ssrc:2060113113 msid:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S 734751a3-05c2-4399-a57a-092888c42f6a
a=ssrc:2060113113 mslabel:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S
a=ssrc:2060113113 label:734751a3-05c2-4399-a57a-092888c42f6a
m=video 61258 UDP/TLS/RTP/SAVPF 96 97 98 99 100 101 102 122 127 121 125 107 108 109 124 120 123 119 114 115 116
c=IN IP4 192.168.0.105
a=rtcp:9 IN IP4 0.0.0.0
a=candidate:1221703924 1 udp 2122260223 192.168.0.105 61258 typ host generation 0 network-id 1 network-cost 10
a=candidate:3419488047 1 udp 2122194687 10.10.5.131 56542 typ host generation 0 network-id 2 network-cost 50
a=candidate:106054660 1 tcp 1518280447 192.168.0.105 9 typ host tcptype active generation 0 network-id 1 network-cost 10
a=candidate:2236793823 1 tcp 1518214911 10.10.5.131 9 typ host tcptype active generation 0 network-id 2 network-cost 50
a=ice-ufrag:Z82v
a=ice-pwd:TGo0VBBobIDFROjA68tzc6SZ
a=ice-options:trickle
a=fingerprint:sha-256 B4:97:1B:41:55:29:A4:39:8D:A7:C3:57:F6:AE:FC:B0:34:BC:AB:7C:7A:48:CD:B9:3C:2D:96:AC:26:15:33:C2
a=setup:actpass
a=mid:1
a=extmap:14 urn:ietf:params:rtp-hdrext:toffset
a=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time
a=extmap:13 urn:3gpp:video-orientation
a=extmap:3 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01
a=extmap:12 http://www.webrtc.org/experiments/rtp-hdrext/playout-delay
a=extmap:11 http://www.webrtc.org/experiments/rtp-hdrext/video-content-type
a=extmap:7 http://www.webrtc.org/experiments/rtp-hdrext/video-timing
a=extmap:8 http://tools.ietf.org/html/draft-ietf-avtext-framemarking-07
a=extmap:9 http://www.webrtc.org/experiments/rtp-hdrext/color-space
a=extmap:4 urn:ietf:params:rtp-hdrext:sdes:mid
a=extmap:5 urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id
a=extmap:6 urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id
a=sendrecv
a=msid:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S 902d625b-b968-4102-a87d-ace4a1ed5931
a=rtcp-mux
a=rtcp-rsize
a=rtpmap:96 VP8/90000
a=rtcp-fb:96 goog-remb
a=rtcp-fb:96 transport-cc
a=rtcp-fb:96 ccm fir
a=rtcp-fb:96 nack
a=rtcp-fb:96 nack pli
a=rtpmap:97 rtx/90000
a=fmtp:97 apt=96
a=rtpmap:98 VP9/90000
a=rtcp-fb:98 goog-remb
a=rtcp-fb:98 transport-cc
a=rtcp-fb:98 ccm fir
a=rtcp-fb:98 nack
a=rtcp-fb:98 nack pli
a=fmtp:98 profile-id=0
a=rtpmap:99 rtx/90000
a=fmtp:99 apt=98
a=rtpmap:100 VP9/90000
a=rtcp-fb:100 goog-remb
a=rtcp-fb:100 transport-cc
a=rtcp-fb:100 ccm fir
a=rtcp-fb:100 nack
a=rtcp-fb:100 nack pli
a=fmtp:100 profile-id=2
a=rtpmap:101 rtx/90000
a=fmtp:101 apt=100
a=rtpmap:102 H264/90000
a=rtcp-fb:102 goog-remb
a=rtcp-fb:102 transport-cc
a=rtcp-fb:102 ccm fir
a=rtcp-fb:102 nack
a=rtcp-fb:102 nack pli
a=fmtp:102 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f
a=rtpmap:122 rtx/90000
a=fmtp:122 apt=102
a=rtpmap:127 H264/90000
a=rtcp-fb:127 goog-remb
a=rtcp-fb:127 transport-cc
a=rtcp-fb:127 ccm fir
a=rtcp-fb:127 nack
a=rtcp-fb:127 nack pli
a=fmtp:127 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f
a=rtpmap:121 rtx/90000
a=fmtp:121 apt=127
a=rtpmap:125 H264/90000
a=rtcp-fb:125 goog-remb
a=rtcp-fb:125 transport-cc
a=rtcp-fb:125 ccm fir
a=rtcp-fb:125 nack
a=rtcp-fb:125 nack pli
a=fmtp:125 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f
a=rtpmap:107 rtx/90000
a=fmtp:107 apt=125
a=rtpmap:108 H264/90000
a=rtcp-fb:108 goog-remb
a=rtcp-fb:108 transport-cc
a=rtcp-fb:108 ccm fir
a=rtcp-fb:108 nack
a=rtcp-fb:108 nack pli
a=fmtp:108 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f
a=rtpmap:109 rtx/90000
a=fmtp:109 apt=108
a=rtpmap:124 H264/90000
a=rtcp-fb:124 goog-remb
a=rtcp-fb:124 transport-cc
a=rtcp-fb:124 ccm fir
a=rtcp-fb:124 nack
a=rtcp-fb:124 nack pli
a=fmtp:124 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=4d0032
a=rtpmap:120 rtx/90000
a=fmtp:120 apt=124
a=rtpmap:123 H264/90000
a=rtcp-fb:123 goog-remb
a=rtcp-fb:123 transport-cc
a=rtcp-fb:123 ccm fir
a=rtcp-fb:123 nack
a=rtcp-fb:123 nack pli
a=fmtp:123 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=640032
a=rtpmap:119 rtx/90000
a=fmtp:119 apt=123
a=rtpmap:114 red/90000
a=rtpmap:115 rtx/90000
a=fmtp:115 apt=114
a=rtpmap:116 ulpfec/90000
a=ssrc-group:FID 1169412873 1092092900
a=ssrc:1169412873 cname:m0iJpygIF0G096YE
a=ssrc:1169412873 msid:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S 902d625b-b968-4102-a87d-ace4a1ed5931
a=ssrc:1169412873 mslabel:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S
a=ssrc:1169412873 label:902d625b-b968-4102-a87d-ace4a1ed5931
a=ssrc:1092092900 cname:m0iJpygIF0G096YE
a=ssrc:1092092900 msid:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S 902d625b-b968-4102-a87d-ace4a1ed5931
a=ssrc:1092092900 mslabel:gBybe8DpDXnbRHnN4dAhyMiHrcmIZizyQn7S
a=ssrc:1092092900 label:902d625b-b968-4102-a87d-ace4a1ed5931

```

answer
```sdp
v=0
o=- 613547399 1587655707 IN IP4 0.0.0.0
s=-
t=0 0
a=fingerprint:sha-256 03:4A:0A:68:E9:46:55:6D:1F:82:08:AB:26:66:3D:26:46:22:B6:36:4E:D9:1F:75:37:B5:BF:75:6D:21:D2:54
a=group:BUNDLE 0 1
m=audio 9 UDP/TLS/RTP/SAVPF 111
c=IN IP4 0.0.0.0
a=setup:passive
a=mid:0
a=ice-ufrag:wrzvSGQqjyNrwnwZ
a=ice-pwd:NfJJvutdGJzlIsmxOaDDGdTXEuZUTqtn
a=rtcp-mux
a=rtcp-rsize
a=rtpmap:111 opus/48000/2
a=fmtp:111 minptime=10;useinbandfec=1
a=ssrc:2854263694 cname:HZYBLVMKViCqAvla
a=ssrc:2854263694 msid:HZYBLVMKViCqAvla GoTnKALOcyHnKDWB
a=ssrc:2854263694 mslabel:HZYBLVMKViCqAvla
a=ssrc:2854263694 label:GoTnKALOcyHnKDWB
a=sendrecv
a=candidate:foundation 1 udp 2130706431 10.10.5.131 59925 typ host generation 0
a=candidate:foundation 2 udp 2130706431 10.10.5.131 59925 typ host generation 0
a=candidate:foundation 1 udp 2130706431 192.168.0.105 65272 typ host generation 0
a=candidate:foundation 2 udp 2130706431 192.168.0.105 65272 typ host generation 0
a=end-of-candidates
m=video 9 UDP/TLS/RTP/SAVPF 96
c=IN IP4 0.0.0.0
a=setup:passive
a=mid:1
a=ice-ufrag:wrzvSGQqjyNrwnwZ
a=ice-pwd:NfJJvutdGJzlIsmxOaDDGdTXEuZUTqtn
a=rtcp-mux
a=rtcp-rsize
a=rtpmap:96 VP8/90000
a=ssrc:1879968118 cname:uQKEWpRyxyPmGkMz
a=ssrc:1879968118 msid:uQKEWpRyxyPmGkMz SmTRzWlKzZADUoZj
a=ssrc:1879968118 mslabel:uQKEWpRyxyPmGkMz
a=ssrc:1879968118 label:SmTRzWlKzZADUoZj
a=sendrecv
a=candidate:foundation 1 udp 2130706431 10.10.5.131 59925 typ host generation 0
a=candidate:foundation 2 udp 2130706431 10.10.5.131 59925 typ host generation 0
a=candidate:foundation 1 udp 2130706431 192.168.0.105 65272 typ host generation 0
a=candidate:foundation 2 udp 2130706431 192.168.0.105 65272 typ host generation 0
a=end-of-candidates

```


  

### 4.为什么 pion 的 ICE 协调过程那么慢？
有时要几十秒。
跟 stun:stun.l.google.com:19302 有关系吗？

测试时，各端都是同一局域网，可以不用 stun ，这样就很快。
```js
    let config = {"iceServers":[],"iceTransportPolicy":"all","iceCandidatePoolSize":"0"}
    let pc = new RTCPeerConnection();
```

### 3.pionConfDemo 要实现的功能
- ConfServer
  offer/answer 协商
  decode SRTP to RTP
  encode RTP to SRTP
- ConfClientH5
  upload  media  over DTLS
  pull media  over DTLS
  show video
- ConfClientGo
  upload mkv file
  pull media
  show video

### 2.编码的区别
pcmu/pcma/g722/opus/vp8/vp9/h264/aac



- AAC(Advanced Audio Coding)

- PCM(Pulse code modulation) 脉冲编码调制
  在计算机应用中，能够达到最高保真水平的就是PCM编码，被广泛用于素材保存及音乐欣赏，CD、DVD以及我们常见的WAV文件中均有应用。因此，PCM约定俗成了无损编码，因为PCM代表了数字音频中最佳的保真水准，并不意味着PCM就能够确保信号绝对保真，PCM也只能做到最大程度的无限接近。要算一个PCM音频流的码率是一件很轻松的事情，采样率值×采样大小值×声道数bps。一个采样率为44.1KHz，采样大小为16bit，双声道的PCM编码的WAV文件，它的数据速率则为 44.1K×16×2=1411.2 Kbps。我们常见的Audio CD就采用了PCM编码，一张光盘的容量只能容纳72分钟的音乐信息。

PCM就是一种模拟信号数字化的过程，将模拟信号抽样， 量化， 编码 然后调制发射。
其中的编码部分最简单的就是G711编码格式了。G711又分为了A律和U律，美洲国家在使用U律，中国使用的是A律
PCM 就是声音原始格式未经过任何压缩，类似位图的BMP格式
https://bbs.csdn.net/topics/392317609

- PCMU PCMA
PCMUandPCMA都能够达到CD音质，但是它们消耗的带宽也最多(64kbps)。如果网络带宽比较低，可以选用低比特速率的编码方法，如G.723或G.729，这两种编码的方法也能达到传统长途电话的音质，但是需要很少的带宽（G723需要5.3/6.3kbps，G729需要8kbps）。如果带宽足够并且需要更好的语音质量，就使用PCMU和 PCMA，甚至可以使用宽带的编码方法G722(64kbps)，这可以提供有高保真度的音质。

> https://baike.baidu.com/item/%E8%AF%AD%E9%9F%B3%E7%BC%96%E8%A7%A3%E7%A0%81#19

[rfc4856 Media Type Registration of Payload Formats in the RTP Profile for Audio and Video Conferences](https://tools.ietf.org/html/rfc4856#page-21)


### 1.why pion not support aac

封装 RTP 协议时，应该不需要过多了解 RTP palyloadType 。
所以，理论上 pion 应该支持任意类型的 playloadType 。
但为什么代码中还是对不同协议做了处理，具体不同 playloadType 在 pion 中有哪些不同处理呢？


```go
132 // Names for the default codecs supported by Pion WebRTC
133 const (         
134         PCMU = "PCMU"
135         PCMA = "PCMA"
136         G722 = "G722"
137         Opus = "opus"
138         VP8  = "VP8"
139         VP9  = "VP9"
140         H264 = "H264"
141 )
142  
"~/go/pkg/mod/github.com/pion/webrtc/v2@v2.1.18/mediaengine.go" [只读] 313 行 --35%--  
```

```go
109 func GetAudioCodec() []JsonCrgwCustomDataCodec {
110         rets := []JsonCrgwCustomDataCodec{}
111         rets = append(rets, ParseCodecItem("opus@121"))
112         rets = append(rets, ParseCodecItem("opus"))
113         rets = append(rets, ParseCodecItem("pcmu"))
114 
115         return rets
116 }
117 
118 func GetVideoCodec() []JsonCrgwCustomDataCodec {
119         rets := []JsonCrgwCustomDataCodec{}
120         rets = append(rets, ParseCodecItem("h264"))
121         rets = append(rets, ParseCodecItem("vp8"))
122 
123         return rets
124 }
125 
126 func GetSsCodec() []JsonCrgwCustomDataCodec {
127         rets := []JsonCrgwCustomDataCodec{}
128         rets = append(rets, ParseCodecItem("h264"))
129         rets = append(rets, ParseCodecItem("vp8"))
130 
131         return rets
132 
```

> test mkv file
https://www.freedesktop.org/software/gstreamer-sdk/data/media/


## ref

https://github.com/liwf616/awesome-live-stream

