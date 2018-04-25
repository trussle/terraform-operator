
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRule describes a AwsSesReceiptRule resource
type AwsSesReceiptRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesReceiptRuleSpec	`json:"spec"`
}


// AwsSesReceiptRuleSpec is the spec for a AwsSesReceiptRule Resource
type AwsSesReceiptRuleSpec struct {
	After	string	`json:"after"`
	RuleSetName	string	`json:"rule_set_name"`
	S3Action	string	`json:"s3_action"`
	SnsAction	string	`json:"sns_action"`
	StopAction	string	`json:"stop_action"`
	BounceAction	string	`json:"bounce_action"`
	LambdaAction	string	`json:"lambda_action"`
	WorkmailAction	string	`json:"workmail_action"`
	Name	string	`json:"name"`
	Recipients	string	`json:"recipients"`
	AddHeaderAction	string	`json:"add_header_action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRuleList is a list of AwsSesReceiptRule resources
type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesReceiptRule	`json:"items"`
}

