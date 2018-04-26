
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
	DeploymentMaximumPercent	int	`json:"deployment_maximum_percent"`
	DeploymentMinimumHealthyPercent	int	`json:"deployment_minimum_healthy_percent"`
	LaunchType	string	`json:"launch_type"`
	NetworkConfiguration	[]Generic	`json:"network_configuration"`
	PlacementConstraints	Generic	`json:"placement_constraints"`
	ServiceRegistries	Generic	`json:"service_registries"`
	TaskDefinition	string	`json:"task_definition"`
	LoadBalancer	Generic	`json:"load_balancer"`
	PlacementStrategy	Generic	`json:"placement_strategy"`
	Name	string	`json:"name"`
	DesiredCount	int	`json:"desired_count"`
	HealthCheckGracePeriodSeconds	int	`json:"health_check_grace_period_seconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsServiceList is a list of AwsEcsService resources
type AwsEcsServiceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEcsService	`json:"items"`
}

