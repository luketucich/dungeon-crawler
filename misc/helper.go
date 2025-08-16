package misc

import (
	"math"
	"math/rand"
	"time"
)

func ShuffleSlice[T any](s []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func Sqrt(num int) int {
	s := int(math.Sqrt(float64(num)))
	if s < 1 {
		s = 1
	}
	return s
}
