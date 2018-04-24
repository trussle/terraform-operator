
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketNotification struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketNotificationSpec	`json:"spec"`
}

type AwsS3BucketNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketNotification	`json:"items"`
}

type AwsS3BucketNotificationSpec struct {
	Queue	[]interface{}	`json:"queue"`
	LambdaFunction	[]interface{}	`json:"lambda_function"`
	Bucket	string	`json:"bucket"`
	Topic	[]interface{}	`json:"topic"`
}
