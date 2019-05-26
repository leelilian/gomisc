package order

type Service interface {
	Get(id string) Order
}

type orderSvr struct {
}

func (svr *orderSvr) Get(id string) Order {
	orders := []Order{
		Order{"1", "hello1", 7.9, 100},
		Order{"2", "hello2", 7.9, 200},
		Order{"3", "hello3", 7.9, 300},
		Order{"4", "hello4", 7.9, 400},
	}
	for index, order := range orders {
		if order.Id == id {
			return orders[index]
		}
	}
	return Order{}
}
