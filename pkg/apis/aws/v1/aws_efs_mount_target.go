
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEfsMountTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEfsMountTargetSpec	`json:"spec"`
}

type AwsEfsMountTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEfsMountTarget	`json:"items"`
}

type AwsEfsMountTargetSpec struct {
	FileSystemId	string	`json:"file_system_id"`
	SubnetId	string	`json:"subnet_id"`
}
