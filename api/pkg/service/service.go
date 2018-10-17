package service

import (
	"context"
	"github.com/king19800105/feigo/api/pkg/io"
)

// ApiService describes the service.
type ApiService interface {
	// 短信api
	SMSSend(ctx context.Context, sendRequest io.SMSSendRequest) (string, error)
	SMSQuery(ctx context.Context, username string, key string) (string, error)
}

type basicApiService struct{}

func (b *basicApiService) SMSSend(ctx context.Context, sendRequest io.SMSSendRequest) (s0 string, e1 error) {
	// TODO implement the business logic of SMSSend
	return s0, e1
}
func (b *basicApiService) SMSQuery(ctx context.Context, username string, key string) (s0 string, e1 error) {
	// TODO implement the business logic of SMSQuery
	return s0, e1
}

// NewBasicApiService returns a naive, stateless implementation of ApiService.
func NewBasicApiService() ApiService {
	return &basicApiService{}
}

// New returns a ApiService with all of the expected middleware wired in.
func New(middleware []Middleware) ApiService {
	var svc ApiService = NewBasicApiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
