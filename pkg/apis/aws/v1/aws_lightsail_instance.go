
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLightsailInstanceSpec	`json:"spec"`
}

type AwsLightsailInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLightsailInstance	`json:"items"`
}

type AwsLightsailInstanceSpec struct {
	BlueprintId	string	`json:"blueprint_id"`
	KeyPairName	string	`json:"key_pair_name"`
	UserData	string	`json:"user_data"`
	AvailabilityZone	string	`json:"availability_zone"`
	BundleId	string	`json:"bundle_id"`
	Name	string	`json:"name"`
}
