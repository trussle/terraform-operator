
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsProxyProtocolPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsProxyProtocolPolicySpec	`json:"spec"`
}

type AwsProxyProtocolPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsProxyProtocolPolicy	`json:"items"`
}

type AwsProxyProtocolPolicySpec struct {
	LoadBalancer	string	`json:"load_balancer"`
	InstancePorts	interface{}	`json:"instance_ports"`
}
