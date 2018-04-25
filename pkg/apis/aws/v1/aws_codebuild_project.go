
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Name	string	`json:"name"`
	Source	string	`json:"source"`
	Timeout	int	`json:"timeout"`
	VpcConfig	[]interface{}	`json:"vpc_config"`
	Artifacts	string	`json:"artifacts"`
	Environment	string	`json:"environment"`
	BuildTimeout	int	`json:"build_timeout"`
	Tags	map[string]interface{}	`json:"tags"`
	Cache	[]interface{}	`json:"cache"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodebuildProjectList is a list of AwsCodebuildProject resources
type AwsCodebuildProjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodebuildProject	`json:"items"`
}

