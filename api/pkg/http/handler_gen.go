// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/king19800105/feigo/api/pkg/endpoint"
	http1 "net/http"
)

//  NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makeSMSSendHandler(m, endpoints, options["SMSSend"])
	makeSMSQueryHandler(m, endpoints, options["SMSQuery"])
	return m
}