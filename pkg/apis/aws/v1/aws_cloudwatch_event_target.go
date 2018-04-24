
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchEventTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchEventTargetSpec	`json:"spec"`
}

type AwsCloudwatchEventTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventTarget	`json:"items"`
}

type AwsCloudwatchEventTargetSpec struct {
	Input	string	`json:"input"`
	InputPath	string	`json:"input_path"`
	RoleArn	string	`json:"role_arn"`
	EcsTarget	[]interface{}	`json:"ecs_target"`
	Arn	string	`json:"arn"`
	InputTransformer	[]interface{}	`json:"input_transformer"`
	Rule	string	`json:"rule"`
	RunCommandTargets	[]interface{}	`json:"run_command_targets"`
}
