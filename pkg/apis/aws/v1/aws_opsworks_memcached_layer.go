
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMemcachedLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksMemcachedLayerSpec	`json:"spec"`
}

type AwsOpsworksMemcachedLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMemcachedLayer	`json:"items"`
}

type AwsOpsworksMemcachedLayerSpec struct {
	SystemPackages	interface{}	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AllocatedMemory	int	`json:"allocated_memory"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	Name	string	`json:"name"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
}
