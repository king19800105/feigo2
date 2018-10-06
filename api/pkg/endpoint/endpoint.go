package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service1 "github.com/king19800105/feigo/api/pkg/service"
	service "github.com/king19800105/feigo/sms/pkg/service"
)

// SMSSendRequest collects the request parameters for the SMSSend method.
type SMSSendRequest struct {
	Username string             `json:"username"`
	Key      string             `json:"key"`
	Content  service.SMSContent `json:"content"`
}

// SMSSendResponse collects the response parameters for the SMSSend method.
type SMSSendResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSMSSendEndpoint returns an endpoint that invokes SMSSend on the service.
func MakeSMSSendEndpoint(s service1.ApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SMSSendRequest)
		s0, e1 := s.SMSSend(ctx, req.Username, req.Key, req.Content)
		return SMSSendResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SMSSendResponse) Failed() error {
	return r.E1
}

// SMSQueryRequest collects the request parameters for the SMSQuery method.
type SMSQueryRequest struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}

// SMSQueryResponse collects the response parameters for the SMSQuery method.
type SMSQueryResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSMSQueryEndpoint returns an endpoint that invokes SMSQuery on the service.
func MakeSMSQueryEndpoint(s service1.ApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SMSQueryRequest)
		s0, e1 := s.SMSQuery(ctx, req.Username, req.Key)
		return SMSQueryResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SMSQueryResponse) Failed() error {
	return r.E1
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SMSSend implements Service. Primarily useful in a client.
func (e Endpoints) SMSSend(ctx context.Context, username string, key string, content service.SMSContent) (s0 string, e1 error) {
	request := SMSSendRequest{
		Content:  content,
		Key:      key,
		Username: username,
	}
	response, err := e.SMSSendEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SMSSendResponse).S0, response.(SMSSendResponse).E1
}

// SMSQuery implements Service. Primarily useful in a client.
func (e Endpoints) SMSQuery(ctx context.Context, username string, key string) (s0 string, e1 error) {
	request := SMSQueryRequest{
		Key:      key,
		Username: username,
	}
	response, err := e.SMSQueryEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SMSQueryResponse).S0, response.(SMSQueryResponse).E1
}
