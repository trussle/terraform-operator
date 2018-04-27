
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListener describes a AwsLbListener resource
type AwsLbListener struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbListenerSpec	`json:"spec"`
}


// AwsLbListenerSpec is the spec for a AwsLbListener Resource
type AwsLbListenerSpec struct {
	LoadBalancerArn	string	`json:"load_balancer_arn"`
	Port	int	`json:"port"`
	Protocol	string	`json:"protocol"`
	CertificateArn	string	`json:"certificate_arn"`
	DefaultAction	[]AwsLbListenerDefaultAction	`json:"default_action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerList is a list of AwsLbListener resources
type AwsLbListenerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbListener	`json:"items"`
}


// AwsLbListenerDefaultAction is a AwsLbListenerDefaultAction
type AwsLbListenerDefaultAction struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	Type	string	`json:"type"`
}
