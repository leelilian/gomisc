package engine

import (
	"log"

	"crawler/consumer"
	"crawler/framework"
	"crawler/scheduler"
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
		w := consumer.Worker{}
		result, err := w.Handle(request)
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
