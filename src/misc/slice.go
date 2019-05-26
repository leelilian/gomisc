package main

import (
	"fmt"
)

type slice []int

func (s *slice) append(a int) {

	*s = append(*s, a)
}

func main() {
	t := make(slice, 10, 20)
	t.append(23)

	fmt.Printf("%v", t)

}
