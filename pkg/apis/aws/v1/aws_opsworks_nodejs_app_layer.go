
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksNodejsAppLayer describes a AwsOpsworksNodejsAppLayer resource
type AwsOpsworksNodejsAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksNodejsAppLayerSpec	`json:"spec"`
}


// AwsOpsworksNodejsAppLayerSpec is the spec for a AwsOpsworksNodejsAppLayer Resource
type AwsOpsworksNodejsAppLayerSpec struct {
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	EbsVolume	string	`json:"ebs_volume"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomJson	string	`json:"custom_json"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	StackId	string	`json:"stack_id"`
	NodejsVersion	string	`json:"nodejs_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksNodejsAppLayerList is a list of AwsOpsworksNodejsAppLayer resources
type AwsOpsworksNodejsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksNodejsAppLayer	`json:"items"`
}

