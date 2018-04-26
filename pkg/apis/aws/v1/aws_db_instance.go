
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
	Iops	int	`json:"iops"`
	ReplicateSourceDb	string	`json:"replicate_source_db"`
	AllowMajorVersionUpgrade	bool	`json:"allow_major_version_upgrade"`
	EnabledCloudwatchLogsExports	[]Generic	`json:"enabled_cloudwatch_logs_exports"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	SecurityGroupNames	Generic	`json:"security_group_names"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	Tags	map[string]Generic	`json:"tags"`
	InstanceClass	string	`json:"instance_class"`
	CopyTagsToSnapshot	bool	`json:"copy_tags_to_snapshot"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	Password	string	`json:"password"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	MonitoringInterval	int	`json:"monitoring_interval"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbInstanceList is a list of AwsDbInstance resources
type AwsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbInstance	`json:"items"`
}

