package request

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"kubeimooc.com/model/base"
)

//@Author: morris
type ServiceAccount struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
}

type Role struct {
	Name      string              `json:"name"`
	Namespace string              `json:"namespace"`
	Labels    []base.ListMapItem  `json:"labels"`
	Rules     []rbacv1.PolicyRule `json:"rules"`
}
type RoleBinding struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	//账号
	Subjects []ServiceAccount `json:"subjects"`
	//角色
	RoleRef string `json:"roleRef"`
}
