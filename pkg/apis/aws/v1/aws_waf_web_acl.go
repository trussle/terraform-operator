
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Rules	Generic	`json:"rules"`
	Name	string	`json:"name"`
	DefaultAction	Generic	`json:"default_action"`
	MetricName	string	`json:"metric_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafWebAclList is a list of AwsWafWebAcl resources
type AwsWafWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafWebAcl	`json:"items"`
}

