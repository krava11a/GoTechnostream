package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(die chan bool) <-chan string { //Возвращаем канал строк только для чтения
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("Boring %d", rand.Intn(100)):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-die:
				fmt.Println("Jobs done!!")
				die <- true
				return
			}
		}
	}()
	return c
}

func main() {

	die := make(chan bool)
	res := boring(die)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q \n", <-res)
	}
	die <- true
	<-die

}
