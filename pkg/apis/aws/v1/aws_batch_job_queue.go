
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBatchJobQueue struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsBatchJobQueueSpec	`json:"spec"`
}

type AwsBatchJobQueueList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchJobQueue	`json:"items"`
}

type AwsBatchJobQueueSpec struct {
	State	string	`json:"state"`
	ComputeEnvironments	[]interface{}	`json:"compute_environments"`
	Name	string	`json:"name"`
	Priority	int	`json:"priority"`
}
