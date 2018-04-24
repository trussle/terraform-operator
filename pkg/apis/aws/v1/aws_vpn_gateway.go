
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpnGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnGatewaySpec	`json:"spec"`
}

type AwsVpnGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnGateway	`json:"items"`
}

type AwsVpnGatewaySpec struct {
	AvailabilityZone	string	`json:"availability_zone"`
	Tags	map[string]interface{}	`json:"tags"`
}
