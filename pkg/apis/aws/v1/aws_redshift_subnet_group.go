
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftSubnetGroup describes a AwsRedshiftSubnetGroup resource
type AwsRedshiftSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftSubnetGroupSpec	`json:"spec"`
}


// AwsRedshiftSubnetGroupSpec is the spec for a AwsRedshiftSubnetGroup Resource
type AwsRedshiftSubnetGroupSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	SubnetIds	string	`json:"subnet_ids"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftSubnetGroupList is a list of AwsRedshiftSubnetGroup resources
type AwsRedshiftSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftSubnetGroup	`json:"items"`
}

