
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogDestinationPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogDestinationPolicySpec	`json:"spec"`
}

type AwsCloudwatchLogDestinationPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogDestinationPolicy	`json:"items"`
}

type AwsCloudwatchLogDestinationPolicySpec struct {
	DestinationName	string	`json:"destination_name"`
	AccessPolicy	string	`json:"access_policy"`
}
