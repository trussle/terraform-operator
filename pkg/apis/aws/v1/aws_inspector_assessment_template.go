
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorAssessmentTemplate describes a AwsInspectorAssessmentTemplate resource
type AwsInspectorAssessmentTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInspectorAssessmentTemplateSpec	`json:"spec"`
}


// AwsInspectorAssessmentTemplateSpec is the spec for a AwsInspectorAssessmentTemplate Resource
type AwsInspectorAssessmentTemplateSpec struct {
	Name	string	`json:"name"`
	TargetArn	string	`json:"target_arn"`
	Duration	int	`json:"duration"`
	RulesPackageArns	string	`json:"rules_package_arns"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorAssessmentTemplateList is a list of AwsInspectorAssessmentTemplate resources
type AwsInspectorAssessmentTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInspectorAssessmentTemplate	`json:"items"`
}

