package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	service "github.com/king19800105/feigo/sms/pkg/service"
)

// Middleware describes a service middleware.
type Middleware func(ApiService) ApiService

type loggingMiddleware struct {
	logger log.Logger
	next   ApiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ApiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ApiService) ApiService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SMSSend(ctx context.Context, username string, key string, content service.SMSContent) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SMSSend", "username", username, "key", key, "content", content, "s0", s0, "e1", e1)
	}()
	return l.next.SMSSend(ctx, username, key, content)
}
func (l loggingMiddleware) SMSQuery(ctx context.Context, username string, key string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SMSQuery", "username", username, "key", key, "s0", s0, "e1", e1)
	}()
	return l.next.SMSQuery(ctx, username, key)
}
