package metrics

import "kubeimooc.com/service"

//@Author: morris

type ApiGroup struct {
	MetricsApi
	PrometheusApi
}

var metricsService = service.ServiceGroupApp.MetricsServiceGroup.MetricsService
