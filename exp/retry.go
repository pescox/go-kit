package exp

func RetryOnFalse(times int, f func() bool) {
	for i := 0; i < times; i++ {
		if f() {
			return
		}
	}
}

func RetryOnError(times int, f func() error) {
	for i := 0; i < times; i++ {
		if err := f(); err == nil {
			return
		}
	}
}
