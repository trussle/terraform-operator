
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
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	StackId	string	`json:"stack_id"`
	RootPasswordOnAllInstances	bool	`json:"root_password_on_all_instances"`
	EbsVolume	Generic	`json:"ebs_volume"`
	Name	string	`json:"name"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	Generic	`json:"system_packages"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	AutoHealing	bool	`json:"auto_healing"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	RootPassword	string	`json:"root_password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMysqlLayerList is a list of AwsOpsworksMysqlLayer resources
type AwsOpsworksMysqlLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMysqlLayer	`json:"items"`
}

