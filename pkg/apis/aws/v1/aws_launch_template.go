
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchTemplate describes a AwsLaunchTemplate resource
type AwsLaunchTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLaunchTemplateSpec	`json:"spec"`
}


// AwsLaunchTemplateSpec is the spec for a AwsLaunchTemplate Resource
type AwsLaunchTemplateSpec struct {
	ImageId	string	`json:"image_id"`
	InstanceMarketOptions	[]InstanceMarketOptions	`json:"instance_market_options"`
	SecurityGroupNames	string	`json:"security_group_names"`
	VpcSecurityGroupIds	string	`json:"vpc_security_group_ids"`
	TagSpecifications	[]TagSpecifications	`json:"tag_specifications"`
	NamePrefix	string	`json:"name_prefix"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	Monitoring	[]Monitoring	`json:"monitoring"`
	NetworkInterfaces	[]NetworkInterfaces	`json:"network_interfaces"`
	Placement	[]Placement	`json:"placement"`
	UserData	string	`json:"user_data"`
	IamInstanceProfile	[]IamInstanceProfile	`json:"iam_instance_profile"`
	CreditSpecification	[]CreditSpecification	`json:"credit_specification"`
	RamDiskId	string	`json:"ram_disk_id"`
	BlockDeviceMappings	[]BlockDeviceMappings	`json:"block_device_mappings"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	ElasticGpuSpecifications	[]ElasticGpuSpecifications	`json:"elastic_gpu_specifications"`
	InstanceType	string	`json:"instance_type"`
	KernelId	string	`json:"kernel_id"`
	KeyName	string	`json:"key_name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchTemplateList is a list of AwsLaunchTemplate resources
type AwsLaunchTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchTemplate	`json:"items"`
}


// SpotOptions is a SpotOptions
type SpotOptions struct {
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	InstanceInterruptionBehavior	string	`json:"instance_interruption_behavior"`
	MaxPrice	string	`json:"max_price"`
	SpotInstanceType	string	`json:"spot_instance_type"`
	ValidUntil	string	`json:"valid_until"`
}

// Ebs is a Ebs
type Ebs struct {
	VolumeSize	int	`json:"volume_size"`
	VolumeType	string	`json:"volume_type"`
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	Encrypted	bool	`json:"encrypted"`
	Iops	int	`json:"iops"`
	KmsKeyId	string	`json:"kms_key_id"`
	SnapshotId	string	`json:"snapshot_id"`
}

// BlockDeviceMappings is a BlockDeviceMappings
type BlockDeviceMappings struct {
	DeviceName	string	`json:"device_name"`
	NoDevice	string	`json:"no_device"`
	VirtualName	string	`json:"virtual_name"`
	Ebs	[]Ebs	`json:"ebs"`
}

// NetworkInterfaces is a NetworkInterfaces
type NetworkInterfaces struct {
	PrivateIpAddress	string	`json:"private_ip_address"`
	Ipv4Addresses	string	`json:"ipv4_addresses"`
	Description	string	`json:"description"`
	SecurityGroups	string	`json:"security_groups"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
	SubnetId	string	`json:"subnet_id"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceIndex	int	`json:"device_index"`
	Ipv6Addresses	string	`json:"ipv6_addresses"`
}

// Placement is a Placement
type Placement struct {
	Affinity	string	`json:"affinity"`
	AvailabilityZone	string	`json:"availability_zone"`
	GroupName	string	`json:"group_name"`
	HostId	string	`json:"host_id"`
	SpreadDomain	string	`json:"spread_domain"`
	Tenancy	string	`json:"tenancy"`
}

// IamInstanceProfile is a IamInstanceProfile
type IamInstanceProfile struct {
	Arn	string	`json:"arn"`
	Name	string	`json:"name"`
}

// CreditSpecification is a CreditSpecification
type CreditSpecification struct {
	CpuCredits	string	`json:"cpu_credits"`
}

// ElasticGpuSpecifications is a ElasticGpuSpecifications
type ElasticGpuSpecifications struct {
	Type	string	`json:"type"`
}

// InstanceMarketOptions is a InstanceMarketOptions
type InstanceMarketOptions struct {
	MarketType	string	`json:"market_type"`
	SpotOptions	[]SpotOptions	`json:"spot_options"`
}

// TagSpecifications is a TagSpecifications
type TagSpecifications struct {
	ResourceType	string	`json:"resource_type"`
	Tags	map[string]string	`json:"tags"`
}

// Monitoring is a Monitoring
type Monitoring struct {
	Enabled	bool	`json:"enabled"`
}
