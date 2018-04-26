
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
	Type	string	`json:"type"`
	DocumentRoot	string	`json:"document_root"`
	DataSourceDatabaseName	string	`json:"data_source_database_name"`
	Environment	string	`json:"environment"`
	Name	string	`json:"name"`
	StackId	string	`json:"stack_id"`
	RailsEnv	string	`json:"rails_env"`
	SslConfiguration	[]CSfkFohs	`json:"ssl_configuration"`
	AutoBundleOnDeploy	string	`json:"auto_bundle_on_deploy"`
	AwsFlowRubySettings	string	`json:"aws_flow_ruby_settings"`
	Description	string	`json:"description"`
	EnableSsl	bool	`json:"enable_ssl"`
	DataSourceType	string	`json:"data_source_type"`
	DataSourceArn	string	`json:"data_source_arn"`
	Domains	[]string	`json:"domains"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksApplicationList is a list of AwsOpsworksApplication resources
type AwsOpsworksApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksApplication	`json:"items"`
}


// KLdTkNyo is a KLdTkNyo
type KLdTkNyo struct {
	Type	string	`json:"type"`
	Url	string	`json:"url"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Revision	string	`json:"revision"`
	SshKey	string	`json:"ssh_key"`
}

// CSfkFohs is a CSfkFohs
type CSfkFohs struct {
	Certificate	string	`json:"certificate"`
	PrivateKey	string	`json:"private_key"`
	Chain	string	`json:"chain"`
}
