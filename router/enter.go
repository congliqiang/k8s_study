package router

import (
	"kubeimooc.com/router/example"
	"kubeimooc.com/router/harbor"
	"kubeimooc.com/router/k8s"
	"kubeimooc.com/router/metrics"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	K8SRouterGroup     k8s.K8sRouter
	HarborRouterGroup  harbor.HarborRouter
	MetricsRouterGroup metrics.MetricsRouter
}

var RouterGroupApp = new(RouterGroup)
