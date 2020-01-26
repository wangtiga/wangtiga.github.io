---
layout: post
title:  "[译]GStream Basic tutorial 10: GStreamer tools"
date:   2019-11-03 12:00:00 +0800
tags:   tech
---

* category
{:toc}



# Basic tutorial 10: GStreamer tools [^GStreamDoc] [^GStreamIBM] [^NoteditGstrem]

GStreamer 作为 GNOME 桌面环境推荐的流媒体应用框架，采用了基于插件（plugin）和管道（pipeline）的体系结构，框架中的所有的功能模块都被实现成可以插拔的组件（component）， 并且在需要的时候能够很方便地安装到任意一个管道上，由于所有插件都通过管道机制进行统一的数据交换，因此很容易利用已有的各种插件“组装”出一个功能完善的多媒体应用程序。

## gst-launch-1.0 


This tool accepts a textual description of a pipeline, instantiates it, and sets it to the PLAYING state. It allows you to quickly check if a given pipeline works, before going through the actual implementation using GStreamer API calls.

`gst-launch-1.0` 是测试文本 pipeline 的命令行工具

Bear in mind that it can only create simple pipelines. In particular, it can only simulate the interaction of the pipeline with the application up to a certain level. In any case, it is extremely handy to test pipelines quickly, and is used by GStreamer developers around the world on a daily basis.

它只能模拟一些简单的 pipeline 命令。

Please note that gst-launch-1.0 is primarily a debugging tool for developers. You should not build applications on top of it. Instead, use the `gst_parse_launch()` function of the GStreamer API as an easy way to construct pipelines from pipeline descriptions.

gst-launch-1.0 主要是给开发者使用的工具，不要使用它开发应用。
开发简单的功能可以用 `gst_parse_launch()` 函数运行 pipleine 。

Although the rules to construct pipeline descriptions are very simple, the concatenation of multiple elements can quickly make such descriptions resemble black magic. Fear not, for everyone learns the gst-launch-1.0 syntax, eventually.

多个简单 pipeline 规则可以组合成十分强大的功能。

The command line for gst-launch-1.0 consists of a list of options followed by a PIPELINE-DESCRIPTION. Some simplified instructions are given next, se the complete documentation at the reference page for gst-launch-1.0.

