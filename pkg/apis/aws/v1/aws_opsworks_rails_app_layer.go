
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
	RubyVersion	string	`json:"ruby_version"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	Name	string	`json:"name"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	PassengerVersion	string	`json:"passenger_version"`
	RubygemsVersion	string	`json:"rubygems_version"`
	AutoHealing	bool	`json:"auto_healing"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	EbsVolume	AwsOpsworksRailsAppLayerEbsVolume	`json:"ebs_volume"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AppServer	string	`json:"app_server"`
	ManageBundler	bool	`json:"manage_bundler"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	BundlerVersion	string	`json:"bundler_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRailsAppLayerList is a list of AwsOpsworksRailsAppLayer resources
type AwsOpsworksRailsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRailsAppLayer	`json:"items"`
}


// AwsOpsworksRailsAppLayerEbsVolume is a AwsOpsworksRailsAppLayerEbsVolume
type AwsOpsworksRailsAppLayerEbsVolume struct {
	Size	int	`json:"size"`
	Type	string	`json:"type"`
	Iops	int	`json:"iops"`
	MountPoint	string	`json:"mount_point"`
	NumberOfDisks	int	`json:"number_of_disks"`
	RaidLevel	string	`json:"raid_level"`
}
