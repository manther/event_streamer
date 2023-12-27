package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/manther/events/stream/memory"
	"github.com/manther/events/stream/process"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go streamProcessEvents(&wg)
	go streamMemStats(&wg)
	wg.Wait()
}

func streamProcessEvents(wg *sync.WaitGroup) {
	manager := process.NewProcessManager(
		time.Second/5,
		20,
	)

	go func() {
		defer wg.Done()
		i := 1
		for event := range manager.Stream(process.BuildProcEvent, process.GetRunningProcs) {
			fmt.Println(event.ToString())
			if i == manager.Amount() {
				break
			}
			i++
		}
	}()
}

func streamMemStats(wg *sync.WaitGroup) {
	manager := memory.NewMemoryManager(
		time.Second/15,
		50,
	)

	go func() {
		defer wg.Done()
		events, quitchan := manager.Stream(memory.BuildMemEvent, memory.GetMemStats)
		i := 1
		for event := range events {

			fmt.Println(event.ToString())
			if i == manager.Amount() {
				close(quitchan)
				break
			}
			i++
		}
	}()
}
