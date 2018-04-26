
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEbsSnapshot describes a AwsEbsSnapshot resource
type AwsEbsSnapshot struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEbsSnapshotSpec	`json:"spec"`
}


// AwsEbsSnapshotSpec is the spec for a AwsEbsSnapshot Resource
type AwsEbsSnapshotSpec struct {
	Tags	map[string]???	`json:"tags"`
	VolumeId	string	`json:"volume_id"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEbsSnapshotList is a list of AwsEbsSnapshot resources
type AwsEbsSnapshotList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEbsSnapshot	`json:"items"`
}

