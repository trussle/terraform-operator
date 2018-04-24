
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultNetworkAcl struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultNetworkAclSpec	`json:"spec"`
}

type AwsDefaultNetworkAclList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultNetworkAcl	`json:"items"`
}

type AwsDefaultNetworkAclSpec struct {
	DefaultNetworkAclId	string	`json:"default_network_acl_id"`
	SubnetIds	interface{}	`json:"subnet_ids"`
	Ingress	interface{}	`json:"ingress"`
	Egress	interface{}	`json:"egress"`
	Tags	map[string]interface{}	`json:"tags"`
}
