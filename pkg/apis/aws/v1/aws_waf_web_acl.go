
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafWebAcl describes a AwsWafWebAcl resource
type AwsWafWebAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafWebAclSpec	`json:"spec"`
}


// AwsWafWebAclSpec is the spec for a AwsWafWebAcl Resource
type AwsWafWebAclSpec struct {
	Name	string	`json:"name"`
	DefaultAction	string	`json:"default_action"`
	MetricName	string	`json:"metric_name"`
	Rules	string	`json:"rules"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafWebAclList is a list of AwsWafWebAcl resources
type AwsWafWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafWebAcl	`json:"items"`
}

