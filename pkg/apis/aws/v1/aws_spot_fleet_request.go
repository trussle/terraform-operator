
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotFleetRequest describes a AwsSpotFleetRequest resource
type AwsSpotFleetRequest struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSpotFleetRequestSpec	`json:"spec"`
}


// AwsSpotFleetRequestSpec is the spec for a AwsSpotFleetRequest Resource
type AwsSpotFleetRequestSpec struct {
	SpotPrice	string	`json:"spot_price"`
	TerminateInstancesWithExpiration	bool	`json:"terminate_instances_with_expiration"`
	ValidUntil	string	`json:"valid_until"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	AllocationStrategy	string	`json:"allocation_strategy"`
	ExcessCapacityTerminationPolicy	string	`json:"excess_capacity_termination_policy"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	ValidFrom	string	`json:"valid_from"`
	IamFleetRole	string	`json:"iam_fleet_role"`
	ReplaceUnhealthyInstances	bool	`json:"replace_unhealthy_instances"`
	LaunchSpecification	string	`json:"launch_specification"`
	TargetCapacity	int	`json:"target_capacity"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotFleetRequestList is a list of AwsSpotFleetRequest resources
type AwsSpotFleetRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotFleetRequest	`json:"items"`
}

