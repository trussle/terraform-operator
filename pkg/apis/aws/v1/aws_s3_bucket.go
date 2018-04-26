
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
	Tags	map[string]string	`json:"tags"`
	BucketPrefix	string	`json:"bucket_prefix"`
	Acl	string	`json:"acl"`
	Website	[]Website	`json:"website"`
	ForceDestroy	bool	`json:"force_destroy"`
	CorsRule	[]CorsRule	`json:"cors_rule"`
	Logging	Logging	`json:"logging"`
	LifecycleRule	[]LifecycleRule	`json:"lifecycle_rule"`
	Policy	string	`json:"policy"`
	ReplicationConfiguration	[]ReplicationConfiguration	`json:"replication_configuration"`
	ServerSideEncryptionConfiguration	[]ServerSideEncryptionConfiguration	`json:"server_side_encryption_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketList is a list of AwsS3Bucket resources
type AwsS3BucketList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3Bucket	`json:"items"`
}


// CorsRule is a CorsRule
type CorsRule struct {
	AllowedMethods	[]string	`json:"allowed_methods"`
	AllowedOrigins	[]string	`json:"allowed_origins"`
	ExposeHeaders	[]string	`json:"expose_headers"`
	MaxAgeSeconds	int	`json:"max_age_seconds"`
	AllowedHeaders	[]string	`json:"allowed_headers"`
}

// NoncurrentVersionExpiration is a NoncurrentVersionExpiration
type NoncurrentVersionExpiration struct {
	Days	int	`json:"days"`
}

// NoncurrentVersionTransition is a NoncurrentVersionTransition
type NoncurrentVersionTransition struct {
	Days	int	`json:"days"`
	StorageClass	string	`json:"storage_class"`
}

// Website is a Website
type Website struct {
	RoutingRules	string	`json:"routing_rules"`
	IndexDocument	string	`json:"index_document"`
	ErrorDocument	string	`json:"error_document"`
	RedirectAllRequestsTo	string	`json:"redirect_all_requests_to"`
}

// Transition is a Transition
type Transition struct {
	Days	int	`json:"days"`
	StorageClass	string	`json:"storage_class"`
	Date	string	`json:"date"`
}

// SseKmsEncryptedObjects is a SseKmsEncryptedObjects
type SseKmsEncryptedObjects struct {
	Enabled	bool	`json:"enabled"`
}

// SourceSelectionCriteria is a SourceSelectionCriteria
type SourceSelectionCriteria struct {
	SseKmsEncryptedObjects	SseKmsEncryptedObjects	`json:"sse_kms_encrypted_objects"`
}

// ApplyServerSideEncryptionByDefault is a ApplyServerSideEncryptionByDefault
type ApplyServerSideEncryptionByDefault struct {
	KmsMasterKeyId	string	`json:"kms_master_key_id"`
	SseAlgorithm	string	`json:"sse_algorithm"`
}

// Versioning is a Versioning
type Versioning struct {
	Enabled	bool	`json:"enabled"`
	MfaDelete	bool	`json:"mfa_delete"`
}

// Logging is a Logging
type Logging struct {
	TargetPrefix	string	`json:"target_prefix"`
	TargetBucket	string	`json:"target_bucket"`
}

// Expiration is a Expiration
type Expiration struct {
	Date	string	`json:"date"`
	Days	int	`json:"days"`
	ExpiredObjectDeleteMarker	bool	`json:"expired_object_delete_marker"`
}

// LifecycleRule is a LifecycleRule
type LifecycleRule struct {
	Tags	map[string]string	`json:"tags"`
	AbortIncompleteMultipartUploadDays	int	`json:"abort_incomplete_multipart_upload_days"`
	Expiration	Expiration	`json:"expiration"`
	NoncurrentVersionExpiration	NoncurrentVersionExpiration	`json:"noncurrent_version_expiration"`
	Transition	Transition	`json:"transition"`
	NoncurrentVersionTransition	NoncurrentVersionTransition	`json:"noncurrent_version_transition"`
	Prefix	string	`json:"prefix"`
	Enabled	bool	`json:"enabled"`
}

// Destination is a Destination
type Destination struct {
	Bucket	string	`json:"bucket"`
	StorageClass	string	`json:"storage_class"`
	ReplicaKmsKeyId	string	`json:"replica_kms_key_id"`
}

// Rules is a Rules
type Rules struct {
	Status	string	`json:"status"`
	Id	string	`json:"id"`
	Destination	Destination	`json:"destination"`
	SourceSelectionCriteria	SourceSelectionCriteria	`json:"source_selection_criteria"`
	Prefix	string	`json:"prefix"`
}

// ReplicationConfiguration is a ReplicationConfiguration
type ReplicationConfiguration struct {
	Role	string	`json:"role"`
	Rules	Rules	`json:"rules"`
}

// Rule is a Rule
type Rule struct {
	ApplyServerSideEncryptionByDefault	[]ApplyServerSideEncryptionByDefault	`json:"apply_server_side_encryption_by_default"`
}

// ServerSideEncryptionConfiguration is a ServerSideEncryptionConfiguration
type ServerSideEncryptionConfiguration struct {
	Rule	[]Rule	`json:"rule"`
}
