
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
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	ReplicationGroupDescription	string	`json:"replication_group_description"`
	SnapshotName	string	`json:"snapshot_name"`
	SnapshotRetentionLimit	int	`json:"snapshot_retention_limit"`
	Tags	map[string]???	`json:"tags"`
	ReplicationGroupId	string	`json:"replication_group_id"`
	AtRestEncryptionEnabled	bool	`json:"at_rest_encryption_enabled"`
	Port	int	`json:"port"`
	AutomaticFailoverEnabled	bool	`json:"automatic_failover_enabled"`
	AuthToken	string	`json:"auth_token"`
	TransitEncryptionEnabled	bool	`json:"transit_encryption_enabled"`
	AvailabilityZones	string	`json:"availability_zones"`
	Engine	string	`json:"engine"`
	SnapshotArns	string	`json:"snapshot_arns"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheReplicationGroupList is a list of AwsElasticacheReplicationGroup resources
type AwsElasticacheReplicationGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheReplicationGroup	`json:"items"`
}


// oVVTYNXf is a oVVTYNXf
type oVVTYNXf struct {
	ReplicasPerNodeGroup	int	`json:"replicas_per_node_group"`
	NumNodeGroups	int	`json:"num_node_groups"`
}
