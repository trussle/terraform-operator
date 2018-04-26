
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
	LaunchSpecification	LaunchSpecification	`json:"launch_specification"`
	ExcessCapacityTerminationPolicy	string	`json:"excess_capacity_termination_policy"`
	ValidUntil	string	`json:"valid_until"`
	TargetCapacity	int	`json:"target_capacity"`
	AllocationStrategy	string	`json:"allocation_strategy"`
	SpotPrice	string	`json:"spot_price"`
	IamFleetRole	string	`json:"iam_fleet_role"`
	ReplaceUnhealthyInstances	bool	`json:"replace_unhealthy_instances"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	TerminateInstancesWithExpiration	bool	`json:"terminate_instances_with_expiration"`
	ValidFrom	string	`json:"valid_from"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotFleetRequestList is a list of AwsSpotFleetRequest resources
type AwsSpotFleetRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotFleetRequest	`json:"items"`
}


// EphemeralBlockDevice is a EphemeralBlockDevice
type EphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}

// RootBlockDevice is a RootBlockDevice
type RootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// EbsBlockDevice is a EbsBlockDevice
type EbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}

// LaunchSpecification is a LaunchSpecification
type LaunchSpecification struct {
	Monitoring	bool	`json:"monitoring"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	InstanceType	string	`json:"instance_type"`
	Ami	string	`json:"ami"`
	PlacementTenancy	string	`json:"placement_tenancy"`
	SpotPrice	string	`json:"spot_price"`
	UserData	string	`json:"user_data"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	WeightedCapacity	string	`json:"weighted_capacity"`
	Tags	map[string]string	`json:"tags"`
}
