
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyDetector describes a AwsGuarddutyDetector resource
type AwsGuarddutyDetector struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGuarddutyDetectorSpec	`json:"spec"`
}


// AwsGuarddutyDetectorSpec is the spec for a AwsGuarddutyDetector Resource
type AwsGuarddutyDetectorSpec struct {
	Enable	bool	`json:"enable"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyDetectorList is a list of AwsGuarddutyDetector resources
type AwsGuarddutyDetectorList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGuarddutyDetector	`json:"items"`
}

