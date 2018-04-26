
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
	InputParameters	string	`json:"input_parameters"`
	MaximumExecutionFrequency	string	`json:"maximum_execution_frequency"`
	Scope	[]hTtUTuWJ	`json:"scope"`
	Source	[]RGiQjOTc	`json:"source"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigConfigRuleList is a list of AwsConfigConfigRule resources
type AwsConfigConfigRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigConfigRule	`json:"items"`
}


// hTtUTuWJ is a hTtUTuWJ
type hTtUTuWJ struct {
	ComplianceResourceId	string	`json:"compliance_resource_id"`
	ComplianceResourceTypes	string	`json:"compliance_resource_types"`
	TagKey	string	`json:"tag_key"`
	TagValue	string	`json:"tag_value"`
}

// RGiQjOTc is a RGiQjOTc
type RGiQjOTc struct {
	SourceDetail	string	`json:"source_detail"`
	SourceIdentifier	string	`json:"source_identifier"`
	Owner	string	`json:"owner"`
}
