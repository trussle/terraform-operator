
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
	SnapshotArns	string	`json:"snapshot_arns"`
	Port	int	`json:"port"`
	SnapshotName	string	`json:"snapshot_name"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	ClusterId	string	`json:"cluster_id"`
	Tags	map[string]???	`json:"tags"`
	AvailabilityZones	string	`json:"availability_zones"`
	SnapshotRetentionLimit	int	`json:"snapshot_retention_limit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheClusterList is a list of AwsElasticacheCluster resources
type AwsElasticacheClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheCluster	`json:"items"`
}


// DOUnXzmD is a DOUnXzmD
type DOUnXzmD struct {
}
