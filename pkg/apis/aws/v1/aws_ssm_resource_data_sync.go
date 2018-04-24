
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmResourceDataSync struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmResourceDataSyncSpec	`json:"spec"`
}

type AwsSsmResourceDataSyncList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmResourceDataSync	`json:"items"`
}

type AwsSsmResourceDataSyncSpec struct {
	Name	string	`json:"name"`
	S3Destination	[]interface{}	`json:"s3_destination"`
}
