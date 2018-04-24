
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSecurityGroupSpec	`json:"spec"`
}

type AwsSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecurityGroup	`json:"items"`
}

type AwsSecurityGroupSpec struct {
	RevokeRulesOnDelete	bool	`json:"revoke_rules_on_delete"`
	NamePrefix	string	`json:"name_prefix"`
	Description	string	`json:"description"`
	Tags	map[string]interface{}	`json:"tags"`
}
