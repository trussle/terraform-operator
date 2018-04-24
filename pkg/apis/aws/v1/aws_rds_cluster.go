
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRdsClusterSpec	`json:"spec"`
}

type AwsRdsClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsCluster	`json:"items"`
}

type AwsRdsClusterSpec struct {
	SourceRegion	string	`json:"source_region"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	Engine	string	`json:"engine"`
	ReplicationSourceIdentifier	string	`json:"replication_source_identifier"`
	Tags	map[string]interface{}	`json:"tags"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	MasterPassword	string	`json:"master_password"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	BackupRetentionPeriod	int	`json:"backup_retention_period"`
	IamRoles	interface{}	`json:"iam_roles"`
}
