
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotDatafeedSubscription describes a AwsSpotDatafeedSubscription resource
type AwsSpotDatafeedSubscription struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSpotDatafeedSubscriptionSpec	`json:"spec"`
}


// AwsSpotDatafeedSubscriptionSpec is the spec for a AwsSpotDatafeedSubscription Resource
type AwsSpotDatafeedSubscriptionSpec struct {
	Prefix	string	`json:"prefix"`
	Bucket	string	`json:"bucket"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotDatafeedSubscriptionList is a list of AwsSpotDatafeedSubscription resources
type AwsSpotDatafeedSubscriptionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotDatafeedSubscription	`json:"items"`
}

