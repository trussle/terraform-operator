
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaEventSourceMapping describes a AwsLambdaEventSourceMapping resource
type AwsLambdaEventSourceMapping struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLambdaEventSourceMappingSpec	`json:"spec"`
}


// AwsLambdaEventSourceMappingSpec is the spec for a AwsLambdaEventSourceMapping Resource
type AwsLambdaEventSourceMappingSpec struct {
	FunctionName	string	`json:"function_name"`
	Enabled	bool	`json:"enabled"`
	EventSourceArn	string	`json:"event_source_arn"`
	StartingPosition	string	`json:"starting_position"`
	BatchSize	int	`json:"batch_size"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLambdaEventSourceMappingList is a list of AwsLambdaEventSourceMapping resources
type AwsLambdaEventSourceMappingList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLambdaEventSourceMapping	`json:"items"`
}

