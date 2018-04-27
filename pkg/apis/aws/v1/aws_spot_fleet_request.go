
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
	LaunchSpecification	AwsSpotFleetRequestLaunchSpecification	`json:"launch_specification"`
	ValidUntil	string	`json:"valid_until"`
	ReplaceUnhealthyInstances	bool	`json:"replace_unhealthy_instances"`
	TargetCapacity	int	`json:"target_capacity"`
	AllocationStrategy	string	`json:"allocation_strategy"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	ValidFrom	string	`json:"valid_from"`
	IamFleetRole	string	`json:"iam_fleet_role"`
	SpotPrice	string	`json:"spot_price"`
	ExcessCapacityTerminationPolicy	string	`json:"excess_capacity_termination_policy"`
	TerminateInstancesWithExpiration	bool	`json:"terminate_instances_with_expiration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotFleetRequestList is a list of AwsSpotFleetRequest resources
type AwsSpotFleetRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotFleetRequest	`json:"items"`
}


// AwsSpotFleetRequestRootBlockDevice is a AwsSpotFleetRequestRootBlockDevice
type AwsSpotFleetRequestRootBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// AwsSpotFleetRequestEphemeralBlockDevice is a AwsSpotFleetRequestEphemeralBlockDevice
type AwsSpotFleetRequestEphemeralBlockDevice struct {
	DeviceName	string	`json:"device_name"`
	VirtualName	string	`json:"virtual_name"`
}

// AwsSpotFleetRequestEbsBlockDevice is a AwsSpotFleetRequestEbsBlockDevice
type AwsSpotFleetRequestEbsBlockDevice struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
	DeviceName	string	`json:"device_name"`
}

// AwsSpotFleetRequestLaunchSpecification is a AwsSpotFleetRequestLaunchSpecification
type AwsSpotFleetRequestLaunchSpecification struct {
	Ami	string	`json:"ami"`
	Monitoring	bool	`json:"monitoring"`
	AssociatePublicIpAddress	bool	`json:"associate_public_ip_address"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	Tags	map[string]string	`json:"tags"`
	InstanceType	string	`json:"instance_type"`
	SpotPrice	string	`json:"spot_price"`
	WeightedCapacity	string	`json:"weighted_capacity"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	PlacementTenancy	string	`json:"placement_tenancy"`
	UserData	string	`json:"user_data"`
}
