
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodebuildProject describes a AwsCodebuildProject resource
type AwsCodebuildProject struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodebuildProjectSpec	`json:"spec"`
}


// AwsCodebuildProjectSpec is the spec for a AwsCodebuildProject Resource
type AwsCodebuildProjectSpec struct {
	VpcConfig	[]AwsCodebuildProjectVpcConfig	`json:"vpc_config"`
	Cache	[]AwsCodebuildProjectCache	`json:"cache"`
	Environment	AwsCodebuildProjectEnvironment	`json:"environment"`
	Source	AwsCodebuildProjectSource	`json:"source"`
	Timeout	int	`json:"timeout"`
	BuildTimeout	int	`json:"build_timeout"`
	Tags	map[string]string	`json:"tags"`
	Artifacts	AwsCodebuildProjectArtifacts	`json:"artifacts"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodebuildProjectList is a list of AwsCodebuildProject resources
type AwsCodebuildProjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodebuildProject	`json:"items"`
}


// AwsCodebuildProjectEnvironment is a AwsCodebuildProjectEnvironment
type AwsCodebuildProjectEnvironment struct {
	Image	string	`json:"image"`
	Type	string	`json:"type"`
	PrivilegedMode	bool	`json:"privileged_mode"`
	ComputeType	string	`json:"compute_type"`
}

// AwsCodebuildProjectAuth is a AwsCodebuildProjectAuth
type AwsCodebuildProjectAuth struct {
	Resource	string	`json:"resource"`
	Type	string	`json:"type"`
}

// AwsCodebuildProjectSource is a AwsCodebuildProjectSource
type AwsCodebuildProjectSource struct {
	Location	string	`json:"location"`
	Type	string	`json:"type"`
	Auth	AwsCodebuildProjectAuth	`json:"auth"`
	Buildspec	string	`json:"buildspec"`
}

// AwsCodebuildProjectArtifacts is a AwsCodebuildProjectArtifacts
type AwsCodebuildProjectArtifacts struct {
	NamespaceType	string	`json:"namespace_type"`
	Packaging	string	`json:"packaging"`
	Path	string	`json:"path"`
	Type	string	`json:"type"`
	Name	string	`json:"name"`
	Location	string	`json:"location"`
}

// AwsCodebuildProjectVpcConfig is a AwsCodebuildProjectVpcConfig
type AwsCodebuildProjectVpcConfig struct {
	Subnets	string	`json:"subnets"`
	SecurityGroupIds	string	`json:"security_group_ids"`
	VpcId	string	`json:"vpc_id"`
}

// AwsCodebuildProjectCache is a AwsCodebuildProjectCache
type AwsCodebuildProjectCache struct {
	Type	string	`json:"type"`
	Location	string	`json:"location"`
}

// AwsCodebuildProjectEnvironmentVariable is a AwsCodebuildProjectEnvironmentVariable
type AwsCodebuildProjectEnvironmentVariable struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
