
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyMember describes a AwsGuarddutyMember resource
type AwsGuarddutyMember struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGuarddutyMemberSpec	`json:"spec"`
}


// AwsGuarddutyMemberSpec is the spec for a AwsGuarddutyMember Resource
type AwsGuarddutyMemberSpec struct {
	AccountId	string	`json:"account_id"`
	DetectorId	string	`json:"detector_id"`
	Email	string	`json:"email"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyMemberList is a list of AwsGuarddutyMember resources
type AwsGuarddutyMemberList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGuarddutyMember	`json:"items"`
}

