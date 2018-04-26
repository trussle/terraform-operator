
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailInstance describes a AwsLightsailInstance resource
type AwsLightsailInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailInstanceSpec	`json:"spec"`
}


// AwsLightsailInstanceSpec is the spec for a AwsLightsailInstance Resource
type AwsLightsailInstanceSpec struct {
	BlueprintId	string	`json:"blueprint_id"`
	UserData	string	`json:"user_data"`
	BundleId	string	`json:"bundle_id"`
	KeyPairName	string	`json:"key_pair_name"`
	Name	string	`json:"name"`
	AvailabilityZone	string	`json:"availability_zone"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailInstanceList is a list of AwsLightsailInstance resources
type AwsLightsailInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailInstance	`json:"items"`
}

