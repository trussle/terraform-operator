
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
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	StackId	string	`json:"stack_id"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	string	`json:"system_packages"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	EbsVolume	string	`json:"ebs_volume"`
	ShortName	string	`json:"short_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksCustomLayerList is a list of AwsOpsworksCustomLayer resources
type AwsOpsworksCustomLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksCustomLayer	`json:"items"`
}

