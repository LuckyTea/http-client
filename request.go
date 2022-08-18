package httpclient

import (
	"time"

	"github.com/valyala/fasthttp"
)

// Do performs the given http request and fills the given http response.
func (c *Client) Do(req *fasthttp.Request, resp *fasthttp.Response) (err error) {
	start := time.Now()

	err = c.HTTP.Do(req, resp)

	if c.latencyMetric != nil {
		c.latencyMetric.WithLabelValues(c.domain).Observe(time.Since(start).Seconds())
	}

	return
}

// DoTimeout performs the given request and waits for response during
// the given timeout duration.
func (c *Client) DoTimeout(req *fasthttp.Request, resp *fasthttp.Response) (err error) {
	start := time.Now()

	err = c.HTTP.DoTimeout(req, resp, c.Timeout)

	if c.latencyMetric != nil {
		c.latencyMetric.WithLabelValues(c.domain).Observe(time.Since(start).Seconds())
	}

	return
}
