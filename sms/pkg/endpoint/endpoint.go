package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/king19800105/feigo/sms/pkg/service"
)

// SendRequest collects the request parameters for the Send method.
type SendRequest struct {
	Username string          `json:"username"`
	Key      string          `json:"key"`
	Info     service.SMSInfo `json:"info"`
}

// SendResponse collects the response parameters for the Send method.
type SendResponse struct {
	I0 int    `json:"i0"`
	S1 string `json:"s1"`
	S2 string `json:"s2"`
}

// MakeSendEndpoint returns an endpoint that invokes Send on the service.
func MakeSendEndpoint(s service.SmsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		i0, s1, s2 := s.Send(ctx, req.Username, req.Key, req.Info)
		return SendResponse{
			I0: i0,
			S1: s1,
			S2: s2,
		}, nil
	}
}

// QueryRequest collects the request parameters for the Query method.
type QueryRequest struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}

// QueryResponse collects the response parameters for the Query method.
type QueryResponse struct {
	F0 float64 `json:"f0"`
	E1 error   `json:"e1"`
}

// MakeQueryEndpoint returns an endpoint that invokes Query on the service.
func MakeQueryEndpoint(s service.SmsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryRequest)
		f0, e1 := s.Query(ctx, req.Username, req.Key)
		return QueryResponse{
			E1: e1,
			F0: f0,
		}, nil
	}
}

// Failed implements Failer.
func (r QueryResponse) Failed() error {
	return r.E1
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Send implements Service. Primarily useful in a client.
func (e Endpoints) Send(ctx context.Context, username string, key string, info service.SMSInfo) (i0 int, s1 string, s2 string) {
	request := SendRequest{
		Info:     info,
		Key:      key,
		Username: username,
	}
	response, err := e.SendEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendResponse).I0, response.(SendResponse).S1, response.(SendResponse).S2
}

// Query implements Service. Primarily useful in a client.
func (e Endpoints) Query(ctx context.Context, username string, key string) (f0 float64, e1 error) {
	request := QueryRequest{
		Key:      key,
		Username: username,
	}
	response, err := e.QueryEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(QueryResponse).F0, response.(QueryResponse).E1
}
