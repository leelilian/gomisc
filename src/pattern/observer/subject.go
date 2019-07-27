package observer

type Subject interface {
	Add(observer Observer)
	Remove(observer Observer)
	notify()
}
