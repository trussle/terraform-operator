
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogGroup describes a AwsCloudwatchLogGroup resource
type AwsCloudwatchLogGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogGroupSpec	`json:"spec"`
}


// AwsCloudwatchLogGroupSpec is the spec for a AwsCloudwatchLogGroup Resource
type AwsCloudwatchLogGroupSpec struct {
	KmsKeyId	string	`json:"kms_key_id"`
	Tags	map[string]string	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	RetentionInDays	int	`json:"retention_in_days"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogGroupList is a list of AwsCloudwatchLogGroup resources
type AwsCloudwatchLogGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogGroup	`json:"items"`
}

