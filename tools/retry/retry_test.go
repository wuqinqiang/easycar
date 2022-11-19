package retry

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRetry(t *testing.T) {
	retry := New(3, WithMaxBackOffTime(1*time.Second), WithFactor(3))

	assert.Equal(t, time.Second, retry.Duration())
	assert.Equal(t, uint32(3), retry.allowAttempt)
	assert.Equal(t, uint32(3), retry.factor)
	assert.Equal(t, uint32(1), retry.currentAttempt)
	err := retry.Run(func() error {
		return nil
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, time.Second, retry.Duration())
	assert.Equal(t, uint32(2), retry.currentAttempt)
}

func TestRetry_Run(t *testing.T) {
	retry := New(3, WithMaxBackOffTime(1*time.Second))
	err := retry.Run(func() error {
		return fmt.Errorf("fn err")
	})
	assert.Equal(t, err.Error(), "fn err")

	expectedVal := 0
	retry = New(3, WithMaxBackOffTime(100*time.Millisecond))
	err = retry.Run(func() error {
		expectedVal++
		if expectedVal < 3 {
			return fmt.Errorf("condition not met")
		}
		return nil
	})
	assert.Equal(t, err, nil)
	assert.Equal(t, uint32(2), retry.currentAttempt)
	assert.Equal(t, 100*time.Millisecond, retry.maxBackOffTime)

}
