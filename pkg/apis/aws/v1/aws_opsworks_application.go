
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
	Description	string	`json:"description"`
	EnableSsl	bool	`json:"enable_ssl"`
	DocumentRoot	string	`json:"document_root"`
	DataSourceType	string	`json:"data_source_type"`
	DataSourceDatabaseName	string	`json:"data_source_database_name"`
	Domains	[]Generic	`json:"domains"`
	Environment	Generic	`json:"environment"`
	SslConfiguration	[]Generic	`json:"ssl_configuration"`
	Name	string	`json:"name"`
	RailsEnv	string	`json:"rails_env"`
	AutoBundleOnDeploy	string	`json:"auto_bundle_on_deploy"`
	AwsFlowRubySettings	string	`json:"aws_flow_ruby_settings"`
	DataSourceArn	string	`json:"data_source_arn"`
	Type	string	`json:"type"`
	StackId	string	`json:"stack_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksApplicationList is a list of AwsOpsworksApplication resources
type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksApplication	`json:"items"`
}

