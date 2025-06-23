package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
	Context with channel leak
*/
func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func ()  {
		defer close(destination)
		counter := 1

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestContextWithCancelLeak(t *testing.T) {
	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	destination := CreateCounterLeak()

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())
}


/*
	Context with Cancel
*/
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func ()  {
		defer close(destination)
		counter := 1

		for {
			select {
			case <- ctx.Done():
					return
				default:
					destination <- counter
					counter++
			}
		}
	}()

	return destination
}
func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	// Init Context
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx) // Context as parameter

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	// Sending signal cancel to Context
	cancel() 

	// Time for make sure goroutines has killed
	time.Sleep(2 * time.Second) 

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())
}