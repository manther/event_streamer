package stream

import (
	"testing"
	"time"

	"fmt"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/stretchr/testify/require"
)

func TestPStream(t *testing.T) {
	pcManger := NewProcessManager(
		time.Second/10,
		10,
	)

	fBuild := func(*process.Process) IEvent {
		return &ProcEvent{}
	}

	fListener := func() []*process.Process {
		return []*process.Process{{}}
	}

	events, quitchan := pcManger.Stream(fBuild, fListener)
	i := 1
	for event := range events {
		fmt.Println(event.ToString())
		if i == pcManger.Amount() {
			close(quitchan)
			return
		}
		i++
	}
	require.Equal(t, i, 10)
}