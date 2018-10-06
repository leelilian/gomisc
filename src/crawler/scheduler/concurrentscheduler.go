package scheduler

import (
	"crawler/framework"
)

type ConcurrentScheduler struct {
	RequestChannel chan framework.Request
}

func (cs *ConcurrentScheduler) Submit(request framework.Request) {
	go func() {
		cs.RequestChannel <- request
	}()

}

func (cs *ConcurrentScheduler) SetRequestChannel(request chan framework.Request) {
	cs.RequestChannel = request

}
