//Don't communicate by sharing memory; share memory by communications

package main

import (
	"errors"
	"fmt"
	"sync"
)

type AccountAsync struct {
	balance     float64
	deltaChan   chan float64
	balanceChan chan float64
	errChan     chan error
}

//Constructor
func NewAccount(balance float64) (a *AccountAsync) {
	a = &AccountAsync{
		balance:     balance,
		deltaChan:   make(chan float64),
		balanceChan: make(chan float64),
		errChan:     make(chan error, 1),
	}

	go a.run()
	return
}

//Бесконечный цикл обработчика счета
//теперь сколько бы горутин не производили операции над этим аккаунтом
//Все они будут синхронизированы здесь, и блокировки уже не нужны
func (a *AccountAsync) run() {
	var delta float64
	for {
		select {
		// Если поступили изменения
		case delta = <-a.deltaChan:
			// Попробуем их применить
			a.errChan<-a.applyDelta(delta)
			// Если кто-то запрашивает баланс
		case a.balanceChan <-a.balance:
			// Не делаем ничего, так как мы уже отправили ответ

		}
	}
}

//Просто читаем из канала баланса
func (a *AccountAsync) Balance() float64 {
	return <-a.balanceChan
}

//Пишем в канал изменения
func (a *AccountAsync) Deposit(amount float64) error {
	a.deltaChan <- amount
	return <-a.errChan
}

//Аналогично, по сути эта функция только для сохранения семантики
func (a *AccountAsync) Withdraw(amount float64) error {
	a.deltaChan <- -amount
	return <-a.errChan
}

//применение изменений к счету
func (a *AccountAsync) applyDelta(amount float64) error {
	stateStr := "Кладем на счет"
	if amount < 0 {
		stateStr = "Снимаем"
	}
	fmt.Println(stateStr, amount)

	newBalance := a.balance + amount
	if newBalance < 0 {
		return errors.New("Insufficient funds")
	}
	a.balance = newBalance
	return nil
}

func main() {
	acc := NewAccount(20)

	с := sync.Once{}

	//Стартуем 10 go рутин
	for i := 0; i < 10; i++ {
		go func() {

			//Выполняем только один раз
			с.Do(func() {
				fmt.Println(acc.Withdraw(100))
			})
			// Каждая из которых, производит операции с аккаунтом
			for j := 0; j<10; j++{
				// Иногда снимает деньги
				if j%2 == 1 {
					acc.Withdraw(50)
					continue
				}
				acc.Deposit(50)
			}
		}()
	}
	fmt.Scanln()
	fmt.Println(acc.Balance())
}
