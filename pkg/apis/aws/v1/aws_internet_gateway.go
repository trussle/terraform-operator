
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInternetGateway describes a AwsInternetGateway resource
type AwsInternetGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInternetGatewaySpec	`json:"spec"`
}


// AwsInternetGatewaySpec is the spec for a AwsInternetGateway Resource
type AwsInternetGatewaySpec struct {
	Tags	map[string]string	`json:"tags"`
	VpcId	string	`json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInternetGatewayList is a list of AwsInternetGateway resources
type AwsInternetGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInternetGateway	`json:"items"`
}

