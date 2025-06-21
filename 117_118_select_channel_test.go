package main

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponses(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "Irwanto Siregar"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponses(channel1)
	go GiveMeResponses(channel2)
	
	counter := 0

	// For with infinity loop
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++	
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		// Counter for break loop
		if counter == 2 {
			break	
		}
	}
}

func TestSelectDefdaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponses(channel1)
	go GiveMeResponses(channel2)
	
	counter := 0

	// For with infinity loop
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++	
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data") // Display when waiting channel process
		}

		// Counter for break loop
		if counter == 2 {
			break	
		}
	}
}
