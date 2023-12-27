package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manther/events/stream"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/stretchr/testify/require"
)

func TestMeMStream(t *testing.T) {
	memManger := NewMemoryManager(
		time.Second/2,
		5,
	)

	fBuild := func(*mem.VirtualMemoryStat) stream.IEvent {

		return &memEvent{
			eventID:           uuid.NewString(),
			timeStamp:         time.Now(),
			VirtualMemoryStat: &mem.VirtualMemoryStat{},
		}
	}

	fListener := func() (*mem.VirtualMemoryStat, error) {
		return &mem.VirtualMemoryStat{}, nil
	}

	events, quitchan := memManger.Stream(fBuild, fListener)
	i := 1

	testFunc := func() {
		for event := range events {
			fmt.Println("Entering", event.ToString())
			if i == memManger.Amount() {
				fmt.Println("Closing", event.ToString())
				close(quitchan)
			}
			i++
			fmt.Println("Exiting", event.ToString())
		}
	}
	require.NotPanics(t,testFunc)
}
