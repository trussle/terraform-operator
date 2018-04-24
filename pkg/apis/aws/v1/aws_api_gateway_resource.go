
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayResource struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayResourceSpec	`json:"spec"`
}

type AwsApiGatewayResourceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayResource	`json:"items"`
}

type AwsApiGatewayResourceSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	ParentId	string	`json:"parent_id"`
	PathPart	string	`json:"path_part"`
}
