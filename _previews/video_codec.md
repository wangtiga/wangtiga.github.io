---
layout: post
title:  "[译]Video Conversion done right: Codecs and Software"
date:   2020-01-01 12:00:00 +0800
tags:   tech
---


# Video Conversion done right: Codecs and Software [^VideoCodec] [^devfest2019]

November 7, 2011  by  [slhck](http://blog.superuser.com/author/slhck/ "Posts by slhck").  [6 comments](http://blog.superuser.com/2011/11/07/video-conversion-done-right-codecs-and-software/#comments)

Videos are everywhere. They come in quite a few different formats – all with their own advantages and disadvantages. Converting videos from one format to another is a very simple task, given the right tools. In this post, we will go through the most popular video codecs and the software you need to get the best results.

## Current Video Codecs and Containers

Before we begin to encode, we have to ask: What is a video codec? “Codec” stands for encoder/decoder. A video codec specifies the software (or the algorithm) to convert video files into a bitstream and then decode that bitstream into a video again. But why is there even a need for codecs? Every video is a sequence of frames, mostly more than 25 frames per second. This is a lot of data. Codecs, when used to compress videos, try to reduce this data using some visual tricks.

什么是 video codec ? ""Codec" 就是 编码器/解码器。
视频编码是将 视频文件转换成 bitstream 。
视频解码是将 bitstream 转换成视频文件。
但为什么需要 codec 呢？
视频是由一组连续的 frame (帧) 组成，一般每秒超过 25 frame 。这些数据量是非常大的。
codec 利用 visual trick (视觉特点)压缩视频大小。


-   **Codecs**  take a video and convert it into a bitstream, mostly with the goal to minimize its size.
-   **Containers**  on the other hand are file formats that are supposed to wrap video and audio into a single file. Video codecs don’t know anything about audio encoding, so this is why there is a need to synchronize them for playback.

-   **Codecs** 将视频转换成 bitstream ，主要用于压缩视频大小
-   **Containers** 是一种文件格式，用于将 video 和 audio 包装成一个文件。 video codec 和 audio encoding 互不关心，因为需要一种同步机制，在回放视频文件时，保持音视频同步。


Codecs mostly encode video with regard to a specification. Almost every specification for video has been written by the Motion Picture Experts Group. This is why most video formats are known under a MPEG acronym. There is always one specification, but there can be an arbitrary number of codecs for this specification.

codec 在 encode video 时都会遵守某个标准。
有关 video 的标准都是由 Motion Picture Experts Group (动态图像专家组)。
所以很多视频格式后缀名都使用 MPEG 缩写。
specification 只有一个，但会有很多 codec 。



These days, there are  [a lot of codecs](http://en.wikipedia.org/wiki/Video_codec#Commonly_used_video_codecs)  around. We will take a look at the best and most widely used ones. The following is partly taken from my answer to the question  [What is a Codec (e.g. DivX?), and how does it differ from a File Format (e.g. MPG)?](http://superuser.com/questions/300897/what-is-a-codec-e-g-divx-and-how-does-it-differ-from-a-file-format-e-g-mpg). Also see this answer for an explanation of current containers.

### MPEG-2

This is mostly used on DVDs and TV broadcasting. It offers high quality at the expense of high file sizes. It’s been the main choice for video compression for many many years now. Encoders are often found embedded into hardware. The encoding scheme is very basic and therefore allows for fast encoding and fast playback.

At some point though, its quality is not good enough for low bitrates. While a satellite or cable transmission offered enough bitrate for a high quality MPEG-2 video, it just wasn’t good enough for the internet and multimedia age.

-   Containers for MPEG-2: Mostly found within an  `.mpg`  container
-   Encoders for MPEG-2: CinemaCraft, TMPGEnc, MainConcept and ffmpeg

Basically, you only want to create MPEG-2 videos if you really don’t care about storage size, or you want to put videos on a Video DVD. Note that many applications for authoring Video DVDs will already convert to MPEG-2 for you.

### MPEG-4 Part II

When it became clear that MPEG-2 was not enough, MPEG-4 was released. Actually, MPEG-4 is a huge standard consisting of many parts, only one of which is about basic video coding: MPEG-4 Part II.

This is probably the one used mostly to encode videos for the web. It offers good quality at practical file sizes, which means one can burn a whole movie onto a CD (whereas with MPEG-2 you would have needed a DVD).

Its drawback is that the quality itself might not be good enough for some viewers.

-   Containers for MPEG-4 Part II: Mostly in  `.avi`
-   Encoders for MPEG-4 Part II:  [DivX](http://www.divx.com/),  [XviD](http://www.xvid.org/), Nero Digital
-   Audio used in MPEG-4 Part II: Mostly MP3.

So, if you want to rip a DVD without taking too much time, and at the same time ensuring best compatibility with current DVD players, go for this.

### MPEG-4 Advanced Video Coding (h.264)

This is the big boss, the number one codec today. It offers superb quality at small file sizes and therefore is perfectly suited for all kinds of video for the internet.

mp4(h.264) 是当今最流行的编码方案，它生成的文件更小，质量更高，因此非常适合在互联网中使用。

Converting video for small bandwidths is its main purpose. Originally, it was called h.264 (and one can see that this has nothing to do with the MPEG). In fact, the  [ITU](http://en.wikipedia.org/wiki/International_Telecommunication_Union)  created it based on h.263 – both were meant for videoconferencing. They however joined efforts and created this new standard.

设计 mp4 编码的主要目的是减少视频的大小，占用更小的 带宽 。
最初，它被称为 h.264 （跟 MPEG 没什么关系）。
实际上 ITU 最初是基于 h.263 设计 －－ 这两个编码方案都是用于视频会议的。

h.264 is found in almost every modern application or device, from phones to camcorders, even on Blu Ray disks: Video is now encoded in h.264. Its quality vs. file size ratio is just so much better than what you could achieve with MPEG-2 or even MPEG-4 Part 2. The main disadvantage is that it is very slow to encode as it has some vast algorithmic improvements over them (which take a lot of time to compute).

-   Containers for h.264:  `.mp4`,  `.mkv`,  `.mov`.
-   Encoders for h.264:  [x264](http://www.videolan.org/developers/x264.html),  [Mainconcept](http://www.mainconcept.com/), QuickTime
-   Audio used in h.264: Mostly  [AAC](http://en.wikipedia.org/wiki/Aac)  (MPEG-4 Audio).

If you upload videos to the web, this is your choice. Also, if you want the best quality at the lowest file size, you should always use h.264. In my experience,  [x264](http://www.videolan.org/developers/x264.html)  is the best free encoder for h.264 out there.

----------

## Set up your Codecs

Most current operating systems don’t come with all codecs (encoders and decoders) installed. This is mainly due to legal and licensing issues. However there are codec packs that you can easily install later.

-   Windows: The  [K-Lite Codec Pack](http://www.codecguide.com/download_kl.htm)  offers almost everything there is. The Standard package should be enough for most users.
-   OS X:  [Perian](http://perian.org/)  is a QuickTime component for OS X that allows native playback for all popular codecs within the QuickTime framework.
-   Linux: Depending on your distribution, there are probably some codec packages in the repositories. For example, in Ubuntu there are  [RestrictedFormats](https://help.ubuntu.com/community/RestrictedFormats)  which can be easily installed later on.

Now that you have your codecs installed, you will be able to play almost every video (and audio) file there is.

----------

## Software for Converting Videos

Of course, there are lots of applications that allow you to convert video from one format to another. Some of them are commercial, but a huge part is free.

Why is that? It turns out that there is a free and open source video toolkit called  [ffmpeg](http://www.ffmpeg.org/). It is based on the  `libavcodec`  and  `libavformat`  libraries, which are a libraries of encoders, decoders, multiplexers and demultiplexers for almost every video codec and format there is today. The fact that it is available for almost every platform, from a Windows PC to an Android phone, makes it extremely versatile. It also offers very good quality when compared to the professional tools.

Because it’s free and open source, it is the basis for almost every free video conversion software. One of the most popular ones is  [Handbrake](http://handbrake.fr/). Let’s take a quick look at it.

### Use Handbrake!

Handbrake is free and available for Windows, Mac OS X and Linux (including  `deb`  and  `rpm`  packages). The installers can be downloaded from the  [Handbrake website](http://handbrake.fr/downloads.php). It is extremely useful for extracting videos from a DVD or converting videos for devices such as iPods or video streaming servers. Note that one downside of Handbrake is that it won’t output  `.avi`  files, which are sometimes needed for playback on legacy devices.

Let’s install it and open the application. It looks a bit different on Windows and Linux, but the functionality is of course the same.

![](http://i.imgur.com/O2E9A.png)

Once you open it, it will ask you for a source video. There are quite a few settings:

-   **Video Codec**: Here you can choose between h.264 and MPEG-4 Part II. Choose whatever fits your needs. If you want a video targeted for modern mobile devices or the web, go for h.264. This should be fine for almost every situation. If you want to create a movie CD for a DVD player capable of playing back XviD/DivX, you should choose MPEG-4. Also, MPEG-4 is still necessary for some older mobile devices. Note that encoding to h.264 takes longer than for MPEG-4.
-   **Framerate**: You should choose to keep the source frame rate unless you have other specific requirements. This makes sure you won’t get any stuttering picture.
-   **Quality**: Now you should think about how whether you have  _file size_  constraints (such as when encoding for a DVD) or  _quality_  constraints. Setting an average bitrate can help when you know that you want to stream the video and you can only stream up to 1500 kilobit per second. Fiddling around with “Constant quality” can sometimes be useful, but it can result in either too large files or too bad quality.

This is probably everything you need. Note that encoding a video that is already encoded might decrease its quality, so if you can use a higher bit rate or quality setting, do it!

Sometimes you might want to resize a video. Use the “Picture Settings” tool located in the toolbar.

![](http://i.imgur.com/KtYye.png)

In order to be able to change the picture size, set “Anamorphic” to “loose”. This will allow you to scale the video keeping the aspect ratio.

### Windows-only alternatives

#### SUPER©

[SUPER](http://www.erightsoft.com/SUPER.html)  is an ffmpeg-based encoder written for Windows. Once you find the download link on the website, you’re good to go. Hint: It’s at the bottom  [of this page](http://www.erightsoft.com/S6Kg1.html). The interface might seem a bit cheap, but SUPER can convert almost everything. It is based on ffmpeg as well, so you should see the same results from either using Handbrake or SUPER.

SUPER’s advantage over Handbrake is that it can output  `.avi`  files, as already stated above. It is also scriptable in case you need to convert a larger number of videos at once.

![](http://i.stack.imgur.com/V5Npc.png)

Once started, SUPER gives you a list of options. You should choose your container and video codec first, depending on your needs. The rest is quite self-explanatory (although the choice of it might seem overwhelming at first). Just make sure you don’t change the video resolution unless absolutely necessary and keep the frame rate the same.

#### FormatFactory

[FormatFactory](http://pcfreetime.com/)  is ad-supported, but freeware. It might also not look that nice, but it definitely gets the job done. Have a look at the options which now should look familiar to you. It also allows ripping DVDs to a more suitable format.

![](http://i.imgur.com/Rbwnj.jpg)

### OS X only alternatives

#### FFmpegX

[ffmpegX](http://www.ffmpegx.com/)  bundles the ffmpeg library into a GUI application. Although the contained version of ffmpeg is not the latest (it’s from 2008), this tool will help you convert videos between different formats.

![](http://i.imgur.com/zNysg.png)

Its video settings allow you to calculate the necessary bitrate for a given file size, or vice-versa. You can also tell it to convert the video so it exactly fits on one CD. What’s very nice about ffmpegX is its collection of tools. With those, you can for example multiplex a video and audio stream, or you can join two separate video files. It also allows you to repair a broken AVI file.

## Wrapping it up

In this post I gave you a quick overview about current codecs and how to choose them for your needs. We’ve seen a few tools, but of course, there are a lot more! Which ones do you use on a regular basis? What is there for Linux?

In my next post, I’d like to show you how to use  `ffmpeg`  as a command line tool, because it offers all the functionality you get from a GUI tool, but has the benefit of being easily scriptable. I’ll try to give you some examples you can bookmark to use whenever you need to convert videos.




[^VideoCodec]: [Video Conversion done right: Codecs and Software](http://blog.superuser.com/2011/11/07/video-conversion-done-right-codecs-and-software/)

[^devfest2019]: [高嵩: Web 音视频技术入坑指南](https://mp.weixin.qq.com/s/MK21kE9Hedgq7qmWMSHEzA)


