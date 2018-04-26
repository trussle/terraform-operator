
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListener describes a AwsAlbListener resource
type AwsAlbListener struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbListenerSpec	`json:"spec"`
}


// AwsAlbListenerSpec is the spec for a AwsAlbListener Resource
type AwsAlbListenerSpec struct {
	Protocol	string	`json:"protocol"`
	CertificateArn	string	`json:"certificate_arn"`
	DefaultAction	[]Generic	`json:"default_action"`
	LoadBalancerArn	string	`json:"load_balancer_arn"`
	Port	int	`json:"port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerList is a list of AwsAlbListener resources
type AwsAlbListenerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListener	`json:"items"`
}

