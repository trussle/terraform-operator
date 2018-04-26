
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstance describes a AwsInstance resource
type AwsInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInstanceSpec	`json:"spec"`
}


// AwsInstanceSpec is the spec for a AwsInstance Resource
type AwsInstanceSpec struct {
	InstanceType	string	`json:"instance_type"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	Monitoring	bool	`json:"monitoring"`
	UserDataBase64	string	`json:"user_data_base64"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	Ami	string	`json:"ami"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	BlockDevice	map[string]string	`json:"block_device"`
	CreditSpecification	[]CreditSpecification	`json:"credit_specification"`
	Tags	map[string]string	`json:"tags"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	GetPasswordData	bool	`json:"get_password_data"`
	UserData	string	`json:"user_data"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstanceList is a list of AwsInstance resources
type AwsInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInstance	`json:"items"`
}


// RootBlockDevice is a RootBlockDevice
type RootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
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
