
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterParameterGroup describes a AwsRdsClusterParameterGroup resource
type AwsRdsClusterParameterGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRdsClusterParameterGroupSpec	`json:"spec"`
}


// AwsRdsClusterParameterGroupSpec is the spec for a AwsRdsClusterParameterGroup Resource
type AwsRdsClusterParameterGroupSpec struct {
	Parameter	string	`json:"parameter"`
	Tags	map[string]???	`json:"tags"`
	Family	string	`json:"family"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterParameterGroupList is a list of AwsRdsClusterParameterGroup resources
type AwsRdsClusterParameterGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsClusterParameterGroup	`json:"items"`
}

