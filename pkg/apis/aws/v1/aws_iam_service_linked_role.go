
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamServiceLinkedRole struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamServiceLinkedRoleSpec	`json:"spec"`
}

type AwsIamServiceLinkedRoleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamServiceLinkedRole	`json:"items"`
}

type AwsIamServiceLinkedRoleSpec struct {
	AwsServiceName	string	`json:"aws_service_name"`
	CustomSuffix	string	`json:"custom_suffix"`
	Description	string	`json:"description"`
}
