
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultNetworkAcl describes a AwsDefaultNetworkAcl resource
type AwsDefaultNetworkAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultNetworkAclSpec	`json:"spec"`
}


// AwsDefaultNetworkAclSpec is the spec for a AwsDefaultNetworkAcl Resource
type AwsDefaultNetworkAclSpec struct {
	DefaultNetworkAclId	string	`json:"default_network_acl_id"`
	SubnetIds	string	`json:"subnet_ids"`
	Ingress	string	`json:"ingress"`
	Egress	string	`json:"egress"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultNetworkAclList is a list of AwsDefaultNetworkAcl resources
type AwsDefaultNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultNetworkAcl	`json:"items"`
}

