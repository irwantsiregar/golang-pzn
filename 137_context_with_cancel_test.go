package main

import (
	"fmt"
	"runtime"
	"testing"
)

func CreateCounter() chan int {
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

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total GoRoutine", runtime.NumGoroutine())

	destination := CreateCounter()

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	fmt.Println("Total GoRoutine", runtime.NumGoroutine())
}