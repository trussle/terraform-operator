
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	StackId	string	`json:"stack_id"`
	EbsVolume	string	`json:"ebs_volume"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	PassengerVersion	string	`json:"passenger_version"`
	BundlerVersion	string	`json:"bundler_version"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	RubyVersion	string	`json:"ruby_version"`
	AppServer	string	`json:"app_server"`
	RubygemsVersion	string	`json:"rubygems_version"`
	ManageBundler	bool	`json:"manage_bundler"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	SystemPackages	string	`json:"system_packages"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRailsAppLayerList is a list of AwsOpsworksRailsAppLayer resources
type AwsOpsworksRailsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRailsAppLayer	`json:"items"`
}

