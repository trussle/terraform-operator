
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksRdsDbInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksRdsDbInstanceSpec	`json:"spec"`
}

type AwsOpsworksRdsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRdsDbInstance	`json:"items"`
}

type AwsOpsworksRdsDbInstanceSpec struct {
	DbPassword	string	`json:"db_password"`
	DbUser	string	`json:"db_user"`
	StackId	string	`json:"stack_id"`
	RdsDbInstanceArn	string	`json:"rds_db_instance_arn"`
}
