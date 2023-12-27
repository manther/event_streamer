package stream

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type manager struct {
	rate   time.Duration
	amount int
}

func (m manager) Amount() int {
	return m.amount
}

func (m manager) Rate() time.Duration {
	return m.rate
}

type testEvent struct {
	eventId   string
	timeStamp time.Time
}

func (e testEvent) TimeStamp() time.Time {
	return e.timeStamp
}

func (e testEvent) ToString() string {
	return fmt.Sprintf("ProcStats: %s, %v", e.timeStamp, e.eventId)
}

// Showing testability of a generic streaming function
// Although the function itself didn't turn out to be 
// very usable. May circle back to it later.
// Also experimenting with close channels.
// This close channel creates a purposful race condition.
// The way the incrementor is put in an else statement 
// creates a continuos loop of attempting to close the chan.
// A defer / recover func allows to recover from the race
// condition.
func TestStream(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("I recovered from an exception.", err)
		}
	}()

	testMangager := manager{
		time.Second / 100,
		10,
	}

	testListener := func() []interface{} {
		return []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	}

	testBuilder := func(interface{}) IEvent {
		return testEvent{
			eventId:   uuid.NewString(),
			timeStamp: time.Now(),
		}
	}

	events, quitchan := Stream(testMangager, testBuilder, testListener)
	i := 1
	timeNow := time.Now()
	for event := range events {
		fmt.Println("Entering", event.ToString())
		require.WithinDuration(t, event.TimeStamp(), timeNow, time.Second/10+1*time.Second)
		if i == testMangager.Amount() {
			fmt.Println("Closing", event.ToString())
			close(quitchan)
		} else {
			i++
		}
		fmt.Println("Exiting", event.ToString())
	}
}
