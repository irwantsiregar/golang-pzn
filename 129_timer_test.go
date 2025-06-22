package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer is representation a event
// TIMER: FOR DELAY JOB

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <- channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func ()  {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())

	group.Wait()
}
