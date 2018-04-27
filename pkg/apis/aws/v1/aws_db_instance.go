
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbInstance describes a AwsDbInstance resource
type AwsDbInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbInstanceSpec	`json:"spec"`
}


// AwsDbInstanceSpec is the spec for a AwsDbInstance Resource
type AwsDbInstanceSpec struct {
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	EnabledCloudwatchLogsExports	[]string	`json:"enabled_cloudwatch_logs_exports"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	InstanceClass	string	`json:"instance_class"`
	Iops	int	`json:"iops"`
	ReplicateSourceDb	string	`json:"replicate_source_db"`
	MonitoringInterval	int	`json:"monitoring_interval"`
	Password	string	`json:"password"`
	S3Import	[]AwsDbInstanceS3Import	`json:"s3_import"`
	Tags	map[string]string	`json:"tags"`
	SecurityGroupNames	string	`json:"security_group_names"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	CopyTagsToSnapshot	bool	`json:"copy_tags_to_snapshot"`
	AllowMajorVersionUpgrade	bool	`json:"allow_major_version_upgrade"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbInstanceList is a list of AwsDbInstance resources
type AwsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbInstance	`json:"items"`
}


// AwsDbInstanceS3Import is a AwsDbInstanceS3Import
type AwsDbInstanceS3Import struct {
	BucketName	string	`json:"bucket_name"`
	BucketPrefix	string	`json:"bucket_prefix"`
	IngestionRole	string	`json:"ingestion_role"`
	SourceEngine	string	`json:"source_engine"`
	SourceEngineVersion	string	`json:"source_engine_version"`
}
