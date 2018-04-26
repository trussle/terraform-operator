
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
	LoadBalancer	LoadBalancer	`json:"load_balancer"`
	PlacementStrategy	PlacementStrategy	`json:"placement_strategy"`
	ServiceRegistries	ServiceRegistries	`json:"service_registries"`
	DeploymentMinimumHealthyPercent	int	`json:"deployment_minimum_healthy_percent"`
	NetworkConfiguration	[]NetworkConfiguration	`json:"network_configuration"`
	LaunchType	string	`json:"launch_type"`
	DesiredCount	int	`json:"desired_count"`
	HealthCheckGracePeriodSeconds	int	`json:"health_check_grace_period_seconds"`
	PlacementConstraints	PlacementConstraints	`json:"placement_constraints"`
	Name	string	`json:"name"`
	DeploymentMaximumPercent	int	`json:"deployment_maximum_percent"`
	TaskDefinition	string	`json:"task_definition"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsServiceList is a list of AwsEcsService resources
type AwsEcsServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsService	`json:"items"`
}


// PlacementConstraints is a PlacementConstraints
type PlacementConstraints struct {
	Type	string	`json:"type"`
	Expression	string	`json:"expression"`
}

// LoadBalancer is a LoadBalancer
type LoadBalancer struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	ContainerName	string	`json:"container_name"`
	ContainerPort	int	`json:"container_port"`
	ElbName	string	`json:"elb_name"`
}

// PlacementStrategy is a PlacementStrategy
type PlacementStrategy struct {
	Type	string	`json:"type"`
	Field	string	`json:"field"`
}

// ServiceRegistries is a ServiceRegistries
type ServiceRegistries struct {
	Port	int	`json:"port"`
	RegistryArn	string	`json:"registry_arn"`
}

// NetworkConfiguration is a NetworkConfiguration
type NetworkConfiguration struct {
	Subnets	string	`json:"subnets"`
	AssignPublicIp	bool	`json:"assign_public_ip"`
	SecurityGroups	string	`json:"security_groups"`
}
