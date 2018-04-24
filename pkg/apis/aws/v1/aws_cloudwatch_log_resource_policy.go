
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogResourcePolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogResourcePolicySpec	`json:"spec"`
}

type AwsCloudwatchLogResourcePolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogResourcePolicy	`json:"items"`
}

type AwsCloudwatchLogResourcePolicySpec struct {
	PolicyName	string	`json:"policy_name"`
	PolicyDocument	string	`json:"policy_document"`
}
