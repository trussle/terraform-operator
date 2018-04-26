
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksNodejsAppLayer describes a AwsOpsworksNodejsAppLayer resource
type AwsOpsworksNodejsAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksNodejsAppLayerSpec	`json:"spec"`
}


// AwsOpsworksNodejsAppLayerSpec is the spec for a AwsOpsworksNodejsAppLayer Resource
type AwsOpsworksNodejsAppLayerSpec struct {
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	NodejsVersion	string	`json:"nodejs_version"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	Name	string	`json:"name"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	StackId	string	`json:"stack_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksNodejsAppLayerList is a list of AwsOpsworksNodejsAppLayer resources
type AwsOpsworksNodejsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksNodejsAppLayer	`json:"items"`
}


// EbsVolume is a EbsVolume
type EbsVolume struct {
	Size	int	`json:"size"`
	Type	string	`json:"type"`
	Iops	int	`json:"iops"`
	MountPoint	string	`json:"mount_point"`
	NumberOfDisks	int	`json:"number_of_disks"`
	RaidLevel	string	`json:"raid_level"`
}
