
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
	Artifacts	string	`json:"artifacts"`
	Cache	[]kURUpiFv	`json:"cache"`
	Name	string	`json:"name"`
	Source	string	`json:"source"`
	BuildTimeout	int	`json:"build_timeout"`
	Tags	map[string]???	`json:"tags"`
	Environment	string	`json:"environment"`
	Timeout	int	`json:"timeout"`
	VpcConfig	[]IZRgBmyA	`json:"vpc_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodebuildProjectList is a list of AwsCodebuildProject resources
type AwsCodebuildProjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodebuildProject	`json:"items"`
}


// kURUpiFv is a kURUpiFv
type kURUpiFv struct {
	Type	string	`json:"type"`
	Location	string	`json:"location"`
}

// IZRgBmyA is a IZRgBmyA
type IZRgBmyA struct {
	VpcId	string	`json:"vpc_id"`
	Subnets	string	`json:"subnets"`
	SecurityGroupIds	string	`json:"security_group_ids"`
}
