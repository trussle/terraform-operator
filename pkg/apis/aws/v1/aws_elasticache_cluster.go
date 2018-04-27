
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheCluster describes a AwsElasticacheCluster resource
type AwsElasticacheCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheClusterSpec	`json:"spec"`
}


// AwsElasticacheClusterSpec is the spec for a AwsElasticacheCluster Resource
type AwsElasticacheClusterSpec struct {
	Port	int	`json:"port"`
	SnapshotArns	string	`json:"snapshot_arns"`
	Tags	map[string]string	`json:"tags"`
	SnapshotName	string	`json:"snapshot_name"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	SnapshotRetentionLimit	int	`json:"snapshot_retention_limit"`
	ClusterId	string	`json:"cluster_id"`
	AvailabilityZones	string	`json:"availability_zones"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheClusterList is a list of AwsElasticacheCluster resources
type AwsElasticacheClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheCluster	`json:"items"`
}


// AwsElasticacheClusterCacheNodes is a AwsElasticacheClusterCacheNodes
type AwsElasticacheClusterCacheNodes struct {
}
