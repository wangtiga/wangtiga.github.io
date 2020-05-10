---
layout: post
title:  "[译]GStream Basic tutorial 08: Short-cutting the pipeline"
date:   2020-05-09 22:00:00 +0800
tags:   tech
---

* category
{:toc}




# Basic tutorial 8: Short-cutting the pipeline 简介 pipeline [^GStreamDoc]:

## Goal 目标

Pipelines constructed with GStreamer do not need to be completely closed. Data can be injected into the pipeline and extracted from it at any time, in a variety of ways. This tutorial shows:

GStreamer 的 Pipeline 组件是开放的。
可以通过多种手段向 pipeline 中注入数据，或者从 pipeline 中取出数据。
本教程包含以下内容：

-   How to inject external data into a general GStreamer pipeline.
    
-   How to extract data from a general GStreamer pipeline.
    
-   How to access and manipulate this data.

- 如果将外部数据注入到 GStreamer pipeline 中。
- 如何从 GStreamer pipeline 中提取数据。
- 如何访问并修改这个数据。
    

[Playback tutorial 3: Short-cutting the pipeline](https://gstreamer.freedesktop.org/documentation/tutorials/playback/short-cutting-the-pipeline.html)  explains how to achieve the same goals in a playbin-based pipeline.

[Playback tutorial 3: Short-cutting the pipeline](https://gstreamer.freedesktop.org/documentation/tutorials/playback/short-cutting-the-pipeline.html) 教程说明了如何在 playbin-based pipeline 中实现类似目标的手段 。



## Introduction 介绍


Applications can interact with the data flowing through a GStreamer pipeline in several ways. This tutorial describes the easiest one, since it uses elements that have been created for this sole purpose.

应用程序可用多种方式与 GStreamer pipeline 的数据进行交互。
此教程只用最简单的方式进行演示。

The element used to inject application data into a GStreamer pipeline is  `appsrc`, and its counterpart, used to extract GStreamer data back to the application is  `appsink`.
To avoid confusing the names, think of it from GStreamer's point of view:  `appsrc`  is just a regular source, that provides data magically fallen from the sky (provided by the application, actually).  `appsink`  is a regular sink, where the data flowing through a GStreamer pipeline goes to die (it is recovered by the application, actually).

用于向 GStreamer pipeline 注入数据的 element 是 `appsrc`，与之对应，
用于从 GStreamer pipeline 提取数据的 element 是 `appsink` 。
可以这样理解：
`appsrc`是数据的源头，即数据来自 application ，流向GStreamer 。
`appsink`是数据的目的地，即数据从 GStreamer 流出，最终流向水槽终止，实际上是流到了 application ，交由应用程序接收并处理媒体流数据）。


`appsrc`  and  `appsink`  are so versatile that they offer their own API (see their documentation), which can be accessed by linking against the  `gstreamer-app`  library. In this tutorial, however, we will use a simpler approach and control them through signals.

可以通过`gstreamer-app` 调用  `appsrc` 和  `appsink` 的完整功能。
这篇教程中，我们仅通过信号实现一些简单的控制功能。

`appsrc`  can work in a variety of modes: in  **pull**  mode, it requests data from the application every time it needs it. In  **push**  mode, the application pushes data at its own pace. Furthermore, in push mode, the application can choose to be blocked in the push function when enough data has already been provided, or it can listen to the  `enough-data`  and  `need-data`  signals to control flow. This example implements the latter approach. Information regarding the other methods can be found in the  `appsrc`  documentation.


`appsrc` 可工作在多种模式中。
**pull** 模式，每次都从 application 中请求数据。
**push** 模式，由 application 按自己的节奏向 GStreamer 推送数据。
push 模式时，如果 GStreamer 所需要的数据已经足够， 可以让 application 调用 push 函数时阻塞，或者 application 监听 enough-data 与 need-data 信号，来调整推送数据的过程。
下面的示例中，使用的第二种方案。
其他方法可以在 `appsrc` 的文档中查看。



### Buffers 缓冲


Data travels through a GStreamer pipeline in chunks called  **buffers**. Since this example produces and consumes data, we need to know about  `GstBuffer`s.

在 GStreamer pipeline 中分块处理的数据被称为 **buffers** 。
在下面的示例中会产生和消费一些数据，所以我们要了解 GstBuffer 的概念。


Source Pads produce buffers, that are consumed by Sink Pads; GStreamer takes these buffers and passes them from element to element.

Source Pad 会产生 buffer ，这些 buffer 可以被 Sink Pad 消费掉。
GStreamer 管理这些 buffer ，将他们从一个 element 传递到另外一个 element 中。



A buffer simply represents a unit of data, do not assume that all buffers will have the same size, or represent the same amount of time. Neither should you assume that if a single buffer enters an element, a single buffer will come out. Elements are free to do with the received buffers as they please.  `GstBuffer`s may also contain more than one actual memory buffer. Actual memory buffers are abstracted away using  `GstMemory`  objects, and a  `GstBuffer`  can contain multiple  `GstMemory`  objects.

一个 buffer 表示一个数据单元，每个 buffer 的大小，持续时长（媒体流的时间）都有可能是不同的。
也不要假设输入一个 buffer 到 element 就一定会输出一个 buffer 。
element 如何处理自己收到的 buffer 是不确定的。
`GstBuffer`可以包含多个真实的 membory buffer 。
真实的 memory buffer 是由 `GstMemory` 对象表示的， `GstBuffer`可以包含多个 `GstMemory` 对象。


Every buffer has attached time-stamps and duration, that describe in which moment the content of the buffer should be decoded, rendered or displayed. Time stamping is a very complex and delicate subject, but this simplified vision should suffice for now.

每个 buffer 都附带时间戳和时长两种信息，用于在解码，渲染及显示时使用。
时间戳是一个复杂又微妙的旆是，但我们这里仅关心最简单的用法。


As an example, a  `filesrc`  (a GStreamer element that reads files) produces buffers with the “ANY” caps and no time-stamping information. 
After demuxing (see  [Basic tutorial 3: Dynamic pipelines](https://gstreamer.freedesktop.org/documentation/tutorials/basic/dynamic-pipelines.html)) buffers can have some specific caps, 
for example “video/x-h264”. After decoding, each buffer will contain a single video frame with raw caps (for example, “video/x-raw-yuv”) and very precise time stamps indicating when should that frame be displayed.

比如，`filesrc`（用于读取文件的 GStreamer element ）会产生 caps 为 "ANY" 且没有时间戳的  buffer 。
但 demuxing buffers 后 （参考 [Basic tutorial 3: Dynamic pipelines](https://gstreamer.freedesktop.org/documentation/tutorials/basic/dynamic-pipelines.html)) 就会附带一些 caps ，比如，"video/x-h264" 。
decode 解码后，每个 buffer 就会包含一个 caps=raw 的原始视频帧（比如 "video/x-raw-yuv" ），和表示视频帧显示时间的精确时间戳。



### This tutorial 教程

This tutorial expands  [Basic tutorial 7: Multithreading and Pad Availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html)  in two ways: 
firstly, the  `audiotestsrc`  is replaced by an  `appsrc`  that will generate the audio data.
Secondly, a new branch is added to the  `tee`  so data going into the audio sink and the wave display is also replicated into an  `appsink`. 
The  `appsink`  uploads the information back into the application, which then just notifies the user that data has been received, but it could obviously perform more complex tasks.

本教程将详细展开在 [Basic tutorial 7: Multithreading and Pad Availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html) 的两方面内容：
首先，使用 `appsrc` 替换 `audiotestsrc` 生成音频数据。

其次，在 `tee` 新增一个分支，以便一份数据送到 audio sink ，还有一份数据用于显示音频波形，将会送到 `appsink` 。TODO 显示波形的应该是视频数据吧。

送到 `appsink` 的数据，其实又回到 application 了，我们这里仅在收到数据时输出日志提示了一个，实现上，我们可以对这些数据做任何复杂的处理。


![](https://gstreamer.freedesktop.org/documentation/tutorials/basic/images/tutorials/basic-tutorial-8.png)



## A crude waveform generator 粗糙的波形生成器

Copy this code into a text file named  `basic-tutorial-8.c`  (or find it in your GStreamer installation).

将以下代码复制到文本文件，并命名为 `basic-tutorial-8.c` 。

```c
#include <gst/gst.h>
#include <gst/audio/audio.h>
#include <string.h>

#define CHUNK_SIZE 1024   /* Amount of bytes we are sending in each buffer */
#define SAMPLE_RATE 44100 /* Samples per second we are sending */

/* Structure to contain all our information, so we can pass it to callbacks */
typedef struct _CustomData {
  GstElement *pipeline, *app_source, *tee, *audio_queue, *audio_convert1, *audio_resample, *audio_sink;
  GstElement *video_queue, *audio_convert2, *visual, *video_convert, *video_sink;
  GstElement *app_queue, *app_sink;

  guint64 num_samples;   /* Number of samples generated so far (for timestamp generation) */
  gfloat a, b, c, d;     /* For waveform generation */

  guint sourceid;        /* To control the GSource */

  GMainLoop *main_loop;  /* GLib's Main Loop */
} CustomData;

/* This method is called by the idle GSource in the mainloop, to feed CHUNK_SIZE bytes into appsrc.
 * The idle handler is added to the mainloop when appsrc requests us to start sending data (need-data signal)
 * and is removed when appsrc has enough data (enough-data signal).
 */
static gboolean push_data (CustomData *data) {
  GstBuffer *buffer;
  GstFlowReturn ret;
  int i;
  GstMapInfo map;
  gint16 *raw;
  gint num_samples = CHUNK_SIZE / 2; /* Because each sample is 16 bits */
  gfloat freq;

  /* Create a new empty buffer */
  buffer = gst_buffer_new_and_alloc (CHUNK_SIZE);

  /* Set its timestamp and duration */
  GST_BUFFER_TIMESTAMP (buffer) = gst_util_uint64_scale (data->num_samples, GST_SECOND, SAMPLE_RATE);
  GST_BUFFER_DURATION (buffer) = gst_util_uint64_scale (num_samples, GST_SECOND, SAMPLE_RATE);

  /* Generate some psychodelic waveforms */
  gst_buffer_map (buffer, &map, GST_MAP_WRITE);
  raw = (gint16 *)map.data;
  data->c += data->d;
  data->d -= data->c / 1000;
  freq = 1100 + 1000 * data->d;
  for (i = 0; i < num_samples; i++) {
    data->a += data->b;
    data->b -= data->a / freq;
    raw[i] = (gint16)(500 * data->a);
  }
  gst_buffer_unmap (buffer, &map);
  data->num_samples += num_samples;

  /* Push the buffer into the appsrc */
  g_signal_emit_by_name (data->app_source, "push-buffer", buffer, &ret);

  /* Free the buffer now that we are done with it */
  gst_buffer_unref (buffer);

  if (ret != GST_FLOW_OK) {
    /* We got some error, stop sending data */
    return FALSE;
  }

  return TRUE;
}

/* This signal callback triggers when appsrc needs data. Here, we add an idle handler
 * to the mainloop to start pushing data into the appsrc */
static void start_feed (GstElement *source, guint size, CustomData *data) {
  if (data->sourceid == 0) {
    g_print ("Start feeding\n");
    data->sourceid = g_idle_add ((GSourceFunc) push_data, data);
  }
}

/* This callback triggers when appsrc has enough data and we can stop sending.
 * We remove the idle handler from the mainloop */
static void stop_feed (GstElement *source, CustomData *data) {
  if (data->sourceid != 0) {
    g_print ("Stop feeding\n");
    g_source_remove (data->sourceid);
    data->sourceid = 0;
  }
}

/* The appsink has received a buffer */
static GstFlowReturn new_sample (GstElement *sink, CustomData *data) {
  GstSample *sample;

  /* Retrieve the buffer */
  g_signal_emit_by_name (sink, "pull-sample", &sample);
  if (sample) {
    /* The only thing we do in this example is print a * to indicate a received buffer */
    g_print ("*");
    gst_sample_unref (sample);
    return GST_FLOW_OK;
  }

  return GST_FLOW_ERROR;
}

/* This function is called when an error message is posted on the bus */
static void error_cb (GstBus *bus, GstMessage *msg, CustomData *data) {
  GError *err;
  gchar *debug_info;

  /* Print error details on the screen */
  gst_message_parse_error (msg, &err, &debug_info);
  g_printerr ("Error received from element %s: %s\n", GST_OBJECT_NAME (msg->src), err->message);
  g_printerr ("Debugging information: %s\n", debug_info ? debug_info : "none");
  g_clear_error (&err);
  g_free (debug_info);

  g_main_loop_quit (data->main_loop);
}

int main(int argc, char *argv[]) {
  CustomData data;
  GstPad *tee_audio_pad, *tee_video_pad, *tee_app_pad;
  GstPad *queue_audio_pad, *queue_video_pad, *queue_app_pad;
  GstAudioInfo info;
  GstCaps *audio_caps;
  GstBus *bus;

  /* Initialize cumstom data structure */
  memset (&data, 0, sizeof (data));
  data.b = 1; /* For waveform generation */
  data.d = 1;

  /* Initialize GStreamer */
  gst_init (&argc, &argv);

  /* Create the elements */
  data.app_source = gst_element_factory_make ("appsrc", "audio_source");
  data.tee = gst_element_factory_make ("tee", "tee");
  data.audio_queue = gst_element_factory_make ("queue", "audio_queue");
  data.audio_convert1 = gst_element_factory_make ("audioconvert", "audio_convert1");
  data.audio_resample = gst_element_factory_make ("audioresample", "audio_resample");
  data.audio_sink = gst_element_factory_make ("autoaudiosink", "audio_sink");
  data.video_queue = gst_element_factory_make ("queue", "video_queue");
  data.audio_convert2 = gst_element_factory_make ("audioconvert", "audio_convert2");
  data.visual = gst_element_factory_make ("wavescope", "visual");
  data.video_convert = gst_element_factory_make ("videoconvert", "video_convert");
  data.video_sink = gst_element_factory_make ("autovideosink", "video_sink");
  data.app_queue = gst_element_factory_make ("queue", "app_queue");
  data.app_sink = gst_element_factory_make ("appsink", "app_sink");

  /* Create the empty pipeline */
  data.pipeline = gst_pipeline_new ("test-pipeline");

  if (!data.pipeline || !data.app_source || !data.tee || !data.audio_queue || !data.audio_convert1 ||
      !data.audio_resample || !data.audio_sink || !data.video_queue || !data.audio_convert2 || !data.visual ||
      !data.video_convert || !data.video_sink || !data.app_queue || !data.app_sink) {
    g_printerr ("Not all elements could be created.\n");
    return -1;
  }

  /* Configure wavescope */
  g_object_set (data.visual, "shader", 0, "style", 0, NULL);

  /* Configure appsrc */
  gst_audio_info_set_format (&info, GST_AUDIO_FORMAT_S16, SAMPLE_RATE, 1, NULL);
  audio_caps = gst_audio_info_to_caps (&info);
  g_object_set (data.app_source, "caps", audio_caps, "format", GST_FORMAT_TIME, NULL);
  g_signal_connect (data.app_source, "need-data", G_CALLBACK (start_feed), &data);
  g_signal_connect (data.app_source, "enough-data", G_CALLBACK (stop_feed), &data);

  /* Configure appsink */
  g_object_set (data.app_sink, "emit-signals", TRUE, "caps", audio_caps, NULL);
  g_signal_connect (data.app_sink, "new-sample", G_CALLBACK (new_sample), &data);
  gst_caps_unref (audio_caps);

  /* Link all elements that can be automatically linked because they have "Always" pads */
  gst_bin_add_many (GST_BIN (data.pipeline), data.app_source, data.tee, data.audio_queue, data.audio_convert1, data.audio_resample,
      data.audio_sink, data.video_queue, data.audio_convert2, data.visual, data.video_convert, data.video_sink, data.app_queue,
      data.app_sink, NULL);
  if (gst_element_link_many (data.app_source, data.tee, NULL) != TRUE ||
      gst_element_link_many (data.audio_queue, data.audio_convert1, data.audio_resample, data.audio_sink, NULL) != TRUE ||
      gst_element_link_many (data.video_queue, data.audio_convert2, data.visual, data.video_convert, data.video_sink, NULL) != TRUE ||
      gst_element_link_many (data.app_queue, data.app_sink, NULL) != TRUE) {
    g_printerr ("Elements could not be linked.\n");
    gst_object_unref (data.pipeline);
    return -1;
  }

  /* Manually link the Tee, which has "Request" pads */
  tee_audio_pad = gst_element_get_request_pad (data.tee, "src_%u");
  g_print ("Obtained request pad %s for audio branch.\n", gst_pad_get_name (tee_audio_pad));
  queue_audio_pad = gst_element_get_static_pad (data.audio_queue, "sink");
  tee_video_pad = gst_element_get_request_pad (data.tee, "src_%u");
  g_print ("Obtained request pad %s for video branch.\n", gst_pad_get_name (tee_video_pad));
  queue_video_pad = gst_element_get_static_pad (data.video_queue, "sink");
  tee_app_pad = gst_element_get_request_pad (data.tee, "src_%u");
  g_print ("Obtained request pad %s for app branch.\n", gst_pad_get_name (tee_app_pad));
  queue_app_pad = gst_element_get_static_pad (data.app_queue, "sink");
  if (gst_pad_link (tee_audio_pad, queue_audio_pad) != GST_PAD_LINK_OK ||
      gst_pad_link (tee_video_pad, queue_video_pad) != GST_PAD_LINK_OK ||
      gst_pad_link (tee_app_pad, queue_app_pad) != GST_PAD_LINK_OK) {
    g_printerr ("Tee could not be linked\n");
    gst_object_unref (data.pipeline);
    return -1;
  }
  gst_object_unref (queue_audio_pad);
  gst_object_unref (queue_video_pad);
  gst_object_unref (queue_app_pad);

  /* Instruct the bus to emit signals for each received message, and connect to the interesting signals */
  bus = gst_element_get_bus (data.pipeline);
  gst_bus_add_signal_watch (bus);
  g_signal_connect (G_OBJECT (bus), "message::error", (GCallback)error_cb, &data);
  gst_object_unref (bus);

  /* Start playing the pipeline */
  gst_element_set_state (data.pipeline, GST_STATE_PLAYING);

  /* Create a GLib Main Loop and set it to run */
  data.main_loop = g_main_loop_new (NULL, FALSE);
  g_main_loop_run (data.main_loop);

  /* Release the request pads from the Tee, and unref them */
  gst_element_release_request_pad (data.tee, tee_audio_pad);
  gst_element_release_request_pad (data.tee, tee_video_pad);
  gst_element_release_request_pad (data.tee, tee_app_pad);
  gst_object_unref (tee_audio_pad);
  gst_object_unref (tee_video_pad);
  gst_object_unref (tee_app_pad);

  /* Free resources */
  gst_element_set_state (data.pipeline, GST_STATE_NULL);
  gst_object_unref (data.pipeline);
  return 0;
}

```

> ![Information](https://gstreamer.freedesktop.org/documentation/tutorials/basic/images/icons/emoticons/information.svg)Need help?
> 
> If you need help to compile this code, refer to the  **Building the tutorials**  section for your platform:  [Linux](https://gstreamer.freedesktop.org/documentation/installing/on-linux.html#InstallingonLinux-Build),  [Mac OS X](https://gstreamer.freedesktop.org/documentation/installing/on-mac-osx.html#InstallingonMacOSX-Build)  or  [Windows](https://gstreamer.freedesktop.org/documentation/installing/on-windows.html#InstallingonWindows-Build), or use this specific command on Linux:
> 
> ``gcc basic-tutorial-8.c -o basic-tutorial-8 `pkg-config --cflags --libs gstreamer-1.0 gstreamer-audio-1.0` ``
> 
> If you need help to run this code, refer to the  **Running the tutorials**  section for your platform:  [Linux](https://gstreamer.freedesktop.org/documentation/installing/on-linux.html#InstallingonLinux-Run),  [Mac OS X](https://gstreamer.freedesktop.org/documentation/installing/on-mac-osx.html#InstallingonMacOSX-Run)  or  [Windows](https://gstreamer.freedesktop.org/documentation/installing/on-windows.html#InstallingonWindows-Run).
> 
> This tutorial plays an audible tone for varying frequency through the audio card and opens a window with a waveform representation of the tone. The waveform should be a sinusoid, but due to the refreshing of the window might not appear so.
> 
> Required libraries:  `gstreamer-1.0`

## Walkthrough[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/short-cutting-the-pipeline.html#walkthrough)

The code to create the pipeline (Lines 131 to 205) is an enlarged version of  [Basic tutorial 7: Multithreading and Pad Availability](https://gstreamer.freedesktop.org/documentation/tutorials/basic/multithreading-and-pad-availability.html). It involves instantiating all the elements, link the elements with Always Pads, and manually link the Request Pads of the  `tee`  element.

Regarding the configuration of the  `appsrc`  and  `appsink`  elements:

```c
/* Configure appsrc */
gst_audio_info_set_format (&info, GST_AUDIO_FORMAT_S16, SAMPLE_RATE, 1, NULL);
audio_caps = gst_audio_info_to_caps (&info);
g_object_set (data.app_source, "caps", audio_caps, NULL);
g_signal_connect (data.app_source, "need-data", G_CALLBACK (start_feed), &data);
g_signal_connect (data.app_source, "enough-data", G_CALLBACK (stop_feed), &data);

```

The first property that needs to be set on the  `appsrc`  is  `caps`. It specifies the kind of data that the element is going to produce, so GStreamer can check if linking with downstream elements is possible (this is, if the downstream elements will understand this kind of data). This property must be a  `GstCaps`  object, which is easily built from a string with  `gst_caps_from_string()`.

We then connect to the  `need-data`  and  `enough-data`  signals. These are fired by  `appsrc`  when its internal queue of data is running low or almost full, respectively. We will use these signals to start and stop (respectively) our signal generation process.

```c
/* Configure appsink */
g_object_set (data.app_sink, "emit-signals", TRUE, "caps", audio_caps, NULL);
g_signal_connect (data.app_sink, "new-sample", G_CALLBACK (new_sample), &data);
gst_caps_unref (audio_caps);

```

Regarding the  `appsink`  configuration, we connect to the  `new-sample`  signal, which is emitted every time the sink receives a buffer. Also, the signal emission needs to be enabled through the  `emit-signals`  property, because, by default, it is disabled.

Starting the pipeline, waiting for messages and final cleanup is done as usual. Let's review the callbacks we have just registered:

```c
/* This signal callback triggers when appsrc needs data. Here, we add an idle handler
 * to the mainloop to start pushing data into the appsrc */
static void start_feed (GstElement *source, guint size, CustomData *data) {
  if (data->sourceid == 0) {
    g_print ("Start feeding\n");
    data->sourceid = g_idle_add ((GSourceFunc) push_data, data);
  }
}

```

This function is called when the internal queue of  `appsrc`  is about to starve (run out of data). The only thing we do here is register a GLib idle function with  `g_idle_add()`  that feeds data to  `appsrc`  until it is full again. A GLib idle function is a method that GLib will call from its main loop whenever it is “idle”, this is, when it has no higher-priority tasks to perform. It requires a GLib  `GMainLoop`  to be instantiated and running, obviously.

This is only one of the multiple approaches that  `appsrc`  allows. In particular, buffers do not need to be fed into  `appsrc`  from the main thread using GLib, and you do not need to use the  `need-data`  and  `enough-data`  signals to synchronize with  `appsrc`  (although this is allegedly the most convenient).

We take note of the sourceid that  `g_idle_add()`  returns, so we can disable it later.

```c
/* This callback triggers when appsrc has enough data and we can stop sending.
 * We remove the idle handler from the mainloop */
static void stop_feed (GstElement *source, CustomData *data) {
  if (data->sourceid != 0) {
    g_print ("Stop feeding\n");
    g_source_remove (data->sourceid);
    data->sourceid = 0;
  }
}

```

This function is called when the internal queue of  `appsrc`  is full enough so we stop pushing data. Here we simply remove the idle function by using  `g_source_remove()`  (The idle function is implemented as a  `GSource`).

```c
/* This method is called by the idle GSource in the mainloop, to feed CHUNK_SIZE bytes into appsrc.
 * The ide handler is added to the mainloop when appsrc requests us to start sending data (need-data signal)
 * and is removed when appsrc has enough data (enough-data signal).
 */
static gboolean push_data (CustomData *data) {
  GstBuffer *buffer;
  GstFlowReturn ret;
  int i;
  gint16 *raw;
  gint num_samples = CHUNK_SIZE / 2; /* Because each sample is 16 bits */
  gfloat freq;

  /* Create a new empty buffer */
  buffer = gst_buffer_new_and_alloc (CHUNK_SIZE);

  /* Set its timestamp and duration */
  GST_BUFFER_TIMESTAMP (buffer) = gst_util_uint64_scale (data->num_samples, GST_SECOND, SAMPLE_RATE);
  GST_BUFFER_DURATION (buffer) = gst_util_uint64_scale (num_samples, GST_SECOND, SAMPLE_RATE);

  /* Generate some psychodelic waveforms */
  raw = (gint16 *)GST_BUFFER_DATA (buffer);

```

This is the function that feeds  `appsrc`. It will be called by GLib at times and rates which are out of our control, but we know that we will disable it when its job is done (when the queue in  `appsrc`  is full).

Its first task is to create a new buffer with a given size (in this example, it is arbitrarily set to 1024 bytes) with  `gst_buffer_new_and_alloc()`.

We count the number of samples that we have generated so far with the  `CustomData.num_samples`  variable, so we can time-stamp this buffer using the  `GST_BUFFER_TIMESTAMP`  macro in  `GstBuffer`.

Since we are producing buffers of the same size, their duration is the same and is set using the  `GST_BUFFER_DURATION`  in  `GstBuffer`.

`gst_util_uint64_scale()`  is a utility function that scales (multiply and divide) numbers which can be large, without fear of overflows.

The bytes that for the buffer can be accessed with GST_BUFFER_DATA in  `GstBuffer`  (Be careful not to write past the end of the buffer: you allocated it, so you know its size).

We will skip over the waveform generation, since it is outside the scope of this tutorial (it is simply a funny way of generating a pretty psychedelic wave).

```c
/* Push the buffer into the appsrc */
g_signal_emit_by_name (data->app_source, "push-buffer", buffer, &ret);

/* Free the buffer now that we are done with it */
gst_buffer_unref (buffer);

```

Once we have the buffer ready, we pass it to  `appsrc`  with the  `push-buffer`  action signal (see information box at the end of  [Playback tutorial 1: Playbin usage](https://gstreamer.freedesktop.org/documentation/tutorials/playback/playbin-usage.html)), and then  `gst_buffer_unref()`  it since we no longer need it.

```c
/* The appsink has received a buffer */
static GstFlowReturn new_sample (GstElement *sink, CustomData *data) {
  GstSample *sample;
  /* Retrieve the buffer */
  g_signal_emit_by_name (sink, "pull-sample", &sample);
  if (sample) {
    /* The only thing we do in this example is print a * to indicate a received buffer */
    g_print ("*");
    gst_sample_unref (sample);
    return GST_FLOW_OK;
  }
  return GST_FLOW_ERROR;
}

```

Finally, this is the function that gets called when the  `appsink`  receives a buffer. We use the  `pull-sample`  action signal to retrieve the buffer and then just print some indicator on the screen. We can retrieve the data pointer using the  `GST_BUFFER_DATA`  macro and the data size using the  `GST_BUFFER_SIZE`  macro in  `GstBuffer`. Remember that this buffer does not have to match the buffer that we produced in the  `push_data`  function, any element in the path could have altered the buffers in any way (Not in this example: there is only a  `tee`  in the path between  `appsrc`  and  `appsink`, and it does not change the content of the buffers).

We then  `gst_buffer_unref()`  the buffer, and this tutorial is done.

## Conclusion[](https://gstreamer.freedesktop.org/documentation/tutorials/basic/short-cutting-the-pipeline.html#conclusion)

This tutorial has shown how applications can:

-   Inject data into a pipeline using the  `appsrc`element.
-   Retrieve data from a pipeline using the  `appsink`  element.
-   Manipulate this data by accessing the  `GstBuffer`.

In a playbin-based pipeline, the same goals are achieved in a slightly different way.  [Playback tutorial 3: Short-cutting the pipeline](https://gstreamer.freedesktop.org/documentation/tutorials/playback/short-cutting-the-pipeline.html)  shows how to do it.

It has been a pleasure having you here, and see you soon!



[^GStreamDoc]:[Short-cutting the pipeline](https://gstreamer.freedesktop.org/documentation/tutorials/basic/short-cutting-the-pipeline.html#basic-tutorial-8-shortcutting-the-pipeline)

