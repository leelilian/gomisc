package main

import (
	"fmt"
)

type slice []int

func (s *slice) append(a int) {
	
	*s = append(*s, a)
}

func main() {
	/*
	t := make(slice, 10, 20)
	t.append(23)
	
	fmt.Printf("%v", t)*/
	urls := make(map[string]string, 3)
	// 这里随便个例子
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"
	
	var s string
	fmt.Println(s)
	
	names := make([]string, 4)
	for index, v := range names {
		
		fmt.Println(index, len(v))
	}
	fmt.Println(names)
	for key, _ := range urls {
		fmt.Println(key)
		
		names = append(names, key)
	}
	
	fmt.Println(names)
	
	fmt.Println(cap(names))
	
}
