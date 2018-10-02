package main

import (
	"fmt"
	"sync/atomic"
)

var i int32 = 0

var ch chan int

func increase() {
	j := atomic.AddInt32(&i, 1)
	fmt.Println("i = ", j)
	ch <- 0
	
}

func main() {
	
	ch = make(chan int)
	for i := 0; i < 10; i++ {
		go increase()
		
	}
	
	for i := 0; i < 10; i++ {
		<-ch
		
	}
	
}
