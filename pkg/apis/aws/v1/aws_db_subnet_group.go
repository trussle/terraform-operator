
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbSubnetGroupSpec	`json:"spec"`
}

type AwsDbSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbSubnetGroup	`json:"items"`
}

type AwsDbSubnetGroupSpec struct {
	Description	string	`json:"description"`
	SubnetIds	interface{}	`json:"subnet_ids"`
	Tags	map[string]interface{}	`json:"tags"`
}
