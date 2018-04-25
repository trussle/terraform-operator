
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamSamlProvider describes a AwsIamSamlProvider resource
type AwsIamSamlProvider struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamSamlProviderSpec	`json:"spec"`
}


// AwsIamSamlProviderSpec is the spec for a AwsIamSamlProvider Resource
type AwsIamSamlProviderSpec struct {
	Name	string	`json:"name"`
	SamlMetadataDocument	string	`json:"saml_metadata_document"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamSamlProviderList is a list of AwsIamSamlProvider resources
type AwsIamSamlProviderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamSamlProvider	`json:"items"`
}

