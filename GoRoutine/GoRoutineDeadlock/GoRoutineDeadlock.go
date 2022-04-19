package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	//channel for players interconnection
	table := make(chan *Ball)
	// starting two players
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) //launch ball in the game
	time.Sleep(1 * time.Second)
	<-table //Stop the game.
}

func player(name string, tab chan *Ball) {
	for true {
		ball := <-tab
		ball.hits++
		fmt.Println(name,ball.hits)
		time.Sleep(100*time.Millisecond)

		tab <- ball
	}
}
