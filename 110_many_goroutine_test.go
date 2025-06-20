package main

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display", number);
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}



/*
- Pada komputer generasi terbaru kebanyakan sudah Multi-Core (ex: 12 core)
- Golang secara default adalah Concurrency (proses berjalan secara bergantian)
- Dikarenakan Multi-Core, maka Go-Routine berjalan secara Concurrency dan Parallel -> (Async Await)
- Eksekusi Go-Routine berjalan secara tidak berurutan seperti layaknya async-await 
*/