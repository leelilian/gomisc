package scheduler

import (
	"crawler/framework"
)

type AsyncScheduler interface {
	Submit(request framework.Request)

	SetRequestChannel(request chan framework.Request)
}
