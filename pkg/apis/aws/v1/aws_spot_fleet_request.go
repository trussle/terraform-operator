
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotFleetRequest struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSpotFleetRequestSpec	`json:"spec"`
}

type AwsSpotFleetRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotFleetRequest	`json:"items"`
}

type AwsSpotFleetRequestSpec struct {
	ReplaceUnhealthyInstances	bool	`json:"replace_unhealthy_instances"`
	LaunchSpecification	interface{}	`json:"launch_specification"`
	IamFleetRole	string	`json:"iam_fleet_role"`
	AllocationStrategy	string	`json:"allocation_strategy"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	ExcessCapacityTerminationPolicy	string	`json:"excess_capacity_termination_policy"`
	ValidFrom	string	`json:"valid_from"`
	ValidUntil	string	`json:"valid_until"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	TargetCapacity	int	`json:"target_capacity"`
	SpotPrice	string	`json:"spot_price"`
	TerminateInstancesWithExpiration	bool	`json:"terminate_instances_with_expiration"`
}
