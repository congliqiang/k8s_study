package request

import (
	"kubeimooc.com/model/base"
	pod_req "kubeimooc.com/model/pod/request"
	pvc_req "kubeimooc.com/model/pvc/request"
)

//@Author: morris

type StatefulSetBase struct {
	Name                 string                          `json:"name"`
	Namespace            string                          `json:"namespace"`
	Replicas             int32                           `json:"replicas"`
	Labels               []base.ListMapItem              `json:"labels"`
	Selector             []base.ListMapItem              `json:"selector"`
	ServiceName          string                          `json:"serviceName"`
	VolumeClaimTemplates []pvc_req.PersistentVolumeClaim `json:"volumeClaimTemplates"`
}
type StatefulSet struct {
	Base     StatefulSetBase `json:"base"`
	Template pod_req.Pod     `json:"template"`
}
