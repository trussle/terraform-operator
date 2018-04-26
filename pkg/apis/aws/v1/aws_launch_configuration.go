
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
	InstanceType	string	`json:"instance_type"`
	VpcClassicLinkSecurityGroups	string	`json:"vpc_classic_link_security_groups"`
	ImageId	string	`json:"image_id"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	SecurityGroups	string	`json:"security_groups"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	NamePrefix	string	`json:"name_prefix"`
	VpcClassicLinkId	string	`json:"vpc_classic_link_id"`
	SpotPrice	string	`json:"spot_price"`
	PlacementTenancy	string	`json:"placement_tenancy"`
	EnableMonitoring	bool	`json:"enable_monitoring"`
	UserData	string	`json:"user_data"`
	UserDataBase64	string	`json:"user_data_base64"`
	EphemeralBlockDevice	EphemeralBlockDevice	`json:"ephemeral_block_device"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchConfigurationList is a list of AwsLaunchConfiguration resources
type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchConfiguration	`json:"items"`
}


// EbsBlockDevice is a EbsBlockDevice
type EbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
	NoDevice	bool	`json:"no_device"`
}

// EphemeralBlockDevice is a EphemeralBlockDevice
type EphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}

// RootBlockDevice is a RootBlockDevice
type RootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}
