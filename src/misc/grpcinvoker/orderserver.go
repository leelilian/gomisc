package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"misc/entity"
)

type orderserver struct {
}

func (server *orderserver) GetOrders(ctx context.Context, req *entity.OrderQueryRequest) (*entity.OrderListResponse, error) {
	
	if req == nil {
		return nil, fmt.Errorf("requst is empty")
		
	}
	rsp := findOrderByNo(req.OrderNo)
	return rsp, nil
}

func findOrderByNo(orderNo string) *entity.OrderListResponse {
	
	var list []*entity.Order
	for _, o := range orders {
		if strings.EqualFold(o.OrderNo, orderNo) {
			list = append(list, o)
		}
	}
	
	rsp := entity.OrderListResponse{
		OrderList: list,
	}
	if len(list) > 0 {
		rsp.ResultCode = 200
		rsp.Message = "OK"
		
	} else {
		rsp.ResultCode = 404
		rsp.Message = "NOT FOUND"
	}
	return &rsp
}

func (server *orderserver) GetOrdersByStream(stream entity.OrderService_GetOrdersByStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		
		fmt.Printf("receive order id:%s\n", req.OrderNo)
		rsp := findOrderByNo(req.OrderNo)
		
		err = stream.Send(rsp)
		if err != nil {
			return err
		}
		
	}
}

func (server *orderserver) GetStreamResponseOrders(req *entity.OrderQueryRequest,
	stream entity.OrderService_GetStreamResponseOrdersServer) error {
	
	if req == nil {
		return fmt.Errorf("request is empty")
	}
	
	list := strings.Split(req.OrderNo, ",")
	for _, no := range list {
		rsp := findOrderByNo(no)
		err := stream.Send(rsp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (server *orderserver) GetOrdersByClientStream(stream entity.OrderService_GetOrdersByClientStreamServer) error {
	
	var rsplist []*entity.Order
	var response entity.OrderListResponse
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				
				response.OrderList = rsplist
				if len(rsplist) > 0 {
					response.ResultCode = 200
					response.Message = "OK"
				} else {
					response.ResultCode = 404
					response.Message = "NOT FOUND"
				}
				//  fmt.Println(response)
				err = stream.SendAndClose(&response)
				return err
				
			}
			
			return err
		}
		rsp := findOrderByNo(req.OrderNo)
		rsplist = append(rsplist, rsp.OrderList...)
		
	}
	return nil
	
}

func generate() {
	if len(orders) == 0 {
		
		items := []*entity.Item{
			{ItemNo: "itemno1", ItemName: "itemname1", Price: 1.23},
			{ItemNo: "itemno2", ItemName: "itemname2", Price: 1.24},
			{ItemNo: "itemno3", ItemName: "itemname3", Price: 1.25},
			{ItemNo: "itemno4", ItemName: "itemname4", Price: 1.26},
			{ItemNo: "itemno5", ItemName: "itemname5", Price: 1.27},
			{ItemNo: "itemno6", ItemName: "itemname6", Price: 1.28},
		}
		
		address := entity.Address{
			Country:  "PRC",
			Province: "GD",
			City:     "SHZ",
			Street1:  "you guess",
			Street2:  "anywhere",
			PostCode: "518000",
		}
		
		contact := entity.Contact{
			FirstName: "mark",
			LastName:  "waterloon",
			Phone:     "12231231dd",
		}
		
		for i := 0; i < 25000; i++ {
			order := &entity.Order{
				OrderNo:       strconv.Itoa(i),
				Items:         items,
				ShipToContact: &contact,
				ShipToAddress: &address,
				BillToContact: &contact,
				BillToAddress: &address,
			}
			
			orders = append(orders, order)
		}
		
	}
	
}

var orders []*entity.Order

func init() {
	generate()
}

func main() {
	server := grpc.NewServer()
	entity.RegisterOrderServiceServer(server, &orderserver{})
	
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("server error:%v", err)
	}
	server.Serve(listener)
}
