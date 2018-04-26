
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMysqlLayer describes a AwsOpsworksMysqlLayer resource
type AwsOpsworksMysqlLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksMysqlLayerSpec	`json:"spec"`
}


// AwsOpsworksMysqlLayerSpec is the spec for a AwsOpsworksMysqlLayer Resource
type AwsOpsworksMysqlLayerSpec struct {
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	SystemPackages	string	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	Name	string	`json:"name"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	AutoHealing	bool	`json:"auto_healing"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	RootPasswordOnAllInstances	bool	`json:"root_password_on_all_instances"`
	RootPassword	string	`json:"root_password"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMysqlLayerList is a list of AwsOpsworksMysqlLayer resources
type AwsOpsworksMysqlLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMysqlLayer	`json:"items"`
}


// EbsVolume is a EbsVolume
type EbsVolume struct {
	RaidLevel	string	`json:"raid_level"`
	Size	int	`json:"size"`
	Type	string	`json:"type"`
	Iops	int	`json:"iops"`
	MountPoint	string	`json:"mount_point"`
	NumberOfDisks	int	`json:"number_of_disks"`
}
