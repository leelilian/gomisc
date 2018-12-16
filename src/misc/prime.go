package main

import (
	"fmt"
)

var index = 0

func main() {
	sieve()
}

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}
func filter(src <-chan int, dst chan<- int, prime int) {
	index++

	for i := range src { // Loop over values received from 'src'.
		fmt.Println("prime=", prime, "   i=", i)
		if i%prime != 0 {
			fmt.Printf("index=%d\n", index)
			fmt.Printf("send=%d\n", i)
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

func sieve() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a subprocess.
	for i := 0; i < 6; i++ {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		// fmt.Printf("%v\n", ch1)
		go filter(ch, ch1, prime)
		ch = ch1
		// fmt.Printf("%v\n", ch)
	}
}
