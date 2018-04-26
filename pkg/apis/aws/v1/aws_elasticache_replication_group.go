
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Engine	string	`json:"engine"`
	Port	int	`json:"port"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	AtRestEncryptionEnabled	bool	`json:"at_rest_encryption_enabled"`
	SnapshotArns	Generic	`json:"snapshot_arns"`
	SnapshotRetentionLimit	int	`json:"snapshot_retention_limit"`
	ReplicationGroupDescription	string	`json:"replication_group_description"`
	SnapshotName	string	`json:"snapshot_name"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	Tags	map[string]Generic	`json:"tags"`
	AutomaticFailoverEnabled	bool	`json:"automatic_failover_enabled"`
	AuthToken	string	`json:"auth_token"`
	ReplicationGroupId	string	`json:"replication_group_id"`
	TransitEncryptionEnabled	bool	`json:"transit_encryption_enabled"`
	AvailabilityZones	Generic	`json:"availability_zones"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheReplicationGroupList is a list of AwsElasticacheReplicationGroup resources
type AwsElasticacheReplicationGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheReplicationGroup	`json:"items"`
}

