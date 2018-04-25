
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsOrganization describes a AwsOrganizationsOrganization resource
type AwsOrganizationsOrganization struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOrganizationsOrganizationSpec	`json:"spec"`
}


// AwsOrganizationsOrganizationSpec is the spec for a AwsOrganizationsOrganization Resource
type AwsOrganizationsOrganizationSpec struct {
	FeatureSet	string	`json:"feature_set"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOrganizationsOrganizationList is a list of AwsOrganizationsOrganization resources
type AwsOrganizationsOrganizationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOrganizationsOrganization	`json:"items"`
}

