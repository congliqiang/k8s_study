package request

import (
	batchv1 "k8s.io/api/batch/v1"
	"kubeimooc.com/model/base"
	pod_req "kubeimooc.com/model/pod/request"
)

//@Author: morris
type CronJobBase struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	//cron表达式
	Schedule string `json:"schedule"`
	//是否暂停cronjob
	Suspend bool `json:"suspend"`
	//并发策略
	ConcurrencyPolicy          batchv1.ConcurrencyPolicy `json:"concurrencyPolicy"`
	SuccessfulJobsHistoryLimit int32                     `json:"successfulJobsHistoryLimit"`
	FailedJobsHistoryLimit     int32                     `json:"failedJobsHistoryLimit"`
	Selector                   []base.ListMapItem        `json:"selector"`
	JobBase                    JobBase                   `json:"jobBase"`
}
type JobBase struct {
	Completions  int32 `json:"completions"`
	BackoffLimit int32 `json:"backoffLimit"`
}
type CronJob struct {
	Base     CronJobBase `json:"base"`
	Template pod_req.Pod `json:"template"`
}
