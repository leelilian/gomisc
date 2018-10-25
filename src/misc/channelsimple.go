package main

import (
	"fmt"
	"math/rand"
)

func generate() chan int {

	out := make(chan int)

	go func() {

		for {

			number := rand.Int()
			fmt.Println("generated:", number)
			out <- number
		}

	}()

	return out
}

func main() {

	ch := generate()

	for {
		fmt.Println("received from channel: ", <-ch)
	}

}
