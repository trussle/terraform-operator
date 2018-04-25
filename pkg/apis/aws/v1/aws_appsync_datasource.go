
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	ServiceRoleArn	string	`json:"service_role_arn"`
	Type	string	`json:"type"`
	DynamodbConfig	[]interface{}	`json:"dynamodb_config"`
	LambdaConfig	[]interface{}	`json:"lambda_config"`
	ElasticsearchConfig	[]interface{}	`json:"elasticsearch_config"`
	ApiId	string	`json:"api_id"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncDatasourceList is a list of AwsAppsyncDatasource resources
type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppsyncDatasource	`json:"items"`
}

