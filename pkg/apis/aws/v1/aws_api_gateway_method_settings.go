
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayMethodSettings struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayMethodSettingsSpec	`json:"spec"`
}

type AwsApiGatewayMethodSettingsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayMethodSettings	`json:"items"`
}

type AwsApiGatewayMethodSettingsSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	StageName	string	`json:"stage_name"`
	MethodPath	string	`json:"method_path"`
	Settings	[]interface{}	`json:"settings"`
}
