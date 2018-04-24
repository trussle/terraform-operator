
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafregionalWebAclAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalWebAclAssociationSpec	`json:"spec"`
}

type AwsWafregionalWebAclAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalWebAclAssociation	`json:"items"`
}

type AwsWafregionalWebAclAssociationSpec struct {
	WebAclId	string	`json:"web_acl_id"`
	ResourceArn	string	`json:"resource_arn"`
}
