
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMemcachedLayer describes a AwsOpsworksMemcachedLayer resource
type AwsOpsworksMemcachedLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksMemcachedLayerSpec	`json:"spec"`
}


// AwsOpsworksMemcachedLayerSpec is the spec for a AwsOpsworksMemcachedLayer Resource
type AwsOpsworksMemcachedLayerSpec struct {
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	StackId	string	`json:"stack_id"`
	AllocatedMemory	int	`json:"allocated_memory"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	string	`json:"system_packages"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMemcachedLayerList is a list of AwsOpsworksMemcachedLayer resources
type AwsOpsworksMemcachedLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMemcachedLayer	`json:"items"`
}


// EbsVolume is a EbsVolume
type EbsVolume struct {
	RaidLevel	string	`json:"raid_level"`
	Size	int	`json:"size"`
	Type	string	`json:"type"`
	Iops	int	`json:"iops"`
	MountPoint	string	`json:"mount_point"`
	NumberOfDisks	int	`json:"number_of_disks"`
}
