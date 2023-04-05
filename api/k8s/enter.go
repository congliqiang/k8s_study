package k8s

import (
	"kubeimooc.com/service"
	"kubeimooc.com/validate"
)

type ApiGroup struct {
	PodApi
	NamespaceApi
	NodeApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
