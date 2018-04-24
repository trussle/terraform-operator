
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksMysqlLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksMysqlLayerSpec	`json:"spec"`
}

type AwsOpsworksMysqlLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMysqlLayer	`json:"items"`
}

type AwsOpsworksMysqlLayerSpec struct {
	SystemPackages	interface{}	`json:"system_packages"`
	RootPassword	string	`json:"root_password"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	RootPasswordOnAllInstances	bool	`json:"root_password_on_all_instances"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	StackId	string	`json:"stack_id"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	EbsVolume	interface{}	`json:"ebs_volume"`
}
