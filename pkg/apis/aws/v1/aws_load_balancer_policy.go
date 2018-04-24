
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLoadBalancerPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLoadBalancerPolicySpec	`json:"spec"`
}

type AwsLoadBalancerPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLoadBalancerPolicy	`json:"items"`
}

type AwsLoadBalancerPolicySpec struct {
	LoadBalancerName	string	`json:"load_balancer_name"`
	PolicyName	string	`json:"policy_name"`
	PolicyTypeName	string	`json:"policy_type_name"`
	PolicyAttribute	interface{}	`json:"policy_attribute"`
}
