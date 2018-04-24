
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNatGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNatGatewaySpec	`json:"spec"`
}

type AwsNatGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNatGateway	`json:"items"`
}

type AwsNatGatewaySpec struct {
	AllocationId	string	`json:"allocation_id"`
	SubnetId	string	`json:"subnet_id"`
	Tags	map[string]interface{}	`json:"tags"`
}
