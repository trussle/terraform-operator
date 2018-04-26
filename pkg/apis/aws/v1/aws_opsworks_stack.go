
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
	ConfigurationManagerName	string	`json:"configuration_manager_name"`
	DefaultRootDeviceType	string	`json:"default_root_device_type"`
	HostnameTheme	string	`json:"hostname_theme"`
	UseOpsworksSecurityGroups	bool	`json:"use_opsworks_security_groups"`
	Name	string	`json:"name"`
	Region	string	`json:"region"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Color	string	`json:"color"`
	ConfigurationManagerVersion	string	`json:"configuration_manager_version"`
	Tags	map[string]???	`json:"tags"`
	DefaultInstanceProfileArn	string	`json:"default_instance_profile_arn"`
	ManageBerkshelf	bool	`json:"manage_berkshelf"`
	BerkshelfVersion	string	`json:"berkshelf_version"`
	CustomJson	string	`json:"custom_json"`
	DefaultOs	string	`json:"default_os"`
	DefaultSshKeyName	string	`json:"default_ssh_key_name"`
	UseCustomCookbooks	bool	`json:"use_custom_cookbooks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStackList is a list of AwsOpsworksStack resources
type AwsOpsworksStackList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStack	`json:"items"`
}


// bLnAIPRY is a bLnAIPRY
type bLnAIPRY struct {
	SshKey	string	`json:"ssh_key"`
	Type	string	`json:"type"`
	Url	string	`json:"url"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Revision	string	`json:"revision"`
}
