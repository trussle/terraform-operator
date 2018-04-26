
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	LaunchSpecification	string	`json:"launch_specification"`
	TargetCapacity	int	`json:"target_capacity"`
	ExcessCapacityTerminationPolicy	string	`json:"excess_capacity_termination_policy"`
	SpotPrice	string	`json:"spot_price"`
	ValidFrom	string	`json:"valid_from"`
	ValidUntil	string	`json:"valid_until"`
	IamFleetRole	string	`json:"iam_fleet_role"`
	ReplaceUnhealthyInstances	bool	`json:"replace_unhealthy_instances"`
	AllocationStrategy	string	`json:"allocation_strategy"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	TerminateInstancesWithExpiration	bool	`json:"terminate_instances_with_expiration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotFleetRequestList is a list of AwsSpotFleetRequest resources
type AwsSpotFleetRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotFleetRequest	`json:"items"`
}

