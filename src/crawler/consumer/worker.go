package consumer

import (
	"log"

	"crawler/framework"
	"crawler/utils"
)

type Worker struct {
	Request chan framework.Request
}

type WorkerReadyNotifier interface {
	ReadyWorker(w Worker)
}

func (w *Worker) Handle(r framework.Request) (*framework.ParseResult, error) {
	log.Printf("fetching url: %s", r.Url)
	resp, err := utils.Fetch(r.Url)
	if err != nil {
		log.Printf("error fetch url: %s, %v", r.Url, err)
		return nil, err

	}
	return r.Parser(resp), nil
}

func (w *Worker) HandleAsync(out chan *framework.ParseResult, notifier WorkerReadyNotifier) {

	go func() {
		// in := make()
		for {
			notifier.ReadyWorker(*w)
			request := <-w.Request
			result, err := w.Handle(request)
			if err != nil {
				continue
			}
			out <- result

		}
	}()

}
