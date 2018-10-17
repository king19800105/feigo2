package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/king19800105/feigo/api/pkg/io"
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

func (l loggingMiddleware) SMSSend(ctx context.Context, r io.SMSSendRequest) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SMSSend", "username", r.Username, "key", r.Key, "content", r.Content, "sign", r.Sign, "extcode", r.ExtensionCode, "s0", s0, "e1", e1)
	}()
	return l.next.SMSSend(ctx, r)
}
func (l loggingMiddleware) SMSQuery(ctx context.Context, username string, key string) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "SMSQuery", "username", username, "key", key, "s0", s0, "e1", e1)
	}()
	return l.next.SMSQuery(ctx, username, key)
}
