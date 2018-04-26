
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStaticWebLayer describes a AwsOpsworksStaticWebLayer resource
type AwsOpsworksStaticWebLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksStaticWebLayerSpec	`json:"spec"`
}


// AwsOpsworksStaticWebLayerSpec is the spec for a AwsOpsworksStaticWebLayer Resource
type AwsOpsworksStaticWebLayerSpec struct {
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	StackId	string	`json:"stack_id"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	SystemPackages	string	`json:"system_packages"`
	Name	string	`json:"name"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStaticWebLayerList is a list of AwsOpsworksStaticWebLayer resources
type AwsOpsworksStaticWebLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStaticWebLayer	`json:"items"`
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
