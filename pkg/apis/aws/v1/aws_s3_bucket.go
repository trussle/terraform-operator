
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	LifecycleRule	[]AwsS3BucketLifecycleRule	`json:"lifecycle_rule"`
	Website	[]AwsS3BucketWebsite	`json:"website"`
	Logging	AwsS3BucketLogging	`json:"logging"`
	ForceDestroy	bool	`json:"force_destroy"`
	ReplicationConfiguration	[]AwsS3BucketReplicationConfiguration	`json:"replication_configuration"`
	BucketPrefix	string	`json:"bucket_prefix"`
	Acl	string	`json:"acl"`
	CorsRule	[]AwsS3BucketCorsRule	`json:"cors_rule"`
	ServerSideEncryptionConfiguration	[]AwsS3BucketServerSideEncryptionConfiguration	`json:"server_side_encryption_configuration"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketList is a list of AwsS3Bucket resources
type AwsS3BucketList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3Bucket	`json:"items"`
}


// AwsS3BucketSourceSelectionCriteria is a AwsS3BucketSourceSelectionCriteria
type AwsS3BucketSourceSelectionCriteria struct {
	SseKmsEncryptedObjects	AwsS3BucketSseKmsEncryptedObjects	`json:"sse_kms_encrypted_objects"`
}

// AwsS3BucketCorsRule is a AwsS3BucketCorsRule
type AwsS3BucketCorsRule struct {
	AllowedMethods	[]string	`json:"allowed_methods"`
	AllowedOrigins	[]string	`json:"allowed_origins"`
	ExposeHeaders	[]string	`json:"expose_headers"`
	MaxAgeSeconds	int	`json:"max_age_seconds"`
	AllowedHeaders	[]string	`json:"allowed_headers"`
}

// AwsS3BucketRule is a AwsS3BucketRule
type AwsS3BucketRule struct {
	ApplyServerSideEncryptionByDefault	[]AwsS3BucketApplyServerSideEncryptionByDefault	`json:"apply_server_side_encryption_by_default"`
}

// AwsS3BucketExpiration is a AwsS3BucketExpiration
type AwsS3BucketExpiration struct {
	Date	string	`json:"date"`
	Days	int	`json:"days"`
	ExpiredObjectDeleteMarker	bool	`json:"expired_object_delete_marker"`
}

// AwsS3BucketNoncurrentVersionTransition is a AwsS3BucketNoncurrentVersionTransition
type AwsS3BucketNoncurrentVersionTransition struct {
	Days	int	`json:"days"`
	StorageClass	string	`json:"storage_class"`
}

// AwsS3BucketReplicationConfiguration is a AwsS3BucketReplicationConfiguration
type AwsS3BucketReplicationConfiguration struct {
	Role	string	`json:"role"`
	Rules	AwsS3BucketRules	`json:"rules"`
}

// AwsS3BucketVersioning is a AwsS3BucketVersioning
type AwsS3BucketVersioning struct {
	Enabled	bool	`json:"enabled"`
	MfaDelete	bool	`json:"mfa_delete"`
}

// AwsS3BucketRules is a AwsS3BucketRules
type AwsS3BucketRules struct {
	Id	string	`json:"id"`
	Destination	AwsS3BucketDestination	`json:"destination"`
	SourceSelectionCriteria	AwsS3BucketSourceSelectionCriteria	`json:"source_selection_criteria"`
	Prefix	string	`json:"prefix"`
	Status	string	`json:"status"`
}

// AwsS3BucketLogging is a AwsS3BucketLogging
type AwsS3BucketLogging struct {
	TargetBucket	string	`json:"target_bucket"`
	TargetPrefix	string	`json:"target_prefix"`
}

// AwsS3BucketLifecycleRule is a AwsS3BucketLifecycleRule
type AwsS3BucketLifecycleRule struct {
	Expiration	AwsS3BucketExpiration	`json:"expiration"`
	NoncurrentVersionExpiration	AwsS3BucketNoncurrentVersionExpiration	`json:"noncurrent_version_expiration"`
	Transition	AwsS3BucketTransition	`json:"transition"`
	Tags	map[string]string	`json:"tags"`
	Enabled	bool	`json:"enabled"`
	AbortIncompleteMultipartUploadDays	int	`json:"abort_incomplete_multipart_upload_days"`
	Prefix	string	`json:"prefix"`
	NoncurrentVersionTransition	AwsS3BucketNoncurrentVersionTransition	`json:"noncurrent_version_transition"`
}

// AwsS3BucketWebsite is a AwsS3BucketWebsite
type AwsS3BucketWebsite struct {
	IndexDocument	string	`json:"index_document"`
	ErrorDocument	string	`json:"error_document"`
	RedirectAllRequestsTo	string	`json:"redirect_all_requests_to"`
	RoutingRules	string	`json:"routing_rules"`
}

// AwsS3BucketDestination is a AwsS3BucketDestination
type AwsS3BucketDestination struct {
	Bucket	string	`json:"bucket"`
	StorageClass	string	`json:"storage_class"`
	ReplicaKmsKeyId	string	`json:"replica_kms_key_id"`
}

// AwsS3BucketSseKmsEncryptedObjects is a AwsS3BucketSseKmsEncryptedObjects
type AwsS3BucketSseKmsEncryptedObjects struct {
	Enabled	bool	`json:"enabled"`
}

// AwsS3BucketApplyServerSideEncryptionByDefault is a AwsS3BucketApplyServerSideEncryptionByDefault
type AwsS3BucketApplyServerSideEncryptionByDefault struct {
	KmsMasterKeyId	string	`json:"kms_master_key_id"`
	SseAlgorithm	string	`json:"sse_algorithm"`
}

// AwsS3BucketServerSideEncryptionConfiguration is a AwsS3BucketServerSideEncryptionConfiguration
type AwsS3BucketServerSideEncryptionConfiguration struct {
	Rule	[]AwsS3BucketRule	`json:"rule"`
}

// AwsS3BucketNoncurrentVersionExpiration is a AwsS3BucketNoncurrentVersionExpiration
type AwsS3BucketNoncurrentVersionExpiration struct {
	Days	int	`json:"days"`
}

// AwsS3BucketTransition is a AwsS3BucketTransition
type AwsS3BucketTransition struct {
	Days	int	`json:"days"`
	StorageClass	string	`json:"storage_class"`
	Date	string	`json:"date"`
}
