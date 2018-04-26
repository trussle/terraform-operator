
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsService describes a AwsEcsService resource
type AwsEcsService struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcsServiceSpec	`json:"spec"`
}


// AwsEcsServiceSpec is the spec for a AwsEcsService Resource
type AwsEcsServiceSpec struct {
	TaskDefinition	string	`json:"task_definition"`
	DesiredCount	int	`json:"desired_count"`
	LoadBalancer	string	`json:"load_balancer"`
	NetworkConfiguration	[]eKovPdQS	`json:"network_configuration"`
	ServiceRegistries	string	`json:"service_registries"`
	Name	string	`json:"name"`
	LaunchType	string	`json:"launch_type"`
	DeploymentMaximumPercent	int	`json:"deployment_maximum_percent"`
	PlacementConstraints	string	`json:"placement_constraints"`
	HealthCheckGracePeriodSeconds	int	`json:"health_check_grace_period_seconds"`
	DeploymentMinimumHealthyPercent	int	`json:"deployment_minimum_healthy_percent"`
	PlacementStrategy	string	`json:"placement_strategy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsServiceList is a list of AwsEcsService resources
type AwsEcsServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsService	`json:"items"`
}


// eKovPdQS is a eKovPdQS
type eKovPdQS struct {
	Subnets	string	`json:"subnets"`
	AssignPublicIp	bool	`json:"assign_public_ip"`
	SecurityGroups	string	`json:"security_groups"`
}
