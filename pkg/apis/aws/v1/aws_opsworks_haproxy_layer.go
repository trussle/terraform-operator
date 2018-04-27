
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
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	StatsEnabled	bool	`json:"stats_enabled"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	string	`json:"system_packages"`
	StatsUrl	string	`json:"stats_url"`
	HealthcheckMethod	string	`json:"healthcheck_method"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	AwsOpsworksHaproxyLayerEbsVolume	`json:"ebs_volume"`
	StatsUser	string	`json:"stats_user"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	Name	string	`json:"name"`
	StatsPassword	string	`json:"stats_password"`
	HealthcheckUrl	string	`json:"healthcheck_url"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksHaproxyLayerList is a list of AwsOpsworksHaproxyLayer resources
type AwsOpsworksHaproxyLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksHaproxyLayer	`json:"items"`
}


// AwsOpsworksHaproxyLayerEbsVolume is a AwsOpsworksHaproxyLayerEbsVolume
type AwsOpsworksHaproxyLayerEbsVolume struct {
	Iops	int	`json:"iops"`
	MountPoint	string	`json:"mount_point"`
	NumberOfDisks	int	`json:"number_of_disks"`
	RaidLevel	string	`json:"raid_level"`
	Size	int	`json:"size"`
	Type	string	`json:"type"`
}
