
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	ClusterIdentifier	string	`json:"cluster_identifier"`
	MasterPassword	string	`json:"master_password"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	OwnerAccount	string	`json:"owner_account"`
	Tags	map[string]string	`json:"tags"`
	ElasticIp	string	`json:"elastic_ip"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	NodeType	string	`json:"node_type"`
	NumberOfNodes	int	`json:"number_of_nodes"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	SnapshotCopy	[]SnapshotCopy	`json:"snapshot_copy"`
	ClusterVersion	string	`json:"cluster_version"`
	Port	int	`json:"port"`
	AllowVersionUpgrade	bool	`json:"allow_version_upgrade"`
	SnapshotClusterIdentifier	string	`json:"snapshot_cluster_identifier"`
	Logging	[]Logging	`json:"logging"`
	MasterUsername	string	`json:"master_username"`
	AutomatedSnapshotRetentionPeriod	int	`json:"automated_snapshot_retention_period"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftClusterList is a list of AwsRedshiftCluster resources
type AwsRedshiftClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftCluster	`json:"items"`
}


// SnapshotCopy is a SnapshotCopy
type SnapshotCopy struct {
	DestinationRegion	string	`json:"destination_region"`
	RetentionPeriod	int	`json:"retention_period"`
	GrantName	string	`json:"grant_name"`
}

// Logging is a Logging
type Logging struct {
	Enable	bool	`json:"enable"`
}
