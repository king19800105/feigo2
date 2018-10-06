// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	endpoint "github.com/king19800105/feigo/api/pkg/endpoint"
	http1 "github.com/king19800105/feigo/api/pkg/http"
	service "github.com/king19800105/feigo/api/pkg/service"
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
		"SMSQuery": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SMSQuery", logger))},
		"SMSSend":  {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "SMSSend", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["SMSSend"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SMSSend")), endpoint.InstrumentingMiddleware(duration.With("method", "SMSSend"))}
	mw["SMSQuery"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "SMSQuery")), endpoint.InstrumentingMiddleware(duration.With("method", "SMSQuery"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"SMSSend", "SMSQuery"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
