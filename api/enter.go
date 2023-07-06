package api

import (
	"kubeimooc.com/api/example"
	"kubeimooc.com/api/harbor"
	"kubeimooc.com/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8SApiGroup     k8s.ApiGroup
	HarborApiGroup  harbor.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
