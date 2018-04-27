
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	DeploymentId	string	`json:"deployment_id"`
	DocumentationVersion	string	`json:"documentation_version"`
	Variables	map[string]string	`json:"variables"`
	ClientCertificateId	string	`json:"client_certificate_id"`
	CacheClusterSize	string	`json:"cache_cluster_size"`
	Description	string	`json:"description"`
	RestApiId	string	`json:"rest_api_id"`
	StageName	string	`json:"stage_name"`
	Tags	map[string]string	`json:"tags"`
	CacheClusterEnabled	bool	`json:"cache_cluster_enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayStageList is a list of AwsApiGatewayStage resources
type AwsApiGatewayStageList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayStage	`json:"items"`
}

