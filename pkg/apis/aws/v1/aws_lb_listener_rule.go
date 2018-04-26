
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
	ListenerArn	string	`json:"listener_arn"`
	Action	[]xdkhCBcg	`json:"action"`
	Condition	string	`json:"condition"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerRuleList is a list of AwsLbListenerRule resources
type AwsLbListenerRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbListenerRule	`json:"items"`
}


// xdkhCBcg is a xdkhCBcg
type xdkhCBcg struct {
	TargetGroupArn	string	`json:"target_group_arn"`
	Type	string	`json:"type"`
}
