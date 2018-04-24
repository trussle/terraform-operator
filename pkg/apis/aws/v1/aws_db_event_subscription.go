
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbEventSubscription struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbEventSubscriptionSpec	`json:"spec"`
}

type AwsDbEventSubscriptionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbEventSubscription	`json:"items"`
}

type AwsDbEventSubscriptionSpec struct {
	Name	string	`json:"name"`
	SourceIds	interface{}	`json:"source_ids"`
	SourceType	string	`json:"source_type"`
	Enabled	bool	`json:"enabled"`
	EventCategories	interface{}	`json:"event_categories"`
	Tags	map[string]interface{}	`json:"tags"`
	SnsTopic	string	`json:"sns_topic"`
}
