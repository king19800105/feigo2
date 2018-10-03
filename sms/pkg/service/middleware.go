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

func (l loggingMiddleware) Send(ctx context.Context, username string, key string, info SMSInfo) (i0 int, s1 string, s2 string) {
	defer func() {
		l.logger.Log("method", "Send", "username", username, "key", key, "info", info, "i0", i0, "s1", s1, "s2", s2)
	}()
	return l.next.Send(ctx, username, key, info)
}
func (l loggingMiddleware) Query(ctx context.Context, username string, key string) (f0 float64, e1 error) {
	defer func() {
		l.logger.Log("method", "Query", "username", username, "key", key, "f0", f0, "e1", e1)
	}()
	return l.next.Query(ctx, username, key)
}
