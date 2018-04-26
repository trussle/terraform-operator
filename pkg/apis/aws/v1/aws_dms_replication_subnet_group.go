
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationSubnetGroup describes a AwsDmsReplicationSubnetGroup resource
type AwsDmsReplicationSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsReplicationSubnetGroupSpec	`json:"spec"`
}


// AwsDmsReplicationSubnetGroupSpec is the spec for a AwsDmsReplicationSubnetGroup Resource
type AwsDmsReplicationSubnetGroupSpec struct {
	ReplicationSubnetGroupDescription	string	`json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId	string	`json:"replication_subnet_group_id"`
	SubnetIds	string	`json:"subnet_ids"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationSubnetGroupList is a list of AwsDmsReplicationSubnetGroup resources
type AwsDmsReplicationSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationSubnetGroup	`json:"items"`
}

