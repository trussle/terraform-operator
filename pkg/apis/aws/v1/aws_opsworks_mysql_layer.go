
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	RootPassword	string	`json:"root_password"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	RootPasswordOnAllInstances	bool	`json:"root_password_on_all_instances"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	Name	string	`json:"name"`
	SystemPackages	string	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	EbsVolume	string	`json:"ebs_volume"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMysqlLayerList is a list of AwsOpsworksMysqlLayer resources
type AwsOpsworksMysqlLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMysqlLayer	`json:"items"`
}

