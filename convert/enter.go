package convert

import (
	"kubeimooc.com/convert/node"
	"kubeimooc.com/convert/pod"
)

//@Author: morris

type ConvertGroup struct {
	PodConvert  pod.PodConvertGroup
	NodeConvert node.Group
}

var ConvertGroupApp = new(ConvertGroup)
