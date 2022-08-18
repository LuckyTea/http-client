package httpclient

import (
	"context"
	"time"

	"github.com/valyala/fasthttp"
)

// Do performs the given http request and fills the given http response.
func (c *Client) Do(req *fasthttp.Request, resp *fasthttp.Response) error {
	if c.latencyMetric != nil {
		defer func(start time.Time) {
			c.latencyMetric.WithLabelValues(c.domain).Observe(time.Since(start).Seconds())
		}(time.Now())
	}

	return c.HTTP.Do(req, resp)
}

// DoTimeout performs the given request and waits for response during
// the given timeout duration.
func (c *Client) DoTimeout(req *fasthttp.Request, resp *fasthttp.Response) error {
	if c.latencyMetric != nil {
		defer func(start time.Time) {
			c.latencyMetric.WithLabelValues(c.domain).Observe(time.Since(start).Seconds())
		}(time.Now())
	}

	return c.HTTP.DoTimeout(req, resp, c.Timeout)
}

// DoContext perform the given request with ctx deadline or waits for response during
// the given timeout duration.
func (c *Client) DoContext(ctx context.Context, req *fasthttp.Request, resp *fasthttp.Response) error {
	if c.latencyMetric != nil {
		defer func(start time.Time) {
			c.latencyMetric.WithLabelValues(c.domain).Observe(time.Since(start).Seconds())
		}(time.Now())
	}

	deadline, ok := ctx.Deadline()

	if ok {
		return c.HTTP.DoDeadline(req, resp, deadline)
	} else {
		return c.HTTP.DoTimeout(req, resp, c.Timeout)
	}
}
