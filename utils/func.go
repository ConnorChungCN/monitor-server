package utils

import (
	"fmt"
	"time"
)

func Retry(f func() error, retries int) error {
	var lastErr error
	for i := 0; i < retries; i++ {
		if err := f(); err != nil {
			lastErr = err
			time.Sleep(time.Second * 1)
			continue
		} else {
			return nil
		}
	}
	return fmt.Errorf("exceeded maximum retry count, %w", lastErr)
}
