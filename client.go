// Package httpclient provides easy way to crate http client with retry and timeout.
package httpclient

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
)

// New - return new http client with given params & backoff.
func New(retry int, timeout time.Duration) Client {
	return Client{
		RetryMax: retry,
		Timeout:  timeout,
		Backoff:  NewBackoff(defaultBackoffMinTime, defaultBackoffMaxTime),
		HTTP: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: true,
		},
	}
}

// NewDefault - return new http client with default params & backoff.
func NewDefault() Client {
	return Client{
		RetryMax: DefaultRetry,
		Timeout:  DefaultTimeout,
		Backoff:  NewBackoff(defaultBackoffMinTime, defaultBackoffMaxTime),
		HTTP: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: true,
		},
	}
}

// NewWithMetric - return new http client with default params & backoff.
func NewWithMetric(domain string, latencyMetric *prometheus.HistogramVec) Client {
	client := Client{
		RetryMax: DefaultRetry,
		Timeout:  DefaultTimeout,
		Backoff:  NewBackoff(defaultBackoffMinTime, defaultBackoffMaxTime),
		domain:   domain,
		HTTP: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: true,
		},
	}

	if latencyMetric != nil {
		client.latencyFunc = func(start time.Time, domain string) {
			latencyMetric.WithLabelValues(domain).Observe(time.Since(start).Seconds())
		}
	}

	return client
}

// NewWithMetricFunc - return new http client with default params & backoff.
func NewWithMetricFunc(domain string, latencyFunc func(start time.Time, domain string)) Client {
	return Client{
		RetryMax:    DefaultRetry,
		Timeout:     DefaultTimeout,
		Backoff:     NewBackoff(defaultBackoffMinTime, defaultBackoffMaxTime),
		latencyFunc: latencyFunc,
		domain:      domain,
		HTTP: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: true,
		},
	}
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}
