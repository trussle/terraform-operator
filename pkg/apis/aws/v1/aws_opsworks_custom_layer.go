
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksCustomLayer describes a AwsOpsworksCustomLayer resource
type AwsOpsworksCustomLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksCustomLayerSpec	`json:"spec"`
}


// AwsOpsworksCustomLayerSpec is the spec for a AwsOpsworksCustomLayer Resource
type AwsOpsworksCustomLayerSpec struct {
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	ShortName	string	`json:"short_name"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	Name	string	`json:"name"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomJson	string	`json:"custom_json"`
	SystemPackages	Generic	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	EbsVolume	Generic	`json:"ebs_volume"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksCustomLayerList is a list of AwsOpsworksCustomLayer resources
type AwsOpsworksCustomLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksCustomLayer	`json:"items"`
}

