
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketNotification describes a AwsS3BucketNotification resource
type AwsS3BucketNotification struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketNotificationSpec	`json:"spec"`
}


// AwsS3BucketNotificationSpec is the spec for a AwsS3BucketNotification Resource
type AwsS3BucketNotificationSpec struct {
	Bucket	string	`json:"bucket"`
	Topic	[]nyNSLTvm	`json:"topic"`
	Queue	[]wZLbKtzi	`json:"queue"`
	LambdaFunction	[]YpTpzlFP	`json:"lambda_function"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketNotificationList is a list of AwsS3BucketNotification resources
type AwsS3BucketNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketNotification	`json:"items"`
}


// nyNSLTvm is a nyNSLTvm
type nyNSLTvm struct {
	Events	string	`json:"events"`
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
	TopicArn	string	`json:"topic_arn"`
}

// wZLbKtzi is a wZLbKtzi
type wZLbKtzi struct {
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
	QueueArn	string	`json:"queue_arn"`
	Events	string	`json:"events"`
}

// YpTpzlFP is a YpTpzlFP
type YpTpzlFP struct {
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
	LambdaFunctionArn	string	`json:"lambda_function_arn"`
	Events	string	`json:"events"`
}
