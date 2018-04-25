
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerListenerPolicy describes a AwsLoadBalancerListenerPolicy resource
type AwsLoadBalancerListenerPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLoadBalancerListenerPolicySpec	`json:"spec"`
}


// AwsLoadBalancerListenerPolicySpec is the spec for a AwsLoadBalancerListenerPolicy Resource
type AwsLoadBalancerListenerPolicySpec struct {
	LoadBalancerName	string	`json:"load_balancer_name"`
	PolicyNames	string	`json:"policy_names"`
	LoadBalancerPort	int	`json:"load_balancer_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerListenerPolicyList is a list of AwsLoadBalancerListenerPolicy resources
type AwsLoadBalancerListenerPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLoadBalancerListenerPolicy	`json:"items"`
}

