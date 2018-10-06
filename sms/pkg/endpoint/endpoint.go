package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/king19800105/feigo/sms/pkg/service"
)

// SendRequest collects the request parameters for the Send method.
type SendRequest struct {
	Content service.SMSContent `json:"content"`
}

// SendResponse collects the response parameters for the Send method.
type SendResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSendEndpoint returns an endpoint that invokes Send on the service.
func MakeSendEndpoint(s service.SmsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendRequest)
		s0, e1 := s.Send(ctx, req.Content)
		return SendResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SendResponse) Failed() error {
	return r.E1
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Send implements Service. Primarily useful in a client.
func (e Endpoints) Send(ctx context.Context, content service.SMSContent) (s0 string, e1 error) {
	request := SendRequest{Content: content}
	response, err := e.SendEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendResponse).S0, response.(SendResponse).E1
}
