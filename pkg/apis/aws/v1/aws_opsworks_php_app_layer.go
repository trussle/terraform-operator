
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
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	StackId	string	`json:"stack_id"`
	EbsVolume	string	`json:"ebs_volume"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomJson	string	`json:"custom_json"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPhpAppLayerList is a list of AwsOpsworksPhpAppLayer resources
type AwsOpsworksPhpAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksPhpAppLayer	`json:"items"`
}

