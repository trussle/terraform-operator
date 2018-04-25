
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftCluster describes a AwsRedshiftCluster resource
type AwsRedshiftCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftClusterSpec	`json:"spec"`
}


// AwsRedshiftClusterSpec is the spec for a AwsRedshiftCluster Resource
type AwsRedshiftClusterSpec struct {
	MasterPassword	string	`json:"master_password"`
	NumberOfNodes	int	`json:"number_of_nodes"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	NodeType	string	`json:"node_type"`
	OwnerAccount	string	`json:"owner_account"`
	SnapshotClusterIdentifier	string	`json:"snapshot_cluster_identifier"`
	Tags	map[string]interface{}	`json:"tags"`
	AutomatedSnapshotRetentionPeriod	int	`json:"automated_snapshot_retention_period"`
	AllowVersionUpgrade	bool	`json:"allow_version_upgrade"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	Port	int	`json:"port"`
	SnapshotCopy	[]interface{}	`json:"snapshot_copy"`
	MasterUsername	string	`json:"master_username"`
	ClusterVersion	string	`json:"cluster_version"`
	ElasticIp	string	`json:"elastic_ip"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	Logging	[]interface{}	`json:"logging"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftClusterList is a list of AwsRedshiftCluster resources
type AwsRedshiftClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftCluster	`json:"items"`
}

