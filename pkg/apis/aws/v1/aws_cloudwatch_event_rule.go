
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventRule describes a AwsCloudwatchEventRule resource
type AwsCloudwatchEventRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchEventRuleSpec	`json:"spec"`
}


// AwsCloudwatchEventRuleSpec is the spec for a AwsCloudwatchEventRule Resource
type AwsCloudwatchEventRuleSpec struct {
	Name	string	`json:"name"`
	ScheduleExpression	string	`json:"schedule_expression"`
	EventPattern	string	`json:"event_pattern"`
	Description	string	`json:"description"`
	RoleArn	string	`json:"role_arn"`
	IsEnabled	bool	`json:"is_enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventRuleList is a list of AwsCloudwatchEventRule resources
type AwsCloudwatchEventRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventRule	`json:"items"`
}

