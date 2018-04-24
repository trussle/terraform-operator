
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamOpenidConnectProvider struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamOpenidConnectProviderSpec	`json:"spec"`
}

type AwsIamOpenidConnectProviderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamOpenidConnectProvider	`json:"items"`
}

type AwsIamOpenidConnectProviderSpec struct {
	ThumbprintList	[]interface{}	`json:"thumbprint_list"`
	Url	string	`json:"url"`
	ClientIdList	[]interface{}	`json:"client_id_list"`
}
