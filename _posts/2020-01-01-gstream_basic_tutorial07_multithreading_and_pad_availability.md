---
layout: post
title:  "[译]GStream Basic tutorial 07: Basic tutorial 7: Multithreading and Pad Availability"
date:   2020-01-01 12:00:00 +0800
tags:   tech
---

* category
{:toc}



# Basic tutorial 7: Multithreading and Pad Availability [^GStreamDoc]


## Goal[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#goal)

GStreamer handles multithreading automatically, but, under some circumstances, you might need to decouple threads manually. This tutorial shows how to do this and, in addition, completes the exposition about Pad Availability. More precisely, this document explains:

GStreamer 会自动管理线程，但某些情况要，需要进行人为干涉。
这篇文章用于说明如何手动管理线程，另外还会涉及一部分 Pad Avaliability 的概念。
准确的说，本文主要包含以下内容：

- How to create new threads of execution for some parts of the pipeline
- What is the Pad Availability
- How to replicate streams

- 如何在 pipeline 中创建 thread
- Pad Availability 是什么
- 如何复制 stream
    

## Introduction[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#introduction)

### Multithreading[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#multithreading)

GStreamer is a multithreaded framework. This means that, internally, it creates and destroys threads as it needs them, for example, to decouple streaming from the application thread. Moreover, plugins are also free to create threads for their own processing, for example, a video decoder could create 4 threads to take full advantage of a CPU with 4 cores.

GStreamer 是一个多线程 framework 。它会在内部按需要创建销毁线程，比如，在 application 线程中分离 stream 时。
另外，plugin 也会在他们自己的进程中创建线程，比如 video decoder 会创建 4 个线程来更有效率的得用四核心的 CPU 。


On top of this, when building the pipeline an application can specify explicitly that a  _branch_  (a part of the pipeline) runs on a different thread (for example, to have the audio and video decoders executing simultaneously).

除此之外，我们在创建 pipeline 时，可以指定某个 _branch_ （即 pipeline 的一部分) 运行在独立的线程上（比如，将 audio video 解码器分别运行在不同的线程中，确保音视频并行处理）。

This is accomplished using the  `queue`  element, which works as follows. The sink pad just enqueues data and returns control. On a different thread, data is dequeued and pushed downstream. This element is also used for buffering, as seen later in the streaming tutorials. The size of the queue can be controlled through properties.

可以使用 `queue` element 实现。
向 queue 的 sink 在送入数据后，会立即返回。
会有另外一个线程取出数据，并送到后续的工作流中使用。
这个 element 也可以用作缓冲，可以在 streaming tutorial 中看到相关例子。
queue 的大小可以通过 property 控制。

> 译：可这样理解，向 queue 的 sink 送入数据后，当前线程的函数调用立即返回。 queue 对应的工作线程会并行从 queue 中取出数据，异步处理。有点像 async write 。


### The example pipeline[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#the-example-pipeline)

This example builds the following pipeline:

> 译：功能类似的 pipline 如下：
> ` gst-launch-1.0 audiotestsrc freq=215 ! tee name=t ! queue ! audioconvert ! audioresample ! autoaudiosink t. ! quque ! wavescope shader=0 style=1 ! videoconvert ! autovideosink `

