# http-client

Simple golang module with fasthttp client and prometheus metric.

## Example

```go
// metric
var NetSourcesLatencyHistogram = func() *prometheus.HistogramVec {
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
client := httpclient.NewWithMetric("example", metrics.NetSourcesLatencyHistogram)

// request
err := p.client.DoTimeout(req, resp)
if err != nil {
    // error handling
}
```
