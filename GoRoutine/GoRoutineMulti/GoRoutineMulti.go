package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving")
}

func fanIn(input1 <-chan string, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for true {
			c <- <-input1
		}
	}()
	go func() {
		for true {
			c <- <-input2
		}
	}()
	return c

}

func boring(msg string) <-chan string { //only read channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}
