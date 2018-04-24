
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStack struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksStackSpec	`json:"spec"`
}

type AwsOpsworksStackList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStack	`json:"items"`
}

type AwsOpsworksStackSpec struct {
	ManageBerkshelf	bool	`json:"manage_berkshelf"`
	BerkshelfVersion	string	`json:"berkshelf_version"`
	DefaultSshKeyName	string	`json:"default_ssh_key_name"`
	Tags	map[string]interface{}	`json:"tags"`
	UseOpsworksSecurityGroups	bool	`json:"use_opsworks_security_groups"`
	Region	string	`json:"region"`
	DefaultRootDeviceType	string	`json:"default_root_device_type"`
	HostnameTheme	string	`json:"hostname_theme"`
	Color	string	`json:"color"`
	ConfigurationManagerVersion	string	`json:"configuration_manager_version"`
	DefaultOs	string	`json:"default_os"`
	UseCustomCookbooks	bool	`json:"use_custom_cookbooks"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	DefaultInstanceProfileArn	string	`json:"default_instance_profile_arn"`
	CustomJson	string	`json:"custom_json"`
	Name	string	`json:"name"`
	ConfigurationManagerName	string	`json:"configuration_manager_name"`
}
