---
layout: post
title:  "WebRTC FAQ"
date:   2020-01-01 12:00:00 +0800
tags:   tech
---


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

#### 3984

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

### 4.为什么 pion 的 ICE 协调过程那么慢？
有时要几十秒。
跟 stun:stun.l.google.com:19302 有关系吗？

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


