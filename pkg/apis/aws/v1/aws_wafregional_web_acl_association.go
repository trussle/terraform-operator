
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalWebAclAssociation describes a AwsWafregionalWebAclAssociation resource
type AwsWafregionalWebAclAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalWebAclAssociationSpec	`json:"spec"`
}


// AwsWafregionalWebAclAssociationSpec is the spec for a AwsWafregionalWebAclAssociation Resource
type AwsWafregionalWebAclAssociationSpec struct {
	ResourceArn	string	`json:"resource_arn"`
	WebAclId	string	`json:"web_acl_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalWebAclAssociationList is a list of AwsWafregionalWebAclAssociation resources
type AwsWafregionalWebAclAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalWebAclAssociation	`json:"items"`
}

