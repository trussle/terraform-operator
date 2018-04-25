
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRdsDbInstance describes a AwsOpsworksRdsDbInstance resource
type AwsOpsworksRdsDbInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksRdsDbInstanceSpec	`json:"spec"`
}


// AwsOpsworksRdsDbInstanceSpec is the spec for a AwsOpsworksRdsDbInstance Resource
type AwsOpsworksRdsDbInstanceSpec struct {
	StackId	string	`json:"stack_id"`
	RdsDbInstanceArn	string	`json:"rds_db_instance_arn"`
	DbPassword	string	`json:"db_password"`
	DbUser	string	`json:"db_user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksRdsDbInstanceList is a list of AwsOpsworksRdsDbInstance resources
type AwsOpsworksRdsDbInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksRdsDbInstance	`json:"items"`
}

