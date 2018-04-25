
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmResourceDataSync describes a AwsSsmResourceDataSync resource
type AwsSsmResourceDataSync struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmResourceDataSyncSpec	`json:"spec"`
}


// AwsSsmResourceDataSyncSpec is the spec for a AwsSsmResourceDataSync Resource
type AwsSsmResourceDataSyncSpec struct {
	Name	string	`json:"name"`
	S3Destination	[]interface{}	`json:"s3_destination"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmResourceDataSyncList is a list of AwsSsmResourceDataSync resources
type AwsSsmResourceDataSyncList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmResourceDataSync	`json:"items"`
}

