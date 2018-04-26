
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPhpAppLayer describes a AwsOpsworksPhpAppLayer resource
type AwsOpsworksPhpAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksPhpAppLayerSpec	`json:"spec"`
}


// AwsOpsworksPhpAppLayerSpec is the spec for a AwsOpsworksPhpAppLayer Resource
type AwsOpsworksPhpAppLayerSpec struct {
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	StackId	string	`json:"stack_id"`
	Name	string	`json:"name"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	string	`json:"system_packages"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPhpAppLayerList is a list of AwsOpsworksPhpAppLayer resources
type AwsOpsworksPhpAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksPhpAppLayer	`json:"items"`
}


// EbsVolume is a EbsVolume
type EbsVolume struct {
	Iops	int	`json:"iops"`
	MountPoint	string	`json:"mount_point"`
	NumberOfDisks	int	`json:"number_of_disks"`
	RaidLevel	string	`json:"raid_level"`
	Size	int	`json:"size"`
	Type	string	`json:"type"`
}
