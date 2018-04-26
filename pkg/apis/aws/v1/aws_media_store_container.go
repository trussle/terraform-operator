
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMediaStoreContainer describes a AwsMediaStoreContainer resource
type AwsMediaStoreContainer struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsMediaStoreContainerSpec	`json:"spec"`
}


// AwsMediaStoreContainerSpec is the spec for a AwsMediaStoreContainer Resource
type AwsMediaStoreContainerSpec struct {
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMediaStoreContainerList is a list of AwsMediaStoreContainer resources
type AwsMediaStoreContainerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMediaStoreContainer	`json:"items"`
}

