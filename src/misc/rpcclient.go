package misc

import (
	"fmt"
	"net/rpc"
)

//
type Agrs struct {
	Num1 int
	Num2 int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:50001")
	if err != nil {
		fmt.Println(err)
		return
	}
	args := &Agrs{Num1: 3, Num2: 4}
	var reply int
	
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("sum is %d\n", reply)
}
