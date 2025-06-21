package main

import (
	"fmt"
	"testing"
	"time"
)

// Special func for sending data through channel
func OnlyIn(channel chan <- string) {
	time.Sleep(2 * time.Second)

	channel <- "Irwant Siregar"
}

// Special func for fetching data from channel
func OnlyOut(channel <- chan string) {
	data := channel

	fmt.Println(data)
}


func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	
	defer close(channel)

	// Func for in and out data via channel
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
