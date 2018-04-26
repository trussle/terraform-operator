
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
	Topic	[]Topic	`json:"topic"`
	Queue	[]Queue	`json:"queue"`
	LambdaFunction	[]LambdaFunction	`json:"lambda_function"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketNotificationList is a list of AwsS3BucketNotification resources
type AwsS3BucketNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketNotification	`json:"items"`
}


// Topic is a Topic
type Topic struct {
	FilterSuffix	string	`json:"filter_suffix"`
	TopicArn	string	`json:"topic_arn"`
	Events	string	`json:"events"`
	FilterPrefix	string	`json:"filter_prefix"`
}

// Queue is a Queue
type Queue struct {
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
	QueueArn	string	`json:"queue_arn"`
	Events	string	`json:"events"`
}

// LambdaFunction is a LambdaFunction
type LambdaFunction struct {
	LambdaFunctionArn	string	`json:"lambda_function_arn"`
	Events	string	`json:"events"`
	FilterPrefix	string	`json:"filter_prefix"`
	FilterSuffix	string	`json:"filter_suffix"`
}
