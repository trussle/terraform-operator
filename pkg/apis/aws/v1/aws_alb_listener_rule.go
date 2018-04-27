
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Action	[]AwsAlbListenerRuleAction	`json:"action"`
	Condition	AwsAlbListenerRuleCondition	`json:"condition"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbListenerRuleList is a list of AwsAlbListenerRule resources
type AwsAlbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbListenerRule	`json:"items"`
}


// AwsAlbListenerRuleAction is a AwsAlbListenerRuleAction
type AwsAlbListenerRuleAction struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	Type	string	`json:"type"`
}

// AwsAlbListenerRuleCondition is a AwsAlbListenerRuleCondition
type AwsAlbListenerRuleCondition struct {
	Field	string	`json:"field"`
	Values	[]string	`json:"values"`
}
