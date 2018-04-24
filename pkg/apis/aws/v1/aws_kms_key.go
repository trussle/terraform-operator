
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsKmsKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsKmsKeySpec	`json:"spec"`
}

type AwsKmsKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsKmsKey	`json:"items"`
}

type AwsKmsKeySpec struct {
	DeletionWindowInDays	int	`json:"deletion_window_in_days"`
	IsEnabled	bool	`json:"is_enabled"`
	EnableKeyRotation	bool	`json:"enable_key_rotation"`
	Tags	map[string]interface{}	`json:"tags"`
}
