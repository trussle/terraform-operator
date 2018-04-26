
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	NamePrefix	string	`json:"name_prefix"`
	Configuration	string	`json:"configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrSecurityConfigurationList is a list of AwsEmrSecurityConfiguration resources
type AwsEmrSecurityConfigurationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrSecurityConfiguration	`json:"items"`
}

