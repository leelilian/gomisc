package scheduler

import (
	"crawler/framework"
)

type SimpleScheduler struct {
	Requests []framework.Request
}

func (this *SimpleScheduler) Submit(request framework.Request) {

	this.Requests = append(this.Requests, request)
}

func (this *SimpleScheduler) Dispatch() framework.Request {
	var request framework.Request
	if len(this.Requests) > 0 {

		request = this.Requests[0]
		this.Requests = this.Requests[1:]

	}
	return request
}
