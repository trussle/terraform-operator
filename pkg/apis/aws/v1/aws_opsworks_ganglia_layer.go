
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksGangliaLayer describes a AwsOpsworksGangliaLayer resource
type AwsOpsworksGangliaLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksGangliaLayerSpec	`json:"spec"`
}


// AwsOpsworksGangliaLayerSpec is the spec for a AwsOpsworksGangliaLayer Resource
type AwsOpsworksGangliaLayerSpec struct {
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	string	`json:"system_packages"`
	Password	string	`json:"password"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	Url	string	`json:"url"`
	Username	string	`json:"username"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	StackId	string	`json:"stack_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksGangliaLayerList is a list of AwsOpsworksGangliaLayer resources
type AwsOpsworksGangliaLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksGangliaLayer	`json:"items"`
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
