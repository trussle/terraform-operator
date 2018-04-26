
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStack describes a AwsOpsworksStack resource
type AwsOpsworksStack struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksStackSpec	`json:"spec"`
}


// AwsOpsworksStackSpec is the spec for a AwsOpsworksStack Resource
type AwsOpsworksStackSpec struct {
	Tags	map[string]string	`json:"tags"`
	UseCustomCookbooks	bool	`json:"use_custom_cookbooks"`
	UseOpsworksSecurityGroups	bool	`json:"use_opsworks_security_groups"`
	Region	string	`json:"region"`
	ManageBerkshelf	bool	`json:"manage_berkshelf"`
	BerkshelfVersion	string	`json:"berkshelf_version"`
	HostnameTheme	string	`json:"hostname_theme"`
	Name	string	`json:"name"`
	DefaultInstanceProfileArn	string	`json:"default_instance_profile_arn"`
	DefaultSshKeyName	string	`json:"default_ssh_key_name"`
	CustomJson	string	`json:"custom_json"`
	DefaultRootDeviceType	string	`json:"default_root_device_type"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Color	string	`json:"color"`
	ConfigurationManagerName	string	`json:"configuration_manager_name"`
	ConfigurationManagerVersion	string	`json:"configuration_manager_version"`
	DefaultOs	string	`json:"default_os"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStackList is a list of AwsOpsworksStack resources
type AwsOpsworksStackList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStack	`json:"items"`
}


// CustomCookbooksSource is a CustomCookbooksSource
type CustomCookbooksSource struct {
	Type	string	`json:"type"`
	Url	string	`json:"url"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Revision	string	`json:"revision"`
	SshKey	string	`json:"ssh_key"`
}
