
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
	ReplicateSourceDb	string	`json:"replicate_source_db"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	S3Import	[]IubeYTND	`json:"s3_import"`
	Password	string	`json:"password"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	EnabledCloudwatchLogsExports	[]string	`json:"enabled_cloudwatch_logs_exports"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	AllowMajorVersionUpgrade	bool	`json:"allow_major_version_upgrade"`
	MonitoringInterval	int	`json:"monitoring_interval"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	CopyTagsToSnapshot	bool	`json:"copy_tags_to_snapshot"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	Tags	map[string]???	`json:"tags"`
	Iops	int	`json:"iops"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	SecurityGroupNames	string	`json:"security_group_names"`
	InstanceClass	string	`json:"instance_class"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbInstanceList is a list of AwsDbInstance resources
type AwsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbInstance	`json:"items"`
}


// IubeYTND is a IubeYTND
type IubeYTND struct {
	SourceEngine	string	`json:"source_engine"`
	SourceEngineVersion	string	`json:"source_engine_version"`
	BucketName	string	`json:"bucket_name"`
	BucketPrefix	string	`json:"bucket_prefix"`
	IngestionRole	string	`json:"ingestion_role"`
}
