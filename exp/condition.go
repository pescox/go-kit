package exp

import (
	"github.com/pescox/go-kit/types"
)

func If[T types.Basic](b bool, v1, v2 T) T {
	if b {
		return v1
	}
	return v2
}

func ValueOrDefault[T types.Basic](v T, def T) T {
	if v == types.ZeroValue(v) {
		return def
	}
	return v
}
