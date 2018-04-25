
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPermission describes a AwsOpsworksPermission resource
type AwsOpsworksPermission struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksPermissionSpec	`json:"spec"`
}


// AwsOpsworksPermissionSpec is the spec for a AwsOpsworksPermission Resource
type AwsOpsworksPermissionSpec struct {
	UserArn	string	`json:"user_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPermissionList is a list of AwsOpsworksPermission resources
type AwsOpsworksPermissionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksPermission	`json:"items"`
}

