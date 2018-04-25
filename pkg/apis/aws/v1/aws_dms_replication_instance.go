
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationInstance describes a AwsDmsReplicationInstance resource
type AwsDmsReplicationInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsReplicationInstanceSpec	`json:"spec"`
}


// AwsDmsReplicationInstanceSpec is the spec for a AwsDmsReplicationInstance Resource
type AwsDmsReplicationInstanceSpec struct {
	ReplicationInstanceId	string	`json:"replication_instance_id"`
	Tags	map[string]interface{}	`json:"tags"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	ReplicationInstanceClass	string	`json:"replication_instance_class"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationInstanceList is a list of AwsDmsReplicationInstance resources
type AwsDmsReplicationInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationInstance	`json:"items"`
}

