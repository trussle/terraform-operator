
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLaunchTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLaunchTemplateSpec	`json:"spec"`
}

type AwsLaunchTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLaunchTemplate	`json:"items"`
}

type AwsLaunchTemplateSpec struct {
	UserData	string	`json:"user_data"`
	Description	string	`json:"description"`
	InstanceMarketOptions	[]interface{}	`json:"instance_market_options"`
	Monitoring	[]interface{}	`json:"monitoring"`
	BlockDeviceMappings	[]interface{}	`json:"block_device_mappings"`
	CreditSpecification	[]interface{}	`json:"credit_specification"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	IamInstanceProfile	[]interface{}	`json:"iam_instance_profile"`
	NetworkInterfaces	[]interface{}	`json:"network_interfaces"`
	NamePrefix	string	`json:"name_prefix"`
	RamDiskId	string	`json:"ram_disk_id"`
	SecurityGroupNames	interface{}	`json:"security_group_names"`
	TagSpecifications	[]interface{}	`json:"tag_specifications"`
	Placement	[]interface{}	`json:"placement"`
	ImageId	string	`json:"image_id"`
	KeyName	string	`json:"key_name"`
	InstanceType	string	`json:"instance_type"`
	KernelId	string	`json:"kernel_id"`
	VpcSecurityGroupIds	interface{}	`json:"vpc_security_group_ids"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	ElasticGpuSpecifications	[]interface{}	`json:"elastic_gpu_specifications"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
}
