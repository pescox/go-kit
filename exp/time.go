package exp

import "time"

func TimeFromSeconds(v int64) time.Time {
	return time.Unix(v/1000, 0)
}

func TimeFromMilliseconds(v int64) time.Time {
	return time.Unix(0, v*int64(time.Millisecond))
}

func UnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
