
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	EventCategories	string	`json:"event_categories"`
	Tags	map[string]string	`json:"tags"`
	Name	string	`json:"name"`
	SnsTopic	string	`json:"sns_topic"`
	SourceIds	string	`json:"source_ids"`
	SourceType	string	`json:"source_type"`
	Enabled	bool	`json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbEventSubscriptionList is a list of AwsDbEventSubscription resources
type AwsDbEventSubscriptionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbEventSubscription	`json:"items"`
}

