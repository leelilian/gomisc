package observer

type Observer interface {
	Update(max float64, min float64, avg float64)
}
