
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
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	string	`json:"system_packages"`
	EbsVolume	string	`json:"ebs_volume"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoHealing	bool	`json:"auto_healing"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	StackId	string	`json:"stack_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStaticWebLayerList is a list of AwsOpsworksStaticWebLayer resources
type AwsOpsworksStaticWebLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStaticWebLayer	`json:"items"`
}

