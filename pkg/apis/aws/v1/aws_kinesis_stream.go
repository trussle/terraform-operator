
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisStream describes a AwsKinesisStream resource
type AwsKinesisStream struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKinesisStreamSpec	`json:"spec"`
}


// AwsKinesisStreamSpec is the spec for a AwsKinesisStream Resource
type AwsKinesisStreamSpec struct {
	ShardCount	int	`json:"shard_count"`
	RetentionPeriod	int	`json:"retention_period"`
	ShardLevelMetrics	string	`json:"shard_level_metrics"`
	EncryptionType	string	`json:"encryption_type"`
	KmsKeyId	string	`json:"kms_key_id"`
	Tags	map[string]???	`json:"tags"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKinesisStreamList is a list of AwsKinesisStream resources
type AwsKinesisStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKinesisStream	`json:"items"`
}

