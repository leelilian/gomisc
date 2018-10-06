package engine

import (
	"log"

	"crawler/framework"
	"crawler/scheduler"
	"crawler/worker"
)

type ConcurrentEngine struct {
	Scheduler   scheduler.AsyncScheduler
	WorkerCount int
}

func (ce ConcurrentEngine) Run(seeds ...framework.Request) {

	for _, request := range seeds {
		ce.Scheduler.Submit(request)
	}

	in := make(chan framework.Request)
	out := make(chan *framework.ParseResult)
	ce.Scheduler.SetRequestChannel(in)

	for i := 0; i < ce.WorkerCount; i++ {
		worker.HandleAsync(in, out)
	}

	for {

		result := <-out
		for _, m := range result.Items {
			log.Printf("item: %v", m)

		}
		for _, request := range result.RequestList {
			ce.Scheduler.Submit(request)

		}

	}
}
