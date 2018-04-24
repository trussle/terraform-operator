
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStaticWebLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksStaticWebLayerSpec	`json:"spec"`
}

type AwsOpsworksStaticWebLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStaticWebLayer	`json:"items"`
}

type AwsOpsworksStaticWebLayerSpec struct {
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	SystemPackages	interface{}	`json:"system_packages"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
}
