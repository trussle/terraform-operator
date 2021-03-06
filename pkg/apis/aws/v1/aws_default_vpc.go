
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultVpc describes a AwsDefaultVpc resource
type AwsDefaultVpc struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultVpcSpec	`json:"spec"`
}


// AwsDefaultVpcSpec is the spec for a AwsDefaultVpc Resource
type AwsDefaultVpcSpec struct {
	EnableDnsSupport	bool	`json:"enable_dns_support"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultVpcList is a list of AwsDefaultVpc resources
type AwsDefaultVpcList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultVpc	`json:"items"`
}

