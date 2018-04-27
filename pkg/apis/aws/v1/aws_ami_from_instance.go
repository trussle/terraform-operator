
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
	Description	string	`json:"description"`
	SnapshotWithoutReboot	bool	`json:"snapshot_without_reboot"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiFromInstanceList is a list of AwsAmiFromInstance resources
type AwsAmiFromInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiFromInstance	`json:"items"`
}


// AwsAmiFromInstanceEbsBlockDevice is a AwsAmiFromInstanceEbsBlockDevice
type AwsAmiFromInstanceEbsBlockDevice struct {
}

// AwsAmiFromInstanceEphemeralBlockDevice is a AwsAmiFromInstanceEphemeralBlockDevice
type AwsAmiFromInstanceEphemeralBlockDevice struct {
}
