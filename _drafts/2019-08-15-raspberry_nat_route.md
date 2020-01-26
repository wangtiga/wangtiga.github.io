


## 3. 使用 ss-redir + iptables-TPROXY 实现透明全局代理（成功）
https://www.zfl9.com/ss-redir.html
https://github.com/zfl9/ss-tproxy.git
https://github.com/shadowsocks/shadowsocks/wiki/Feature-Comparison-across-Different-Versions
https://github.com/shadowsocks/shadowsocks
https://github.com/shadowsocksr-backup/shadowsocksr-libev



## 2. 使用 ss-local + badvpn-tun2socks 实现透明全局代理（失败）
https://github.com/eycorsican/go-tun2socks
https://blog.csdn.net/dog250/article/details/70343230
https://awesomeopensource.com/project/eycorsican/go-tun2socks
https://code.google.com/archive/p/badvpn/wikis/Examples.wikihttps://github.com/ambrop72/badvpn

https://github.com/YahuiWong/ss-tun2socks
https://github.com/eycorsican/go-tun2socks
https://blog.csdn.net/dog250/article/details/70343230
https://github.com/ambrop72/badvpn
https://blog.csdn.net/u011068702/article/details/53899537


## 1. Setting up a Raspberry Pi as an access point in a standalone network (NAT)
[Why assign MAC and IP addresses on Bridge interface](https://unix.stackexchange.com/questions/319979/why-assign-mac-and-ip-addresses-on-bridge-interface)
[raspberrypi documentation git](https://github.com/raspberrypi/documentation.git)
[raspberrypi documentation configuration wireless/access-point](https://www.raspberrypi.org/documentation/configuration/wireless/access-point.md)

[vbrid linux iptables 实现 NAT](http://cn.linux.vbird.org/linux_server/0250simple_firewall.php)  
[csdn ip route](https://blog.csdn.net/u011068702/article/details/53899537)
[arch linux iptables](https://wiki.archlinux.org/index.php/Iptables_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87))


```shell
# Setting up a Raspberry Pi as an access point in a standalone network (NAT)
$ sudo apt install dnsmasq hostapd
$ sudo systemctl stop dnsmasq
$ sudo systemctl stop hostapd

# Configuring a static IP
$ sudo nano /etc/dhcpcd.conf
 interface wlan0
    static ip_address=192.168.123.1/24
    nohook wpa_supplicant


$ sudo service dhcpcd restart

# Configuring the DHCP server (dnsmasq)
$ sudo mv /etc/dnsmasq.conf /etc/dnsmasq.conf.orig
$ sudo nano /etc/dnsmasq.conf
interface=wlan0      # Use the require wireless interface - usually wlan0
dhcp-range=192.168.123.2,192.168.123.20,255.255.255.0,24h
$ sudo systemctl reload dnsmasq

# Configuring the access point host software (hostapd)
# To use the 5 GHz band, you can change the operations mode from hw_mode=g to hw_mode=a. Possible values for hw_mode are:
#   a = IEEE 802.11a (5 GHz)
#   b = IEEE 802.11b (2.4 GHz)
#   g = IEEE 802.11g (2.4 GHz)
#   ad = IEEE 802.11ad (60 GHz)
$ sudo nano /etc/hostapd/hostapd.conf
interface=wlan0
#bridge=br0
driver=nl80211
ssid=wspi
hw_mode=g
channel=7
wmm_enabled=0
macaddr_acl=0
auth_algs=1
ignore_broadcast_ssid=0
wpa=2
wpa_passphrase=252114997
wpa_key_mgmt=WPA-PSK
wpa_pairwise=TKIP
rsn_pairwise=CCMP

# We now need to tell the system where to find this configuration file.
$ sudo nano /etc/default/hostapd
DAEMON_CONF="/etc/hostapd/hostapd.conf"

# Start it up
sudo systemctl unmask hostapd
sudo systemctl enable hostapd
sudo systemctl start hostapd

sudo systemctl status hostapd
sudo systemctl status dnsmasq

# Add routing and masquerade
$ sudo nano /etc/sysctl.conf
net.ipv4.ip_forward=1
$ sudo iptables -t nat -A  POSTROUTING -o eth0 -j MASQUERADE
$ sudo sh -c "iptables-save > /etc/iptables.ipv4.nat"

# Edit /etc/rc.local and add this just above "exit 0" to install these rules on boot.
$ sudo nano /etc/rc.local
iptables-restore < /etc/iptables.ipv4.nat

```

<!--stackedit_data:
eyJoaXN0b3J5IjpbODA4Njk0NDMzLC0xMDM5NDAxOTMwLC0xMj
c1ODkwNzU4LDY4MTIxNjkxMCwxMzc2NjU5OTM2LC0xNjUxODQ3
MzksMjA5MzY2NjQ3MywxMDA4ODY2MjE2LDE3NzMwOTA5NzUsMj
A1OTY3NDA1OCwxNDgzMTEwNDM2LDg3MTE1Njk5OSwtMjEzMjM4
MDYwNSwtMTY0MDI4NzAzMiwtMTQyNjU1NTU4OF19
-->