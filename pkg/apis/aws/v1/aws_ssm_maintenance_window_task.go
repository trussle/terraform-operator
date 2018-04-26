
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
	MaxErrors	string	`json:"max_errors"`
	TaskType	string	`json:"task_type"`
	TaskArn	string	`json:"task_arn"`
	Targets	[]vAUkAwww	`json:"targets"`
	WindowId	string	`json:"window_id"`
	ServiceRoleArn	string	`json:"service_role_arn"`
	Priority	int	`json:"priority"`
	LoggingInfo	[]TndUJHiQ	`json:"logging_info"`
	TaskParameters	[]ecbxzvqz	`json:"task_parameters"`
	MaxConcurrency	string	`json:"max_concurrency"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTaskList is a list of AwsSsmMaintenanceWindowTask resources
type AwsSsmMaintenanceWindowTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTask	`json:"items"`
}


// vAUkAwww is a vAUkAwww
type vAUkAwww struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// TndUJHiQ is a TndUJHiQ
type TndUJHiQ struct {
	S3BucketName	string	`json:"s3_bucket_name"`
	S3Region	string	`json:"s3_region"`
	S3BucketPrefix	string	`json:"s3_bucket_prefix"`
}

// ecbxzvqz is a ecbxzvqz
type ecbxzvqz struct {
	Name	string	`json:"name"`
	Values	[]string	`json:"values"`
}
