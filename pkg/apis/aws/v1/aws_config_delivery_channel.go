
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigDeliveryChannel describes a AwsConfigDeliveryChannel resource
type AwsConfigDeliveryChannel struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsConfigDeliveryChannelSpec	`json:"spec"`
}


// AwsConfigDeliveryChannelSpec is the spec for a AwsConfigDeliveryChannel Resource
type AwsConfigDeliveryChannelSpec struct {
	S3KeyPrefix	string	`json:"s3_key_prefix"`
	SnsTopicArn	string	`json:"sns_topic_arn"`
	SnapshotDeliveryProperties	[]Generic	`json:"snapshot_delivery_properties"`
	Name	string	`json:"name"`
	S3BucketName	string	`json:"s3_bucket_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsConfigDeliveryChannelList is a list of AwsConfigDeliveryChannel resources
type AwsConfigDeliveryChannelList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsConfigDeliveryChannel	`json:"items"`
}

