package task

import (
	"time"
)

type intervalOption struct {
	interval       time.Duration
	maxConcurrency int
}

func newIntervalOption(interval time.Duration, maxConcurrency int) *intervalOption {
	return &intervalOption{
		interval:       interval,
		maxConcurrency: maxConcurrency,
	}
}

func (o *intervalOption) Interval() time.Duration {
	return o.interval
}

func (o *intervalOption) MaxConcurrency() int {
	if o.maxConcurrency <= 0 {
		return 1
	}
	return o.maxConcurrency
}

func DoInterval(opt intervalOption, f func()) {
	timer := time.NewTicker(opt.Interval())
	defer timer.Stop()
	next := make(chan struct{}, opt.MaxConcurrency())
	for {
		select {
		case <-timer.C:
			next <- struct{}{}
			go func() {
				f()
				<-next
			}()
		}
	}
}
