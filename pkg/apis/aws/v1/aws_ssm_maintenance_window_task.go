
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
	Priority	int	`json:"priority"`
	WindowId	string	`json:"window_id"`
	MaxConcurrency	string	`json:"max_concurrency"`
	MaxErrors	string	`json:"max_errors"`
	TaskType	string	`json:"task_type"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Targets	[]interface{}	`json:"targets"`
	TaskArn	string	`json:"task_arn"`
	LoggingInfo	[]interface{}	`json:"logging_info"`
	TaskParameters	[]interface{}	`json:"task_parameters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTaskList is a list of AwsSsmMaintenanceWindowTask resources
type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTask	`json:"items"`
}

