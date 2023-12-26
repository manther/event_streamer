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

type EventTypes int

const (
	Memory    EventTypes = iota + 1
	Process                       
	Network                                        
)

func (e EventTypes) String() string {
	return []string{"Memory", "Process", "Network"}[e-1]
}

func (e EventTypes) EnumIndex() int {
	return int(e)
}

var eventTypes = []string{Memory.String(), Process.String(), Network.String()}

func RandomEventType() string {
	return eventTypes[RandomInt(0, 2)]
}

