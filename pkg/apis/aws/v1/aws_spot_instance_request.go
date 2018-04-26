
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
	Tags	map[string]string	`json:"tags"`
	BlockDevice	map[string]string	`json:"block_device"`
	InstanceType	string	`json:"instance_type"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	SpotType	string	`json:"spot_type"`
	LaunchGroup	string	`json:"launch_group"`
	CreditSpecification	[]CreditSpecification	`json:"credit_specification"`
	SpotPrice	string	`json:"spot_price"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	Monitoring	bool	`json:"monitoring"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	Ami	string	`json:"ami"`
	UserDataBase64	string	`json:"user_data_base64"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	UserData	string	`json:"user_data"`
	VolumeTags	map[string]string	`json:"volume_tags"`
	GetPasswordData	bool	`json:"get_password_data"`
	EbsOptimized	bool	`json:"ebs_optimized"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequestList is a list of AwsSpotInstanceRequest resources
type AwsSpotInstanceRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotInstanceRequest	`json:"items"`
}


// EbsBlockDevice is a EbsBlockDevice
type EbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}

// EphemeralBlockDevice is a EphemeralBlockDevice
type EphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
	NoDevice	bool	`json:"no_device"`
}

// CreditSpecification is a CreditSpecification
type CreditSpecification struct {
	CpuCredits	string	`json:"cpu_credits"`
}

// NetworkInterface is a NetworkInterface
type NetworkInterface struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
	DeviceIndex	int	`json:"device_index"`
}

// RootBlockDevice is a RootBlockDevice
type RootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}
