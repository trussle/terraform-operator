
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
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	string	`json:"system_packages"`
	EbsVolume	string	`json:"ebs_volume"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	ShortName	string	`json:"short_name"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksCustomLayerList is a list of AwsOpsworksCustomLayer resources
type AwsOpsworksCustomLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksCustomLayer	`json:"items"`
}

