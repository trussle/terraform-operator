
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiFromInstance describes a AwsAmiFromInstance resource
type AwsAmiFromInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAmiFromInstanceSpec	`json:"spec"`
}


// AwsAmiFromInstanceSpec is the spec for a AwsAmiFromInstance Resource
type AwsAmiFromInstanceSpec struct {
	Tags	map[string]string	`json:"tags"`
	SourceInstanceId	string	`json:"source_instance_id"`
	SnapshotWithoutReboot	bool	`json:"snapshot_without_reboot"`
	Description	string	`json:"description"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiFromInstanceList is a list of AwsAmiFromInstance resources
type AwsAmiFromInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiFromInstance	`json:"items"`
}


// EbsBlockDevice is a EbsBlockDevice
type EbsBlockDevice struct {
}

// EphemeralBlockDevice is a EphemeralBlockDevice
type EphemeralBlockDevice struct {
}
