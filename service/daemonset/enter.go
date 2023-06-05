package daemonset

import "kubeimooc.com/convert"

//@Author: morris
type ServiceGroup struct {
	DaemonSetService
}

var podConvert = convert.ConvertGroupApp.PodConvert
