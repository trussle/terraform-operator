
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
	Type	string	`json:"type"`
	ElasticsearchConfig	[]FpLSjFbc	`json:"elasticsearch_config"`
	LambdaConfig	[]XoEFfRsW	`json:"lambda_config"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	DynamodbConfig	[]xPLDnJOb	`json:"dynamodb_config"`
	ServiceRoleArn	string	`json:"service_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncDatasourceList is a list of AwsAppsyncDatasource resources
type AwsAppsyncDatasourceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppsyncDatasource	`json:"items"`
}


// FpLSjFbc is a FpLSjFbc
type FpLSjFbc struct {
	Region	string	`json:"region"`
	Endpoint	string	`json:"endpoint"`
}

// XoEFfRsW is a XoEFfRsW
type XoEFfRsW struct {
	FunctionArn	string	`json:"function_arn"`
}

// xPLDnJOb is a xPLDnJOb
type xPLDnJOb struct {
	UseCallerCredentials	bool	`json:"use_caller_credentials"`
	Region	string	`json:"region"`
	TableName	string	`json:"table_name"`
}
