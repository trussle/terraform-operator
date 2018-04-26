
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbSslNegotiationPolicy describes a AwsLbSslNegotiationPolicy resource
type AwsLbSslNegotiationPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbSslNegotiationPolicySpec	`json:"spec"`
}


// AwsLbSslNegotiationPolicySpec is the spec for a AwsLbSslNegotiationPolicy Resource
type AwsLbSslNegotiationPolicySpec struct {
	LoadBalancer	string	`json:"load_balancer"`
	LbPort	int	`json:"lb_port"`
	Attribute	Attribute	`json:"attribute"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbSslNegotiationPolicyList is a list of AwsLbSslNegotiationPolicy resources
type AwsLbSslNegotiationPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbSslNegotiationPolicy	`json:"items"`
}


// Attribute is a Attribute
type Attribute struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
