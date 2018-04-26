
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
	Artifacts	Artifacts	`json:"artifacts"`
	BuildTimeout	int	`json:"build_timeout"`
	VpcConfig	[]VpcConfig	`json:"vpc_config"`
	Timeout	int	`json:"timeout"`
	Tags	map[string]string	`json:"tags"`
	Cache	[]Cache	`json:"cache"`
	Environment	Environment	`json:"environment"`
	Name	string	`json:"name"`
	Source	Source	`json:"source"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodebuildProjectList is a list of AwsCodebuildProject resources
type AwsCodebuildProjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodebuildProject	`json:"items"`
}


// VpcConfig is a VpcConfig
type VpcConfig struct {
	VpcId	string	`json:"vpc_id"`
	Subnets	string	`json:"subnets"`
	SecurityGroupIds	string	`json:"security_group_ids"`
}

// Cache is a Cache
type Cache struct {
	Type	string	`json:"type"`
	Location	string	`json:"location"`
}

// EnvironmentVariable is a EnvironmentVariable
type EnvironmentVariable struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}

// Environment is a Environment
type Environment struct {
	Image	string	`json:"image"`
	Type	string	`json:"type"`
	PrivilegedMode	bool	`json:"privileged_mode"`
	ComputeType	string	`json:"compute_type"`
}

// Auth is a Auth
type Auth struct {
	Resource	string	`json:"resource"`
	Type	string	`json:"type"`
}

// Source is a Source
type Source struct {
	Type	string	`json:"type"`
	Auth	Auth	`json:"auth"`
	Buildspec	string	`json:"buildspec"`
	Location	string	`json:"location"`
}

// Artifacts is a Artifacts
type Artifacts struct {
	Packaging	string	`json:"packaging"`
	Path	string	`json:"path"`
	Type	string	`json:"type"`
	Name	string	`json:"name"`
	Location	string	`json:"location"`
	NamespaceType	string	`json:"namespace_type"`
}
