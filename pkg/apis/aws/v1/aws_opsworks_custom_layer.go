
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksCustomLayer describes a AwsOpsworksCustomLayer resource
type AwsOpsworksCustomLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksCustomLayerSpec	`json:"spec"`
}


// AwsOpsworksCustomLayerSpec is the spec for a AwsOpsworksCustomLayer Resource
type AwsOpsworksCustomLayerSpec struct {
	StackId	string	`json:"stack_id"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	ShortName	string	`json:"short_name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksCustomLayerList is a list of AwsOpsworksCustomLayer resources
type AwsOpsworksCustomLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksCustomLayer	`json:"items"`
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
