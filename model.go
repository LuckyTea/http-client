package httpclient

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
)

const (
	// DefaultRetry - количество ретраев под текущую реализацию переключения между основными api и failover.
	DefaultRetry = 2
	// DefaultTimeout - стандартное время ожидания ответа от сервера.
	DefaultTimeout = 3 * time.Second

	defaultBackoffMinTime = 150 * time.Millisecond
	defaultBackoffMaxTime = 2 * time.Second
)

// Client is a convenient API to make HTTP calls.
// Client also handles automatically retrying failed HTTP requests.
type Client struct {
	RetryMax      int              // Количество попыток запроса.
	Timeout       time.Duration    // Таймаут для http запроса.
	Backoff       *Backoff         // Рассчёт времени ожидания между попытками запроса.
	HTTP          *fasthttp.Client // HTTP клиент.
	latencyMetric *prometheus.HistogramVec
	domain        string
}

// Backoff - provide essention backoff parameters.
type Backoff struct {
	Min       time.Duration
	Max       time.Duration
	Calculate func(int) time.Duration
}
