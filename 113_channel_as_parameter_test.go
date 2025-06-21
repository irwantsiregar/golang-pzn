package main

import (
	"fmt"
	"testing"
	"time"
)

// Func with channel parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
		
	channel <- "Irwanto"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	// This func waiting data from channel
	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
