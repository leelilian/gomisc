package scheduler

import (
	"crawler/consumer"
	"crawler/framework"
)

type ConcurrentScheduler struct {
	RequestChannel chan framework.Request
}

func (cs *ConcurrentScheduler) DetermineRequestChannel() chan framework.Request {
	return cs.RequestChannel
}

func (cs *ConcurrentScheduler) ReadyWorker(w consumer.Worker) {

}

func (cs *ConcurrentScheduler) Run() {
	cs.RequestChannel = make(chan framework.Request)
}

func (cs *ConcurrentScheduler) Submit(request framework.Request) {
	go func() {
		cs.RequestChannel <- request
	}()

}
