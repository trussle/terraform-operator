
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRdsClusterParameterGroupSpec	`json:"spec"`
}

type AwsRdsClusterParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsClusterParameterGroup	`json:"items"`
}

type AwsRdsClusterParameterGroupSpec struct {
	Family	string	`json:"family"`
	Description	string	`json:"description"`
	Parameter	interface{}	`json:"parameter"`
	Tags	map[string]interface{}	`json:"tags"`
}
