
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	StackId	string	`json:"stack_id"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	string	`json:"ebs_volume"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	AutoHealing	bool	`json:"auto_healing"`
	SystemPackages	string	`json:"system_packages"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPhpAppLayerList is a list of AwsOpsworksPhpAppLayer resources
type AwsOpsworksPhpAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksPhpAppLayer	`json:"items"`
}

