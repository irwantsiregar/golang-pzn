package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Read Write Mutex
type BankAccount struct {
	RWMutex sync.RWMutex 
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}


func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance 
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	// GoRoutine with handling race condition
	for i := 1; i <= 100; i++ {
		go func ()  {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)

				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Total Balance: ", account.GetBalance())
}
