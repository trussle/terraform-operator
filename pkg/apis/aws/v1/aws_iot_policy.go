
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotPolicy describes a AwsIotPolicy resource
type AwsIotPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIotPolicySpec	`json:"spec"`
}


// AwsIotPolicySpec is the spec for a AwsIotPolicy Resource
type AwsIotPolicySpec struct {
	Name	string	`json:"name"`
	Policy	string	`json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotPolicyList is a list of AwsIotPolicy resources
type AwsIotPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIotPolicy	`json:"items"`
}

