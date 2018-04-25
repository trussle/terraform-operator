
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
	KinesisSourceConfiguration	[]interface{}	`json:"kinesis_source_configuration"`
	S3Configuration	[]interface{}	`json:"s3_configuration"`
	ExtendedS3Configuration	[]interface{}	`json:"extended_s3_configuration"`
	Name	string	`json:"name"`
	Destination	string	`json:"destination"`
	RedshiftConfiguration	[]interface{}	`json:"redshift_configuration"`
	ElasticsearchConfiguration	[]interface{}	`json:"elasticsearch_configuration"`
	SplunkConfiguration	[]interface{}	`json:"splunk_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisFirehoseDeliveryStreamList is a list of AwsKinesisFirehoseDeliveryStream resources
type AwsKinesisFirehoseDeliveryStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisFirehoseDeliveryStream	`json:"items"`
}

