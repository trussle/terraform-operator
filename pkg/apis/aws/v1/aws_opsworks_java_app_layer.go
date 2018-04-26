
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	SystemPackages	string	`json:"system_packages"`
	EbsVolume	string	`json:"ebs_volume"`
	Name	string	`json:"name"`
	JvmType	string	`json:"jvm_type"`
	JvmVersion	string	`json:"jvm_version"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomJson	string	`json:"custom_json"`
	AppServer	string	`json:"app_server"`
	AppServerVersion	string	`json:"app_server_version"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	AutoHealing	bool	`json:"auto_healing"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	JvmOptions	string	`json:"jvm_options"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksJavaAppLayerList is a list of AwsOpsworksJavaAppLayer resources
type AwsOpsworksJavaAppLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksJavaAppLayer	`json:"items"`
}

