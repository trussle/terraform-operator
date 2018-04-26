
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerRule describes a AwsAlbListenerRule resource
type AwsAlbListenerRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbListenerRuleSpec	`json:"spec"`
}


// AwsAlbListenerRuleSpec is the spec for a AwsAlbListenerRule Resource
type AwsAlbListenerRuleSpec struct {
	ListenerArn	string	`json:"listener_arn"`
	Action	[]Generic	`json:"action"`
	Condition	Generic	`json:"condition"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerRuleList is a list of AwsAlbListenerRule resources
type AwsAlbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListenerRule	`json:"items"`
}

