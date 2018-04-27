
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Action	[]AwsLbListenerRuleAction	`json:"action"`
	Condition	AwsLbListenerRuleCondition	`json:"condition"`
	ListenerArn	string	`json:"listener_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerRuleList is a list of AwsLbListenerRule resources
type AwsLbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbListenerRule	`json:"items"`
}


// AwsLbListenerRuleAction is a AwsLbListenerRuleAction
type AwsLbListenerRuleAction struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	Type	string	`json:"type"`
}

// AwsLbListenerRuleCondition is a AwsLbListenerRuleCondition
type AwsLbListenerRuleCondition struct {
	Values	[]string	`json:"values"`
	Field	string	`json:"field"`
}
