
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
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	PassengerVersion	string	`json:"passenger_version"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	SystemPackages	Generic	`json:"system_packages"`
	ManageBundler	bool	`json:"manage_bundler"`
	BundlerVersion	string	`json:"bundler_version"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	EbsVolume	Generic	`json:"ebs_volume"`
	Name	string	`json:"name"`
	RubyVersion	string	`json:"ruby_version"`
	RubygemsVersion	string	`json:"rubygems_version"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	AppServer	string	`json:"app_server"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRailsAppLayerList is a list of AwsOpsworksRailsAppLayer resources
type AwsOpsworksRailsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRailsAppLayer	`json:"items"`
}

