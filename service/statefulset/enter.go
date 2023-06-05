package statefulset

import "kubeimooc.com/convert"

//@Author: morris
type ServiceGroup struct {
	StatefulSetService
}

var podConvert = convert.ConvertGroupApp.PodConvert
