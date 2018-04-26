
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbSslNegotiationPolicy describes a AwsLbSslNegotiationPolicy resource
type AwsLbSslNegotiationPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbSslNegotiationPolicySpec	`json:"spec"`
}


// AwsLbSslNegotiationPolicySpec is the spec for a AwsLbSslNegotiationPolicy Resource
type AwsLbSslNegotiationPolicySpec struct {
	Attribute	Generic	`json:"attribute"`
	Name	string	`json:"name"`
	LoadBalancer	string	`json:"load_balancer"`
	LbPort	int	`json:"lb_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbSslNegotiationPolicyList is a list of AwsLbSslNegotiationPolicy resources
type AwsLbSslNegotiationPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbSslNegotiationPolicy	`json:"items"`
}

