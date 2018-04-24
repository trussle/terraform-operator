
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksGangliaLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksGangliaLayerSpec	`json:"spec"`
}

type AwsOpsworksGangliaLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksGangliaLayer	`json:"items"`
}

type AwsOpsworksGangliaLayerSpec struct {
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	interface{}	`json:"system_packages"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	Name	string	`json:"name"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	Url	string	`json:"url"`
}
