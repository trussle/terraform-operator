
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
	ComputeEnvironmentName	string	`json:"compute_environment_name"`
	ComputeResources	[]AwsBatchComputeEnvironmentComputeResources	`json:"compute_resources"`
	ServiceRole	string	`json:"service_role"`
	State	string	`json:"state"`
	Type	string	`json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBatchComputeEnvironmentList is a list of AwsBatchComputeEnvironment resources
type AwsBatchComputeEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsBatchComputeEnvironment	`json:"items"`
}


// AwsBatchComputeEnvironmentComputeResources is a AwsBatchComputeEnvironmentComputeResources
type AwsBatchComputeEnvironmentComputeResources struct {
	InstanceRole	string	`json:"instance_role"`
	InstanceType	string	`json:"instance_type"`
	MaxVcpus	int	`json:"max_vcpus"`
	MinVcpus	int	`json:"min_vcpus"`
	Tags	map[string]string	`json:"tags"`
	Type	string	`json:"type"`
	BidPercentage	int	`json:"bid_percentage"`
	Ec2KeyPair	string	`json:"ec2_key_pair"`
	ImageId	string	`json:"image_id"`
	SecurityGroupIds	string	`json:"security_group_ids"`
	SpotIamFleetRole	string	`json:"spot_iam_fleet_role"`
	Subnets	string	`json:"subnets"`
	DesiredVcpus	int	`json:"desired_vcpus"`
}
