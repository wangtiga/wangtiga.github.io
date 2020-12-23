---
layout: post
title:  "InfluxDB 问题记录"
date:   2019-10-09 19:09:00 +0800
tags: tool
---

* category
{:toc}



## InfluxDB [^InfluxdbGettingStart]

### 1.measurement 相当于 SQL 的表， tags 有索引，fields 没有

First, a short primer on the datastore. Data in InfluxDB is organized by `time series`, which contain a measured value, like `cpu_load` or `temperature`. Time series have zero to many points, one for each discrete sample of the metric. Points consist of time (a timestamp), a measurement (`cpu_load`, for example), at least one key-value field (the measured value itself, e.g. `value=0.64`, or `temperature=21.2`), and zero to many key-value tags containing any metadata about the value (e.g. `host=server01`, `region=EMEA`, `dc=Frankfurt`).

简单来说，InfluxDB 中由 "时序" 数据组成。
主要用于记录测量到的指标数据，比如 CPU负载，气温 等类型的数据。
时序数据可理解为，由零到多个时间上离散的取样值组成。每个取样被称为 点 （Point)。

每个 Point 包含：
- time 时间
- measurement 取样类型
  * 比如记录CPU负载的 `cpu_load` 。
- field 取样的指标
  * 至少一个取样指标，否则取样无意义。
  * 比如 负载 value=4 或 温度 temperature=21.2 。
- tag 取样的标签
  * 可以有零到多个标签，用于给取样分类。
  * 比如 所属主机 host=server01 或 所属区域 region=EMEA  或 所属地区 dc=Frankfurt （法兰克福，德国的一个城市）


Conceptually you can think of a measurement as an SQL table, where the primary index is always time. tags and fields are effectively columns in the table. tags are indexed, and fields are not. The difference is that, with InfluxDB, you can have millions of measurements, you don’t have to define schemas up-front, and null values aren’t stored.

可以把 measurement 当做 SQL table ,
time 是主索引，
tag 是加了索引的 column ，
field 是没加索引的 column 。
唯一的不同是， InfluxDB 中可以保存上百万的数据，而且不用在插入数据前定义表结构 (schema) ，而且数据中也不会出现 null 值。

