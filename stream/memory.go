package stream

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/manther/events/util"
	"github.com/shirou/gopsutil/v3/mem"
)

type memEvent struct {
	eventType       string
	eventID         string
	*mem.VirtualMemoryStat
	timeStamp time.Time
}

func (e memEvent) ToString() string {
	return fmt.Sprintf("MemStats - Event Type: %s. Event ID: %s. Total: %v, Free:%v, UsedPercent:%f%%", e.eventType, e.eventID, e.Total, e.Free, e.UsedPercent)
}

func (e memEvent) TimeStamp() time.Time {
	return e.timeStamp
}

type MemManager struct {
	rate   time.Duration
	amount int
}

func NewMemoryManager(rate time.Duration, amount int) MemManager {
	return MemManager{
		rate:   rate,
		amount: amount,
	}
}

func (m MemManager) Amount() int {
	return m.amount
}

func (m MemManager) Rate() time.Duration {
	return m.rate
}

func (m MemManager) Stream(fBuilder func(memStat *mem.VirtualMemoryStat) IEvent, fListener func() (*mem.VirtualMemoryStat, error)) (<-chan IEvent, chan<- struct{}) {
	eventChan := make(chan IEvent)
	quitChan := make(chan struct{})
	go func() {
		for {
			out, err := fListener()
			if err != nil {
				fmt.Println("err getting mem stats.", err)
				return
			}
			pcEvent := fBuilder(out)
			time.Sleep(m.Rate())
			eventChan <- pcEvent

			select {
			case <-quitChan:
				close(eventChan)
				return
			default:
			}
		}

	}()

	return eventChan, quitChan
}

func GetMemStats() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

func BuildMemEvent(memStat *mem.VirtualMemoryStat) IEvent {
	return memEvent{
		eventType:         util.Memory.String(),
		eventID:           uuid.NewString(),
		VirtualMemoryStat: memStat,
		timeStamp:         time.Now(),
	}
}
