
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEgressOnlyInternetGateway describes a AwsEgressOnlyInternetGateway resource
type AwsEgressOnlyInternetGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEgressOnlyInternetGatewaySpec	`json:"spec"`
}


// AwsEgressOnlyInternetGatewaySpec is the spec for a AwsEgressOnlyInternetGateway Resource
type AwsEgressOnlyInternetGatewaySpec struct {
	VpcId	string	`json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEgressOnlyInternetGatewayList is a list of AwsEgressOnlyInternetGateway resources
type AwsEgressOnlyInternetGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEgressOnlyInternetGateway	`json:"items"`
}

