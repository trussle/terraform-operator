
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
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	RubyVersion	string	`json:"ruby_version"`
	AppServer	string	`json:"app_server"`
	PassengerVersion	string	`json:"passenger_version"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	RubygemsVersion	string	`json:"rubygems_version"`
	BundlerVersion	string	`json:"bundler_version"`
	ManageBundler	bool	`json:"manage_bundler"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	string	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRailsAppLayerList is a list of AwsOpsworksRailsAppLayer resources
type AwsOpsworksRailsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRailsAppLayer	`json:"items"`
}


// EbsVolume is a EbsVolume
type EbsVolume struct {
	Type	string	`json:"type"`
	Iops	int	`json:"iops"`
	MountPoint	string	`json:"mount_point"`
	NumberOfDisks	int	`json:"number_of_disks"`
	RaidLevel	string	`json:"raid_level"`
	Size	int	`json:"size"`
}
