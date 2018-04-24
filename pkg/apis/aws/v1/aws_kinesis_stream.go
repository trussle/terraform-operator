
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKinesisStream struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKinesisStreamSpec	`json:"spec"`
}

type AwsKinesisStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisStream	`json:"items"`
}

type AwsKinesisStreamSpec struct {
	RetentionPeriod	int	`json:"retention_period"`
	ShardLevelMetrics	interface{}	`json:"shard_level_metrics"`
	EncryptionType	string	`json:"encryption_type"`
	KmsKeyId	string	`json:"kms_key_id"`
	Tags	map[string]interface{}	`json:"tags"`
	Name	string	`json:"name"`
	ShardCount	int	`json:"shard_count"`
}
