
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamSamlProvider struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamSamlProviderSpec	`json:"spec"`
}

type AwsIamSamlProviderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamSamlProvider	`json:"items"`
}

type AwsIamSamlProviderSpec struct {
	Name	string	`json:"name"`
	SamlMetadataDocument	string	`json:"saml_metadata_document"`
}