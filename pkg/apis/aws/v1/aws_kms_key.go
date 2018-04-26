
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsKey describes a AwsKmsKey resource
type AwsKmsKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKmsKeySpec	`json:"spec"`
}


// AwsKmsKeySpec is the spec for a AwsKmsKey Resource
type AwsKmsKeySpec struct {
	IsEnabled	bool	`json:"is_enabled"`
	EnableKeyRotation	bool	`json:"enable_key_rotation"`
	DeletionWindowInDays	int	`json:"deletion_window_in_days"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsKmsKeyList is a list of AwsKmsKey resources
type AwsKmsKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKmsKey	`json:"items"`
}

