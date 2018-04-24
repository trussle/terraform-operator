
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksRailsAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksRailsAppLayerSpec	`json:"spec"`
}

type AwsOpsworksRailsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRailsAppLayer	`json:"items"`
}

type AwsOpsworksRailsAppLayerSpec struct {
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	PassengerVersion	string	`json:"passenger_version"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	interface{}	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	BundlerVersion	string	`json:"bundler_version"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	StackId	string	`json:"stack_id"`
	Name	string	`json:"name"`
	RubyVersion	string	`json:"ruby_version"`
	RubygemsVersion	string	`json:"rubygems_version"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	AppServer	string	`json:"app_server"`
	ManageBundler	bool	`json:"manage_bundler"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
}
