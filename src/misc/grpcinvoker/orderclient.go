package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gpmgo/gopm/modules/log"
	"google.golang.org/grpc"
	"misc/entity"
)

func main() {
	conn, err := grpc.Dial(":12345", grpc.WithInsecure())
	if err != nil {
		log.Fatal(" error:\n ", err)
	}
	defer conn.Close()

	client := entity.NewOrderServiceClient(conn)

	ctx := context.Background()
	/*
		rsp, err := client.GetOrders(ctx, &entity.OrderQueryRequest{OrderNo: "1"})
		if err != nil {
			fmt.Printf("error from server, %v\n", err)
		}
		printJson(rsp)

		streamclient, err := client.GetOrdersByStream(ctx)
		if err != nil {
			log.Fatal("init stream client error: %v\n", err)
		}

		go func() {
			for {
				i := rand.Intn(200)
				fmt.Printf("send to server, order no:%d\n", i)
				err := streamclient.Send(&entity.OrderQueryRequest{
					OrderNo: strconv.Itoa(i),
				})

				if err != nil {
					fmt.Printf("send to server: %v\n", err)
				}
				if i == 159 {
					break
				}
			}

		}()
		go func() {
			for {
				rsp, err := streamclient.Recv()

				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("receive error from server: %v\n", err)
				}

				printJson(rsp)

			}
		}()
	*/

	c, err := client.GetStreamResponseOrders(ctx, &entity.OrderQueryRequest{OrderNo: "1,2,3,4,5,6"})
	if err != nil {
		log.Fatal("c error:%v", err)
	}
	var count int
	for {
		rsp, err := c.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("rsp error:%v", err)
		}
		count++
		fmt.Println(count)
		printJson(rsp)
	}

	/*
		stream, err := client.GetOrdersByClientStream(ctx)
		if err != nil {
			panic(err)
		}
		for i := 0; i < 5; i++ {
			err = stream.Send(
				&entity.OrderQueryRequest{
					OrderNo: strconv.Itoa(i),
				})
			if err != nil {
				panic(err)
			}
		}

		rsp, err := stream.CloseAndRecv()
		fmt.Println(rsp)
		if err != nil {
			panic(err)
		}
		printJson(rsp)
	*/
	// select {}

}

func printJson(data *entity.OrderListResponse) {
	txt, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("marshalling to json error, %v\n", err)
		return
	}
	fmt.Printf("time:%s, %s\n", time.Now(), txt)
	file, err := os.OpenFile("txt.txt", os.O_CREATE|os.O_RDWR, 0666)
	w := bufio.NewWriter(file)
	w.WriteString((string)(txt))
	w.Flush()
}
