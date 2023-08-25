# http-client

Simple golang package with fasthttp client and prometheus metric.

## Example NewWithMetric

```go
// metric
var netSourcesLatencyHistogram = func() *prometheus.HistogramVec {
    var metric = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Namespace: "service",
            Subsystem: "net",
            Name:      "sources_latency",
            Help:      "Third party response latency histogram.",
            Buckets:   prometheus.ExponentialBuckets(0.05, 2, 8),
        }, []string{"source"})
    prometheus.MustRegister(metric)

    return metric
}()

// create
client := httpclient.NewWithMetric("domain", netSourcesLatencyHistogram)

// request
if err := p.client.DoTimeout(req, resp); err != nil {
    // error handling
}
```

## Example NewWithMetricFunc

```golang
// metric func
var latencyFunc = func(start time.Time, domain string) {
    latencyMetric.WithLabelValues(domain).Observe(float64(time.Since(start).Nanoseconds()) / 1000000)
}

// func
client := httpclient.NewWithMetricFunc("domain", latencyFunc)

// request
if err := p.client.DoTimeout(req, resp); err != nil {
    // error handling
}
```
