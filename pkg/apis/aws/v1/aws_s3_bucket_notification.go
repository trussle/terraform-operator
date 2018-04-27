
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
	Topic	[]AwsS3BucketNotificationTopic	`json:"topic"`
	Queue	[]AwsS3BucketNotificationQueue	`json:"queue"`
	LambdaFunction	[]AwsS3BucketNotificationLambdaFunction	`json:"lambda_function"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketNotificationList is a list of AwsS3BucketNotification resources
type AwsS3BucketNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketNotification	`json:"items"`
}


// AwsS3BucketNotificationTopic is a AwsS3BucketNotificationTopic
type AwsS3BucketNotificationTopic struct {
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
	TopicArn	string	`json:"topic_arn"`
	Events	string	`json:"events"`
}

// AwsS3BucketNotificationQueue is a AwsS3BucketNotificationQueue
type AwsS3BucketNotificationQueue struct {
	QueueArn	string	`json:"queue_arn"`
	Events	string	`json:"events"`
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
}

// AwsS3BucketNotificationLambdaFunction is a AwsS3BucketNotificationLambdaFunction
type AwsS3BucketNotificationLambdaFunction struct {
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
	LambdaFunctionArn	string	`json:"lambda_function_arn"`
	Events	string	`json:"events"`
}
