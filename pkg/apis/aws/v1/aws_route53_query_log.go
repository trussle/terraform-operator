
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53QueryLog describes a AwsRoute53QueryLog resource
type AwsRoute53QueryLog struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53QueryLogSpec	`json:"spec"`
}


// AwsRoute53QueryLogSpec is the spec for a AwsRoute53QueryLog Resource
type AwsRoute53QueryLogSpec struct {
	CloudwatchLogGroupArn	string	`json:"cloudwatch_log_group_arn"`
	ZoneId	string	`json:"zone_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53QueryLogList is a list of AwsRoute53QueryLog resources
type AwsRoute53QueryLogList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53QueryLog	`json:"items"`
}

