# http-client

Simple golang package with fasthttp client and prometheus metric.

## Example

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
client := httpclient.NewWithMetric("example", netSourcesLatencyHistogram)

// request
err := p.client.DoTimeout(req, resp)
if err != nil {
    // error handling
}
```
