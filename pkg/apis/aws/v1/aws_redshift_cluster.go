
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftClusterSpec	`json:"spec"`
}

type AwsRedshiftClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftCluster	`json:"items"`
}

type AwsRedshiftClusterSpec struct {
	MasterPassword	string	`json:"master_password"`
	Port	int	`json:"port"`
	AllowVersionUpgrade	bool	`json:"allow_version_upgrade"`
	NumberOfNodes	int	`json:"number_of_nodes"`
	NodeType	string	`json:"node_type"`
	AutomatedSnapshotRetentionPeriod	int	`json:"automated_snapshot_retention_period"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	SnapshotCopy	[]interface{}	`json:"snapshot_copy"`
	MasterUsername	string	`json:"master_username"`
	ElasticIp	string	`json:"elastic_ip"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	Tags	map[string]interface{}	`json:"tags"`
	ClusterVersion	string	`json:"cluster_version"`
	OwnerAccount	string	`json:"owner_account"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
	Logging	[]interface{}	`json:"logging"`
	SnapshotClusterIdentifier	string	`json:"snapshot_cluster_identifier"`
}
