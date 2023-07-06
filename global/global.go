package global

import (
	"k8s.io/client-go/kubernetes"
	"kubeimooc.com/config"
	"kubeimooc.com/plugins/harbor"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
	HarborClient  *harbor.Harbor
)
