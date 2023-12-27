package stream

import (
	"fmt"
	"time"
)

type IStreamManager interface {
	Amount() int
	Rate() time.Duration
}

type IEvent interface {
	ToString() string
	TimeStamp() time.Time
}

type EventTypes int

const (
	Memory EventTypes = iota + 1
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

// Attempt to make a streamer that could take any build / listen function
func Stream(manager IStreamManager, fBuilder func(any) IEvent, fListener func() []any) (<-chan IEvent, chan<- struct{}) {
	eventChan := make(chan IEvent)
	quitChan := make(chan struct{})
	go func() {
		for {
			for _, out := range fListener() {
				pcEvent := fBuilder(out)
				time.Sleep(manager.Rate())
				eventChan <- pcEvent

				select {
				case recQuit := <-quitChan:
					close(eventChan)
					fmt.Println("I recieved a quit request.", recQuit)
					return
				default:
				}
			}
		}

	}()

	return eventChan, quitChan
}
