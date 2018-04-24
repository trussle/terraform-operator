
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerListenerPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLoadBalancerListenerPolicySpec	`json:"spec"`
}

type AwsLoadBalancerListenerPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLoadBalancerListenerPolicy	`json:"items"`
}

type AwsLoadBalancerListenerPolicySpec struct {
	LoadBalancerName	string	`json:"load_balancer_name"`
	PolicyNames	interface{}	`json:"policy_names"`
	LoadBalancerPort	int	`json:"load_balancer_port"`
}
