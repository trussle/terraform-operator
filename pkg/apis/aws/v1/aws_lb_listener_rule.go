
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerRule describes a AwsLbListenerRule resource
type AwsLbListenerRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbListenerRuleSpec	`json:"spec"`
}


// AwsLbListenerRuleSpec is the spec for a AwsLbListenerRule Resource
type AwsLbListenerRuleSpec struct {
	Condition	string	`json:"condition"`
	ListenerArn	string	`json:"listener_arn"`
	Action	[]interface{}	`json:"action"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerRuleList is a list of AwsLbListenerRule resources
type AwsLbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbListenerRule	`json:"items"`
}

