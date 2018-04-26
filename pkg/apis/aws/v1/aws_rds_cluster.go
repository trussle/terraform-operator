
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	BackupRetentionPeriod	int	`json:"backup_retention_period"`
	SourceRegion	string	`json:"source_region"`
	IamRoles	string	`json:"iam_roles"`
	Engine	string	`json:"engine"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	StorageEncrypted	bool	`json:"storage_encrypted"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	Tags	map[string]???	`json:"tags"`
	ReplicationSourceIdentifier	string	`json:"replication_source_identifier"`
	IamDatabaseAuthenticationEnabled	bool	`json:"iam_database_authentication_enabled"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	MasterPassword	string	`json:"master_password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterList is a list of AwsRdsCluster resources
type AwsRdsClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsCluster	`json:"items"`
}

