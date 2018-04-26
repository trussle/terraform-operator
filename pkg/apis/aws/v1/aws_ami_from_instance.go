
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Description	string	`json:"description"`
	SnapshotWithoutReboot	bool	`json:"snapshot_without_reboot"`
	Name	string	`json:"name"`
	SourceInstanceId	string	`json:"source_instance_id"`
	Tags	map[string]Generic	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiFromInstanceList is a list of AwsAmiFromInstance resources
type AwsAmiFromInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiFromInstance	`json:"items"`
}

