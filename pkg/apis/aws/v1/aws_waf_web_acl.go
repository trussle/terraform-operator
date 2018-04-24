
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafWebAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafWebAclSpec	`json:"spec"`
}

type AwsWafWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafWebAcl	`json:"items"`
}

type AwsWafWebAclSpec struct {
	Name	string	`json:"name"`
	DefaultAction	interface{}	`json:"default_action"`
	MetricName	string	`json:"metric_name"`
	Rules	interface{}	`json:"rules"`
}
