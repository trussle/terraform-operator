
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyMember struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGuarddutyMemberSpec	`json:"spec"`
}

type AwsGuarddutyMemberList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGuarddutyMember	`json:"items"`
}

type AwsGuarddutyMemberSpec struct {
	AccountId	string	`json:"account_id"`
	DetectorId	string	`json:"detector_id"`
	Email	string	`json:"email"`
}
