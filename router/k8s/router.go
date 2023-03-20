package k8s

import (
	"github.com/gin-gonic/gin"
	"kubeimooc.com/api"
)

type K8sRouter struct {
}

func (*K8sRouter) InitK8SRouter(r *gin.Engine) {
	group := r.Group("/k8s")
	apiGroup := api.ApiGroupApp.K8SApiGroup
	group.GET("/listPod", apiGroup.GetPodList)
}
