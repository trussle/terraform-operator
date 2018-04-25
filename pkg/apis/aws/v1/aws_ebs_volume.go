
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEbsVolume describes a AwsEbsVolume resource
type AwsEbsVolume struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEbsVolumeSpec	`json:"spec"`
}


// AwsEbsVolumeSpec is the spec for a AwsEbsVolume Resource
type AwsEbsVolumeSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	AvailabilityZone	string	`json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEbsVolumeList is a list of AwsEbsVolume resources
type AwsEbsVolumeList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEbsVolume	`json:"items"`
}

