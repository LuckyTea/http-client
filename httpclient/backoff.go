package httpclient

import (
	"math"
	"time"
)

// NewBackoff - return function for new backoff calculation.
func NewBackoff(min, max time.Duration) *Backoff {
	b := Backoff{
		Min: min,
		Max: max,
	}

	b.Calculace = b.defaultBackoff

	return &b
}

// DefaultBackoff provides a default callback for Client.Backoff which
// will perform exponential backoff based on the attempt number and limited
// by the provided minimum and maximum durations.
func (b *Backoff) defaultBackoff(attemptNum int) time.Duration {
	mult := math.Pow(2, float64(attemptNum)) * float64(b.Min) //nolint: gomnd
	sleep := time.Duration(mult)

	if float64(sleep) != mult || sleep > b.Max {
		sleep = b.Max
	}

	return sleep
}
