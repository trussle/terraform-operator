
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncDatasource struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppsyncDatasourceSpec	`json:"spec"`
}

type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppsyncDatasource	`json:"items"`
}

type AwsAppsyncDatasourceSpec struct {
	Description	string	`json:"description"`
	LambdaConfig	[]interface{}	`json:"lambda_config"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	ElasticsearchConfig	[]interface{}	`json:"elasticsearch_config"`
	ApiId	string	`json:"api_id"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	DynamodbConfig	[]interface{}	`json:"dynamodb_config"`
}
