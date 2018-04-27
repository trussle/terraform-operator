
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
	LaunchType	string	`json:"launch_type"`
	DeploymentMinimumHealthyPercent	int	`json:"deployment_minimum_healthy_percent"`
	HealthCheckGracePeriodSeconds	int	`json:"health_check_grace_period_seconds"`
	DeploymentMaximumPercent	int	`json:"deployment_maximum_percent"`
	NetworkConfiguration	[]AwsEcsServiceNetworkConfiguration	`json:"network_configuration"`
	PlacementConstraints	AwsEcsServicePlacementConstraints	`json:"placement_constraints"`
	Name	string	`json:"name"`
	LoadBalancer	AwsEcsServiceLoadBalancer	`json:"load_balancer"`
	PlacementStrategy	AwsEcsServicePlacementStrategy	`json:"placement_strategy"`
	ServiceRegistries	AwsEcsServiceServiceRegistries	`json:"service_registries"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsServiceList is a list of AwsEcsService resources
type AwsEcsServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsService	`json:"items"`
}


// AwsEcsServicePlacementConstraints is a AwsEcsServicePlacementConstraints
type AwsEcsServicePlacementConstraints struct {
	Type	string	`json:"type"`
	Expression	string	`json:"expression"`
}

// AwsEcsServiceLoadBalancer is a AwsEcsServiceLoadBalancer
type AwsEcsServiceLoadBalancer struct {
	ContainerPort	int	`json:"container_port"`
	ElbName	string	`json:"elb_name"`
	TargetGroupArn	string	`json:"target_group_arn"`
	ContainerName	string	`json:"container_name"`
}

// AwsEcsServicePlacementStrategy is a AwsEcsServicePlacementStrategy
type AwsEcsServicePlacementStrategy struct {
	Type	string	`json:"type"`
	Field	string	`json:"field"`
}

// AwsEcsServiceServiceRegistries is a AwsEcsServiceServiceRegistries
type AwsEcsServiceServiceRegistries struct {
	Port	int	`json:"port"`
	RegistryArn	string	`json:"registry_arn"`
}

// AwsEcsServiceNetworkConfiguration is a AwsEcsServiceNetworkConfiguration
type AwsEcsServiceNetworkConfiguration struct {
	SecurityGroups	string	`json:"security_groups"`
	Subnets	string	`json:"subnets"`
	AssignPublicIp	bool	`json:"assign_public_ip"`
}
