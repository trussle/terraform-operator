
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchConfiguration describes a AwsLaunchConfiguration resource
type AwsLaunchConfiguration struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLaunchConfigurationSpec	`json:"spec"`
}


// AwsLaunchConfigurationSpec is the spec for a AwsLaunchConfiguration Resource
type AwsLaunchConfigurationSpec struct {
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	UserData	string	`json:"user_data"`
	SpotPrice	string	`json:"spot_price"`
	PlacementTenancy	string	`json:"placement_tenancy"`
	EnableMonitoring	bool	`json:"enable_monitoring"`
	InstanceType	string	`json:"instance_type"`
	UserDataBase64	string	`json:"user_data_base64"`
	SecurityGroups	string	`json:"security_groups"`
	VpcClassicLinkSecurityGroups	string	`json:"vpc_classic_link_security_groups"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	NamePrefix	string	`json:"name_prefix"`
	ImageId	string	`json:"image_id"`
	VpcClassicLinkId	string	`json:"vpc_classic_link_id"`
	EphemeralBlockDevice	AwsLaunchConfigurationEphemeralBlockDevice	`json:"ephemeral_block_device"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchConfigurationList is a list of AwsLaunchConfiguration resources
type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchConfiguration	`json:"items"`
}


// AwsLaunchConfigurationRootBlockDevice is a AwsLaunchConfigurationRootBlockDevice
type AwsLaunchConfigurationRootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// AwsLaunchConfigurationEbsBlockDevice is a AwsLaunchConfigurationEbsBlockDevice
type AwsLaunchConfigurationEbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
	NoDevice	bool	`json:"no_device"`
}

// AwsLaunchConfigurationEphemeralBlockDevice is a AwsLaunchConfigurationEphemeralBlockDevice
type AwsLaunchConfigurationEphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}
