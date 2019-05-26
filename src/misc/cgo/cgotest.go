package main

import "C"
import (
	"fmt"
)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L/Users/ethan.li/Documents/src/go/src/misc/cgo/ -lfoo

#include "foo.h"
*/
import "C"

func main() {
	fmt.Println("cgo")
	C.foo()
}
