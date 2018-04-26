
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteTable describes a AwsRouteTable resource
type AwsRouteTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRouteTableSpec	`json:"spec"`
}


// AwsRouteTableSpec is the spec for a AwsRouteTable Resource
type AwsRouteTableSpec struct {
	Tags	map[string]???	`json:"tags"`
	VpcId	string	`json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRouteTableList is a list of AwsRouteTable resources
type AwsRouteTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRouteTable	`json:"items"`
}

