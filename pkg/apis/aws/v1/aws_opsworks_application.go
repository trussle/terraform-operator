
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksApplication describes a AwsOpsworksApplication resource
type AwsOpsworksApplication struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksApplicationSpec	`json:"spec"`
}


// AwsOpsworksApplicationSpec is the spec for a AwsOpsworksApplication Resource
type AwsOpsworksApplicationSpec struct {
	DataSourceArn	string	`json:"data_source_arn"`
	SslConfiguration	[]interface{}	`json:"ssl_configuration"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	Description	string	`json:"description"`
	EnableSsl	bool	`json:"enable_ssl"`
	StackId	string	`json:"stack_id"`
	RailsEnv	string	`json:"rails_env"`
	AutoBundleOnDeploy	string	`json:"auto_bundle_on_deploy"`
	DataSourceType	string	`json:"data_source_type"`
	Domains	[]interface{}	`json:"domains"`
	DocumentRoot	string	`json:"document_root"`
	AwsFlowRubySettings	string	`json:"aws_flow_ruby_settings"`
	DataSourceDatabaseName	string	`json:"data_source_database_name"`
	Environment	string	`json:"environment"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksApplicationList is a list of AwsOpsworksApplication resources
type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksApplication	`json:"items"`
}

