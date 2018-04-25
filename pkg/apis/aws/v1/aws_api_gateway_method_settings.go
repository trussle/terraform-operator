
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodSettings describes a AwsApiGatewayMethodSettings resource
type AwsApiGatewayMethodSettings struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayMethodSettingsSpec	`json:"spec"`
}


// AwsApiGatewayMethodSettingsSpec is the spec for a AwsApiGatewayMethodSettings Resource
type AwsApiGatewayMethodSettingsSpec struct {
	MethodPath	string	`json:"method_path"`
	Settings	[]interface{}	`json:"settings"`
	RestApiId	string	`json:"rest_api_id"`
	StageName	string	`json:"stage_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodSettingsList is a list of AwsApiGatewayMethodSettings resources
type AwsApiGatewayMethodSettingsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayMethodSettings	`json:"items"`
}

