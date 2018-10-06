package service

import (
	"context"

	smsservice "github.com/king19800105/feigo/sms/pkg/service"
)

// 外部服务接口
// 对外提供http访问接口，调用内部响应grpc完成业务
type ApiService interface {
	// 短信api
	SMSSend(ctx context.Context, username string, key string, content smsservice.SMSContent) (string, error)
	SMSQuery(ctx context.Context, username string, key string) (string, error)
}

type basicApiService struct{}

func (b *basicApiService) SMSSend(ctx context.Context, username string, key string, content smsservice.SMSContent) (s0 string, e1 error) {
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
