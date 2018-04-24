
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkAclSpec	`json:"spec"`
}

type AwsNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkAcl	`json:"items"`
}

type AwsNetworkAclSpec struct {
	VpcId	string	`json:"vpc_id"`
	SubnetId	string	`json:"subnet_id"`
	Tags	map[string]interface{}	`json:"tags"`
}
