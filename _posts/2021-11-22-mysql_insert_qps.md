---
layout: post
title: "MySQL 插入耗时大概是多少？"
date: 2021-11-22 12:00:00 +0800
tags: tech
---

* category
{:toc}



##  MySQL 插入耗时大概是多少？

[^qcrao_architecture]: [搞定系统设计 02：估算的一些方法 ](https://mp.weixin.qq.com/s/fH-AJpE99ulSLbC_1jxlqw)
[^gozero_shorturl]: [gozero shorturl ](https://go-zero.dev/cn/shorturl.html?h=short)


百万条数据内，耗时在 ms 级别，即 1000qps 上下。

> 具体耗时情况肯定有很多影响因素。这里只做一个简易测试，用于系统设计时进行估算。[^qcrao_architecture]


### 百万数据下 0.001 sec [^gozero_shorturl]


```sql
mysql> set profiling=1;
Query OK, 0 rows affected, 1 warning (0.00 sec)

mysql> INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test18', 'https://test.com/test18', '0');
Query OK, 1 row affected (0.01 sec)

mysql> select count(shorten) from shorturl;
+----------------+
| count(shorten) |
+----------------+
|         970772 |
+----------------+
1 row in set (0.57 sec)

mysql> SHOW PROFILES;
+----------+------------+----------------------------------------------------------------------------------------------------------------+
| Query_ID | Duration   | Query                                                                                                          |
+----------+------------+----------------------------------------------------------------------------------------------------------------+
|        2 | 0.00132700 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test7', 'https://test.com/test7', '0')   |
|        3 | 0.00133600 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test8', 'https://test.com/test8', '0')   |
|        4 | 0.00118900 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test9', 'https://test.com/test9', '0')   |
|        5 | 0.00140300 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test10', 'https://test.com/test10', '0') |
|        6 | 0.00125100 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test11', 'https://test.com/test11', '0') |
|        7 | 0.00106000 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test12', 'https://test.com/test12', '0') |
|        8 | 0.00058900 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test13', 'https://test.com/test13', '0') |
|        9 | 0.00157900 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test14', 'https://test.com/test14', '0') |
|       10 | 0.00111500 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test15', 'https://test.com/test15', '0') |
|       11 | 0.00114600 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test16', 'https://test.com/test16', '0') |
|       12 | 0.00140800 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test17', 'https://test.com/test17', '0') |
|       13 | 0.00115100 | desc shorturl                                                                                                  |
|       14 | 0.00028900 | set profiling=1                                                                                                |
|       15 | 0.00119400 | INSERT INTO `gozero`.`shorturl` (`shorten`, `url`, `visits`) VALUES ('test18', 'https://test.com/test18', '0') |
|       16 | 0.57214400 | select count(shorten) from shorturl                                                                            |
+----------+------------+----------------------------------------------------------------------------------------------------------------+
15 rows in set, 1 warning (0.00 sec)

mysql> desc shorturl;
+----------+--------------+------+-----+-------------------+-----------------------------+
| Field    | Type         | Null | Key | Default           | Extra                       |
+----------+--------------+------+-----+-------------------+-----------------------------+
| shorten  | varchar(255) | NO   | PRI | NULL              |                             |
| url      | varchar(255) | NO   |     | NULL              |                             |
| visits   | int(11)      | NO   |     | 0                 |                             |
| createat | timestamp    | NO   |     | CURRENT_TIMESTAMP |                             |
| updateat | timestamp    | NO   |     | CURRENT_TIMESTAMP | on update CURRENT_TIMESTAMP |
+----------+--------------+------+-----+-------------------+-----------------------------+
5 rows in set (0.01 sec)
```

### 1秒钟 5000 次 insert

```sh
% mysqlslap -h127.0.0.1 -uroot -p123456 --concurrency=1 --iterations=1 --auto-generate-sql --auto-generate-sql-load-type=write --auto-generate-sql-add-autoincrement --engine=innodb --number-of-queries=5000

mysqlslap: [Warning] Using a password on the command line interface can be insecure. 
Benchmark 
Running for engine innodb 
Average number of seconds to run all queries: 1.198 seconds 
Minimum number of seconds to run all queries: 1.198 seconds 
Maximum number of seconds to run all queries: 1.198 seconds 
Number of clients running queries: 1 
Average number of queries per client: 5000 

% mysqlslap -h127.0.0.1 -uroot -p123456 --concurrency=1 --iterations=1 --auto-generate-sql --auto-generate-sql-load-type=write --auto-generate-sql-add-autoincrement --engine=innodb --number-of-queries=1000000 

mysqlslap: [Warning] Using a password on the command line interface can be insecure. 
Benchmark 
Running for engine innodb 
Average number of seconds to run all queries: 226.226 seconds 
Minimum number of seconds to run all queries: 226.226 seconds 
Maximum number of seconds to run all queries: 226.226 seconds 
Number of clients running queries: 1 
Average number of queries per client: 1000000 


% mysqlslap -h127.0.0.1 -uroot -p123456 --concurrency=1 --iterations=1 --auto-generate-sql --auto-generate-sql-load-type=write --auto-generate-sql-add-autoincrement --engine=innodb --number-of-queries=1 --only-print 

DROP SCHEMA IF EXISTS `mysqlslap`; 
CREATE SCHEMA `mysqlslap`; 
use mysqlslap; 
set default_storage_engine=`innodb`; 
CREATE TABLE `t1` (id serial,intcol1 INT(32) ,charcol1 VARCHAR(128)); 
INSERT INTO t1 VALUES (NULL,1804289383,'mxvtvmC9127qJNm06sGB8R92q2j7vTiiITRDGXM9ZLzkdekbWtmXKwZ2qG1llkRw5m9DHOFilEREk3q7oce8O3BEJC0woJsm6uzFAEynLH2xCsw1KQ1lT4zg9rdxBL'); 
...
INSERT INTO t1 VALUES (NULL,2010720737,'iLEkea9eOuLJweKJjQ3IJlW1dXTSs6dbJjJpHTRZFgkBJ4XuNQ4iCRchy51r3WQEEdvNyMxDvZHEeg0164bINBsM9l54GNBM06lrTa4G8LWe1Mf8I6IiA24Jo1FwOQ'); 
DROP SCHEMA IF EXISTS `mysqlslap`;
```


### 系统版本及配置

```sh
mysql> SHOW VARIABLES LIKE "%version%";
+-------------------------+-----------------------+
| Variable_name           | Value                 |
+-------------------------+-----------------------+
| innodb_version          | 5.7.35                |
| protocol_version        | 10                    |
| slave_type_conversions  |                       |
| tls_version             | TLSv1,TLSv1.1,TLSv1.2 |
| version                 | 5.7.35                |
| version_comment         | Homebrew              |
| version_compile_machine | x86_64                |
| version_compile_os      | osx10.15              |
+-------------------------+-----------------------+
8 rows in set (0.00 sec)
```

```sh
% system_profiler SPHardwareDataType SPSoftwareDataType
Hardware:

    Hardware Overview:

      Model Name: MacBook Pro
      Model Identifier: MacBookPro16,2
      Processor Name: Quad-Core Intel Core i5
      Processor Speed: 2 GHz
      Number of Processors: 1
      Total Number of Cores: 4
      L2 Cache (per Core): 512 KB
      L3 Cache: 6 MB
      Hyper-Threading Technology: Enabled
      Memory: 16 GB
      Boot ROM Version: 1554.120.19.0.0 (iBridge: 18.16.14663.0.0,0)
      Serial Number (system): tiga
      Hardware UUID: tiga-tiga-tiga-tiga-tiga
      Activation Lock Status: Disabled

Software:

    System Software Overview:

      System Version: macOS 10.15.7 (19H15)
      Kernel Version: Darwin 19.6.0
      Boot Volume: tiga202106
      Boot Mode: Normal
      Computer Name: tiga的MacBook Pro
      User Name: tiga (tiga)
      Secure Virtual Memory: Enabled
      System Integrity Protection: Enabled
      Time since boot: 7:24

```

