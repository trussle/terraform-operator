
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmParameter describes a AwsSsmParameter resource
type AwsSsmParameter struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmParameterSpec	`json:"spec"`
}


// AwsSsmParameterSpec is the spec for a AwsSsmParameter Resource
type AwsSsmParameterSpec struct {
	Type	string	`json:"type"`
	Value	string	`json:"value"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Overwrite	bool	`json:"overwrite"`
	AllowedPattern	string	`json:"allowed_pattern"`
	Tags	map[string]Generic	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmParameterList is a list of AwsSsmParameter resources
type AwsSsmParameterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmParameter	`json:"items"`
}

