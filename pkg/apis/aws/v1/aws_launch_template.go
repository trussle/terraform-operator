
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	InstanceType	string	`json:"instance_type"`
	KeyName	string	`json:"key_name"`
	Monitoring	[]Generic	`json:"monitoring"`
	NetworkInterfaces	[]Generic	`json:"network_interfaces"`
	Description	string	`json:"description"`
	IamInstanceProfile	[]Generic	`json:"iam_instance_profile"`
	TagSpecifications	[]Generic	`json:"tag_specifications"`
	VpcSecurityGroupIds	Generic	`json:"vpc_security_group_ids"`
	UserData	string	`json:"user_data"`
	BlockDeviceMappings	[]Generic	`json:"block_device_mappings"`
	CreditSpecification	[]Generic	`json:"credit_specification"`
	InstanceMarketOptions	[]Generic	`json:"instance_market_options"`
	KernelId	string	`json:"kernel_id"`
	RamDiskId	string	`json:"ram_disk_id"`
	SecurityGroupNames	Generic	`json:"security_group_names"`
	NamePrefix	string	`json:"name_prefix"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	ElasticGpuSpecifications	[]Generic	`json:"elastic_gpu_specifications"`
	ImageId	string	`json:"image_id"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	Placement	[]Generic	`json:"placement"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchTemplateList is a list of AwsLaunchTemplate resources
type AwsLaunchTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchTemplate	`json:"items"`
}

