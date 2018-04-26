
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
	PlacementTenancy	string	`json:"placement_tenancy"`
	ImageId	string	`json:"image_id"`
	NamePrefix	string	`json:"name_prefix"`
	UserDataBase64	string	`json:"user_data_base64"`
	VpcClassicLinkId	string	`json:"vpc_classic_link_id"`
	SpotPrice	string	`json:"spot_price"`
	EnableMonitoring	bool	`json:"enable_monitoring"`
	SecurityGroups	string	`json:"security_groups"`
	UserData	string	`json:"user_data"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	EphemeralBlockDevice	string	`json:"ephemeral_block_device"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchConfigurationList is a list of AwsLaunchConfiguration resources
type AwsLaunchConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchConfiguration	`json:"items"`
}


// alGqARmn is a alGqARmn
type alGqARmn struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}
