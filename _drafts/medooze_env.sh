sudo apt-get install -y  zlib1g-dev
sudo apt-get install -y  libpng-dev
sudo apt-get install -y  libjpeg-dev

sudo apt-get install -y  build-essential

sudo apt-get install -y  libxmlrpc-c++8-dev
sudo apt-get install -y  libgsm1-dev
sudo apt-get install -y  libspeex-dev
sudo apt-get install -y  libopus-dev
sudo apt-get install -y  libavresample-dev
sudo apt-get install -y  libx264-dev
sudo apt-get install -y  libvpx-dev
sudo apt-get install -y  libswscale-dev
sudo apt-get install -y  libavformat-dev
sudo apt-get install -y  libmp4v2-dev
sudo apt-get install -y  libgcrypt11-dev
sudo apt-get install -y  libssl1.0-dev
sudo apt-get install -y  ninja-build


# nodejs
sudo apt-get install curl
curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
sudo apt-get install nodejs
node -v 
npm -v 


# MediaDevices.getUserMedia` undefined 的问题
# https://www.cnblogs.com/Wayou/p/using_MediaDevices_getUserMedia_wihtout_https.html
# chrome://flags/#unsafely-treat-insecure-origin-as-secure

# media-server-go
sudo apt install -y  autoconf
sudo apt install -y  libtool
sudo apt install -y  automake
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test
sudo apt-get update -qq
sudo apt-get install g++-7
sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 90

# media-server-go-native
# git clone --recurse-submodules https://github.com/notedit/media-server-go-native.git  
# cd media-server-go-native
# make
# 如果编译不通过，反复多次编译后，就成功了。 未找到具体原因
# 类似这样： make clean && make mediaserver && cd media-server/ && make  && cd ../ && make clean && make

# https://github.com/notedit/media-server-go-demo
# gstreamer1.0.plugins-bad gstreamer1.0.plugins-good 的区别参考 
sudo apt-get install libgstreamer1.0-0 gstreamer1.0-plugins-base gstreamer1.0-libav gstreamer1.0-plugins-bad libgstreamer-plugins-bad1.0-dev


Notes

### 1.medooze 的 media-server 有没有提供混屏混音的功能呢？
作者说可以用它提供的接口实现 mcs ，那么应该是有相关接口吧。

### 2.media-server 使用的 webrtc 是与 chrome 一样的代码吗？
与 media-process-server 是一样的 webrtc 代码吗？
webrtc 代码实现的是哪此功能？流量控制，dtls srtp 这些吗？有没有音视频编码和解码相关的吗？

### 3.TODO 使用 media-server-go-demo 制作一个 sfu conference demo
or webrtc-to-hls


# openssl lib conflit  
# about pkg-config  https://www.cnblogs.com/rainsoul/p/10567390.html
```shell
sudo locate libcrypto.a libmediaserver.a libmp4v2.a libsrtp.a libssl.a  | xargs  md5sum
```

```shell
# mp3 to mp4 with image
ffmpeg -i sample.mp3 -f image2 -i 1.jpg -acodec aac -strict -2 -vcodec h264 -ar 22050 -ab 128k -ac 2 -pix_fmt yuvj420p -y 1.mp4

```

