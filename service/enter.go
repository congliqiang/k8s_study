package service

import (
	"kubeimooc.com/service/configmap"
	"kubeimooc.com/service/ingress"
	"kubeimooc.com/service/ingroute"
	"kubeimooc.com/service/node"
	"kubeimooc.com/service/pod"
	"kubeimooc.com/service/pv"
	"kubeimooc.com/service/pvc"
	"kubeimooc.com/service/sc"
	"kubeimooc.com/service/secret"
	"kubeimooc.com/service/svc"
)

//@Author: morris
type ServiceGroup struct {
	PodServiceGroup       pod.PodServiceGroup
	NodeServiceGroup      node.Group
	ConfigMapServiceGroup configmap.ServiceGroup
	SecretServiceGroup    secret.SeviceGroup
	PVServiceGroup        pv.ServiceGroup
	PVCServiceGroup       pvc.ServiceGroup
	SCServiceGroup        sc.SCServiceGroup
	SvcServiceGroup       svc.ServiceGroup
	IngressServiceGroup   ingress.ServiceGroup
	IngRouteServiceGroup  ingroute.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
