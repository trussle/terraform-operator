
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKinesisFirehoseDeliveryStream struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKinesisFirehoseDeliveryStreamSpec	`json:"spec"`
}

type AwsKinesisFirehoseDeliveryStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisFirehoseDeliveryStream	`json:"items"`
}

type AwsKinesisFirehoseDeliveryStreamSpec struct {
	Name	string	`json:"name"`
	Destination	string	`json:"destination"`
	S3Configuration	[]interface{}	`json:"s3_configuration"`
	SplunkConfiguration	[]interface{}	`json:"splunk_configuration"`
	KinesisSourceConfiguration	[]interface{}	`json:"kinesis_source_configuration"`
	ExtendedS3Configuration	[]interface{}	`json:"extended_s3_configuration"`
	RedshiftConfiguration	[]interface{}	`json:"redshift_configuration"`
	ElasticsearchConfiguration	[]interface{}	`json:"elasticsearch_configuration"`
}
