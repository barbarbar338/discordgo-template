package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(a []string) string {
	idx := rand.Intn(len(a) - 1)
	elem := a[idx]
	return elem
}
