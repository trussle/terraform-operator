
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesEventDestination struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesEventDestinationSpec	`json:"spec"`
}

type AwsSesEventDestinationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesEventDestination	`json:"items"`
}

type AwsSesEventDestinationSpec struct {
	KinesisDestination	interface{}	`json:"kinesis_destination"`
	SnsDestination	interface{}	`json:"sns_destination"`
	Name	string	`json:"name"`
	ConfigurationSetName	string	`json:"configuration_set_name"`
	Enabled	bool	`json:"enabled"`
	MatchingTypes	interface{}	`json:"matching_types"`
	CloudwatchDestination	interface{}	`json:"cloudwatch_destination"`
}
