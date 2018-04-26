
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventPermission describes a AwsCloudwatchEventPermission resource
type AwsCloudwatchEventPermission struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchEventPermissionSpec	`json:"spec"`
}


// AwsCloudwatchEventPermissionSpec is the spec for a AwsCloudwatchEventPermission Resource
type AwsCloudwatchEventPermissionSpec struct {
	StatementId	string	`json:"statement_id"`
	Action	string	`json:"action"`
	Principal	string	`json:"principal"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventPermissionList is a list of AwsCloudwatchEventPermission resources
type AwsCloudwatchEventPermissionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventPermission	`json:"items"`
}

