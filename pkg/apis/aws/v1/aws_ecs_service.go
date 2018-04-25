
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Name	string	`json:"name"`
	LoadBalancer	string	`json:"load_balancer"`
	PlacementStrategy	string	`json:"placement_strategy"`
	PlacementConstraints	string	`json:"placement_constraints"`
	ServiceRegistries	string	`json:"service_registries"`
	DeploymentMaximumPercent	int	`json:"deployment_maximum_percent"`
	DeploymentMinimumHealthyPercent	int	`json:"deployment_minimum_healthy_percent"`
	NetworkConfiguration	[]interface{}	`json:"network_configuration"`
	TaskDefinition	string	`json:"task_definition"`
	HealthCheckGracePeriodSeconds	int	`json:"health_check_grace_period_seconds"`
	DesiredCount	int	`json:"desired_count"`
	LaunchType	string	`json:"launch_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsServiceList is a list of AwsEcsService resources
type AwsEcsServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsService	`json:"items"`
}

