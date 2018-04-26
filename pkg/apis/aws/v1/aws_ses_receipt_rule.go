
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
	Name	string	`json:"name"`
	RuleSetName	string	`json:"rule_set_name"`
	AddHeaderAction	Generic	`json:"add_header_action"`
	S3Action	Generic	`json:"s3_action"`
	After	string	`json:"after"`
	Recipients	Generic	`json:"recipients"`
	BounceAction	Generic	`json:"bounce_action"`
	StopAction	Generic	`json:"stop_action"`
	WorkmailAction	Generic	`json:"workmail_action"`
	LambdaAction	Generic	`json:"lambda_action"`
	SnsAction	Generic	`json:"sns_action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRuleList is a list of AwsSesReceiptRule resources
type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesReceiptRule	`json:"items"`
}

