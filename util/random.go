package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnop"

var r = *rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

func RandomString() string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < k; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// func RandomEventType() string {
// 	return eventTypes[RandomInt(0, 2)]
// }

