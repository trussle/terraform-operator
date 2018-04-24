
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListener struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbListenerSpec	`json:"spec"`
}

type AwsLbListenerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbListener	`json:"items"`
}

type AwsLbListenerSpec struct {
	Port	int	`json:"port"`
	Protocol	string	`json:"protocol"`
	CertificateArn	string	`json:"certificate_arn"`
	DefaultAction	[]interface{}	`json:"default_action"`
	LoadBalancerArn	string	`json:"load_balancer_arn"`
}
