
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Tags	map[string]interface{}	`json:"tags"`
	Password	string	`json:"password"`
	SecurityGroupNames	string	`json:"security_group_names"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	MonitoringInterval	int	`json:"monitoring_interval"`
	CopyTagsToSnapshot	bool	`json:"copy_tags_to_snapshot"`
	EnabledCloudwatchLogsExports	[]interface{}	`json:"enabled_cloudwatch_logs_exports"`
	ReplicateSourceDb	string	`json:"replicate_source_db"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	AllowMajorVersionUpgrade	bool	`json:"allow_major_version_upgrade"`
	InstanceClass	string	`json:"instance_class"`
	Iops	int	`json:"iops"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbInstanceList is a list of AwsDbInstance resources
type AwsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbInstance	`json:"items"`
}

