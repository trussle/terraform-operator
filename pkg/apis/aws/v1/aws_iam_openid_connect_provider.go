
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamOpenidConnectProvider describes a AwsIamOpenidConnectProvider resource
type AwsIamOpenidConnectProvider struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamOpenidConnectProviderSpec	`json:"spec"`
}


// AwsIamOpenidConnectProviderSpec is the spec for a AwsIamOpenidConnectProvider Resource
type AwsIamOpenidConnectProviderSpec struct {
	Url	string	`json:"url"`
	ClientIdList	[]string	`json:"client_id_list"`
	ThumbprintList	[]string	`json:"thumbprint_list"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamOpenidConnectProviderList is a list of AwsIamOpenidConnectProvider resources
type AwsIamOpenidConnectProviderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamOpenidConnectProvider	`json:"items"`
}

