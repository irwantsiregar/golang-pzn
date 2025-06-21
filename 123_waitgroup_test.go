package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // To mark that the goroutine process has been completed

	group.Add(1) // To mark that there is a goroutine process
	
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T){
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait() // Waiting all goroutine process has been completed
	
	fmt.Println("Selesai")

}
