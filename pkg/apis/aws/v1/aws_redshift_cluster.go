
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
	SnapshotCopy	[]AwsRedshiftClusterSnapshotCopy	`json:"snapshot_copy"`
	Logging	[]AwsRedshiftClusterLogging	`json:"logging"`
	MasterPassword	string	`json:"master_password"`
	Tags	map[string]string	`json:"tags"`
	MasterUsername	string	`json:"master_username"`
	Port	int	`json:"port"`
	SnapshotClusterIdentifier	string	`json:"snapshot_cluster_identifier"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	ClusterVersion	string	`json:"cluster_version"`
	AllowVersionUpgrade	bool	`json:"allow_version_upgrade"`
	ElasticIp	string	`json:"elastic_ip"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
	AutomatedSnapshotRetentionPeriod	int	`json:"automated_snapshot_retention_period"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	OwnerAccount	string	`json:"owner_account"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	NodeType	string	`json:"node_type"`
	NumberOfNodes	int	`json:"number_of_nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftClusterList is a list of AwsRedshiftCluster resources
type AwsRedshiftClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftCluster	`json:"items"`
}


// AwsRedshiftClusterSnapshotCopy is a AwsRedshiftClusterSnapshotCopy
type AwsRedshiftClusterSnapshotCopy struct {
	DestinationRegion	string	`json:"destination_region"`
	RetentionPeriod	int	`json:"retention_period"`
	GrantName	string	`json:"grant_name"`
}

// AwsRedshiftClusterLogging is a AwsRedshiftClusterLogging
type AwsRedshiftClusterLogging struct {
	Enable	bool	`json:"enable"`
}
