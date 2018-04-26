
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
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	MonitoringInterval	int	`json:"monitoring_interval"`
	Password	string	`json:"password"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	EnabledCloudwatchLogsExports	[]string	`json:"enabled_cloudwatch_logs_exports"`
	S3Import	[]S3Import	`json:"s3_import"`
	SecurityGroupNames	string	`json:"security_group_names"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	Iops	int	`json:"iops"`
	InstanceClass	string	`json:"instance_class"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	CopyTagsToSnapshot	bool	`json:"copy_tags_to_snapshot"`
	ReplicateSourceDb	string	`json:"replicate_source_db"`
	AllowMajorVersionUpgrade	bool	`json:"allow_major_version_upgrade"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	Tags	map[string]string	`json:"tags"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbInstanceList is a list of AwsDbInstance resources
type AwsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbInstance	`json:"items"`
}


// S3Import is a S3Import
type S3Import struct {
	BucketPrefix	string	`json:"bucket_prefix"`
	IngestionRole	string	`json:"ingestion_role"`
	SourceEngine	string	`json:"source_engine"`
	SourceEngineVersion	string	`json:"source_engine_version"`
	BucketName	string	`json:"bucket_name"`
}
