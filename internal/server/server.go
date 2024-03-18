package server

import (
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

var (
	_metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "server requests duration(sec).",
		Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})

	_metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_code_total",
		Help: "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
)

func init() {
	prometheus.MustRegister(_metricSeconds, _metricRequests)
}
