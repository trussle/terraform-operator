
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
	Password	string	`json:"password"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	Name	string	`json:"name"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	Generic	`json:"ebs_volume"`
	Url	string	`json:"url"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	StackId	string	`json:"stack_id"`
	SystemPackages	Generic	`json:"system_packages"`
	Username	string	`json:"username"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksGangliaLayerList is a list of AwsOpsworksGangliaLayer resources
type AwsOpsworksGangliaLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksGangliaLayer	`json:"items"`
}

