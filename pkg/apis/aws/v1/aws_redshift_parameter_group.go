
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftParameterGroupSpec	`json:"spec"`
}

type AwsRedshiftParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftParameterGroup	`json:"items"`
}

type AwsRedshiftParameterGroupSpec struct {
	Name	string	`json:"name"`
	Family	string	`json:"family"`
	Description	string	`json:"description"`
	Parameter	interface{}	`json:"parameter"`
}
