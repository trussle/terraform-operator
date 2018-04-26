
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksJavaAppLayer describes a AwsOpsworksJavaAppLayer resource
type AwsOpsworksJavaAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksJavaAppLayerSpec	`json:"spec"`
}


// AwsOpsworksJavaAppLayerSpec is the spec for a AwsOpsworksJavaAppLayer Resource
type AwsOpsworksJavaAppLayerSpec struct {
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomConfigureRecipes	[]Generic	`json:"custom_configure_recipes"`
	CustomUndeployRecipes	[]Generic	`json:"custom_undeploy_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	Generic	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	JvmVersion	string	`json:"jvm_version"`
	JvmOptions	string	`json:"jvm_options"`
	AppServer	string	`json:"app_server"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomSecurityGroupIds	Generic	`json:"custom_security_group_ids"`
	AutoHealing	bool	`json:"auto_healing"`
	EbsVolume	Generic	`json:"ebs_volume"`
	CustomSetupRecipes	[]Generic	`json:"custom_setup_recipes"`
	CustomJson	string	`json:"custom_json"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	JvmType	string	`json:"jvm_type"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomDeployRecipes	[]Generic	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]Generic	`json:"custom_shutdown_recipes"`
	Name	string	`json:"name"`
	AppServerVersion	string	`json:"app_server_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksJavaAppLayerList is a list of AwsOpsworksJavaAppLayer resources
type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksJavaAppLayer	`json:"items"`
}

