
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
	OwnerAccount	string	`json:"owner_account"`
	ElasticIp	string	`json:"elastic_ip"`
	SkipFinalSnapshot	bool	`json:"skip_final_snapshot"`
	SnapshotIdentifier	string	`json:"snapshot_identifier"`
	Tags	map[string]???	`json:"tags"`
	MasterPassword	string	`json:"master_password"`
	FinalSnapshotIdentifier	string	`json:"final_snapshot_identifier"`
	SnapshotCopy	[]OtaRLUtm	`json:"snapshot_copy"`
	AutomatedSnapshotRetentionPeriod	int	`json:"automated_snapshot_retention_period"`
	ClusterVersion	string	`json:"cluster_version"`
	NodeType	string	`json:"node_type"`
	Logging	[]YgmSVYBA	`json:"logging"`
	Port	int	`json:"port"`
	NumberOfNodes	int	`json:"number_of_nodes"`
	AllowVersionUpgrade	bool	`json:"allow_version_upgrade"`
	SnapshotClusterIdentifier	string	`json:"snapshot_cluster_identifier"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
	MasterUsername	string	`json:"master_username"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftClusterList is a list of AwsRedshiftCluster resources
type AwsRedshiftClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftCluster	`json:"items"`
}


// OtaRLUtm is a OtaRLUtm
type OtaRLUtm struct {
	DestinationRegion	string	`json:"destination_region"`
	RetentionPeriod	int	`json:"retention_period"`
	GrantName	string	`json:"grant_name"`
}

// YgmSVYBA is a YgmSVYBA
type YgmSVYBA struct {
	Enable	bool	`json:"enable"`
}
