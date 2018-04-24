
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsReplicationInstanceSpec	`json:"spec"`
}

type AwsDmsReplicationInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationInstance	`json:"items"`
}

type AwsDmsReplicationInstanceSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	ReplicationInstanceId	string	`json:"replication_instance_id"`
	ReplicationInstanceClass	string	`json:"replication_instance_class"`
}
