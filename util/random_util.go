package util

import (
	"math/rand"
	"time"
)

func RandomInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
