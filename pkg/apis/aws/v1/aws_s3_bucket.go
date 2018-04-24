
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3Bucket struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketSpec	`json:"spec"`
}

type AwsS3BucketList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3Bucket	`json:"items"`
}

type AwsS3BucketSpec struct {
	ServerSideEncryptionConfiguration	[]interface{}	`json:"server_side_encryption_configuration"`
	ForceDestroy	bool	`json:"force_destroy"`
	LifecycleRule	[]interface{}	`json:"lifecycle_rule"`
	ReplicationConfiguration	[]interface{}	`json:"replication_configuration"`
	Tags	map[string]interface{}	`json:"tags"`
	Website	[]interface{}	`json:"website"`
	Logging	interface{}	`json:"logging"`
	Acl	string	`json:"acl"`
	BucketPrefix	string	`json:"bucket_prefix"`
	Policy	string	`json:"policy"`
	CorsRule	[]interface{}	`json:"cors_rule"`
}
