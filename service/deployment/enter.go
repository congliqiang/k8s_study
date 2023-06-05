package deployment

import "kubeimooc.com/convert"

//@Author: morris
type ServiceGroup struct {
	DeploymentService
}

var podConvert = convert.ConvertGroupApp.PodConvert
