---
layout: post
title:  "[译]Prometheus 指标类型"
date:   2021-04-24 23:23:00 +0800
tags:   tech
---

* category
{:toc}




# METRIC TYPES 指标类型 [^PROMETHEUS_METRIC_TYPES]

-   [Counter](https://prometheus.io/docs/concepts/metric_types/#counter)
-   [Gauge](https://prometheus.io/docs/concepts/metric_types/#gauge)
-   [Histogram](https://prometheus.io/docs/concepts/metric_types/#histogram)
-   [Summary](https://prometheus.io/docs/concepts/metric_types/#summary)

The Prometheus client libraries offer four core metric types. 
These are currently only differentiated in the client libraries (to enable APIs tailored to the usage of the specific types) and in the wire protocol. 
The Prometheus server does not yet make use of the type information and flattens all data into untyped time series. This may change in the future.


- Counter（计数器）
- Gauge（仪表盘）
- Histogram（直方图）
- Summary（摘要）

Prometheus 客户端库提供4种指标类型。
不过，目前仅在 Prometheus 客户端及网络协议区分了这几种类型（用于根据具体类型来定制API使用方式）。
实际上 Prometheus 的服务端并不区分指标类型，而是简单地把这些指标统一视为无类型的时序数据。
这个特性未来可能还会有变化。


## Counter 计数器

A  _counter_  is a cumulative metric that represents a single  [monotonically increasing counter](https://en.wikipedia.org/wiki/Monotonic_function)  whose value can only increase or be reset to zero on restart. For example, you can use a counter to represent the number of requests served, tasks completed, or errors.

指标数据单调递增的数据用 _计数器_ 表示，这类数据只会增加，不会减少（但会重置为零）。
比如，可以用计数器来表示请求数量，已经完成的任务数，错误发生的数量。

Do not use a counter to expose a value that can decrease. For example, do not use a counter for the number of currently running processes; instead use a gauge.

不要使用计数器记录可能会减少的数据。
比如，不要用计数器记录当前运行中的任务数量；这种数据请使用仪表盘表示。

Client library usage documentation for counters:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Counter)
-   [Java](https://github.com/prometheus/client_java#counter)
-   [Python](https://github.com/prometheus/client_python#counter)
-   [Ruby](https://github.com/prometheus/client_ruby#counter)

## Gauge 仪表盘

A  _gauge_  is a metric that represents a single numerical value that can arbitrarily go up and down.

指标数据可增可减的数据用 _仪表盘_ 表示。

Gauges are typically used for measured values like temperatures or current memory usage, but also "counts" that can go up and down, like the number of concurrent requests.

可以用仪表盘记录 当前温度，当前内存使用率，或者当前请求数量 这类"总数"会随时变化的数据。

Client library usage documentation for gauges:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Gauge)
-   [Java](https://github.com/prometheus/client_java#gauge)
-   [Python](https://github.com/prometheus/client_python#gauge)
-   [Ruby](https://github.com/prometheus/client_ruby#gauge)

## Histogram 直方图（柱状图）

A  _histogram_  samples observations (usually things like request durations or response sizes) and counts them in configurable buckets. It also provides a sum of all observed values.

_直方图_ 用于对指标数据（比如请求耗时或响应包大小）进行采样，可以按取值区间对不同指标数据的出现次数进行统计。它也会记录所有样本的总和。

> NOTE: 比如统计 "语文考试中，大于90分的学生有多少人，大于80分的有多少人，大于60分的有多少人" 时，就可以用直方图表示。

> NOTE: bucket 桶。每段不同的样本区间都可以理解为一个桶，当一个采样值落在某个样本区间时，可以认为这个样本值落到对就的桶里，此桶的计数就可以 +1 了。


A histogram with a base metric name of  `<basename>`  exposes multiple time series during a scrape:

-   cumulative counters for the observation buckets, exposed as  `<basename>_bucket{le="<upper inclusive bound>"}`
-   the  **total sum**  of all observed values, exposed as  `<basename>_sum`
-   the  **count**  of events that have been observed, exposed as  `<basename>_count`  (identical to  `<basename>_bucket{le="+Inf"}`  above)

使用直方图类型抓取数据时，会生成几种不同名称的时序数据。假设指标名称是 `<basename>` ：

- 在不同取值区间的样本统计结果，保存在  `<basename>_bucket{le="<上边界>"}` 中。le表示 lower equal , 用于筛选小于等于上边界的数据。
- 所有观测到的样本值的总和，保存在 `<basename>_sum` 中。
- 所有观测到的样本总个数，保存在 `<basename>_count` 中。


> NOTE: 假设最近一次统计周期（如15s）共请求 `/metrics` 接口3次，每次耗时分别为 4s 2s 3s ，那么
> 1. `prometheus_http_request_duration_seconds_bucket{handler="/metrics",le="0"}` 显示 0 条数据
> 2. `prometheus_http_request_duration_seconds_bucket{handler="/metrics",le="1"}` 显示 0 条数据
> 3. `prometheus_http_request_duration_seconds_bucket{handler="/metrics",le="2"}` 显示 {2s} 1 条数据
> 4. `prometheus_http_request_duration_seconds_bucket{handler="/metrics",le="4"}` 显示 {4s 2s 3s} 3 条数据
> 5. `prometheus_http_request_duration_seconds_bucket{handler="/metrics",le="6"}` 显示 {4s 2s 3s} 3 条数据
> 6. `prometheus_http_request_duration_seconds_sum{handler="/metrics"}` 显示值应该是 4+2+3 ＝ 9s
> 7. `prometheus_http_request_duration_seconds_count{handler="/metrics"}` 显示值应该是 1+1+1 ＝3次

Use the  [`histogram_quantile()`  function](https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile)  to calculate quantiles from histograms or even aggregations of histograms. 
A histogram is also suitable to calculate an  [Apdex score](https://en.wikipedia.org/wiki/Apdex). 
When operating on buckets, remember that the histogram is  [cumulative](https://en.wikipedia.org/wiki/Histogram#Cumulative_histogram). 
See  [histograms and summaries](https://prometheus.io/docs/practices/histograms)  for details of histogram usage and differences to  [summaries](https://prometheus.io/docs/concepts/metric_types/#summary).

使用 `histogram_quantile()`  函数可以从直方图（甚至直方图的聚合结果）中计算分位数。
直方图还能用来生成性能指数。
另外，要劳记，直方图在不同取值区间的统计结果是累积的。
查看 [histograms and summaries](https://prometheus.io/docs/practices/histograms)  了解 直方图 与 概要图的详细区别。


> NOTE: `histogram_quantile（φ float，b instant-vector）` 0 < φ < 1 ，比如 0.9 表示取 90% 分位数，即第 90 个百分位数。 

> NOTE: 以后面解释分数数的示例数据为例子，将样本划分为四个取值区间： {0~3, 3~5.5, 5.5~7, 7~9}
>
> - ordinary histogram 显示的样本统计结果
>   1. {0~3}  共3个样本；
>   2. {3~5.5}共2个样本；
>   3. {5.5~7}共3个样本；
>   4. {7~9}  共3个样本；
>
> - cumulative histogram 显示的样本统计结果
>   1. {0~3}  共3个样本；
>   2. {3~5.5}共5个样本；
>   3. {5.5~7}共8个样本；
>   4. {7~9}  共10个样本；

> NOTE: 性能指数(Apdex score) Apdex 是一个国际通用标准，是对用户体验满意度的量化值。响应时间(Response time) 

> NOTE: 分位数是什么意义? 可以参考
> 1. [四分位数 数学乐](https://www.shuxuele.com/data/quartiles.html)理解.
> 2. [第95个百分位（95th percentile）是什么概念？](https://www.zhihu.com/question/20575291)

```txt 
 把一组数据排序，然后分为四等份:
   {1, 3,  3,  4, 5,   6, 6, 7, 8, 8}
           |         |       | 
           |         |       | 
           V         V       V 
           Q1        Q2      Q3

 Q1: 第一个四分位数，也可以叫"第25个百分位"。这里是 3 。
 Q2: 第二个四分位数，也可以叫"第50个百分位"。这里是 (5+6) / 2 = 5.5 。
 Q3: 第三个四分位数，也可以叫"第75个百分位"。这里是 7 。

```

Client library usage documentation for histograms:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Histogram)
-   [Java](https://github.com/prometheus/client_java#histogram)
-   [Python](https://github.com/prometheus/client_python#histogram)
-   [Ruby](https://github.com/prometheus/client_ruby#histogram)

## Summary 概要图

Similar to a  _histogram_, a  _summary_  samples observations (usually things like request durations and response sizes). While it also provides a total count of observations and a sum of all observed values, it calculates configurable quantiles over a sliding time window.

与 直方图 类型， _概要图_ 用于对指标数据（比如请求耗时或响应包大小）进行采样，但它能直接按分位数对指标数据的出现次数和样本值总和进行统计。

A summary with a base metric name of  `<basename>`  exposes multiple time series during a scrape:

-   streaming  **φ-quantiles**  (0 ≤ φ ≤ 1) of observed events, exposed as  `<basename>{quantile="<φ>"}`
-   the  **total sum**  of all observed values, exposed as  `<basename>_sum`
-   the  **count**  of events that have been observed, exposed as  `<basename>_count`

使用概要图类型抓取数据时，也会生成几种不同名称的时序数据。假设指标名称是 `<basename>` ：

- 从 `<basename>{quantile="<φ>"}` 中获取特定分位数，**φ-quantiles** 取值范围是 0 ≤ φ ≤ 1
- 所有观测到的样本值的总和，保存在 `<basename>_sum` 中。
- 所有观测到的样本总个数，保存在 `<basename>_count` 中。


See  [histograms and summaries](https://prometheus.io/docs/practices/histograms)  for detailed explanations of φ-quantiles, summary usage, and differences to  [histograms](https://prometheus.io/docs/concepts/metric_types/#histogram).

Client library usage documentation for summaries:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Summary)
-   [Java](https://github.com/prometheus/client_java#summary)
-   [Python](https://github.com/prometheus/client_python#summary)
-   [Ruby](https://github.com/prometheus/client_ruby#summary)



# 笔记


- 如何要得到网络上传下载速率？

  被监控的程序(exporter 端)只要记录"当前值"，即 exporter 端只要记录总上传下载字节数。
  经过 prometheus server 的定时抓取后，就能利用各种查询语法得到 不同时间和精度的 速率；

  不需要在 exporter 中计算每秒的速率，再主动上报。

- 如何分析接口平均耗时如何？耗时5s的接口总共有多少？耗时低于100的接口又有多少？

  这就是直方图的用武之地了。
  在 exporter 中还是只需要记录"当前值"，
  但是要利用相关 client library 计算出不同取值区间的"当前值"。



- 统计网络速率

[prometheus querying basics](https://prometheus.io/docs/prometheus/latest/querying/basics/)

```proql

rate(node_netstat_IpExt_InOctets{}[5m])

rate(windows_net_bytes_received_total{}[5m])

rate(node_network_receive_bytes_total{}[5m])

# 5m 表示对时序数据分成 5m 区间的多段数据，然后用 rate() 计算每秒钟的网络下载速率

```

- 程序启动命令

```sh

 ./prometheus --config.file=./prometheus.yml
./node_exporter --collector.netstat.fields="^(.*_(InErrors|InErrs)|Ip_Forwarding|Ip(6|Ext)_(InOctets|OutOctets)|Icmp6?_(InMsgs|OutMsgs)|TcpExt_(Listen.*|Syncookies.*|TCPSynRetrans)|Tcp_(ActiveOpens|InSegs|OutSegs|OutRsts|PassiveOpens|RetransSegs|CurrEstab)|Udp6?_(InDatagrams|OutDatagrams|NoPorts|RcvbufErrors|SndbufErrors))$" --web.listen-address 0.0.0.0:8082
./windows_exporter-0.16.0-amd64.exe --log.level="debug"


https://github.com/prometheus/prometheus

https://github.com/prometheus/node_exporter

https://github.com/prometheus-community/windows_exporter
```

- 配置文件

```sh
# my global config
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
  - static_configs:
    - targets:
      # - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'prometheus'
    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.
    static_configs:
    - targets: ['localhost:9090']

  - job_name:       'node_windows'
    scrape_interval: 5s
    static_configs:
    - targets: ['192.168.0.18:9182']

  - job_name:       'node_ubuntu'
    scrape_interval: 5s
    static_configs:
      - targets: ['localhost:8082']
        labels:
          group: 'canary'

  - job_name:       'node_mac'
    scrape_interval: 5s
    static_configs:
    - targets: ['localhost:9100']

# "prometheus.yml" 39 行 --38%--  
```

[^PROMETHEUS_METRIC_TYPES]: [Prometheus Metric Types](https://prometheus.io/docs/concepts/metric_types/#metric-types)
