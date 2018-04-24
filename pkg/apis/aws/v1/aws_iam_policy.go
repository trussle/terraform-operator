
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamPolicySpec	`json:"spec"`
}

type AwsIamPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamPolicy	`json:"items"`
}

type AwsIamPolicySpec struct {
	Description	string	`json:"description"`
	Path	string	`json:"path"`
	Policy	string	`json:"policy"`
	NamePrefix	string	`json:"name_prefix"`
}
