
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxCluster describes a AwsDaxCluster resource
type AwsDaxCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDaxClusterSpec	`json:"spec"`
}


// AwsDaxClusterSpec is the spec for a AwsDaxCluster Resource
type AwsDaxClusterSpec struct {
	ReplicationFactor	int	`json:"replication_factor"`
	IamRoleArn	string	`json:"iam_role_arn"`
	Tags	map[string]Generic	`json:"tags"`
	ClusterName	string	`json:"cluster_name"`
	Description	string	`json:"description"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	NodeType	string	`json:"node_type"`
	AvailabilityZones	Generic	`json:"availability_zones"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxClusterList is a list of AwsDaxCluster resources
type AwsDaxClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDaxCluster	`json:"items"`
}