[Writing and exploring data](https://docs.influxdata.com/influxdb/v1.7/introduction/getting-started/#writing-and-exploring-data)



### 2.measurement 中一行数据的 tagKey, tagValue 和 timestamp 三者数值一样，那么 field 会被最新的值替换

A point is uniquely identified by the measurement name, tag set, and timestamp. If you submit a new point with the same measurement, tag set, and timestamp as an existing point, the field set becomes the union of the old field set and the new field set, where any ties go to the new field set. This is the intended behavior.

point 由 measurement name , tag set 和 time 唯一标识。
如果插入的新 point 值中，这三种类型的取值完全一样，那么 point 中的 field set 值将是 新旧数据的并集，对于重复的 field name ，则新数据覆盖旧数据。


>  如果出现 point 覆盖的问题，说明取样的时间精度太粗。比如原本每秒取样，提高为毫秒（ms）甚至纳秒（ns）。
> 还有可能是表结构设计不对，应该增加 tag 。比如统计两台主机的温度时，如果不加 host 的 tag ，那么统计到两台主机在同一时间点的样本值，肯定也应该丢失一个数据。


[InfluxDB FAQ 如何处理重复数据](https://docs.influxdata.com/influxdb/v1.7/troubleshooting/frequently-asked-questions/#how-does-influxdb-handle-duplicate-points)



### 3.select 查询必须包含 fields 字段，否则不返回数据

A query requires at least one field key in the SELECT clause to return data. If the SELECT clause only includes a single tag key or several tag keys, the query returns an empty response. Please see the Data Exploration page for additional information.

SELECT 查询子句必须至少包含一个 feild 。
如果 SELECT 查询子句中仅包含一个或几个 tag ，只会返回一个空的响应。
详细说明参考[Data exploration](https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/#common-issues-with-the-select-statement) 。

[Tag keys in the SELECT clause](https://docs.influxdata.com/influxdb/v1.7/troubleshooting/frequently-asked-questions/#tag-keys-in-the-select-clause)



### 4.max-values-per-tag 限制总的 tag 种类不能超过 10 万 

The number of unique database, measurement, tag set, and field key combinations in an InfluxDB instance.

[series cardinality 数据量级](https://docs.influxdata.com/influxdb/v1.1/concepts/glossary/#series-cardinality) 是 database, measurement, tag set, field key 组合时，非重复数据的数量。

For example, assume that an InfluxDB instance has a single database and one measurement. The single measurement has two tag keys: email and status. If there are three different emails, and each email address is associated with two different statuses then the series cardinality for the measurement is 6 (3 * 2 = 6):

假设 InfluxDB 有一个 database 和一个 measurement 。
这个 measurement 中只有两个 tag key ： email 和 status 。
如果有三个不同的 email 地址，每个 email 又有两种不同的状态，那么总的 series cardinality 值是 `3 *2 = 6` 。

- 越多的 series cardinality 会占用越高的 RAM （内存）。
  详细情况参考 [When do I need more RAM?](https://docs.influxdata.com/influxdb/v1.7/guides/hardware_sizing/#when-do-i-need-more-ram) 。
  所以尽量避免使用 `UUID / hash / 随机字符串` 等种类过多的数据作为 tag ，来[避免 series 数量过多](https://docs.influxdata.com/influxdb/v1.7/concepts/schema_and_data_layout/#don-t-have-too-many-series)
- InfluxDB v1.1 版本时，增加了 [max-values-per-tag](https://github.com/influxdata/influxdb/blob/1.1/CHANGELOG.md#data-section) 选项来限制 series cardinality 。
  此选项限制的是 tag 的种类数量，默认是 100000 。注意它限制的不是十万个数据，而是十万种不同的数据。
  所以即使 measurement 中只有一个 tag ，总的数据量也能超过十万。假设有十一万个数据，每个数据的 host tag 值都是 server01 。
  另外，即使种类超过限制，InfluxDB 也能正常使用，只有 write 操作新增了 host tag 种类数年时，才会报错`max-values-per-tag limit exceeded` 。


[InfluxDB写入失败-partial write: max-values-per-tag limit exceeded](http://xiajunhust.github.io/2016/12/13/InfluxDB%E5%86%99%E5%85%A5%E5%A4%B1%E8%B4%A5-partial-write-max-values-per-tag-limit-exceeded/)

[InfluxDB 单点运行时，对硬件配置的要求](https://docs.influxdata.com/influxdb/v1.7/guides/hardware_sizing/#general-hardware-guidelines-for-a-single-node)



#### 4.1 downsampling and retention (降低采样和自动删除）

InfluxDB can handle hundreds of thousands of data points per second. Working with that much data over a long period of time can create storage concerns. A natural solution is to downsample the data; keep the high precision raw data for only a limited time, and store the lower precision, summarized data for much longer or forever.

InfluxDB 每秒钟能处理成百上千个 point 数据。
长年累月地运行，肯定会遇到存储空间不足的问题。
很自然的解决方案是， downsample 数据。
只保留最近一段时间的高精度数据；
太旧的数据，降低精度，只保留一些概要的统计结果。


InfluxDB offers two features - continuous queries (CQ) and retention policies (RP) - that automate the process of downsampling data and expiring old data. This guide describes a practical use case for CQs and RPs and covers how to set up those features in InfluxDB databases.

InfluxDB 提供了两种功能 - 自动查询(CQ) 和 自动删除(RP) - 两者都用于在旧数据中自动执行降低采样的功能。
[InfluxDB Downsampling and data retention](https://docs.influxdata.com/influxdb/v1.7/guides/downsampling_and_retention/) 文档描述如何配置并使用 CQ 和 RP 这两个功能。

- [A continuous query (CQ)](https://docs.influxdata.com/influxdb/v1.7/query_language/continuous_queries/)
  * is an InfluxQL query that runs automatically and periodically within a database.
    CQs require a function in the SELECT clause and must include a GROUP BY time() clause.
  * CQ 在会定期在数据库中自动执行查询。启用 CQ 功能时，必须配置 SELECT 查询子句，并且其中还要包含`GROUP BY time`。

- [A retention policy (RP)](https://docs.influxdata.com/influxdb/v1.7/query_language/database_management/#retention-policy-management)
  * is the part of InfluxDB data structure that describes for how long InfluxDB keeps data.
    InfluxDB compares your local server’s timestamp to the timestamps on your data and deletes data that are older than the RP’s DURATION. A single database can have several RPs and RPs are unique per database.
  * RP 用于配置 InfluxDB 最多保留数据多长时间。它会自动比较 InfluxDB 所处服务器时间与数据的时间，并删除早于 RP 策略配置的数据。
    一个数据库可以有多个 RP ，但每个 RP 仅能用于一个数据库。

> 设置 CQ 和 RP 策略，就能自动归档数据。也就能间接保证不超过 max-values-per-tag limit，而且保留了旧数据

```log


influx -precision rfc3339 -port 8086

> show databases;
name: databases
----
_internal
wsdb
mydb

> use mydb;
Using database mydb


> show measurements;
name: measurements                                                                                                             │                                                                                                                               
name                                                                                                                           │                                                                                                                               
----                                                                                                                           │                                                                                                                               
shapes                                                                                                                         │                                                                                                                               
tab_monitor_cpu                                                                                                                │                                                                                                                               
tab_monitor_net 


# 自动归档策略
> SHOW RETENTION POLICIES;
> CREATE RETENTION POLICY "auto_delete" ON "mydb" DURATION 5d REPLICATION 1;
> DROP RETENTION POLICY "auto_delete" ON "mydb";

# 选取第一条记录
> SELECT * FROM tab_monitor_cpu GROUP BY * ORDER BY DESC LIMIT 1
> SELECT * FROM tab_monitor_cpu GROUP BY * ORDER BY ASC  LIMIT 1


# 修改归档策略
> ALTER RETENTION POLICY autogen ON telegraf DURATION 0s REPLICATION 1 SHARD DURATION 168h0m0s default
> CREATE RETENTION POLICY autogen ON telegraf DURATION 0s REPLICATION 1 SHARD DURATION 168h0m0s default
> DROP RETENTION POLICY autogen ON telegraf 
```



### 5.InfluxDB 不支持 order by tag 

InfluxDB 仅支持 order by time ，有人 2015 年在 github 提了一个 [feature request issue2964](https://github.com/influxdata/influxdb/issues/3954) ，但现在都 2019 了，还没实现。
所以对此功能不抱过多希望。

同样的原因， influxdb 也无法支持 [order by count](https://stackoverflow.com/questions/40117511/how-to-do-group-by-with-count-and-ordering-by-count-in-influxdb) 。


- tag 是加了索引的字段
在 influxdb 中主要用于 where 条件过滤数据，官方文档可参考[influxdb 1.7 tag-key](https://docs.influxdata.com/influxdb/v1.7/concepts/glossary/#tag-key)
```sql
> show tag keys on sdb
name: ServerConf
tagKey
------
ConfId
StartTime
```

- 下面这种查询是失败的
因为 StartTime 是自定义的 tag 
```sql
> select StartTime, ConfId, TimeCost, ConfName, MemberCount from ServerConf  order by StartTime desc  limit 24 offset 0
ERR: error parsing query: only ORDER BY time supported at this time
> 
```

- 必须按 time 排序
```sql
> select StartTime, ConfId, TimeCost, ConfName, MemberCount from ServerConf  order by time desc  limit 24 offset 0
name: ServerConf
time                StartTime                 ConfId TimeCost ConfName MemberCount
----                ---------                 ------ -------- -------- -----------
1569756198304656259 2019-09-29T17:23:12+08:00 100008 7206     003的会议   1       
1569748457288307886 2019-09-29T16:57:34+08:00 100006 1003     003的会议   1       
1569747020536547155 2019-09-29T16:36:43+08:00 100004 817      003的会议   1       
1569745107424092383 2019-09-29T15:50:41+08:00 100005 1666     admin的会议 4       
1569743501334874861 2019-09-29T15:50:11+08:00 100002 90       003的会议   1       
```


[^InfluxdbGettingStart]: [InfluxDB 入门简介](https://docs.influxdata.com/influxdb/v1.6/introduction/getting-started/)



