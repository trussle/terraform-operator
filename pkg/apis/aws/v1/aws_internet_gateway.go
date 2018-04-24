
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInternetGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInternetGatewaySpec	`json:"spec"`
}

type AwsInternetGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInternetGateway	`json:"items"`
}

type AwsInternetGatewaySpec struct {
	VpcId	string	`json:"vpc_id"`
	Tags	map[string]interface{}	`json:"tags"`
}
