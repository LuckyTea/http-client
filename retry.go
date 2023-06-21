package httpclient

import "time"

// CheckRetry will return a boolean value depending on whether the attempts ended or not.
func (c *Client) CheckRetry(try int) bool {
	return c.RetryMax-try <= 0
}

// CheckRetryWithBackoff will return a boolean value depending on whether the attempts ended or not.
// If the attempts are over - the value will be returned immediately,
// if not - the value will be returned after the waiting time set via the backoff function.
func (c *Client) CheckRetryWithBackoff(try int) bool {
	if c.RetryMax-try <= 0 {
		return false
	}

	time.Sleep(c.Backoff.Calculate(try))

	return true
}
