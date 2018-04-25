
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsCluster describes a AwsRdsCluster resource
type AwsRdsCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRdsClusterSpec	`json:"spec"`
}


// AwsRdsClusterSpec is the spec for a AwsRdsCluster Resource
type AwsRdsClusterSpec struct {
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	MasterPassword	string	`json:"master_password"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	BackupRetentionPeriod	int	`json:"backup_retention_period"`
	SourceRegion	string	`json:"source_region"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	ReplicationSourceIdentifier	string	`json:"replication_source_identifier"`
	Engine	string	`json:"engine"`
	IamRoles	string	`json:"iam_roles"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterList is a list of AwsRdsCluster resources
type AwsRdsClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsCluster	`json:"items"`
}

