package service

import "context"

// 对外开放的短信服务，内部调用api微服务，走http方式
type SmsService interface {
	Send(ctx context.Context, username string, key string, info SMSInfo) (int, string, string)
	Query(ctx context.Context, username string, key string) (float64, error)
}

// 短信实体结构
type SMSInfo struct {
	Sign    string // 短信签名
	Content string // 短信内容
	SubCode string // 短信扩展码
	Type    int    // 短信类型
}

type basicSmsService struct{}

func (b *basicSmsService) Send(ctx context.Context, username string, key string, info SMSInfo) (i0 int, s1 string, s2 string) {
	// TODO implement the business logic of Send
	return i0, s1, s2
}
func (b *basicSmsService) Query(ctx context.Context, username string, key string) (f0 float64, e1 error) {
	// TODO implement the business logic of Query
	return f0, e1
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
