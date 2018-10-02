package misc

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

//
type Agrs struct {
	Num1 int
	Num2 int
}

//
type Arith struct {
}

//
func (arith *Arith) Add(args *Agrs, reply *int) error {
	
	*reply = args.Num1 + args.Num2
	return nil
	
}

func main() {
	arth := new(Arith)
	rpc.Register(arth)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":50001")
	if err != nil {
		fmt.Println(err)
	}
	http.Serve(l, nil)
}
