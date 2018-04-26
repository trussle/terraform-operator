
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGateway describes a AwsVpnGateway resource
type AwsVpnGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpnGatewaySpec	`json:"spec"`
}


// AwsVpnGatewaySpec is the spec for a AwsVpnGateway Resource
type AwsVpnGatewaySpec struct {
	AvailabilityZone	string	`json:"availability_zone"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpnGatewayList is a list of AwsVpnGateway resources
type AwsVpnGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpnGateway	`json:"items"`
}

