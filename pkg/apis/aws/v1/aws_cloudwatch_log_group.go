
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogGroupSpec	`json:"spec"`
}

type AwsCloudwatchLogGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogGroup	`json:"items"`
}

type AwsCloudwatchLogGroupSpec struct {
	KmsKeyId	string	`json:"kms_key_id"`
	Tags	map[string]interface{}	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	RetentionInDays	int	`json:"retention_in_days"`
}
