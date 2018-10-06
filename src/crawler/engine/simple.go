package engine

import (
	"log"

	"crawler/framework"
	"crawler/scheduler"
	"crawler/worker"
)

type SimpleEngine struct {
	Scheduler scheduler.SyncScheduler
}

func (this *SimpleEngine) Run(seeds ...framework.Request) {

	for _, m := range seeds {

		this.Scheduler.Submit(m)
	}
	for {
		request := this.Scheduler.Dispatch()
		result, err := worker.Handle(request)
		if err != nil {
			continue
		}

		for _, req := range result.RequestList {

			this.Scheduler.Submit(req)
		}

		for _, item := range result.Items {
			log.Printf("item: %v", item)
		}

	}

}
