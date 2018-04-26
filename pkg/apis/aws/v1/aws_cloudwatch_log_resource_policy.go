
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogResourcePolicy describes a AwsCloudwatchLogResourcePolicy resource
type AwsCloudwatchLogResourcePolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogResourcePolicySpec	`json:"spec"`
}


// AwsCloudwatchLogResourcePolicySpec is the spec for a AwsCloudwatchLogResourcePolicy Resource
type AwsCloudwatchLogResourcePolicySpec struct {
	PolicyName	string	`json:"policy_name"`
	PolicyDocument	string	`json:"policy_document"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogResourcePolicyList is a list of AwsCloudwatchLogResourcePolicy resources
type AwsCloudwatchLogResourcePolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogResourcePolicy	`json:"items"`
}

