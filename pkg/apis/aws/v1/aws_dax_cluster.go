
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
	ReplicationFactor	int	`json:"replication_factor"`
	Description	string	`json:"description"`
	NotificationTopicArn	string	`json:"notification_topic_arn"`
	Tags	map[string]string	`json:"tags"`
	ClusterName	string	`json:"cluster_name"`
	NodeType	string	`json:"node_type"`
	AvailabilityZones	string	`json:"availability_zones"`
	IamRoleArn	string	`json:"iam_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDaxClusterList is a list of AwsDaxCluster resources
type AwsDaxClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDaxCluster	`json:"items"`
}


// Nodes is a Nodes
type Nodes struct {
}
