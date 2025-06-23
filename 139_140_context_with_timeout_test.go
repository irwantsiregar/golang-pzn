package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
	Using context with timeout is very suitable when,
	Example: we are querying a database or http api, but want to determine the maximum timeout limit.
*/

// Context will cancel when duration is timout
func CreateCounter2(ctx context.Context) chan int {
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
					time.Sleep(1 * time.Second) // Slow simulation
			}
		}
	}()

	return destination
}

// TIMEOUT
func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	// Init Context
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	
	defer cancel() // auto cancel when duration has timeout

	destination := CreateCounter2(ctx) // Context as parameter

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	// Time for make sure goroutines has killed
	time.Sleep(2 * time.Second) 

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())
}


// DEADLINE
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second)) // time deadline
	
	defer cancel() // auto cancel when duration has timeout

	destination := CreateCounter2(ctx)

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second) 

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())
}