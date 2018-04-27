
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncDatasource describes a AwsAppsyncDatasource resource
type AwsAppsyncDatasource struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppsyncDatasourceSpec	`json:"spec"`
}


// AwsAppsyncDatasourceSpec is the spec for a AwsAppsyncDatasource Resource
type AwsAppsyncDatasourceSpec struct {
	ApiId	string	`json:"api_id"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	DynamodbConfig	[]AwsAppsyncDatasourceDynamodbConfig	`json:"dynamodb_config"`
	ElasticsearchConfig	[]AwsAppsyncDatasourceElasticsearchConfig	`json:"elasticsearch_config"`
	LambdaConfig	[]AwsAppsyncDatasourceLambdaConfig	`json:"lambda_config"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncDatasourceList is a list of AwsAppsyncDatasource resources
type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppsyncDatasource	`json:"items"`
}


// AwsAppsyncDatasourceDynamodbConfig is a AwsAppsyncDatasourceDynamodbConfig
type AwsAppsyncDatasourceDynamodbConfig struct {
	Region	string	`json:"region"`
	TableName	string	`json:"table_name"`
	UseCallerCredentials	bool	`json:"use_caller_credentials"`
}

// AwsAppsyncDatasourceElasticsearchConfig is a AwsAppsyncDatasourceElasticsearchConfig
type AwsAppsyncDatasourceElasticsearchConfig struct {
	Region	string	`json:"region"`
	Endpoint	string	`json:"endpoint"`
}

// AwsAppsyncDatasourceLambdaConfig is a AwsAppsyncDatasourceLambdaConfig
type AwsAppsyncDatasourceLambdaConfig struct {
	FunctionArn	string	`json:"function_arn"`
}
