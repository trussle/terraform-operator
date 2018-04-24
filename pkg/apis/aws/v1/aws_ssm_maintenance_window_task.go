
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTask struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmMaintenanceWindowTaskSpec	`json:"spec"`
}

type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTask	`json:"items"`
}

type AwsSsmMaintenanceWindowTaskSpec struct {
	TaskType	string	`json:"task_type"`
	TaskArn	string	`json:"task_arn"`
	Priority	int	`json:"priority"`
	LoggingInfo	[]interface{}	`json:"logging_info"`
	TaskParameters	[]interface{}	`json:"task_parameters"`
	WindowId	string	`json:"window_id"`
	MaxConcurrency	string	`json:"max_concurrency"`
	MaxErrors	string	`json:"max_errors"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Targets	[]interface{}	`json:"targets"`
}
