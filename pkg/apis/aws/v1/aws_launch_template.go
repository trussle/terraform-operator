
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
	BlockDeviceMappings	[]interface{}	`json:"block_device_mappings"`
	KeyName	string	`json:"key_name"`
	Description	string	`json:"description"`
	IamInstanceProfile	[]interface{}	`json:"iam_instance_profile"`
	InstanceType	string	`json:"instance_type"`
	NetworkInterfaces	[]interface{}	`json:"network_interfaces"`
	TagSpecifications	[]interface{}	`json:"tag_specifications"`
	UserData	string	`json:"user_data"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	InstanceMarketOptions	[]interface{}	`json:"instance_market_options"`
	Monitoring	[]interface{}	`json:"monitoring"`
	RamDiskId	string	`json:"ram_disk_id"`
	SecurityGroupNames	string	`json:"security_group_names"`
	VpcSecurityGroupIds	string	`json:"vpc_security_group_ids"`
	NamePrefix	string	`json:"name_prefix"`
	CreditSpecification	[]interface{}	`json:"credit_specification"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	ElasticGpuSpecifications	[]interface{}	`json:"elastic_gpu_specifications"`
	ImageId	string	`json:"image_id"`
	KernelId	string	`json:"kernel_id"`
	Placement	[]interface{}	`json:"placement"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLaunchTemplateList is a list of AwsLaunchTemplate resources
type AwsLaunchTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchTemplate	`json:"items"`
}

