
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyThreatintelset describes a AwsGuarddutyThreatintelset resource
type AwsGuarddutyThreatintelset struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGuarddutyThreatintelsetSpec	`json:"spec"`
}


// AwsGuarddutyThreatintelsetSpec is the spec for a AwsGuarddutyThreatintelset Resource
type AwsGuarddutyThreatintelsetSpec struct {
	DetectorId	string	`json:"detector_id"`
	Name	string	`json:"name"`
	Format	string	`json:"format"`
	Location	string	`json:"location"`
	Activate	bool	`json:"activate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyThreatintelsetList is a list of AwsGuarddutyThreatintelset resources
type AwsGuarddutyThreatintelsetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGuarddutyThreatintelset	`json:"items"`
}

