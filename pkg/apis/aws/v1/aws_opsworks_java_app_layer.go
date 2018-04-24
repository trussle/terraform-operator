
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksJavaAppLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksJavaAppLayerSpec	`json:"spec"`
}

type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksJavaAppLayer	`json:"items"`
}

type AwsOpsworksJavaAppLayerSpec struct {
	CustomShutdownRecipes	[]interface{}	`json:"custom_shutdown_recipes"`
	CustomJson	string	`json:"custom_json"`
	AutoHealing	bool	`json:"auto_healing"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	Name	string	`json:"name"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]interface{}	`json:"custom_setup_recipes"`
	CustomConfigureRecipes	[]interface{}	`json:"custom_configure_recipes"`
	CustomDeployRecipes	[]interface{}	`json:"custom_deploy_recipes"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AppServerVersion	string	`json:"app_server_version"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	JvmVersion	string	`json:"jvm_version"`
	JvmType	string	`json:"jvm_type"`
	CustomSecurityGroupIds	interface{}	`json:"custom_security_group_ids"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	SystemPackages	interface{}	`json:"system_packages"`
	EbsVolume	interface{}	`json:"ebs_volume"`
	AppServer	string	`json:"app_server"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	JvmOptions	string	`json:"jvm_options"`
	CustomUndeployRecipes	[]interface{}	`json:"custom_undeploy_recipes"`
}
