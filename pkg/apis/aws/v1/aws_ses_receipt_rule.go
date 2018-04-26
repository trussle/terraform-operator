
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	RuleSetName	string	`json:"rule_set_name"`
	After	string	`json:"after"`
	Recipients	string	`json:"recipients"`
	LambdaAction	LambdaAction	`json:"lambda_action"`
	StopAction	StopAction	`json:"stop_action"`
	Name	string	`json:"name"`
	WorkmailAction	WorkmailAction	`json:"workmail_action"`
	AddHeaderAction	AddHeaderAction	`json:"add_header_action"`
	BounceAction	BounceAction	`json:"bounce_action"`
	S3Action	S3Action	`json:"s3_action"`
	SnsAction	SnsAction	`json:"sns_action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRuleList is a list of AwsSesReceiptRule resources
type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesReceiptRule	`json:"items"`
}


// StopAction is a StopAction
type StopAction struct {
	Scope	string	`json:"scope"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}

// WorkmailAction is a WorkmailAction
type WorkmailAction struct {
	OrganizationArn	string	`json:"organization_arn"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}

// AddHeaderAction is a AddHeaderAction
type AddHeaderAction struct {
	HeaderName	string	`json:"header_name"`
	HeaderValue	string	`json:"header_value"`
	Position	int	`json:"position"`
}

// BounceAction is a BounceAction
type BounceAction struct {
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
	Message	string	`json:"message"`
	Sender	string	`json:"sender"`
	SmtpReplyCode	string	`json:"smtp_reply_code"`
	StatusCode	string	`json:"status_code"`
}

// S3Action is a S3Action
type S3Action struct {
	BucketName	string	`json:"bucket_name"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	ObjectKeyPrefix	string	`json:"object_key_prefix"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}

// SnsAction is a SnsAction
type SnsAction struct {
	Position	int	`json:"position"`
	TopicArn	string	`json:"topic_arn"`
}

// LambdaAction is a LambdaAction
type LambdaAction struct {
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
	FunctionArn	string	`json:"function_arn"`
}
