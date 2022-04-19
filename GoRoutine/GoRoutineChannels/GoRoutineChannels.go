package main

import "fmt"

var c chan int

func main() {
	c := make(chan string)

	go greet(c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c, ",", <-c)
	}
	// buffered channel , length 7
	stuff := make(chan int, 7)
	for i := 0; i < 19; i = i + 3 {
		stuff <- i
	}

	close(stuff)

	fmt.Println("Res ",process(stuff))



}

func process(input <-chan int) (res int) {
	for r := range input {
		res +=r
	}
	return
}

func greet(c chan<- string) {
	for true {
		c <- fmt.Sprintf("Владыка")
		c <- fmt.Sprintf("Штурмовик")
	}
}
