
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	ReplicationInstanceClass	string	`json:"replication_instance_class"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	Tags	map[string]string	`json:"tags"`
	ReplicationInstanceId	string	`json:"replication_instance_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationInstanceList is a list of AwsDmsReplicationInstance resources
type AwsDmsReplicationInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationInstance	`json:"items"`
}

