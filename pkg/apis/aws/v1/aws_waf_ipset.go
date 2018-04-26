
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafIpset describes a AwsWafIpset resource
type AwsWafIpset struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafIpsetSpec	`json:"spec"`
}


// AwsWafIpsetSpec is the spec for a AwsWafIpset Resource
type AwsWafIpsetSpec struct {
	IpSetDescriptors	IpSetDescriptors	`json:"ip_set_descriptors"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafIpsetList is a list of AwsWafIpset resources
type AwsWafIpsetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafIpset	`json:"items"`
}


// IpSetDescriptors is a IpSetDescriptors
type IpSetDescriptors struct {
	Type	string	`json:"type"`
	Value	string	`json:"value"`
}
