
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnapshotCreateVolumePermission describes a AwsSnapshotCreateVolumePermission resource
type AwsSnapshotCreateVolumePermission struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSnapshotCreateVolumePermissionSpec	`json:"spec"`
}


// AwsSnapshotCreateVolumePermissionSpec is the spec for a AwsSnapshotCreateVolumePermission Resource
type AwsSnapshotCreateVolumePermissionSpec struct {
	SnapshotId	string	`json:"snapshot_id"`
	AccountId	string	`json:"account_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSnapshotCreateVolumePermissionList is a list of AwsSnapshotCreateVolumePermission resources
type AwsSnapshotCreateVolumePermissionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSnapshotCreateVolumePermission	`json:"items"`
}

