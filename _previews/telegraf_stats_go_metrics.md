---
layout: post
title:  "telegraf stats go metric"
date:   2020-02-08 18:00:00 +0800
tags:   tech
---

* category
{:toc}




# Metric registries [^JavaMetric]:
The five metric types: Gauges, Counters, Histograms, Meters, and Timers.


## Gauges

A gauge is the simplest metric type. It just returns a  _value_. If, for example, your application has a value which is maintained by a third-party library, you can easily expose it by registering a  `Gauge`  instance which returns that value:

registry.register(name(SessionStore.class, "cache-evictions"), new Gauge<Integer>() {
    @Override
    public Integer getValue() {
        return cache.getEvictionsCount();
    }
});

This will create a new gauge named  `com.example.proj.auth.SessionStore.cache-evictions`  which will return the number of evictions from the cache.


## Counters

A counter is a simple incrementing and decrementing 64-bit integer:

final Counter evictions = registry.counter(name(SessionStore.class, "cache-evictions"));
evictions.inc();
evictions.inc(3);
evictions.dec();
evictions.dec(2);

All  `Counter`  metrics start out at 0.


## Histograms

A  `Histogram`  measures the distribution of values in a stream of data: e.g., the number of results returned by a search:

final Histogram resultCounts = registry.histogram(name(ProductDAO.class, "result-counts");
resultCounts.update(results.size());

`Histogram`  metrics allow you to measure not just easy things like the min, mean, max, and standard deviation of values, but also  [quantiles](http://en.wikipedia.org/wiki/Quantile)  like the median or 95th percentile.

Traditionally, the way the median (or any other quantile) is calculated is to take the entire data set, sort it, and take the value in the middle (or 1% from the end, for the 99th percentile). This works for small data sets, or batch processing systems, but not for high-throughput, low-latency services.

The solution for this is to sample the data as it goes through. By maintaining a small, manageable reservoir which is statistically representative of the data stream as a whole, we can quickly and easily calculate quantiles which are valid approximations of the actual quantiles. This technique is called  **reservoir sampling**.

Metrics provides a number of different  `Reservoir`  implementations, each of which is useful.


## Meters

A meter measures the  _rate_  at which a set of events occur:

final Meter getRequests = registry.meter(name(WebProxy.class, "get-requests", "requests"));
getRequests.mark();
getRequests.mark(requests.size());

Meters measure the rate of the events in a few different ways. The  _mean_  rate is the average rate of events. It’s generally useful for trivia, but as it represents the total rate for your application’s entire lifetime (e.g., the total number of requests handled, divided by the number of seconds the process has been running), it doesn’t offer a sense of recency. Luckily, meters also record three different  _exponentially-weighted moving average_  rates: the 1-, 5-, and 15-minute moving averages.


## Timers

A timer is basically a  [histogram](https://metrics.dropwizard.io/4.1.1/manual/core.html#man-core-histograms)  of the duration of a type of event and a  [meter](https://metrics.dropwizard.io/4.1.1/manual/core.html#man-core-meters)  of the rate of its occurrence.

final Timer timer = registry.timer(name(WebProxy.class, "get-requests"));

final Timer.Context context = timer.time();
try {
    // handle request
} finally {
    context.stop();
}

Note

Elapsed times for it events are measured internally in nanoseconds, using Java’s high-precision  `System.nanoTime()`  method. Its precision and accuracy vary depending on operating system and hardware.


## Metric Sets

Metrics can also be grouped together into reusable metric sets using the  `MetricSet`  interface. This allows library authors to provide a single entry point for the instrumentation of a wide variety of functionality.



## 说明

- 设计监控数据的表格结构与设计业务数据表结构不同

  监控数据更多用于统计分析，所以每种数据放一个独立的表即可。

- 善于使用业界通用的库

  通用库归纳了一些常用统计方法，避免重复踩坑。

  比如， 
  统计并发量，只要 recv 时记录数量即可，短期内的数据可以聚合。
  统计接口延迟，只要取样采集数据即可，记录所有请求的耗时意义不大。



The five metric types: Gauges, Counters, Histograms, Meters, and Timers.

#### Guages
A gauge is the simplest metric type. It just returns a value. If, for example, your application has a value which is maintained by a third-party library, you can easily expose it by registering a Gauge instance which returns that value
:

#### Counters
A counter is a simple incrementing and decrementing 64-bit integer:

#### Histograms

`Histogram`  metrics allow you to measure not just easy things like the min, mean, max, and standard deviation of values, but also  [quantiles](http://en.wikipedia.org/wiki/Quantile)  like the median or 95th percentile.

The solution for this is to sample the data as it goes through. By maintaining a small, manageable reservoir which is statistically representative of the data stream as a whole, we can quickly and easily calculate quantiles which are v
alid approximations of the actual quantiles. This technique is called  **reservoir sampling**.

#### Metric Sets

Metrics can also be grouped together into reusable metric sets using the  `MetricSet`  interface. This allows library authors to provide a single entry point for the instrumentation of a wide variety of functionality.


### TelegrafStatsdSink

SetGaugeWithLabels 直接记录统计的数据。

IncrCounterWithLabels 会自动缓存并聚合数据。通过 inc() 或 dec() 实现

AddSampleWithLabels() 中什么意思？与 Histograms 类似 ？



### 1.接口并发请求量

IncrCounterWithLabels

### 2.接口耗时

AddSampleWithLabels

### 3.ERROR信息，启动停止事件

SetGaugeWithLabels

### 4.成员数量，时长

SetGaugeWithLabels




[^JavaMetric]:[Java Metric lib 核心概念](https://metrics.dropwizard.io/4.1.1/manual/core.html)

