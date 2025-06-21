package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // 3 is buffered

	defer close(channel)

	go func ()  {
		channel <- "Irwan"
		channel <- "Siregar"
		channel <- "Sigadong"
	}()

	// This is can fetch multiple data from channel without save to variable, because existing buffered in channel
	go func ()  {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}
