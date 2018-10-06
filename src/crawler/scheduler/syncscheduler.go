package scheduler

import (
	"crawler/framework"
)

type SyncScheduler interface {
	Submit(request framework.Request)
	Dispatch() framework.Request
}
