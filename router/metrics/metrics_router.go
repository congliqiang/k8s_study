package metrics

import (
	"github.com/gin-gonic/gin"
	"kubeimooc.com/api"
)

//@Author: morris
func initMetricsRouter(group *gin.RouterGroup) {
	metricsApiGroup := api.ApiGroupApp.MetricsApiGroup.MetricsApi
	prometheusApiGroup := api.ApiGroupApp.MetricsApiGroup.PrometheusApi
	group.GET("/dashboard", metricsApiGroup.GetDashboardData)
	group.GET("/prometheus", prometheusApiGroup.GetMetrics)
}
