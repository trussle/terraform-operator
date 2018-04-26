
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorAssessmentTarget describes a AwsInspectorAssessmentTarget resource
type AwsInspectorAssessmentTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInspectorAssessmentTargetSpec	`json:"spec"`
}


// AwsInspectorAssessmentTargetSpec is the spec for a AwsInspectorAssessmentTarget Resource
type AwsInspectorAssessmentTargetSpec struct {
	ResourceGroupArn	string	`json:"resource_group_arn"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorAssessmentTargetList is a list of AwsInspectorAssessmentTarget resources
type AwsInspectorAssessmentTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInspectorAssessmentTarget	`json:"items"`
}

