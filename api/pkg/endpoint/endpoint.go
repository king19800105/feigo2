package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/king19800105/feigo/api/pkg/io"
	service1 "github.com/king19800105/feigo/api/pkg/service"
)

// MakeSMSSendEndpoint returns an endpoint that invokes SMSSend on the service.
func MakeSMSSendEndpoint(s service1.ApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(io.SMSSendRequest)
		s0, e1 := s.SMSSend(ctx, req)
		return io.ApiResponse{
			Err:    e1,
			Result: s0,
		}, nil
	}
}

// MakeSMSQueryEndpoint returns an endpoint that invokes SMSQuery on the service.
func MakeSMSQueryEndpoint(s service1.ApiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(io.SMSQueryRequest)
		s0, e1 := s.SMSQuery(ctx, req.Username, req.Key)
		return io.ApiResponse{
			Err:    e1,
			Result: s0,
		}, nil
	}
}

// SMSSend implements Service. Primarily useful in a client.
func (e Endpoints) SMSSend(ctx context.Context, r io.SMSSendRequest) (s0 string, e1 error) {
	request := io.SMSSendRequest{
		Key:           r.Key,
		Username:      r.Username,
		Content:       r.Content,
		ExtensionCode: r.ExtensionCode,
	}
	response, err := e.SMSSendEndpoint(ctx, request)

	if err != nil {
		return
	}

	return response.(io.ApiResponse).Result, response.(io.ApiResponse).Err
}

// SMSQuery implements Service. Primarily useful in a client.
func (e Endpoints) SMSQuery(ctx context.Context, username string, key string) (s0 string, e1 error) {
	request := io.SMSQueryRequest{
		Key:      key,
		Username: username,
	}
	response, err := e.SMSQueryEndpoint(ctx, request)

	if err != nil {
		return
	}

	return response.(io.ApiResponse).Result, response.(io.ApiResponse).Err
}
