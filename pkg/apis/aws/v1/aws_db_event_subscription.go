
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbEventSubscription describes a AwsDbEventSubscription resource
type AwsDbEventSubscription struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbEventSubscriptionSpec	`json:"spec"`
}


// AwsDbEventSubscriptionSpec is the spec for a AwsDbEventSubscription Resource
type AwsDbEventSubscriptionSpec struct {
	SnsTopic	string	`json:"sns_topic"`
	SourceType	string	`json:"source_type"`
	Name	string	`json:"name"`
	EventCategories	string	`json:"event_categories"`
	SourceIds	string	`json:"source_ids"`
	Enabled	bool	`json:"enabled"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbEventSubscriptionList is a list of AwsDbEventSubscription resources
type AwsDbEventSubscriptionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbEventSubscription	`json:"items"`
}

