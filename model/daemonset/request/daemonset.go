package request

import (
	"kubeimooc.com/model/base"
	pod_req "kubeimooc.com/model/pod/request"
)

//@Author: morris
type DaemonSetBase struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	Selector  []base.ListMapItem `json:"selector"`
}
type DaemonSet struct {
	Base     DaemonSetBase `json:"base"`
	Template pod_req.Pod   `json:"template"`
}
