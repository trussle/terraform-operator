
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEbsVolume struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEbsVolumeSpec	`json:"spec"`
}

type AwsEbsVolumeList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEbsVolume	`json:"items"`
}

type AwsEbsVolumeSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	AvailabilityZone	string	`json:"availability_zone"`
}
