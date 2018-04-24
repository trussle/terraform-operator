
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayStage struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayStageSpec	`json:"spec"`
}

type AwsApiGatewayStageList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayStage	`json:"items"`
}

type AwsApiGatewayStageSpec struct {
	CacheClusterEnabled	bool	`json:"cache_cluster_enabled"`
	CacheClusterSize	string	`json:"cache_cluster_size"`
	DocumentationVersion	string	`json:"documentation_version"`
	StageName	string	`json:"stage_name"`
	Variables	map[string]interface{}	`json:"variables"`
	ClientCertificateId	string	`json:"client_certificate_id"`
	DeploymentId	string	`json:"deployment_id"`
	Description	string	`json:"description"`
	RestApiId	string	`json:"rest_api_id"`
}
