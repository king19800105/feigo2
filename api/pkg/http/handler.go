package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "github.com/go-kit/kit/transport/http"
	endpoint "github.com/king19800105/feigo/api/pkg/endpoint"
	"net/http"
)

// makeSMSSendHandler creates the handler logic
func makeSMSSendHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/smssend", http1.NewServer(endpoints.SMSSendEndpoint, decodeSMSSendRequest, encodeSMSSendResponse, options...))
}

// decodeSMSSendResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSMSSendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SMSSendRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSMSSendResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSMSSendResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSMSQueryHandler creates the handler logic
func makeSMSQueryHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/smsquery", http1.NewServer(endpoints.SMSQueryEndpoint, decodeSMSQueryRequest, encodeSMSQueryResponse, options...))
}

// decodeSMSQueryResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSMSQueryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SMSQueryRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSMSQueryResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSMSQueryResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
