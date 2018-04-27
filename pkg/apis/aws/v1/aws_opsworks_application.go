
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
	StackId	string	`json:"stack_id"`
	RailsEnv	string	`json:"rails_env"`
	AutoBundleOnDeploy	string	`json:"auto_bundle_on_deploy"`
	DataSourceDatabaseName	string	`json:"data_source_database_name"`
	Domains	[]string	`json:"domains"`
	EnableSsl	bool	`json:"enable_ssl"`
	Environment	AwsOpsworksApplicationEnvironment	`json:"environment"`
	Name	string	`json:"name"`
	Type	string	`json:"type"`
	DocumentRoot	string	`json:"document_root"`
	AwsFlowRubySettings	string	`json:"aws_flow_ruby_settings"`
	DataSourceArn	string	`json:"data_source_arn"`
	Description	string	`json:"description"`
	DataSourceType	string	`json:"data_source_type"`
	SslConfiguration	[]AwsOpsworksApplicationSslConfiguration	`json:"ssl_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksApplicationList is a list of AwsOpsworksApplication resources
type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksApplication	`json:"items"`
}


// AwsOpsworksApplicationEnvironment is a AwsOpsworksApplicationEnvironment
type AwsOpsworksApplicationEnvironment struct {
	Key	string	`json:"key"`
	Value	string	`json:"value"`
	Secure	bool	`json:"secure"`
}

// AwsOpsworksApplicationAppSource is a AwsOpsworksApplicationAppSource
type AwsOpsworksApplicationAppSource struct {
	Type	string	`json:"type"`
	Url	string	`json:"url"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Revision	string	`json:"revision"`
	SshKey	string	`json:"ssh_key"`
}

// AwsOpsworksApplicationSslConfiguration is a AwsOpsworksApplicationSslConfiguration
type AwsOpsworksApplicationSslConfiguration struct {
	Certificate	string	`json:"certificate"`
	PrivateKey	string	`json:"private_key"`
	Chain	string	`json:"chain"`
}
