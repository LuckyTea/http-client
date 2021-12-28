package httpclient

import (
	"time"

	"github.com/valyala/fasthttp"
)

func (c *Client) Do(req *fasthttp.Request, resp *fasthttp.Response) (err error) {
	start := time.Now()

	err = c.HTTP.Do(req, resp)

	c.latencyMetric.WithLabelValues(c.domain).Observe(time.Since(start).Seconds())

	return
}

func (c *Client) DoTimeout(req *fasthttp.Request, resp *fasthttp.Response) (err error) {
	start := time.Now()

	err = c.HTTP.DoTimeout(req, resp, c.Timeout)

	c.latencyMetric.WithLabelValues(c.domain).Observe(time.Since(start).Seconds())

	return
}
