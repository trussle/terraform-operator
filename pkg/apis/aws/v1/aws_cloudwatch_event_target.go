
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventTarget describes a AwsCloudwatchEventTarget resource
type AwsCloudwatchEventTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchEventTargetSpec	`json:"spec"`
}


// AwsCloudwatchEventTargetSpec is the spec for a AwsCloudwatchEventTarget Resource
type AwsCloudwatchEventTargetSpec struct {
	Arn	string	`json:"arn"`
	InputPath	string	`json:"input_path"`
	InputTransformer	[]interface{}	`json:"input_transformer"`
	Rule	string	`json:"rule"`
	Input	string	`json:"input"`
	RoleArn	string	`json:"role_arn"`
	RunCommandTargets	[]interface{}	`json:"run_command_targets"`
	EcsTarget	[]interface{}	`json:"ecs_target"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchEventTargetList is a list of AwsCloudwatchEventTarget resources
type AwsCloudwatchEventTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchEventTarget	`json:"items"`
}

