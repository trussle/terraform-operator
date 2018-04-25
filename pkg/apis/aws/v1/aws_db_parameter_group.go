
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbParameterGroup describes a AwsDbParameterGroup resource
type AwsDbParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbParameterGroupSpec	`json:"spec"`
}


// AwsDbParameterGroupSpec is the spec for a AwsDbParameterGroup Resource
type AwsDbParameterGroupSpec struct {
	Family	string	`json:"family"`
	Description	string	`json:"description"`
	Parameter	string	`json:"parameter"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbParameterGroupList is a list of AwsDbParameterGroup resources
type AwsDbParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbParameterGroup	`json:"items"`
}

