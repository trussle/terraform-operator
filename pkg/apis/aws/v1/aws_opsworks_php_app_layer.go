
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksPhpAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksPhpAppLayerSpec	`json:"spec"`
}

type AwsOpsworksPhpAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksPhpAppLayer	`json:"items"`
}

type AwsOpsworksPhpAppLayerSpec struct {
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomJson	string	`json:"custom_json"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	SystemPackages	interface{}	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	StackId	string	`json:"stack_id"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
}
