
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheSubnetGroup describes a AwsElasticacheSubnetGroup resource
type AwsElasticacheSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheSubnetGroupSpec	`json:"spec"`
}


// AwsElasticacheSubnetGroupSpec is the spec for a AwsElasticacheSubnetGroup Resource
type AwsElasticacheSubnetGroupSpec struct {
	Description	string	`json:"description"`
	Name	string	`json:"name"`
	SubnetIds	string	`json:"subnet_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheSubnetGroupList is a list of AwsElasticacheSubnetGroup resources
type AwsElasticacheSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheSubnetGroup	`json:"items"`
}

