package main

import "fmt"

func main() {

	fmt.Println("start")

	//функция
	go process(0)
	//анонимная функция
	go func() {
		fmt.Println("Anonymus start")
	}()

	for i :=0;i < 1000; i++ {
		go process(i)
	}

	//Для завершения результата (ожидание)
	fmt.Scanln()
	
}

func process(i int) {
	fmt.Println("handle:   ", i)
}
