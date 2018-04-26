
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
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	Generic	`json:"system_packages"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	NodejsVersion	string	`json:"nodejs_version"`
	CustomJson	string	`json:"custom_json"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	AutoHealing	bool	`json:"auto_healing"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	StackId	string	`json:"stack_id"`
	EbsVolume	Generic	`json:"ebs_volume"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksNodejsAppLayerList is a list of AwsOpsworksNodejsAppLayer resources
type AwsOpsworksNodejsAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksNodejsAppLayer	`json:"items"`
}

