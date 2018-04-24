
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDaxCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDaxClusterSpec	`json:"spec"`
}

type AwsDaxClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDaxCluster	`json:"items"`
}

type AwsDaxClusterSpec struct {
	ReplicationFactor	int	`json:"replication_factor"`
	NodeType	string	`json:"node_type"`
	AvailabilityZones	interface{}	`json:"availability_zones"`
	Tags	map[string]interface{}	`json:"tags"`
	IamRoleArn	string	`json:"iam_role_arn"`
	Description	string	`json:"description"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	ClusterName	string	`json:"cluster_name"`
}
