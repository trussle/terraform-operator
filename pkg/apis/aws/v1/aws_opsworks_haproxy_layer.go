
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
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	HealthcheckMethod	string	`json:"healthcheck_method"`
	StatsUrl	string	`json:"stats_url"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	StatsUser	string	`json:"stats_user"`
	StatsPassword	string	`json:"stats_password"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	EbsVolume	string	`json:"ebs_volume"`
	StatsEnabled	bool	`json:"stats_enabled"`
	HealthcheckUrl	string	`json:"healthcheck_url"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	StackId	string	`json:"stack_id"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksHaproxyLayerList is a list of AwsOpsworksHaproxyLayer resources
type AwsOpsworksHaproxyLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksHaproxyLayer	`json:"items"`
}

