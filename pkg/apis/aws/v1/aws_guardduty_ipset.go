
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyIpset struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGuarddutyIpsetSpec	`json:"spec"`
}

type AwsGuarddutyIpsetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGuarddutyIpset	`json:"items"`
}

type AwsGuarddutyIpsetSpec struct {
	DetectorId	string	`json:"detector_id"`
	Name	string	`json:"name"`
	Format	string	`json:"format"`
	Location	string	`json:"location"`
	Activate	bool	`json:"activate"`
}
