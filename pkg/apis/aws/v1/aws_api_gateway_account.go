
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayAccount describes a AwsApiGatewayAccount resource
type AwsApiGatewayAccount struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayAccountSpec	`json:"spec"`
}


// AwsApiGatewayAccountSpec is the spec for a AwsApiGatewayAccount Resource
type AwsApiGatewayAccountSpec struct {
	CloudwatchRoleArn	string	`json:"cloudwatch_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayAccountList is a list of AwsApiGatewayAccount resources
type AwsApiGatewayAccountList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayAccount	`json:"items"`
}


// AwsApiGatewayAccountThrottleSettings is a AwsApiGatewayAccountThrottleSettings
type AwsApiGatewayAccountThrottleSettings struct {
}
