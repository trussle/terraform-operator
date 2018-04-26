
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrSecurityConfiguration describes a AwsEmrSecurityConfiguration resource
type AwsEmrSecurityConfiguration struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEmrSecurityConfigurationSpec	`json:"spec"`
}


// AwsEmrSecurityConfigurationSpec is the spec for a AwsEmrSecurityConfiguration Resource
type AwsEmrSecurityConfigurationSpec struct {
	Configuration	string	`json:"configuration"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrSecurityConfigurationList is a list of AwsEmrSecurityConfiguration resources
type AwsEmrSecurityConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrSecurityConfiguration	`json:"items"`
}

