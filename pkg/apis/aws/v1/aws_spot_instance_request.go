
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequest describes a AwsSpotInstanceRequest resource
type AwsSpotInstanceRequest struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSpotInstanceRequestSpec	`json:"spec"`
}


// AwsSpotInstanceRequestSpec is the spec for a AwsSpotInstanceRequest Resource
type AwsSpotInstanceRequestSpec struct {
	DisableApiTermination	bool	`json:"disable_api_termination"`
	Monitoring	bool	`json:"monitoring"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	Tags	map[string]string	`json:"tags"`
	BlockDevice	map[string]string	`json:"block_device"`
	SpotType	string	`json:"spot_type"`
	GetPasswordData	bool	`json:"get_password_data"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	UserDataBase64	string	`json:"user_data_base64"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	LaunchGroup	string	`json:"launch_group"`
	VolumeTags	map[string]string	`json:"volume_tags"`
	InstanceType	string	`json:"instance_type"`
	UserData	string	`json:"user_data"`
	Ami	string	`json:"ami"`
	CreditSpecification	[]AwsSpotInstanceRequestCreditSpecification	`json:"credit_specification"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	SpotPrice	string	`json:"spot_price"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequestList is a list of AwsSpotInstanceRequest resources
type AwsSpotInstanceRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotInstanceRequest	`json:"items"`
}


// AwsSpotInstanceRequestCreditSpecification is a AwsSpotInstanceRequestCreditSpecification
type AwsSpotInstanceRequestCreditSpecification struct {
	CpuCredits	string	`json:"cpu_credits"`
}

// AwsSpotInstanceRequestEphemeralBlockDevice is a AwsSpotInstanceRequestEphemeralBlockDevice
type AwsSpotInstanceRequestEphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
	NoDevice	bool	`json:"no_device"`
}

// AwsSpotInstanceRequestNetworkInterface is a AwsSpotInstanceRequestNetworkInterface
type AwsSpotInstanceRequestNetworkInterface struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
	DeviceIndex	int	`json:"device_index"`
}

// AwsSpotInstanceRequestRootBlockDevice is a AwsSpotInstanceRequestRootBlockDevice
type AwsSpotInstanceRequestRootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// AwsSpotInstanceRequestEbsBlockDevice is a AwsSpotInstanceRequestEbsBlockDevice
type AwsSpotInstanceRequestEbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}
