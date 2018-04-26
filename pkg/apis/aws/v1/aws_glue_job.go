
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueJob describes a AwsGlueJob resource
type AwsGlueJob struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlueJobSpec	`json:"spec"`
}


// AwsGlueJobSpec is the spec for a AwsGlueJob Resource
type AwsGlueJobSpec struct {
	AllocatedCapacity	int	`json:"allocated_capacity"`
	Command	[]Generic	`json:"command"`
	Description	string	`json:"description"`
	MaxRetries	int	`json:"max_retries"`
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
	Connections	[]Generic	`json:"connections"`
	DefaultArguments	map[string]Generic	`json:"default_arguments"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueJobList is a list of AwsGlueJob resources
type AwsGlueJobList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueJob	`json:"items"`
}

