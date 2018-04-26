
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalWebAcl describes a AwsWafregionalWebAcl resource
type AwsWafregionalWebAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalWebAclSpec	`json:"spec"`
}


// AwsWafregionalWebAclSpec is the spec for a AwsWafregionalWebAcl Resource
type AwsWafregionalWebAclSpec struct {
	MetricName	string	`json:"metric_name"`
	Rule	Generic	`json:"rule"`
	Name	string	`json:"name"`
	DefaultAction	[]Generic	`json:"default_action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalWebAclList is a list of AwsWafregionalWebAcl resources
type AwsWafregionalWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalWebAcl	`json:"items"`
}

