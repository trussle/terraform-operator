
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
	NamePrefix	string	`json:"name_prefix"`
	BlockDeviceMappings	[]CwFkDifI	`json:"block_device_mappings"`
	KeyName	string	`json:"key_name"`
	NetworkInterfaces	[]diTskZoQ	`json:"network_interfaces"`
	Placement	[]JMqrTICT	`json:"placement"`
	SecurityGroupNames	string	`json:"security_group_names"`
	VpcSecurityGroupIds	string	`json:"vpc_security_group_ids"`
	Description	string	`json:"description"`
	ElasticGpuSpecifications	[]ojIYxyeS	`json:"elastic_gpu_specifications"`
	ImageId	string	`json:"image_id"`
	InstanceMarketOptions	[]xZyfroRO	`json:"instance_market_options"`
	InstanceType	string	`json:"instance_type"`
	Monitoring	[]PNRWCJPM	`json:"monitoring"`
	TagSpecifications	[]HDtJmHAY	`json:"tag_specifications"`
	UserData	string	`json:"user_data"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	IamInstanceProfile	[]ORsUfUMA	`json:"iam_instance_profile"`
	KernelId	string	`json:"kernel_id"`
	CreditSpecification	[]psVgzHbl	`json:"credit_specification"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	RamDiskId	string	`json:"ram_disk_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchTemplateList is a list of AwsLaunchTemplate resources
type AwsLaunchTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchTemplate	`json:"items"`
}


// BuufFMoW is a BuufFMoW
type BuufFMoW struct {
	KmsKeyId	string	`json:"kms_key_id"`
	SnapshotId	string	`json:"snapshot_id"`
	VolumeSize	int	`json:"volume_size"`
	VolumeType	string	`json:"volume_type"`
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	Encrypted	bool	`json:"encrypted"`
	Iops	int	`json:"iops"`
}

// CwFkDifI is a CwFkDifI
type CwFkDifI struct {
	DeviceName	string	`json:"device_name"`
	NoDevice	string	`json:"no_device"`
	VirtualName	string	`json:"virtual_name"`
	Ebs	[]BuufFMoW	`json:"ebs"`
}

// diTskZoQ is a diTskZoQ
type diTskZoQ struct {
	SecurityGroups	string	`json:"security_groups"`
	NetworkInterfaceId	string	`json:"network_interface_id"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	DeviceIndex	int	`json:"device_index"`
	Ipv6Addresses	string	`json:"ipv6_addresses"`
	PrivateIpAddress	string	`json:"private_ip_address"`
	Ipv4Addresses	string	`json:"ipv4_addresses"`
	SubnetId	string	`json:"subnet_id"`
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	Description	string	`json:"description"`
}

// JMqrTICT is a JMqrTICT
type JMqrTICT struct {
	AvailabilityZone	string	`json:"availability_zone"`
	GroupName	string	`json:"group_name"`
	HostId	string	`json:"host_id"`
	SpreadDomain	string	`json:"spread_domain"`
	Tenancy	string	`json:"tenancy"`
	Affinity	string	`json:"affinity"`
}

// ojIYxyeS is a ojIYxyeS
type ojIYxyeS struct {
	Type	string	`json:"type"`
}

// DMbNDRZn is a DMbNDRZn
type DMbNDRZn struct {
	SpotInstanceType	string	`json:"spot_instance_type"`
	ValidUntil	string	`json:"valid_until"`
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	InstanceInterruptionBehavior	string	`json:"instance_interruption_behavior"`
	MaxPrice	string	`json:"max_price"`
}

// xZyfroRO is a xZyfroRO
type xZyfroRO struct {
	MarketType	string	`json:"market_type"`
	SpotOptions	[]DMbNDRZn	`json:"spot_options"`
}

// PNRWCJPM is a PNRWCJPM
type PNRWCJPM struct {
	Enabled	bool	`json:"enabled"`
}

// HDtJmHAY is a HDtJmHAY
type HDtJmHAY struct {
	ResourceType	string	`json:"resource_type"`
	Tags	map[string]???	`json:"tags"`
}

// ORsUfUMA is a ORsUfUMA
type ORsUfUMA struct {
	Arn	string	`json:"arn"`
	Name	string	`json:"name"`
}

// psVgzHbl is a psVgzHbl
type psVgzHbl struct {
	CpuCredits	string	`json:"cpu_credits"`
}
