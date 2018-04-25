
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	string	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	EbsVolume	string	`json:"ebs_volume"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	AllocatedMemory	int	`json:"allocated_memory"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMemcachedLayerList is a list of AwsOpsworksMemcachedLayer resources
type AwsOpsworksMemcachedLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMemcachedLayer	`json:"items"`
}

