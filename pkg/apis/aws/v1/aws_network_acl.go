
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Tags	map[string]interface{}	`json:"tags"`
	VpcId	string	`json:"vpc_id"`
	SubnetId	string	`json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkAclList is a list of AwsNetworkAcl resources
type AwsNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkAcl	`json:"items"`
}

