
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchEventRuleSpec	`json:"spec"`
}

type AwsCloudwatchEventRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventRule	`json:"items"`
}

type AwsCloudwatchEventRuleSpec struct {
	Name	string	`json:"name"`
	ScheduleExpression	string	`json:"schedule_expression"`
	EventPattern	string	`json:"event_pattern"`
	Description	string	`json:"description"`
	RoleArn	string	`json:"role_arn"`
	IsEnabled	bool	`json:"is_enabled"`
}
