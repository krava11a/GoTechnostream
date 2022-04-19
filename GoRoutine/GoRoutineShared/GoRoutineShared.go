package main

import (
	"fmt"
	"log"
)

type Account struct {
	balance float64
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	log.Printf("depositing: %f", amount)
	a.balance += amount
}

func (a *Account) WithDraw(amount float64) {
	if amount > a.balance {
		return
	}
	log.Printf("withdrawing: %f", amount)
	a.balance -= amount
}

func main() {
	acc := Account{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
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
	//closure()
}

func closure() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Got",i)
		}(i)
	}
	fmt.Scanln()
}
