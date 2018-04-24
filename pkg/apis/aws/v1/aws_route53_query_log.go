
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53QueryLog struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53QueryLogSpec	`json:"spec"`
}

type AwsRoute53QueryLogList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53QueryLog	`json:"items"`
}

type AwsRoute53QueryLogSpec struct {
	CloudwatchLogGroupArn	string	`json:"cloudwatch_log_group_arn"`
	ZoneId	string	`json:"zone_id"`
}
