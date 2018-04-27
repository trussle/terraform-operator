
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
	GetPasswordData	bool	`json:"get_password_data"`
	UserData	string	`json:"user_data"`
	InstanceType	string	`json:"instance_type"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	Ami	string	`json:"ami"`
	Tags	map[string]string	`json:"tags"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	BlockDevice	map[string]string	`json:"block_device"`
	UserDataBase64	string	`json:"user_data_base64"`
	Monitoring	bool	`json:"monitoring"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	CreditSpecification	[]AwsInstanceCreditSpecification	`json:"credit_specification"`
	SourceDestCheck	bool	`json:"source_dest_check"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstanceList is a list of AwsInstance resources
type AwsInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInstance	`json:"items"`
}


// AwsInstanceNetworkInterface is a AwsInstanceNetworkInterface
type AwsInstanceNetworkInterface struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
	DeviceIndex	int	`json:"device_index"`
}

// AwsInstanceEphemeralBlockDevice is a AwsInstanceEphemeralBlockDevice
type AwsInstanceEphemeralBlockDevice struct {
	NoDevice	bool	`json:"no_device"`
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}

// AwsInstanceEbsBlockDevice is a AwsInstanceEbsBlockDevice
type AwsInstanceEbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}

// AwsInstanceCreditSpecification is a AwsInstanceCreditSpecification
type AwsInstanceCreditSpecification struct {
	CpuCredits	string	`json:"cpu_credits"`
}

// AwsInstanceRootBlockDevice is a AwsInstanceRootBlockDevice
type AwsInstanceRootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}
