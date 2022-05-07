package types

type Basic interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | string | bool | complex64 | complex128
}

func ZeroValue[T Basic](t T) T {
	var zero T
	return zero
}
