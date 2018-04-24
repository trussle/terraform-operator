
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLambdaEventSourceMapping struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLambdaEventSourceMappingSpec	`json:"spec"`
}

type AwsLambdaEventSourceMappingList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaEventSourceMapping	`json:"items"`
}

type AwsLambdaEventSourceMappingSpec struct {
	EventSourceArn	string	`json:"event_source_arn"`
	FunctionName	string	`json:"function_name"`
	StartingPosition	string	`json:"starting_position"`
	BatchSize	int	`json:"batch_size"`
	Enabled	bool	`json:"enabled"`
}
