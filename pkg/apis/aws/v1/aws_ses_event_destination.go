
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesEventDestination describes a AwsSesEventDestination resource
type AwsSesEventDestination struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesEventDestinationSpec	`json:"spec"`
}


// AwsSesEventDestinationSpec is the spec for a AwsSesEventDestination Resource
type AwsSesEventDestinationSpec struct {
	Name	string	`json:"name"`
	ConfigurationSetName	string	`json:"configuration_set_name"`
	Enabled	bool	`json:"enabled"`
	MatchingTypes	Generic	`json:"matching_types"`
	CloudwatchDestination	Generic	`json:"cloudwatch_destination"`
	KinesisDestination	Generic	`json:"kinesis_destination"`
	SnsDestination	Generic	`json:"sns_destination"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesEventDestinationList is a list of AwsSesEventDestination resources
type AwsSesEventDestinationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesEventDestination	`json:"items"`
}

