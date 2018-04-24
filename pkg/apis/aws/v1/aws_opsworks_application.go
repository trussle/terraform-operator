
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksApplication struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksApplicationSpec	`json:"spec"`
}

type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksApplication	`json:"items"`
}

type AwsOpsworksApplicationSpec struct {
	Name	string	`json:"name"`
	StackId	string	`json:"stack_id"`
	AwsFlowRubySettings	string	`json:"aws_flow_ruby_settings"`
	DataSourceDatabaseName	string	`json:"data_source_database_name"`
	DocumentRoot	string	`json:"document_root"`
	Description	string	`json:"description"`
	Type	string	`json:"type"`
	RailsEnv	string	`json:"rails_env"`
	DataSourceArn	string	`json:"data_source_arn"`
	Domains	[]interface{}	`json:"domains"`
	EnableSsl	bool	`json:"enable_ssl"`
	SslConfiguration	[]interface{}	`json:"ssl_configuration"`
	AutoBundleOnDeploy	string	`json:"auto_bundle_on_deploy"`
	DataSourceType	string	`json:"data_source_type"`
	Environment	interface{}	`json:"environment"`
}
