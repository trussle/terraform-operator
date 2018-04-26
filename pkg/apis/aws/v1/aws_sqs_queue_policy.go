
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSqsQueuePolicy describes a AwsSqsQueuePolicy resource
type AwsSqsQueuePolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSqsQueuePolicySpec	`json:"spec"`
}


// AwsSqsQueuePolicySpec is the spec for a AwsSqsQueuePolicy Resource
type AwsSqsQueuePolicySpec struct {
	QueueUrl	string	`json:"queue_url"`
	Policy	string	`json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSqsQueuePolicyList is a list of AwsSqsQueuePolicy resources
type AwsSqsQueuePolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSqsQueuePolicy	`json:"items"`
}

