
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNatGateway describes a AwsNatGateway resource
type AwsNatGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNatGatewaySpec	`json:"spec"`
}


// AwsNatGatewaySpec is the spec for a AwsNatGateway Resource
type AwsNatGatewaySpec struct {
	Tags	map[string]string	`json:"tags"`
	AllocationId	string	`json:"allocation_id"`
	SubnetId	string	`json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNatGatewayList is a list of AwsNatGateway resources
type AwsNatGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNatGateway	`json:"items"`
}

