
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbSslNegotiationPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbSslNegotiationPolicySpec	`json:"spec"`
}

type AwsLbSslNegotiationPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbSslNegotiationPolicy	`json:"items"`
}

type AwsLbSslNegotiationPolicySpec struct {
	Name	string	`json:"name"`
	LoadBalancer	string	`json:"load_balancer"`
	LbPort	int	`json:"lb_port"`
	Attribute	interface{}	`json:"attribute"`
}
