package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T){
	pool := sync.Pool{
		New: func () interface{}  {
			return "New Pool"
		},
	}

	pool.Put("Irwan")
	pool.Put("Siregar")
	pool.Put("Gadong")

	for i := 0; i < 10; i++ {
		go func ()  {
			data := pool.Get()

			fmt.Println(data)

			time.Sleep(1 * time.Second)

			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)

	fmt.Println("Selesai")
}


/*
	Pool: Biasanya digunakan ketika ingin melakukan koneksi ke database (ex: 10 connection)
	Pool: Has secure from race condition
*/