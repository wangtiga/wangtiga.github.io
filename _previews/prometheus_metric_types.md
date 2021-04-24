

# METRIC TYPES [^PROMETHEUS_METRIC_TYPES]

-   [Counter](https://prometheus.io/docs/concepts/metric_types/#counter)
-   [Gauge](https://prometheus.io/docs/concepts/metric_types/#gauge)
-   [Histogram](https://prometheus.io/docs/concepts/metric_types/#histogram)
-   [Summary](https://prometheus.io/docs/concepts/metric_types/#summary)

The Prometheus client libraries offer four core metric types. These are currently only differentiated in the client libraries (to enable APIs tailored to the usage of the specific types) and in the wire protocol. The Prometheus server does not yet make use of the type information and flattens all data into untyped time series. This may change in the future.

## Counter[](https://prometheus.io/docs/concepts/metric_types/#counter)

A  _counter_  is a cumulative metric that represents a single  [monotonically increasing counter](https://en.wikipedia.org/wiki/Monotonic_function)  whose value can only increase or be reset to zero on restart. For example, you can use a counter to represent the number of requests served, tasks completed, or errors.

Do not use a counter to expose a value that can decrease. For example, do not use a counter for the number of currently running processes; instead use a gauge.

Client library usage documentation for counters:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Counter)
-   [Java](https://github.com/prometheus/client_java#counter)
-   [Python](https://github.com/prometheus/client_python#counter)
-   [Ruby](https://github.com/prometheus/client_ruby#counter)

## Gauge[](https://prometheus.io/docs/concepts/metric_types/#gauge)

A  _gauge_  is a metric that represents a single numerical value that can arbitrarily go up and down.

Gauges are typically used for measured values like temperatures or current memory usage, but also "counts" that can go up and down, like the number of concurrent requests.

Client library usage documentation for gauges:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Gauge)
-   [Java](https://github.com/prometheus/client_java#gauge)
-   [Python](https://github.com/prometheus/client_python#gauge)
-   [Ruby](https://github.com/prometheus/client_ruby#gauge)

## Histogram[](https://prometheus.io/docs/concepts/metric_types/#histogram)

A  _histogram_  samples observations (usually things like request durations or response sizes) and counts them in configurable buckets. It also provides a sum of all observed values.

A histogram with a base metric name of  `<basename>`  exposes multiple time series during a scrape:

-   cumulative counters for the observation buckets, exposed as  `<basename>_bucket{le="<upper inclusive bound>"}`
-   the  **total sum**  of all observed values, exposed as  `<basename>_sum`
-   the  **count**  of events that have been observed, exposed as  `<basename>_count`  (identical to  `<basename>_bucket{le="+Inf"}`  above)

Use the  [`histogram_quantile()`  function](https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile)  to calculate quantiles from histograms or even aggregations of histograms. A histogram is also suitable to calculate an  [Apdex score](https://en.wikipedia.org/wiki/Apdex). When operating on buckets, remember that the histogram is  [cumulative](https://en.wikipedia.org/wiki/Histogram#Cumulative_histogram). See  [histograms and summaries](https://prometheus.io/docs/practices/histograms)  for details of histogram usage and differences to  [summaries](https://prometheus.io/docs/concepts/metric_types/#summary).

Client library usage documentation for histograms:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Histogram)
-   [Java](https://github.com/prometheus/client_java#histogram)
-   [Python](https://github.com/prometheus/client_python#histogram)
-   [Ruby](https://github.com/prometheus/client_ruby#histogram)

## Summary[](https://prometheus.io/docs/concepts/metric_types/#summary)

Similar to a  _histogram_, a  _summary_  samples observations (usually things like request durations and response sizes). While it also provides a total count of observations and a sum of all observed values, it calculates configurable quantiles over a sliding time window.

A summary with a base metric name of  `<basename>`  exposes multiple time series during a scrape:

-   streaming  **φ-quantiles**  (0 ≤ φ ≤ 1) of observed events, exposed as  `<basename>{quantile="<φ>"}`
-   the  **total sum**  of all observed values, exposed as  `<basename>_sum`
-   the  **count**  of events that have been observed, exposed as  `<basename>_count`

See  [histograms and summaries](https://prometheus.io/docs/practices/histograms)  for detailed explanations of φ-quantiles, summary usage, and differences to  [histograms](https://prometheus.io/docs/concepts/metric_types/#histogram).

Client library usage documentation for summaries:

-   [Go](https://godoc.org/github.com/prometheus/client_golang/prometheus#Summary)
-   [Java](https://github.com/prometheus/client_java#summary)
-   [Python](https://github.com/prometheus/client_python#summary)
-   [Ruby](https://github.com/prometheus/client_ruby#summary)


# https://prometheus.io/docs/prometheus/latest/querying/basics/

# Note

```proql

rate(node_netstat_IpExt_InOctets{}[5m])

rate(windows_net_bytes_received_total{}[5m])

# TODO 5m 具体含义

```

```sh

 ./prometheus --config.file=./prometheus.yml
./node_exporter --collector.netstat.fields="^(.*_(InErrors|InErrs)|Ip_Forwarding|Ip(6|Ext)_(InOctets|OutOctets)|Icmp6?_(InMsgs|OutMsgs)|TcpExt_(Listen.*|Syncookies.*|TCPSynRetrans)|Tcp_(ActiveOpens
|InSegs|OutSegs|OutRsts|PassiveOpens|RetransSegs|CurrEstab)|Udp6?_(InDatagrams|OutDatagrams|NoPorts|RcvbufErrors|SndbufErrors))$" --web.listen-address 0.0.0.0:8082
./windows_exporter-0.16.0-amd64.exe --log.level="debug"


https://github.com/prometheus/prometheus

https://github.com/prometheus/node_exporter

https://github.com/prometheus-community/windows_exporter
```

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
                                                                                                                                                                                                                                                                           
# "prometheus.yml" 39 行 --38%--  
```

[^PROMETHEUS_METRIC_TYPES]: [Prometheus Metric Types](https://prometheus.io/docs/concepts/metric_types/#metric-types)
