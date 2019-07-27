package observer

import (
	"fmt"
)

type Bank struct {
}

func (bank *Bank) Update(max float64, min float64, avg float64) {

	fmt.Printf("bank max: %f, min: %f, avg: %f\n", max, min, avg)
}
