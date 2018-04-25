
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbSnapshot describes a AwsDbSnapshot resource
type AwsDbSnapshot struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbSnapshotSpec	`json:"spec"`
}


// AwsDbSnapshotSpec is the spec for a AwsDbSnapshot Resource
type AwsDbSnapshotSpec struct {
	DbSnapshotIdentifier	string	`json:"db_snapshot_identifier"`
	DbInstanceIdentifier	string	`json:"db_instance_identifier"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbSnapshotList is a list of AwsDbSnapshot resources
type AwsDbSnapshotList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbSnapshot	`json:"items"`
}

