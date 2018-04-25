
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksGangliaLayer describes a AwsOpsworksGangliaLayer resource
type AwsOpsworksGangliaLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksGangliaLayerSpec	`json:"spec"`
}


// AwsOpsworksGangliaLayerSpec is the spec for a AwsOpsworksGangliaLayer Resource
type AwsOpsworksGangliaLayerSpec struct {
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Password	string	`json:"password"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	Url	string	`json:"url"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	EbsVolume	string	`json:"ebs_volume"`
	Name	string	`json:"name"`
	Username	string	`json:"username"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	string	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksGangliaLayerList is a list of AwsOpsworksGangliaLayer resources
type AwsOpsworksGangliaLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksGangliaLayer	`json:"items"`
}

