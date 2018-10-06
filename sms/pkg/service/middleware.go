package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(SmsService) SmsService

type loggingMiddleware struct {
	logger log.Logger
	next   SmsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SmsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SmsService) SmsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Send(ctx context.Context, content SMSContent) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "Send", "content", content, "s0", s0, "e1", e1)
	}()
	return l.next.Send(ctx, content)
}
