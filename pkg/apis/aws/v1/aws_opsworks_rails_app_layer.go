
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRailsAppLayer describes a AwsOpsworksRailsAppLayer resource
type AwsOpsworksRailsAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksRailsAppLayerSpec	`json:"spec"`
}


// AwsOpsworksRailsAppLayerSpec is the spec for a AwsOpsworksRailsAppLayer Resource
type AwsOpsworksRailsAppLayerSpec struct {
	Name	string	`json:"name"`
	RubyVersion	string	`json:"ruby_version"`
	PassengerVersion	string	`json:"passenger_version"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	StackId	string	`json:"stack_id"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	RubygemsVersion	string	`json:"rubygems_version"`
	BundlerVersion	string	`json:"bundler_version"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	string	`json:"system_packages"`
	EbsVolume	string	`json:"ebs_volume"`
	ManageBundler	bool	`json:"manage_bundler"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	AppServer	string	`json:"app_server"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRailsAppLayerList is a list of AwsOpsworksRailsAppLayer resources
type AwsOpsworksRailsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRailsAppLayer	`json:"items"`
}

