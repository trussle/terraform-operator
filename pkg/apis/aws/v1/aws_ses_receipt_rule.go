
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesReceiptRuleSpec	`json:"spec"`
}

type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesReceiptRule	`json:"items"`
}

type AwsSesReceiptRuleSpec struct {
	Name	string	`json:"name"`
	BounceAction	interface{}	`json:"bounce_action"`
	RuleSetName	string	`json:"rule_set_name"`
	LambdaAction	interface{}	`json:"lambda_action"`
	StopAction	interface{}	`json:"stop_action"`
	SnsAction	interface{}	`json:"sns_action"`
	After	string	`json:"after"`
	Recipients	interface{}	`json:"recipients"`
	AddHeaderAction	interface{}	`json:"add_header_action"`
	S3Action	interface{}	`json:"s3_action"`
	WorkmailAction	interface{}	`json:"workmail_action"`
}
