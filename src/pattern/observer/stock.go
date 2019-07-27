package observer

import (
	"fmt"
)

type Stock struct {
	max       float64
	min       float64
	avg       float64
	observers []Observer
}

func (stock *Stock) Add(observer Observer) {
	stock.observers = append(stock.observers, observer)

}

func (stock *Stock) Remove(observer Observer) {
	for index, value := range stock.observers {
		if observer == value {
			fmt.Println(index)
			stock.observers = append(stock.observers[:index], stock.observers[index+1:]...)
			return
		}
	}
}

func (stock *Stock) notify() {
	for index := range stock.observers {
		stock.observers[index].Update(stock.max, stock.min, stock.avg)
	}
}

func (stock *Stock) Receive(max float64, min float64, avg float64) {
	stock.avg = avg
	stock.max = max
	stock.min = min
	stock.notify()
}
