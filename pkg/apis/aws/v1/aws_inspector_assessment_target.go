
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInspectorAssessmentTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInspectorAssessmentTargetSpec	`json:"spec"`
}

type AwsInspectorAssessmentTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInspectorAssessmentTarget	`json:"items"`
}

type AwsInspectorAssessmentTargetSpec struct {
	Name	string	`json:"name"`
	ResourceGroupArn	string	`json:"resource_group_arn"`
}
