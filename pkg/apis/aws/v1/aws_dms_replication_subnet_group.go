
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsReplicationSubnetGroupSpec	`json:"spec"`
}

type AwsDmsReplicationSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationSubnetGroup	`json:"items"`
}

type AwsDmsReplicationSubnetGroupSpec struct {
	ReplicationSubnetGroupDescription	string	`json:"replication_subnet_group_description"`
	ReplicationSubnetGroupId	string	`json:"replication_subnet_group_id"`
	SubnetIds	interface{}	`json:"subnet_ids"`
	Tags	map[string]interface{}	`json:"tags"`
}
