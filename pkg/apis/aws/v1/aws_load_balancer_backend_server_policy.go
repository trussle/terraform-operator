
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerBackendServerPolicy describes a AwsLoadBalancerBackendServerPolicy resource
type AwsLoadBalancerBackendServerPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLoadBalancerBackendServerPolicySpec	`json:"spec"`
}


// AwsLoadBalancerBackendServerPolicySpec is the spec for a AwsLoadBalancerBackendServerPolicy Resource
type AwsLoadBalancerBackendServerPolicySpec struct {
	LoadBalancerName	string	`json:"load_balancer_name"`
	PolicyNames	Generic	`json:"policy_names"`
	InstancePort	int	`json:"instance_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLoadBalancerBackendServerPolicyList is a list of AwsLoadBalancerBackendServerPolicy resources
type AwsLoadBalancerBackendServerPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLoadBalancerBackendServerPolicy	`json:"items"`
}

