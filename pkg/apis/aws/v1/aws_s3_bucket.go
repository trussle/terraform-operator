
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
	CorsRule	[]pwvkbeRK	`json:"cors_rule"`
	ReplicationConfiguration	[]DknQXeAJ	`json:"replication_configuration"`
	ServerSideEncryptionConfiguration	[]nkxLkNKW	`json:"server_side_encryption_configuration"`
	Tags	map[string]???	`json:"tags"`
	Acl	string	`json:"acl"`
	Website	[]MKcSiStq	`json:"website"`
	ForceDestroy	bool	`json:"force_destroy"`
	Logging	string	`json:"logging"`
	BucketPrefix	string	`json:"bucket_prefix"`
	Policy	string	`json:"policy"`
	LifecycleRule	[]DUZsoLEz	`json:"lifecycle_rule"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketList is a list of AwsS3Bucket resources
type AwsS3BucketList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3Bucket	`json:"items"`
}


// pwvkbeRK is a pwvkbeRK
type pwvkbeRK struct {
	ExposeHeaders	[]string	`json:"expose_headers"`
	MaxAgeSeconds	int	`json:"max_age_seconds"`
	AllowedHeaders	[]string	`json:"allowed_headers"`
	AllowedMethods	[]string	`json:"allowed_methods"`
	AllowedOrigins	[]string	`json:"allowed_origins"`
}

// vUdUyRSg is a vUdUyRSg
type vUdUyRSg struct {
	Enabled	bool	`json:"enabled"`
	MfaDelete	bool	`json:"mfa_delete"`
}

// DknQXeAJ is a DknQXeAJ
type DknQXeAJ struct {
	Role	string	`json:"role"`
	Rules	string	`json:"rules"`
}

// xvjTsnzI is a xvjTsnzI
type xvjTsnzI struct {
	KmsMasterKeyId	string	`json:"kms_master_key_id"`
	SseAlgorithm	string	`json:"sse_algorithm"`
}

// WUEqSfzv is a WUEqSfzv
type WUEqSfzv struct {
	ApplyServerSideEncryptionByDefault	[]xvjTsnzI	`json:"apply_server_side_encryption_by_default"`
}

// nkxLkNKW is a nkxLkNKW
type nkxLkNKW struct {
	Rule	[]WUEqSfzv	`json:"rule"`
}

// MKcSiStq is a MKcSiStq
type MKcSiStq struct {
	IndexDocument	string	`json:"index_document"`
	ErrorDocument	string	`json:"error_document"`
	RedirectAllRequestsTo	string	`json:"redirect_all_requests_to"`
	RoutingRules	string	`json:"routing_rules"`
}

// DUZsoLEz is a DUZsoLEz
type DUZsoLEz struct {
	NoncurrentVersionTransition	string	`json:"noncurrent_version_transition"`
	Expiration	string	`json:"expiration"`
	NoncurrentVersionExpiration	string	`json:"noncurrent_version_expiration"`
	Transition	string	`json:"transition"`
	Enabled	bool	`json:"enabled"`
	AbortIncompleteMultipartUploadDays	int	`json:"abort_incomplete_multipart_upload_days"`
	Prefix	string	`json:"prefix"`
	Tags	map[string]???	`json:"tags"`
}
