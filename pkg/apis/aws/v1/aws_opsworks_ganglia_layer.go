
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
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	Url	string	`json:"url"`
	Password	string	`json:"password"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	SystemPackages	string	`json:"system_packages"`
	Name	string	`json:"name"`
	Username	string	`json:"username"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	EbsVolume	string	`json:"ebs_volume"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksGangliaLayerList is a list of AwsOpsworksGangliaLayer resources
type AwsOpsworksGangliaLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksGangliaLayer	`json:"items"`
}

