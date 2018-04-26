
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayVpcLink describes a AwsApiGatewayVpcLink resource
type AwsApiGatewayVpcLink struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayVpcLinkSpec	`json:"spec"`
}


// AwsApiGatewayVpcLinkSpec is the spec for a AwsApiGatewayVpcLink Resource
type AwsApiGatewayVpcLinkSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	TargetArns	Generic	`json:"target_arns"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayVpcLinkList is a list of AwsApiGatewayVpcLink resources
type AwsApiGatewayVpcLinkList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayVpcLink	`json:"items"`
}

