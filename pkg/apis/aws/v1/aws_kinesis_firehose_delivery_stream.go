
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisFirehoseDeliveryStream describes a AwsKinesisFirehoseDeliveryStream resource
type AwsKinesisFirehoseDeliveryStream struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKinesisFirehoseDeliveryStreamSpec	`json:"spec"`
}


// AwsKinesisFirehoseDeliveryStreamSpec is the spec for a AwsKinesisFirehoseDeliveryStream Resource
type AwsKinesisFirehoseDeliveryStreamSpec struct {
	Name	string	`json:"name"`
	S3Configuration	[]Generic	`json:"s3_configuration"`
	ElasticsearchConfiguration	[]Generic	`json:"elasticsearch_configuration"`
	SplunkConfiguration	[]Generic	`json:"splunk_configuration"`
	KinesisSourceConfiguration	[]Generic	`json:"kinesis_source_configuration"`
	Destination	string	`json:"destination"`
	ExtendedS3Configuration	[]Generic	`json:"extended_s3_configuration"`
	RedshiftConfiguration	[]Generic	`json:"redshift_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisFirehoseDeliveryStreamList is a list of AwsKinesisFirehoseDeliveryStream resources
type AwsKinesisFirehoseDeliveryStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisFirehoseDeliveryStream	`json:"items"`
}

