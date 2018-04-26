
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	TaskParameters	[]Generic	`json:"task_parameters"`
	WindowId	string	`json:"window_id"`
	MaxConcurrency	string	`json:"max_concurrency"`
	Targets	[]Generic	`json:"targets"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Priority	int	`json:"priority"`
	LoggingInfo	[]Generic	`json:"logging_info"`
	MaxErrors	string	`json:"max_errors"`
	TaskType	string	`json:"task_type"`
	TaskArn	string	`json:"task_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTaskList is a list of AwsSsmMaintenanceWindowTask resources
type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTask	`json:"items"`
}

