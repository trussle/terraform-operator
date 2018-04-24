
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbListenerRuleSpec	`json:"spec"`
}

type AwsLbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbListenerRule	`json:"items"`
}

type AwsLbListenerRuleSpec struct {
	Condition	interface{}	`json:"condition"`
	ListenerArn	string	`json:"listener_arn"`
	Action	[]interface{}	`json:"action"`
}
