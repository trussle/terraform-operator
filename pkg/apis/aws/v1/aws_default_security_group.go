
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultSecurityGroupSpec	`json:"spec"`
}

type AwsDefaultSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultSecurityGroup	`json:"items"`
}

type AwsDefaultSecurityGroupSpec struct {
	Ingress	interface{}	`json:"ingress"`
	RevokeRulesOnDelete	bool	`json:"revoke_rules_on_delete"`
	Tags	map[string]interface{}	`json:"tags"`
	Egress	interface{}	`json:"egress"`
}
