package retry

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var (
	// DefaultRetries everything only three times when you on china
	DefaultRetries uint32 = 3
	DefaultFactor  uint32 = 2
	// MaxWaitBackOff 10 minutes
	MaxWaitBackOff = time.Second * 10 * 60
	// ErrMaxAttempt is returned when the max number of attempts has been reached
	ErrMaxAttempt = fmt.Errorf("max attempt")
)

type (
	Fn    func() error
	Retry struct {
		//min, max       time.Duration
		factor         uint32
		allowAttempt   uint32
		currentAttempt uint32
		timer          *time.Timer
		fn             Fn
		AfterFn        func() // AfterFn is called after the last attempt
	}
)

func NewRetry(allowRetries uint32, factor uint32, fn Fn) *Retry {
	if allowRetries == 0 {
		allowRetries = DefaultRetries
	}
	if factor == 0 {
		factor = DefaultFactor
	}
	b := &Retry{
		allowAttempt: allowRetries,
		factor:       factor,
		fn:           fn,
		timer:        time.NewTimer(0),
	}
	return b
}
func (b *Retry) Duration() time.Duration {
	backDuration := time.Duration(math.Pow(float64(b.factor),
		float64(b.currentAttempt))) * time.Second
	if backDuration > MaxWaitBackOff {
		backDuration = MaxWaitBackOff
	}
	return backDuration
}

func (b *Retry) Run() error {
	atomic.AddUint32(&b.currentAttempt, 1)
	if b.currentAttempt > b.allowAttempt {
		return ErrMaxAttempt
	}
	<-b.timer.C
	err := b.fn()
	if err == nil {
		b.timer.Stop()
		return nil
	}
	b.timer.Reset(b.Duration())

	if b.currentAttempt == b.allowAttempt {
		return err
	}
	return b.Run()
}
