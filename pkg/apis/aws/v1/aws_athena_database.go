
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAthenaDatabase struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAthenaDatabaseSpec	`json:"spec"`
}

type AwsAthenaDatabaseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAthenaDatabase	`json:"items"`
}

type AwsAthenaDatabaseSpec struct {
	ForceDestroy	bool	`json:"force_destroy"`
	Name	string	`json:"name"`
	Bucket	string	`json:"bucket"`
}
