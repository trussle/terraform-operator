
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsConfigConfigRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsConfigConfigRuleSpec	`json:"spec"`
}

type AwsConfigConfigRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigRule	`json:"items"`
}

type AwsConfigConfigRuleSpec struct {
	Description	string	`json:"description"`
	InputParameters	string	`json:"input_parameters"`
	MaximumExecutionFrequency	string	`json:"maximum_execution_frequency"`
	Scope	[]interface{}	`json:"scope"`
	Source	[]interface{}	`json:"source"`
	Name	string	`json:"name"`
}
