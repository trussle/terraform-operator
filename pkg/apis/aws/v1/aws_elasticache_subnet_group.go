
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheSubnetGroupSpec	`json:"spec"`
}

type AwsElasticacheSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheSubnetGroup	`json:"items"`
}

type AwsElasticacheSubnetGroupSpec struct {
	Description	string	`json:"description"`
	Name	string	`json:"name"`
	SubnetIds	interface{}	`json:"subnet_ids"`
}
