package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/manther/events/stream"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go streamProcessEvents(&wg)
	go streamMemStats(&wg)
	wg.Wait()
}

func streamProcessEvents(wg *sync.WaitGroup) {
	manager := stream.NewProcessManager(
		time.Second/5,
		20,
	)

	go func() {
		defer wg.Done()
		events, quitchan := manager.Stream(stream.BuildProcEvent, stream.GetRunningProcs)
		i := 1
		for event := range events {
			fmt.Println(event.ToString())
			if i == manager.Amount() {
				close(quitchan)
				return
			}
			i++
		}
	}()
}

func streamMemStats(wg *sync.WaitGroup) {
	manager := stream.NewMemoryManager(
		time.Second/15,
		50,
	)

	go func() {
		defer wg.Done()
		events, quitchan := manager.Stream(stream.BuildMemEvent, stream.GetMemStats)
		i := 1
		for event := range events {
			
			fmt.Println(event.ToString())
			if i == manager.Amount() {
				close(quitchan)
				return
			}
			i++
		}
	}()
}

