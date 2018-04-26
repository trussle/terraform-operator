
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Source	[]Source	`json:"source"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	InputParameters	string	`json:"input_parameters"`
	MaximumExecutionFrequency	string	`json:"maximum_execution_frequency"`
	Scope	[]Scope	`json:"scope"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigRuleList is a list of AwsConfigConfigRule resources
type AwsConfigConfigRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigRule	`json:"items"`
}


// SourceDetail is a SourceDetail
type SourceDetail struct {
	MaximumExecutionFrequency	string	`json:"maximum_execution_frequency"`
	MessageType	string	`json:"message_type"`
	EventSource	string	`json:"event_source"`
}

// Source is a Source
type Source struct {
	Owner	string	`json:"owner"`
	SourceDetail	SourceDetail	`json:"source_detail"`
	SourceIdentifier	string	`json:"source_identifier"`
}

// Scope is a Scope
type Scope struct {
	TagKey	string	`json:"tag_key"`
	TagValue	string	`json:"tag_value"`
	ComplianceResourceId	string	`json:"compliance_resource_id"`
	ComplianceResourceTypes	string	`json:"compliance_resource_types"`
}
