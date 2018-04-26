
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkAcl describes a AwsNetworkAcl resource
type AwsNetworkAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkAclSpec	`json:"spec"`
}


// AwsNetworkAclSpec is the spec for a AwsNetworkAcl Resource
type AwsNetworkAclSpec struct {
	VpcId	string	`json:"vpc_id"`
	SubnetId	string	`json:"subnet_id"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkAclList is a list of AwsNetworkAcl resources
type AwsNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkAcl	`json:"items"`
}

