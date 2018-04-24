
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnStateMachine struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSfnStateMachineSpec	`json:"spec"`
}

type AwsSfnStateMachineList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSfnStateMachine	`json:"items"`
}

type AwsSfnStateMachineSpec struct {
	Definition	string	`json:"definition"`
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
}
