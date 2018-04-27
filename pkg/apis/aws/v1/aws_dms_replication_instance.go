
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
	ApplyImmediately	bool	`json:"apply_immediately"`
	ReplicationInstanceId	string	`json:"replication_instance_id"`
	ReplicationInstanceClass	string	`json:"replication_instance_class"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationInstanceList is a list of AwsDmsReplicationInstance resources
type AwsDmsReplicationInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationInstance	`json:"items"`
}

