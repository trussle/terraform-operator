
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	ClusterName	string	`json:"cluster_name"`
	Tags	map[string]string	`json:"tags"`
	IamRoleArn	string	`json:"iam_role_arn"`
	ReplicationFactor	int	`json:"replication_factor"`
	Description	string	`json:"description"`
	NodeType	string	`json:"node_type"`
	AvailabilityZones	string	`json:"availability_zones"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxClusterList is a list of AwsDaxCluster resources
type AwsDaxClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDaxCluster	`json:"items"`
}


// AwsDaxClusterNodes is a AwsDaxClusterNodes
type AwsDaxClusterNodes struct {
}
