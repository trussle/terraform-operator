
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListener struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbListenerSpec	`json:"spec"`
}

type AwsAlbListenerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListener	`json:"items"`
}

type AwsAlbListenerSpec struct {
	DefaultAction	[]interface{}	`json:"default_action"`
	LoadBalancerArn	string	`json:"load_balancer_arn"`
	Port	int	`json:"port"`
	Protocol	string	`json:"protocol"`
	CertificateArn	string	`json:"certificate_arn"`
}
