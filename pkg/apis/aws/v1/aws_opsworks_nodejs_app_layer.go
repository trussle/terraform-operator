
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksNodejsAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksNodejsAppLayerSpec	`json:"spec"`
}

type AwsOpsworksNodejsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksNodejsAppLayer	`json:"items"`
}

type AwsOpsworksNodejsAppLayerSpec struct {
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	interface{}	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	NodejsVersion	string	`json:"nodejs_version"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	EbsVolume	interface{}	`json:"ebs_volume"`
}
