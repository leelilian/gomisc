package observer

import (
	"fmt"
)

type Agent struct {
}

func (agent *Agent) Update(max float64, min float64, avg float64) {

	fmt.Printf("agent max: %f, min: %f, avg: %f\n", max, min, avg)
}
