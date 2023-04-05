package service

import (
	"kubeimooc.com/service/node"
	"kubeimooc.com/service/pod"
)

//@Author: morris
type ServiceGroup struct {
	PodServiceGroup  pod.PodServiceGroup
	NodeServiceGroup node.Group
}

var ServiceGroupApp = new(ServiceGroup)
