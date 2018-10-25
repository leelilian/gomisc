package main

import (
	"net"

	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"misc/helloworld"
)

func main() {
	host := ":12345"
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal("server failed to listen %s: %v", host, err)

	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, server{})
	err = s.Serve(listener)
	if err != nil {
		log.Fatal("server failed to listen %s: %v", host, err)
	}
}

type server struct {
}

func (s server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}
