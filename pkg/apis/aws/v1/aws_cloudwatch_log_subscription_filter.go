
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogSubscriptionFilter describes a AwsCloudwatchLogSubscriptionFilter resource
type AwsCloudwatchLogSubscriptionFilter struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogSubscriptionFilterSpec	`json:"spec"`
}


// AwsCloudwatchLogSubscriptionFilterSpec is the spec for a AwsCloudwatchLogSubscriptionFilter Resource
type AwsCloudwatchLogSubscriptionFilterSpec struct {
	FilterPattern	string	`json:"filter_pattern"`
	LogGroupName	string	`json:"log_group_name"`
	Distribution	string	`json:"distribution"`
	Name	string	`json:"name"`
	DestinationArn	string	`json:"destination_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogSubscriptionFilterList is a list of AwsCloudwatchLogSubscriptionFilter resources
type AwsCloudwatchLogSubscriptionFilterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogSubscriptionFilter	`json:"items"`
}

