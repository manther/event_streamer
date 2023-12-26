package stream

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Showing testability of a generic streaming function
func TestStream(t *testing.T) {
	manager := NewProcessManager(
		time.Second/10,
		10,
	)

	testListener := func() []interface{} {
		return []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	}

	testBuilder := func(interface{}) IEvent {
		return ProcEvent{
			timeStamp: time.Now(),
		}
	}

	events, quitchan := Stream(manager, testBuilder, testListener)
	i := 1
	timeNow := time.Now()
	for event := range events {
		require.WithinDuration(t, event.TimeStamp(), timeNow, time.Second/10 + 1*time.Second)
		if i == manager.Amount() {
			close(quitchan)
			return
		}
		i++
	}

}
