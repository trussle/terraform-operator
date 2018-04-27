
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayApiKey describes a AwsApiGatewayApiKey resource
type AwsApiGatewayApiKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayApiKeySpec	`json:"spec"`
}


// AwsApiGatewayApiKeySpec is the spec for a AwsApiGatewayApiKey Resource
type AwsApiGatewayApiKeySpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Enabled	bool	`json:"enabled"`
	StageKey	AwsApiGatewayApiKeyStageKey	`json:"stage_key"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayApiKeyList is a list of AwsApiGatewayApiKey resources
type AwsApiGatewayApiKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayApiKey	`json:"items"`
}


// AwsApiGatewayApiKeyStageKey is a AwsApiGatewayApiKeyStageKey
type AwsApiGatewayApiKeyStageKey struct {
	RestApiId	string	`json:"rest_api_id"`
	StageName	string	`json:"stage_name"`
}
