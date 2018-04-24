
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchConfiguration struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLaunchConfigurationSpec	`json:"spec"`
}

type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchConfiguration	`json:"items"`
}

type AwsLaunchConfigurationSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	InstanceType	string	`json:"instance_type"`
	UserData	string	`json:"user_data"`
	EnableMonitoring	bool	`json:"enable_monitoring"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	VpcClassicLinkSecurityGroups	interface{}	`json:"vpc_classic_link_security_groups"`
	SpotPrice	string	`json:"spot_price"`
	EphemeralBlockDevice	interface{}	`json:"ephemeral_block_device"`
	ImageId	string	`json:"image_id"`
	SecurityGroups	interface{}	`json:"security_groups"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	VpcClassicLinkId	string	`json:"vpc_classic_link_id"`
	PlacementTenancy	string	`json:"placement_tenancy"`
}
