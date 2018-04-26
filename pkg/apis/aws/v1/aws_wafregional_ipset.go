
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalIpset describes a AwsWafregionalIpset resource
type AwsWafregionalIpset struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalIpsetSpec	`json:"spec"`
}


// AwsWafregionalIpsetSpec is the spec for a AwsWafregionalIpset Resource
type AwsWafregionalIpsetSpec struct {
	Name	string	`json:"name"`
	IpSetDescriptor	IpSetDescriptor	`json:"ip_set_descriptor"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalIpsetList is a list of AwsWafregionalIpset resources
type AwsWafregionalIpsetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalIpset	`json:"items"`
}


// IpSetDescriptor is a IpSetDescriptor
type IpSetDescriptor struct {
	Type	string	`json:"type"`
	Value	string	`json:"value"`
}
