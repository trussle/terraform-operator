
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksJavaAppLayer describes a AwsOpsworksJavaAppLayer resource
type AwsOpsworksJavaAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksJavaAppLayerSpec	`json:"spec"`
}


// AwsOpsworksJavaAppLayerSpec is the spec for a AwsOpsworksJavaAppLayer Resource
type AwsOpsworksJavaAppLayerSpec struct {
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	SystemPackages	string	`json:"system_packages"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	EbsVolume	EbsVolume	`json:"ebs_volume"`
	JvmVersion	string	`json:"jvm_version"`
	AppServerVersion	string	`json:"app_server_version"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	JvmType	string	`json:"jvm_type"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	Name	string	`json:"name"`
	JvmOptions	string	`json:"jvm_options"`
	AppServer	string	`json:"app_server"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksJavaAppLayerList is a list of AwsOpsworksJavaAppLayer resources
type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksJavaAppLayer	`json:"items"`
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
