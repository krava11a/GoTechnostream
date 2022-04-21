package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func boring(wg *sync.WaitGroup,die chan bool) <-chan string { //Возвращаем канал строк только для чтения
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("Boring %d", rand.Intn(100)):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-die:
				fmt.Println("Jobs done!!")
				wg.Done()
				return
			}
		}
	}()
	return c
}

func main() {

	die := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	res1 := boring(&wg,die)
	res2 := boring(&wg,die)
	res3 := boring(&wg,die)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say  1 : %q \n", <-res1)
		fmt.Printf("You say  2 : %q \n", <-res2)
		fmt.Printf("You say  3 : %q \n", <-res3)
	}
	die <- true
	fmt.Println("GoodBye")
	wg.Wait()


}
