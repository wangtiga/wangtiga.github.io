---
layout: post
title: "数据速率、带宽与数据量的单位"
date: 2019-07-14 00:00:00 +0800
tags: network
---

* category
{:toc}

本文主要试图理清楚以下两个问题：
- 问题1：为什么我办理的百兆带宽，下载速度不到100Mb/s？
- 问题2：为什么买了一个256GB的硬盘，实际可用空间没有这么大呢？

## 一、数据速率(datarate)
数据进行传送(通信)的速率，也可叫信息速率

### 比特速率(bit/s)
每秒传送的信息量

### 码元速率(Baud/s)
每秒传送的码元数量。
调制方法不同，码一个码元可能对应多个比特。
GMSK调制(GSM/GMRS)一个码元对应一个比特。
详细参考《大话移动通信》[^3]#848

## 二、带宽(bandwidth)
通信线路中带宽越宽，其单位时间能传送的数据越多。
与此类似，马路越宽，单位时间能通过的车辆也越多。


### 信号的频带宽度(Hz)
信号所包含的不同频率成分所占据的频率范围，是一种频域称谓。
由于过去很长时间，通信线路传送的是模拟信号，所以使用赫兹(Hz)衡量信道的带宽。
```txt
1kHz = 1000Hz
1MHz = 1000kHz
1GHz = 1000MHz
1THz = 1000GHz
```

### 最高数据率(bit/s)
单位时间内，信道能通过的“最高数据率"，是一种时域称谓。

#### 问题1：为什么我办理的百兆带宽，下载速度不到100Mb/s？
- 宽带运营商(ISP)表示带宽的常用单位是 bit/s ，所以办理宽带时常说的百兆带宽，其实是指下行最高数据率为`100Mb/s=100Mbit/s=100*1000*1000bit/s`

- 浏览器等下载工具显示的单位是 Byte/s ，所以在百兆带宽中下载软件时，理论最快速度是 `100Mb/s = 12.5*8Mb/s=12.5MB/s=11.9MiB/s`，约10多兆每秒。

```txt
1 Byte = 8 bit
1024 Byte = 1 Mega binary Byte
1000 Byte = 1 Mega Byte

100 Mb
=100*1000*1000 bit
=100*1000*1000/8 Byte
=12.5 Mega Byte
=12.5 MB

100 Mb
=100*1000*1000/1024/1024 Mega binary bit
=95.367 Mega binary bit
=95.367/8 Mega binary Byte
=11.920 Mega binary Byte
=11.920 MiB
```

## 三、数据量(datasize)

### KiB是kilo binary byte缩写，指千位二进制字节
国际电工委员会二进制乘数词头(International Electrotechnical Commission)
，IEEE 1541-2002标准[^1]
```txt
1KiB=1024Byte
1MiB=1024KiB=1024*1024Byte
```

