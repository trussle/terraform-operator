
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
	BounceAction	AwsSesReceiptRuleBounceAction	`json:"bounce_action"`
	LambdaAction	AwsSesReceiptRuleLambdaAction	`json:"lambda_action"`
	Recipients	string	`json:"recipients"`
	AddHeaderAction	AwsSesReceiptRuleAddHeaderAction	`json:"add_header_action"`
	Name	string	`json:"name"`
	RuleSetName	string	`json:"rule_set_name"`
	S3Action	AwsSesReceiptRuleS3Action	`json:"s3_action"`
	SnsAction	AwsSesReceiptRuleSnsAction	`json:"sns_action"`
	StopAction	AwsSesReceiptRuleStopAction	`json:"stop_action"`
	WorkmailAction	AwsSesReceiptRuleWorkmailAction	`json:"workmail_action"`
	After	string	`json:"after"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRuleList is a list of AwsSesReceiptRule resources
type AwsSesReceiptRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesReceiptRule	`json:"items"`
}


// AwsSesReceiptRuleWorkmailAction is a AwsSesReceiptRuleWorkmailAction
type AwsSesReceiptRuleWorkmailAction struct {
	OrganizationArn	string	`json:"organization_arn"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}

// AwsSesReceiptRuleBounceAction is a AwsSesReceiptRuleBounceAction
type AwsSesReceiptRuleBounceAction struct {
	SmtpReplyCode	string	`json:"smtp_reply_code"`
	StatusCode	string	`json:"status_code"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
	Message	string	`json:"message"`
	Sender	string	`json:"sender"`
}

// AwsSesReceiptRuleLambdaAction is a AwsSesReceiptRuleLambdaAction
type AwsSesReceiptRuleLambdaAction struct {
	FunctionArn	string	`json:"function_arn"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}

// AwsSesReceiptRuleAddHeaderAction is a AwsSesReceiptRuleAddHeaderAction
type AwsSesReceiptRuleAddHeaderAction struct {
	Position	int	`json:"position"`
	HeaderName	string	`json:"header_name"`
	HeaderValue	string	`json:"header_value"`
}

// AwsSesReceiptRuleS3Action is a AwsSesReceiptRuleS3Action
type AwsSesReceiptRuleS3Action struct {
	BucketName	string	`json:"bucket_name"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	ObjectKeyPrefix	string	`json:"object_key_prefix"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}

// AwsSesReceiptRuleSnsAction is a AwsSesReceiptRuleSnsAction
type AwsSesReceiptRuleSnsAction struct {
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}

// AwsSesReceiptRuleStopAction is a AwsSesReceiptRuleStopAction
type AwsSesReceiptRuleStopAction struct {
	Scope	string	`json:"scope"`
	TopicArn	string	`json:"topic_arn"`
	Position	int	`json:"position"`
}
