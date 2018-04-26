
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
	ClientCertificateId	string	`json:"client_certificate_id"`
	DeploymentId	string	`json:"deployment_id"`
	Description	string	`json:"description"`
	StageName	string	`json:"stage_name"`
	Variables	map[string]string	`json:"variables"`
	Tags	map[string]string	`json:"tags"`
	CacheClusterEnabled	bool	`json:"cache_cluster_enabled"`
	CacheClusterSize	string	`json:"cache_cluster_size"`
	DocumentationVersion	string	`json:"documentation_version"`
	RestApiId	string	`json:"rest_api_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayStageList is a list of AwsApiGatewayStage resources
type AwsApiGatewayStageList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayStage	`json:"items"`
}

