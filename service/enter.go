package service

import "kubeimooc.com/service/pod"

//@Author: morris
type ServiceGroup struct {
	PodServiceGroup pod.PodServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
