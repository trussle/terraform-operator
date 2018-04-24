
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventPermission struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchEventPermissionSpec	`json:"spec"`
}

type AwsCloudwatchEventPermissionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventPermission	`json:"items"`
}

type AwsCloudwatchEventPermissionSpec struct {
	Action	string	`json:"action"`
	Principal	string	`json:"principal"`
	StatementId	string	`json:"statement_id"`
}
