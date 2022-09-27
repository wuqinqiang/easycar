package retry

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRetry(t *testing.T) {
	retry := New(3, WithMaxBackOffTime(1*time.Second))

	assert.Equal(t, time.Second, retry.Duration())
	assert.Equal(t, uint32(3), retry.allowAttempt)
	assert.Equal(t, uint32(2), retry.factor)
	assert.Equal(t, uint32(0), retry.currentAttempt)
	err := retry.Run(func() error {
		return nil
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, time.Second, retry.Duration())
	assert.Equal(t, uint32(1), retry.currentAttempt)
}
