
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
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	StackId	string	`json:"stack_id"`
	EbsVolume	Generic	`json:"ebs_volume"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	SystemPackages	Generic	`json:"system_packages"`
	Name	string	`json:"name"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksPhpAppLayerList is a list of AwsOpsworksPhpAppLayer resources
type AwsOpsworksPhpAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksPhpAppLayer	`json:"items"`
}

