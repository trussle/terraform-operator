
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSubnet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultSubnetSpec	`json:"spec"`
}

type AwsDefaultSubnetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultSubnet	`json:"items"`
}

type AwsDefaultSubnetSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	AvailabilityZone	string	`json:"availability_zone"`
}