### KB是kilo byte的缩写，指千字节
国际单位制词头(SI)，[International System of Units Prefix 标准[^2]
```txt
1KB=1000Byte
1MB=1000KB=1000*1000Byte
```

#### 问题2：为什么买了一个256GB的硬盘，实际可用空间没有这么大呢？
- 硬件、内存等硬件厂商标称的单位是 KB(kilo byte)，但Windows系统显示的是KiB(kilo binary byte)
所以标称256GB的硬件，实际空间是238GiB，再扣除分区空间后，实际可用空间只有230GiB多一点
256GB = 256*1000*1000*1000 Byte= 256*1000*1000*1000/1024/1024/1024=238.41GiB

- Windows系统表示信息量(文件大小）的单位是KiB，但实际显示的单位是KB。
[IEEE 1541-2002]标准[^1]是在2002年制定的。在2002之前，windows系统已经存在十几年了，其显示单位也一直是KB，但实际使用的进KiB定义的单位。虽然新制定的标准与Windows不一致，微软也没有修改自己的问题。不过，在Mac和基于Linux的多数OS中都遵守了[IEEE 1541-2002]标准[^1]。

- KiB主要在计算机领域使用
因为计算机内存使用二进制寻址,使用二进制计量单位相对方便。在计算机发展早期，工程师们使用KB表示1024Byte，因为1024Byte与1000Byte十分接近，所以这个错误影响并不明显，没有人着急修复。
但随着时间发展，常用单位变成GB时，这个误差就十分大，所以必须制定一个新的单位与Kilo Byte区分。Kilo Binary Byte(KiB)自然而然出现。
```txt
1GiB-1GB
=1024*1024*1024Byte - 1000*1000*1000Byte
=73741824Byte
=70Mega binary Byte
```

## 四、有关数据单位的详细说明
### 小b，表示bite
```txt
b = bit = binary digit
b/s = bps = bit per second
```

### 大B，表示Byte
```txt
B = byte = binary byte
B/s = Bps = Byte per second
```

### K, Kilo，表示十进制千
- Kb  = Kilo bit
- 1Kb = 1000 bit

- KB  = Kilo Byte
- 1KB = 1000 Byte

```txt
k(kilo)  = 1000  = 千
M(Mega)  = 1000k = 兆
G(Giga)  = 1000M = 吉
T(Tera)  = 1000G = 太
P(Peta)  = 1000T = 拍
E(Exa)   = 1000P = 艾
Z(Zetta) = 1000E = 泽
```

### Kib，kilo binary，表示二进制千
因为 $2^{10} = 1024$ 与十进制1000很接近，所以使用1024作为二进制千

- Kib = `k`ilo b`i`nary `b`it
- 1kib = 1024bit

- KiB = `k`ilo b`i`nary `B`yte
- 1KiB = 1024Byte


```txt
Kibi(Kilo binary)  = 1024    = 千
Mebi(Mega binary)  = 1024 Ki = 兆
Gibi(Giga binary)  = 1024 Mi = 吉
Tebi(Tera binary)  = 1024 Gi = 太
Pebi(Peta binary)  = 1024 Ti = 拍
Exbi(Exa binary)   = 1024 Pi = 艾
Zebi(Zetta binary) = 1024 Ei = 泽
```

Ei 缩写由来
E = `E`xa
i = b`i`nary


### filesize在不同系统中的表现

#### windows
使用 mintty2.7.3(x86_64-pc-msys)
```shell
dodo@dodo-PC MINGW64 ~/Desktop/tmp2/file
$ dd if=/dev/zero bs=1024 count=1024 of=./1MiB.file
1024+0 records in
1024+0 records out
1048576 bytes (1.0 MB, 1.0 MiB) copied, 0.00339663 s, 309 MB/s

dodo@dodo-PC MINGW64 ~/Desktop/tmp2/file
$ dd if=/dev/zero bs=1000 count=1000 of=./1MB.file
1000+0 records in
1000+0 records out
1000000 bytes (1.0 MB, 977 KiB) copied, 0.00335651 s, 298 MB/s

dodo@dodo-PC MINGW64 ~/Desktop/tmp2/file
$ dd if=/dev/zero bs=1000 count=1000000 of=./1GB.file
1000000+0 records in
1000000+0 records out
1000000000 bytes (1.0 GB, 954 MiB) copied, 3.7019 s, 270 MB/s

dodo@dodo-PC MINGW64 ~/Desktop/tmp2/file
$ dd if=/dev/zero bs=1024 count=1048576 of=./1GiB.file
1048576+0 records in
1048576+0 records out
1073741824 bytes (1.1 GB, 1.0 GiB) copied, 2.10021 s, 511 MB/s

dodo@dodo-PC MINGW64 ~/Desktop/tmp2/file
$ ls -l
total 2027144
-rw-r--r-- 1 dodo 197121 1000000000 Jul 16 14:41 1GB.file
-rw-r--r-- 1 dodo 197121 1073741824 Jul 16 14:41 1GiB.file
-rw-r--r-- 1 dodo 197121    1000000 Jul 16 14:41 1MB.file
-rw-r--r-- 1 dodo 197121    1048576 Jul 16 14:41 1MiB.file

dodo@dodo-PC MINGW64 ~/Desktop/tmp2/file
$ ls -lh
total 2.0G
-rw-r--r-- 1 dodo 197121 954M Jul 16 14:41 1GB.file
-rw-r--r-- 1 dodo 197121 1.0G Jul 16 14:41 1GiB.file
-rw-r--r-- 1 dodo 197121 977K Jul 16 14:41 1MB.file
-rw-r--r-- 1 dodo 197121 1.0M Jul 16 14:41 1MiB.file

dodo@dodo-PC MINGW64 ~/Desktop/tmp2/file
$ ls -lh --si
total 2.1G
-rw-r--r-- 1 dodo 197121 1.0G Jul 16 14:41 1GB.file
-rw-r--r-- 1 dodo 197121 1.1G Jul 16 14:41 1GiB.file
-rw-r--r-- 1 dodo 197121 1.0M Jul 16 14:41 1MB.file
-rw-r--r-- 1 dodo 197121 1.1M Jul 16 14:41 1MiB.file
```

#### linux
使用 CentOS 7.2
```shell
$ cat /etc/redhat-release 
CentOS Linux release 7.2.1511 (Core) 
$ uname -a
Linux centos 3.10.0-327.el7.x86_64 #1 SMP Thu Nov 19 22:10:57 UTC 2015 x86_64 x86_64 x86_64 GNU/Linux

$ bc 1024*1024=1048576
$ dd if=/dev/zero bs=1024 count=1024 of=./1MiB.file
$ dd if=/dev/zero bs=1000 count=1000 of=./1MB.file
$ dd if=/dev/zero bs=1000 count=1000000 of=./1GB.file
$ dd if=/dev/zero bs=1024 count=1048576 of=./1GiB.file
$ ls -l
总用量 2027152
-rw-r--r--. 1 yun yun 1000000000 7月  16 14:15 1GB.file
-rw-r--r--. 1 yun yun 1073741824 7月  16 14:10 1GiB.file
-rw-r--r--. 1 yun yun    1000000 7月  16 14:13 1MB.file
-rw-r--r--. 1 yun yun    1048576 7月  16 14:12 1MiB.file
$ ls -lh
总用量 2.0G
-rw-r--r--. 1 yun yun 954M 7月  16 14:15 1GB.file
-rw-r--r--. 1 yun yun 1.0G 7月  16 14:10 1GiB.file
-rw-r--r--. 1 yun yun 977K 7月  16 14:13 1MB.file
-rw-r--r--. 1 yun yun 1.0M 7月  16 14:12 1MiB.file
$ ls -lh --si
总用量 2.1G
-rw-r--r--. 1 yun yun 1.0G 7月  16 14:15 1GB.file
-rw-r--r--. 1 yun yun 1.1G 7月  16 14:10 1GiB.file
-rw-r--r--. 1 yun yun 1.0M 7月  16 14:13 1MB.file
-rw-r--r--. 1 yun yun 1.1M 7月  16 14:12 1MiB.file
```

```shell
$ man ls
       -h, --human-readable
              with -l, print sizes in human readable format (e.g., 1K 234M 2G)
       --si   likewise, but use powers of 1000 not 1024
```


## 参考
- [《计算机网络》P21 第七版 计算机网络的性能指标](https://book.douban.com/subject/2970300/)
- [《通信基础》P31，4742自考2008年版，信号举例](https://www.google.com/search?q=通信基础+4742)
- [zhihu](https://www.zhihu.com/question/24601215)
- [seagate](https://www.seagate.com/cn/zh/support/kb/storage-capacity-measurement-standards-194563en/)

[^1]: [IEEE 1541-2002](https://en.m.wikipedia.org/wiki/IEEE_1541-2002)
[^2]: [International System of Units Prefix](https://en.m.wikipedia.org/wiki/Metric_prefix)
[^3]: [《大话移动通信》#833，无线信道之烦恼](https://book.douban.com/subject/6876925/)

<!--stackedit_data:
eyJoaXN0b3J5IjpbMTU4NjAwMjgwMyw4NzYxMzcwNDUsOTc3OD
MwMjU2LC0xMzE2OTE3MDU1LDE2NTAxNzg0NDksOTQ0MDAwOTky
LC0xNzIxMDU4NTAzLC02ODE5NDkyNzddfQ==
-->
