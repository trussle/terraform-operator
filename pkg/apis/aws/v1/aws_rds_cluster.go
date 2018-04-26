
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
	StorageEncrypted	bool	`json:"storage_encrypted"`
	BackupRetentionPeriod	int	`json:"backup_retention_period"`
	Tags	map[string]Generic	`json:"tags"`
	ReplicationSourceIdentifier	string	`json:"replication_source_identifier"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	Engine	string	`json:"engine"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	IamRoles	Generic	`json:"iam_roles"`
	SourceRegion	string	`json:"source_region"`
	MasterPassword	string	`json:"master_password"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterList is a list of AwsRdsCluster resources
type AwsRdsClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsCluster	`json:"items"`
}

