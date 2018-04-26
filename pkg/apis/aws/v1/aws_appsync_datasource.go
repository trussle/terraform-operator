
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
	Name	string	`json:"name"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	ElasticsearchConfig	[]ElasticsearchConfig	`json:"elasticsearch_config"`
	LambdaConfig	[]LambdaConfig	`json:"lambda_config"`
	Type	string	`json:"type"`
	Description	string	`json:"description"`
	DynamodbConfig	[]DynamodbConfig	`json:"dynamodb_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncDatasourceList is a list of AwsAppsyncDatasource resources
type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppsyncDatasource	`json:"items"`
}


// ElasticsearchConfig is a ElasticsearchConfig
type ElasticsearchConfig struct {
	Region	string	`json:"region"`
	Endpoint	string	`json:"endpoint"`
}

// LambdaConfig is a LambdaConfig
type LambdaConfig struct {
	FunctionArn	string	`json:"function_arn"`
}

// DynamodbConfig is a DynamodbConfig
type DynamodbConfig struct {
	Region	string	`json:"region"`
	TableName	string	`json:"table_name"`
	UseCallerCredentials	bool	`json:"use_caller_credentials"`
}
