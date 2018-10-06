package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/king19800105/feigo/sms/pkg/endpoint"
	pb "github.com/king19800105/feigo/sms/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeSendHandler creates the handler logic
func makeSendHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEndpoint, decodeSendRequest, encodeSendResponse, options...)
}

// decodeSendResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeSendRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Sms' Decoder is not impelemented")
}

// encodeSendResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Sms' Encoder is not impelemented")
}
func (g *grpcServer) Send(ctx context1.Context, req *pb.SendRequest) (*pb.SendReply, error) {
	_, rep, err := g.send.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendReply), nil
}
