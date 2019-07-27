package main

import (
	"pattern/observer"
)

func main() {

	bank := &observer.Bank{}

	agent := &observer.Agent{}
	agent1 := &observer.Agent{}

	stock := &observer.Stock{}

	stock.Add(bank)

	stock.Add(agent)
	stock.Add(agent1)

	stock.Remove(agent)

	for i := 0; i < 1000000000; i++ {
		stock.Receive((float64)(i*1.0), (float64)(i*1.0), (float64)(i*1.0))
	}

}
