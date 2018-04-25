
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksHaproxyLayer describes a AwsOpsworksHaproxyLayer resource
type AwsOpsworksHaproxyLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksHaproxyLayerSpec	`json:"spec"`
}


// AwsOpsworksHaproxyLayerSpec is the spec for a AwsOpsworksHaproxyLayer Resource
type AwsOpsworksHaproxyLayerSpec struct {
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	StackId	string	`json:"stack_id"`
	EbsVolume	string	`json:"ebs_volume"`
	Name	string	`json:"name"`
	StatsEnabled	bool	`json:"stats_enabled"`
	CustomJson	string	`json:"custom_json"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	HealthcheckUrl	string	`json:"healthcheck_url"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	HealthcheckMethod	string	`json:"healthcheck_method"`
	StatsUrl	string	`json:"stats_url"`
	StatsUser	string	`json:"stats_user"`
	StatsPassword	string	`json:"stats_password"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksHaproxyLayerList is a list of AwsOpsworksHaproxyLayer resources
type AwsOpsworksHaproxyLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksHaproxyLayer	`json:"items"`
}

