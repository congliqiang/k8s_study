package request

import (
	"kubeimooc.com/model/base"
	ingroute_k8s "kubeimooc.com/model/ingroute/k8s"
)

//@Author: morris

type IngressRoute struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	ingroute_k8s.IngressRouteSpec
}
