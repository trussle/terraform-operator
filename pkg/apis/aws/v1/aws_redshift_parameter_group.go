
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftParameterGroup describes a AwsRedshiftParameterGroup resource
type AwsRedshiftParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftParameterGroupSpec	`json:"spec"`
}


// AwsRedshiftParameterGroupSpec is the spec for a AwsRedshiftParameterGroup Resource
type AwsRedshiftParameterGroupSpec struct {
	Parameter	AwsRedshiftParameterGroupParameter	`json:"parameter"`
	Name	string	`json:"name"`
	Family	string	`json:"family"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftParameterGroupList is a list of AwsRedshiftParameterGroup resources
type AwsRedshiftParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftParameterGroup	`json:"items"`
}


// AwsRedshiftParameterGroupParameter is a AwsRedshiftParameterGroupParameter
type AwsRedshiftParameterGroupParameter struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
