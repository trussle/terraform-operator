
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbInstanceSpec	`json:"spec"`
}

type AwsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbInstance	`json:"items"`
}

type AwsDbInstanceSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	InstanceClass	string	`json:"instance_class"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	CopyTagsToSnapshot	bool	`json:"copy_tags_to_snapshot"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	EnabledCloudwatchLogsExports	[]interface{}	`json:"enabled_cloudwatch_logs_exports"`
	Password	string	`json:"password"`
	ReplicateSourceDb	string	`json:"replicate_source_db"`
	MonitoringInterval	int	`json:"monitoring_interval"`
	Iops	int	`json:"iops"`
	SecurityGroupNames	interface{}	`json:"security_group_names"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	AllowMajorVersionUpgrade	bool	`json:"allow_major_version_upgrade"`
}
