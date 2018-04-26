
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorResourceGroup describes a AwsInspectorResourceGroup resource
type AwsInspectorResourceGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInspectorResourceGroupSpec	`json:"spec"`
}


// AwsInspectorResourceGroupSpec is the spec for a AwsInspectorResourceGroup Resource
type AwsInspectorResourceGroupSpec struct {
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInspectorResourceGroupList is a list of AwsInspectorResourceGroup resources
type AwsInspectorResourceGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInspectorResourceGroup	`json:"items"`
}

