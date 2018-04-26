
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchJobQueue describes a AwsBatchJobQueue resource
type AwsBatchJobQueue struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsBatchJobQueueSpec	`json:"spec"`
}


// AwsBatchJobQueueSpec is the spec for a AwsBatchJobQueue Resource
type AwsBatchJobQueueSpec struct {
	ComputeEnvironments	[]string	`json:"compute_environments"`
	Name	string	`json:"name"`
	Priority	int	`json:"priority"`
	State	string	`json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchJobQueueList is a list of AwsBatchJobQueue resources
type AwsBatchJobQueueList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchJobQueue	`json:"items"`
}

