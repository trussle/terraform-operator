
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogDestinationPolicy describes a AwsCloudwatchLogDestinationPolicy resource
type AwsCloudwatchLogDestinationPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogDestinationPolicySpec	`json:"spec"`
}


// AwsCloudwatchLogDestinationPolicySpec is the spec for a AwsCloudwatchLogDestinationPolicy Resource
type AwsCloudwatchLogDestinationPolicySpec struct {
	AccessPolicy	string	`json:"access_policy"`
	DestinationName	string	`json:"destination_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogDestinationPolicyList is a list of AwsCloudwatchLogDestinationPolicy resources
type AwsCloudwatchLogDestinationPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogDestinationPolicy	`json:"items"`
}