gst-launch-1.0 的命令行参数主要是 不同选项组成的 PIPELINE-DESCRIPTION 。
下面简单介绍几个，完整文档参考 [gst-launch-1.0 reference page](https://gstreamer.freedesktop.org/documentation/tools/gst-launch.html) 。



### Elements [元件](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html#elements)

In simple form, a PIPELINE-DESCRIPTION is a list of element types separated by exclamation marks (!). Go ahead and type in the following command:

简单的 PIPELINE-DESCRIPTION 是用感叹号 `!` 分隔的多个 element 组成。
如下所示：

```shell
gst-launch-1.0 videotestsrc ! videoconvert ! autovideosink
```

执行完毕后，会看到一个视频窗口。
在终端中按 CTRL+C 能结束这个程序。

You should see a windows with an animated video pattern. Use CTRL+C on the terminal to stop the program.

This instantiates a new element of type  `videotestsrc`  (an element which generates a sample video pattern), an  `videoconvert`  (an element which does raw video format conversion, making sure other elements can understand each other), and an  `autovideosink`  (a window to which video is rendered). Then, GStreamer tries to link the output of each element to the input of the element appearing on its right in the description. If more than one input or output Pad is available, the Pad Caps are used to find two compatible Pads.

这个命令实例化一个 `videotestsrc` 类型的 element ，它专门用于生成一个示例图像。
随后，经过 `videoconvert` element 的处理，视频被转换成 raw （原始） 视频格式，以便其他 element 能处理这个视频数据，
接着，`autovideosink` element 将视频图像显示在 windows 窗口中。
 GStreamer 会把每个 element 的 output input 链接到一起。
如果有超过一个 input/output Pad 可用，会根据 Pad 的能力选择。



### Properties [属性](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html#properties)

Properties may be appended to elements, in the form `property=value` (multiple properties can be specified, separated by spaces). Use the  `gst-inspect-1.0`  tool (explained next) to find out the available properties for an element.

可以设置 element 的 property ，格式为 `property2=value property2=value` （使用空格分隔多个属性）。
使用 `gst-inspect-1.0` 工具查找 element 的可用属性（稍后会介绍）。

```shell
gst-launch-1.0 videotestsrc pattern=11 ! videoconvert ! autovideosink
```

You should see a static video pattern, made of circles.

执行上述命令后，会看到像年轮一样的圆圈组成的静态视频。



### Named elements [命名元件](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html#named-elements)

Elements can be named using the  `name`  property, in this way complex pipelines involving branches can be created. Names allow linking to elements created previously in the description, and are indispensable to use elements with multiple output pads, like demuxers or tees, for example.

使用 `name` 属性能给 element 起名称，善用这种方法，能用 pipeline 构造出强大的功能。

Named elements are referred to using their name followed by a dot.

引用命名属性时，要在名称后面加一个点`.`。

```shell
gst-launch-1.0 videotestsrc ! videoconvert ! tee name=t ! queue ! autovideosink t. ! queue ! autovideosink
```

You should see two video windows, showing the same sample video pattern. If you see only one, try to move it, since it is probably on top of the second window.

执行完毕后，会看到两个视频窗口，都显示相同的视频图案。
如果你只看到一个窗口，试试移动一下，因为两个窗口有可能重叠在一起了。


This example instantiates a  `videotestsrc`, linked to a  `videoconvert`, linked to a  `tee`  (Remember from  [Basic tutorial 7: Multithreading and Pad Availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html)  that a  `tee`  copies to each of its output pads everything coming through its input pad). The  `tee`  is named simply ‘t’ (using the  `name`  property) and then linked to a  `queue`  and an  `autovideosink`. The same  `tee`  is referred to using ‘t.’ (mind the dot) and then linked to a second  `queue`  and a second  `autovideosink`.

以上示例初始化了一个 `videotestsrc` 视频源, 然后视频数据输出到 `videoconvert` 转换为原始格式,然后输出到 `tee` (详细作用请复习 [Basic tutorial 7: Multithreading and Pad Availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html) 。
这个被 `name` 属性命名为 ‘t’ 的 `tee` 会把从 input sink 中收到的数据复制到每个 output src 中，然后输出到 `queue` 和 `autovideosink` 处理) 。
随后，又使用 ‘t.’ (注意最后的点) 再次引用了刚才定义的 `tee` element ，再把视频数据输出到第二个 `queue` 和第二个 `autovideosink` 处理。

To learn why the queues are necessary read  [Basic tutorial 7: Multithreading and Pad Availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html).

阅读 [Basic tutorial 7: Multithreading and Pad Availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html) 了解为什么要使用 `queue` element 。



### Pads [衬垫](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html#pads)

Instead of letting GStreamer choose which Pad to use when linking two elements, you may want to specify the Pads directly. You can do this by adding a dot plus the Pad name after the name of the element (it must be a named element). Learn the names of the Pads of an element by using the  `gst-inspect-1.0`  tool.

利用 pad 名称属性，可以自由地控制链接 element 时使用的 pad 。
使用`gst-inspect-1.0`工具可以查询 element 包含哪些已经命名的 pad 。 

This is useful, for example, when you want to retrieve one particular stream out of a demuxer:

在你希望获取 demux （解码，将音频视频分离) 后某个特定流时，这是很有用的。

```shell
gst-launch-1.0 souphttpsrc location=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm ! matroskademux name=d d.video_00 ! matroskamux ! filesink location=sintel_video.mkv
```

