
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbListenerRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbListenerRuleSpec	`json:"spec"`
}

type AwsAlbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListenerRule	`json:"items"`
}

type AwsAlbListenerRuleSpec struct {
	Action	[]interface{}	`json:"action"`
	Condition	interface{}	`json:"condition"`
	ListenerArn	string	`json:"listener_arn"`
}
