
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEfsMountTarget describes a AwsEfsMountTarget resource
type AwsEfsMountTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEfsMountTargetSpec	`json:"spec"`
}


// AwsEfsMountTargetSpec is the spec for a AwsEfsMountTarget Resource
type AwsEfsMountTargetSpec struct {
	SubnetId	string	`json:"subnet_id"`
	FileSystemId	string	`json:"file_system_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEfsMountTargetList is a list of AwsEfsMountTarget resources
type AwsEfsMountTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEfsMountTarget	`json:"items"`
}

