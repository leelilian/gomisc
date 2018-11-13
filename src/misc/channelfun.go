package main

import (
	"fmt"
)

var ch chan string

func main() {
	ch = make(chan string)
	go pingpong()
	msg := "ping"
	for {

		ch <- msg
		pong := <-ch
		fmt.Println("receive from pingpong:", pong)
	}
}

func pingpong() {
	for {
		ping := <-ch
		fmt.Println("receive from main: ", ping)
		ch <- "pong"
	}

	// ch <- ping

}
