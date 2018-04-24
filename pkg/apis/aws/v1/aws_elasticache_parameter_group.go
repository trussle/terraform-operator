
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheParameterGroupSpec	`json:"spec"`
}

type AwsElasticacheParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheParameterGroup	`json:"items"`
}

type AwsElasticacheParameterGroupSpec struct {
	Family	string	`json:"family"`
	Description	string	`json:"description"`
	Parameter	interface{}	`json:"parameter"`
	Name	string	`json:"name"`
}
