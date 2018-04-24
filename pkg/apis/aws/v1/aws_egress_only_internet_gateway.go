
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEgressOnlyInternetGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEgressOnlyInternetGatewaySpec	`json:"spec"`
}

type AwsEgressOnlyInternetGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEgressOnlyInternetGateway	`json:"items"`
}

type AwsEgressOnlyInternetGatewaySpec struct {
	VpcId	string	`json:"vpc_id"`
}
