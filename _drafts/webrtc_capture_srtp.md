
> 来源 
> [https://webrtchacks.com/video_replay/](https://webrtchacks.com/video_replay/)
> [webrtc-debug ](https://blog.codeship.com/webrtc-issues-and-how-to-debug-them/?__s=sgkgganpatrhthvch4js)

## webrtcH4cKS: ~ How to capture & replay WebRTC video streams with video_replay (Stian Selnes)

Posted by  [Chad Hart](https://webrtchacks.com/author/chad/ "View all posts by Chad Hart")  on  [August 31, 2017](https://webrtchacks.com/video_replay/)

**Posted in:**  [Guide](https://webrtchacks.com/category/guide/).

**Tagged:**  [debug](https://webrtchacks.com/tag/debug/),  [stian selnes](https://webrtchacks.com/tag/stian-selnes/),  [video](https://webrtchacks.com/tag/video/),  [video_replay](https://webrtchacks.com/tag/video_replay/),  [wireshark](https://webrtchacks.com/tag/wireshark/).

[10 comments](https://webrtchacks.com/video_replay/#comments)

Decoding video when there is packet loss is not an easy task. Recent Chrome versions have been plagued by video corruption issues related to a new video jitter buffer introduced in Chrome 58. These issues are hard to debug since they occur only when certain packets are lost. To combat these issues, webrtc.org has a pretty powerful tool to reproduce and analyze them called  _video_replay_. When I saw another video corruption issue filed by Stian Selnes I told him about that tool. With an easy reproduction of the stream, the WebRTC video team at Google made short work of the bug. Unfortunately this process is not too well documented, so we asked Stian to walk us through the process of capturing the necessary data and using the  _video_replay_tool. Stian, who works at  [Pexip](https://www.pexip.com/), has been dealing with real-time communication for more than 10 years. He has experience in large parts of the media stack with a special interest in video codecs and other types of signal processing, network protocols and error resilience.

{“intro-by”: “[Philipp Hancke](https://webrtchacks.com/about/fippo/)“}

![corrupted video, oh my!](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/08/corruption.png)

You have a problem if your WebRTC video stream looks like this. Fortunately there are tools to help like [video_replay](https://chromium.googlesource.com/external/webrtc/trunk/webrtc.git/+/master/video/replay.cc).

WebRTC contains a nice and little known tool called  [video_replay](https://chromium.googlesource.com/external/webrtc/trunk/webrtc.git/+/master/video/replay.cc) which has proven very useful for debugging video decoding issues. It’s purpose? To easily replay a capture of a WebRTC call to reproduce an observed behavior.  _video_replay_  takes a captured  [RTP](https://tools.ietf.org/html/rfc3550) stream of video as an input file, decodes the stream with the WebRTC framework “offline”, and then displays the resulting output on screen.

To illustrate, recently I was working on an  [issue](https://bugs.chromium.org/p/webrtc/issues/detail?id=8028) where Chrome suddenly displayed the incoming video as a corrupt image shown above. Eventually, after using  _video_replay_  to debug, the WebRTC team found that Chrome’s jitter buffer reimplementation introduced a bug that made the video stream corrupt in certain cases. This caused the internal state of the VP8 decoder to be wrong and output frames that looked like random data.

Video Player

00:00

00:15

Video coding issues are often among the toughest problems to solve. Initially I was able to develop a test setup that reproduced the issue in about 1 out of every 20 calls. Replicating the issue this way was very time consuming, often frustrating, and ultimately did not make it easy for the WebRTC team to find the fix. To consistently reproduce the problem I managed to get a  [Wireshark](https://www.wireshark.org/) capture of one of the failing calls and fed the capture into the  _video_replay_  tool. Voilà! Now I had a test case that reproduced this rare issue every single time. When an issue is as reproducible as this it’s usually very attractive for someone to grab and implement a fix quickly. A classic win-win situation.

In this post I’ll demonstrate how to use  _video_replay_  by going through an example where we capture the RTP traffic of a WebRTC call, identify and extract the received video stream, and finally feed it into video_replay to display the captured video on screen.

# **Capture unencrypted RTP**

video_replay takes the input file and feeds it into the RTP stack, depacketizer and decoder. However, it currently lacks the capability to decrypt  [S](https://tools.ietf.org/html/rfc3711)[RTP](https://tools.ietf.org/html/rfc3711) packets for encrypted calls. Encrypted calls is the norm for both Chrome and Firefox. Decrypting a WebRTC call is not a trivial process, in particular since  [DTLS](https://tools.ietf.org/html/rfc4347) is used to securely share the secret key that’s used for SRTP so the key is not easy to get. To get around this, it’s better to use Chromium or Chrome Canary since they have a flag that will disable SRTP encryption. Simply start the browser with the command-line flag  [`--disable-webrtc-encryption`](https://peter.sh/experiments/chromium-command-line-switches/#disable-webrtc-encryption) and you should see warning should be displayed at the top of your window that you’re using an unsupported command-line flag. Note that it’s necessary that both parties in the call run unencrypted; if not, the call will fail to connect.

First, start to capture packets using Wireshark. It’s important to start the capture before the call starts sending media so that you make sure the entire stream is recorded. If the beginning of the stream is missing from the capture, the video decoder will fail to decode it.

Second, open a tab and go to  chrome://webrtc-internals  (or Fippo’s new  [webrtc-externals](https://webrtchacks.com/webrtc-externals/)). Do this prior to the call in order to get all the information needed, specifically the negotiated  [SDP](https://tools.ietf.org/html/draft-nandakumar-rtcweb-sdp) (see the webrtcHacks  [SDP Anatomy](https://webrtchacks.com/sdp-anatomy/)  guide here for help interpreting this).

Finally, it’s time to make the call. We’ll use  [appr.tc](https://appr.tc/) as an example, but any call that’s using WebRTC will work. Open a second tab and go to  [https://appr.tc/?ipv6=false](https://appr.tc/?ipv6=false). In this particular example I disabled IPv6 because of a currently unresolved  [issue](https://bugs.chromium.org/p/webrtc/issues/detail?id=8075) with with  _video_replay_, but I expect that to be fixed very soon.

Now, join a video room. RTP will start flowing when a second participant joins the same room. It doesn’t matter who joins first, except that  _chrome://webrtc-internals_  will look slightly different. The screenshots below were taken when dialing into an existing room.

# **Gather information**

In order to successfully extract the RTP packets for the received video stream and successfully play them with  _video_replay_  we need to gather some details about the RTP stream. There are a few ways to go about this, but I’ll stick to what I believe will give the clearest instructions. What we need is:

-   Video codec
-   RTP SSRC
-   RTP payload types
-   IP address and port

### Gathering stats with webrtc-internals

First, expand the stats table for the received video stream, which will have a name likessrc_4075734755_recv . There may be be more than one, typically a second stream for audio, and possibly a couple that has the _send suffix which holds the equivalent statistics for the sent streams. The stats for the received video stream are easily identified with the  _recv suffix andmediaType=video . Make a note of the fields  ssrc ,  googCodecName and  transportId , which is 4075734755, VP9 and Channel-audio-1 respectively.

You may ask why the video stream has the same transportId as the audio channel? This means that  [BUNDLE](https://tools.ietf.org/html/draft-ietf-mmusic-sdp-bundle-negotiation) is in use which makes audio and video share the same channel. If BUNDLE is not negotiated and in use, audio and video will use separate channels.

![webrtc details recv ssrc](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/08/webrtc-details-recv-ssrc.png)

Next, we’ll look at the negotiated SDP to get the RTP payload types (PT). In addition to the PT for the video codec in use, we must also find the PT for  [RED](https://tools.ietf.org/html/rfc2198) which WebRTC uses to encapsulate the video packets. SDP describes the receive capabilities of a video client, so in order to find the payload types we receive we must look at the SDP our browser offers to the other participant. This is found by expanding the  setLocalDescription API call, find the section  m=video and the following rtpmap attributes that defines the PT for each supported codec. Since our video codec is VP9, we’ll note the fields  a=rtpmap:98  VP9/90000 and  a=rtpmap:102  red/9000 , which tells us that the payload types for VP9 and RED are 98 and 102, respectively.

![webrtc details setLocalDescription](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/08/webrtc-details-setLocalDescription.png)

If you’re looking for information regarding the sent stream instead of the received stream, you should look at what the other participant has signaled by expanding  setRemoteDescription instead. In practice, however, the payload types should be symmetrical, so whether you look atsetLocalDescription or  setRemoteDescription shouldn’t matter, but in the world of real-time video communication you never know…

In order to quickly identify the correct RTP stream in Wireshark it’s very useful to know the IP address and port that’s being used. Whether you look at the remote or local address doesn’t matter as long as the appropriate Wireshark filter is applied. For this example we’re going to use the local address, which is the destination of the packets we’re interested in since we want to extract the received stream.  chrome://webrtc-internals  contains connection stats in the sections that starts with  Conn-audio and  Conn-video . The active ones will be highlighted with bold font, and based on the  transportId in previous step we know wether to look at the audio or video channel. To find the local address in our example we need to expand Conn-audio-1-0 and note the field  googLocalAddress , which has the value  10.47.4.245:52740

![view of webrtc connection details in chrome://webrtc-internals](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/08/webrtc-details-connection.png)

webrtc connection details in chrome://webrtc-internals

### RTP marking in Wireshark

Now that we have collected all the information necessary to easily identify and extract the received video stream in our call. Wireshark will likely show the captured RTP packets simply as UDP packets. We want to tell Wireshark that these are RTP packets so that we can export them to rtpdump format.

First, apply a display filter on address and port, e.g.  ip.dst  ==  10.47.4.245 and  udp.dstport==  52740. Then, right click a packet, select Decode As, and choose RTP.

Second, list all RTP streams by selecting from the menu  _Telephony → RTP → RTP Streams_. The SSRC of our received video stream will be listed together with other streams. Select it and export it as rtpdump format. (_video_replay_  also supports the pcap format, but because of very limited support for various link-layers I generally recommend to use rtpdump.) Finally we have a file that contains only the received video packets which can be fed into  _video_replay_.

![wireshark decode as](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/08/wireshark-decode-as.png)

![wireshark export rtpdump](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/08/wireshark-export-rtpdump.png)

# **Build WebRTC and video_replay**

You need to build  _video_replay_  from the WebRTC source before you can use it. General instructions on how to set up the environment, get the code and compile are found at  [https://webrtc.org/native-code/development](https://webrtc.org/native-code/development). Note that in order to build  _video_replay_  the tool has to be explicitly listed as a target when compiling. In short, after making sure that the  [prerequisite](https://webrtc.org/native-code/development/prerequisite-sw) software is installed, the following commands will get the code and build  _video_replay_.

```shell
mkdir  webrtc-checkout
cd  webrtc-checkout/
fetch  --nohooks webrtc
gclient sync
cd  src
gn gen out/Default
ninja  -C  out/Default  video_replay
```

# **Use video_replay to playback the capture**

Finally we’re ready to replay our captured stream and hopefully display it exactly how it appeared at  [appr.tc](https://appr.tc/)  originally. The minimal command line to do this for our example is:

```Shell
out/Default/video_replay  -input_file received-video.rtpdump  -codec VP9  -media_payload_type  98  -red_payload_type  102  -ssrc  4075734755
```

(note: the media_payload_type argument was named payload previously)
The command line arguments are easy to understood from the previous text.


### video_replay arguments

If your goal is to reproduce a WebRTC issue and  [post a bug](https://webrtc.org/bugs/), providing the rtpdump together with the command line arguments to replay it will be tremendously helpful for certain issues. However, if you want to do some more debugging on your own  _video_replay_  has a a few more command line options that may be of interest.

Let’s look at the current help text and explain what the different options do.

Flags from  ../../webrtc/video/replay.cc  at the time of this writing:

```shell
-abs_send_time_id  (RTP extension ID for  abs-send-time)  type:  int32 default:  -1

-codec  (Video codec)  type:  string  default:  "VP8"

-decoder_bitstream_filename  (Decoder bitstream output file)  type:  string  default:  ""

-fec_payload_type  (ULPFEC payload type)  type:  int32 default:  -1

-input_file  (input file)  type:  string  default:  ""

-out_base  (Basename  (excluding  .yuv)  for  raw output)  type:  string  default:  ""

-payload_type  (Payload type)  type:  int32 default:  123

-payload_type_rtx  (RTX payload type)  type:  int32 default:  98

-red_payload_type  (RED payload type)  type:  int32 default:  -1

-ssrc  (Incoming SSRC)  type:  uint64 default:  12648429

-ssrc_rtx  (Incoming RTX SSRC)  type:  uint64 default:  195939069

-transmission_offset_id  (RTP extension ID for  transmission-offset)  type:  int32 default:  -1

And here is what they mean with some more explanation:

```

Flag

Description

Required?

**abs_send_time**

The RTP extension ID for the  [abs-send-time](https://webrtc.org/experiments/rtp-hdrext/abs-send-time/) extension

Optional

**codec**

The video codec associated with payload_type and defines the codec used to decode the stream

Required

**decoder_bitstream_filename**

File to save the received bitstream after depayloading and before decoding. The bitstream will be dumped to file without any type of container or frame delimiters and cannot be decoded as far as I know. It may, however, be useful for visual inspection  
of the bits

Optional

**fec_payload_type**

The RTP payload type for  [Forward Error Correction](https://tools.ietf.org/html/draft-ietf-rtcweb-fec) packets. FEC is a mechanism to improve quality when there is packet loss. The payload type for  
FEC can be parsed from the SDP, e.ga=rtpmap:127  ulpfec/90000

Optional

**input_file**

The file to replay. This can either be rtpdump or pcap. Currently only  [LINKTYPE_NULL and LINKTYPE_ETHERNET](http://www.tcpdump.org/linktypes.html) is supported for pcaps, which makes its use case a  
bit limited

Required

**out_base**

Save the decoded frames to file in uncompressed format (I420)

Optional

**payload_type**

The payload type of the video packets to be decoded, as explained earlier

Required

**payload_type_rtx**

The payload type of  [RTP retransmission (RTX)](https://tools.ietf.org/html/rfc4588) packets in the case of packet loss. The payload type of such packets is also found in the SDP. Note that there are  
multiple payload types for RTX. Make sure to use the one that is associated with the correct payload_type, e.g.  a=rtpmap:99  rtx/90000a=fmtp:99  apt=98  says that payload type 99 is retransmission  
of packets with payload type 98

Optional

**red_payload_type**

The payload type of RED, as explained earlier

Optional

**ssrc**

The SSRC of the stream to decode, as explained earlier

Required

**ssrc_rtx**

The SSRC of the RTX stream. This value is actually found in the SDP offered by the sender side, as opposed to the receiver side as everything else we’ve looked at. It’s found in the attribute ssrc-group, e.g.  a=ssrc-group:FID  40757347553957275412,  
where  3957275412 is the SSRC for RTX

Optional

**transmission_offset_id**

The RTP extension ID for the  [transmission offset](https://tools.ietf.org/html/rfc5450) extension

Optional

# **Shortcuts**

When you get familiar with process above there are a couple of shortcuts you can apply in order to be more effective.

First, you can often identify the RTP video packets in Wireshark without looking at  chrome://webrtc-internals. Most video packets are usually more than 1000 bytes, while audio packets are more like a couple of hundred bytes. Decode the video packet as RTP and Wireshark will display both the SSRC and payload type. Wireshark doesn’t automatically know is whether RED is used or not, but from experience you can probably guess that too since payload types normally don’t change between calls.

Second, if your pcap is supported by  _video_replay_  you can feed the original pcap directly into video_replay. There will be a lot of noise on the command line output as it ignores all the unknown packets, but it will decode and display the specified stream.

{“author”, “[stian selnes](http://stianselnes.com/)“}

Want to keep up on our latest posts? Please click [here](http://webrtchacks.com/subscribe/) to subscribe to our mailing list if you have not already. We only email post updates. You can also follow us on twitter at [@webrtcHacks](http://twitter.com/webrtcHacks) for blog updates.

### Related Posts

-   [![Debugging VP8 is more fun than it used to be](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/03/27180224005_7a2067dce3_b-150x150.jpg)](https://webrtchacks.com/wireshark-debug-vp8/)[Debugging VP8 is more fun than it used to be](https://webrtchacks.com/wireshark-debug-vp8/)
-   [![WebRTC Video Codec Debate:  Is There No End in Sight? (Chris Wendt)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2014/01/ChrisWendt-150x150.jpg)](https://webrtchacks.com/video-codec-debate-commentary-chris-wednt/)[WebRTC Video Codec Debate: Is There No End in Sight? (Chris Wendt)](https://webrtchacks.com/video-codec-debate-commentary-chris-wednt/)
-   [![WebRTC mandatory video codec discussion: the final duel?](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2013/10/20071115180910Yevgeny_Onegin_by_Repin-150x150.jpg)](https://webrtchacks.com/webrtc-video-codec-discussion/)[WebRTC mandatory video codec discussion: the final duel?](https://webrtchacks.com/webrtc-video-codec-discussion/)
-   [![Bisecting Browser Bugs (Arne Georg Gisnås Gleditsch)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2019/05/5109528393_b81fe0f685_b-150x150.jpg)](https://webrtchacks.com/bisecting-chrome-regressions/)[Bisecting Browser Bugs (Arne Georg Gisnås Gleditsch)](https://webrtchacks.com/bisecting-chrome-regressions/)
-   [![WebRTC Externals – the cross-browser WebRTC debug extension](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2017/07/7718086616_d2da413f3d_h-150x150.jpg)](https://webrtchacks.com/webrtc-externals/)[WebRTC Externals – the cross-browser WebRTC debug extension](https://webrtchacks.com/webrtc-externals/)
-   [![How Janus Battled libFuzzer and Won (Alessandro Toppi)](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/uploads/2019/03/janus-vs-the-fuzzer-150x150.png)](https://webrtchacks.com/fuzzing-janus/)[How Janus Battled libFuzzer and Won (Alessandro Toppi)](https://webrtchacks.com/fuzzing-janus/)

[![Tweet about this on Twitter](https://webrtchacks.com/wp-content/themes/parament/images/Twitter.png "Twitter")](http://twitter.com/share?url=https://webrtchacks.com/video_replay/&text=How%20to%20capture%20%26%20replay%20WebRTC%20video%20streams%20with%20video_replay%20%28Stian%20Selnes%29%20)[![Share on Facebook](https://webrtchacks.com/wp-content/themes/parament/images/facebook-favicon.png "Facebook")](http://www.facebook.com/sharer.php?u=https://webrtchacks.com/video_replay/)[![Share on LinkedIn](https://ltwus2ix28x10gixx34jeigv-wpengine.netdna-ssl.com/wp-content/plugins/simple-share-buttons-adder/buttons/plain/linkedin.png "LinkedIn")](http://www.linkedin.com/shareArticle?mini=true&url=https://webrtchacks.com/video_replay/)[![Share on Google+](https://webrtchacks.com/wp-content/themes/parament/images/google+-favicon.png "Google+")](https://plus.google.com/share?url=https://webrtchacks.com/video_replay/)

# Posts navigation

←  [WebRTC Externals – the cross-browser WebRTC debug extension](https://webrtchacks.com/webrtc-externals/)

[SDP: Your Fears Are Unleashed (Iñaki Baz Castillo)](https://webrtchacks.com/webrtc-sdp-inaki-baz-castillo/)  →

## 10 comments on “How to capture & replay WebRTC video streams with video_replay (Stian Selnes)”

1.  ![](https://secure.gravatar.com/avatar/481a37ee40baa8a2d95708ec09b4dac2?s=40&d=retro&r=r)
    
    [Lorenzo Miniero](https://janus.conf.meetecho.com/)on  [August 31, 2017 at 11:44 am](https://webrtchacks.com/video_replay/#comment-17813)  said:
    
    Interesting article, thanks for sharing!
    
    To share our experience, we have a similar feature in Janus, which we mostly use for recording but sometimes for debugging as well. Our recordings basically consist in a structured dump of the plain RTP packets we receive (so after the Janus core has decrypted the SRTP layer) to a custom format, and we have a simple tool to go through that file offline and depacketize the payload to a playable format (e.g., webm), in a similar but probably less sophisticated way to how video_replay seems to work. The cool thing is that we don’t need to disable encryption, which wouldn’t work anyway when setting up the PeerConnection with Janus, but the custom format makes it less reusable (e.g., for bug reports). Adding support for rtpdump and/or pcap as well, besides our custom target file, would indeed be an interesting way to make debugging even easier, so thanks again for this cool tutorial!
    
    [Reply](https://webrtchacks.com/video_replay/#comment-17813)
    
    -   ![](https://secure.gravatar.com/avatar/d977dd1b0b2d64745ccddc9c4f00520d?s=40&d=retro&r=r)
        
        [Philipp Hancke](https://github.com/fippo)on  [August 31, 2017 at 1:39 pm](https://webrtchacks.com/video_replay/#comment-17814)  said:
        
        on the serverside logging in text2pcap format is pretty simple and powerful. grep from logs + convert with  `text2pcap -D -n -u 1000,2000 -t "%H:%M:%S." \!:1 \!:2`  is magic (courtesy of randell jesup)  ![😉](https://s.w.org/images/core/emoji/11.2.0/svg/1f609.svg)
        
        [Reply](https://webrtchacks.com/video_replay/#comment-17814)
        
        -   ![](https://secure.gravatar.com/avatar/481a37ee40baa8a2d95708ec09b4dac2?s=40&d=retro&r=r)
            
            [Lorenzo Miniero](https://janus.conf.meetecho.com/)on  [September 1, 2017 at 8:07 am](https://webrtchacks.com/video_replay/#comment-17817)said:
            
            Will definitely have a look at that, thanks for the tip!
            
            [Reply](https://webrtchacks.com/video_replay/#comment-17817)
            
            -   ![](https://secure.gravatar.com/avatar/d977dd1b0b2d64745ccddc9c4f00520d?s=40&d=retro&r=r)
                
                [Philipp Hancke](https://github.com/fippo)on  [September 5, 2017 at 4:05 pm](https://webrtchacks.com/video_replay/#comment-17820)  said:
                
                apparently we have to not only thank Randell but also Michael Tüxen who originally came up with it. So… thank you Michael  ![🙂](https://s.w.org/images/core/emoji/11.2.0/svg/1f642.svg)
                
                [Reply](https://webrtchacks.com/video_replay/#comment-17820)
                
                -   ![](https://secure.gravatar.com/avatar/481a37ee40baa8a2d95708ec09b4dac2?s=40&d=retro&r=r)
                    
                    [Lorenzo Miniero](https://janus.conf.meetecho.com/)on  [September 6, 2017 at 5:01 am](https://webrtchacks.com/video_replay/#comment-17821)  said:
                    
                    Yep, usrsctp has the same feature, very helpful!
                    
2.  Pingback:  [Debugging WebRTC errors with video_replay – Pexip | Your world connected](http://blog.pexip.com/2017/08/31/debugging-webrtc-errors-with-video_replay/)
    
3.  ![](https://secure.gravatar.com/avatar/0d54f31b3a508bdae4fb3e26cedfd406?s=40&d=retro&r=r)
    
    Matthieuon  [September 3, 2017 at 4:23 am](https://webrtchacks.com/video_replay/#comment-17818)  said:
    
    This is really good stuff !  
    However I have a problem generating the Ninja project file, Running the [gn gen out/Default] command on a windows environment (from a command prompt started with administrator privilege).  
    I don’t see any error in the previous phases, but when running the generation I get the below error. Has any one seen this ?  
    (unfortunately I only have a windows environment at home…)
    
    ERROR at //third_party/protobuf/proto_library.gni:229:15: File is not inside out  
    put directory.  
    outputs = get_path_info(protogens, “abspath”)  
    ^———————————  
    The given file should be in the output directory. Normally you would specify  
    “$target_out_dir/foo” or “$target_gen_dir/foo”. I interpreted this as  
    “//out/Default/gen/webrtc/rtc_tools/event_log_visualizer/chart.pb.h”.  
    See //webrtc/rtc_tools/BUILD.gn:184:3: whence it was called.  
    proto_library(“chart_proto”) {  
    ^—————————–  
    See //BUILD.gn:16:5: which caused the file to be included.  
    “//webrtc/rtc_tools”,  
    ^——————-  
    Traceback (most recent call last):  
    File “D:/temp/webrtc-checkout/src/build/vs_toolchain.py”, line 459, in  
    sys.exit(main())  
    File “D:/temp/webrtc-checkout/src/build/vs_toolchain.py”, line 455, in main  
    return commands[sys.argv[1]](*sys.argv[2:])  
    File “D:/temp/webrtc-checkout/src/build/vs_toolchain.py”, line 431, in GetTool  
    chainDir  
    win_sdk_dir = SetEnvironmentAndGetSDKDir()  
    File “D:/temp/webrtc-checkout/src/build/vs_toolchain.py”, line 424, in SetEnvi  
    ronmentAndGetSDKDir  
    return NormalizePath(os.environ[‘WINDOWSSDKDIR’])  
    File “D:\temp\depot_tools\win_tools-2_7_6_bin\python\bin\lib\os.py”, line 423,  
    in __getitem__  
    return self.data[key.upper()]  
    KeyError: ‘WINDOWSSDKDIR’
    
    [Reply](https://webrtchacks.com/video_replay/#comment-17818)
    
4.  ![](https://secure.gravatar.com/avatar/32f702d68608c3c35fa767ea7f53973c?s=40&d=retro&r=r)
    
    Kevin Connoron  [September 5, 2017 at 3:12 am](https://webrtchacks.com/video_replay/#comment-17819)  said:
    
    For replay of audio RTP streams, captured or simulated, I’ve found SIPP to be a useful tool to exercise jitter buffers and reproduce errors as it plays out with more or less correct timing, and handles the sip signalling. Does anyone know any other RTP regurgitators? Fantastic post, thanks for sharing and documenting!
    
    [Reply](https://webrtchacks.com/video_replay/#comment-17819)
    
5.  Pingback:  [Ways to capture incoming WebRTC video streams (client side) – program faq](http://www.programfaqs.com/faq/ways-to-capture-incoming-webrtc-video-streams-client-side/)
    
6.  ![](https://secure.gravatar.com/avatar/240e5c5e27bde3065041c0568dafe04e?s=40&d=retro&r=r)
    
    Anandon  [January 16, 2018 at 3:09 pm](https://webrtchacks.com/video_replay/#comment-17871)  said:
    
    Great Article!!
    
    To decode H264, you need to explicitly request it at makefile creation time
    
    gn gen out/Default –args=’rtc_use_h264=true ffmpeg_branding=”Chrome”‘
    
    ninja -C out/Default video_replay
    
    video_replay –input_file d3.rtpdump –ssrc 2176537790 –media_payload_type 96 –codec H264
<!--stackedit_data:
eyJkaXNjdXNzaW9ucyI6eyJvUmxYMExvSUNDN3prd2VlIjp7In
N0YXJ0IjozOTg5LCJlbmQiOjQwMDEsInRleHQiOiJyZXByb2R1
Y2libGUifSwiSm82Sjh4SHU1dE9YVzJ1ViI6eyJzdGFydCI6ND
AyOCwiZW5kIjo0MDM4LCJ0ZXh0IjoiYXR0cmFjdGl2ZSJ9LCJt
Mk1PSXR0VGdEa1FPekduIjp7InN0YXJ0IjozNjE3LCJlbmQiOj
M2MjgsInRleHQiOiJmcnVzdHJhdGluZyJ9LCJOT3g0bzZ0MW84
bzBsTjJuIjp7InN0YXJ0Ijo0NzY1LCJlbmQiOjQ3NzIsInRleH
QiOiJ0cml2aWFsIn0sIlRtNHdNaFlGV2s3RTU3bnAiOnsic3Rh
cnQiOjQ5OTgsImVuZCI6NTAwNCwidGV4dCI6IkNhbmFyeSJ9LC
JZMkdvelZpM1JrVFh2UDJuIjp7InN0YXJ0Ijo2MDc4LCJlbmQi
OjYxMjQsInRleHQiOiJBbmF0b215XShodHRwczovL3dlYnJ0Y2
hhY2tzLmNvbS9zZHAtYW5hdG9teS8pIn0sInM5YXBETHpGd2Jt
M1h1TUkiOnsic3RhcnQiOjcxMzAsImVuZCI6NzEzNSwidGV4dC
I6InN0aWNrIn0sInh3bFB1bE1Oa1VZaEVMaGQiOnsic3RhcnQi
Ojc5NDcsImVuZCI6ODAyNCwidGV4dCI6IkJVTkRMRV0oaHR0cH
M6Ly90b29scy5pZXRmLm9yZy9odG1sL2RyYWZ0LWlldGYtbW11
c2ljLXNkcC1idW5kbGUtbmVnb3RpYXRpb24pIn0sIlRsWW8wVG
JSdGNsbEc1ZGwiOnsic3RhcnQiOjEwMjQ1LCJlbmQiOjEwMjQ5
LCJ0ZXh0IjoiYm9sZCJ9LCJLVVN2Uks3ODloeUxZc01nIjp7In
N0YXJ0IjoxMjQ0OSwiZW5kIjoxMjUyMiwidGV4dCI6InByZXJl
cXVpc2l0ZV0oaHR0cHM6Ly93ZWJydGMub3JnL25hdGl2ZS1jb2
RlL2RldmVsb3BtZW50L3ByZXJlcXVpc2l0ZS1zdykifSwieDE4
ZWlFMHZYV0Q1UDIzbCI6eyJzdGFydCI6MzQyMiwiZW5kIjozND
MwLCJ0ZXh0IjoidG91Z2hlc3QifX0sImNvbW1lbnRzIjp7ImJH
NzdwN25PUTdOd1VJWkQiOnsiZGlzY3Vzc2lvbklkIjoib1JsWD
BMb0lDQzd6a3dlZSIsInN1YiI6ImdvOjExMjYxMTYzMzY5ODE0
NDMyNTQyMCIsInRleHQiOiJyZXByb2R1Y2libGUgWyxyaXByyZ
knZGrKinPJmWJsXSDor6bnu4bCuyBhZGouIOWPr+WGjeeUn+ea
hO+8m+WPr+e5geaulueahO+8m+WPr+WkjeWGmeeahCIsImNyZW
F0ZWQiOjE1NjA0ODc2MjE0MDB9LCJxRzkxdDJQRkhWZU5aajl6
Ijp7ImRpc2N1c3Npb25JZCI6IkpvNko4eEh1NXRPWFcydVYiLC
JzdWIiOiJnbzoxMTI2MTE2MzM2OTgxNDQzMjU0MjAiLCJ0ZXh0
IjoiJCB5ZGljdCBhdHRyYWN0aXZlIFxuXG4gICAg6IuxOiBbyZ
kndHLDpmt0yap2XSAgICDnvo46IFvJmSd0csOma3TJqnZdXG5c
biAgICAgYWRqLiDlkLjlvJXkurrnmoTvvJvmnInprYXlipvnmo
TvvJvlvJXkurrms6jnm67nmoQiLCJjcmVhdGVkIjoxNTYwNDg5
NTUzNjY2fSwidW1INlhCQ3AweUs2RE9IWCI6eyJkaXNjdXNzaW
9uSWQiOiJtMk1PSXR0VGdEa1FPekduIiwic3ViIjoiZ286MTEy
NjExNjMzNjk4MTQ0MzI1NDIwIiwidGV4dCI6IiQgeSBmcnVzdH
JhdGluZ1xuXG4gICAg6IuxOiBbZnLKjCdzdHJlyap0yarFi10g
ICAg576OOiBbJ2ZyyoxzdHJldMmqxYtdXG5cbiAgICAgYWRqLi
Dku6Tkurrmsq7kuKfnmoRcbiAgICAgdi4g5L2/5rKu5Lin77yI
ZnJ1c3RyYXRl55qEaW5n5b2i5byP77yJIiwiY3JlYXRlZCI6MT
U2MDQ4OTY1MzU2Mn0sImUzRDNGd0xNTHMwbVJGTTQiOnsiZGlz
Y3Vzc2lvbklkIjoiTk94NG82dDFvOG8wbE4ybiIsInN1YiI6Im
dvOjExMjYxMTYzMzY5ODE0NDMyNTQyMCIsInRleHQiOiIkIHkg
dHJpdmlhbFxuXG4gICAg6IuxOiBbJ3Ryyap2yarJmWxdICAgIO
e+jjogWyd0csmqdsmqyZlsXVxuXG4gICAgIGFkai4g5LiN6YeN
6KaB55qE77yM55CQ56KO55qE77yb55CQ57uG55qEIiwiY3JlYX
RlZCI6MTU2MDQ5MjAxMDkwOH0sInJ4UjlsNGZEVTJsMFk3b2Yi
OnsiZGlzY3Vzc2lvbklkIjoiVG00d01oWUZXazdFNTducCIsIn
N1YiI6ImdvOjExMjYxMTYzMzY5ODE0NDMyNTQyMCIsInRleHQi
OiIkIHkgY2FuYXJ5XG5cbiAgICDoi7E6IFtryZknbmXJmXLJql
0gICAg576OOiBba8mZJ27Jm3JpXVxuXG4gICAgIG4uIOmHkeS4
nembgO+8m+a3oem7hOiJsu+8jOmynOm7hOiJsu+8m+WKoOe6s+
WIqeeUnOmFklxuICAgICBuLiAoQ2FuYXJ5KSDvvIjnvo7jgIHl
jbDjgIHoi7HvvInljaHnurPph4wiLCJjcmVhdGVkIjoxNTYwND
kyMDg5MTg4fSwieVMyYWZMUUE5b0J0b0VWOCI6eyJkaXNjdXNz
aW9uSWQiOiJZMkdvelZpM1JrVFh2UDJuIiwic3ViIjoiZ286MT
EyNjExNjMzNjk4MTQ0MzI1NDIwIiwidGV4dCI6IiQgeSBhbmF0
b215XG5cbiAgICDoi7E6IFvJmSduw6Z0yZltyapdICAgIOe+jj
ogW8mZJ27DpnTJmW1pXVxuXG4gICAgIG4uIOino+WJlu+8m+in
o+WJluWtpu+8m+WJluaekO+8m+mqqOmqvCIsImNyZWF0ZWQiOj
E1NjA0OTI4NDQ4Mzd9LCJTTUdIMGJLd2JFWEk4bVZaIjp7ImRp
c2N1c3Npb25JZCI6InM5YXBETHpGd2JtM1h1TUkiLCJzdWIiOi
JnbzoxMTI2MTE2MzM2OTgxNDQzMjU0MjAiLCJ0ZXh0IjoiJCB5
IHN0aWNrXG5cblxuICAgIOiLsTogW3N0yaprXSAgICDnvo46IF
tzdMmqa11cblxuICAgICB2dC4g5Yi677yM5oiz77yb5Ly45Ye6
77yb57KY6LS0XG4gICAgIHZpLiDlnZrmjIHvvJvkvLjlh7rvvJ
vnspjkvY9cbiAgICAgbi4g5qON77yb5omL5p2W77yb5ZGG5aS0
5ZGG6ISR55qE5Lq6XG4gICAgIG4uIChTdGljaynkurrlkI3vvJ
so6IqsKeaWr+iSguWFiyIsImNyZWF0ZWQiOjE1NjA0OTMxMjY0
MTN9LCJ2VzBhZXoxQ2pMSUpsN3VzIjp7ImRpc2N1c3Npb25JZC
I6Inh3bFB1bE1Oa1VZaEVMaGQiLCJzdWIiOiJnbzoxMTI2MTE2
MzM2OTgxNDQzMjU0MjAiLCJ0ZXh0Ijoi6IuxOiBbJ2LKjG5kKM
mZKWxdICAgIOe+jjogWydiyoxuZGxdXG5cbiAgICAgbi4g5p2f
77yb5o2GXG4gICAgIHZ0LiDmjYZcbiAgICAgdmkuIOWMhuW/me
emu+W8gCIsImNyZWF0ZWQiOjE1NjA0OTQyNDAyMzh9LCJiN0hJ
ODVwNXBrMmplVUpMIjp7ImRpc2N1c3Npb25JZCI6IlRsWW8wVG
JSdGNsbEc1ZGwiLCJzdWIiOiJnbzoxMTI2MTE2MzM2OTgxNDQz
MjU0MjAiLCJ0ZXh0IjoiJCB5IGJvbGRcblxuICAgIOiLsTogW2
LJmcqKbGRdICAgIOe+jjogW2JvbGRdXG5cbiAgICAgYWRqLiDl
pKfog4bnmoTvvIzoi7Hli4fnmoTvvJvpu5HkvZPnmoTvvJvljp
rpopzml6DogLvnmoTvvJvpmanls7vnmoRcbiAgICAgbi4gKEJv
bGQp5Lq65ZCN77ybKOiLseOAgeW+t+OAgee9l+OAgeaNt+OAge
eRnuWFuCnljZrlsJTlvrciLCJjcmVhdGVkIjoxNTYwNDk0NTQ2
NzY2fSwiajl1V1hYWm1Ja1Z1bk1lQiI6eyJkaXNjdXNzaW9uSW
QiOiJLVVN2Uks3ODloeUxZc01nIiwic3ViIjoiZ286MTEyNjEx
NjMzNjk4MTQ0MzI1NDIwIiwidGV4dCI6IiQgeSBwcmVyZXF1aX
NpdGVcblxuICAgIOiLsTogW3ByacuQJ3Jla3fJqnrJqnRdICAg
IOe+jjogWyxwcmkncsmba3fJmXrJqnRdXG5cbiAgICAgbi4g5Y
WI5Yaz5p2h5Lu2XG4gICAgIGFkai4g6aaW6KaB5b+F5aSH55qE
IiwiY3JlYXRlZCI6MTU2MDQ5NDg4MDcyNn0sIjBxZklIVUY1cU
pJQ3NHMXoiOnsiZGlzY3Vzc2lvbklkIjoieDE4ZWlFMHZYV0Q1
UDIzbCIsInN1YiI6ImdoOjM5NzY3MDMyIiwidGV4dCI6IiQgeS
B0b3VnaFxuXG4gICAg6IuxOiBbdMqMZl0gICAg576OOiBbdMqM
Zl1cblxuICAgICBhZGouIOiJsOiLpueahO+8jOWbsOmavueahO
+8m+WdmuW8uueahO+8jOS4jeWxiOS4jeaMoOeahO+8m+Wdmumf
p+eahO+8jOeJouWbuueahO+8m+W8uuWjrueahO+8jOe7k+Wunu
eahFxuICAgICBuLiDmgbbmo41cbiAgICAgdnQuIOWdmuaMge+8
m+W/jeWPl++8jOW/jeiAkFxuICAgICBhZHYuIOW8uuehrOWcsO
+8jOmhveW8uuWcsFxuICAgICBuLiAoVG91Z2gp5Lq65ZCN77yb
KOiLsSnlm77otasiLCJjcmVhdGVkIjoxNTYxMzY2ODA1MjE5fX
0sImhpc3RvcnkiOlstNzMxNDg3OTI4LDU5NzU2OTA3MSwtNDM0
NjY5MjM4XX0=
-->