
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Port	int	`json:"port"`
	Protocol	string	`json:"protocol"`
	CertificateArn	string	`json:"certificate_arn"`
	DefaultAction	[]IdXaKUaq	`json:"default_action"`
	LoadBalancerArn	string	`json:"load_balancer_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerList is a list of AwsAlbListener resources
type AwsAlbListenerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListener	`json:"items"`
}


// IdXaKUaq is a IdXaKUaq
type IdXaKUaq struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	Type	string	`json:"type"`
}
