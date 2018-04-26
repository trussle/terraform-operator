
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
	AllocatedCapacity	int	`json:"allocated_capacity"`
	Command	[]TjPcDeMi	`json:"command"`
	Connections	[]string	`json:"connections"`
	DefaultArguments	map[string]???	`json:"default_arguments"`
	Description	string	`json:"description"`
	MaxRetries	int	`json:"max_retries"`
	Name	string	`json:"name"`
	RoleArn	string	`json:"role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueJobList is a list of AwsGlueJob resources
type AwsGlueJobList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueJob	`json:"items"`
}


// TjPcDeMi is a TjPcDeMi
type TjPcDeMi struct {
	Name	string	`json:"name"`
	ScriptLocation	string	`json:"script_location"`
}

// xVduzfhi is a xVduzfhi
type xVduzfhi struct {
	MaxConcurrentRuns	int	`json:"max_concurrent_runs"`
}
