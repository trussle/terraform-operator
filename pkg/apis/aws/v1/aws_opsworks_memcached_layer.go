
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
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	StackId	string	`json:"stack_id"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	Name	string	`json:"name"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	EbsVolume	string	`json:"ebs_volume"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AllocatedMemory	int	`json:"allocated_memory"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	string	`json:"system_packages"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMemcachedLayerList is a list of AwsOpsworksMemcachedLayer resources
type AwsOpsworksMemcachedLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMemcachedLayer	`json:"items"`
}

