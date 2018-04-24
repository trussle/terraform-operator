
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksCustomLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksCustomLayerSpec	`json:"spec"`
}

type AwsOpsworksCustomLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksCustomLayer	`json:"items"`
}

type AwsOpsworksCustomLayerSpec struct {
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	ShortName	string	`json:"short_name"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	SystemPackages	interface{}	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
}
