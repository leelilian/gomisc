package scheduler

import (
	"crawler/consumer"
	"crawler/framework"
)

type QueueScheduler struct {
	RequestChannel chan framework.Request
	WorkerChannel  chan consumer.Worker
}

func (sch *QueueScheduler) DetermineRequestChannel() chan framework.Request {
	return make(chan framework.Request)
}

func (sch *QueueScheduler) Submit(request framework.Request) {
	sch.RequestChannel <- request
}

func (sch *QueueScheduler) ReadyWorker(w consumer.Worker) {
	sch.WorkerChannel <- w
}

func (sch *QueueScheduler) Run() {
	sch.WorkerChannel = make(chan consumer.Worker)
	sch.RequestChannel = make(chan framework.Request)
	go func() {
		rq := framework.RequestQueue{}
		wq := consumer.WorkerQueue{}

		for {
			var activeRequest framework.Request
			var activeWorker consumer.Worker
			if rq.Len() > 0 && wq.Len() > 0 {
				activeRequest = rq.Peek(0)

				activeWorker = wq.Peek(0)
			}

			select {
			case r := <-sch.RequestChannel:
				// log.Printf("request received: %v", r)
				rq.Enqueue(r)
			case w := <-sch.WorkerChannel:
				// log.Printf("worker received: %v", w)
				wq.Enqueue(w)
			case activeWorker.Request <- activeRequest:
				// log.Printf("workload received:")
				rq.Dequeue()
				wq.Dequeue()
			default:
				// log.Printf("well, I have no idea what to do")
			}

		}
	}()
}
