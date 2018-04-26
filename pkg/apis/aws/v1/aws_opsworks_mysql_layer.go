
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMysqlLayer describes a AwsOpsworksMysqlLayer resource
type AwsOpsworksMysqlLayer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksMysqlLayerSpec	`json:"spec"`
}


// AwsOpsworksMysqlLayerSpec is the spec for a AwsOpsworksMysqlLayer Resource
type AwsOpsworksMysqlLayerSpec struct {
	DrainElbOnShutdown	bool	`json:"drain_elb_on_shutdown"`
	UseEbsOptimizedInstances	bool	`json:"use_ebs_optimized_instances"`
	Name	string	`json:"name"`
	CustomInstanceProfileArn	string	`json:"custom_instance_profile_arn"`
	CustomConfigureRecipes	[]string	`json:"custom_configure_recipes"`
	CustomShutdownRecipes	[]string	`json:"custom_shutdown_recipes"`
	CustomSecurityGroupIds	string	`json:"custom_security_group_ids"`
	CustomJson	string	`json:"custom_json"`
	CustomDeployRecipes	[]string	`json:"custom_deploy_recipes"`
	InstallUpdatesOnBoot	bool	`json:"install_updates_on_boot"`
	StackId	string	`json:"stack_id"`
	RootPasswordOnAllInstances	bool	`json:"root_password_on_all_instances"`
	AutoAssignElasticIps	bool	`json:"auto_assign_elastic_ips"`
	ElasticLoadBalancer	string	`json:"elastic_load_balancer"`
	CustomSetupRecipes	[]string	`json:"custom_setup_recipes"`
	CustomUndeployRecipes	[]string	`json:"custom_undeploy_recipes"`
	SystemPackages	string	`json:"system_packages"`
	AutoAssignPublicIps	bool	`json:"auto_assign_public_ips"`
	AutoHealing	bool	`json:"auto_healing"`
	InstanceShutdownTimeout	int	`json:"instance_shutdown_timeout"`
	EbsVolume	string	`json:"ebs_volume"`
	RootPassword	string	`json:"root_password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksMysqlLayerList is a list of AwsOpsworksMysqlLayer resources
type AwsOpsworksMysqlLayerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksMysqlLayer	`json:"items"`
}