This fetches a media file from the internet using  `souphttpsrc`, which is in webm format 
(a special kind of Matroska container, see  [Basic tutorial 2: GStreamer concepts](https://gstreamer.freedesktop.org/documentation/tutorials/basic/concepts.html)).
We then open the container using  `matroskademux`. This media contains both audio and video, so  `matroskademux`  will create two output Pads, named  `video_00`  and  `audio_00`. We link  `video_00`  to a  `matroskamux`  element to re-pack the video stream into a new container, and finally link it to a  `filesink`, which will write the stream into a file named "sintel_video.mkv" (the  `location`  property specifies the name of the file).

以上命令使用 `souphttpsrc` 属性从互联网上获取 webm 格式的媒体文件。
(这是一种 Matroska container （容器）, 详细了解  [Basic tutorial 2: GStreamer concepts](https://gstreamer.freedesktop.org/documentation/tutorials/basic/concepts.html))。
我们使用 `matroskademux` 打开这个 container  。
因为这种媒体中同时包含音频和视频，`matroskademux` 会产生两个 output pad ，分别命名为  `video_00`  and  `audio_00` 。
把 `video_00` link 到 `matroskamux` 后会把视频流 re-pack （打包）到文件 "sintel\_video.mkv" 中（location 属性用于设置文件名称）。

TODO 此命令能否把 mp4/flv 等文件的音频提取出来？不同的媒体格式，需要使用不同的参数吗？

All in all, we took a webm file, stripped it of audio, and generated a new matroska file with the video. If we wanted to keep only the audio:

总而言之，我们将一个 webm 文件剥离音频数据，然后生成了一个仅包含视频的文件。
如果我们只想保留音频，可以这样做：

```shell
gst-launch-1.0 souphttpsrc location=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm ! matroskademux name=d d.audio_00 ! vorbisparse ! matroskamux ! filesink location=sintel_audio.mka
```

The  `vorbisparse`  element is required to extract some information from the stream and put it in the Pad Caps, so the next element,  `matroskamux`, knows how to deal with the stream. In the case of video this was not necessary, because  `matroskademux`  already extracted this information and added it to the Caps.

这里必须用 `vorbisparse` element 把媒体流中的信息提取出来放到 Pad Caps 中，以便下一个 `matroskamux` element 获取这些信息进行处理。
上面处理 video 时不用这样，是因为 `matroskademux` 已经把这些信息提取到 Caps 中了。


Note that in the above two examples no media has been decoded or played. We have just moved from one container to another (demultiplexing and re-multiplexing again).

注意，以上两条示例中，没有播放或解码任何媒体流。
我们仅仅是把它们从一个 container 移动到另一个 container 中（分解视频流，再合并视频流）。



### Caps filters [能力过滤器](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html#caps-filters)

When an element has more than one output pad, it might happen that the link to the next element is ambiguous: the next element may have more than one compatible input pad, or its input pad may be compatible with the Pad Caps of all the output pads. In these cases GStreamer will link using the first pad that is available, which pretty much amounts to saying that GStreamer will choose one output pad at random.

当 element 有多个 output pad 时， link 到另一个 element 的过程就会出现歧义：
下一个 element 也许有多个兼容的 input pad ，甚至所有的 input pad 都能兼容 out pad 的 Pad Caps 。
此时， GStreamer 会选择第一个可用的 pad ，可以理解成随机先了一个 output pad 使用。

Consider the following pipeline:

考虑以下 pipeline ：

```shell
gst-launch-1.0 souphttpsrc location=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm ! matroskademux ! filesink location=test
```

This is the same media file and demuxer as in the previous example. The input Pad Caps of  `filesink`  are  `ANY`, meaning that it can accept any kind of media. Which one of the two output pads of  `matroskademux`  will be linked against the filesink?  `video_00`  or  `audio_00`? You cannot know.

假设媒体文件与分流器与之前示例相同。
`filesink`的 input Pad Caps （输入垫的能力）是 `ANY` ，即它能接受任意类型的 media 。
那么从 `matroskademux` 分流出来的 `video_00` 和 `audio_00` 中，哪个会被链接到 filesink 呢？
不太确定吧。


You can remove this ambiguity, though, by using named pads, as in the previous sub-section, or by using  **Caps Filters**:

可以像上一节中那样使用命名 pad 解决这个歧义，还能用 **Cap Filter 能力过滤器** 。


```shell
gst-launch-1.0 souphttpsrc location=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm ! matroskademux ! video/x-vp8 ! matroskamux ! filesink location=sintel_video.mkv
```

A Caps Filter behaves like a pass-through element which does nothing and only accepts media with the given Caps, effectively resolving the ambiguity. In this example, between  `matroskademux`  and  `matroskamux`  we added a  `video/x-vp8`  Caps Filter to specify that we are interested in the output pad of  `matroskademux`  which can produce this kind of video.

Cap Filter 会直接将允许的 media 信息透传到下一个 element ，除此以外，什么也不做。
这个示例中，我们在 `matroskademux` 和 `matroskamux` 之间添加了一个 `video/x-vp8` Cap Filter ，表示，我们需要  matroskademux 输出的这种视频。


To find out the Caps an element accepts and produces, use the  `gst-inspect-1.0`  tool. To find out the Caps contained in a particular file, use the  `gst-discoverer-1.0`  tool. To find out the Caps an element is producing for a particular pipeline, run  `gst-launch-1.0`  as usual, with the  `–v`  option to print Caps information.

要知道 element 能产生或消费哪些 Cap ，可使用 gst-inspect-1.0 工具。
要知道 file 中包含哪些 Cap ，可使用 gst-discoverer-1.0 工具。
要想知道 pipeline 中的 element 具体会产生什么 Cap ，可使用使用 gst-launch-1.0 工具，加上 `-v` 选项就行。



### Examples [示例](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html#examples)

Play a media file using  `playbin`  (as in  [Basic tutorial 1: Hello world!](https://gstreamer.freedesktop.org/documentation/tutorials/basic/hello-world.html)):

使用 playbin 播放媒体文件：

```shell
gst-launch-1.0 playbin uri=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm
```

A fully operation playback pipeline, with audio and video (more or less the same pipeline that  `playbin`  will create internally):

实现与 playbin 类似的播放命令：

```shell
gst-launch-1.0 souphttpsrc location=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm ! matroskademux name=d ! queue ! vp8dec ! videoconvert ! autovideosink d. ! queue ! vorbisdec ! audioconvert ! audioresample ! autoaudiosink
```

A transcoding pipeline, which opens the webm container and decodes both streams (via uridecodebin), then re-encodes the audio and video branches with a different codec, and puts them back together in an Ogg container (just for the sake of it).

一个转换的命令，将 webm 格式的视频解码后，再重新编码。

```shell
gst-launch-1.0 uridecodebin uri=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm name=d ! queue ! theoraenc ! oggmux name=m ! filesink location=sintel.ogg d. ! queue ! audioconvert ! audioresample ! flacenc ! m.
```

A rescaling pipeline. The  `videoscale`  element performs a rescaling operation whenever the frame size is different in the input and the output caps. The output caps are set by the Caps Filter to 320x200.

一个缩放的命令：
输出 320x200 分辨率的视频。


```shell
gst-launch-1.0 uridecodebin uri=https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm ! queue ! videoscale ! video/x-raw-yuv,width=320,height=200 ! videoconvert ! autovideosink
```

This short description of  `gst-launch-1.0`  should be enough to get you started. Remember that you have the  [complete documentation available here](https://gstreamer.freedesktop.org/documentation/tools/gst-launch.html).

以上有关 `gst-launch-1.0` 的简要描述已经足够入门使用了。
更多详细内容参考[完整的文档](https://gstreamer.freedesktop.org/documentation/tools/gst-launch.html) 。


#### [videotestsrc](https://gstreamer.freedesktop.org/documentation/videotestsrc/index.html?gi-language=c)

The videotestsrc element is used to produce test video data in a wide variety of formats. The video test data produced can be controlled with the "pattern" property.

videotestsrc element 用于产生任意格式的视频。具体视频数据可以使用 "pattern" 属性控制。


By default the videotestsrc will generate data indefinitely, but if the num-buffers property is non-zero it will instead generate a fixed number of video frames and then send EOS.

videotestsrc 默认会产生无限的数据，如果给 num-buffer 属性设置了非0的值，就会产生固定数量的视频帧数。

#### [tee](https://gstreamer.freedesktop.org/documentation/coreelements/tee.html?gi-language=c#tee-page)

Split data to multiple pads. Branching the data flow is useful when e.g. capturing a video where the video is shown on the screen and also encoded and written to a file. Another example is playing music and hooking up a visualisation module.

将数据分隔成多个 pad 。
需要把数据送到多个地方进行处理时比较有用。
比如捕获视频后，既要显示到屏幕上，又要编码写入到文件中。
或者播放音乐的同时还要显示可视化的音波图形时。

One needs to use separate queue elements (or a multiqueue) in each branch to provide separate threads for each branch. Otherwise a blocked dataflow in one branch would stall the other branches.


#### [appsrc](https://gstreamer.freedesktop.org/documentation/app/appsrc.html?gi-language=c#appsrc-page)

The appsrc element can be used by applications to insert data into a GStreamer pipeline. Unlike most GStreamer elements, Appsrc provides external API functions.

appsrc element 用于从 application 中向 GStreamer pipeline 中插入数据。
与其他 GStreamer element 不同， appsrc 还提供了 API 函数。


For the documentation of the API, please see the libgstapp section in the GStreamer Plugins Base Libraries documentation.

有关 API 的文档，可以参考 GStreamer Plugin Base Libraries 文档中 libgstapp 一节。


##### [GstAppSrc](https://gstreamer.freedesktop.org/documentation/applib/gstappsrc.html?gi-language=c#page-description)

The appsrc element can be used by applications to insert data into a GStreamer pipeline. Unlike most GStreamer elements, appsrc provides external API functions.



appsrc can be used by linking with the libgstapp library to access the methods directly or by using the appsrc action signals.



Before operating appsrc, the caps property must be set to fixed caps describing the format of the data that will be pushed with appsrc. An exception to this is when pushing buffers with unknown caps, in which case no caps should be set. This is typically true of file-like sources that push raw byte buffers. If you don't want to explicitly set the caps, you can use gst_app_src_push_sample. This method gets the caps associated with the sample and sets them on the appsrc replacing any previously set caps (if different from sample's caps).

使用 appsrc 前，要设置 caps 属性，描述 push 到 appsrc 中的数据格式。
如果没有设置 caps 就 push 数据，会出现异常。
一般来自文件的数据源，直接 push 原始的 byte buffer 。
如果不相显式设置 caps 可以使用 gst_app_src_push_sample 。TODO


The main way of handing data to the appsrc element is by calling the gst_app_src_push_buffer method or by emitting the push-buffer action signal. This will put the buffer onto a queue from which appsrc will read from in its streaming thread. It is important to note that data transport will not happen from the thread that performed the push-buffer call.

调用 gst_app_src_push_buffer 方法或发送 push-buffer 信号，就会使 appsrc element 开始 handing data 。
这时 appsrc 会从 streaming thread 读取数据，然后把 buffer 放到 queue 中。
注意，传送数据的过程，不是发生在调用 push-buffer 函数的线程中。


The "max-bytes" property controls how much data can be queued in appsrc before appsrc considers the queue full. A filled internal queue will always signal the "enough-data" signal, which signals the application that it should stop pushing data into appsrc. The "block" property will cause appsrc to block the push-buffer method until free data becomes available again.

"max-bytes" 属性用于控制 apsrc 最多能进行排除的数据。超过这个数量，表示队列已经满了。
内部队列满了以后，会发出 "enough-data" 的信号，此时应该停止向 appsrc push data 。
"block"属于会使 push-buffer 方法阻塞，直到有足够之间可用时才返回。


When the internal queue is running out of data, the "need-data" signal is emitted, which signals the application that it should start pushing more data into appsrc.

当内部队列的数据即将处理完毕时，会发送 "need-data" 信号，表示 application 应该 push 更多数据到 appsrc 。


In addition to the "need-data" and "enough-data" signals, appsrc can emit the "seek-data" signal when the "stream-mode" property is set to "seekable" or "random-access". The signal argument will contain the new desired position in the stream expressed in the unit set with the "format" property. After receiving the seek-data signal, the application should push-buffers from the new position.

如果 appsrc 的 "stream-mode" 属性设置了 "seekable" 或 "random-access" 时，除 "need-data" 和 "enough-data" 信号外， appsrc 还能发出 "seek-data" 信号。
这个信息的参数中会包含所期望的数据流位置，此参数格式由 "format" 属于设置。
当收到 seek-data 信号时，应用应该 push-buffer 指定时间点的流数据。


These signals allow the application to operate the appsrc in two different ways:

这些信号允许 application 用两种不同的方法操作 appsrc ：


The push mode, in which the application repeatedly calls the push-buffer/push-sample method with a new buffer/sample. Optionally, the queue size in the appsrc can be controlled with the enough-data and need-data signals by respectively stopping/starting the push-buffer/push-sample calls. This is a typical mode of operation for the stream-type "stream" and "seekable". Use this mode when implementing various network protocols or hardware devices.

Push mode :
由 application 反复调用 push-buffer/push-sample 方法送入 buffer/sample 。
根据 enough-data 和 need-data 信号分别调用 stopping/starting 方法，以便控制 appsrc 中 push-buffer/push-sample 的队列大小。
这是 "stream" "seekable" 这种 stream-type 的典型使用方法。
这种方法适用于各种网络协议和硬件设备。


The pull mode, in which the need-data signal triggers the next push-buffer call. This mode is typically used in the "random-access" stream-type. Use this mode for file access or other randomly accessible sources. In this mode, a buffer of exactly the amount of bytes given by the need-data signal should be pushed into appsrc.

Pull mode :
只在收到 need-data 信号时，才调用 push-buffer 。
这种模式一般用于 "random-access" stream-type 。
访问文件或其他可随机访问的数据源时，可使用这种模式。
这种模式下，每次收到 need-data 信号时，只 push 所指定大小的数据到 appsrc 中。


In all modes, the size property on appsrc should contain the total stream size in bytes. Setting this property is mandatory in the random-access mode. For the stream and seekable modes, setting this property is optional but recommended.

appsrc 的 size 属性表示总的 stream byte 大小。
random-access 模式下， size 属性是必选参数。
stream 和 seekable 模式下，size 属性是可选项，但建议指定这个参数。


When the application has finished pushing data into appsrc, it should call gst_app_src_end_of_stream or emit the end-of-stream action signal. After this call, no more buffers can be pushed into appsrc until a flushing seek occurs or the state of the appsrc has gone through READY.

当 application 结束 push data 到 appsrc 时，应该调用 gst_app_src_end_of_stream 或发送 end-of-stream 信号。
而且，调用结束方法后，不能再向 appsrc push 数据，除非执行了 flushing seek ，或者 appsrc 的状态变为 READY 后。



> ## 说明
> They process the data as it flows downstream from 
> the source elements (data producers) to 
> the sink elements (data consumers), passing through 
> filter elements.
> 
> GStreamer bus  
> Your application should always keep an eye on the bus to be notified of errors and other playback-related issues.
> 



[^GStreamIBM]: [用 GStreamer 简化 Linux 多媒体开发 肖文鹏 2014](https://www.ibm.com/developerworks/cn/linux/l-gstreamer/)

[^NoteditGstrem]: [gstreamer commond](https://github.com/notedit/gstreamer-command)

[^GStreamDoc]:[gstreamer-tools](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html?gi-language=c)


