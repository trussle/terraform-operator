
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheReplicationGroup describes a AwsElasticacheReplicationGroup resource
type AwsElasticacheReplicationGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheReplicationGroupSpec	`json:"spec"`
}


// AwsElasticacheReplicationGroupSpec is the spec for a AwsElasticacheReplicationGroup Resource
type AwsElasticacheReplicationGroupSpec struct {
	TransitEncryptionEnabled	bool	`json:"transit_encryption_enabled"`
	Engine	string	`json:"engine"`
	SnapshotName	string	`json:"snapshot_name"`
	Tags	map[string]string	`json:"tags"`
	SnapshotRetentionLimit	int	`json:"snapshot_retention_limit"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	AtRestEncryptionEnabled	bool	`json:"at_rest_encryption_enabled"`
	AvailabilityZones	string	`json:"availability_zones"`
	AuthToken	string	`json:"auth_token"`
	Port	int	`json:"port"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	ReplicationGroupId	string	`json:"replication_group_id"`
	ReplicationGroupDescription	string	`json:"replication_group_description"`
	SnapshotArns	string	`json:"snapshot_arns"`
	AutomaticFailoverEnabled	bool	`json:"automatic_failover_enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheReplicationGroupList is a list of AwsElasticacheReplicationGroup resources
type AwsElasticacheReplicationGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheReplicationGroup	`json:"items"`
}


// AwsElasticacheReplicationGroupClusterMode is a AwsElasticacheReplicationGroupClusterMode
type AwsElasticacheReplicationGroupClusterMode struct {
	ReplicasPerNodeGroup	int	`json:"replicas_per_node_group"`
	NumNodeGroups	int	`json:"num_node_groups"`
}
