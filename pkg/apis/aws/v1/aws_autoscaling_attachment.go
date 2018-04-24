
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAutoscalingAttachmentSpec	`json:"spec"`
}

type AwsAutoscalingAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAutoscalingAttachment	`json:"items"`
}

type AwsAutoscalingAttachmentSpec struct {
	Elb	string	`json:"elb"`
	AlbTargetGroupArn	string	`json:"alb_target_group_arn"`
	AutoscalingGroupName	string	`json:"autoscaling_group_name"`
}
