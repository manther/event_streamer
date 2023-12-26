package stream

import (
	"time"
)

// Attempt to make a streamer that could take any build / listen function
func Stream(manager IManager, fBuilder func(any) IEvent, fListener func() []any) (<-chan IEvent, chan<- struct{}) {
	eventChan := make(chan IEvent)
	quitChan := make(chan struct{})
	go func() {
		for {
			for _, out := range fListener() {
				pcEvent := fBuilder(out)
				time.Sleep(manager.Rate())
				eventChan <- pcEvent

				select {
				case <-quitChan:
					close(eventChan)
					return
				default:
				}
			}
		}

	}()

	return eventChan, quitChan
}
