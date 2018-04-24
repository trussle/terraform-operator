
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSnapshot struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbSnapshotSpec	`json:"spec"`
}

type AwsDbSnapshotList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbSnapshot	`json:"items"`
}

type AwsDbSnapshotSpec struct {
	DbInstanceIdentifier	string	`json:"db_instance_identifier"`
	DbSnapshotIdentifier	string	`json:"db_snapshot_identifier"`
}
