
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotDatafeedSubscription struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSpotDatafeedSubscriptionSpec	`json:"spec"`
}

type AwsSpotDatafeedSubscriptionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotDatafeedSubscription	`json:"items"`
}

type AwsSpotDatafeedSubscriptionSpec struct {
	Bucket	string	`json:"bucket"`
	Prefix	string	`json:"prefix"`
}
