
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
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	InputParameters	string	`json:"input_parameters"`
	MaximumExecutionFrequency	string	`json:"maximum_execution_frequency"`
	Scope	[]AwsConfigConfigRuleScope	`json:"scope"`
	Source	[]AwsConfigConfigRuleSource	`json:"source"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigRuleList is a list of AwsConfigConfigRule resources
type AwsConfigConfigRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigRule	`json:"items"`
}


// AwsConfigConfigRuleSource is a AwsConfigConfigRuleSource
type AwsConfigConfigRuleSource struct {
	SourceIdentifier	string	`json:"source_identifier"`
	Owner	string	`json:"owner"`
	SourceDetail	AwsConfigConfigRuleSourceDetail	`json:"source_detail"`
}

// AwsConfigConfigRuleScope is a AwsConfigConfigRuleScope
type AwsConfigConfigRuleScope struct {
	ComplianceResourceId	string	`json:"compliance_resource_id"`
	ComplianceResourceTypes	string	`json:"compliance_resource_types"`
	TagKey	string	`json:"tag_key"`
	TagValue	string	`json:"tag_value"`
}

// AwsConfigConfigRuleSourceDetail is a AwsConfigConfigRuleSourceDetail
type AwsConfigConfigRuleSourceDetail struct {
	EventSource	string	`json:"event_source"`
	MaximumExecutionFrequency	string	`json:"maximum_execution_frequency"`
	MessageType	string	`json:"message_type"`
}
