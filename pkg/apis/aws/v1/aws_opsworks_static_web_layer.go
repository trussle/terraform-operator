
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
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	string	`json:"ebs_volume"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	SystemPackages	string	`json:"system_packages"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	CustomJson	string	`json:"custom_json"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	StackId	string	`json:"stack_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStaticWebLayerList is a list of AwsOpsworksStaticWebLayer resources
type AwsOpsworksStaticWebLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStaticWebLayer	`json:"items"`
}

