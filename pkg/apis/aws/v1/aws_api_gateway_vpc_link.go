
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayVpcLink struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayVpcLinkSpec	`json:"spec"`
}

type AwsApiGatewayVpcLinkList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayVpcLink	`json:"items"`
}

type AwsApiGatewayVpcLinkSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	TargetArns	interface{}	`json:"target_arns"`
}
