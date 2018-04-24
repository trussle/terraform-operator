
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheReplicationGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheReplicationGroupSpec	`json:"spec"`
}

type AwsElasticacheReplicationGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheReplicationGroup	`json:"items"`
}

type AwsElasticacheReplicationGroupSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	AtRestEncryptionEnabled	bool	`json:"at_rest_encryption_enabled"`
	ReplicationGroupId	string	`json:"replication_group_id"`
	TransitEncryptionEnabled	bool	`json:"transit_encryption_enabled"`
	Engine	string	`json:"engine"`
	SnapshotRetentionLimit	int	`json:"snapshot_retention_limit"`
	SnapshotArns	interface{}	`json:"snapshot_arns"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	AutomaticFailoverEnabled	bool	`json:"automatic_failover_enabled"`
	AuthToken	string	`json:"auth_token"`
	SnapshotName	string	`json:"snapshot_name"`
	Port	int	`json:"port"`
	AvailabilityZones	interface{}	`json:"availability_zones"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	ReplicationGroupDescription	string	`json:"replication_group_description"`
}
