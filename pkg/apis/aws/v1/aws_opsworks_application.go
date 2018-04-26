
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	SslConfiguration	[]SslConfiguration	`json:"ssl_configuration"`
	RailsEnv	string	`json:"rails_env"`
	AutoBundleOnDeploy	string	`json:"auto_bundle_on_deploy"`
	DataSourceType	string	`json:"data_source_type"`
	Type	string	`json:"type"`
	Domains	[]string	`json:"domains"`
	Description	string	`json:"description"`
	Environment	Environment	`json:"environment"`
	EnableSsl	bool	`json:"enable_ssl"`
	Name	string	`json:"name"`
	StackId	string	`json:"stack_id"`
	DocumentRoot	string	`json:"document_root"`
	AwsFlowRubySettings	string	`json:"aws_flow_ruby_settings"`
	DataSourceDatabaseName	string	`json:"data_source_database_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksApplicationList is a list of AwsOpsworksApplication resources
type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksApplication	`json:"items"`
}


// SslConfiguration is a SslConfiguration
type SslConfiguration struct {
	Certificate	string	`json:"certificate"`
	PrivateKey	string	`json:"private_key"`
	Chain	string	`json:"chain"`
}

// AppSource is a AppSource
type AppSource struct {
	Type	string	`json:"type"`
	Url	string	`json:"url"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Revision	string	`json:"revision"`
	SshKey	string	`json:"ssh_key"`
}

// Environment is a Environment
type Environment struct {
	Key	string	`json:"key"`
	Value	string	`json:"value"`
	Secure	bool	`json:"secure"`
}
