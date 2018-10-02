package main

import (
	"fmt"
)

type flyer interface {
	fly()
}

type bird struct {
}

func (b *bird) fly() {
	
	fmt.Println("bird is flying")
	
}

type plane struct {
}

func (p *plane) fly() {
	fmt.Println("aircraft is flying")
	
}

func main() {
	var item flyer
	item = &plane{}
	item.fly()
	
	item = &bird{}
	item.fly()
	
}
