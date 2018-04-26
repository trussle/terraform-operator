
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	RoleArn	string	`json:"role_arn"`
	AllocatedCapacity	int	`json:"allocated_capacity"`
	DefaultArguments	map[string]string	`json:"default_arguments"`
	Name	string	`json:"name"`
	Command	[]Command	`json:"command"`
	Connections	[]string	`json:"connections"`
	Description	string	`json:"description"`
	MaxRetries	int	`json:"max_retries"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueJobList is a list of AwsGlueJob resources
type AwsGlueJobList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueJob	`json:"items"`
}


// ExecutionProperty is a ExecutionProperty
type ExecutionProperty struct {
	MaxConcurrentRuns	int	`json:"max_concurrent_runs"`
}

// Command is a Command
type Command struct {
	Name	string	`json:"name"`
	ScriptLocation	string	`json:"script_location"`
}
