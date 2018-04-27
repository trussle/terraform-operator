
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
	LoggingInfo	[]AwsSsmMaintenanceWindowTaskLoggingInfo	`json:"logging_info"`
	TaskParameters	[]AwsSsmMaintenanceWindowTaskTaskParameters	`json:"task_parameters"`
	WindowId	string	`json:"window_id"`
	MaxErrors	string	`json:"max_errors"`
	TaskArn	string	`json:"task_arn"`
	Priority	int	`json:"priority"`
	MaxConcurrency	string	`json:"max_concurrency"`
	TaskType	string	`json:"task_type"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Targets	[]AwsSsmMaintenanceWindowTaskTargets	`json:"targets"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTaskList is a list of AwsSsmMaintenanceWindowTask resources
type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTask	`json:"items"`
}


// AwsSsmMaintenanceWindowTaskLoggingInfo is a AwsSsmMaintenanceWindowTaskLoggingInfo
type AwsSsmMaintenanceWindowTaskLoggingInfo struct {
	S3BucketName	string	`json:"s3_bucket_name"`
	S3Region	string	`json:"s3_region"`
	S3BucketPrefix	string	`json:"s3_bucket_prefix"`
}

// AwsSsmMaintenanceWindowTaskTaskParameters is a AwsSsmMaintenanceWindowTaskTaskParameters
type AwsSsmMaintenanceWindowTaskTaskParameters struct {
	Name	string	`json:"name"`
	Values	[]string	`json:"values"`
}

// AwsSsmMaintenanceWindowTaskTargets is a AwsSsmMaintenanceWindowTaskTargets
type AwsSsmMaintenanceWindowTaskTargets struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}
