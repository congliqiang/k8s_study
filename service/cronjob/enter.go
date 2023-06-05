package cronjob

import "kubeimooc.com/convert"

//@Author: morris
type ServiceGroup struct {
	CronJobSetService
}

var podConvert = convert.ConvertGroupApp.PodConvert
