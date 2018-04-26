
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxSubnetGroup describes a AwsDaxSubnetGroup resource
type AwsDaxSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDaxSubnetGroupSpec	`json:"spec"`
}


// AwsDaxSubnetGroupSpec is the spec for a AwsDaxSubnetGroup Resource
type AwsDaxSubnetGroupSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	SubnetIds	string	`json:"subnet_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxSubnetGroupList is a list of AwsDaxSubnetGroup resources
type AwsDaxSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDaxSubnetGroup	`json:"items"`
}

