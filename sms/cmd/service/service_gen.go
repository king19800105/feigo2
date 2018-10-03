// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/king19800105/feigo/sms/pkg/endpoint"
	http1 "github.com/king19800105/feigo/sms/pkg/http"
	service "github.com/king19800105/feigo/sms/pkg/service"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Query": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Query", logger))},
		"Send":  {http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Send", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["Send"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Send")), endpoint.InstrumentingMiddleware(duration.With("method", "Send"))}
	mw["Query"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Query")), endpoint.InstrumentingMiddleware(duration.With("method", "Query"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Send", "Query"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
