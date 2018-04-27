
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamServiceLinkedRole describes a AwsIamServiceLinkedRole resource
type AwsIamServiceLinkedRole struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamServiceLinkedRoleSpec	`json:"spec"`
}


// AwsIamServiceLinkedRoleSpec is the spec for a AwsIamServiceLinkedRole Resource
type AwsIamServiceLinkedRoleSpec struct {
	CustomSuffix	string	`json:"custom_suffix"`
	Description	string	`json:"description"`
	AwsServiceName	string	`json:"aws_service_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamServiceLinkedRoleList is a list of AwsIamServiceLinkedRole resources
type AwsIamServiceLinkedRoleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamServiceLinkedRole	`json:"items"`
}

