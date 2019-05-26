package stringserver

import (
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   StringService
}

func (mw LoggingMiddleware) ToUpper(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.ToUpper(s)
	return
}
