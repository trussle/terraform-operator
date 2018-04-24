
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalWebAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalWebAclSpec	`json:"spec"`
}

type AwsWafregionalWebAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalWebAcl	`json:"items"`
}

type AwsWafregionalWebAclSpec struct {
	DefaultAction	[]interface{}	`json:"default_action"`
	MetricName	string	`json:"metric_name"`
	Rule	interface{}	`json:"rule"`
	Name	string	`json:"name"`
}
