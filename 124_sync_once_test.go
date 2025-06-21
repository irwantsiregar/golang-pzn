package main

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func RunOnce(group *sync.WaitGroup, once *sync.Once) {
	group.Add(1)

	once.Do(OnlyOnce) // Only once execute goroutine process

	group.Done()
}

func TestOnce(t *testing.T){
	once := &sync.Once{}
	group := &sync.WaitGroup{}
	
	for i := 0; i < 100; i++ {
		go func ()  {
			RunOnce(group, once)
		}()
	}

	group.Wait()
	
	fmt.Println(counter)
	fmt.Println("Selesai")
}
