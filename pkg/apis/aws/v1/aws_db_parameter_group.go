
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Parameter	AwsDbParameterGroupParameter	`json:"parameter"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbParameterGroupList is a list of AwsDbParameterGroup resources
type AwsDbParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbParameterGroup	`json:"items"`
}


// AwsDbParameterGroupParameter is a AwsDbParameterGroupParameter
type AwsDbParameterGroupParameter struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
	ApplyMethod	string	`json:"apply_method"`
}
