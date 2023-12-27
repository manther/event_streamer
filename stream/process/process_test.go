package process

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/manther/events/stream"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/stretchr/testify/require"
)

func TestPStream(t *testing.T) {
	pcManger := NewProcessManager(
		time.Second/2,
		5,
	)

	fBuild := func(*process.Process) stream.IEvent {

		return &ProcEvent{
			eventID:   uuid.NewString(),
			timeStamp: time.Now(),
		}
	}

	fListener := func() []*process.Process {
		return []*process.Process{{}}
	}

	i := 1
	for {
		event := <- pcManger.Stream(fBuild, fListener)
		fmt.Println("Entering", event.ToString(), i)
		if i == pcManger.Amount() {
			fmt.Println("Closing", event.ToString())
			break
		}
		fmt.Println("Exiting", event.ToString(), i)
		i++
	}
	require.Equal(t, i, pcManger.Amount())
}
