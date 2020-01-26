---
layout: post
title:  "[译]FFmpeg: The ultimate Video and Audio Manipulation Tool"
date:   2020-01-01 12:00:00 +0800
tags:   tech
---


# [FFmpeg: The ultimate Video and Audio Manipulation Tool](http://blog.superuser.com/2012/02/24/ffmpeg-the-ultimate-video-and-audio-manipulation-tool/)

February 24, 2012  by  [slhck](http://blog.superuser.com/author/slhck/ "Posts by slhck").  [17 comments](http://blog.superuser.com/2012/02/24/ffmpeg-the-ultimate-video-and-audio-manipulation-tool/#comments)

# What is FFmpeg?

Chances are you’ve probably heard of  [FFmpeg](http://ffmpeg.org/)  already. It’s a set of tools dedicated to decoding, encoding and transcoding video and audio. FFmpeg is based on the popular  `libavcodec`  and  `libavformat`  libraries that can be found in many other video conversion applications, like  [Handbrake](http://handbrake.fr/).

So why would you need FFmpeg? Got a video in an obscure format that every other player couldn’t recognize? Transcode it with FFmpeg. Want to automate cutting video segments out of movies? Write a short batch script that uses FFmpeg. Where I work, I constantly have to encode and process video material, and I’d never want to go back to using a GUI tool for doing this.

This post is a follow-up on  [Video Conversion done right: Codecs and Software](http://blog.superuser.com/2011/11/07/video-conversion-done-right-codecs-and-software/), where I discussed the various codecs and containers that you can find these days. For a quick overview, I’d suggest to read this one as well, because it covers some important basics.

Now, let’s dive into the more practical aspects.

# Installation

FFmpeg is free and open source, and it’s cross-platform. You’ll be able to install it on your Windows PC as well as your Linux server. The only thing you need to be comfortable with is using your command line. As it’s actively developed, you should try to update your version from time to time. Sometimes, new features are added, and bugs are fixed. You won’t believe how many times updating FFmpeg solved encoding problems I had.

## Windows

Compiling FFmpeg on Windows is not too easy. That’s why there are (semi-) automated builds of the latest version available online, most prominently the ones  [from Zeranoe.com](http://ffmpeg.zeranoe.com/builds/). Download the latest build labeled as “static” for your 32 or 64 Bit system. Those builds work fairly well, but they might not include all possible codecs. For the commands where I used  `libfaac`  below, with the Zeranoe builds you have to choose another AAC codec instead, like  `libvo_aacenc`, or  `aac`, with the additional options  `-strict experimental`.

The Zeranoe download will include  `ffmpeg.exe`, which is the main tool we are going to use. And you’re done. That was easy, was it?

## OS X

On OS X, you have the tools for building programs from source. It’s very easy to install FFmpeg if you use a package manager like  [Homebrew](http://www.ffmpegx.com/)  or  [MacPorts](http://www.ffmpegx.com/). Install one of them if you haven’t already, and then run either

```
brew install ffmpeg --with-ffplay 
```

or

```
sudo port install ffmpeg-devel 
```

This takes a while, because dependencies have to be built too. When you’re done, you should be able to call FFmpeg by running  `ffmpeg`  from the command line. You should see a copyright header and some build details.

If you’re lazy, you can also  [download a static build](http://www.evermeet.cx/ffmpeg/), which just works out of the box. You only need to extract the archive and you’ll find an  `ffmpeg`  binary.

Note that there’s also  [ffmpegX](http://www.ffmpegx.com/), but its bundled version always is a bit behind, and I wouldn’t recommend using it unless you have a good reason for doing so. The  `.app`  package contains a compiled version of the  `ffmpeg`  command line tool already, which you  _could_  theoretically use.

## Linux

On many Linux distributions, FFmpeg comes in the default packages (e.g. on Ubuntu, where  [Libav](http://libav.org/)  – an FFmpeg fork – is bundled). However, mostly this is a very old release version. Do yourself a favor and compile from source – it sounds harder than it actually is. FFmpeg has a few dependencies, so therefore you’ll want to install encoders such as LAME for MP3 and x264 for videos as well. To compile FFmpeg, look at the FFmpeg wiki, which has  [a few compilation guides](https://ffmpeg.org/trac/ffmpeg/wiki/CompilationGuide)  for various Linux distributions.

Similar to Windows, static builds for FFmpeg are  [also available](http://ffmpeg.gusari.org/static/). You can download these if you don’t want to compile yourself or install the outdate version from your package manager.

----------

# Our first encoded Video

Now that we have FFmpeg working, we can start with our video or audio conversion. Note that this is always a very resource intensive task, especially with video. Expect some conversions to take a little while, for example when re-encoding whole movies. A quad-core CPU like the i7 can easily be saturated with a good encoder and performs very fast then.

Before you handle videos, you should know what the difference between “codec” and “container” (also, “format”) is, as I will be using these terms without explaining them any further. I know these can be easily mixed up and aren’t as self-explaining as you’d like them to be. To read more about this, see this Super User answer:  [What is a Codec (e.g. DivX?), and how does it differ from a File Format (e.g. MPG)?](http://superuser.com/a/300997/48078)

## Basic Syntax

Okay, now let’s start. The most basic form of an FFmpeg command goes like this – note the absence of an output option like  `-o`:

```
ffmpeg -i input output 
```

And you’re done. For example, you could convert an MP3 into a WAV file, or an MKV video into MP4 (but don’t do that just yet, please). You could also just extract audio from a video file.

```
ffmpeg -i audio.mp3 audio.wav
```

FFmpeg will guess which codecs you want to use depending on the format specifier (e.g. “.wav” container obviously needs a WAV codec inside, and MKV video will use x264 video and AC3 audio). But that’s not always what we want. Let’s say you have an MP4 container – which video codec should it include? There are a couple of possible answers. So, we’re better off specifying these codecs ourselves, since some codecs are better than others or have particularities that you might need.

## Specifying Video and Audio codecs

In most cases, you want a specific output codec, like h.264, or AAC audio. Today, there’s almost no reason  _not_ to encode to h.264, as it offers incredible quality at small file sizes. But which codecs do really work in FFmpeg? Luckily, it has a ton of built-in codecs and formats, which you can get a list of. Check the output: The  `D`  and  `E`  stand for “decoding” and/or “encoding” capabilities, respectively.

```
ffmpeg -codecs
```

Now, we can use any of those to perform our transcoding. Use the  `-c`  option to specify a codec, with a  `:a`  or  `:v`  for audio and video codecs respectively. These are called “streams” in FFmpeg – normally there is one video and one audio stream, but you could of course use more. Let’s start by encoding an audio file with just one audio stream, and then a video file with a video stream and an audio stream.

```
ffmpeg -i input.wav -c:a libfaac output.mp4
```

The possible combinations are countless, but obviously restricted by the format/container. Especially those pesky old  [AVI](http://en.wikipedia.org/wiki/Audio_Video_Interleave)  containers don’t like all codecs, so you’re better off encoding to MP4 or MKV containers these days, because they accept more codecs and can be easily handled by any modern player.

## When to copy, when to encode?

You might not have thought about this before, but sometimes, you want to just copy the contents of a video and not re-encode. This is actually very critical when just cutting out portions from a file, or only changing containers (and  _not_  the codecs inside). In the example I gave above, we wanted to change MP4 to MKV.

The following command is  **wrong**: It will re-encode your video. It will take forever, and  [the result will (probably) look bad](http://en.wikipedia.org/wiki/Generation_loss).

```
ffmpeg -i input.mp4 output.mkv
```

What FFmpeg did here was encoding the file yet another time, because you didn’t tell it to copy. This command however does it right:

```
ffmpeg -i input.mp4 -c:v copy -c:a copy output.mkv 
```

If you were to re-encode your videos all the time when you cut them, instead of just copying the codec contents, you’d eventually end up with something like this (exaggerated, but still highly entertaining)  [example of generation loss](http://www.youtube.com/watch?v=icruGcSsPp0). Looks scary, does it?

If changing the container without encoding does not work for you, you can try the  `mkvmerge`  tool that  [mkvtoolnix](http://www.bunkus.org/videotools/mkvtoolnix/)  offers. This will for example create an MKV file.

```
mkvmerge input.avi -o output.mkv
```

Similarly,  [MP4Box](http://gpac.wp.mines-telecom.fr/)  can create MP4 files from existing video.

----------

# Advanced Options

If you’ve used some test videos or audio files for the above sequences, you might have seen that the quality wasn’t good enough to be useful. This is completely normal. Also, you will probably have constraints about bit rate, file size or target devices, such as when encoding a video for the PlayStation 3 or your Android phone.

## Quality Settings

Quality comes first. What is “quality”, even? Generally, the more bits you can spend on a file, the better it will look or sound. This means, that for the same codec, larger file size (or larger bit rate) will equal in better quality. In most cases, that is.

Most people associate quality with “bit rate”. This is a relict from the good old days where everyone would encode MP3 files from their CDs. Of course, this is a relatively simple approach, and you can get a good feeling for the bit rates you might need for certain videos. If you want to set the bit rate, you can do that with the  `-b`  option, again specifying the audio or video stream (just to resolve any ambiguities). The bit rate can be specified with suffixes like “K” for kBit/s or “M” for mBit/s.

```
ffmpeg -i input.wav -b:a 192K out.mp3
```

But often, bit rate is not enough. You only need to restrict bit rate when you have a very specific target file size to reach. In all other cases, use a better concept, called “constant quality”. The reason is that sometimes you don’t want to spend the same amount of bits to a segment of a file. That’s when you should use variable bit rate. Actually, this concept is very well known for MP3 audio, where VBR rips are commonly found.

In video, this means setting a certain value called  [Constant Rate Factor](http://homepage.univie.ac.at/werner.robitza/crf.html). For x264 (the h.264 encoder you should use), this is very easy:

```
ffmpeg -i source.mp4 -c:v libx264 -crf 23 out.mp4 
```

The CRF can be anything within 0 and 51, with the reasonable range being 17 to 26. The lower, the better the quality, the higher the file size. The default value is 23 here. For MPEG-4 video (like XviD), a similar concept exists, called “qscale”.

```
ffmpeg -i source.mp4 -c:v mpeg4 -qscale:v 3 out.mp4 
```

Here, the  `qscale:v`  can range from 1 to 31. The lower, the higher the quality, with values of 3 to 5 giving a good enough result in most cases.

In general, the best bet is to just  _try and see for yourself_ what looks good. Take into account the result file size you want and how much quality you can trade in for smaller file sizes. It’s all up to you. As a last note, don’t ever fall for the  `sameq`  option.  [Sameq does not mean same quality](http://superuser.com/a/478550/48078)  and you should never use it.

## Cutting Video

Often, you want to just cut out a portion from a file. FFmpeg supports basic cutting with the  `-t`  and  `-ss`  options. The first one will specify the duration of the output, and the second one will specify the start point. For example, to get the first five seconds, starting from one minute and 30 seconds:

```
ffmpeg -ss 00:01:30 -i input.mov -c:v copy -c:a copy -t 5 output.mov 
```

The time values can either be seconds or in the form of  `HH:MM:SS.ms`. So, you could also cut one minute, ten seconds and 500 milliseconds:

```
ffmpeg -ss 00:01:30 -i input.mov -c:v copy -c:a copy -t 00:01:30.500 output.mov 
```

Note that we’ve again just copied the contents instead of re-encoding because we used the  `copy`  codec. Also, see how the  `-ss`  option is  _before_  the actual input? This will first seek to the point in the file, and  _then_  start to encode and cut. This is especially useful when you’re actually not just copying content, but encoding it, because it’ll speed up your conversion (although it’s less accurate). If you want to be more accurate, you’ll have to tell FFmpeg to encode the whole video and just discard the output until the position indicated by  `-ss`:

```
ffmpeg -i input.mov -ss 00:01:30 -c:v copy -c:a copy -t 5 output.mov
```

## Resizing

Got a High Definition 1080p video that you want to watch on your mobile phone? Scale it down first! FFmpeg supports software-based scaling when encoding with the  `-s`  option. The following example will do that:

```
ffmpeg -i hd-movie.mkv -c:v libx264 -s:v 854x480 -c:a copy out.mp4 
```

Phones have some restrictions on video playback, which can be tackled by using so-called “profiles” when encoding. More about this can be found in the resources at the bottom.

## Other useful Examples

On Super User, I’m currently very active in the  [FFmpeg tag](http://superuser.com/questions/tagged/ffmpeg). You could browse the questions there, or see this list of questions and answers I’ve compiled.

General issues:

-   [Resources To Use FFMPEG Effectively](http://superuser.com/questions/373018/resources-to-use-ffmpeg-effectively/373024#373024), including some general links and resources
-   [What do the numbers 240 and 360 mean when downloading video? How can I tell which video is more compressed?](http://superuser.com/questions/359730/what-do-the-numbers-240-and-360-mean-when-downloading-video-how-can-i-tell-whic/359734#359734), which focuses on tradeoffs between bit rate, file size, video dimensions, et cetera.

Specific conversion problems:

-   [Convert from MOV to MP4 Container format](http://superuser.com/questions/378726/convert-from-mov-to-mp4-container-format/378761#378761), which talks about switching containers only
-   [How to convert any video to DVD (.vob) in good video quality?](http://superuser.com/questions/299550/how-to-convert-any-video-to-dvd-vob-in-good-video-quality/299557#299557), about the limitations of retaining quality when converting
-   [What parameters should I be looking at to reduce the size of a .MOV file?](http://superuser.com/questions/383903/what-parameters-should-i-be-looking-at-to-reduce-the-size-of-a-mov-file/383946#383946), in depth explanation about how to use the Constant Rate Factor
-   [Use DivX settings to encode to mp4 with ffmpeg](http://superuser.com/questions/380633/use-divx-settings-to-encode-to-mp4-with-ffmpeg), more examples about CRF and how to set advanced x264 options like presets and profiles
-   [Convert any video to HD Video](http://superuser.com/questions/290136/convert-any-video-to-hd-video/290155#290155), about up- or downscaling video
-   [How to cut at exact frames using ffmpeg?](http://superuser.com/questions/459313/how-to-cut-at-exact-frames-using-ffmpeg/459488#459488), which shows how to calculate cutting points for video

Of course, if you have any specific problem, feel free to ask it on the main site and we Super Users can help you out! However, make sure you always post the command you’re using and the full, uncut output FFmpeg gives you. Also, always try to use the latest version if possible. Ideally, we’d like to see a sample video as well.


## Other by tiga



### 推流命令
```shell
./ffmpeg -re -i VID_20190806_151825.mp4 -vcodec copy -an -f rtp rtp://192.168.199.198:12345 -acodec copy -vn -f rtp rtp://192.168.199.198:12356 > t.sdp
./ffmpeg -re -i VID_20190806_151825.mp4 -vcodec copy -an -f rtp rtp://192.168.199.198:12346 > tv.sdp
./ffmpeg -re -i VID_20190806_151825.mp4 -acodec copy -vn -f rtp rtp://192.168.199.198:12345 > ta.sdp
./ffplay.exe  -protocol_whitelist "file,http,https,rtp,udp,tcp,tls" ./t.sdp
./ffmpeg -i VID_20190806_151825.mp4 -c:a libopus -b:a 96K t.opus

gst-launch-1.0 filesrc location="./VID_20190806_151825.mp4" ! decodebin ! x264enc tune=zerolatency subme=1 speed-preset=1 qp-min=30 key-int-max=15 ! rtph264pay ! udpsink host=192.168.199.1 port=12346 filesrc location="./VID_20190806_151825.mp4" ! decodebin ! audioconvert ! audioresample ! mulawenc  ! rtppcmupay max-ptime=20000000 min-ptime=20000000  ! udpsink host=192.168.199.1 port=12345 
```

我们音频只支持pcum opus，上行aac不支持

gstream 下载
[https://gstreamer.freedesktop.org/data/pkg/windows/1.16.0/](https://gstreamer.freedesktop.org/data/pkg/windows/1.16.0/)

ffmpeg 转码命令
[https://gist.github.com/protrolium/e0dbd4bb0f1a396fcb55](https://gist.github.com/protrolium/e0dbd4bb0f1a396fcb55)

#### ffplay无法播放声音
在mingw中运行 ffplay.exe 时会提示以下错误，原因参考这是 [# FFplay: WASAPI can't initialize audio client](https://trac.ffmpeg.org/ticket/6891)
1.使用原生的CMD运行ffplay.exe
2.set SDL_AUDIODRIVER=directsound && ./ffplay.exe Kalimba.mp3 
```shell
dodo@dodo-PC MINGW64 /d/Program Files/ffmpeg-4.1.4-win64-static/bin
$ set SDL_AUDIODRIVER=directsound && ./ffplay.exe Kalimba.mp3
ffplay version 4.1.4 Copyright (c) 2003-2019 the FFmpeg developers
  built with gcc 9.1.1 (GCC) 20190716
  configuration: --enable-gpl --enable-version3 --enable-sdl2 --enable-fontconfig --enable-gnutls --enable-iconv --enable-libass --enable-libbluray --enable-libfreetype --enable-libmp3lame --enable-libopencore-amrnb --enable-libopencore-amrwb --enable-libopenjpeg --enable-libopus --enable-libshine --enable-libsnappy --enable-libsoxr --enable-libtheora --enable-libtwolame --enable-libvpx --enable-libwavpack --enable-libwebp --enable-libx264 --enable-libx265 --enable-libxml2 --enable-libzimg --enable-lzma --enable-zlib --enable-gmp --enable-libvidstab --enable-libvorbis --enable-libvo-amrwbenc --enable-libmysofa --enable-libspeex --enable-libxvid --enable-libaom --enable-libmfx --enable-amf --enable-ffnvcodec --enable-cuvid --enable-d3d11va --enable-nvenc --enable-nvdec --enable-dxva2 --enable-avisynth
  libavutil      56. 22.100 / 56. 22.100
  libavcodec     58. 35.100 / 58. 35.100
  libavformat    58. 20.100 / 58. 20.100
  libavdevice    58.  5.100 / 58.  5.100
  libavfilter     7. 40.101 /  7. 40.101
  libswscale      5.  3.100 /  5.  3.100
  libswresample   3.  3.100 /  3.  3.100
  libpostproc    55.  3.100 / 55.  3.100
[mp3 @ 0000000000534d80] Skipping 386 bytes of junk at 60603.
[mp3 @ 0000000000534d80] Estimating duration from bitrate, this may be inaccurate
Input #0, mp3, from 'Kalimba.mp3':
  Metadata:
    comment         : Ninja Tune Records
    publisher       : Ninja Tune
    track           : 1
    album           : Ninja Tuna
    artist          : Mr. Scruff
    album_artist    : Mr. Scruff
    title           : Kalimba
    genre           : Electronic
    composer        : A. Carthy and A. Kingslow
    date            : 2008
    id3v2_priv.WM/UniqueFileIdentifier: A\x00M\x00G\x00a\x00_\x00i\x00d\x00=\x00R\x00 \x00 \x001\x004\x004\x007\x008\x007\x005\x00;\x00A\x00M\x00G\x00p\x00_\x00i\x00d\x00=\x00P\x00 \x00 \x00 \x002\x002\x004\x001\x005\x007\x00;\x00A\x00M\x00G\x00t\x00_\x00i\x00d\x00=\x00T\x00 \x001\x005\x000\x00
    id3v2_priv.WM/WMCollectionGroupID: 5]\xa0_\x82\xa6\xf6J\x96\xf7\x07s\xe4-M\x16
    id3v2_priv.WM/WMCollectionID: 5]\xa0_\x82\xa6\xf6J\x96\xf7\x07s\xe4-M\x16
    id3v2_priv.WM/WMContentID: \xf3\xa0\x0fO\x95=\x1aG\xb0\xd2\x9d\xcb0\xa9\xbb\xae
    id3v2_priv.WM/Provider: A\x00M\x00G\x00\x00\x00
    id3v2_priv.WM/MediaClassPrimaryID: \xbc}`\xd1#\xe3\xe2K\x86\xa1H\xa4*(D\x1e
    id3v2_priv.WM/MediaClassSecondaryID: \x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00
  Duration: 00:05:48.06, start: 0.000000, bitrate: 193 kb/s
    Stream #0:0: Audio: mp3, 44100 Hz, stereo, fltp, 192 kb/s
    Stream #0:1: Video: mjpeg, yuvj420p(pc, bt470bg/unknown/unknown), 512x512, 90k tbr, 90k tbn, 90k tbc
    Metadata:
      title           : thumbnail
      comment         : Cover (front)
SDL_OpenAudio (2 channels, 44100 Hz): WASAPI can't initialize audio client: 尚未调用 CoInitialize。

SDL_OpenAudio (1 channels, 44100 Hz): WASAPI can't initialize audio client: 尚未调用 CoInitialize。

No more combinations to try, audio open failed
[swscaler @ 0000000006b52cc0] deprecated pixel format used, make sure you did set range correctly
    nan M-V:    nan fd=   0 aq=    0KB vq=    0KB sq=    0B f=0/0
```


```shell
# convert mp4 to flv
# https://stackoverflow.com/questions/8504923/how-to-convert-mp4-video-file-into-flv-format-using-ffmpeg
ffmpeg -i source.mp4 -c:v libx264 -crf 19 destinationfile.flv
ffmpeg -i source.mp4 -c:v libx264 -ar 22050 -crf 28 destinationfile.flv

# push mp4 file to rtmp server
ffmpeg -re -i ./big_buck_bunny.mp4 -c copy -f flv rtmp://localhost:1935/live/live

# pull rtmp media stream
ffplay "rtmp://localhost:1935/live/live"


# mp4 to rtmp (push stream)
# https://stackoverflow.com/questions/29794053/streaming-mp4-video-file-on-gstreamer
# https://raspberrypi.stackexchange.com/questions/40110/how-to-stream-video-file-to-rtmp-server-with-gstreamer-on-rpi2
gst-launch-1.0 \
  filesrc location="big_buck_bunny.mp4" ! decodebin name=t \
  t. ! videoconvert ! x264enc ! queue ! flvmux name=mux \
  t. ! audioconvert ! voaacenc bitrate=128000 ! queue ! mux. \
  mux. ! rtmpsink location="rtmp://127.0.0.1:1935/live/live live=1"

# play rtmp  (pull stream)
ffplay rtmp://127.0.0.1:1935/live/live

# play rtp
TODO

# mp4 to rtp
gst-launch-1.0 filesrc location=/Users/xiang/Movies/test2.mp4 ! qtdemux !  video/x-h264 ! h264parse ! rtph264pay config-interval=-1 ! udpsink  host=127.0.0.1 port=5000

# rtmp to rtp
gst-launch-1.0 -v rtmpsrc location=rtmp://localhost/live/stream  ! flvdemux name=demux demux.audio ! queue ! decodebin ! audioconvert ! audioresample ! opusenc ! rtpopuspay timestamp-offset=0  ! udpsink  host=127.0.0.1 port=5000 demux.video! queue ! h264parse ! rtph264pay timestamp-offset=0 config-interval=-1  ! udpsink  host=127.0.0.1 port=5002

# appto rtmp
appsrc is-live=true do-timestamp=true name=src ! h264parse ! video/x-h264,stream-format=(string)avc ! flvmux ! rtmpsink sync=false location=%s

```

#### decodebin

GstBin that auto-magically constructs a decoding pipeline using available decoders and demuxers via auto-plugging.

decodebin is considered stable now and replaces the old decodebin element. uridecodebin uses decodebin internally and is often more convenient to use, as it creates a suitable source element as well.

