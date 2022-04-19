package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//ограничение по ЦПУ
	runtime.GOMAXPROCS(runtime.NumCPU())

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")

			}
			//В этот момент можно прерывать
			//runtime.Gosched()
		}
	}()
	fmt.Scanln()


}

