package service

import "context"

// 短信业务服务
// 短信核心业务处理，以grpc方式对api提供业务接口
type SmsService interface {
	Send(ctx context.Context, content SMSContent) (string, error)
}

type SMSContent struct {
}

type basicSmsService struct{}

func (b *basicSmsService) Send(ctx context.Context, content SMSContent) (s0 string, e1 error) {
	// TODO implement the business logic of Send
	return s0, e1
}

// NewBasicSmsService returns a naive, stateless implementation of SmsService.
func NewBasicSmsService() SmsService {
	return &basicSmsService{}
}

// New returns a SmsService with all of the expected middleware wired in.
func New(middleware []Middleware) SmsService {
	var svc SmsService = NewBasicSmsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
