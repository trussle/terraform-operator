
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodecommitTrigger struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodecommitTriggerSpec	`json:"spec"`
}

type AwsCodecommitTriggerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodecommitTrigger	`json:"items"`
}

type AwsCodecommitTriggerSpec struct {
	Trigger	interface{}	`json:"trigger"`
	RepositoryName	string	`json:"repository_name"`
}
