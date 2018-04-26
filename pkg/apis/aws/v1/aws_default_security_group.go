
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSecurityGroup describes a AwsDefaultSecurityGroup resource
type AwsDefaultSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultSecurityGroupSpec	`json:"spec"`
}


// AwsDefaultSecurityGroupSpec is the spec for a AwsDefaultSecurityGroup Resource
type AwsDefaultSecurityGroupSpec struct {
	Ingress	string	`json:"ingress"`
	RevokeRulesOnDelete	bool	`json:"revoke_rules_on_delete"`
	Egress	string	`json:"egress"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDefaultSecurityGroupList is a list of AwsDefaultSecurityGroup resources
type AwsDefaultSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultSecurityGroup	`json:"items"`
}

