
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSubnet describes a AwsDefaultSubnet resource
type AwsDefaultSubnet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultSubnetSpec	`json:"spec"`
}


// AwsDefaultSubnetSpec is the spec for a AwsDefaultSubnet Resource
type AwsDefaultSubnetSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	AvailabilityZone	string	`json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSubnetList is a list of AwsDefaultSubnet resources
type AwsDefaultSubnetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultSubnet	`json:"items"`
}

