
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogDestination describes a AwsCloudwatchLogDestination resource
type AwsCloudwatchLogDestination struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogDestinationSpec	`json:"spec"`
}


// AwsCloudwatchLogDestinationSpec is the spec for a AwsCloudwatchLogDestination Resource
type AwsCloudwatchLogDestinationSpec struct {
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
	TargetArn	string	`json:"target_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogDestinationList is a list of AwsCloudwatchLogDestination resources
type AwsCloudwatchLogDestinationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogDestination	`json:"items"`
}

