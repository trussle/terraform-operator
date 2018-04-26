
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
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AllocatedMemory	int	`json:"allocated_memory"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	EbsVolume	Generic	`json:"ebs_volume"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	Generic	`json:"system_packages"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	StackId	string	`json:"stack_id"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMemcachedLayerList is a list of AwsOpsworksMemcachedLayer resources
type AwsOpsworksMemcachedLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMemcachedLayer	`json:"items"`
}

