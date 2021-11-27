package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PromResponseOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "invest_api_http_response_200",
		Help: "Total number of 200 requests",
	})
	PromResponseBadRequest = promauto.NewCounter(prometheus.CounterOpts{
		Name: "invest_api_http_response_400",
		Help: "Total number of 400 requests",
	})
	PromResponseUnauthorized = promauto.NewCounter(prometheus.CounterOpts{
		Name: "invest_api_http_response_401",
		Help: "Total number of 401 requests",
	})
	PromResponseForbidden = promauto.NewCounter(prometheus.CounterOpts{
		Name: "invest_api_http_response_403",
		Help: "Total number of 403 requests",
	})
	PromResponseNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "invest_api_http_response_404",
		Help: "Total number of 404 requests",
	})
	PromResponseInternalError = promauto.NewCounter(prometheus.CounterOpts{
		Name: "invest_api_http_response_500",
		Help: "Total number of 500 requests",
	})
)

