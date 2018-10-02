package misc

import (
	"fmt"
)

func main() {
	
	/*
		c = make(chan string)

		go input()

		for {
			var s string
			fmt.Scan(&s)
			c <- s
		}
	*/
	c := gen(2, 3, 4, 5, 6)
	out := sq(c)
	
	// Consume the output.
	for index := 1; index <= 5; index++ {
		fmt.Println(<-out)
	}
	// 4
	// fmt.Println(<-out) // 9
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
