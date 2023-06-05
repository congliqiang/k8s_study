package k8s

import (
	"github.com/gin-gonic/gin"
	"kubeimooc.com/api"
)

//@Author: morris
func initCronJobRouter(group *gin.RouterGroup) {
	cronJobApiGroup := api.ApiGroupApp.K8SApiGroup.CronJobApi
	group.POST("/cronjob", cronJobApiGroup.CreateOrUpdateCronJob)
	group.GET("/cronjob/:namespace", cronJobApiGroup.GetCronJobDetailOrList)
	group.DELETE("/cronjob/:namespace/:name", cronJobApiGroup.DeleteCronJob)
}
