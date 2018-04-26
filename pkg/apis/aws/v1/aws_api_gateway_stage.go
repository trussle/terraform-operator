
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayStage describes a AwsApiGatewayStage resource
type AwsApiGatewayStage struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayStageSpec	`json:"spec"`
}


// AwsApiGatewayStageSpec is the spec for a AwsApiGatewayStage Resource
type AwsApiGatewayStageSpec struct {
	Description	string	`json:"description"`
	DocumentationVersion	string	`json:"documentation_version"`
	StageName	string	`json:"stage_name"`
	Variables	map[string]Generic	`json:"variables"`
	CacheClusterEnabled	bool	`json:"cache_cluster_enabled"`
	CacheClusterSize	string	`json:"cache_cluster_size"`
	RestApiId	string	`json:"rest_api_id"`
	ClientCertificateId	string	`json:"client_certificate_id"`
	DeploymentId	string	`json:"deployment_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayStageList is a list of AwsApiGatewayStage resources
type AwsApiGatewayStageList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayStage	`json:"items"`
}

