
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigRule describes a AwsConfigConfigRule resource
type AwsConfigConfigRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsConfigConfigRuleSpec	`json:"spec"`
}


// AwsConfigConfigRuleSpec is the spec for a AwsConfigConfigRule Resource
type AwsConfigConfigRuleSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	InputParameters	string	`json:"input_parameters"`
	MaximumExecutionFrequency	string	`json:"maximum_execution_frequency"`
	Scope	[]interface{}	`json:"scope"`
	Source	[]interface{}	`json:"source"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigRuleList is a list of AwsConfigConfigRule resources
type AwsConfigConfigRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigRule	`json:"items"`
}

