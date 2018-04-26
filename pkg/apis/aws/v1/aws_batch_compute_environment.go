
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchComputeEnvironment describes a AwsBatchComputeEnvironment resource
type AwsBatchComputeEnvironment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsBatchComputeEnvironmentSpec	`json:"spec"`
}


// AwsBatchComputeEnvironmentSpec is the spec for a AwsBatchComputeEnvironment Resource
type AwsBatchComputeEnvironmentSpec struct {
	ServiceRole	string	`json:"service_role"`
	ComputeResources	[]ComputeResources	`json:"compute_resources"`
	State	string	`json:"state"`
	Type	string	`json:"type"`
	ComputeEnvironmentName	string	`json:"compute_environment_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchComputeEnvironmentList is a list of AwsBatchComputeEnvironment resources
type AwsBatchComputeEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchComputeEnvironment	`json:"items"`
}


// ComputeResources is a ComputeResources
type ComputeResources struct {
	Subnets	string	`json:"subnets"`
	Tags	map[string]string	`json:"tags"`
	Type	string	`json:"type"`
	MinVcpus	int	`json:"min_vcpus"`
	SpotIamFleetRole	string	`json:"spot_iam_fleet_role"`
	Ec2KeyPair	string	`json:"ec2_key_pair"`
	ImageId	string	`json:"image_id"`
	InstanceRole	string	`json:"instance_role"`
	InstanceType	string	`json:"instance_type"`
	MaxVcpus	int	`json:"max_vcpus"`
	SecurityGroupIds	string	`json:"security_group_ids"`
	BidPercentage	int	`json:"bid_percentage"`
	DesiredVcpus	int	`json:"desired_vcpus"`
}
