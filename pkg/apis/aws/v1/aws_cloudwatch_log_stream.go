
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogStream describes a AwsCloudwatchLogStream resource
type AwsCloudwatchLogStream struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogStreamSpec	`json:"spec"`
}


// AwsCloudwatchLogStreamSpec is the spec for a AwsCloudwatchLogStream Resource
type AwsCloudwatchLogStreamSpec struct {
	Name	string	`json:"name"`
	LogGroupName	string	`json:"log_group_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogStreamList is a list of AwsCloudwatchLogStream resources
type AwsCloudwatchLogStreamList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogStream	`json:"items"`
}

