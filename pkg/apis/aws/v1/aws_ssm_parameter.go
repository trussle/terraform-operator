
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmParameter struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmParameterSpec	`json:"spec"`
}

type AwsSsmParameterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmParameter	`json:"items"`
}

type AwsSsmParameterSpec struct {
	Overwrite	bool	`json:"overwrite"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Type	string	`json:"type"`
	Value	string	`json:"value"`
	AllowedPattern	string	`json:"allowed_pattern"`
	Tags	map[string]interface{}	`json:"tags"`
}
