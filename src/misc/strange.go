package main


import (
	"fmt"
)

func main() {
	
	s := []byte("")
	
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s1 := append(s, 'a')
	
	s2 := append(s, 'b')
	// fmt.Printf("%p, %p, %p\n", s, s1, s2)
	
	// fmt.Println(s1, "***", s2)
	fmt.Println(string(s1), string(s2))
	// fmt.Printf("%p, %p, %p\n", s, s1, s2)
}
