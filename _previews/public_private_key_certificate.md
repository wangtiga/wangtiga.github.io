---
layout: post
title:  "公钥私钥和证书"
date:   2022-02-04 18:00:00 +0800
tags:   tech
---

* category
{:toc}


### 生成RSA密钥对
生成根密钥和根证书
```sh
openssl genrsa -out certstd/cakey.pem 2048
openssl req -new -x509 -key certstd/cakey.pem -out certstd/cacert.pem -days 7300 -subj /C=CN/ST=BeiJing/O=QIANXIN/OU=Data_Security_Group/CN=QuickCa
```

生成app密钥对
```sh
openssl genrsa -out certstd/app1.keypair.pem 2048
openssl pkcs8 -topk8 -inform PEM -in certstd/app1.keypair.pem -outform PEM -nocrypt -out certstd/app1.keypair.p8.pem
```

生成证书
```sh
openssl req -new -key certstd/app1.keypair.p8.pem -out certstd/app1.cert.csr -subj /C=CN/CN=app1
openssl ca -in certstd/app1.cert.csr -out certstd/app1.cert.cer -days 3650 -batch -policy policy_anything -cert certstd/cacert.pem -keyfile certstd/cakey.pem
```
> 必须有这些文件才能生成 appcert ： mkdir -p ./demoCA/newcerts && touch ./demoCA/index.txt && cp certgm/rootcert.srl demoCA/serial

生成证书 p12 格式
```sh
openssl pkcs12 -export -in certstd/app1.cert.cer -inkey certstd/app1.keypair.pem -certfile certstd/cacert.pem -out certstd/app1.keypair.cert.p12 -password pass:pwd123
```

### 生成GM密钥对
生成根证书
```sh
./gmssl ecparam -config ./ -genkey -name sm2p256v1 -out certgm/cakey.pem 
./gmssl req -config ./ -new -sm3 -key certgm/cakey.pem -out certgm/cacsr.pem -subj "/C=CN/ST=Beijing/L=Beijing/O=QAX/OU=MM/CN=Quick/emailAddress=quick@qianxin.com" 
./gmssl x509 -req -set_serial 1 -in certgm/cacsr.pem -days 3650 -sm3 -outform pem -out certgm/cacert.pem -signkey certgm/cakey.pem 
```

生成密钥对（经过master加密，可直接存储mysql）
```sh
./kmstool gen-asym-key -c ../../../custom/config.yaml -o certgm/app2.keypair.bin.cipher.base64
```

将密钥对解密并转换成 p8.pem 格式，方便 gmssl 使用
```sh
./kmstool decrypt-by-master-key -c ../../../custom/config.yaml -i certgm/app2.keypair.bin.cipher.base64 -o certgm/app2.keypair.bin
./kmstool keypair -i certgm/app2.keypair.bin -o ./certgm/app2.keypair.p8.pem
```

生成证书
```sh
./gmssl req -config ./ -new -utf8 -sm3 -key certgm/app2.keypair.p8.pem -out certgm/app2.cert.csr -subj "/CN=Quick"
./gmssl x509 -req -in certgm/app2.cert.csr -extensions v3_req -days 3650 -sm3 -outform pem -out certgm/app2.cert.cer -CA certgm/cacert.pem -CAkey certgm/cakey.pem -CAserial certgm/rootcert.srl
```

使用证书公钥加密
```sh
./kmstool encrypt-by-pubkey -c ../../../custom/config.yaml -p certgm/app2.cert.cer -t 12345
```

使用密钥对私钥解密
```sh
 ./kmstool decrypt-by-privkey -c ../../../custom/config.yaml -k certgm/app2.keypair.bin -t MG0CIB+7GZcCttflib0W5S/xJryjmtuMNfvqgniODrAM9QEVAiABx0A5wBa/EyO8Q+ftjJl7SH6zXuBxsoKqjz9VB1MqNgQg0vExRDBriV+WPN4xLufV5qkp26jYDKv5wrq7R5LCP9cEBVIzCKk9
```


导入 p8.der.txt 格式密钥对(偶现 ./kmstool p8key2bin Segmentation fault)
```sh
./gmssl pkcs8 -topk8 -inform PEM -in certgm/app2.keypair.p8.pem -outform PEM -nocrypt -out certgm/app2.import.keypair.p8.pem
./gmssl base64 -config ./ -d -in certgm/app2.import.keypair.p8.pem -out  certgm/app2.import.keypair.p8.der.bin
xxd -ps -c 256 certgm/app2.import.keypair.p8.der.bin > certgm/app2.import.keypair.p8.der.txt
./kmstool p8key2bin -i certgm/app2.import.keypair.p8.der.txt -o certgm/app2.import.keypair.bin
```

### tmp

生成密钥对（使用 gmssl 生成）
```sh
./gmssl ecparam -config ./ -genkey -name sm2p256v1 -out certgm/app3.keypair.tmp.pem 
./gmssl pkcs8 -topk8 -inform PEM -in certgm/app3.keypair.tmp.pem -outform PEM -nocrypt -out certgm/app3.keypair.p8.pem
./gmssl base64 -config ./ -d -in certgm/app3.keypair.p8.pem -out  certgm/app3.keypair.p8.der.bin
xxd -ps -c 256 certgm/app3.keypair.p8.der.bin > certgm/app3.keypair.p8.der.txt
./kmstool p8key2bin -i certgm/app3.keypair.p8.der.txt -o certgm/app3.keypair.bin
./kmstool encrypt-by-master-key -c ../../../custom/config.yaml -i certgm/app3.keypair.bin -o certgm/app3.keypair.bin.cipher.base64
```


生成证书
```sh
./gmssl req -config ./ -new -utf8 -sm3 -key certgm/app3.keypair.p8.pem -out certgm/app3.cert.csr -subj "/CN=Quick"
./gmssl x509 -req -in certgm/app3.cert.csr -extensions v3_req -days 3650 -sm3 -outform pem -out certgm/app3.cert.cer -CA certgm/cacert.pem -CAkey certgm/cakey.pem -CAserial certgm/rootcert.srl
```

使用证书公钥加密
```sh
./kmstool encrypt-by-pubkey -c ../../../custom/config.yaml -p certgm/app3.cert.cer -t 12345
```

使用密钥对私钥解密
```sh
 ./kmstool decrypt-by-privkey -c ../../../custom/config.yaml -k certgm/app3.keypair.bin -t MG0CIE99AmcUU4VWwdi0qc/2PYi+IRYdjC3WKqkXkcAWHhZVAiAKGG0jZ0Gce1cn4xAdOZ9W4UvuUjw+CdDrewjaUtnpuwQgPstLj7y3IbIyz6lK4YnLCxviMt3KX28hn48h+eLyJUcEBaL40bHt
```
