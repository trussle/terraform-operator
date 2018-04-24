
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInspectorAssessmentTemplateSpec	`json:"spec"`
}

type AwsInspectorAssessmentTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInspectorAssessmentTemplate	`json:"items"`
}

type AwsInspectorAssessmentTemplateSpec struct {
	Name	string	`json:"name"`
	TargetArn	string	`json:"target_arn"`
	Duration	int	`json:"duration"`
	RulesPackageArns	interface{}	`json:"rules_package_arns"`
}
