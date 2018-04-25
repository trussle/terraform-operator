
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3Bucket describes a AwsS3Bucket resource
type AwsS3Bucket struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketSpec	`json:"spec"`
}


// AwsS3BucketSpec is the spec for a AwsS3Bucket Resource
type AwsS3BucketSpec struct {
	Policy	string	`json:"policy"`
	Website	[]interface{}	`json:"website"`
	ReplicationConfiguration	[]interface{}	`json:"replication_configuration"`
	Acl	string	`json:"acl"`
	ForceDestroy	bool	`json:"force_destroy"`
	ServerSideEncryptionConfiguration	[]interface{}	`json:"server_side_encryption_configuration"`
	Tags	map[string]interface{}	`json:"tags"`
	BucketPrefix	string	`json:"bucket_prefix"`
	CorsRule	[]interface{}	`json:"cors_rule"`
	LifecycleRule	[]interface{}	`json:"lifecycle_rule"`
	Logging	string	`json:"logging"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketList is a list of AwsS3Bucket resources
type AwsS3BucketList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3Bucket	`json:"items"`
}