![](https://gstreamer.freedesktop.org/documentation/tutorials/basic/images/tutorials/basic-tutorial-7.png)

The source is a synthetic audio signal (a continuous tone) which is split using a  `tee`  element (it sends through its source pads everything it receives through its sink pad). One branch then sends the signal to the audio card, and the other renders a video of the waveform and sends it to the screen.

source 是连续的声音信号。
使用 `tee` element 将它从 sink pad 收到的数据原样转发到 source pad 。
可以把一份数据副本送到声卡播放。
再把另一份数据副本的声音波形图绘制到屏幕显示出来。


As seen in the picture, queues create a new thread, so this pipeline runs in 3 threads. Pipelines with more than one sink usually need to be multithreaded, because, to be synchronized, sinks usually block execution until all other sinks are ready, and they cannot get ready if there is only one thread, being blocked by the first sink.

queue 用于创建 thread ，如上图所示的 pipeline 运行在 3 个 thread 中。
包含超过一个 sink 的 pipeline 通常都要运行在 multithread 中，为了保持同步，必须保证所有 sink 都处于 ready 状态时，才会继续执行，否则会 block （阻塞）。
如果只有一个 thread 那么所有 sink 都会 block 。


### Request pads[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#request-pads)

In  [Basic tutorial 3: Dynamic pipelines](https://gstreamer.freedesktop.org/documentation/tutorials/basic/dynamic-pipelines.html)  we saw an element (`uridecodebin`) which had no pads to begin with, and they appeared as data started to flow and the element learned about the media. These are called  **Sometimes Pads**, and contrast with the regular pads which are always available and are called  **Always Pads**.

在 Basic tutorial 3: Dynamic pipelines 中能看到 (`uridecodebin`) element 是没有 pads 的， 但当媒体流数据出现，并且获取到相关媒体信息时，就会产生  **Sometimes Pads** ，与以对应，一直存在的 pad 称作  **Always Pads** 。

The third kind of pad is the  **Request Pad**, which is created on demand. The classical example is the  `tee`  element, which has one sink pad and no initial source pads: they need to be requested and then  `tee`  adds them. In this way, an input stream can be replicated any number of times. The disadvantage is that linking elements with Request Pads is not as automatic, as linking Always Pads, as the walkthrough for this example will show.

还有一类按需创建的 pad 称为 **Request Pad** 。
典型的例子就是 `tee` element ，它只有一个 sink pad ，默认没有任何 source pad ：在收到相关 request 后，才会自动创建 source pad 。
这样就能复制多份 input stream 。
缺点是 link element 与 Request Pad 的过程不是自动完成的，后面的示例会说明具体如何操作。


Also, to request (or release) pads in the  `PLAYING`  or  `PAUSED`  states, you need to take additional cautions (Pad blocking) which are not described in this tutorial. It is safe to request (or release) pads in the  `NULL`  or  `READY`  states, though.

另外，在 PLAYING PAUSED 状态时 request (or release) pad ，要注意防止 Pad blocking ，本文不会详细说明这方面的内容。
在 NULL READY 状态时 request pad 是安全的。

Without further delay, let's see the code.

下面直接看代码吧。



## Simple multithreaded example[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#simple-multithreaded-example)

Copy this code into a text file named  `basic-tutorial-7.c`  (or find it in your GStreamer installation).

**basic-tutorial-7.c**

```c
#include <gst/gst.h>

int main(int argc, char *argv[]) {
  GstElement *pipeline, *audio_source, *tee, *audio_queue, *audio_convert, *audio_resample, *audio_sink;
  GstElement *video_queue, *visual, *video_convert, *video_sink;
  GstBus *bus;
  GstMessage *msg;
  GstPad *tee_audio_pad, *tee_video_pad;
  GstPad *queue_audio_pad, *queue_video_pad;

  /* Initialize GStreamer */
  gst_init (&argc, &argv);

  /* Create the elements */
  audio_source = gst_element_factory_make ("audiotestsrc", "audio_source");
  tee = gst_element_factory_make ("tee", "tee");
  audio_queue = gst_element_factory_make ("queue", "audio_queue");
  audio_convert = gst_element_factory_make ("audioconvert", "audio_convert");
  audio_resample = gst_element_factory_make ("audioresample", "audio_resample");
  audio_sink = gst_element_factory_make ("autoaudiosink", "audio_sink");
  video_queue = gst_element_factory_make ("queue", "video_queue");
  visual = gst_element_factory_make ("wavescope", "visual");
  video_convert = gst_element_factory_make ("videoconvert", "csp");
  video_sink = gst_element_factory_make ("autovideosink", "video_sink");

  /* Create the empty pipeline */
  pipeline = gst_pipeline_new ("test-pipeline");

  if (!pipeline || !audio_source || !tee || !audio_queue || !audio_convert || !audio_resample || !audio_sink ||
      !video_queue || !visual || !video_convert || !video_sink) {
    g_printerr ("Not all elements could be created.\n");
    return -1;
  }

  /* Configure elements */
  g_object_set (audio_source, "freq", 215.0f, NULL);
  g_object_set (visual, "shader", 0, "style", 1, NULL);

  /* Link all elements that can be automatically linked because they have "Always" pads */
  gst_bin_add_many (GST_BIN (pipeline), audio_source, tee, audio_queue, audio_convert, audio_resample, audio_sink,
      video_queue, visual, video_convert, video_sink, NULL);
  if (gst_element_link_many (audio_source, tee, NULL) != TRUE ||
      gst_element_link_many (audio_queue, audio_convert, audio_resample, audio_sink, NULL) != TRUE ||
      gst_element_link_many (video_queue, visual, video_convert, video_sink, NULL) != TRUE) {
    g_printerr ("Elements could not be linked.\n");
    gst_object_unref (pipeline);
    return -1;
  }

  /* Manually link the Tee, which has "Request" pads */
  tee_audio_pad = gst_element_get_request_pad (tee, "src_%u");
  g_print ("Obtained request pad %s for audio branch.\n", gst_pad_get_name (tee_audio_pad));
  queue_audio_pad = gst_element_get_static_pad (audio_queue, "sink");
  tee_video_pad = gst_element_get_request_pad (tee, "src_%u");
  g_print ("Obtained request pad %s for video branch.\n", gst_pad_get_name (tee_video_pad));
  queue_video_pad = gst_element_get_static_pad (video_queue, "sink");
  if (gst_pad_link (tee_audio_pad, queue_audio_pad) != GST_PAD_LINK_OK ||
      gst_pad_link (tee_video_pad, queue_video_pad) != GST_PAD_LINK_OK) {
    g_printerr ("Tee could not be linked.\n");
    gst_object_unref (pipeline);
    return -1;
  }
  gst_object_unref (queue_audio_pad);
  gst_object_unref (queue_video_pad);

  /* Start playing the pipeline */
  gst_element_set_state (pipeline, GST_STATE_PLAYING);

  /* Wait until error or EOS */
  bus = gst_element_get_bus (pipeline);
  msg = gst_bus_timed_pop_filtered (bus, GST_CLOCK_TIME_NONE, GST_MESSAGE_ERROR | GST_MESSAGE_EOS);

  /* Release the request pads from the Tee, and unref them */
  gst_element_release_request_pad (tee, tee_audio_pad);
  gst_element_release_request_pad (tee, tee_video_pad);
  gst_object_unref (tee_audio_pad);
  gst_object_unref (tee_video_pad);

  /* Free resources */
  if (msg != NULL)
    gst_message_unref (msg);
  gst_object_unref (bus);
  gst_element_set_state (pipeline, GST_STATE_NULL);

  gst_object_unref (pipeline);
  return 0;
}

```

> Need help?
> 
> If you need help to compile this code, refer to the  **Building the tutorials**  section for your platform:  [Linux](https://gstreamer.freedesktop.org/documentation/installing/on-linux.html#InstallingonLinux-Build),  [Mac OS X](https://gstreamer.freedesktop.org/documentation/installing/on-mac-osx.html#InstallingonMacOSX-Build)  or  [Windows](https://gstreamer.freedesktop.org/documentation/installing/on-windows.html#InstallingonWindows-Build), or use this specific command on Linux:
> 
> ``gcc basic-tutorial-7.c -o basic-tutorial-7 `pkg-config --cflags --libs gstreamer-1.0` ``
> 
> If you need help to run this code, refer to the  **Running the tutorials**  section for your platform:  [Linux](https://gstreamer.freedesktop.org/documentation/installing/on-linux.html#InstallingonLinux-Run),  [Mac OS X](https://gstreamer.freedesktop.org/documentation/installing/on-mac-osx.html#InstallingonMacOSX-Run)  or  [Windows](https://gstreamer.freedesktop.org/documentation/installing/on-windows.html#InstallingonWindows-Run).
> 
> This tutorial plays an audible tone through the audio card and opens a window with a waveform representation of the tone. The waveform should be a sinusoid, but due to the refreshing of the window might not appear so.
> 
> Required libraries:  `gstreamer-1.0`

## Walkthrough[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#walkthrough)

```c
/* Create the elements */
audio_source = gst_element_factory_make ("audiotestsrc", "audio_source");
tee = gst_element_factory_make ("tee", "tee");
audio_queue = gst_element_factory_make ("queue", "audio_queue");
audio_convert = gst_element_factory_make ("audioconvert", "audio_convert");
audio_resample = gst_element_factory_make ("audioresample", "audio_resample");
audio_sink = gst_element_factory_make ("autoaudiosink", "audio_sink");
video_queue = gst_element_factory_make ("queue", "video_queue");
visual = gst_element_factory_make ("wavescope", "visual");
video_convert = gst_element_factory_make ("videoconvert", "video_convert");
video_sink = gst_element_factory_make ("autovideosink", "video_sink");

```

All the elements in the above picture are instantiated here:

`audiotestsrc`  produces a synthetic tone.
`wavescope`  consumes an audio signal and renders a waveform as if it was an (admittedly cheap) oscilloscope.
We have already worked with the  `autoaudiosink`  and  `autovideosink`.

`audiotestsrc` 可以产生电子音调。

`wavescope` 消费音频信号，并渲染出它的波形，可理解成一个简易示波器。

`autoaudiosink` 会自动播放音频（自动给声卡发信号，播放声音）。

`autovideosink` 会自动播放视频（自动打开一个窗口，显示这个视频画面）。


The conversion elements (`audioconvert`,  `audioresample`  and  `videoconvert`) are necessary to guarantee that the pipeline can be linked. Indeed, the Capabilities of the audio and video sinks depend on the hardware, and you do not know at design time if they will match the Caps produced by the  `audiotestsrc`  and  `wavescope`. If the Caps matched, though, these elements act in “pass-through” mode and do not modify the signal, having negligible impact on performance.

conversion element (audioconvert, audioresample, videoconvert) 是为了保证 pipeline 能顺利 link 在一起。
因为具体的音视频与硬件有关，设计命令时，无法确定 audiotestsrc 产生的视频格式是否与 wavescope 支持的格式匹配。
如果格式匹配， conversion element 会直接透传数据，不做任何修改，对性能的影响可以忽略不计。

```c
/* Configure elements */
g_object_set (audio_source, "freq", 215.0f, NULL);
g_object_set (visual, "shader", 0, "style", 1, NULL);

```

Small adjustments for better demonstration: The “freq” property of  `audiotestsrc`  controls the frequency of the wave (215Hz makes the wave appear almost stationary in the window), and this style and shader for  `wavescope`  make the wave continuous. Use the  `gst-inspect-1.0`  tool described in  [Basic tutorial 10: GStreamer tools](https://gstreamer.freedesktop.org/documentation/tutorials/basic/gstreamer-tools.html)  to learn all the properties of these elements.

```c
/* Link all elements that can be automatically linked because they have "Always" pads */
gst_bin_add_many (GST_BIN (pipeline), audio_source, tee, audio_queue, audio_convert, audio_sink,
    video_queue, visual, video_convert, video_sink, NULL);
if (gst_element_link_many (audio_source, tee, NULL) != TRUE ||
    gst_element_link_many (audio_queue, audio_convert, audio_sink, NULL) != TRUE ||
    gst_element_link_many (video_queue, visual, video_convert, video_sink, NULL) != TRUE) {
  g_printerr ("Elements could not be linked.\n");
  gst_object_unref (pipeline);
  return -1;
}

```

This code block adds all elements to the pipeline and then links the ones that can be automatically linked (the ones with Always Pads, as the comment says).

> ![Warning](https://gstreamer.freedesktop.org/documentation/tutorials/basic/images/icons/emoticons/warning.svg)`gst_element_link_many()`  can actually link elements with Request Pads. It internally requests the Pads so you do not have worry about the elements being linked having Always or Request Pads. Strange as it might seem, this is actually inconvenient, because you still need to release the requested Pads afterwards, and, if the Pad was requested automatically by  `gst_element_link_many()`, it is easy to forget. Stay out of trouble by always requesting Request Pads manually, as shown in the next code block.

```c
/* Manually link the Tee, which has "Request" pads */
tee_audio_pad = gst_element_get_request_pad (tee, "src_%u");
g_print ("Obtained request pad %s for audio branch.\n", gst_pad_get_name (tee_audio_pad));
queue_audio_pad = gst_element_get_static_pad (audio_queue, "sink");
tee_video_pad = gst_element_get_request_pad (tee, "src_%u");
g_print ("Obtained request pad %s for video branch.\n", gst_pad_get_name (tee_video_pad));
queue_video_pad = gst_element_get_static_pad (video_queue, "sink");
if (gst_pad_link (tee_audio_pad, queue_audio_pad) != GST_PAD_LINK_OK ||
    gst_pad_link (tee_video_pad, queue_video_pad) != GST_PAD_LINK_OK) {
  g_printerr ("Tee could not be linked.\n");
  gst_object_unref (pipeline);
  return -1;
}
gst_object_unref (queue_audio_pad);
gst_object_unref (queue_video_pad);

```

To link Request Pads, they need to be obtained by “requesting” them to the element. An element might be able to produce different kinds of Request Pads, so, when requesting them, the desired Pad Template name must be provided. In the documentation for the  `tee`  element we see that it has two pad templates named “sink” (for its sink Pads) and “src_%u” (for the Request Pads). We request two Pads from the tee (for the audio and video branches) with  `gst_element_get_request_pad()`.

We then obtain the Pads from the downstream elements to which these Request Pads need to be linked. These are normal Always Pads, so we obtain them with  `gst_element_get_static_pad()`.

Finally, we link the pads with  `gst_pad_link()`. This is the function that  `gst_element_link()`  and  `gst_element_link_many()`  use internally.

The sink Pads we have obtained need to be released with  `gst_object_unref()`. The Request Pads will be released when we no longer need them, at the end of the program.

We then set the pipeline to playing as usual, and wait until an error message or an EOS is produced. The only thing left to so is cleanup the requested Pads:

```c
/* Release the request pads from the Tee, and unref them */
gst_element_release_request_pad (tee, tee_audio_pad);
gst_element_release_request_pad (tee, tee_video_pad);
gst_object_unref (tee_audio_pad);
gst_object_unref (tee_video_pad);

```

`gst_element_release_request_pad()`  releases the pad from the  `tee`, but it still needs to be unreferenced (freed) with  `gst_object_unref()`.

## Conclusion[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#conclusion)

This tutorial has shown:

-   How to make parts of a pipeline run on a different thread by using  `queue`  elements.
    
-   What is a Request Pad and how to link elements with request pads, with  `gst_element_get_request_pad()`,  `gst_pad_link()`  and  `gst_element_release_request_pad()`.
    
-   How to have the same stream available in different branches by using  `tee`  elements.
    

The next tutorial builds on top of this one to show how data can be manually injected into and extracted from a running pipeline.

It has been a pleasure having you here, and see you soon!



> ## 说明
> 相同作用的 pipline 如下
> gst-launch-1.0 audiotestsrc freq=215 ! tee name=t ! queue ! audioconvert ! audioresample ! autoaudiosink t. ! queue ! wavescope shader=0 style=1 ! videoconvert ! autovideosink
> 
> 


[^GStreamDoc]:[multithreading-and-pad-availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html#basic-tutorial-7-multithreading-and-pad-availability)

