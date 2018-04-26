
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksHaproxyLayer describes a AwsOpsworksHaproxyLayer resource
type AwsOpsworksHaproxyLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksHaproxyLayerSpec	`json:"spec"`
}


// AwsOpsworksHaproxyLayerSpec is the spec for a AwsOpsworksHaproxyLayer Resource
type AwsOpsworksHaproxyLayerSpec struct {
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	Name	string	`json:"name"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	StatsUser	string	`json:"stats_user"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	HealthcheckMethod	string	`json:"healthcheck_method"`
	StatsEnabled	bool	`json:"stats_enabled"`
	StatsUrl	string	`json:"stats_url"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	SystemPackages	string	`json:"system_packages"`
	StatsPassword	string	`json:"stats_password"`
	HealthcheckUrl	string	`json:"healthcheck_url"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksHaproxyLayerList is a list of AwsOpsworksHaproxyLayer resources
type AwsOpsworksHaproxyLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksHaproxyLayer	`json:"items"`
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
