package tools

import "time"

func DoWithRetry(f func() bool, retry int, interval time.Duration) (int, bool) {
	retried := 0
	for i := 0; i < retry; i++ {
		if f() {
			return retried, true
		}
		retried++
		time.Sleep(interval)
	}
	return retried, false
}
