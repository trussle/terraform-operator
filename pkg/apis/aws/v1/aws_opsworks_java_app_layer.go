
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
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	EbsVolume	string	`json:"ebs_volume"`
	AppServerVersion	string	`json:"app_server_version"`
	JvmType	string	`json:"jvm_type"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomJson	string	`json:"custom_json"`
	SystemPackages	string	`json:"system_packages"`
	StackId	string	`json:"stack_id"`
	JvmOptions	string	`json:"jvm_options"`
	AppServer	string	`json:"app_server"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	Name	string	`json:"name"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	JvmVersion	string	`json:"jvm_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksJavaAppLayerList is a list of AwsOpsworksJavaAppLayer resources
type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksJavaAppLayer	`json:"items"`
}

