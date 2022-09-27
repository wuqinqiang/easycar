package retry

import "time"

type RetryFn func(*Retry)

func WithFactor(factor uint32) RetryFn {
	return func(retry *Retry) {
		retry.factor = factor
	}
}
func WithMaxBackOffTime(maxBackOffTime time.Duration) RetryFn {
	return func(retry *Retry) {
		retry.maxBackOffTime = maxBackOffTime
	}
}
