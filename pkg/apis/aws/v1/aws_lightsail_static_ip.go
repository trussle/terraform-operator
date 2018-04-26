
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailStaticIp describes a AwsLightsailStaticIp resource
type AwsLightsailStaticIp struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailStaticIpSpec	`json:"spec"`
}


// AwsLightsailStaticIpSpec is the spec for a AwsLightsailStaticIp Resource
type AwsLightsailStaticIpSpec struct {
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailStaticIpList is a list of AwsLightsailStaticIp resources
type AwsLightsailStaticIpList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailStaticIp	`json:"items"`
}

