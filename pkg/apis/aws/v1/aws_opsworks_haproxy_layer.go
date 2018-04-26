
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
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	StackId	string	`json:"stack_id"`
	StatsEnabled	bool	`json:"stats_enabled"`
	HealthcheckUrl	string	`json:"healthcheck_url"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	EbsVolume	Generic	`json:"ebs_volume"`
	StatsUser	string	`json:"stats_user"`
	StatsPassword	string	`json:"stats_password"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	Generic	`json:"system_packages"`
	Name	string	`json:"name"`
	HealthcheckMethod	string	`json:"healthcheck_method"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	StatsUrl	string	`json:"stats_url"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksHaproxyLayerList is a list of AwsOpsworksHaproxyLayer resources
type AwsOpsworksHaproxyLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksHaproxyLayer	`json:"items"`
}

