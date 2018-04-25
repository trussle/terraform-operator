
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheParameterGroup describes a AwsElasticacheParameterGroup resource
type AwsElasticacheParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheParameterGroupSpec	`json:"spec"`
}


// AwsElasticacheParameterGroupSpec is the spec for a AwsElasticacheParameterGroup Resource
type AwsElasticacheParameterGroupSpec struct {
	Name	string	`json:"name"`
	Family	string	`json:"family"`
	Description	string	`json:"description"`
	Parameter	string	`json:"parameter"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheParameterGroupList is a list of AwsElasticacheParameterGroup resources
type AwsElasticacheParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheParameterGroup	`json:"items"`
}

