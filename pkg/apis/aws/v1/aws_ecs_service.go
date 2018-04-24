
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsService struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEcsServiceSpec	`json:"spec"`
}

type AwsEcsServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsService	`json:"items"`
}

type AwsEcsServiceSpec struct {
	PlacementConstraints	interface{}	`json:"placement_constraints"`
	Name	string	`json:"name"`
	HealthCheckGracePeriodSeconds	int	`json:"health_check_grace_period_seconds"`
	PlacementStrategy	interface{}	`json:"placement_strategy"`
	LoadBalancer	interface{}	`json:"load_balancer"`
	TaskDefinition	string	`json:"task_definition"`
	DesiredCount	int	`json:"desired_count"`
	ServiceRegistries	interface{}	`json:"service_registries"`
	NetworkConfiguration	[]interface{}	`json:"network_configuration"`
	LaunchType	string	`json:"launch_type"`
	DeploymentMaximumPercent	int	`json:"deployment_maximum_percent"`
	DeploymentMinimumHealthyPercent	int	`json:"deployment_minimum_healthy_percent"`
}
