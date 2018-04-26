
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroup describes a AwsSecurityGroup resource
type AwsSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSecurityGroupSpec	`json:"spec"`
}


// AwsSecurityGroupSpec is the spec for a AwsSecurityGroup Resource
type AwsSecurityGroupSpec struct {
	Description	string	`json:"description"`
	Tags	map[string]Generic	`json:"tags"`
	RevokeRulesOnDelete	bool	`json:"revoke_rules_on_delete"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityGroupList is a list of AwsSecurityGroup resources
type AwsSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecurityGroup	`json:"items"`
}

