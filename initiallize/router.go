package initiallize

import (
	"github.com/gin-gonic/gin"
	"kubeimooc.com/middleware"
	"kubeimooc.com/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors)
	examplGroup := router.RouterGroupApp.ExampleRouterGroup
	k8sGroup := router.RouterGroupApp.K8SRouterGroup
	harborRouterGroup := router.RouterGroupApp.HarborRouterGroup
	examplGroup.InitExample(r)
	k8sGroup.InitK8SRouter(r)
	harborRouterGroup.InitHarborRouter(r)
	return r
	//r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
