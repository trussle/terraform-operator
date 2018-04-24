
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayBasePathMapping struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayBasePathMappingSpec	`json:"spec"`
}

type AwsApiGatewayBasePathMappingList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayBasePathMapping	`json:"items"`
}

type AwsApiGatewayBasePathMappingSpec struct {
	ApiId	string	`json:"api_id"`
	BasePath	string	`json:"base_path"`
	StageName	string	`json:"stage_name"`
	DomainName	string	`json:"domain_name"`
}
