
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyIpset describes a AwsGuarddutyIpset resource
type AwsGuarddutyIpset struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGuarddutyIpsetSpec	`json:"spec"`
}


// AwsGuarddutyIpsetSpec is the spec for a AwsGuarddutyIpset Resource
type AwsGuarddutyIpsetSpec struct {
	Activate	bool	`json:"activate"`
	DetectorId	string	`json:"detector_id"`
	Name	string	`json:"name"`
	Format	string	`json:"format"`
	Location	string	`json:"location"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyIpsetList is a list of AwsGuarddutyIpset resources
type AwsGuarddutyIpsetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGuarddutyIpset	`json:"items"`
}

