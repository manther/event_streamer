package process

import (
	"fmt"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/manther/events/stream"
	"github.com/shirou/gopsutil/v3/process"
)

// Process event structure that will take the state of gopsutil array elements
type ProcEvent struct {
	eventType       string
	eventID         string
	pid             string
	procUtilization string
	timeStamp       time.Time
}

func (e ProcEvent) TimeStamp() time.Time {
	return e.timeStamp
}

// The work that in the end will be shown for streaming information.
// in a real app this info would be used to some purpose. 
func (e ProcEvent) ToString() string {
	return fmt.Sprintf("ProcStats - TimeStamp: %v. Pid: %s. Event ID: %s. Event Type: %s. CPU: %%:%s", e.TimeStamp(), e.pid, e.eventID, e.eventType, e.procUtilization)
}

type ProcessManager struct {
	rate   time.Duration
	amount int
}

func (m ProcessManager) Amount() int {
	return m.amount
}

func (m ProcessManager) Rate() time.Duration {
	return m.rate
}

func NewProcessManager(rate time.Duration, amount int) ProcessManager {
	return ProcessManager{
		rate:   rate,
		amount: amount,
	}
}

func (m ProcessManager) Stream(fBuilder func(proc *process.Process) stream.IEvent, fListener func() []*process.Process) <-chan stream.IEvent {
	eventChan := make(chan stream.IEvent)
	go func() {
		for {
			for _, out := range fListener() {
				pcEvent := fBuilder(out)
				time.Sleep(m.Rate())
				eventChan <- pcEvent
			}
		}

	}()

	return eventChan
}

func BuildProcEvent(proc *process.Process) stream.IEvent {
	// TODO use with context to allow a timeout.
	cpuPerc, err := proc.CPUPercent()
	// TODO this test is not evaluating correctly.
	if syscall.EPERM.Is(err) {
		return nil
	} else if err != nil {
		panic(err)
	}

	return ProcEvent{
		eventType:       stream.Process.String(),
		eventID:         uuid.NewString(),
		pid:             fmt.Sprintf("%d", proc.Pid),
		procUtilization: fmt.Sprintf("%f", cpuPerc),
		timeStamp:       time.Now(),
	}
}

func GetRunningProcs() []*process.Process {
	ps, err := process.Processes()
	if err != nil {
		panic(err)
	}
	return ps
}