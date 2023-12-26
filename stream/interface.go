package stream

import "time"

type IManager interface {
	Amount() int
	Rate() time.Duration
}

type IListener interface {
	
}

type IEvent interface {
	ToString() string
	TimeStamp() time.Time
}
