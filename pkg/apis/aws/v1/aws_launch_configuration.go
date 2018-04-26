
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	PlacementTenancy	string	`json:"placement_tenancy"`
	EnableMonitoring	bool	`json:"enable_monitoring"`
	EphemeralBlockDevice	Generic	`json:"ephemeral_block_device"`
	NamePrefix	string	`json:"name_prefix"`
	VpcClassicLinkId	string	`json:"vpc_classic_link_id"`
	SpotPrice	string	`json:"spot_price"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	ImageId	string	`json:"image_id"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	UserData	string	`json:"user_data"`
	SecurityGroups	Generic	`json:"security_groups"`
	VpcClassicLinkSecurityGroups	Generic	`json:"vpc_classic_link_security_groups"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchConfigurationList is a list of AwsLaunchConfiguration resources
type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchConfiguration	`json:"items"`
}

