
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodebuildProject struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodebuildProjectSpec	`json:"spec"`
}

type AwsCodebuildProjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodebuildProject	`json:"items"`
}

type AwsCodebuildProjectSpec struct {
	Environment	interface{}	`json:"environment"`
	Source	interface{}	`json:"source"`
	Timeout	int	`json:"timeout"`
	BuildTimeout	int	`json:"build_timeout"`
	VpcConfig	[]interface{}	`json:"vpc_config"`
	Artifacts	interface{}	`json:"artifacts"`
	Cache	[]interface{}	`json:"cache"`
	Tags	map[string]interface{}	`json:"tags"`
	Name	string	`json:"name"`
}
