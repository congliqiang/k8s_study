package k8s

import (
	"kubeimooc.com/service"
	"kubeimooc.com/validate"
)

type ApiGroup struct {
	PodApi
	NamespaceApi
	NodeApi
	ConfigMapApi
	SecretApi
	PVApi
	PVCApi
	SCApi
	SvcApi
	IngressApi
	IngRouteApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
var configMapService = service.ServiceGroupApp.ConfigMapServiceGroup.ConfigMapService
var secretService = service.ServiceGroupApp.SecretServiceGroup.SecretService
var pvService = service.ServiceGroupApp.PVServiceGroup.PVService
var pvcService = service.ServiceGroupApp.PVCServiceGroup.PVCService
var scService = service.ServiceGroupApp.SCServiceGroup.SCService
var svcService = service.ServiceGroupApp.SvcServiceGroup.SvcService
var ingressService = service.ServiceGroupApp.IngressServiceGroup.IngressService
var ingRouteService = service.ServiceGroupApp.IngRouteServiceGroup.IngRouteService
