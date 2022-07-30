package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRetry(t *testing.T) {

	now := time.Now()
	retry := NewRetry(3, 2, func() error {
		fmt.Println("retry1")
		return nil
	})
	fmt.Println(retry.Duration())

	if err := retry.Run(); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("done!")
	}

	fmt.Println("time.Now()-now:", time.Since(now))
}
