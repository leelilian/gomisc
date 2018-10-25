package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Printf("hello, %d\n", x)
		x, y = y, x+y
	}
	close(c)
}

func main() {
	for i := 0; i <= 100; i++ {
		c := make(chan int, 10)
		go fibonacci(cap(c), c)
		for i := range c {
			fmt.Println(i)
		}
	}

}
