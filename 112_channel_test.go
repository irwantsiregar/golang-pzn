package main

import (
	"fmt"
	"testing"
	"time"
)

func TestConceptChannel(t *testing.T) {
	// Create Channel
	channel := make(chan string)

	// Send data to channel
	channel <- "Irwant"
	
	// Fetch data from channel
	data := <- channel

	fmt.Println(data)
	time.Sleep(2 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel) // Close channel

	go func ()  {
		time.Sleep(2 * time.Second)
		
		channel <- "Irwanto Siregar"

		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
