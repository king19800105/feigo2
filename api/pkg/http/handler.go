package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/king19800105/feigo/api/internal"
	"github.com/king19800105/feigo/api/pkg/endpoint"
	"github.com/king19800105/feigo/api/pkg/io"
	"github.com/king19800105/feigo/api/tools"
	http1 "net/http"
)

// makeSMSSendHandler creates the handler logic
func makeSMSSendHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/sms/send").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.SMSSendEndpoint, decodeSMSSendRequest, encodeSMSSendResponse, options...)))
}

// decodeSMSSendResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSMSSendRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := io.SMSSendRequest{}

	if err := tools.DecodeFormData(r, &req); nil != err {
		return req, internal.ErrReqData
	}

	return req, nil
}

// encodeSMSSendResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSMSSendResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(io.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSMSQueryHandler creates the handler logic
func makeSMSQueryHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/sms-query").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.SMSQueryEndpoint, decodeSMSQueryRequest, encodeSMSQueryResponse, options...)))
}

// decodeSMSQueryResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSMSQueryRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := io.SMSQueryRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSMSQueryResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSMSQueryResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(io.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	// todo... 解析json，需要判断一下是Err的子类，还是普通Error的子类。
	// todo... 如果是前者，则对于的传递code，message，result。如果是后者，则给予code，result默认的值。
	// todo... 在tools里面编写同一的responseApi方法，然后再这里调用一下，同一格式，无论是正确还是错误响应。
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
