
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAthenaDatabase describes a AwsAthenaDatabase resource
type AwsAthenaDatabase struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAthenaDatabaseSpec	`json:"spec"`
}


// AwsAthenaDatabaseSpec is the spec for a AwsAthenaDatabase Resource
type AwsAthenaDatabaseSpec struct {
	Name	string	`json:"name"`
	Bucket	string	`json:"bucket"`
	ForceDestroy	bool	`json:"force_destroy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAthenaDatabaseList is a list of AwsAthenaDatabase resources
type AwsAthenaDatabaseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAthenaDatabase	`json:"items"`
}

