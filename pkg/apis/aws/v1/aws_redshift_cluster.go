
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
	AutomatedSnapshotRetentionPeriod	int	`json:"automated_snapshot_retention_period"`
	NodeType	string	`json:"node_type"`
	MasterPassword	string	`json:"master_password"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	Logging	[]Generic	`json:"logging"`
	Tags	map[string]Generic	`json:"tags"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	ElasticIp	string	`json:"elastic_ip"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	MasterUsername	string	`json:"master_username"`
	ClusterVersion	string	`json:"cluster_version"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
	NumberOfNodes	int	`json:"number_of_nodes"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	SnapshotCopy	[]Generic	`json:"snapshot_copy"`
	SnapshotClusterIdentifier	string	`json:"snapshot_cluster_identifier"`
	Port	int	`json:"port"`
	OwnerAccount	string	`json:"owner_account"`
	AllowVersionUpgrade	bool	`json:"allow_version_upgrade"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftClusterList is a list of AwsRedshiftCluster resources
type AwsRedshiftClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftCluster	`json:"items"`
}

