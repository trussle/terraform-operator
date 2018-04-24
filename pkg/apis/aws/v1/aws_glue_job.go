
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueJob struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlueJobSpec	`json:"spec"`
}

type AwsGlueJobList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueJob	`json:"items"`
}

type AwsGlueJobSpec struct {
	Connections	[]interface{}	`json:"connections"`
	Description	string	`json:"description"`
	Name	string	`json:"name"`
	AllocatedCapacity	int	`json:"allocated_capacity"`
	Command	[]interface{}	`json:"command"`
	MaxRetries	int	`json:"max_retries"`
	RoleArn	string	`json:"role_arn"`
	DefaultArguments	map[string]interface{}	`json:"default_arguments"`
}
