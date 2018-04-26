
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTask describes a AwsSsmMaintenanceWindowTask resource
type AwsSsmMaintenanceWindowTask struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmMaintenanceWindowTaskSpec	`json:"spec"`
}


// AwsSsmMaintenanceWindowTaskSpec is the spec for a AwsSsmMaintenanceWindowTask Resource
type AwsSsmMaintenanceWindowTaskSpec struct {
	WindowId	string	`json:"window_id"`
	MaxConcurrency	string	`json:"max_concurrency"`
	Priority	int	`json:"priority"`
	LoggingInfo	[]LoggingInfo	`json:"logging_info"`
	TaskParameters	[]TaskParameters	`json:"task_parameters"`
	MaxErrors	string	`json:"max_errors"`
	TaskType	string	`json:"task_type"`
	TaskArn	string	`json:"task_arn"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Targets	[]Targets	`json:"targets"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTaskList is a list of AwsSsmMaintenanceWindowTask resources
type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTask	`json:"items"`
}


// Targets is a Targets
type Targets struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// LoggingInfo is a LoggingInfo
type LoggingInfo struct {
	S3BucketPrefix	string	`json:"s3_bucket_prefix"`
	S3BucketName	string	`json:"s3_bucket_name"`
	S3Region	string	`json:"s3_region"`
}

// TaskParameters is a TaskParameters
type TaskParameters struct {
	Values	[]string	`json:"values"`
	Name	string	`json:"name"`
}
