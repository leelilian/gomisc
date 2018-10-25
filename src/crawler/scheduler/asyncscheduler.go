package scheduler

import (
	"crawler/consumer"
	"crawler/framework"
)

type AsyncScheduler interface {
	consumer.WorkerReadyNotifier
	Submit(request framework.Request)

	DetermineRequestChannel() chan framework.Request

	Run()
}
