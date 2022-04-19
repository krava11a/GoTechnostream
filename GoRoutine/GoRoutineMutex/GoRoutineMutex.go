package main

import (
	"fmt"
	"log"
	"sync"
)

type AccountProtected struct {
	sync.RWMutex
	//sync.Mutex
	balance float64
}

func (a *AccountProtected) Balance() float64 {
	a.RLock()
	//a.Lock()
	defer a.RUnlock()
	//defer a.Unlock()
	return a.balance
}

func (a *AccountProtected) Deposit(amount float64) {
	a.Lock()
	defer a.Unlock()
	log.Printf("depositing: %f", amount)
	a.balance += amount
}

func (a *AccountProtected) WithDraw(amount float64) {
	a.Lock()
	defer a.Unlock()
	if amount > a.balance {
		return
	}
	log.Printf("withdrawing: %f", amount)
	a.balance -= amount
}

func main() {
	acc := AccountProtected{}
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				if j%2 == 1 {
					acc.WithDraw(50)
					continue
				}
				acc.Deposit(50)
			}
		}()
	}
	fmt.Scanln()
	fmt.Println(acc.Balance())
}
