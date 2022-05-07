package task

import (
	"time"

	"github.com/pescox/go-kit/log"
)

func DoAfter(d time.Duration, f func()) {
	timer := time.AfterFunc(d, f)
	defer timer.Stop()

	select {
	case <-timer.C:
	}
}

func DoTimeout(d time.Duration, f func()) {
	timeout := time.NewTimer(d)
	defer timeout.Stop()
	done := make(chan struct{}, 1)
	go func() {
		f()
		done <- struct{}{}
	}()

	select {
	case <-timeout.C:
		log.Info("timeout")
		return
	case <-done:
		close(done)
		return
	}
}
