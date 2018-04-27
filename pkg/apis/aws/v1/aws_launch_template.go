
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
	CreditSpecification	[]AwsLaunchTemplateCreditSpecification	`json:"credit_specification"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	ImageId	string	`json:"image_id"`
	InstanceMarketOptions	[]AwsLaunchTemplateInstanceMarketOptions	`json:"instance_market_options"`
	KernelId	string	`json:"kernel_id"`
	KeyName	string	`json:"key_name"`
	InstanceType	string	`json:"instance_type"`
	Placement	[]AwsLaunchTemplatePlacement	`json:"placement"`
	RamDiskId	string	`json:"ram_disk_id"`
	TagSpecifications	[]AwsLaunchTemplateTagSpecifications	`json:"tag_specifications"`
	UserData	string	`json:"user_data"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	SecurityGroupNames	string	`json:"security_group_names"`
	NamePrefix	string	`json:"name_prefix"`
	BlockDeviceMappings	[]AwsLaunchTemplateBlockDeviceMappings	`json:"block_device_mappings"`
	ElasticGpuSpecifications	[]AwsLaunchTemplateElasticGpuSpecifications	`json:"elastic_gpu_specifications"`
	IamInstanceProfile	[]AwsLaunchTemplateIamInstanceProfile	`json:"iam_instance_profile"`
	VpcSecurityGroupIds	string	`json:"vpc_security_group_ids"`
	Description	string	`json:"description"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	Monitoring	[]AwsLaunchTemplateMonitoring	`json:"monitoring"`
	NetworkInterfaces	[]AwsLaunchTemplateNetworkInterfaces	`json:"network_interfaces"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchTemplateList is a list of AwsLaunchTemplate resources
type AwsLaunchTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchTemplate	`json:"items"`
}


// AwsLaunchTemplateTagSpecifications is a AwsLaunchTemplateTagSpecifications
type AwsLaunchTemplateTagSpecifications struct {
	ResourceType	string	`json:"resource_type"`
	Tags	map[string]string	`json:"tags"`
}

// AwsLaunchTemplateNetworkInterfaces is a AwsLaunchTemplateNetworkInterfaces
type AwsLaunchTemplateNetworkInterfaces struct {
	SubnetId	string	`json:"subnet_id"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	Description	string	`json:"description"`
	SecurityGroups	string	`json:"security_groups"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
	PrivateIpAddress	string	`json:"private_ip_address"`
	Ipv4Addresses	string	`json:"ipv4_addresses"`
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceIndex	int	`json:"device_index"`
	Ipv6Addresses	string	`json:"ipv6_addresses"`
}

// AwsLaunchTemplateCreditSpecification is a AwsLaunchTemplateCreditSpecification
type AwsLaunchTemplateCreditSpecification struct {
	CpuCredits	string	`json:"cpu_credits"`
}

// AwsLaunchTemplateSpotOptions is a AwsLaunchTemplateSpotOptions
type AwsLaunchTemplateSpotOptions struct {
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	InstanceInterruptionBehavior	string	`json:"instance_interruption_behavior"`
	MaxPrice	string	`json:"max_price"`
	SpotInstanceType	string	`json:"spot_instance_type"`
	ValidUntil	string	`json:"valid_until"`
}

// AwsLaunchTemplatePlacement is a AwsLaunchTemplatePlacement
type AwsLaunchTemplatePlacement struct {
	Tenancy	string	`json:"tenancy"`
	Affinity	string	`json:"affinity"`
	AvailabilityZone	string	`json:"availability_zone"`
	GroupName	string	`json:"group_name"`
	HostId	string	`json:"host_id"`
	SpreadDomain	string	`json:"spread_domain"`
}

// AwsLaunchTemplateElasticGpuSpecifications is a AwsLaunchTemplateElasticGpuSpecifications
type AwsLaunchTemplateElasticGpuSpecifications struct {
	Type	string	`json:"type"`
}

// AwsLaunchTemplateIamInstanceProfile is a AwsLaunchTemplateIamInstanceProfile
type AwsLaunchTemplateIamInstanceProfile struct {
	Arn	string	`json:"arn"`
	Name	string	`json:"name"`
}

// AwsLaunchTemplateMonitoring is a AwsLaunchTemplateMonitoring
type AwsLaunchTemplateMonitoring struct {
	Enabled	bool	`json:"enabled"`
}

// AwsLaunchTemplateInstanceMarketOptions is a AwsLaunchTemplateInstanceMarketOptions
type AwsLaunchTemplateInstanceMarketOptions struct {
	MarketType	string	`json:"market_type"`
	SpotOptions	[]AwsLaunchTemplateSpotOptions	`json:"spot_options"`
}

// AwsLaunchTemplateEbs is a AwsLaunchTemplateEbs
type AwsLaunchTemplateEbs struct {
	SnapshotId	string	`json:"snapshot_id"`
	VolumeSize	int	`json:"volume_size"`
	VolumeType	string	`json:"volume_type"`
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	Encrypted	bool	`json:"encrypted"`
	Iops	int	`json:"iops"`
	KmsKeyId	string	`json:"kms_key_id"`
}

// AwsLaunchTemplateBlockDeviceMappings is a AwsLaunchTemplateBlockDeviceMappings
type AwsLaunchTemplateBlockDeviceMappings struct {
	Ebs	[]AwsLaunchTemplateEbs	`json:"ebs"`
	DeviceName	string	`json:"device_name"`
	NoDevice	string	`json:"no_device"`
	VirtualName	string	`json:"virtual_name"`
}
