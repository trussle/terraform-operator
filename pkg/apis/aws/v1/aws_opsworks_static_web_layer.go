
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Name	string	`json:"name"`
	CustomJson	string	`json:"custom_json"`
	EbsVolume	Generic	`json:"ebs_volume"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	Generic	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	AutoHealing	bool	`json:"auto_healing"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStaticWebLayerList is a list of AwsOpsworksStaticWebLayer resources
type AwsOpsworksStaticWebLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStaticWebLayer	`json:"items"`
}

