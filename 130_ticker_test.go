package main

import (
	"fmt"
	"testing"
	"time"
)

// Ticker is representation events loop
// TIMER: FOR DELAY JOB

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func ()  {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}