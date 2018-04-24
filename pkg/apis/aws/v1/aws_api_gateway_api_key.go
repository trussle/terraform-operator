
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayApiKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayApiKeySpec	`json:"spec"`
}

type AwsApiGatewayApiKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayApiKey	`json:"items"`
}

type AwsApiGatewayApiKeySpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Enabled	bool	`json:"enabled"`
	StageKey	interface{}	`json:"stage_key"`
}
