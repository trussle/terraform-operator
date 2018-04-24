
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsSnapshot struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEbsSnapshotSpec	`json:"spec"`
}

type AwsEbsSnapshotList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEbsSnapshot	`json:"items"`
}

type AwsEbsSnapshotSpec struct {
	Description	string	`json:"description"`
	Tags	map[string]interface{}	`json:"tags"`
	VolumeId	string	`json:"volume_id"`
}
