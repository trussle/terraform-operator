
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
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Description	string	`json:"description"`
	LambdaConfig	[]Generic	`json:"lambda_config"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	ApiId	string	`json:"api_id"`
	DynamodbConfig	[]Generic	`json:"dynamodb_config"`
	ElasticsearchConfig	[]Generic	`json:"elasticsearch_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncDatasourceList is a list of AwsAppsyncDatasource resources
type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppsyncDatasource	`json:"items"`
}

