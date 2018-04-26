
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopicPolicy describes a AwsSnsTopicPolicy resource
type AwsSnsTopicPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnsTopicPolicySpec	`json:"spec"`
}


// AwsSnsTopicPolicySpec is the spec for a AwsSnsTopicPolicy Resource
type AwsSnsTopicPolicySpec struct {
	Arn	string	`json:"arn"`
	Policy	string	`json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnsTopicPolicyList is a list of AwsSnsTopicPolicy resources
type AwsSnsTopicPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnsTopicPolicy	`json:"items"`
}

