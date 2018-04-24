
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksHaproxyLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksHaproxyLayerSpec	`json:"spec"`
}

type AwsOpsworksHaproxyLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksHaproxyLayer	`json:"items"`
}

type AwsOpsworksHaproxyLayerSpec struct {
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	StatsUser	string	`json:"stats_user"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	interface{}	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	Name	string	`json:"name"`
	StatsEnabled	bool	`json:"stats_enabled"`
	HealthcheckUrl	string	`json:"healthcheck_url"`
	HealthcheckMethod	string	`json:"healthcheck_method"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StatsPassword	string	`json:"stats_password"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	StatsUrl	string	`json:"stats_url"`
}
