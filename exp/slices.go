package exp

import (
	"math/rand"
	"time"

	"github.com/pescox/go-kit/types"
)

func Contains[T types.Basic](slices []T, e T) bool {
	for _, a := range slices {
		if a == e {
			return true
		}
	}
	return false
}

func Distinct[T types.Basic](slices []T) []T {
	var keys map[T]bool
	for _, e := range slices {
		keys[e] = true
	}
	result := make([]T, len(keys))
	for k := range keys {
		result = append(result, k)
	}
	return result
}

func Shuffle[T types.Basic](s []T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(s) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse[T types.Basic](s []T) {
	l := len(s)
	for i := l/2 - 1; i >= 0; i-- {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
