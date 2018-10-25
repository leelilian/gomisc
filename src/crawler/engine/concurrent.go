package engine

import (
	"log"

	"crawler/consumer"
	"crawler/framework"
	"crawler/scheduler"
)

type ConcurrentEngine struct {
	Scheduler   scheduler.AsyncScheduler
	WorkerCount int
}

func (ce *ConcurrentEngine) Run(seeds ...framework.Request) {

	out := make(chan *framework.ParseResult)
	ce.Scheduler.Run()

	for i := 0; i < ce.WorkerCount; i++ {
		in := ce.Scheduler.DetermineRequestChannel()
		w := consumer.Worker{Request: in}
		w.HandleAsync(out, ce.Scheduler)
	}

	for _, request := range seeds {
		ce.Scheduler.Submit(request)
	}

	for {

		result := <-out
		for _, item := range result.Items {
			log.Printf("item: %v", item)

		}
		for _, request := range result.RequestList {
			ce.Scheduler.Submit(request)

		}

	}
}
