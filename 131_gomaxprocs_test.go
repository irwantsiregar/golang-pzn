package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}

	// Make goroutine any 100 and 2 built in. (100 + 2 = 102 goroutine)
	for i := 0; i < 100; i++ {
		group.Add(1)

		go func ()  {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU() // Total CPU
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(20) // For UPDATE total thread
	totalThread := runtime.GOMAXPROCS(-1)  // Default check thread is value -1, for update use value >0
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine() // Total GoRoutine
	fmt.Println("Total Goroutine", totalGoroutine)
	
	group.Wait()
}