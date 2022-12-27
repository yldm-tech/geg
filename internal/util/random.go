package util

import (
	"math/rand"
)

func RandomRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
