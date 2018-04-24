
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheClusterSpec	`json:"spec"`
}

type AwsElasticacheClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheCluster	`json:"items"`
}

type AwsElasticacheClusterSpec struct {
	AvailabilityZones	interface{}	`json:"availability_zones"`
	SnapshotArns	interface{}	`json:"snapshot_arns"`
	ClusterId	string	`json:"cluster_id"`
	Port	int	`json:"port"`
	SnapshotRetentionLimit	int	`json:"snapshot_retention_limit"`
	Tags	map[string]interface{}	`json:"tags"`
	SnapshotName	string	`json:"snapshot_name"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
}
