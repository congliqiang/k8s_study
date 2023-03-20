package convert

import "kubeimooc.com/convert/pod"

//@Author: morris

type ConvertGroup struct {
	PodConvert pod.PodConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
