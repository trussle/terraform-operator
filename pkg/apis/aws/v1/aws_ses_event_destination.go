
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	KinesisDestination	AwsSesEventDestinationKinesisDestination	`json:"kinesis_destination"`
	SnsDestination	AwsSesEventDestinationSnsDestination	`json:"sns_destination"`
	Name	string	`json:"name"`
	ConfigurationSetName	string	`json:"configuration_set_name"`
	Enabled	bool	`json:"enabled"`
	MatchingTypes	string	`json:"matching_types"`
	CloudwatchDestination	AwsSesEventDestinationCloudwatchDestination	`json:"cloudwatch_destination"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesEventDestinationList is a list of AwsSesEventDestination resources
type AwsSesEventDestinationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesEventDestination	`json:"items"`
}


// AwsSesEventDestinationKinesisDestination is a AwsSesEventDestinationKinesisDestination
type AwsSesEventDestinationKinesisDestination struct {
	StreamArn	string	`json:"stream_arn"`
	RoleArn	string	`json:"role_arn"`
}

// AwsSesEventDestinationSnsDestination is a AwsSesEventDestinationSnsDestination
type AwsSesEventDestinationSnsDestination struct {
	TopicArn	string	`json:"topic_arn"`
}

// AwsSesEventDestinationCloudwatchDestination is a AwsSesEventDestinationCloudwatchDestination
type AwsSesEventDestinationCloudwatchDestination struct {
	ValueSource	string	`json:"value_source"`
	DefaultValue	string	`json:"default_value"`
	DimensionName	string	`json:"dimension_name"`
}
